package main

import (
	"context"
	"dagger/ci/internal/dagger"
	"fmt"
)

type Ci struct{}

// Dispatch function for pocketci
func (m *Ci) Dispatch(ctx context.Context, src *dagger.Directory, eventTrigger *dagger.File) error {
	fmt.Println(eventTrigger.Contents(ctx))
	return nil
}
