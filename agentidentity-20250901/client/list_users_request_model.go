// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersRequest
	GetNextToken() *string
	SetUserPoolName(v string) *ListUsersRequest
	GetUserPoolName() *string
}

type ListUsersRequest struct {
	// The maximum number of entries to return per page. Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to query the next page of data. You do not need to specify this parameter for the first call. If a NextToken is returned in the response of the previous call, specify it as the input parameter for the current call to obtain the next page.
	//
	// example:
	//
	// TGlzdFVzZXJzOjoxMA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the user pool.
	//
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetNextToken(v string) *ListUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersRequest) SetUserPoolName(v string) *ListUsersRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
