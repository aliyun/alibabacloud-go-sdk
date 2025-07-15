// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamPreloadTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *DescribeLiveStreamPreloadTasksResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamPreloadTasksResponseBody
	GetPageSize() *int32
	SetPreloadTasks(v *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks) *DescribeLiveStreamPreloadTasksResponseBody
	GetPreloadTasks() *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks
	SetRequestId(v string) *DescribeLiveStreamPreloadTasksResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveStreamPreloadTasksResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveStreamPreloadTasksResponseBody
	GetTotalPage() *int32
}

type DescribeLiveStreamPreloadTasksResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details of the prefetch task.
	PreloadTasks *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks `json:"PreloadTasks,omitempty" xml:"PreloadTasks,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E1564CBC-DCFE-5E1B-8B78-8DED9A39F334
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveStreamPreloadTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPreloadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) GetPreloadTasks() *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks {
	return s.PreloadTasks
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) SetPageNum(v int32) *DescribeLiveStreamPreloadTasksResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) SetPageSize(v int32) *DescribeLiveStreamPreloadTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) SetPreloadTasks(v *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks) *DescribeLiveStreamPreloadTasksResponseBody {
	s.PreloadTasks = v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) SetRequestId(v string) *DescribeLiveStreamPreloadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) SetTotalNum(v int32) *DescribeLiveStreamPreloadTasksResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) SetTotalPage(v int32) *DescribeLiveStreamPreloadTasksResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks struct {
	PreloadTask []*DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask `json:"PreloadTask,omitempty" xml:"PreloadTask,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks) GetPreloadTask() []*DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	return s.PreloadTask
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks) SetPreloadTask(v []*DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks {
	s.PreloadTask = v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask struct {
	// The acceleration region where the live content is prefetched. Valid values:
	//
	// 	- domestic: regions in the Chinese mainland.
	//
	// 	- overseas: regions outside the Chinese mainland.
	//
	// 	- global: regions in and outside the Chinese mainland.
	//
	// >  If this parameter is left empty, the acceleration region configured for the domain name is returned.
	//
	// example:
	//
	// domestic
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time when the prefetch task was created.
	//
	// example:
	//
	// 2016-06-29T19:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the prefetch task is successful. Valid values:
	//
	// 	- Successfully
	//
	// 	- InternalError
	//
	// example:
	//
	// Successfully
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The streaming domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The streaming URL.
	PlayUrl *string `json:"PlayUrl,omitempty" xml:"PlayUrl,omitempty"`
	// The time when the prefetch task ended.
	//
	// example:
	//
	// 2016-06-30T19:00:00Z
	PreloadedEndTime *string `json:"PreloadedEndTime,omitempty" xml:"PreloadedEndTime,omitempty"`
	// The time when the prefetch task started.
	//
	// example:
	//
	// 2016-06-29T19:00:00Z
	PreloadedStartTime *string `json:"PreloadedStartTime,omitempty" xml:"PreloadedStartTime,omitempty"`
	// The progress of the prefetch task.
	//
	// example:
	//
	// 100%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The status of the prefetch task. Valid values:
	//
	// 	- Success
	//
	// 	- Failed
	//
	// >  Success is returned only if the prefetch task is configured for all streaming URLs.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the prefetch task.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetArea() *string {
	return s.Area
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetPlayUrl() *string {
	return s.PlayUrl
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetPreloadedEndTime() *string {
	return s.PreloadedEndTime
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetPreloadedStartTime() *string {
	return s.PreloadedStartTime
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetProcess() *string {
	return s.Process
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetArea(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.Area = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetCreateTime(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetDescription(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.Description = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetDomainName(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetPlayUrl(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.PlayUrl = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetPreloadedEndTime(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.PreloadedEndTime = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetPreloadedStartTime(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.PreloadedStartTime = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetProcess(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.Process = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetStatus(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.Status = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) SetTaskId(v string) *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask {
	s.TaskId = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponseBodyPreloadTasksPreloadTask) Validate() error {
	return dara.Validate(s)
}
