// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullToPushListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstUrl(v string) *DescribeLivePullToPushListRequest
	GetDstUrl() *string
	SetOwnerId(v int64) *DescribeLivePullToPushListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLivePullToPushListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLivePullToPushListRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeLivePullToPushListRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLivePullToPushListRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeLivePullToPushListRequest
	GetTaskId() *string
	SetTaskName(v string) *DescribeLivePullToPushListRequest
	GetTaskName() *string
}

type DescribeLivePullToPushListRequest struct {
	// The destination URL. Fuzzy search is performed based on the destination URL.
	//
	// example:
	//
	// rtmp://qd
	DstUrl  *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// >  The value must be greater than 0 and not greater than the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// >  Valid values: [1,100]. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region of the live center. Valid values:
	//
	// 	- ap-southeast-1: Singapore
	//
	// 	- ap-southeast-5: Indonesia (Jakarta)
	//
	// 	- cn-beijing: China (Beijing)
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID. Fuzzy search is performed based on the task ID.
	//
	// >  The ID can be up to 55 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// 861009
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name. Fuzzy search is performed based on the task name.
	//
	// example:
	//
	// task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeLivePullToPushListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushListRequest) GetDstUrl() *string {
	return s.DstUrl
}

func (s *DescribeLivePullToPushListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePullToPushListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLivePullToPushListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLivePullToPushListRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLivePullToPushListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePullToPushListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLivePullToPushListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeLivePullToPushListRequest) SetDstUrl(v string) *DescribeLivePullToPushListRequest {
	s.DstUrl = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) SetOwnerId(v int64) *DescribeLivePullToPushListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) SetPageNumber(v int32) *DescribeLivePullToPushListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) SetPageSize(v int32) *DescribeLivePullToPushListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) SetRegion(v string) *DescribeLivePullToPushListRequest {
	s.Region = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) SetRegionId(v string) *DescribeLivePullToPushListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) SetTaskId(v string) *DescribeLivePullToPushListRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) SetTaskName(v string) *DescribeLivePullToPushListRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeLivePullToPushListRequest) Validate() error {
	return dara.Validate(s)
}
