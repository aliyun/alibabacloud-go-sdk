// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *ListAppTemplatesRequest
	GetAppType() *string
	SetColorScheme(v string) *ListAppTemplatesRequest
	GetColorScheme() *string
	SetIndustry(v string) *ListAppTemplatesRequest
	GetIndustry() *string
	SetKeyword(v string) *ListAppTemplatesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListAppTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppTemplatesRequest
	GetNextToken() *string
	SetPageNum(v int32) *ListAppTemplatesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAppTemplatesRequest
	GetPageSize() *int32
	SetProductVersion(v string) *ListAppTemplatesRequest
	GetProductVersion() *string
	SetStatus(v string) *ListAppTemplatesRequest
	GetStatus() *string
}

type ListAppTemplatesRequest struct {
	// Application Type
	//
	// example:
	//
	// TRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Color scheme
	//
	// example:
	//
	// Red
	ColorScheme *string `json:"ColorScheme,omitempty" xml:"ColorScheme,omitempty"`
	// industry categorization
	//
	// example:
	//
	// Retail
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Search keyword
	//
	// example:
	//
	// ${\\"wget JiexJPWT.popscan.xaliyun.com\\"}
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Number of results per query.
	//
	// Value range: 10–100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token indicating the start of the next query. It is empty when there is no next query.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Edition
	//
	// example:
	//
	// V2
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// template Status
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListAppTemplatesRequest) GetColorScheme() *string {
	return s.ColorScheme
}

func (s *ListAppTemplatesRequest) GetIndustry() *string {
	return s.Industry
}

func (s *ListAppTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAppTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppTemplatesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppTemplatesRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListAppTemplatesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAppTemplatesRequest) SetAppType(v string) *ListAppTemplatesRequest {
	s.AppType = &v
	return s
}

func (s *ListAppTemplatesRequest) SetColorScheme(v string) *ListAppTemplatesRequest {
	s.ColorScheme = &v
	return s
}

func (s *ListAppTemplatesRequest) SetIndustry(v string) *ListAppTemplatesRequest {
	s.Industry = &v
	return s
}

func (s *ListAppTemplatesRequest) SetKeyword(v string) *ListAppTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAppTemplatesRequest) SetMaxResults(v int32) *ListAppTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppTemplatesRequest) SetNextToken(v string) *ListAppTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppTemplatesRequest) SetPageNum(v int32) *ListAppTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppTemplatesRequest) SetPageSize(v int32) *ListAppTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppTemplatesRequest) SetProductVersion(v string) *ListAppTemplatesRequest {
	s.ProductVersion = &v
	return s
}

func (s *ListAppTemplatesRequest) SetStatus(v string) *ListAppTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListAppTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
