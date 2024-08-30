package main

import (
	"context"
	"dagger/ci/internal/dagger"
	"fmt"
)

type Ci struct{}

// Dispatch function for pocketci
func (m *Ci) Dispatch(ctx context.Context, vendor, event, filter string, src *dagger.Directory, eventTrigger *dagger.File) error {
	fmt.Println(eventTrigger.Contents(ctx))

	ci := dag.Pocketci(eventTrigger)
	if ci.OnPullRequest() != nil {
		fmt.Println("Received On Pull Request")
	}
	if ci.OnCommitPush() != nil {
		fmt.Println("Receinved OnCommitPush")
	}

	return nil
}

func (m *Ci) OnPullRequest(ctx context.Context) {
}

func (m *Ci) OnCommitPush(ctx context.Context, src *dagger.Directory, eventTrigger *dagger.File) {
}
