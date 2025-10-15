// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationAccountFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccountId(v string) *RemoveApplicationAccountFromUserRequest
	GetApplicationAccountId() *string
	SetApplicationId(v string) *RemoveApplicationAccountFromUserRequest
	GetApplicationId() *string
	SetInstanceId(v string) *RemoveApplicationAccountFromUserRequest
	GetInstanceId() *string
	SetUserId(v string) *RemoveApplicationAccountFromUserRequest
	GetUserId() *string
}

type RemoveApplicationAccountFromUserRequest struct {
	// 应用账号Id
	//
	// This parameter is required.
	//
	// example:
	//
	// act_example
	ApplicationAccountId *string `json:"ApplicationAccountId,omitempty" xml:"ApplicationAccountId,omitempty"`
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

func (s RemoveApplicationAccountFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationAccountFromUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveApplicationAccountFromUserRequest) GetApplicationAccountId() *string {
	return s.ApplicationAccountId
}

func (s *RemoveApplicationAccountFromUserRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemoveApplicationAccountFromUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveApplicationAccountFromUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemoveApplicationAccountFromUserRequest) SetApplicationAccountId(v string) *RemoveApplicationAccountFromUserRequest {
	s.ApplicationAccountId = &v
	return s
}

func (s *RemoveApplicationAccountFromUserRequest) SetApplicationId(v string) *RemoveApplicationAccountFromUserRequest {
	s.ApplicationId = &v
	return s
}

func (s *RemoveApplicationAccountFromUserRequest) SetInstanceId(v string) *RemoveApplicationAccountFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveApplicationAccountFromUserRequest) SetUserId(v string) *RemoveApplicationAccountFromUserRequest {
	s.UserId = &v
	return s
}

func (s *RemoveApplicationAccountFromUserRequest) Validate() error {
	return dara.Validate(s)
}
