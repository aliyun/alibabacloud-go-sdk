// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVirtualMFARequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *DisableVirtualMFARequest
	GetUserPrincipalName() *string
}

type DisableVirtualMFARequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DisableVirtualMFARequest) String() string {
	return dara.Prettify(s)
}

func (s DisableVirtualMFARequest) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFARequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DisableVirtualMFARequest) SetUserPrincipalName(v string) *DisableVirtualMFARequest {
	s.UserPrincipalName = &v
	return s
}

func (s *DisableVirtualMFARequest) Validate() error {
	return dara.Validate(s)
}
