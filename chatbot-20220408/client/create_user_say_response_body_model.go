// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserSayResponseBody
	GetRequestId() *string
	SetUserSayId(v int64) *CreateUserSayResponseBody
	GetUserSayId() *int64
}

type CreateUserSayResponseBody struct {
	// example:
	//
	// 8g4n8bnd236fg79
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 46456176856
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s CreateUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserSayResponseBody) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *CreateUserSayResponseBody) SetRequestId(v string) *CreateUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserSayResponseBody) SetUserSayId(v int64) *CreateUserSayResponseBody {
	s.UserSayId = &v
	return s
}

func (s *CreateUserSayResponseBody) Validate() error {
	return dara.Validate(s)
}
