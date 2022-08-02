package module

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProtoPackageName(t *testing.T) {
	cases := []struct {
		name   string
		app    string
		module string
		want   string
	}{
		{
			name:   "name",
			app:    "baseclass",
			module: "test",
			want:   "baseclass.test",
		},
		{
			name:   "path",
			app:    "baseclass/cli",
			module: "test",
			want:   "baseclass.cli.test",
		},
		{
			name:   "path with dash",
			app:    "baseclass/c-li",
			module: "test",
			want:   "baseclass.cli.test",
		},
		{
			name:   "path with number prefix",
			app:    "0ignite/cli",
			module: "test",
			want:   "_0ignite.cli.test",
		},
		{
			name:   "path with number prefix and dash",
			app:    "0ignite/cli",
			module: "test",
			want:   "_0ignite.cli.test",
		},
		{
			name:   "module with dash",
			app:    "baseclass",
			module: "test-mod",
			want:   "baseclass.testmod",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, ProtoPackageName(tt.app, tt.module))
		})
	}
}
