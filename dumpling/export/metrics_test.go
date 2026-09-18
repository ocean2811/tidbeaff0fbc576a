// Copyright 2020 PingCAP, Inc. Licensed under Apache-2.0.

package export

import (
	"testing"

	"github.com/ocean2811/tidbeaff0fbc576a/pkg/util/promutil"
)

func TestMetricsRegistration(t *testing.T) {
	m := newMetrics(promutil.NewDefaultFactory(), nil)
	registry := promutil.NewDefaultRegistry()
	m.registerTo(registry)
	m.unregisterFrom(registry)
}
