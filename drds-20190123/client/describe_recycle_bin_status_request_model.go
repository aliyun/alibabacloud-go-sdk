// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecycleBinStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeRecycleBinStatusRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeRecycleBinStatusRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribeRecycleBinStatusRequest
	GetRegionId() *string
}

type DescribeRecycleBinStatusRequest struct {
	// The name of the database that is created in the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRecycleBinStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecycleBinStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeRecycleBinStatusRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeRecycleBinStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRecycleBinStatusRequest) SetDbName(v string) *DescribeRecycleBinStatusRequest {
	s.DbName = &v
	return s
}

func (s *DescribeRecycleBinStatusRequest) SetDrdsInstanceId(v string) *DescribeRecycleBinStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRecycleBinStatusRequest) SetRegionId(v string) *DescribeRecycleBinStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRecycleBinStatusRequest) Validate() error {
	return dara.Validate(s)
}
