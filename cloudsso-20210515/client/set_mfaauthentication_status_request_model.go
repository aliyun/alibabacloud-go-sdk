// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMFAAuthenticationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SetMFAAuthenticationStatusRequest
	GetDirectoryId() *string
	SetMFAAuthenticationStatus(v string) *SetMFAAuthenticationStatusRequest
	GetMFAAuthenticationStatus() *string
}

type SetMFAAuthenticationStatusRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The status of MFA. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	MFAAuthenticationStatus *string `json:"MFAAuthenticationStatus,omitempty" xml:"MFAAuthenticationStatus,omitempty"`
}

func (s SetMFAAuthenticationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetMFAAuthenticationStatusRequest) GoString() string {
	return s.String()
}

func (s *SetMFAAuthenticationStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetMFAAuthenticationStatusRequest) GetMFAAuthenticationStatus() *string {
	return s.MFAAuthenticationStatus
}

func (s *SetMFAAuthenticationStatusRequest) SetDirectoryId(v string) *SetMFAAuthenticationStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetMFAAuthenticationStatusRequest) SetMFAAuthenticationStatus(v string) *SetMFAAuthenticationStatusRequest {
	s.MFAAuthenticationStatus = &v
	return s
}

func (s *SetMFAAuthenticationStatusRequest) Validate() error {
	return dara.Validate(s)
}
