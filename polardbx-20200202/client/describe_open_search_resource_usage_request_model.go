// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchResourceUsageRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenSearchResourceUsageRequest
	GetRegionId() *string
}

type DescribeOpenSearchResourceUsageRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpenSearchResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchResourceUsageRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchResourceUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchResourceUsageRequest) SetDBInstanceName(v string) *DescribeOpenSearchResourceUsageRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageRequest) SetRegionId(v string) *DescribeOpenSearchResourceUsageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
