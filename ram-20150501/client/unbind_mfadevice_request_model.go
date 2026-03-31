// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindMFADeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *UnbindMFADeviceRequest
	GetUserName() *string
}

type UnbindMFADeviceRequest struct {
	// Specifies the username.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UnbindMFADeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceRequest) GetUserName() *string {
	return s.UserName
}

func (s *UnbindMFADeviceRequest) SetUserName(v string) *UnbindMFADeviceRequest {
	s.UserName = &v
	return s
}

func (s *UnbindMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
