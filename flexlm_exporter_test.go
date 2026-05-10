package main

import (
	"log/slog"
	"os"
	"testing"
)

func TestNewHandler(t *testing.T) {
	t.Parallel()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	configPath := "config/fixtures/licenses.yml"

	testCases := []struct {
		name                   string
		includeExporterMetrics bool
		maxRequests            int
	}{
		{
			name:                   "WithExporterMetrics",
			includeExporterMetrics: true,
			maxRequests:            40,
		},
		{
			name:                   "WithoutExporterMetrics",
			includeExporterMetrics: false,
			maxRequests:            0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			h := newHandler(tc.includeExporterMetrics, configPath, tc.maxRequests, logger)

			if h == nil {
				t.Fatal("expected handler to not be nil")
			}

			if h.includeExporterMetrics != tc.includeExporterMetrics {
				t.Errorf("expected includeExporterMetrics to be %v, got %v", tc.includeExporterMetrics, h.includeExporterMetrics)
			}

			if h.configPath != configPath {
				t.Errorf("expected configPath to be %s, got %s", configPath, h.configPath)
			}

			if h.maxRequests != tc.maxRequests {
				t.Errorf("expected maxRequests to be %d, got %d", tc.maxRequests, h.maxRequests)
			}

			if h.logger != logger {
				t.Error("expected logger to match the provided logger")
			}

			if h.exporterMetricsRegistry == nil {
				t.Error("expected exporterMetricsRegistry to be initialized")
			}

			if h.unfilteredHandler == nil {
				t.Error("expected unfilteredHandler to be initialized")
			}
		})
	}
}
