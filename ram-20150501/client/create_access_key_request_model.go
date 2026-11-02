// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *CreateAccessKeyRequest
	GetUserName() *string
}

type CreateAccessKeyRequest struct {
	// The username. When a Resource Access Management (RAM) user invokes this operation, an AccessKey pair is created for the RAM user by default.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateAccessKeyRequest) SetUserName(v string) *CreateAccessKeyRequest {
	s.UserName = &v
	return s
}

func (s *CreateAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
