// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateUserRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *UpdateUserRequest
	GetEmail() *string
	SetStatus(v string) *UpdateUserRequest
	GetStatus() *string
	SetUserName(v string) *UpdateUserRequest
	GetUserName() *string
	SetUserPoolName(v string) *UpdateUserRequest
	GetUserPoolName() *string
}

type UpdateUserRequest struct {
	// The description of the user. The description must be 0 to 256 characters in length.
	//
	// example:
	//
	// Senior data analyst
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user. The display name must be 0 to 128 characters in length.
	//
	// example:
	//
	// Alice Wang
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// alice.wang@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The status of the user. Valid values:
	//
	// - Enabled: The user is enabled.
	//
	// - Disabled: The user is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The username of the user to be updated. The username can contain letters, digits, underscores (_), periods (.), and hyphens (-), and must be 3 to 128 characters in length.
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The name of the user pool.
	//
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateUserRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *UpdateUserRequest) SetDescription(v string) *UpdateUserRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserRequest) SetDisplayName(v string) *UpdateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetStatus(v string) *UpdateUserRequest {
	s.Status = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

func (s *UpdateUserRequest) SetUserPoolName(v string) *UpdateUserRequest {
	s.UserPoolName = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
