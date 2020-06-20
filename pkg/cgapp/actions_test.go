package cgapp

import (
	"os"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCreateCLIAction(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success action",
			args{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateCLIAction(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateCLIAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	files := []string{".editorconfig", ".gitignore", "Makefile", "docker-compose.yml", "docker-compose.prod.yml"}
	for _, name := range files {
		os.Remove(name)
	}
}