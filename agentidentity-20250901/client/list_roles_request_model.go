// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRolesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRolesRequest
	GetNextToken() *string
	SetUserPoolName(v string) *ListRolesRequest
	GetUserPoolName() *string
}

type ListRolesRequest struct {
	// The maximum number of entries to return per page. Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to query the next page of data. This parameter is not required for the first call. If a NextToken is returned in the response of a previous call, specify it in the current call to retrieve the next page.
	//
	// example:
	//
	// TGlzdFJvbGVzOjoxMA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the user pool.
	//
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRolesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRolesRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListRolesRequest) SetMaxResults(v int32) *ListRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRolesRequest) SetNextToken(v string) *ListRolesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRolesRequest) SetUserPoolName(v string) *ListRolesRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
