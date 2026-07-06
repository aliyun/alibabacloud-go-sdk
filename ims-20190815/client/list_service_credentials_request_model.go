// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllUsers(v bool) *ListServiceCredentialsRequest
	GetAllUsers() *bool
	SetMaxResults(v int32) *ListServiceCredentialsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceCredentialsRequest
	GetNextToken() *string
	SetServiceName(v string) *ListServiceCredentialsRequest
	GetServiceName() *string
	SetUserPrincipalName(v string) *ListServiceCredentialsRequest
	GetUserPrincipalName() *string
}

type ListServiceCredentialsRequest struct {
	// Specifies whether to query service credentials for all Resource Access Management (RAM) users under the Alibaba Cloud account.
	//
	// If this parameter is set to true, you cannot specify UserPrincipalName at the same time.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AllUsers *bool `json:"AllUsers,omitempty" xml:"AllUsers,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. You do not need to specify this parameter for the first API call.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// EXAMPLE*******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service name of the Alibaba Cloud service.
	//
	// example:
	//
	// xxx.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The logon name of the Resource Access Management (RAM) user.
	//
	// Queries the service credentials of the specified RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListServiceCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceCredentialsRequest) GetAllUsers() *bool {
	return s.AllUsers
}

func (s *ListServiceCredentialsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceCredentialsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceCredentialsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServiceCredentialsRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListServiceCredentialsRequest) SetAllUsers(v bool) *ListServiceCredentialsRequest {
	s.AllUsers = &v
	return s
}

func (s *ListServiceCredentialsRequest) SetMaxResults(v int32) *ListServiceCredentialsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceCredentialsRequest) SetNextToken(v string) *ListServiceCredentialsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceCredentialsRequest) SetServiceName(v string) *ListServiceCredentialsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListServiceCredentialsRequest) SetUserPrincipalName(v string) *ListServiceCredentialsRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *ListServiceCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
