// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListTemplateVersionsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTemplateVersionsRequest
	GetNextToken() *string
	SetTemplateId(v string) *ListTemplateVersionsRequest
	GetTemplateId() *string
}

type ListTemplateVersionsRequest struct {
	// The maximum number of results to be returned in a single call when NextToken is used for the query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplateVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTemplateVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateVersionsRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplateVersionsRequest) SetMaxResults(v int64) *ListTemplateVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetNextToken(v string) *ListTemplateVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetTemplateId(v string) *ListTemplateVersionsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListTemplateVersionsRequest) Validate() error {
	return dara.Validate(s)
}
