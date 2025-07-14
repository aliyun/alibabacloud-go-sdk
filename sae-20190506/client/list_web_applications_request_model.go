// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListWebApplicationsRequest
	GetLimit() *int32
	SetNamespaceId(v string) *ListWebApplicationsRequest
	GetNamespaceId() *string
	SetNextToken(v string) *ListWebApplicationsRequest
	GetNextToken() *string
	SetPrefix(v string) *ListWebApplicationsRequest
	GetPrefix() *string
}

type ListWebApplicationsRequest struct {
	// The number of applications returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// MTIzNCNhYmM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix of the application name.
	//
	// example:
	//
	// my-application
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListWebApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListWebApplicationsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListWebApplicationsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListWebApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebApplicationsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListWebApplicationsRequest) SetLimit(v int32) *ListWebApplicationsRequest {
	s.Limit = &v
	return s
}

func (s *ListWebApplicationsRequest) SetNamespaceId(v string) *ListWebApplicationsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListWebApplicationsRequest) SetNextToken(v string) *ListWebApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWebApplicationsRequest) SetPrefix(v string) *ListWebApplicationsRequest {
	s.Prefix = &v
	return s
}

func (s *ListWebApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
