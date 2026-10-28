// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeInstanceStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRealtimeInstanceStatsResponseBody
	GetCode() *string
	SetData(v *ListRealtimeInstanceStatsResponseBodyData) *ListRealtimeInstanceStatsResponseBody
	GetData() *ListRealtimeInstanceStatsResponseBodyData
	SetHttpStatusCode(v int32) *ListRealtimeInstanceStatsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRealtimeInstanceStatsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListRealtimeInstanceStatsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListRealtimeInstanceStatsResponseBody
	GetRequestId() *string
}

type ListRealtimeInstanceStatsResponseBody struct {
	// The internal error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The call records.
	Data *ListRealtimeInstanceStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRealtimeInstanceStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeInstanceStatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeInstanceStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRealtimeInstanceStatsResponseBody) GetData() *ListRealtimeInstanceStatsResponseBodyData {
	return s.Data
}

func (s *ListRealtimeInstanceStatsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRealtimeInstanceStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRealtimeInstanceStatsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListRealtimeInstanceStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRealtimeInstanceStatsResponseBody) SetCode(v string) *ListRealtimeInstanceStatsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBody) SetData(v *ListRealtimeInstanceStatsResponseBodyData) *ListRealtimeInstanceStatsResponseBody {
	s.Data = v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBody) SetHttpStatusCode(v int32) *ListRealtimeInstanceStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBody) SetMessage(v string) *ListRealtimeInstanceStatsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBody) SetParams(v []*string) *ListRealtimeInstanceStatsResponseBody {
	s.Params = v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBody) SetRequestId(v string) *ListRealtimeInstanceStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRealtimeInstanceStatsResponseBodyData struct {
	// The list of real-time statistics for instances.
	List []*ListRealtimeInstanceStatsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRealtimeInstanceStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeInstanceStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRealtimeInstanceStatsResponseBodyData) GetList() []*ListRealtimeInstanceStatsResponseBodyDataList {
	return s.List
}

func (s *ListRealtimeInstanceStatsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRealtimeInstanceStatsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRealtimeInstanceStatsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRealtimeInstanceStatsResponseBodyData) SetList(v []*ListRealtimeInstanceStatsResponseBodyDataList) *ListRealtimeInstanceStatsResponseBodyData {
	s.List = v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyData) SetPageNumber(v int32) *ListRealtimeInstanceStatsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyData) SetPageSize(v int32) *ListRealtimeInstanceStatsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyData) SetTotalCount(v int32) *ListRealtimeInstanceStatsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRealtimeInstanceStatsResponseBodyDataList struct {
	// The configured concurrency.
	//
	// example:
	//
	// 12
	ConfiguredConcurrency *int64 `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 88b64330a8d34e63b60c9d272f8b3950
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time of the statistics.
	//
	// example:
	//
	// 1776837256000
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// The used concurrency.
	//
	// example:
	//
	// 10
	UsedConcurrency *int64 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s ListRealtimeInstanceStatsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeInstanceStatsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) GetConfiguredConcurrency() *int64 {
	return s.ConfiguredConcurrency
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) GetUsedConcurrency() *int64 {
	return s.UsedConcurrency
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) SetConfiguredConcurrency(v int64) *ListRealtimeInstanceStatsResponseBodyDataList {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) SetInstanceId(v string) *ListRealtimeInstanceStatsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) SetStatsTime(v int64) *ListRealtimeInstanceStatsResponseBodyDataList {
	s.StatsTime = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) SetUsedConcurrency(v int64) *ListRealtimeInstanceStatsResponseBodyDataList {
	s.UsedConcurrency = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
