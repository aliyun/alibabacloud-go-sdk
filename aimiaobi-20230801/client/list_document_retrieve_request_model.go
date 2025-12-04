// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentRetrieveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *ListDocumentRetrieveRequest
	GetContentType() *string
	SetElementScope(v string) *ListDocumentRetrieveRequest
	GetElementScope() *string
	SetEndDate(v string) *ListDocumentRetrieveRequest
	GetEndDate() *string
	SetMaxResults(v int32) *ListDocumentRetrieveRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDocumentRetrieveRequest
	GetNextToken() *string
	SetOffice(v string) *ListDocumentRetrieveRequest
	GetOffice() *string
	SetQuery(v string) *ListDocumentRetrieveRequest
	GetQuery() *string
	SetRegion(v string) *ListDocumentRetrieveRequest
	GetRegion() *string
	SetSource(v string) *ListDocumentRetrieveRequest
	GetSource() *string
	SetStartDate(v string) *ListDocumentRetrieveRequest
	GetStartDate() *string
	SetSubContentType(v string) *ListDocumentRetrieveRequest
	GetSubContentType() *string
	SetWordSize(v string) *ListDocumentRetrieveRequest
	GetWordSize() *string
	SetWorkspaceId(v string) *ListDocumentRetrieveRequest
	GetWorkspaceId() *string
}

type ListDocumentRetrieveRequest struct {
	// example:
	//
	// 1
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 0
	ElementScope *string `json:"ElementScope,omitempty" xml:"ElementScope,omitempty"`
	// example:
	//
	// 2025-07-03
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 94
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cEoBWREAXdxaOyjq/cqAbg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 国务院办公室
	Office *string `json:"Office,omitempty" xml:"Office,omitempty"`
	// example:
	//
	// 检索Query
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 北京市
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 2025-10-10
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 1
	SubContentType *string `json:"SubContentType,omitempty" xml:"SubContentType,omitempty"`
	// example:
	//
	// 宁民规〔2020〕5号
	WordSize *string `json:"WordSize,omitempty" xml:"WordSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDocumentRetrieveRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentRetrieveRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentRetrieveRequest) GetContentType() *string {
	return s.ContentType
}

func (s *ListDocumentRetrieveRequest) GetElementScope() *string {
	return s.ElementScope
}

func (s *ListDocumentRetrieveRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListDocumentRetrieveRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentRetrieveRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentRetrieveRequest) GetOffice() *string {
	return s.Office
}

func (s *ListDocumentRetrieveRequest) GetQuery() *string {
	return s.Query
}

func (s *ListDocumentRetrieveRequest) GetRegion() *string {
	return s.Region
}

func (s *ListDocumentRetrieveRequest) GetSource() *string {
	return s.Source
}

func (s *ListDocumentRetrieveRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListDocumentRetrieveRequest) GetSubContentType() *string {
	return s.SubContentType
}

func (s *ListDocumentRetrieveRequest) GetWordSize() *string {
	return s.WordSize
}

func (s *ListDocumentRetrieveRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDocumentRetrieveRequest) SetContentType(v string) *ListDocumentRetrieveRequest {
	s.ContentType = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetElementScope(v string) *ListDocumentRetrieveRequest {
	s.ElementScope = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetEndDate(v string) *ListDocumentRetrieveRequest {
	s.EndDate = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetMaxResults(v int32) *ListDocumentRetrieveRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetNextToken(v string) *ListDocumentRetrieveRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetOffice(v string) *ListDocumentRetrieveRequest {
	s.Office = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetQuery(v string) *ListDocumentRetrieveRequest {
	s.Query = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetRegion(v string) *ListDocumentRetrieveRequest {
	s.Region = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetSource(v string) *ListDocumentRetrieveRequest {
	s.Source = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetStartDate(v string) *ListDocumentRetrieveRequest {
	s.StartDate = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetSubContentType(v string) *ListDocumentRetrieveRequest {
	s.SubContentType = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetWordSize(v string) *ListDocumentRetrieveRequest {
	s.WordSize = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetWorkspaceId(v string) *ListDocumentRetrieveRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDocumentRetrieveRequest) Validate() error {
	return dara.Validate(s)
}
