// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeletePermissionRuleRequest
	GetClientToken() *string
}

type DeletePermissionRuleRequest struct {
	// The client token that ensures the idempotence of the request.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeletePermissionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePermissionRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeletePermissionRuleRequest) SetClientToken(v string) *DeletePermissionRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletePermissionRuleRequest) Validate() error {
	return dara.Validate(s)
}
