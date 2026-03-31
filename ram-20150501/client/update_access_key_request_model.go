// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpdateAccessKeyRequest
	GetStatus() *string
	SetUserAccessKeyId(v string) *UpdateAccessKeyRequest
	GetUserAccessKeyId() *string
	SetUserName(v string) *UpdateAccessKeyRequest
	GetUserName() *string
}

type UpdateAccessKeyRequest struct {
	// The status of the AccessKey pair. Valid values: `Active` and `Inactive`.
	//
	// example:
	//
	// Inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The AccessKey ID in the AccessKey pair whose status you want to change.``
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

func (s UpdateAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAccessKeyRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *UpdateAccessKeyRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateAccessKeyRequest) SetStatus(v string) *UpdateAccessKeyRequest {
	s.Status = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserAccessKeyId(v string) *UpdateAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserName(v string) *UpdateAccessKeyRequest {
	s.UserName = &v
	return s
}

func (s *UpdateAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
