// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *GetAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetInstanceId(v string) *GetAuthorizationRuleRequest
	GetInstanceId() *string
}

type GetAuthorizationRuleRequest struct {
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

func (s GetAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *GetAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *GetAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *GetAuthorizationRuleRequest) SetInstanceId(v string) *GetAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
