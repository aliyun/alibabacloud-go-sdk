// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDataAgentThemeResponseBodyData) *DescribeDataAgentThemeResponseBody
	GetData() *DescribeDataAgentThemeResponseBodyData
	SetErrorCode(v string) *DescribeDataAgentThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeDataAgentThemeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeDataAgentThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDataAgentThemeResponseBody
	GetSuccess() *bool
}

type DescribeDataAgentThemeResponseBody struct {
	// The response struct.
	Data *DescribeDataAgentThemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request is abnormal.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
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
}

func (s DescribeDataAgentThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentThemeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentThemeResponseBody) GetData() *DescribeDataAgentThemeResponseBodyData {
	return s.Data
}

func (s *DescribeDataAgentThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDataAgentThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDataAgentThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataAgentThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataAgentThemeResponseBody) SetData(v *DescribeDataAgentThemeResponseBodyData) *DescribeDataAgentThemeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataAgentThemeResponseBody) SetErrorCode(v string) *DescribeDataAgentThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBody) SetErrorMessage(v string) *DescribeDataAgentThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBody) SetRequestId(v string) *DescribeDataAgentThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBody) SetSuccess(v bool) *DescribeDataAgentThemeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataAgentThemeResponseBodyData struct {
	// The common scenarios. Valid values: report, infographic, and others.
	//
	// example:
	//
	// custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2025-06-15T08:30:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The description of the theme.
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
	// The theme tracing information. This field is currently not enabled.
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
	// example:
	//
	// custom
	ThemeFrom *string `json:"ThemeFrom,omitempty" xml:"ThemeFrom,omitempty"`
	// The business ID of the theme.
	//
	// example:
	//
	// 0f8b2c1d***********9a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// The display name of the theme.
	//
	// example:
	//
	// weekly report
	ThemeName *string `json:"ThemeName,omitempty" xml:"ThemeName,omitempty"`
	// The theme stage. Valid values:
	//
	// - design: design.md only.
	//
	// - template: complete and renderable.
	//
	// example:
	//
	// template
	ThemeType *string `json:"ThemeType,omitempty" xml:"ThemeType,omitempty"`
}

func (s DescribeDataAgentThemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentThemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentThemeResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *DescribeDataAgentThemeResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeDataAgentThemeResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataAgentThemeResponseBodyData) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *DescribeDataAgentThemeResponseBodyData) GetReferTo() *string {
	return s.ReferTo
}

func (s *DescribeDataAgentThemeResponseBodyData) GetThemeFrom() *string {
	return s.ThemeFrom
}

func (s *DescribeDataAgentThemeResponseBodyData) GetThemeId() *string {
	return s.ThemeId
}

func (s *DescribeDataAgentThemeResponseBodyData) GetThemeName() *string {
	return s.ThemeName
}

func (s *DescribeDataAgentThemeResponseBodyData) GetThemeType() *string {
	return s.ThemeType
}

func (s *DescribeDataAgentThemeResponseBodyData) SetCategory(v string) *DescribeDataAgentThemeResponseBodyData {
	s.Category = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetCreatedAt(v string) *DescribeDataAgentThemeResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetDescription(v string) *DescribeDataAgentThemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetModifiedAt(v string) *DescribeDataAgentThemeResponseBodyData {
	s.ModifiedAt = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetReferTo(v string) *DescribeDataAgentThemeResponseBodyData {
	s.ReferTo = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetThemeFrom(v string) *DescribeDataAgentThemeResponseBodyData {
	s.ThemeFrom = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetThemeId(v string) *DescribeDataAgentThemeResponseBodyData {
	s.ThemeId = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetThemeName(v string) *DescribeDataAgentThemeResponseBodyData {
	s.ThemeName = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) SetThemeType(v string) *DescribeDataAgentThemeResponseBodyData {
	s.ThemeType = &v
	return s
}

func (s *DescribeDataAgentThemeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
