// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateUserAttributeRequest
	GetClusterId() *string
	SetDescription(v string) *UpdateUserAttributeRequest
	GetDescription() *string
	SetPassword(v string) *UpdateUserAttributeRequest
	GetPassword() *string
	SetRegionId(v string) *UpdateUserAttributeRequest
	GetRegionId() *string
	SetUserId(v string) *UpdateUserAttributeRequest
	GetUserId() *string
	SetUserName(v string) *UpdateUserAttributeRequest
	GetUserName() *string
}

type UpdateUserAttributeRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The remarks of the user.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user password.
	//
	// example:
	//
	// 1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// The user ID.
	//
	// example:
	//
	// 125046002175****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// yun****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserAttributeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateUserAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateUserAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateUserAttributeRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserAttributeRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateUserAttributeRequest) SetClusterId(v string) *UpdateUserAttributeRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateUserAttributeRequest) SetDescription(v string) *UpdateUserAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserAttributeRequest) SetPassword(v string) *UpdateUserAttributeRequest {
	s.Password = &v
	return s
}

func (s *UpdateUserAttributeRequest) SetRegionId(v string) *UpdateUserAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateUserAttributeRequest) SetUserId(v string) *UpdateUserAttributeRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserAttributeRequest) SetUserName(v string) *UpdateUserAttributeRequest {
	s.UserName = &v
	return s
}

func (s *UpdateUserAttributeRequest) Validate() error {
	return dara.Validate(s)
}
