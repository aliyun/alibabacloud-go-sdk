// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *GetUserIdByEmailRequest
	GetEmail() *string
}

type GetUserIdByEmailRequest struct {
	// The email address of the user who owns the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s GetUserIdByEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByEmailRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailRequest) GetEmail() *string {
	return s.Email
}

func (s *GetUserIdByEmailRequest) SetEmail(v string) *GetUserIdByEmailRequest {
	s.Email = &v
	return s
}

func (s *GetUserIdByEmailRequest) Validate() error {
	return dara.Validate(s)
}
