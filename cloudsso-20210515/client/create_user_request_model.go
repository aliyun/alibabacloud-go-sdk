// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateUserRequest
	GetDirectoryId() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetFirstName(v string) *CreateUserRequest
	GetFirstName() *string
	SetLastName(v string) *CreateUserRequest
	GetLastName() *string
	SetStatus(v string) *CreateUserRequest
	GetStatus() *string
	SetTags(v []*CreateUserRequestTags) *CreateUserRequest
	GetTags() []*CreateUserRequestTags
	SetUserName(v string) *CreateUserRequest
	GetUserName() *string
}

type CreateUserRequest struct {
	// The description of the user.
	//
	// Maximum length: 1024 characters.
	//
	// example:
	//
	// This is a user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The display name of the user.
	//
	// Maximum length: 256 characters.
	//
	// example:
	//
	// Alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user. The email address must be unique within the directory.
	//
	// Maximum length: 128 characters.
	//
	// example:
	//
	// Alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the user.
	//
	// Maximum length: 64 characters.
	//
	// example:
	//
	// Alice
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	//
	// Maximum length: 64 characters.
	//
	// example:
	//
	// Lee
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The status of the user. Valid values:
	//
	// - Enabled (default): Enabled.
	//
	// - Disabled: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*CreateUserRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The username. The username must be unique within the directory and cannot be modified.
	//
	// Format: Can contain digits, letters, and the following special characters: `@_-.`
	//
	// Maximum length: 64 characters.
	//
	// example:
	//
	// Alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetFirstName() *string {
	return s.FirstName
}

func (s *CreateUserRequest) GetLastName() *string {
	return s.LastName
}

func (s *CreateUserRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateUserRequest) GetTags() []*CreateUserRequestTags {
	return s.Tags
}

func (s *CreateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetDirectoryId(v string) *CreateUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetFirstName(v string) *CreateUserRequest {
	s.FirstName = &v
	return s
}

func (s *CreateUserRequest) SetLastName(v string) *CreateUserRequest {
	s.LastName = &v
	return s
}

func (s *CreateUserRequest) SetStatus(v string) *CreateUserRequest {
	s.Status = &v
	return s
}

func (s *CreateUserRequest) SetTags(v []*CreateUserRequestTags) *CreateUserRequest {
	s.Tags = v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateUserRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateUserRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequestTags) GoString() string {
	return s.String()
}

func (s *CreateUserRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateUserRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateUserRequestTags) SetKey(v string) *CreateUserRequestTags {
	s.Key = &v
	return s
}

func (s *CreateUserRequestTags) SetValue(v string) *CreateUserRequestTags {
	s.Value = &v
	return s
}

func (s *CreateUserRequestTags) Validate() error {
	return dara.Validate(s)
}
