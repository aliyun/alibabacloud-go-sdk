// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomHeaders(v map[string]interface{}) *UpdateMcpServerRequest
	GetCustomHeaders() map[string]interface{}
	SetName(v string) *UpdateMcpServerRequest
	GetName() *string
	SetTransport(v string) *UpdateMcpServerRequest
	GetTransport() *string
	SetUrl(v string) *UpdateMcpServerRequest
	GetUrl() *string
	SetVisibility(v string) *UpdateMcpServerRequest
	GetVisibility() *string
	SetVisibilityScope(v *UpdateMcpServerRequestVisibilityScope) *UpdateMcpServerRequest
	GetVisibilityScope() *UpdateMcpServerRequestVisibilityScope
}

type UpdateMcpServerRequest struct {
	// The new custom request headers, specified as key-value pairs.
	//
	// example:
	//
	// {}
	CustomHeaders map[string]interface{} `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// The name of the MCP Server to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new transport protocol.
	//
	// example:
	//
	// SSE
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	// The new service address. The address must start with`https://`.
	//
	// example:
	//
	// https://example.com/mcp/sse
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The new visibility level.
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The new visibility scope. The fields in this object depend on the value of the `Visibility` parameter.
	VisibilityScope *UpdateMcpServerRequestVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s UpdateMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerRequest) GetCustomHeaders() map[string]interface{} {
	return s.CustomHeaders
}

func (s *UpdateMcpServerRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMcpServerRequest) GetTransport() *string {
	return s.Transport
}

func (s *UpdateMcpServerRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateMcpServerRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateMcpServerRequest) GetVisibilityScope() *UpdateMcpServerRequestVisibilityScope {
	return s.VisibilityScope
}

func (s *UpdateMcpServerRequest) SetCustomHeaders(v map[string]interface{}) *UpdateMcpServerRequest {
	s.CustomHeaders = v
	return s
}

func (s *UpdateMcpServerRequest) SetName(v string) *UpdateMcpServerRequest {
	s.Name = &v
	return s
}

func (s *UpdateMcpServerRequest) SetTransport(v string) *UpdateMcpServerRequest {
	s.Transport = &v
	return s
}

func (s *UpdateMcpServerRequest) SetUrl(v string) *UpdateMcpServerRequest {
	s.Url = &v
	return s
}

func (s *UpdateMcpServerRequest) SetVisibility(v string) *UpdateMcpServerRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateMcpServerRequest) SetVisibilityScope(v *UpdateMcpServerRequestVisibilityScope) *UpdateMcpServerRequest {
	s.VisibilityScope = v
	return s
}

func (s *UpdateMcpServerRequest) Validate() error {
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpServerRequestVisibilityScope struct {
	// The list of workspace IDs that can access the MCP Server. This parameter takes effect only when `Visibility` is set to `PROJECT`.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of user IDs that can access the MCP Server. This parameter takes effect only when `Visibility` is set to `USER`.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateMcpServerRequestVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerRequestVisibilityScope) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerRequestVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *UpdateMcpServerRequestVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateMcpServerRequestVisibilityScope) SetProjectIds(v []*string) *UpdateMcpServerRequestVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *UpdateMcpServerRequestVisibilityScope) SetUserIds(v []*string) *UpdateMcpServerRequestVisibilityScope {
	s.UserIds = v
	return s
}

func (s *UpdateMcpServerRequestVisibilityScope) Validate() error {
	return dara.Validate(s)
}
