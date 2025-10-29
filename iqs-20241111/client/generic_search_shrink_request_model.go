// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenericSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedParamsShrink(v string) *GenericSearchShrinkRequest
	GetAdvancedParamsShrink() *string
	SetEnableRerank(v bool) *GenericSearchShrinkRequest
	GetEnableRerank() *bool
	SetIndustry(v string) *GenericSearchShrinkRequest
	GetIndustry() *string
	SetPage(v int32) *GenericSearchShrinkRequest
	GetPage() *int32
	SetQuery(v string) *GenericSearchShrinkRequest
	GetQuery() *string
	SetReturnMainText(v bool) *GenericSearchShrinkRequest
	GetReturnMainText() *bool
	SetReturnMarkdownText(v bool) *GenericSearchShrinkRequest
	GetReturnMarkdownText() *bool
	SetReturnRichMainBody(v bool) *GenericSearchShrinkRequest
	GetReturnRichMainBody() *bool
	SetReturnSummary(v bool) *GenericSearchShrinkRequest
	GetReturnSummary() *bool
	SetSessionId(v string) *GenericSearchShrinkRequest
	GetSessionId() *string
	SetTimeRange(v string) *GenericSearchShrinkRequest
	GetTimeRange() *string
}

type GenericSearchShrinkRequest struct {
	AdvancedParamsShrink *string `json:"advancedParams,omitempty" xml:"advancedParams,omitempty"`
	EnableRerank         *bool   `json:"enableRerank,omitempty" xml:"enableRerank,omitempty"`
	Industry             *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// This parameter is required.
	Query              *string `json:"query,omitempty" xml:"query,omitempty"`
	ReturnMainText     *bool   `json:"returnMainText,omitempty" xml:"returnMainText,omitempty"`
	ReturnMarkdownText *bool   `json:"returnMarkdownText,omitempty" xml:"returnMarkdownText,omitempty"`
	ReturnRichMainBody *bool   `json:"returnRichMainBody,omitempty" xml:"returnRichMainBody,omitempty"`
	ReturnSummary      *bool   `json:"returnSummary,omitempty" xml:"returnSummary,omitempty"`
	SessionId          *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GenericSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenericSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenericSearchShrinkRequest) GetAdvancedParamsShrink() *string {
	return s.AdvancedParamsShrink
}

func (s *GenericSearchShrinkRequest) GetEnableRerank() *bool {
	return s.EnableRerank
}

func (s *GenericSearchShrinkRequest) GetIndustry() *string {
	return s.Industry
}

func (s *GenericSearchShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *GenericSearchShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *GenericSearchShrinkRequest) GetReturnMainText() *bool {
	return s.ReturnMainText
}

func (s *GenericSearchShrinkRequest) GetReturnMarkdownText() *bool {
	return s.ReturnMarkdownText
}

func (s *GenericSearchShrinkRequest) GetReturnRichMainBody() *bool {
	return s.ReturnRichMainBody
}

func (s *GenericSearchShrinkRequest) GetReturnSummary() *bool {
	return s.ReturnSummary
}

func (s *GenericSearchShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenericSearchShrinkRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *GenericSearchShrinkRequest) SetAdvancedParamsShrink(v string) *GenericSearchShrinkRequest {
	s.AdvancedParamsShrink = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetEnableRerank(v bool) *GenericSearchShrinkRequest {
	s.EnableRerank = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetIndustry(v string) *GenericSearchShrinkRequest {
	s.Industry = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetPage(v int32) *GenericSearchShrinkRequest {
	s.Page = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetQuery(v string) *GenericSearchShrinkRequest {
	s.Query = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetReturnMainText(v bool) *GenericSearchShrinkRequest {
	s.ReturnMainText = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetReturnMarkdownText(v bool) *GenericSearchShrinkRequest {
	s.ReturnMarkdownText = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetReturnRichMainBody(v bool) *GenericSearchShrinkRequest {
	s.ReturnRichMainBody = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetReturnSummary(v bool) *GenericSearchShrinkRequest {
	s.ReturnSummary = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetSessionId(v string) *GenericSearchShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *GenericSearchShrinkRequest) SetTimeRange(v string) *GenericSearchShrinkRequest {
	s.TimeRange = &v
	return s
}

func (s *GenericSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
