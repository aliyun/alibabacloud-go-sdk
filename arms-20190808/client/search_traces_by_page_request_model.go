// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *SearchTracesByPageRequest
	GetEndTime() *int64
	SetExclusionFilters(v []*SearchTracesByPageRequestExclusionFilters) *SearchTracesByPageRequest
	GetExclusionFilters() []*SearchTracesByPageRequestExclusionFilters
	SetIsError(v bool) *SearchTracesByPageRequest
	GetIsError() *bool
	SetMinDuration(v int64) *SearchTracesByPageRequest
	GetMinDuration() *int64
	SetOperationName(v string) *SearchTracesByPageRequest
	GetOperationName() *string
	SetPageNumber(v int32) *SearchTracesByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchTracesByPageRequest
	GetPageSize() *int32
	SetPid(v string) *SearchTracesByPageRequest
	GetPid() *string
	SetRegionId(v string) *SearchTracesByPageRequest
	GetRegionId() *string
	SetReverse(v bool) *SearchTracesByPageRequest
	GetReverse() *bool
	SetServiceIp(v string) *SearchTracesByPageRequest
	GetServiceIp() *string
	SetServiceName(v string) *SearchTracesByPageRequest
	GetServiceName() *string
	SetStartTime(v int64) *SearchTracesByPageRequest
	GetStartTime() *int64
	SetTags(v []*SearchTracesByPageRequestTags) *SearchTracesByPageRequest
	GetTags() []*SearchTracesByPageRequestTags
}

type SearchTracesByPageRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595210400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter conditions.
	ExclusionFilters []*SearchTracesByPageRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
	// Specifies whether to include the traces of abnormal calls.
	//
	// 	- `true`: No
	//
	// 	- `false` (default): Yes
	//
	// example:
	//
	// false
	IsError *bool `json:"IsError,omitempty" xml:"IsError,omitempty"`
	// The minimum amount of time consumed by traces. Unit: milliseconds.
	//
	// example:
	//
	// 2
	MinDuration *int64 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	// The name of the traced span.
	//
	// example:
	//
	// /demo/queryNotExistDB/11
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The application ID.
	//
	// example:
	//
	// b590lhguqs@9781be0f44dXXXX
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to sort the query results in chronological order or reverse chronological order. Default value: `false`.
	//
	// 	- `true`: sorts the query results in reverse chronological order.
	//
	// 	- `false`: sorts the query results in chronological order.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The IP address of the host where the application resides.
	//
	// example:
	//
	// 172.20.XX.XX
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// arms-k8s-demo-subcomponent
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595174400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of tags.
	Tags []*SearchTracesByPageRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchTracesByPageRequest) GetExclusionFilters() []*SearchTracesByPageRequestExclusionFilters {
	return s.ExclusionFilters
}

func (s *SearchTracesByPageRequest) GetIsError() *bool {
	return s.IsError
}

func (s *SearchTracesByPageRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *SearchTracesByPageRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTracesByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTracesByPageRequest) GetPid() *string {
	return s.Pid
}

func (s *SearchTracesByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTracesByPageRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *SearchTracesByPageRequest) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *SearchTracesByPageRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesByPageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchTracesByPageRequest) GetTags() []*SearchTracesByPageRequestTags {
	return s.Tags
}

func (s *SearchTracesByPageRequest) SetEndTime(v int64) *SearchTracesByPageRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetExclusionFilters(v []*SearchTracesByPageRequestExclusionFilters) *SearchTracesByPageRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesByPageRequest) SetIsError(v bool) *SearchTracesByPageRequest {
	s.IsError = &v
	return s
}

func (s *SearchTracesByPageRequest) SetMinDuration(v int64) *SearchTracesByPageRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOperationName(v string) *SearchTracesByPageRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageNumber(v int32) *SearchTracesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageSize(v int32) *SearchTracesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPid(v string) *SearchTracesByPageRequest {
	s.Pid = &v
	return s
}

func (s *SearchTracesByPageRequest) SetRegionId(v string) *SearchTracesByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesByPageRequest) SetReverse(v bool) *SearchTracesByPageRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceIp(v string) *SearchTracesByPageRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceName(v string) *SearchTracesByPageRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetStartTime(v int64) *SearchTracesByPageRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetTags(v []*SearchTracesByPageRequestTags) *SearchTracesByPageRequest {
	s.Tags = v
	return s
}

func (s *SearchTracesByPageRequest) Validate() error {
	if s.ExclusionFilters != nil {
		for _, item := range s.ExclusionFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTracesByPageRequestExclusionFilters struct {
	// The key that is used to filter the query results.
	//
	// example:
	//
	// http.status_code
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the key that is used to filter the query results.
	//
	// example:
	//
	// 404
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesByPageRequestExclusionFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequestExclusionFilters) GetKey() *string {
	return s.Key
}

func (s *SearchTracesByPageRequestExclusionFilters) GetValue() *string {
	return s.Value
}

func (s *SearchTracesByPageRequestExclusionFilters) SetKey(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesByPageRequestExclusionFilters) SetValue(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Value = &v
	return s
}

func (s *SearchTracesByPageRequestExclusionFilters) Validate() error {
	return dara.Validate(s)
}

type SearchTracesByPageRequestTags struct {
	// The key of the tag. The following system preset fields are provided:
	//
	// 	- traceId: the ID of the trace.
	//
	// 	- serverApp: the name of the server application.
	//
	// 	- clientApp: the name of the client application.
	//
	// 	- service: the name of the interface.
	//
	// 	- rpc: the type of the call.
	//
	// 	- msOfSpan: the duration exceeds a specific value.
	//
	// 	- clientIp: the IP address of the client.
	//
	// 	- serverIp: the IP address of the server.
	//
	// 	- isError: specifies whether the call is abnormal.
	//
	// example:
	//
	// http.status_code
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesByPageRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageRequestTags) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequestTags) GetKey() *string {
	return s.Key
}

func (s *SearchTracesByPageRequestTags) GetValue() *string {
	return s.Value
}

func (s *SearchTracesByPageRequestTags) SetKey(v string) *SearchTracesByPageRequestTags {
	s.Key = &v
	return s
}

func (s *SearchTracesByPageRequestTags) SetValue(v string) *SearchTracesByPageRequestTags {
	s.Value = &v
	return s
}

func (s *SearchTracesByPageRequestTags) Validate() error {
	return dara.Validate(s)
}
