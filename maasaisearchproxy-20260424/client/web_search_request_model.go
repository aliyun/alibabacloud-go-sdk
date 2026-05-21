// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *WebSearchRequest
	GetEndTime() *string
	SetExcludeDomain(v []*string) *WebSearchRequest
	GetExcludeDomain() []*string
	SetIncludeDomain(v []*string) *WebSearchRequest
	GetIncludeDomain() []*string
	SetLimit(v int32) *WebSearchRequest
	GetLimit() *int32
	SetQuery(v string) *WebSearchRequest
	GetQuery() *string
	SetRegion(v string) *WebSearchRequest
	GetRegion() *string
	SetSearchType(v string) *WebSearchRequest
	GetSearchType() *string
	SetStartTime(v string) *WebSearchRequest
	GetStartTime() *string
}

type WebSearchRequest struct {
	// example:
	//
	// 1754973000000
	EndTime       *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExcludeDomain []*string `json:"excludeDomain,omitempty" xml:"excludeDomain,omitempty" type:"Repeated"`
	IncludeDomain []*string `json:"includeDomain,omitempty" xml:"includeDomain,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// {\\"entityFilter\\":{\\"domain\\":\\"synthetics\\",\\"filters\\":[],\\"type\\":\\"synthetics.task\\"},\\"metric\\":\\"availability\\",\\"metricSet\\":\\"synthetics.metric.task\\",\\"type\\":\\"METRIC_SET_QUERY\\"}
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region     *string `json:"region,omitempty" xml:"region,omitempty"`
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	// example:
	//
	// 2026-03-06 10:04:45
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s WebSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s WebSearchRequest) GoString() string {
	return s.String()
}

func (s *WebSearchRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *WebSearchRequest) GetExcludeDomain() []*string {
	return s.ExcludeDomain
}

func (s *WebSearchRequest) GetIncludeDomain() []*string {
	return s.IncludeDomain
}

func (s *WebSearchRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *WebSearchRequest) GetQuery() *string {
	return s.Query
}

func (s *WebSearchRequest) GetRegion() *string {
	return s.Region
}

func (s *WebSearchRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *WebSearchRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *WebSearchRequest) SetEndTime(v string) *WebSearchRequest {
	s.EndTime = &v
	return s
}

func (s *WebSearchRequest) SetExcludeDomain(v []*string) *WebSearchRequest {
	s.ExcludeDomain = v
	return s
}

func (s *WebSearchRequest) SetIncludeDomain(v []*string) *WebSearchRequest {
	s.IncludeDomain = v
	return s
}

func (s *WebSearchRequest) SetLimit(v int32) *WebSearchRequest {
	s.Limit = &v
	return s
}

func (s *WebSearchRequest) SetQuery(v string) *WebSearchRequest {
	s.Query = &v
	return s
}

func (s *WebSearchRequest) SetRegion(v string) *WebSearchRequest {
	s.Region = &v
	return s
}

func (s *WebSearchRequest) SetSearchType(v string) *WebSearchRequest {
	s.SearchType = &v
	return s
}

func (s *WebSearchRequest) SetStartTime(v string) *WebSearchRequest {
	s.StartTime = &v
	return s
}

func (s *WebSearchRequest) Validate() error {
	return dara.Validate(s)
}
