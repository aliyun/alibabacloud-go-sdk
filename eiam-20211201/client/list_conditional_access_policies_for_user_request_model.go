// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListConditionalAccessPoliciesForUserRequest
	GetInstanceId() *string
	SetUserId(v string) *ListConditionalAccessPoliciesForUserRequest
	GetUserId() *string
}

type ListConditionalAccessPoliciesForUserRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 条件访问策略关联的用户ID
	//
	// This parameter is required.
	//
	// example:
	//
	// user_xxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListConditionalAccessPoliciesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListConditionalAccessPoliciesForUserRequest) SetInstanceId(v string) *ListConditionalAccessPoliciesForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserRequest) SetUserId(v string) *ListConditionalAccessPoliciesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserRequest) Validate() error {
	return dara.Validate(s)
}
