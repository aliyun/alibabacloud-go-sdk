// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *CreateMcpServerRequestConfig) *CreateMcpServerRequest
	GetConfig() *CreateMcpServerRequestConfig
	SetName(v string) *CreateMcpServerRequest
	GetName() *string
	SetVisibility(v string) *CreateMcpServerRequest
	GetVisibility() *string
	SetVisibilityScope(v *CreateMcpServerRequestVisibilityScope) *CreateMcpServerRequest
	GetVisibilityScope() *CreateMcpServerRequestVisibilityScope
}

type CreateMcpServerRequest struct {
	// example:
	//
	// -
	Config *CreateMcpServerRequestConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TENANT
	Visibility      *string                                `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VisibilityScope *CreateMcpServerRequestVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s CreateMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequest) GetConfig() *CreateMcpServerRequestConfig {
	return s.Config
}

func (s *CreateMcpServerRequest) GetName() *string {
	return s.Name
}

func (s *CreateMcpServerRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateMcpServerRequest) GetVisibilityScope() *CreateMcpServerRequestVisibilityScope {
	return s.VisibilityScope
}

func (s *CreateMcpServerRequest) SetConfig(v *CreateMcpServerRequestConfig) *CreateMcpServerRequest {
	s.Config = v
	return s
}

func (s *CreateMcpServerRequest) SetName(v string) *CreateMcpServerRequest {
	s.Name = &v
	return s
}

func (s *CreateMcpServerRequest) SetVisibility(v string) *CreateMcpServerRequest {
	s.Visibility = &v
	return s
}

func (s *CreateMcpServerRequest) SetVisibilityScope(v *CreateMcpServerRequestVisibilityScope) *CreateMcpServerRequest {
	s.VisibilityScope = v
	return s
}

func (s *CreateMcpServerRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpServerRequestConfig struct {
	// example:
	//
	// {}
	CustomHeaders map[string]interface{} `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// example:
	//
	// SSE
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	// example:
	//
	// https://example.com/mcp/sse
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateMcpServerRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestConfig) GetCustomHeaders() map[string]interface{} {
	return s.CustomHeaders
}

func (s *CreateMcpServerRequestConfig) GetTransport() *string {
	return s.Transport
}

func (s *CreateMcpServerRequestConfig) GetUrl() *string {
	return s.Url
}

func (s *CreateMcpServerRequestConfig) SetCustomHeaders(v map[string]interface{}) *CreateMcpServerRequestConfig {
	s.CustomHeaders = v
	return s
}

func (s *CreateMcpServerRequestConfig) SetTransport(v string) *CreateMcpServerRequestConfig {
	s.Transport = &v
	return s
}

func (s *CreateMcpServerRequestConfig) SetUrl(v string) *CreateMcpServerRequestConfig {
	s.Url = &v
	return s
}

func (s *CreateMcpServerRequestConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMcpServerRequestVisibilityScope struct {
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	UserIds    []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateMcpServerRequestVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestVisibilityScope) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *CreateMcpServerRequestVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateMcpServerRequestVisibilityScope) SetProjectIds(v []*string) *CreateMcpServerRequestVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *CreateMcpServerRequestVisibilityScope) SetUserIds(v []*string) *CreateMcpServerRequestVisibilityScope {
	s.UserIds = v
	return s
}

func (s *CreateMcpServerRequestVisibilityScope) Validate() error {
	return dara.Validate(s)
}
