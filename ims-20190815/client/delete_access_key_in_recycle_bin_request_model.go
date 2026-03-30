// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessKeyInRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *DeleteAccessKeyInRecycleBinRequest
	GetUserAccessKeyId() *string
	SetUserId(v string) *DeleteAccessKeyInRecycleBinRequest
	GetUserId() *string
}

type DeleteAccessKeyInRecycleBinRequest struct {
	// The AccessKey ID of the RAM user.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The ID of the RAM user.
	//
	// > - If you use an Alibaba Cloud account to call the operation, you must specify the parameter.
	//
	// > - If you use a RAM user to call the operation, you can leave the parameter empty. In this case, the ID of the RAM user is used by default.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteAccessKeyInRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessKeyInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyInRecycleBinRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *DeleteAccessKeyInRecycleBinRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteAccessKeyInRecycleBinRequest) SetUserAccessKeyId(v string) *DeleteAccessKeyInRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *DeleteAccessKeyInRecycleBinRequest) SetUserId(v string) *DeleteAccessKeyInRecycleBinRequest {
	s.UserId = &v
	return s
}

func (s *DeleteAccessKeyInRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
