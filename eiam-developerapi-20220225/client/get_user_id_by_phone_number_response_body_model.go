// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetUserIdByPhoneNumberResponseBody
	GetUserId() *string
}

type GetUserIdByPhoneNumberResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdByPhoneNumberResponseBody) SetUserId(v string) *GetUserIdByPhoneNumberResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdByPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
