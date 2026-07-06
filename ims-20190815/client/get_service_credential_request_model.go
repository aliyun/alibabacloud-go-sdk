// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCredentialId(v string) *GetServiceCredentialRequest
	GetServiceCredentialId() *string
	SetUserPrincipalName(v string) *GetServiceCredentialRequest
	GetUserPrincipalName() *string
}

type GetServiceCredentialRequest struct {
	// The service credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// SC****************
	ServiceCredentialId *string `json:"ServiceCredentialId,omitempty" xml:"ServiceCredentialId,omitempty"`
	// The logon name of the Resource Access Management (RAM) user.
	//
	// If not specified, the service credential of the current caller identity that invokes this operation is retrieved.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetServiceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetServiceCredentialRequest) GetServiceCredentialId() *string {
	return s.ServiceCredentialId
}

func (s *GetServiceCredentialRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetServiceCredentialRequest) SetServiceCredentialId(v string) *GetServiceCredentialRequest {
	s.ServiceCredentialId = &v
	return s
}

func (s *GetServiceCredentialRequest) SetUserPrincipalName(v string) *GetServiceCredentialRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetServiceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
