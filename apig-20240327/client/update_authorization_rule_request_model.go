// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResources(v []*string) *UpdateAuthorizationRuleRequest
	GetResources() []*string
}

type UpdateAuthorizationRuleRequest struct {
	// The new list of authorized resources. In MCP Tool scenarios, this is a set of tool names. An empty array returns ErrInvalidParameterWithEmpty.
	Resources []*string `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s UpdateAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleRequest) GetResources() []*string {
	return s.Resources
}

func (s *UpdateAuthorizationRuleRequest) SetResources(v []*string) *UpdateAuthorizationRuleRequest {
	s.Resources = v
	return s
}

func (s *UpdateAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
