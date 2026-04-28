// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *DeleteUserRequest
	GetUserId() *string
}

type DeleteUserRequest struct {
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
