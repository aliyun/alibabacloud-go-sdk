// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenericSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedParams(v map[string]interface{}) *GenericSearchRequest
	GetAdvancedParams() map[string]interface{}
	SetEnableRerank(v bool) *GenericSearchRequest
	GetEnableRerank() *bool
	SetIndustry(v string) *GenericSearchRequest
	GetIndustry() *string
	SetPage(v int32) *GenericSearchRequest
	GetPage() *int32
	SetQuery(v string) *GenericSearchRequest
	GetQuery() *string
	SetReturnMainText(v bool) *GenericSearchRequest
	GetReturnMainText() *bool
	SetReturnMarkdownText(v bool) *GenericSearchRequest
	GetReturnMarkdownText() *bool
	SetReturnRichMainBody(v bool) *GenericSearchRequest
	GetReturnRichMainBody() *bool
	SetReturnSummary(v bool) *GenericSearchRequest
	GetReturnSummary() *bool
	SetSessionId(v string) *GenericSearchRequest
	GetSessionId() *string
	SetTimeRange(v string) *GenericSearchRequest
	GetTimeRange() *string
}

type GenericSearchRequest struct {
	AdvancedParams map[string]interface{} `json:"advancedParams,omitempty" xml:"advancedParams,omitempty"`
	EnableRerank   *bool                  `json:"enableRerank,omitempty" xml:"enableRerank,omitempty"`
	Industry       *string                `json:"industry,omitempty" xml:"industry,omitempty"`
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

func (s GenericSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s GenericSearchRequest) GoString() string {
	return s.String()
}

func (s *GenericSearchRequest) GetAdvancedParams() map[string]interface{} {
	return s.AdvancedParams
}

func (s *GenericSearchRequest) GetEnableRerank() *bool {
	return s.EnableRerank
}

func (s *GenericSearchRequest) GetIndustry() *string {
	return s.Industry
}

func (s *GenericSearchRequest) GetPage() *int32 {
	return s.Page
}

func (s *GenericSearchRequest) GetQuery() *string {
	return s.Query
}

func (s *GenericSearchRequest) GetReturnMainText() *bool {
	return s.ReturnMainText
}

func (s *GenericSearchRequest) GetReturnMarkdownText() *bool {
	return s.ReturnMarkdownText
}

func (s *GenericSearchRequest) GetReturnRichMainBody() *bool {
	return s.ReturnRichMainBody
}

func (s *GenericSearchRequest) GetReturnSummary() *bool {
	return s.ReturnSummary
}

func (s *GenericSearchRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenericSearchRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *GenericSearchRequest) SetAdvancedParams(v map[string]interface{}) *GenericSearchRequest {
	s.AdvancedParams = v
	return s
}

func (s *GenericSearchRequest) SetEnableRerank(v bool) *GenericSearchRequest {
	s.EnableRerank = &v
	return s
}

func (s *GenericSearchRequest) SetIndustry(v string) *GenericSearchRequest {
	s.Industry = &v
	return s
}

func (s *GenericSearchRequest) SetPage(v int32) *GenericSearchRequest {
	s.Page = &v
	return s
}

func (s *GenericSearchRequest) SetQuery(v string) *GenericSearchRequest {
	s.Query = &v
	return s
}

func (s *GenericSearchRequest) SetReturnMainText(v bool) *GenericSearchRequest {
	s.ReturnMainText = &v
	return s
}

func (s *GenericSearchRequest) SetReturnMarkdownText(v bool) *GenericSearchRequest {
	s.ReturnMarkdownText = &v
	return s
}

func (s *GenericSearchRequest) SetReturnRichMainBody(v bool) *GenericSearchRequest {
	s.ReturnRichMainBody = &v
	return s
}

func (s *GenericSearchRequest) SetReturnSummary(v bool) *GenericSearchRequest {
	s.ReturnSummary = &v
	return s
}

func (s *GenericSearchRequest) SetSessionId(v string) *GenericSearchRequest {
	s.SessionId = &v
	return s
}

func (s *GenericSearchRequest) SetTimeRange(v string) *GenericSearchRequest {
	s.TimeRange = &v
	return s
}

func (s *GenericSearchRequest) Validate() error {
	return dara.Validate(s)
}
