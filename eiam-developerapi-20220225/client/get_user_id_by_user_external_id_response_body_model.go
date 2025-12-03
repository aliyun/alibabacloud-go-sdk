// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUserExternalIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetUserIdByUserExternalIdResponseBody
	GetUserId() *string
}

type GetUserIdByUserExternalIdResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByUserExternalIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUserExternalIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdByUserExternalIdResponseBody) SetUserId(v string) *GetUserIdByUserExternalIdResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdByUserExternalIdResponseBody) Validate() error {
	return dara.Validate(s)
}
