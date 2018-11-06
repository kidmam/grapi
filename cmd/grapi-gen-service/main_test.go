package main

import (
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/izumin5210/grapi/pkg/cli"
	"github.com/izumin5210/grapi/pkg/gencmd"
	gencmdtesting "github.com/izumin5210/grapi/pkg/gencmd/testing"
	"github.com/izumin5210/grapi/pkg/grapicmd"
	"github.com/izumin5210/grapi/pkg/protoc"
	"github.com/izumin5210/grapi/pkg/svcgen"
	svcgentesting "github.com/izumin5210/grapi/pkg/svcgen/testing"
)

func TestRun(t *testing.T) {
	cases := []svcgentesting.Case{
		{
			Test: "simple",
			Args: []string{"foo"},
			Files: []string{
				"api/protos/foo.proto",
				"app/server/foo_server.go",
				"app/server/foo_server_register_funcs.go",
				"app/server/foo_server_test.go",
			},
		},
		{
			Test: "specify package",
			Args: []string{"foo"},
			Files: []string{
				"api/protos/foo.proto",
				"app/server/foo_server.go",
				"app/server/foo_server_register_funcs.go",
				"app/server/foo_server_test.go",
			},
			PkgName: "testcompany.testapp",
		},
		{
			Test: "nested",
			Args: []string{"foo/bar"},
			Files: []string{
				"api/protos/foo/bar.proto",
				"app/server/foo/bar_server.go",
				"app/server/foo/bar_server_register_funcs.go",
				"app/server/foo/bar_server_test.go",
			},
		},
		{
			Test: "nested with specify pacakge",
			Args: []string{"foo/bar"},
			Files: []string{
				"api/protos/foo/bar.proto",
				"app/server/foo/bar_server.go",
				"app/server/foo/bar_server_register_funcs.go",
				"app/server/foo/bar_server_test.go",
			},
			PkgName: "testcompany.testapp",
		},
		{
			Test: "snake_case name",
			Args: []string{"foo/bar_baz"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test: "kebab-case name",
			Args: []string{"foo/bar-baz"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test: "with some standard methods",
			Args: []string{"foo/bar-baz", "list", "create", "delete"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test: "with non-standard methods",
			Args: []string{"foo/bar-baz", "list", "create", "rename", "delete", "move_move"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test: "specify proto dir",
			Args: []string{"qux"},
			Files: []string{
				"pkg/foo/protos/qux.proto",
				"app/server/qux_server.go",
				"app/server/qux_server_register_funcs.go",
				"app/server/qux_server_test.go",
			},
			ProtoDir: "pkg/foo/protos",
		},
		{
			Test: "specify proto out dir",
			Args: []string{"quux"},
			Files: []string{
				"api/protos/quux.proto",
				"app/server/quux_server.go",
				"app/server/quux_server_register_funcs.go",
				"app/server/quux_server_test.go",
			},
			ProtoOutDir: "api/out",
		},
		{
			Test: "specify server dir",
			Args: []string{"corge"},
			Files: []string{
				"api/protos/corge.proto",
				"pkg/foo/server/corge_server.go",
				"pkg/foo/server/corge_server_register_funcs.go",
				"pkg/foo/server/corge_server_test.go",
			},
			ServerDir: "pkg/foo/server",
		},
		{
			Test: "skip tests",
			Args: []string{"--skip-test", "book"},
			Files: []string{
				"api/protos/book.proto",
				"app/server/book_server.go",
				"app/server/book_server_register_funcs.go",
			},
			SkippedFiles: map[string]struct{}{
				"app/server/book_server_test.go": {},
			},
		},
		{
			Test: "specify resource name",
			Args: []string{"library", "--resource-name=book"},
			Files: []string{
				"api/protos/library.proto",
				"app/server/library_server.go",
				"app/server/library_server_register_funcs.go",
				"app/server/library_server_test.go",
			},
		},
	}

	rootDir := cli.RootDir("/home/src/testapp")

	createSvcApp := func(ctx *gencmd.Ctx, cmd *gencmd.Command) (*svcgen.App, error) {
		return svcgentesting.NewTestApp(ctx, cmd, &fakeProtocWrapper{}, cli.NopUI)
	}
	createGenApp := func(ctx *gencmd.Ctx, cmd *gencmd.Command) (*gencmd.App, error) {
		return gencmdtesting.NewTestApp(ctx, cmd, cli.NopUI)
	}
	createCmd := func(t *testing.T, fs afero.Fs, tc svcgentesting.Case) *cobra.Command {
		ctx := &grapicmd.Ctx{
			FS:      fs,
			RootDir: rootDir,
			Config: grapicmd.Config{
				Package: tc.PkgName,
			},
			ProtocConfig: protoc.Config{
				ProtosDir: tc.ProtoDir,
				OutDir:    tc.ProtoOutDir,
			},
		}
		ctx.Config.Grapi.ServerDir = tc.ServerDir
		return gencmd.NewCommand("service", &gencmd.Ctx{
			Ctx:           ctx,
			CreateAppFunc: createGenApp,
			GenerateCmd:   NewGenerateCommand(createSvcApp),
		})
	}

	ctx := &svcgentesting.Ctx{
		GOPATH:    "/home",
		RootDir:   rootDir,
		CreateCmd: createCmd,
		Cases:     cases,
	}

	svcgentesting.Run(t, ctx)
}

type fakeProtocWrapper struct{}

func (*fakeProtocWrapper) Exec(context.Context) error { return nil }
