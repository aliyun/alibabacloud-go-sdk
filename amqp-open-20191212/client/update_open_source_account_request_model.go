// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpenSourceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateOpenSourceAccountRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateOpenSourceAccountRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateOpenSourceAccountRequest
	GetInstanceId() *string
	SetPassword(v string) *UpdateOpenSourceAccountRequest
	GetPassword() *string
	SetUserName(v string) *UpdateOpenSourceAccountRequest
	GetUserName() *string
}

type UpdateOpenSourceAccountRequest struct {
	// The client token that is used to ensure the idempotence of the request and prevent repeated submissions.
	//
	// example:
	//
	// f6af6bb123988d191bfe01c9a9b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The remarks.
	//
	// example:
	//
	// project name pass the check
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-20p****04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account password.
	//
	// example:
	//
	// 123456Aa
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// myname
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateOpenSourceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpenSourceAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpenSourceAccountRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateOpenSourceAccountRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOpenSourceAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateOpenSourceAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateOpenSourceAccountRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateOpenSourceAccountRequest) SetClientToken(v string) *UpdateOpenSourceAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOpenSourceAccountRequest) SetDescription(v string) *UpdateOpenSourceAccountRequest {
	s.Description = &v
	return s
}

func (s *UpdateOpenSourceAccountRequest) SetInstanceId(v string) *UpdateOpenSourceAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOpenSourceAccountRequest) SetPassword(v string) *UpdateOpenSourceAccountRequest {
	s.Password = &v
	return s
}

func (s *UpdateOpenSourceAccountRequest) SetUserName(v string) *UpdateOpenSourceAccountRequest {
	s.UserName = &v
	return s
}

func (s *UpdateOpenSourceAccountRequest) Validate() error {
	return dara.Validate(s)
}
