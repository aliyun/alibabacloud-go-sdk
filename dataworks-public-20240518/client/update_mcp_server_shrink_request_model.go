// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomHeadersShrink(v string) *UpdateMcpServerShrinkRequest
	GetCustomHeadersShrink() *string
	SetName(v string) *UpdateMcpServerShrinkRequest
	GetName() *string
	SetTransport(v string) *UpdateMcpServerShrinkRequest
	GetTransport() *string
	SetUrl(v string) *UpdateMcpServerShrinkRequest
	GetUrl() *string
	SetVisibility(v string) *UpdateMcpServerShrinkRequest
	GetVisibility() *string
	SetVisibilityScopeShrink(v string) *UpdateMcpServerShrinkRequest
	GetVisibilityScopeShrink() *string
}

type UpdateMcpServerShrinkRequest struct {
	// example:
	//
	// {}
	CustomHeadersShrink *string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SSE
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	// example:
	//
	// https://example.com/mcp/sse
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// TENANT
	Visibility            *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VisibilityScopeShrink *string `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty"`
}

func (s UpdateMcpServerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerShrinkRequest) GetCustomHeadersShrink() *string {
	return s.CustomHeadersShrink
}

func (s *UpdateMcpServerShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMcpServerShrinkRequest) GetTransport() *string {
	return s.Transport
}

func (s *UpdateMcpServerShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateMcpServerShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateMcpServerShrinkRequest) GetVisibilityScopeShrink() *string {
	return s.VisibilityScopeShrink
}

func (s *UpdateMcpServerShrinkRequest) SetCustomHeadersShrink(v string) *UpdateMcpServerShrinkRequest {
	s.CustomHeadersShrink = &v
	return s
}

func (s *UpdateMcpServerShrinkRequest) SetName(v string) *UpdateMcpServerShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMcpServerShrinkRequest) SetTransport(v string) *UpdateMcpServerShrinkRequest {
	s.Transport = &v
	return s
}

func (s *UpdateMcpServerShrinkRequest) SetUrl(v string) *UpdateMcpServerShrinkRequest {
	s.Url = &v
	return s
}

func (s *UpdateMcpServerShrinkRequest) SetVisibility(v string) *UpdateMcpServerShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateMcpServerShrinkRequest) SetVisibilityScopeShrink(v string) *UpdateMcpServerShrinkRequest {
	s.VisibilityScopeShrink = &v
	return s
}

func (s *UpdateMcpServerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
