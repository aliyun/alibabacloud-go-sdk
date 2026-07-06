// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListServiceCredentialsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListServiceCredentialsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceCredentialsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceCredentialsResponseBody
	GetRequestId() *string
	SetServiceCredentials(v []*ListServiceCredentialsResponseBodyServiceCredentials) *ListServiceCredentialsResponseBody
	GetServiceCredentials() []*ListServiceCredentialsResponseBodyServiceCredentials
}

type ListServiceCredentialsResponseBody struct {
	// Indicates whether there is a next page of results.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// EXAMPLE*******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D80A0F97-6F12-5CD1-A70A-77A03BF4CFC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of service credentials.
	ServiceCredentials []*ListServiceCredentialsResponseBodyServiceCredentials `json:"ServiceCredentials,omitempty" xml:"ServiceCredentials,omitempty" type:"Repeated"`
}

func (s ListServiceCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceCredentialsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListServiceCredentialsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceCredentialsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceCredentialsResponseBody) GetServiceCredentials() []*ListServiceCredentialsResponseBodyServiceCredentials {
	return s.ServiceCredentials
}

func (s *ListServiceCredentialsResponseBody) SetIsTruncated(v bool) *ListServiceCredentialsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetMaxResults(v int32) *ListServiceCredentialsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetNextToken(v string) *ListServiceCredentialsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetRequestId(v string) *ListServiceCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetServiceCredentials(v []*ListServiceCredentialsResponseBodyServiceCredentials) *ListServiceCredentialsResponseBody {
	s.ServiceCredentials = v
	return s
}

func (s *ListServiceCredentialsResponseBody) Validate() error {
	if s.ServiceCredentials != nil {
		for _, item := range s.ServiceCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceCredentialsResponseBodyServiceCredentials struct {
	// The time when the service credential was created.
	//
	// example:
	//
	// 2026-05-07T05:49:57Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time. This field is not returned for permanent service credentials.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2026-06-07T05:49:57Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The ID of the service credential.
	//
	// example:
	//
	// SC*************
	ServiceCredentialId *string `json:"ServiceCredentialId,omitempty" xml:"ServiceCredentialId,omitempty"`
	// The name of the service credential.
	//
	// example:
	//
	// test
	ServiceCredentialName *string `json:"ServiceCredentialName,omitempty" xml:"ServiceCredentialName,omitempty"`
	// The service name of the Alibaba Cloud service.
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
	// The logon name of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListServiceCredentialsResponseBodyServiceCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCredentialsResponseBodyServiceCredentials) GoString() string {
	return s.String()
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetServiceCredentialId() *string {
	return s.ServiceCredentialId
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetServiceCredentialName() *string {
	return s.ServiceCredentialName
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetStatus() *string {
	return s.Status
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetCreateTime(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetExpirationTime(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.ExpirationTime = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetServiceCredentialId(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.ServiceCredentialId = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetServiceCredentialName(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.ServiceCredentialName = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetServiceName(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.ServiceName = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetStatus(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.Status = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetUserPrincipalName(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.UserPrincipalName = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) Validate() error {
	return dara.Validate(s)
}
