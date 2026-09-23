// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCredentialId(v string) *UpdateServiceCredentialRequest
	GetServiceCredentialId() *string
	SetServiceCredentialName(v string) *UpdateServiceCredentialRequest
	GetServiceCredentialName() *string
	SetStatus(v string) *UpdateServiceCredentialRequest
	GetStatus() *string
	SetUserPrincipalName(v string) *UpdateServiceCredentialRequest
	GetUserPrincipalName() *string
}

type UpdateServiceCredentialRequest struct {
	// The ID of the service credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// SC****************
	ServiceCredentialId *string `json:"ServiceCredentialId,omitempty" xml:"ServiceCredentialId,omitempty"`
	// The name of the service credential. The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). You must specify at least one of Status and ServiceCredentialName.
	//
	// example:
	//
	// yourServiceCredentialName
	ServiceCredentialName *string `json:"ServiceCredentialName,omitempty" xml:"ServiceCredentialName,omitempty"`
	// The status of the service credential. Valid values: Active, Inactive. You must specify at least one of Status and ServiceCredentialName.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logon name of the Resource Access Management (RAM) user. If not specified, the service credential of the current invoke identity is updated.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateServiceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceCredentialRequest) GetServiceCredentialId() *string {
	return s.ServiceCredentialId
}

func (s *UpdateServiceCredentialRequest) GetServiceCredentialName() *string {
	return s.ServiceCredentialName
}

func (s *UpdateServiceCredentialRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateServiceCredentialRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *UpdateServiceCredentialRequest) SetServiceCredentialId(v string) *UpdateServiceCredentialRequest {
	s.ServiceCredentialId = &v
	return s
}

func (s *UpdateServiceCredentialRequest) SetServiceCredentialName(v string) *UpdateServiceCredentialRequest {
	s.ServiceCredentialName = &v
	return s
}

func (s *UpdateServiceCredentialRequest) SetStatus(v string) *UpdateServiceCredentialRequest {
	s.Status = &v
	return s
}

func (s *UpdateServiceCredentialRequest) SetUserPrincipalName(v string) *UpdateServiceCredentialRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateServiceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
