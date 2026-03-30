// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindMFADeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *UnbindMFADeviceRequest
	GetUserPrincipalName() *string
}

type UnbindMFADeviceRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UnbindMFADeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *UnbindMFADeviceRequest) SetUserPrincipalName(v string) *UnbindMFADeviceRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UnbindMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
