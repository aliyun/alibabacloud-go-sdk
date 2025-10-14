// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *StopDataImportTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *StopDataImportTaskRequest
	GetPageSize() *int32
	SetRegionId(v string) *StopDataImportTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *StopDataImportTaskRequest
	GetSlinkTaskId() *string
}

type StopDataImportTaskRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s StopDataImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDataImportTaskRequest) GoString() string {
	return s.String()
}

func (s *StopDataImportTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *StopDataImportTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *StopDataImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDataImportTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *StopDataImportTaskRequest) SetPageNumber(v int32) *StopDataImportTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *StopDataImportTaskRequest) SetPageSize(v int32) *StopDataImportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *StopDataImportTaskRequest) SetRegionId(v string) *StopDataImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopDataImportTaskRequest) SetSlinkTaskId(v string) *StopDataImportTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *StopDataImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
