package cases

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/warden-protocol/wardenprotocol/tests/framework"
	"github.com/warden-protocol/wardenprotocol/tests/framework/exec"
	"github.com/warden-protocol/wardenprotocol/warden/x/warden/types"
)

type Test_CreateSpace struct {
	w *exec.WardenNode
}

func (c *Test_CreateSpace) Setup(t *testing.T, ctx context.Context, build framework.BuildResult) {
	c.w = exec.NewWardenNode(t, build.Wardend)

	go c.w.Start(t, ctx, "./testdata/snapshot-base")
	c.w.WaitRunnning(t)
}

func (c *Test_CreateSpace) Run(t *testing.T, ctx context.Context, build framework.BuildResult) {
	alice := exec.NewWardend(c.w, "alice")
	alice.Run(t, "tx warden new-space")

	client := c.w.GRPCClient(t)

	require.Eventually(t, func() bool {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		res, err := client.Warden.Spaces(ctx, &types.QuerySpacesRequest{})
		return err == nil && len(res.Spaces) == 1
	}, 10*time.Second, 10*time.Millisecond)
}
