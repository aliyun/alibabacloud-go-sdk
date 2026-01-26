// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *SearchTracesRequest
	GetEndTime() *int64
	SetExclusionFilters(v []*SearchTracesRequestExclusionFilters) *SearchTracesRequest
	GetExclusionFilters() []*SearchTracesRequestExclusionFilters
	SetMinDuration(v int64) *SearchTracesRequest
	GetMinDuration() *int64
	SetOperationName(v string) *SearchTracesRequest
	GetOperationName() *string
	SetPid(v string) *SearchTracesRequest
	GetPid() *string
	SetRegionId(v string) *SearchTracesRequest
	GetRegionId() *string
	SetReverse(v bool) *SearchTracesRequest
	GetReverse() *bool
	SetServiceIp(v string) *SearchTracesRequest
	GetServiceIp() *string
	SetServiceName(v string) *SearchTracesRequest
	GetServiceName() *string
	SetStartTime(v int64) *SearchTracesRequest
	GetStartTime() *int64
	SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest
	GetTag() []*SearchTracesRequestTag
}

type SearchTracesRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595210400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter conditions.
	ExclusionFilters []*SearchTracesRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
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
	Tag []*SearchTracesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchTracesRequest) GetExclusionFilters() []*SearchTracesRequestExclusionFilters {
	return s.ExclusionFilters
}

func (s *SearchTracesRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *SearchTracesRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesRequest) GetPid() *string {
	return s.Pid
}

func (s *SearchTracesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTracesRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *SearchTracesRequest) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *SearchTracesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchTracesRequest) GetTag() []*SearchTracesRequestTag {
	return s.Tag
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetExclusionFilters(v []*SearchTracesRequestExclusionFilters) *SearchTracesRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetPid(v string) *SearchTracesRequest {
	s.Pid = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetReverse(v bool) *SearchTracesRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesRequest) SetServiceIp(v string) *SearchTracesRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

func (s *SearchTracesRequest) Validate() error {
	if s.ExclusionFilters != nil {
		for _, item := range s.ExclusionFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTracesRequestExclusionFilters struct {
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

func (s SearchTracesRequestExclusionFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestExclusionFilters) GetKey() *string {
	return s.Key
}

func (s *SearchTracesRequestExclusionFilters) GetValue() *string {
	return s.Value
}

func (s *SearchTracesRequestExclusionFilters) SetKey(v string) *SearchTracesRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestExclusionFilters) SetValue(v string) *SearchTracesRequestExclusionFilters {
	s.Value = &v
	return s
}

func (s *SearchTracesRequestExclusionFilters) Validate() error {
	return dara.Validate(s)
}

type SearchTracesRequestTag struct {
	// The tag key. The following system preset fields are provided:
	//
	// 	- serverApp: the name of the server application.
	//
	// 	- clientApp: the name of the client application.
	//
	// 	- service: the name of the operation.
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
	// 	- hasTprof: contains only thread profiling.
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

func (s SearchTracesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestTag) GetKey() *string {
	return s.Key
}

func (s *SearchTracesRequestTag) GetValue() *string {
	return s.Value
}

func (s *SearchTracesRequestTag) SetKey(v string) *SearchTracesRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestTag) SetValue(v string) *SearchTracesRequestTag {
	s.Value = &v
	return s
}

func (s *SearchTracesRequestTag) Validate() error {
	return dara.Validate(s)
}
