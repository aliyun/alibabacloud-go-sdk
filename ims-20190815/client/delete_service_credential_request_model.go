// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCredentialId(v string) *DeleteServiceCredentialRequest
	GetServiceCredentialId() *string
	SetUserPrincipalName(v string) *DeleteServiceCredentialRequest
	GetUserPrincipalName() *string
}

type DeleteServiceCredentialRequest struct {
	// The service credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// SC****************
	ServiceCredentialId *string `json:"ServiceCredentialId,omitempty" xml:"ServiceCredentialId,omitempty"`
	// The logon name of the Resource Access Management (RAM) user. If this parameter is not specified, the service credential of the identity that invokes this operation is deleted.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteServiceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceCredentialRequest) GetServiceCredentialId() *string {
	return s.ServiceCredentialId
}

func (s *DeleteServiceCredentialRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DeleteServiceCredentialRequest) SetServiceCredentialId(v string) *DeleteServiceCredentialRequest {
	s.ServiceCredentialId = &v
	return s
}

func (s *DeleteServiceCredentialRequest) SetUserPrincipalName(v string) *DeleteServiceCredentialRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *DeleteServiceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
