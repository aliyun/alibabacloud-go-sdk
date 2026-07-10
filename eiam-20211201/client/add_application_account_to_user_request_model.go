// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApplicationAccountToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AddApplicationAccountToUserRequest
	GetApplicationId() *string
	SetApplicationUsername(v string) *AddApplicationAccountToUserRequest
	GetApplicationUsername() *string
	SetInstanceId(v string) *AddApplicationAccountToUserRequest
	GetInstanceId() *string
	SetUserId(v string) *AddApplicationAccountToUserRequest
	GetUserId() *string
}

type AddApplicationAccountToUserRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_11111
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application account name.
	//
	// This parameter is required.
	//
	// example:
	//
	// zhangsan
	ApplicationUsername *string `json:"ApplicationUsername,omitempty" xml:"ApplicationUsername,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_example
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddApplicationAccountToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddApplicationAccountToUserRequest) GoString() string {
	return s.String()
}

func (s *AddApplicationAccountToUserRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AddApplicationAccountToUserRequest) GetApplicationUsername() *string {
	return s.ApplicationUsername
}

func (s *AddApplicationAccountToUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddApplicationAccountToUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddApplicationAccountToUserRequest) SetApplicationId(v string) *AddApplicationAccountToUserRequest {
	s.ApplicationId = &v
	return s
}

func (s *AddApplicationAccountToUserRequest) SetApplicationUsername(v string) *AddApplicationAccountToUserRequest {
	s.ApplicationUsername = &v
	return s
}

func (s *AddApplicationAccountToUserRequest) SetInstanceId(v string) *AddApplicationAccountToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AddApplicationAccountToUserRequest) SetUserId(v string) *AddApplicationAccountToUserRequest {
	s.UserId = &v
	return s
}

func (s *AddApplicationAccountToUserRequest) Validate() error {
	return dara.Validate(s)
}
