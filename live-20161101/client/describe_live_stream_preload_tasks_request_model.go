// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamPreloadTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamPreloadTasksRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamPreloadTasksRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamPreloadTasksRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveStreamPreloadTasksRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamPreloadTasksRequest
	GetPageSize() *int32
	SetPlayUrl(v string) *DescribeLiveStreamPreloadTasksRequest
	GetPlayUrl() *string
	SetRegionId(v string) *DescribeLiveStreamPreloadTasksRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamPreloadTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeLiveStreamPreloadTasksRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeLiveStreamPreloadTasksRequest
	GetTaskId() *string
}

type DescribeLiveStreamPreloadTasksRequest struct {
	// The streaming domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2016-06-30T19:00:00Z. The interval between the start time and end time cannot exceed three days.
	//
	// example:
	//
	// 2016-06-30T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The streaming URL. You can specify up to 100 streaming URLs in a request. Separate multiple streaming URLs with commas (,).
	PlayUrl  *string `json:"PlayUrl,omitempty" xml:"PlayUrl,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2016-06-29T19:00:00Z. You can query only data in the previous three days.
	//
	// example:
	//
	// 2016-06-29T19:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the prefetch task. Valid values:
	//
	// 	- complete
	//
	// 	- pending
	//
	// 	- preloading
	//
	// 	- failed
	//
	// example:
	//
	// complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the prefetch task. The task ID is returned when you call the [SetLiveStreamPreloadTasks](https://help.aliyun.com/document_detail/2519938.html) operation to configure the prefetch task.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeLiveStreamPreloadTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPreloadTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetPlayUrl() *string {
	return s.PlayUrl
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveStreamPreloadTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetDomainName(v string) *DescribeLiveStreamPreloadTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetEndTime(v string) *DescribeLiveStreamPreloadTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetOwnerId(v int64) *DescribeLiveStreamPreloadTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetPageNum(v int32) *DescribeLiveStreamPreloadTasksRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetPageSize(v int32) *DescribeLiveStreamPreloadTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetPlayUrl(v string) *DescribeLiveStreamPreloadTasksRequest {
	s.PlayUrl = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetRegionId(v string) *DescribeLiveStreamPreloadTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetStartTime(v string) *DescribeLiveStreamPreloadTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetStatus(v string) *DescribeLiveStreamPreloadTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) SetTaskId(v string) *DescribeLiveStreamPreloadTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksRequest) Validate() error {
	return dara.Validate(s)
}
