// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserSayResponseBody
	GetRequestId() *string
	SetUserSayId(v int64) *DeleteUserSayResponseBody
	GetUserSayId() *int64
}

type DeleteUserSayResponseBody struct {
	// example:
	//
	// dfgdg324gf34t34g34g3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4562121234
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s DeleteUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserSayResponseBody) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *DeleteUserSayResponseBody) SetRequestId(v string) *DeleteUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserSayResponseBody) SetUserSayId(v int64) *DeleteUserSayResponseBody {
	s.UserSayId = &v
	return s
}

func (s *DeleteUserSayResponseBody) Validate() error {
	return dara.Validate(s)
}
