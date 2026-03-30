// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserInRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *DeleteUserInRecycleBinRequest
	GetUserId() *string
}

type DeleteUserInRecycleBinRequest struct {
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserInRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserInRecycleBinRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserInRecycleBinRequest) SetUserId(v string) *DeleteUserInRecycleBinRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserInRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
