// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *GetUserMFAInfoRequest
	GetUserName() *string
}

type GetUserMFAInfoRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserMFAInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetUserMFAInfoRequest) SetUserName(v string) *GetUserMFAInfoRequest {
	s.UserName = &v
	return s
}

func (s *GetUserMFAInfoRequest) Validate() error {
	return dara.Validate(s)
}
