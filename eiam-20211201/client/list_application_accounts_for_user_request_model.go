// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccountsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListApplicationAccountsForUserRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListApplicationAccountsForUserRequest
	GetInstanceId() *string
	SetUserId(v string) *ListApplicationAccountsForUserRequest
	GetUserId() *string
}

type ListApplicationAccountsForUserRequest struct {
	// IDaaS的应用主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// app_11111
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用户Id
	//
	// This parameter is required.
	//
	// example:
	//
	// user_example
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListApplicationAccountsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsForUserRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationAccountsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationAccountsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListApplicationAccountsForUserRequest) SetApplicationId(v string) *ListApplicationAccountsForUserRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationAccountsForUserRequest) SetInstanceId(v string) *ListApplicationAccountsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationAccountsForUserRequest) SetUserId(v string) *ListApplicationAccountsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListApplicationAccountsForUserRequest) Validate() error {
	return dara.Validate(s)
}
