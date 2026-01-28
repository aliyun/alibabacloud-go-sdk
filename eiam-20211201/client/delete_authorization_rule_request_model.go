// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *DeleteAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetInstanceId(v string) *DeleteAuthorizationRuleRequest
	GetInstanceId() *string
}

type DeleteAuthorizationRuleRequest struct {
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *DeleteAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *DeleteAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *DeleteAuthorizationRuleRequest) SetInstanceId(v string) *DeleteAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
