// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeRecordFileRequest
	GetDesktopId() *string
	SetEndTime(v string) *DescribeRecordFileRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeRecordFileRequest
	GetEndUserId() *string
	SetFileName(v string) *DescribeRecordFileRequest
	GetFileName() *string
	SetOrderBy(v string) *DescribeRecordFileRequest
	GetOrderBy() *string
	SetOrderSort(v string) *DescribeRecordFileRequest
	GetOrderSort() *string
	SetPageNumber(v int32) *DescribeRecordFileRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRecordFileRequest
	GetPageSize() *int32
	SetRecordType(v string) *DescribeRecordFileRequest
	GetRecordType() *string
	SetRegionId(v string) *DescribeRecordFileRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeRecordFileRequest
	GetResourceType() *string
	SetStartTime(v string) *DescribeRecordFileRequest
	GetStartTime() *string
	SetStatus(v int32) *DescribeRecordFileRequest
	GetStatus() *int32
}

type DescribeRecordFileRequest struct {
	// The Cloud Desktop ID.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The end time of the query. Format: yyyyMMddHHmmss.
	//
	// example:
	//
	// 20251218205715
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// Alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The name of the recording file.
	//
	// example:
	//
	// Task7
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The field used for sorting. If this parameter is not specified, the results are sorted by recording start time in descending order. Valid values:
	//
	// - startTime: recording start time.
	//
	// example:
	//
	// startTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The sort order. Valid values:
	//
	// - asc: ascending order.
	//
	// - desc: descending order.
	//
	// example:
	//
	// asc
	OrderSort *string `json:"OrderSort,omitempty" xml:"OrderSort,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of rows per page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the screen recording file. Valid values:
	//
	// - alltime: full-time recording.
	//
	// - period: interval recording.
	//
	// - event: event-based recording.
	//
	// - session: session recording.
	//
	// example:
	//
	// alltime
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time of the query. Format: yyyyMMddHHmmss.
	//
	// example:
	//
	// 20251218175715
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the screen recording file. Valid values:
	//
	// - 0: uploaded.
	//
	// - 1: uploading.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRecordFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordFileRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeRecordFileRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRecordFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeRecordFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribeRecordFileRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeRecordFileRequest) GetOrderSort() *string {
	return s.OrderSort
}

func (s *DescribeRecordFileRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRecordFileRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRecordFileRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *DescribeRecordFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRecordFileRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRecordFileRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRecordFileRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeRecordFileRequest) SetDesktopId(v string) *DescribeRecordFileRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeRecordFileRequest) SetEndTime(v string) *DescribeRecordFileRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordFileRequest) SetEndUserId(v string) *DescribeRecordFileRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeRecordFileRequest) SetFileName(v string) *DescribeRecordFileRequest {
	s.FileName = &v
	return s
}

func (s *DescribeRecordFileRequest) SetOrderBy(v string) *DescribeRecordFileRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeRecordFileRequest) SetOrderSort(v string) *DescribeRecordFileRequest {
	s.OrderSort = &v
	return s
}

func (s *DescribeRecordFileRequest) SetPageNumber(v int32) *DescribeRecordFileRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordFileRequest) SetPageSize(v int32) *DescribeRecordFileRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordFileRequest) SetRecordType(v string) *DescribeRecordFileRequest {
	s.RecordType = &v
	return s
}

func (s *DescribeRecordFileRequest) SetRegionId(v string) *DescribeRecordFileRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRecordFileRequest) SetResourceType(v string) *DescribeRecordFileRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRecordFileRequest) SetStartTime(v string) *DescribeRecordFileRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordFileRequest) SetStatus(v int32) *DescribeRecordFileRequest {
	s.Status = &v
	return s
}

func (s *DescribeRecordFileRequest) Validate() error {
	return dara.Validate(s)
}
