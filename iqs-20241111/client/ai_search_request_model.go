// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndustry(v string) *AiSearchRequest
	GetIndustry() *string
	SetPage(v int32) *AiSearchRequest
	GetPage() *int32
	SetQuery(v string) *AiSearchRequest
	GetQuery() *string
	SetSessionId(v string) *AiSearchRequest
	GetSessionId() *string
	SetTimeRange(v string) *AiSearchRequest
	GetTimeRange() *string
}

type AiSearchRequest struct {
	// example:
	//
	// finance
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// {\\"total_count\\": 6851, \\"page_number\\": 54, \\"page_size\\": 100}
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// 17dc8bcd-f34a-46d1-a7a3-0fa3d1ce3824
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AiSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s AiSearchRequest) GoString() string {
	return s.String()
}

func (s *AiSearchRequest) GetIndustry() *string {
	return s.Industry
}

func (s *AiSearchRequest) GetPage() *int32 {
	return s.Page
}

func (s *AiSearchRequest) GetQuery() *string {
	return s.Query
}

func (s *AiSearchRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AiSearchRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *AiSearchRequest) SetIndustry(v string) *AiSearchRequest {
	s.Industry = &v
	return s
}

func (s *AiSearchRequest) SetPage(v int32) *AiSearchRequest {
	s.Page = &v
	return s
}

func (s *AiSearchRequest) SetQuery(v string) *AiSearchRequest {
	s.Query = &v
	return s
}

func (s *AiSearchRequest) SetSessionId(v string) *AiSearchRequest {
	s.SessionId = &v
	return s
}

func (s *AiSearchRequest) SetTimeRange(v string) *AiSearchRequest {
	s.TimeRange = &v
	return s
}

func (s *AiSearchRequest) Validate() error {
	return dara.Validate(s)
}
