// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetErrorsRequest
	GetAppKey() *int64
	SetBizModule(v string) *GetErrorsRequest
	GetBizModule() *string
	SetDigestHash(v string) *GetErrorsRequest
	GetDigestHash() *string
	SetFilter(v *GetErrorsRequestFilter) *GetErrorsRequest
	GetFilter() *GetErrorsRequestFilter
	SetOs(v string) *GetErrorsRequest
	GetOs() *string
	SetPageIndex(v int32) *GetErrorsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetErrorsRequest
	GetPageSize() *int32
	SetTimeRange(v *GetErrorsRequestTimeRange) *GetErrorsRequest
	GetTimeRange() *GetErrorsRequestTimeRange
	SetUtdid(v string) *GetErrorsRequest
	GetUtdid() *string
}

type GetErrorsRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crash
	BizModule *string `json:"BizModule,omitempty" xml:"BizModule,omitempty"`
	// example:
	//
	// -3481243636390427020
	DigestHash *string                 `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	Filter     *GetErrorsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TimeRange *GetErrorsRequestTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
	// utdid
	//
	// example:
	//
	// Z70g6V/MXJ8DABtD53eHzn4X
	Utdid *string `json:"Utdid,omitempty" xml:"Utdid,omitempty"`
}

func (s GetErrorsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsRequest) GoString() string {
	return s.String()
}

func (s *GetErrorsRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetErrorsRequest) GetBizModule() *string {
	return s.BizModule
}

func (s *GetErrorsRequest) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetErrorsRequest) GetFilter() *GetErrorsRequestFilter {
	return s.Filter
}

func (s *GetErrorsRequest) GetOs() *string {
	return s.Os
}

func (s *GetErrorsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetErrorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetErrorsRequest) GetTimeRange() *GetErrorsRequestTimeRange {
	return s.TimeRange
}

func (s *GetErrorsRequest) GetUtdid() *string {
	return s.Utdid
}

func (s *GetErrorsRequest) SetAppKey(v int64) *GetErrorsRequest {
	s.AppKey = &v
	return s
}

func (s *GetErrorsRequest) SetBizModule(v string) *GetErrorsRequest {
	s.BizModule = &v
	return s
}

func (s *GetErrorsRequest) SetDigestHash(v string) *GetErrorsRequest {
	s.DigestHash = &v
	return s
}

func (s *GetErrorsRequest) SetFilter(v *GetErrorsRequestFilter) *GetErrorsRequest {
	s.Filter = v
	return s
}

func (s *GetErrorsRequest) SetOs(v string) *GetErrorsRequest {
	s.Os = &v
	return s
}

func (s *GetErrorsRequest) SetPageIndex(v int32) *GetErrorsRequest {
	s.PageIndex = &v
	return s
}

func (s *GetErrorsRequest) SetPageSize(v int32) *GetErrorsRequest {
	s.PageSize = &v
	return s
}

func (s *GetErrorsRequest) SetTimeRange(v *GetErrorsRequestTimeRange) *GetErrorsRequest {
	s.TimeRange = v
	return s
}

func (s *GetErrorsRequest) SetUtdid(v string) *GetErrorsRequest {
	s.Utdid = &v
	return s
}

func (s *GetErrorsRequest) Validate() error {
	return dara.Validate(s)
}

type GetErrorsRequestFilter struct {
	// example:
	//
	// osVersion
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// and
	Operator   *string       `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SubFilters []*string     `json:"SubFilters,omitempty" xml:"SubFilters,omitempty" type:"Repeated"`
	Values     []interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetErrorsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetErrorsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetErrorsRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *GetErrorsRequestFilter) GetSubFilters() []*string {
	return s.SubFilters
}

func (s *GetErrorsRequestFilter) GetValues() []interface{} {
	return s.Values
}

func (s *GetErrorsRequestFilter) SetKey(v string) *GetErrorsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetErrorsRequestFilter) SetOperator(v string) *GetErrorsRequestFilter {
	s.Operator = &v
	return s
}

func (s *GetErrorsRequestFilter) SetSubFilters(v []*string) *GetErrorsRequestFilter {
	s.SubFilters = v
	return s
}

func (s *GetErrorsRequestFilter) SetValues(v []interface{}) *GetErrorsRequestFilter {
	s.Values = v
	return s
}

func (s *GetErrorsRequestFilter) Validate() error {
	return dara.Validate(s)
}

type GetErrorsRequestTimeRange struct {
	// This parameter is required.
	//
	// example:
	//
	// 1740499200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1739894400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetErrorsRequestTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsRequestTimeRange) GoString() string {
	return s.String()
}

func (s *GetErrorsRequestTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetErrorsRequestTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetErrorsRequestTimeRange) SetEndTime(v int64) *GetErrorsRequestTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetErrorsRequestTimeRange) SetStartTime(v int64) *GetErrorsRequestTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetErrorsRequestTimeRange) Validate() error {
	return dara.Validate(s)
}
