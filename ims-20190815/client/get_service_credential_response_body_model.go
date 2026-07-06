// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceCredentialResponseBody
	GetRequestId() *string
	SetServiceCredential(v *GetServiceCredentialResponseBodyServiceCredential) *GetServiceCredentialResponseBody
	GetServiceCredential() *GetServiceCredentialResponseBodyServiceCredential
}

type GetServiceCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 66815255-7CCE-4759-AC37-9755794C3626
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service credential information.
	ServiceCredential *GetServiceCredentialResponseBodyServiceCredential `json:"ServiceCredential,omitempty" xml:"ServiceCredential,omitempty" type:"Struct"`
}

func (s GetServiceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceCredentialResponseBody) GetServiceCredential() *GetServiceCredentialResponseBodyServiceCredential {
	return s.ServiceCredential
}

func (s *GetServiceCredentialResponseBody) SetRequestId(v string) *GetServiceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceCredentialResponseBody) SetServiceCredential(v *GetServiceCredentialResponseBodyServiceCredential) *GetServiceCredentialResponseBody {
	s.ServiceCredential = v
	return s
}

func (s *GetServiceCredentialResponseBody) Validate() error {
	if s.ServiceCredential != nil {
		if err := s.ServiceCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceCredentialResponseBodyServiceCredential struct {
	// The creation time.
	//
	// example:
	//
	// 2026-03-15T09:20:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time. This field is not returned for permanent service credentials.
	//
	// example:
	//
	// 2026-04-15T09:20:58Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The service credential ID.
	//
	// example:
	//
	// SC***************
	ServiceCredentialId *string `json:"ServiceCredentialId,omitempty" xml:"ServiceCredentialId,omitempty"`
	// The service credential name.
	//
	// example:
	//
	// test
	ServiceCredentialName *string `json:"ServiceCredentialName,omitempty" xml:"ServiceCredentialName,omitempty"`
	// The Alibaba Cloud service name.
	//
	// example:
	//
	// xxx.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service credential status.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logon name of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetServiceCredentialResponseBodyServiceCredential) String() string {
	return dara.Prettify(s)
}

func (s GetServiceCredentialResponseBodyServiceCredential) GoString() string {
	return s.String()
}

func (s *GetServiceCredentialResponseBodyServiceCredential) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetServiceCredentialResponseBodyServiceCredential) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *GetServiceCredentialResponseBodyServiceCredential) GetServiceCredentialId() *string {
	return s.ServiceCredentialId
}

func (s *GetServiceCredentialResponseBodyServiceCredential) GetServiceCredentialName() *string {
	return s.ServiceCredentialName
}

func (s *GetServiceCredentialResponseBodyServiceCredential) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceCredentialResponseBodyServiceCredential) GetStatus() *string {
	return s.Status
}

func (s *GetServiceCredentialResponseBodyServiceCredential) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetServiceCredentialResponseBodyServiceCredential) SetCreateTime(v string) *GetServiceCredentialResponseBodyServiceCredential {
	s.CreateTime = &v
	return s
}

func (s *GetServiceCredentialResponseBodyServiceCredential) SetExpirationTime(v string) *GetServiceCredentialResponseBodyServiceCredential {
	s.ExpirationTime = &v
	return s
}

func (s *GetServiceCredentialResponseBodyServiceCredential) SetServiceCredentialId(v string) *GetServiceCredentialResponseBodyServiceCredential {
	s.ServiceCredentialId = &v
	return s
}

func (s *GetServiceCredentialResponseBodyServiceCredential) SetServiceCredentialName(v string) *GetServiceCredentialResponseBodyServiceCredential {
	s.ServiceCredentialName = &v
	return s
}

func (s *GetServiceCredentialResponseBodyServiceCredential) SetServiceName(v string) *GetServiceCredentialResponseBodyServiceCredential {
	s.ServiceName = &v
	return s
}

func (s *GetServiceCredentialResponseBodyServiceCredential) SetStatus(v string) *GetServiceCredentialResponseBodyServiceCredential {
	s.Status = &v
	return s
}

func (s *GetServiceCredentialResponseBodyServiceCredential) SetUserPrincipalName(v string) *GetServiceCredentialResponseBodyServiceCredential {
	s.UserPrincipalName = &v
	return s
}

func (s *GetServiceCredentialResponseBodyServiceCredential) Validate() error {
	return dara.Validate(s)
}
