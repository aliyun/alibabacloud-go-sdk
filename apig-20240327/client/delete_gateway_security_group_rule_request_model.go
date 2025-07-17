// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCascadingDelete(v bool) *DeleteGatewaySecurityGroupRuleRequest
	GetCascadingDelete() *bool
}

type DeleteGatewaySecurityGroupRuleRequest struct {
	// Whether to cascade delete the security group rules.
	//
	// example:
	//
	// true
	CascadingDelete *bool `json:"cascadingDelete,omitempty" xml:"cascadingDelete,omitempty"`
}

func (s DeleteGatewaySecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecurityGroupRuleRequest) GetCascadingDelete() *bool {
	return s.CascadingDelete
}

func (s *DeleteGatewaySecurityGroupRuleRequest) SetCascadingDelete(v bool) *DeleteGatewaySecurityGroupRuleRequest {
	s.CascadingDelete = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
