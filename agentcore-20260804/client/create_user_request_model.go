// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateUserRequestBody) *CreateUserRequest
	GetBody() *CreateUserRequestBody
	SetClientToken(v string) *CreateUserRequest
	GetClientToken() *string
}

type CreateUserRequest struct {
	// The request body for creating a user.
	Body *CreateUserRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Not supported.
	//
	// example:
	//
	// Not supported
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetBody() *CreateUserRequestBody {
	return s.Body
}

func (s *CreateUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateUserRequest) SetBody(v *CreateUserRequestBody) *CreateUserRequest {
	s.Body = v
	return s
}

func (s *CreateUserRequest) SetClientToken(v string) *CreateUserRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserRequestBody struct {
	// The display name of the user. The name must be 1 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// John Smith
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user. The address can be up to 256 characters in length.
	//
	// example:
	//
	// user-01@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The username. The name must be unique within the workspace and can contain only lowercase letters, digits, and hyphens. It must start and end with a lowercase letter or digit and must be 1 to 32 characters in length. Reserved names such as manager, admin, or names starting with worker- cannot be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The remarks of the user. The remarks can be up to 1024 characters in length.
	//
	// example:
	//
	// Member of the agent operations team
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// The initial password of the user. The password must be 8 to 32 characters in length and must contain uppercase letters, lowercase letters, digits, and special characters. The password cannot contain the username. If this parameter is not specified, the server generates a random password and returns it in the initialPassword field of the response.
	//
	// example:
	//
	// Example@2026
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s CreateUserRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequestBody) GoString() string {
	return s.String()
}

func (s *CreateUserRequestBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequestBody) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateUserRequestBody) GetNote() *string {
	return s.Note
}

func (s *CreateUserRequestBody) GetPassword() *string {
	return s.Password
}

func (s *CreateUserRequestBody) SetDisplayName(v string) *CreateUserRequestBody {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequestBody) SetEmail(v string) *CreateUserRequestBody {
	s.Email = &v
	return s
}

func (s *CreateUserRequestBody) SetName(v string) *CreateUserRequestBody {
	s.Name = &v
	return s
}

func (s *CreateUserRequestBody) SetNote(v string) *CreateUserRequestBody {
	s.Note = &v
	return s
}

func (s *CreateUserRequestBody) SetPassword(v string) *CreateUserRequestBody {
	s.Password = &v
	return s
}

func (s *CreateUserRequestBody) Validate() error {
	return dara.Validate(s)
}
