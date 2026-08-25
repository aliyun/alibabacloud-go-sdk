// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSCIMServerCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSCIMServerCredentialsResponseBody
	GetRequestId() *string
	SetSCIMServerCredentials(v []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials) *ListSCIMServerCredentialsResponseBody
	GetSCIMServerCredentials() []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials
	SetTotalCounts(v int32) *ListSCIMServerCredentialsResponseBody
	GetTotalCounts() *int32
}

type ListSCIMServerCredentialsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE4B7037-C315-5DD5-826E-57A87950BCD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SCIM credentials.
	SCIMServerCredentials []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials `json:"SCIMServerCredentials,omitempty" xml:"SCIMServerCredentials,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListSCIMServerCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSCIMServerCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSCIMServerCredentialsResponseBody) GetSCIMServerCredentials() []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	return s.SCIMServerCredentials
}

func (s *ListSCIMServerCredentialsResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListSCIMServerCredentialsResponseBody) SetRequestId(v string) *ListSCIMServerCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBody) SetSCIMServerCredentials(v []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials) *ListSCIMServerCredentialsResponseBody {
	s.SCIMServerCredentials = v
	return s
}

func (s *ListSCIMServerCredentialsResponseBody) SetTotalCounts(v int32) *ListSCIMServerCredentialsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBody) Validate() error {
	if s.SCIMServerCredentials != nil {
		for _, item := range s.SCIMServerCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSCIMServerCredentialsResponseBodySCIMServerCredentials struct {
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
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSCIMServerCredentialsResponseBodySCIMServerCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GetCredentialType() *string {
	return s.CredentialType
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GetStatus() *string {
	return s.Status
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetCreateTime(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetCredentialId(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.CredentialId = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetCredentialType(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.CredentialType = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetDirectoryId(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.DirectoryId = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetExpireTime(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.ExpireTime = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetStatus(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.Status = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) Validate() error {
	return dara.Validate(s)
}
