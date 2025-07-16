// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinaryLogListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeBinaryLogListRequest
	GetDBInstanceName() *string
	SetEndTime(v string) *DescribeBinaryLogListRequest
	GetEndTime() *string
	SetInstanceName(v string) *DescribeBinaryLogListRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *DescribeBinaryLogListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBinaryLogListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBinaryLogListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeBinaryLogListRequest
	GetStartTime() *string
}

type DescribeBinaryLogListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hz1fds
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-09 10:27:46
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-09 10:27:46
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBinaryLogListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinaryLogListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeBinaryLogListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBinaryLogListRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBinaryLogListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBinaryLogListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBinaryLogListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBinaryLogListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBinaryLogListRequest) SetDBInstanceName(v string) *DescribeBinaryLogListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetEndTime(v string) *DescribeBinaryLogListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetInstanceName(v string) *DescribeBinaryLogListRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetPageNumber(v int32) *DescribeBinaryLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetPageSize(v int32) *DescribeBinaryLogListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetRegionId(v string) *DescribeBinaryLogListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetStartTime(v string) *DescribeBinaryLogListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBinaryLogListRequest) Validate() error {
	return dara.Validate(s)
}
