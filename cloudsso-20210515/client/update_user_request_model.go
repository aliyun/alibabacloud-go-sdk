// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateUserRequest
	GetDirectoryId() *string
	SetNewDescription(v string) *UpdateUserRequest
	GetNewDescription() *string
	SetNewDisplayName(v string) *UpdateUserRequest
	GetNewDisplayName() *string
	SetNewEmail(v string) *UpdateUserRequest
	GetNewEmail() *string
	SetNewFirstName(v string) *UpdateUserRequest
	GetNewFirstName() *string
	SetNewLastName(v string) *UpdateUserRequest
	GetNewLastName() *string
	SetUserId(v string) *UpdateUserRequest
	GetUserId() *string
}

type UpdateUserRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new description of the user.
	//
	// example:
	//
	// This is a user.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new display name of the user.
	//
	// example:
	//
	// AliceLee
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The new email address of the user.
	//
	// example:
	//
	// AliceLee@example.com
	NewEmail *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	// The new first name of the user.
	//
	// example:
	//
	// Alice
	NewFirstName *string `json:"NewFirstName,omitempty" xml:"NewFirstName,omitempty"`
	// The new last name of the user.
	//
	// example:
	//
	// Lee
	NewLastName *string `json:"NewLastName,omitempty" xml:"NewLastName,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateUserRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateUserRequest) GetNewDisplayName() *string {
	return s.NewDisplayName
}

func (s *UpdateUserRequest) GetNewEmail() *string {
	return s.NewEmail
}

func (s *UpdateUserRequest) GetNewFirstName() *string {
	return s.NewFirstName
}

func (s *UpdateUserRequest) GetNewLastName() *string {
	return s.NewLastName
}

func (s *UpdateUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserRequest) SetDirectoryId(v string) *UpdateUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserRequest) SetNewDescription(v string) *UpdateUserRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateUserRequest) SetNewDisplayName(v string) *UpdateUserRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetNewEmail(v string) *UpdateUserRequest {
	s.NewEmail = &v
	return s
}

func (s *UpdateUserRequest) SetNewFirstName(v string) *UpdateUserRequest {
	s.NewFirstName = &v
	return s
}

func (s *UpdateUserRequest) SetNewLastName(v string) *UpdateUserRequest {
	s.NewLastName = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
