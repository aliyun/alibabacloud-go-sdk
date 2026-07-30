// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseOrgsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseOrgsRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *DescribeLangfuseOrgsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLangfuseOrgsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLangfuseOrgsRequest
	GetRegionId() *string
}

type DescribeLangfuseOrgsRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLangfuseOrgsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseOrgsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLangfuseOrgsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLangfuseOrgsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseOrgsRequest) SetDBInstanceId(v string) *DescribeLangfuseOrgsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseOrgsRequest) SetPageNumber(v int32) *DescribeLangfuseOrgsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseOrgsRequest) SetPageSize(v int32) *DescribeLangfuseOrgsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseOrgsRequest) SetRegionId(v string) *DescribeLangfuseOrgsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseOrgsRequest) Validate() error {
	return dara.Validate(s)
}
