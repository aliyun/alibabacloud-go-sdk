// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSCIMServerCredentialStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSCIMServerCredentialStatusResponseBody
	GetRequestId() *string
	SetSCIMServerCredential(v *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) *UpdateSCIMServerCredentialStatusResponseBody
	GetSCIMServerCredential() *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential
}

type UpdateSCIMServerCredentialStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7C086C2F-1C66-57B3-B14E-2C1DA70727CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SCIM credential.
	SCIMServerCredential *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential `json:"SCIMServerCredential,omitempty" xml:"SCIMServerCredential,omitempty" type:"Struct"`
}

func (s UpdateSCIMServerCredentialStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSCIMServerCredentialStatusResponseBody) GetSCIMServerCredential() *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	return s.SCIMServerCredential
}

func (s *UpdateSCIMServerCredentialStatusResponseBody) SetRequestId(v string) *UpdateSCIMServerCredentialStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBody) SetSCIMServerCredential(v *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) *UpdateSCIMServerCredentialStatusResponseBody {
	s.SCIMServerCredential = v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBody) Validate() error {
	if s.SCIMServerCredential != nil {
		if err := s.SCIMServerCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential struct {
	// The time when the SCIM credential was created.
	//
	// example:
	//
	// 2021-11-09T08:12:52Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the SCIM credential.
	//
	// example:
	//
	// scimcred-004whl0kvfwcypbi****
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The type of the SCIM credential.
	//
	// example:
	//
	// BearerToken
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The time when the SCIM credential expires.
	//
	// example:
	//
	// 2022-11-09T08:12:52Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the SCIM credential. Valid values:
	//
	// - Enabled: The SCIM credential is enabled.
	//
	// - Disabled: The SCIM credential is disabled.
	//
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) String() string {
	return dara.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GetCredentialType() *string {
	return s.CredentialType
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GetStatus() *string {
	return s.Status
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetCreateTime(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.CreateTime = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetCredentialId(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.CredentialId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetCredentialType(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.CredentialType = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetDirectoryId(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.DirectoryId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetExpireTime(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.ExpireTime = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetStatus(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.Status = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) Validate() error {
	return dara.Validate(s)
}
