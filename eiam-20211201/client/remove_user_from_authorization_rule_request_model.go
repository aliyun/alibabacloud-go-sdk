// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *RemoveUserFromAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetInstanceId(v string) *RemoveUserFromAuthorizationRuleRequest
	GetInstanceId() *string
	SetUserId(v string) *RemoveUserFromAuthorizationRuleRequest
	GetUserId() *string
}

type RemoveUserFromAuthorizationRuleRequest struct {
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
	// 账户ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUserFromAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *RemoveUserFromAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveUserFromAuthorizationRuleRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemoveUserFromAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *RemoveUserFromAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *RemoveUserFromAuthorizationRuleRequest) SetInstanceId(v string) *RemoveUserFromAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUserFromAuthorizationRuleRequest) SetUserId(v string) *RemoveUserFromAuthorizationRuleRequest {
	s.UserId = &v
	return s
}

func (s *RemoveUserFromAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
