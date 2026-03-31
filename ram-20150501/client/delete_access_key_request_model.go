// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *DeleteAccessKeyRequest
	GetUserAccessKeyId() *string
	SetUserName(v string) *DeleteAccessKeyRequest
	GetUserName() *string
}

type DeleteAccessKeyRequest struct {
	// The AccessKey ID in the AccessKey pair that you want to delete.``
	//
	// example:
	//
	// 0wNEpMMlzy7s****
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *DeleteAccessKeyRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteAccessKeyRequest) SetUserAccessKeyId(v string) *DeleteAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *DeleteAccessKeyRequest) SetUserName(v string) *DeleteAccessKeyRequest {
	s.UserName = &v
	return s
}

func (s *DeleteAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
