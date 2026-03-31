// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListReportTemplatesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListReportTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListReportTemplatesRequest
	GetNextToken() *string
}

type ListReportTemplatesRequest struct {
	// example:
	//
	// test-description
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListReportTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReportTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListReportTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListReportTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReportTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListReportTemplatesRequest) SetKeyword(v string) *ListReportTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListReportTemplatesRequest) SetMaxResults(v int32) *ListReportTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListReportTemplatesRequest) SetNextToken(v string) *ListReportTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListReportTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
