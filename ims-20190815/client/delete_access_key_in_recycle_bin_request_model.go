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
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
