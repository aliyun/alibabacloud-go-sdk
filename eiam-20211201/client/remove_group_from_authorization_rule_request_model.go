// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupFromAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *RemoveGroupFromAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetGroupId(v string) *RemoveGroupFromAuthorizationRuleRequest
	GetGroupId() *string
	SetInstanceId(v string) *RemoveGroupFromAuthorizationRuleRequest
	GetInstanceId() *string
}

type RemoveGroupFromAuthorizationRuleRequest struct {
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 组ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RemoveGroupFromAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupFromAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupFromAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *RemoveGroupFromAuthorizationRuleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveGroupFromAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveGroupFromAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *RemoveGroupFromAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *RemoveGroupFromAuthorizationRuleRequest) SetGroupId(v string) *RemoveGroupFromAuthorizationRuleRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveGroupFromAuthorizationRuleRequest) SetInstanceId(v string) *RemoveGroupFromAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveGroupFromAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
