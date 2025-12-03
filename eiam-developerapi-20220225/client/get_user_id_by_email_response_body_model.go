// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetUserIdByEmailResponseBody
	GetUserId() *string
}

type GetUserIdByEmailResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByEmailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdByEmailResponseBody) SetUserId(v string) *GetUserIdByEmailResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdByEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
