// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataAgentThemeResponseBodyData) *ListDataAgentThemeResponseBody
	GetData() []*ListDataAgentThemeResponseBodyData
	SetErrorCode(v string) *ListDataAgentThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentThemeResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataAgentThemeResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentThemeResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataAgentThemeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAgentThemeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataAgentThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentThemeResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListDataAgentThemeResponseBody
	GetTotal() *int32
	SetTotalPages(v int32) *ListDataAgentThemeResponseBody
	GetTotalPages() *int32
}

type ListDataAgentThemeResponseBody struct {
	// The response struct.
	Data []*ListDataAgentThemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned when the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The current page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDataAgentThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentThemeResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentThemeResponseBody) GetData() []*ListDataAgentThemeResponseBodyData {
	return s.Data
}

func (s *ListDataAgentThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentThemeResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentThemeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentThemeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAgentThemeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAgentThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentThemeResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDataAgentThemeResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListDataAgentThemeResponseBody) SetData(v []*ListDataAgentThemeResponseBodyData) *ListDataAgentThemeResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetErrorCode(v string) *ListDataAgentThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetErrorMessage(v string) *ListDataAgentThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetMaxResults(v int32) *ListDataAgentThemeResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetNextToken(v string) *ListDataAgentThemeResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetPageNumber(v int32) *ListDataAgentThemeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetPageSize(v int32) *ListDataAgentThemeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetRequestId(v string) *ListDataAgentThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetSuccess(v bool) *ListDataAgentThemeResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetTotal(v int32) *ListDataAgentThemeResponseBody {
	s.Total = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) SetTotalPages(v int32) *ListDataAgentThemeResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListDataAgentThemeResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAgentThemeResponseBodyData struct {
	// The common scenarios. Valid values: report, infographic, and others.
	//
	// example:
	//
	// report
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2025-06-15T08:30:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The description.
	//
	// example:
	//
	// weekly report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The modification time in ISO 8601 format.
	//
	// example:
	//
	// 2025-06-20T10:15:30Z
	ModifiedAt *string `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	// The tracing reference that points to the UUID of the source theme.
	//
	// example:
	//
	// 6d1e3f9a-****-****-****-2b8c4e6f0a1d
	ReferTo *string `json:"ReferTo,omitempty" xml:"ReferTo,omitempty"`
	// The source of the theme. Valid values:
	//
	// - system
	//
	// - custom
	//
	// - derived
	//
	// example:
	//
	// custom
	ThemeFrom *string `json:"ThemeFrom,omitempty" xml:"ThemeFrom,omitempty"`
	// The business identifier of the theme.
	//
	// example:
	//
	// 0f8b2c1d-****-****-****-9a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// The display name of the theme.
	//
	// example:
	//
	// weekly report
	ThemeName *string `json:"ThemeName,omitempty" xml:"ThemeName,omitempty"`
	// The theme stage. Valid values:
	//
	// - design: contains only design.md.
	//
	// - template: complete and renderable.
	//
	// example:
	//
	// template
	ThemeType *string `json:"ThemeType,omitempty" xml:"ThemeType,omitempty"`
}

func (s ListDataAgentThemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentThemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentThemeResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *ListDataAgentThemeResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListDataAgentThemeResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListDataAgentThemeResponseBodyData) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *ListDataAgentThemeResponseBodyData) GetReferTo() *string {
	return s.ReferTo
}

func (s *ListDataAgentThemeResponseBodyData) GetThemeFrom() *string {
	return s.ThemeFrom
}

func (s *ListDataAgentThemeResponseBodyData) GetThemeId() *string {
	return s.ThemeId
}

func (s *ListDataAgentThemeResponseBodyData) GetThemeName() *string {
	return s.ThemeName
}

func (s *ListDataAgentThemeResponseBodyData) GetThemeType() *string {
	return s.ThemeType
}

func (s *ListDataAgentThemeResponseBodyData) SetCategory(v string) *ListDataAgentThemeResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetCreatedAt(v string) *ListDataAgentThemeResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetDescription(v string) *ListDataAgentThemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetModifiedAt(v string) *ListDataAgentThemeResponseBodyData {
	s.ModifiedAt = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetReferTo(v string) *ListDataAgentThemeResponseBodyData {
	s.ReferTo = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetThemeFrom(v string) *ListDataAgentThemeResponseBodyData {
	s.ThemeFrom = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetThemeId(v string) *ListDataAgentThemeResponseBodyData {
	s.ThemeId = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetThemeName(v string) *ListDataAgentThemeResponseBodyData {
	s.ThemeName = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) SetThemeType(v string) *ListDataAgentThemeResponseBodyData {
	s.ThemeType = &v
	return s
}

func (s *ListDataAgentThemeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
