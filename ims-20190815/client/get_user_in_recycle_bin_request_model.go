// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *GetUserInRecycleBinRequest
	GetUserAccessKeyId() *string
	SetUserId(v string) *GetUserInRecycleBinRequest
	GetUserId() *string
}

type GetUserInRecycleBinRequest struct {
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserInRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetUserInRecycleBinRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserInRecycleBinRequest) SetUserAccessKeyId(v string) *GetUserInRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetUserInRecycleBinRequest) SetUserId(v string) *GetUserInRecycleBinRequest {
	s.UserId = &v
	return s
}

func (s *GetUserInRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
