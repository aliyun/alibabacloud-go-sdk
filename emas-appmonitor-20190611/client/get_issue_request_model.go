// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetIssueRequest
	GetAppKey() *int64
	SetBizModule(v string) *GetIssueRequest
	GetBizModule() *string
	SetDigestHash(v string) *GetIssueRequest
	GetDigestHash() *string
	SetFilter(v *GetIssueRequestFilter) *GetIssueRequest
	GetFilter() *GetIssueRequestFilter
	SetOs(v string) *GetIssueRequest
	GetOs() *string
	SetTimeRange(v *GetIssueRequestTimeRange) *GetIssueRequest
	GetTimeRange() *GetIssueRequestTimeRange
}

type GetIssueRequest struct {
	// AppKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// h5Resource
	BizModule *string `json:"BizModule,omitempty" xml:"BizModule,omitempty"`
	// example:
	//
	// 2963475858785631
	DigestHash *string                `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	Filter     *GetIssueRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// h5
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	TimeRange *GetIssueRequestTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
}

func (s GetIssueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIssueRequest) GoString() string {
	return s.String()
}

func (s *GetIssueRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetIssueRequest) GetBizModule() *string {
	return s.BizModule
}

func (s *GetIssueRequest) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetIssueRequest) GetFilter() *GetIssueRequestFilter {
	return s.Filter
}

func (s *GetIssueRequest) GetOs() *string {
	return s.Os
}

func (s *GetIssueRequest) GetTimeRange() *GetIssueRequestTimeRange {
	return s.TimeRange
}

func (s *GetIssueRequest) SetAppKey(v int64) *GetIssueRequest {
	s.AppKey = &v
	return s
}

func (s *GetIssueRequest) SetBizModule(v string) *GetIssueRequest {
	s.BizModule = &v
	return s
}

func (s *GetIssueRequest) SetDigestHash(v string) *GetIssueRequest {
	s.DigestHash = &v
	return s
}

func (s *GetIssueRequest) SetFilter(v *GetIssueRequestFilter) *GetIssueRequest {
	s.Filter = v
	return s
}

func (s *GetIssueRequest) SetOs(v string) *GetIssueRequest {
	s.Os = &v
	return s
}

func (s *GetIssueRequest) SetTimeRange(v *GetIssueRequestTimeRange) *GetIssueRequest {
	s.TimeRange = v
	return s
}

func (s *GetIssueRequest) Validate() error {
	return dara.Validate(s)
}

type GetIssueRequestFilter struct {
	// example:
	//
	// MySQL_IOPS
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Equal
	Operator   *string       `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SubFilters []*string     `json:"SubFilters,omitempty" xml:"SubFilters,omitempty" type:"Repeated"`
	Values     []interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetIssueRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetIssueRequestFilter) GoString() string {
	return s.String()
}

func (s *GetIssueRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetIssueRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *GetIssueRequestFilter) GetSubFilters() []*string {
	return s.SubFilters
}

func (s *GetIssueRequestFilter) GetValues() []interface{} {
	return s.Values
}

func (s *GetIssueRequestFilter) SetKey(v string) *GetIssueRequestFilter {
	s.Key = &v
	return s
}

func (s *GetIssueRequestFilter) SetOperator(v string) *GetIssueRequestFilter {
	s.Operator = &v
	return s
}

func (s *GetIssueRequestFilter) SetSubFilters(v []*string) *GetIssueRequestFilter {
	s.SubFilters = v
	return s
}

func (s *GetIssueRequestFilter) SetValues(v []interface{}) *GetIssueRequestFilter {
	s.Values = v
	return s
}

func (s *GetIssueRequestFilter) Validate() error {
	return dara.Validate(s)
}

type GetIssueRequestTimeRange struct {
	// example:
	//
	// 2024-08-23T02:12:00Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Host
	Granularity *int32 `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// minute
	GranularityUnit *string `json:"GranularityUnit,omitempty" xml:"GranularityUnit,omitempty"`
	// example:
	//
	// 2024-12-18 00:00:00
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetIssueRequestTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetIssueRequestTimeRange) GoString() string {
	return s.String()
}

func (s *GetIssueRequestTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetIssueRequestTimeRange) GetGranularity() *int32 {
	return s.Granularity
}

func (s *GetIssueRequestTimeRange) GetGranularityUnit() *string {
	return s.GranularityUnit
}

func (s *GetIssueRequestTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetIssueRequestTimeRange) SetEndTime(v int64) *GetIssueRequestTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetIssueRequestTimeRange) SetGranularity(v int32) *GetIssueRequestTimeRange {
	s.Granularity = &v
	return s
}

func (s *GetIssueRequestTimeRange) SetGranularityUnit(v string) *GetIssueRequestTimeRange {
	s.GranularityUnit = &v
	return s
}

func (s *GetIssueRequestTimeRange) SetStartTime(v int64) *GetIssueRequestTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetIssueRequestTimeRange) Validate() error {
	return dara.Validate(s)
}
