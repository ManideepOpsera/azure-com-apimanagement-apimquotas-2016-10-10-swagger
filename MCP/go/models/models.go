package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// QuotaCounterCollection represents the QuotaCounterCollection schema from the OpenAPI specification
type QuotaCounterCollection struct {
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []QuotaCounterContract `json:"value,omitempty"` // Quota counter values.
}

// QuotaCounterContract represents the QuotaCounterContract schema from the OpenAPI specification
type QuotaCounterContract struct {
	Kbtransferred float64 `json:"kbTransferred,omitempty"` // Data Transferred in KiloBytes.
	Callscount int `json:"callsCount,omitempty"` // Number of times Counter was called.
}

// QuotaCounterValueContract represents the QuotaCounterValueContract schema from the OpenAPI specification
type QuotaCounterValueContract struct {
	Kbtransferred float64 `json:"kbTransferred,omitempty"` // Data Transferred in KiloBytes.
	Callscount int `json:"callsCount,omitempty"` // Number of times Counter was called.
}
