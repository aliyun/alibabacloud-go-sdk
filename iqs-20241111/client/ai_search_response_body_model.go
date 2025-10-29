// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *AiSearchResponseBodyHeader) *AiSearchResponseBody
	GetHeader() *AiSearchResponseBodyHeader
	SetPayload(v string) *AiSearchResponseBody
	GetPayload() *string
	SetRequestId(v string) *AiSearchResponseBody
	GetRequestId() *string
}

type AiSearchResponseBody struct {
	Header *AiSearchResponseBodyHeader `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	// example:
	//
	// {"header":{"eventId":"6f617de0-204f-406f-a9be-34779c06d498","event":"on_common_search_start","responseTime":120},"payload":"","requestId":"715d01a0-de7e-42c3-abca-b901fcd79b39"}
	Payload *string `json:"payload,omitempty" xml:"payload,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AiSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AiSearchResponseBody) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBody) GetHeader() *AiSearchResponseBodyHeader {
	return s.Header
}

func (s *AiSearchResponseBody) GetPayload() *string {
	return s.Payload
}

func (s *AiSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AiSearchResponseBody) SetHeader(v *AiSearchResponseBodyHeader) *AiSearchResponseBody {
	s.Header = v
	return s
}

func (s *AiSearchResponseBody) SetPayload(v string) *AiSearchResponseBody {
	s.Payload = &v
	return s
}

func (s *AiSearchResponseBody) SetRequestId(v string) *AiSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *AiSearchResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiSearchResponseBodyHeader struct {
	// example:
	//
	// on_common_search_end
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 988021f0-951a-43d0-ba4d-785359e7e7be
	EventId      *string                                 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	QueryContext *AiSearchResponseBodyHeaderQueryContext `json:"queryContext,omitempty" xml:"queryContext,omitempty" type:"Struct"`
	// example:
	//
	// 1293
	ResponseTime *int64 `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
}

func (s AiSearchResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s AiSearchResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *AiSearchResponseBodyHeader) GetEventId() *string {
	return s.EventId
}

func (s *AiSearchResponseBodyHeader) GetQueryContext() *AiSearchResponseBodyHeaderQueryContext {
	return s.QueryContext
}

func (s *AiSearchResponseBodyHeader) GetResponseTime() *int64 {
	return s.ResponseTime
}

func (s *AiSearchResponseBodyHeader) SetEvent(v string) *AiSearchResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *AiSearchResponseBodyHeader) SetEventId(v string) *AiSearchResponseBodyHeader {
	s.EventId = &v
	return s
}

func (s *AiSearchResponseBodyHeader) SetQueryContext(v *AiSearchResponseBodyHeaderQueryContext) *AiSearchResponseBodyHeader {
	s.QueryContext = v
	return s
}

func (s *AiSearchResponseBodyHeader) SetResponseTime(v int64) *AiSearchResponseBodyHeader {
	s.ResponseTime = &v
	return s
}

func (s *AiSearchResponseBodyHeader) Validate() error {
	if s.QueryContext != nil {
		if err := s.QueryContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiSearchResponseBodyHeaderQueryContext struct {
	OriginalQuery *AiSearchResponseBodyHeaderQueryContextOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty" type:"Struct"`
	Rewrite       *AiSearchResponseBodyHeaderQueryContextRewrite       `json:"rewrite,omitempty" xml:"rewrite,omitempty" type:"Struct"`
}

func (s AiSearchResponseBodyHeaderQueryContext) String() string {
	return dara.Prettify(s)
}

func (s AiSearchResponseBodyHeaderQueryContext) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeaderQueryContext) GetOriginalQuery() *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	return s.OriginalQuery
}

func (s *AiSearchResponseBodyHeaderQueryContext) GetRewrite() *AiSearchResponseBodyHeaderQueryContextRewrite {
	return s.Rewrite
}

func (s *AiSearchResponseBodyHeaderQueryContext) SetOriginalQuery(v *AiSearchResponseBodyHeaderQueryContextOriginalQuery) *AiSearchResponseBodyHeaderQueryContext {
	s.OriginalQuery = v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContext) SetRewrite(v *AiSearchResponseBodyHeaderQueryContextRewrite) *AiSearchResponseBodyHeaderQueryContext {
	s.Rewrite = v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContext) Validate() error {
	if s.OriginalQuery != nil {
		if err := s.OriginalQuery.Validate(); err != nil {
			return err
		}
	}
	if s.Rewrite != nil {
		if err := s.Rewrite.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiSearchResponseBodyHeaderQueryContextOriginalQuery struct {
	Industry  *string `json:"industry,omitempty" xml:"industry,omitempty"`
	Page      *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AiSearchResponseBodyHeaderQueryContextOriginalQuery) String() string {
	return dara.Prettify(s)
}

func (s AiSearchResponseBodyHeaderQueryContextOriginalQuery) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) GetIndustry() *string {
	return s.Industry
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) GetPage() *int32 {
	return s.Page
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) GetQuery() *string {
	return s.Query
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) GetTimeRange() *string {
	return s.TimeRange
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetIndustry(v string) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.Industry = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetPage(v int32) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.Page = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetQuery(v string) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.Query = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetTimeRange(v string) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.TimeRange = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) Validate() error {
	return dara.Validate(s)
}

type AiSearchResponseBodyHeaderQueryContextRewrite struct {
	Enabled   *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AiSearchResponseBodyHeaderQueryContextRewrite) String() string {
	return dara.Prettify(s)
}

func (s AiSearchResponseBodyHeaderQueryContextRewrite) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeaderQueryContextRewrite) GetEnabled() *bool {
	return s.Enabled
}

func (s *AiSearchResponseBodyHeaderQueryContextRewrite) GetTimeRange() *string {
	return s.TimeRange
}

func (s *AiSearchResponseBodyHeaderQueryContextRewrite) SetEnabled(v bool) *AiSearchResponseBodyHeaderQueryContextRewrite {
	s.Enabled = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextRewrite) SetTimeRange(v string) *AiSearchResponseBodyHeaderQueryContextRewrite {
	s.TimeRange = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextRewrite) Validate() error {
	return dara.Validate(s)
}
