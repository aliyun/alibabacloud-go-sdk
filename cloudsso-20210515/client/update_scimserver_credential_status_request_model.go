// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSCIMServerCredentialStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *UpdateSCIMServerCredentialStatusRequest
	GetCredentialId() *string
	SetDirectoryId(v string) *UpdateSCIMServerCredentialStatusRequest
	GetDirectoryId() *string
	SetNewStatus(v string) *UpdateSCIMServerCredentialStatusRequest
	GetNewStatus() *string
}

type UpdateSCIMServerCredentialStatusRequest struct {
	// The ID of the SCIM credential.
	//
	// example:
	//
	// scimcred-004whl0kvfwcypbi****
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new status of the SCIM credential. Valid values:
	//
	// - Enabled: The SCIM credential is enabled.
	//
	// - Disabled: The SCIM credential is disabled.
	//
	// example:
	//
	// Disabled
	NewStatus *string `json:"NewStatus,omitempty" xml:"NewStatus,omitempty"`
}

func (s UpdateSCIMServerCredentialStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateSCIMServerCredentialStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateSCIMServerCredentialStatusRequest) GetNewStatus() *string {
	return s.NewStatus
}

func (s *UpdateSCIMServerCredentialStatusRequest) SetCredentialId(v string) *UpdateSCIMServerCredentialStatusRequest {
	s.CredentialId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusRequest) SetDirectoryId(v string) *UpdateSCIMServerCredentialStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusRequest) SetNewStatus(v string) *UpdateSCIMServerCredentialStatusRequest {
	s.NewStatus = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusRequest) Validate() error {
	return dara.Validate(s)
}
