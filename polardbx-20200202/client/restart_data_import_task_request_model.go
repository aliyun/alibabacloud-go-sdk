// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *RestartDataImportTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *RestartDataImportTaskRequest
	GetPageSize() *int32
	SetRegionId(v string) *RestartDataImportTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *RestartDataImportTaskRequest
	GetSlinkTaskId() *string
}

type RestartDataImportTaskRequest struct {
	// The page number. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: ***30*****50*****100**. Default value: **30**.
	//
	// example:
	//
	// 0
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the instance resides. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the destination task.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s RestartDataImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDataImportTaskRequest) GoString() string {
	return s.String()
}

func (s *RestartDataImportTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RestartDataImportTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RestartDataImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartDataImportTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *RestartDataImportTaskRequest) SetPageNumber(v int32) *RestartDataImportTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *RestartDataImportTaskRequest) SetPageSize(v int32) *RestartDataImportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *RestartDataImportTaskRequest) SetRegionId(v string) *RestartDataImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDataImportTaskRequest) SetSlinkTaskId(v string) *RestartDataImportTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *RestartDataImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
