// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserSayResponseBody
	GetRequestId() *string
	SetUserSayId(v int64) *UpdateUserSayResponseBody
	GetUserSayId() *int64
}

type UpdateUserSayResponseBody struct {
	// example:
	//
	// 2356fg3wf34634vdt23wef2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 34512323
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s UpdateUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserSayResponseBody) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *UpdateUserSayResponseBody) SetRequestId(v string) *UpdateUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserSayResponseBody) SetUserSayId(v int64) *UpdateUserSayResponseBody {
	s.UserSayId = &v
	return s
}

func (s *UpdateUserSayResponseBody) Validate() error {
	return dara.Validate(s)
}
