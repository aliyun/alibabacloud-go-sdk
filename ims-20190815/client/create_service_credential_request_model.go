// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialAgeDays(v int32) *CreateServiceCredentialRequest
	GetCredentialAgeDays() *int32
	SetServiceCredentialName(v string) *CreateServiceCredentialRequest
	GetServiceCredentialName() *string
	SetServiceName(v string) *CreateServiceCredentialRequest
	GetServiceName() *string
	SetUserPrincipalName(v string) *CreateServiceCredentialRequest
	GetUserPrincipalName() *string
}

type CreateServiceCredentialRequest struct {
	// The expiration time of the service credential, in days.
	//
	// Valid values: 1 to 36600.
	//
	// If this parameter is not specified, the service credential is permanently valid.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 30
	CredentialAgeDays *int32 `json:"CredentialAgeDays,omitempty" xml:"CredentialAgeDays,omitempty"`
	// The service credential name.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourServiceCredentialName
	ServiceCredentialName *string `json:"ServiceCredentialName,omitempty" xml:"ServiceCredentialName,omitempty"`
	// The Alibaba Cloud service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The logon name of the RAM user.
	//
	// If this parameter is left empty, a service credential is created for the current user by default.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateServiceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceCredentialRequest) GetCredentialAgeDays() *int32 {
	return s.CredentialAgeDays
}

func (s *CreateServiceCredentialRequest) GetServiceCredentialName() *string {
	return s.ServiceCredentialName
}

func (s *CreateServiceCredentialRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServiceCredentialRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *CreateServiceCredentialRequest) SetCredentialAgeDays(v int32) *CreateServiceCredentialRequest {
	s.CredentialAgeDays = &v
	return s
}

func (s *CreateServiceCredentialRequest) SetServiceCredentialName(v string) *CreateServiceCredentialRequest {
	s.ServiceCredentialName = &v
	return s
}

func (s *CreateServiceCredentialRequest) SetServiceName(v string) *CreateServiceCredentialRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceCredentialRequest) SetUserPrincipalName(v string) *CreateServiceCredentialRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateServiceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
