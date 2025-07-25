// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIspFlushCacheTasks(v []*DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) *DescribeIspFlushCacheTasksResponseBody
	GetIspFlushCacheTasks() []*DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks
	SetPageNumber(v int32) *DescribeIspFlushCacheTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIspFlushCacheTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIspFlushCacheTasksResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeIspFlushCacheTasksResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeIspFlushCacheTasksResponseBody
	GetTotalPages() *int32
}

type DescribeIspFlushCacheTasksResponseBody struct {
	IspFlushCacheTasks []*DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks `json:"IspFlushCacheTasks,omitempty" xml:"IspFlushCacheTasks,omitempty" type:"Repeated"`
	PageNumber         *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItems         *int32                                                      `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages         *int32                                                      `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeIspFlushCacheTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTasksResponseBody) GetIspFlushCacheTasks() []*DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	return s.IspFlushCacheTasks
}

func (s *DescribeIspFlushCacheTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIspFlushCacheTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIspFlushCacheTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIspFlushCacheTasksResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeIspFlushCacheTasksResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeIspFlushCacheTasksResponseBody) SetIspFlushCacheTasks(v []*DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) *DescribeIspFlushCacheTasksResponseBody {
	s.IspFlushCacheTasks = v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBody) SetPageNumber(v int32) *DescribeIspFlushCacheTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBody) SetPageSize(v int32) *DescribeIspFlushCacheTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBody) SetRequestId(v string) *DescribeIspFlushCacheTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBody) SetTotalItems(v int32) *DescribeIspFlushCacheTasksResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBody) SetTotalPages(v int32) *DescribeIspFlushCacheTasksResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Isp             *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus      *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetIsp() *string {
	return s.Isp
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetCreateTime(v string) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetCreateTimestamp(v int64) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetDomainName(v string) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.DomainName = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetInstanceId(v string) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.InstanceId = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetInstanceName(v string) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.InstanceName = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetIsp(v string) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.Isp = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetTaskId(v string) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) SetTaskStatus(v string) *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks {
	s.TaskStatus = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponseBodyIspFlushCacheTasks) Validate() error {
	return dara.Validate(s)
}
