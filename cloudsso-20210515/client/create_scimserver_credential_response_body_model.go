// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSCIMServerCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSCIMServerCredentialResponseBody
	GetRequestId() *string
	SetSCIMServerCredential(v *CreateSCIMServerCredentialResponseBodySCIMServerCredential) *CreateSCIMServerCredentialResponseBody
	GetSCIMServerCredential() *CreateSCIMServerCredentialResponseBodySCIMServerCredential
}

type CreateSCIMServerCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D2E5180-7ACF-57FF-A56C-26A49ABEBFF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SCIM credential.
	SCIMServerCredential *CreateSCIMServerCredentialResponseBodySCIMServerCredential `json:"SCIMServerCredential,omitempty" xml:"SCIMServerCredential,omitempty" type:"Struct"`
}

func (s CreateSCIMServerCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSCIMServerCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSCIMServerCredentialResponseBody) GetSCIMServerCredential() *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	return s.SCIMServerCredential
}

func (s *CreateSCIMServerCredentialResponseBody) SetRequestId(v string) *CreateSCIMServerCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBody) SetSCIMServerCredential(v *CreateSCIMServerCredentialResponseBodySCIMServerCredential) *CreateSCIMServerCredentialResponseBody {
	s.SCIMServerCredential = v
	return s
}

func (s *CreateSCIMServerCredentialResponseBody) Validate() error {
	if s.SCIMServerCredential != nil {
		if err := s.SCIMServerCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSCIMServerCredentialResponseBodySCIMServerCredential struct {
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
	// The SCIM credential.
	//
	// > The SCIM credential is returned only when it is created. After the SCIM credential is created, you cannot query it. Keep the SCIM credential confidential.
	//
	// example:
	//
	// 8aAJCtpbyPJ8saXeYDgyw****
	CredentialSecret *string `json:"CredentialSecret,omitempty" xml:"CredentialSecret,omitempty"`
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
	// The status of the SCIM credential. The value is fixed as Enabled, which indicates that the SCIM credential is enabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSCIMServerCredentialResponseBodySCIMServerCredential) String() string {
	return dara.Prettify(s)
}

func (s CreateSCIMServerCredentialResponseBodySCIMServerCredential) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) GetCredentialSecret() *string {
	return s.CredentialSecret
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) GetStatus() *string {
	return s.Status
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCreateTime(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CreateTime = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCredentialId(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CredentialId = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCredentialSecret(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CredentialSecret = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCredentialType(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CredentialType = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetDirectoryId(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.DirectoryId = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetExpireTime(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.ExpireTime = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetStatus(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.Status = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) Validate() error {
	return dara.Validate(s)
}
