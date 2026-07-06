// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceCredentialResponseBody
	GetRequestId() *string
	SetServiceCredential(v *CreateServiceCredentialResponseBodyServiceCredential) *CreateServiceCredentialResponseBody
	GetServiceCredential() *CreateServiceCredentialResponseBodyServiceCredential
}

type CreateServiceCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81313F5E-3C85-478F-BCC9-E1B70E4556DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service credential information.
	ServiceCredential *CreateServiceCredentialResponseBodyServiceCredential `json:"ServiceCredential,omitempty" xml:"ServiceCredential,omitempty" type:"Struct"`
}

func (s CreateServiceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceCredentialResponseBody) GetServiceCredential() *CreateServiceCredentialResponseBodyServiceCredential {
	return s.ServiceCredential
}

func (s *CreateServiceCredentialResponseBody) SetRequestId(v string) *CreateServiceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceCredentialResponseBody) SetServiceCredential(v *CreateServiceCredentialResponseBodyServiceCredential) *CreateServiceCredentialResponseBody {
	s.ServiceCredential = v
	return s
}

func (s *CreateServiceCredentialResponseBody) Validate() error {
	if s.ServiceCredential != nil {
		if err := s.ServiceCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceCredentialResponseBodyServiceCredential struct {
	// The time when the service credential was created.
	//
	// example:
	//
	// 2026-01-01T10:05:24Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the service credential.
	//
	// This field is not returned for permanently valid service credentials.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2026-02-01T10:05:24Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The service credential ID.
	//
	// example:
	//
	// SC*************
	ServiceCredentialId *string `json:"ServiceCredentialId,omitempty" xml:"ServiceCredentialId,omitempty"`
	// The service credential name.
	//
	// example:
	//
	// yourServiceCredentialName
	ServiceCredentialName *string `json:"ServiceCredentialName,omitempty" xml:"ServiceCredentialName,omitempty"`
	// The secret of the service credential.
	//
	// example:
	//
	// yourServiceCredentialSecret
	ServiceCredentialSecret *string `json:"ServiceCredentialSecret,omitempty" xml:"ServiceCredentialSecret,omitempty"`
	// The Alibaba Cloud service name.
	//
	// example:
	//
	// xxx.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The status of the service credential.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateServiceCredentialResponseBodyServiceCredential) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCredentialResponseBodyServiceCredential) GoString() string {
	return s.String()
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetServiceCredentialId() *string {
	return s.ServiceCredentialId
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetServiceCredentialName() *string {
	return s.ServiceCredentialName
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetServiceCredentialSecret() *string {
	return s.ServiceCredentialSecret
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetCreateTime(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.CreateTime = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetExpirationTime(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.ExpirationTime = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetServiceCredentialId(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.ServiceCredentialId = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetServiceCredentialName(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.ServiceCredentialName = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetServiceCredentialSecret(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.ServiceCredentialSecret = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetServiceName(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetStatus(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.Status = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) SetUserPrincipalName(v string) *CreateServiceCredentialResponseBodyServiceCredential {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateServiceCredentialResponseBodyServiceCredential) Validate() error {
	return dara.Validate(s)
}
