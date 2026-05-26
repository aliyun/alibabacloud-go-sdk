// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserPoolClientsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserPoolClientsRequest
	GetNextToken() *string
	SetUserPoolName(v string) *ListUserPoolClientsRequest
	GetUserPoolName() *string
}

type ListUserPoolClientsRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdENsaWVudHM6OjEw
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListUserPoolClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolClientsRequest) GoString() string {
	return s.String()
}

func (s *ListUserPoolClientsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserPoolClientsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserPoolClientsRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListUserPoolClientsRequest) SetMaxResults(v int32) *ListUserPoolClientsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserPoolClientsRequest) SetNextToken(v string) *ListUserPoolClientsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserPoolClientsRequest) SetUserPoolName(v string) *ListUserPoolClientsRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListUserPoolClientsRequest) Validate() error {
	return dara.Validate(s)
}
