// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationFromAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemoveApplicationFromAuthorizationRuleRequest
	GetApplicationId() *string
	SetAuthorizationRuleId(v string) *RemoveApplicationFromAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetInstanceId(v string) *RemoveApplicationFromAuthorizationRuleRequest
	GetInstanceId() *string
}

type RemoveApplicationFromAuthorizationRuleRequest struct {
	// 应用 ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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

func (s RemoveApplicationFromAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationFromAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *RemoveApplicationFromAuthorizationRuleRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemoveApplicationFromAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *RemoveApplicationFromAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveApplicationFromAuthorizationRuleRequest) SetApplicationId(v string) *RemoveApplicationFromAuthorizationRuleRequest {
	s.ApplicationId = &v
	return s
}

func (s *RemoveApplicationFromAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *RemoveApplicationFromAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *RemoveApplicationFromAuthorizationRuleRequest) SetInstanceId(v string) *RemoveApplicationFromAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveApplicationFromAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
