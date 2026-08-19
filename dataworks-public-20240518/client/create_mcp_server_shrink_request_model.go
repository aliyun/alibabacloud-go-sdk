// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigShrink(v string) *CreateMcpServerShrinkRequest
	GetConfigShrink() *string
	SetName(v string) *CreateMcpServerShrinkRequest
	GetName() *string
	SetVisibility(v string) *CreateMcpServerShrinkRequest
	GetVisibility() *string
	SetVisibilityScopeShrink(v string) *CreateMcpServerShrinkRequest
	GetVisibilityScopeShrink() *string
}

type CreateMcpServerShrinkRequest struct {
	// The connection configuration of the MCP Server.
	//
	// example:
	//
	// -
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the MCP Server. The name must be unique at the tenant level. It must start with a lowercase letter and can contain only `a-z`, `0-9`, `_`, and `-`.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The visibility level.
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The visibility scope. The corresponding field is used based on the Visibility value.
	VisibilityScopeShrink *string `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty"`
}

func (s CreateMcpServerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpServerShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *CreateMcpServerShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMcpServerShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateMcpServerShrinkRequest) GetVisibilityScopeShrink() *string {
	return s.VisibilityScopeShrink
}

func (s *CreateMcpServerShrinkRequest) SetConfigShrink(v string) *CreateMcpServerShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *CreateMcpServerShrinkRequest) SetName(v string) *CreateMcpServerShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMcpServerShrinkRequest) SetVisibility(v string) *CreateMcpServerShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *CreateMcpServerShrinkRequest) SetVisibilityScopeShrink(v string) *CreateMcpServerShrinkRequest {
	s.VisibilityScopeShrink = &v
	return s
}

func (s *CreateMcpServerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
