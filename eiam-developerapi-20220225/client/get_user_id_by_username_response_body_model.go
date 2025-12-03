// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUsernameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetUserIdByUsernameResponseBody
	GetUserId() *string
}

type GetUserIdByUsernameResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByUsernameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUsernameResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdByUsernameResponseBody) SetUserId(v string) *GetUserIdByUsernameResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdByUsernameResponseBody) Validate() error {
	return dara.Validate(s)
}
