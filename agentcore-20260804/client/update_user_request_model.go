// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateUserRequestBody) *UpdateUserRequest
	GetBody() *UpdateUserRequestBody
	SetClientToken(v string) *UpdateUserRequest
	GetClientToken() *string
}

type UpdateUserRequest struct {
	// The request body for updating a user.
	Body *UpdateUserRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Not supported.
	//
	// example:
	//
	// Not supported
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetBody() *UpdateUserRequestBody {
	return s.Body
}

func (s *UpdateUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateUserRequest) SetBody(v *UpdateUserRequestBody) *UpdateUserRequest {
	s.Body = v
	return s
}

func (s *UpdateUserRequest) SetClientToken(v string) *UpdateUserRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserRequestBody struct {
	// The display name of the user. The name must be 1 to 32 characters in length. At least one of displayName, email, and note must be specified.
	//
	// example:
	//
	// John
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user. The address can be up to 256 characters in length.
	//
	// example:
	//
	// user-01@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The note for the user. The note can be up to 1,024 characters in length.
	//
	// example:
	//
	// Agent operations team member
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
}

func (s UpdateUserRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateUserRequestBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserRequestBody) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserRequestBody) GetNote() *string {
	return s.Note
}

func (s *UpdateUserRequestBody) SetDisplayName(v string) *UpdateUserRequestBody {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserRequestBody) SetEmail(v string) *UpdateUserRequestBody {
	s.Email = &v
	return s
}

func (s *UpdateUserRequestBody) SetNote(v string) *UpdateUserRequestBody {
	s.Note = &v
	return s
}

func (s *UpdateUserRequestBody) Validate() error {
	return dara.Validate(s)
}
