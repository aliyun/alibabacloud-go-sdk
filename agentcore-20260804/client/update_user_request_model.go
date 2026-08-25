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
	Body *UpdateUserRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// 暂不支持
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
	// example:
	//
	// 张三
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// user-01@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 智能体运营组成员
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
