// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwarelibSoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassifyId(v string) *ListSoftwarelibSoftwareRequest
	GetClassifyId() *string
	SetCurrentPage(v int64) *ListSoftwarelibSoftwareRequest
	GetCurrentPage() *int64
	SetMaxResults(v int32) *ListSoftwarelibSoftwareRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSoftwarelibSoftwareRequest
	GetNextToken() *string
	SetOs(v string) *ListSoftwarelibSoftwareRequest
	GetOs() *string
	SetPageSize(v int64) *ListSoftwarelibSoftwareRequest
	GetPageSize() *int64
	SetSoftwareName(v string) *ListSoftwarelibSoftwareRequest
	GetSoftwareName() *string
	SetSourceType(v string) *ListSoftwarelibSoftwareRequest
	GetSourceType() *string
}

type ListSoftwarelibSoftwareRequest struct {
	// The software category ID. You can call [ListSoftwarelibClassify](~~ListSoftwarelibClassify~~) to obtain the value.
	//
	// example:
	//
	// softwarelib-classify-61b7ccc63cae****
	ClassifyId *string `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// The page number of the current page displayed during paging in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page. This parameter is not supported by the current operation. Use CurrentPage and PageSize for pagination.
	//
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. This parameter is not supported by the current operation. Use CurrentPage and PageSize for pagination.
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh27/Jy4SXvlU9WgqeV7az+t
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The operating system applicable to the software package. Valid values:
	//
	// - **Windows**: Windows.
	//
	// - **Mac(Apple)**: macOS with Apple silicon.
	//
	// - **Mac(Intel)**: macOS with Intel chips.
	//
	// example:
	//
	// Windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The Settings for the number of entries per page during paging in a paged query. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The software name. Fuzzy match is supported.
	//
	// example:
	//
	// Thunder
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// The software source. Valid values:
	//
	// - **custom**: custom software.
	//
	// - **builtin**: built-in software library.
	//
	// example:
	//
	// builtin
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListSoftwarelibSoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwarelibSoftwareRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwarelibSoftwareRequest) GetClassifyId() *string {
	return s.ClassifyId
}

func (s *ListSoftwarelibSoftwareRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListSoftwarelibSoftwareRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSoftwarelibSoftwareRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSoftwarelibSoftwareRequest) GetOs() *string {
	return s.Os
}

func (s *ListSoftwarelibSoftwareRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSoftwarelibSoftwareRequest) GetSoftwareName() *string {
	return s.SoftwareName
}

func (s *ListSoftwarelibSoftwareRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSoftwarelibSoftwareRequest) SetClassifyId(v string) *ListSoftwarelibSoftwareRequest {
	s.ClassifyId = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) SetCurrentPage(v int64) *ListSoftwarelibSoftwareRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) SetMaxResults(v int32) *ListSoftwarelibSoftwareRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) SetNextToken(v string) *ListSoftwarelibSoftwareRequest {
	s.NextToken = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) SetOs(v string) *ListSoftwarelibSoftwareRequest {
	s.Os = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) SetPageSize(v int64) *ListSoftwarelibSoftwareRequest {
	s.PageSize = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) SetSoftwareName(v string) *ListSoftwarelibSoftwareRequest {
	s.SoftwareName = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) SetSourceType(v string) *ListSoftwarelibSoftwareRequest {
	s.SourceType = &v
	return s
}

func (s *ListSoftwarelibSoftwareRequest) Validate() error {
	return dara.Validate(s)
}
