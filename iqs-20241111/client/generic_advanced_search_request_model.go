// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenericAdvancedSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndustry(v string) *GenericAdvancedSearchRequest
	GetIndustry() *string
	SetQuery(v string) *GenericAdvancedSearchRequest
	GetQuery() *string
	SetSessionId(v string) *GenericAdvancedSearchRequest
	GetSessionId() *string
	SetTimeRange(v string) *GenericAdvancedSearchRequest
	GetTimeRange() *string
}

type GenericAdvancedSearchRequest struct {
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// job-4065bee3-e7aa-49fc-aad2-a8e3a7fd6acd
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GenericAdvancedSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s GenericAdvancedSearchRequest) GoString() string {
	return s.String()
}

func (s *GenericAdvancedSearchRequest) GetIndustry() *string {
	return s.Industry
}

func (s *GenericAdvancedSearchRequest) GetQuery() *string {
	return s.Query
}

func (s *GenericAdvancedSearchRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenericAdvancedSearchRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *GenericAdvancedSearchRequest) SetIndustry(v string) *GenericAdvancedSearchRequest {
	s.Industry = &v
	return s
}

func (s *GenericAdvancedSearchRequest) SetQuery(v string) *GenericAdvancedSearchRequest {
	s.Query = &v
	return s
}

func (s *GenericAdvancedSearchRequest) SetSessionId(v string) *GenericAdvancedSearchRequest {
	s.SessionId = &v
	return s
}

func (s *GenericAdvancedSearchRequest) SetTimeRange(v string) *GenericAdvancedSearchRequest {
	s.TimeRange = &v
	return s
}

func (s *GenericAdvancedSearchRequest) Validate() error {
	return dara.Validate(s)
}
