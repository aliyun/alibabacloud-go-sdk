// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBVersionInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceMode(v string) *DescribeDBVersionInfosRequest
	GetDBInstanceMode() *string
	SetDBVersion(v string) *DescribeDBVersionInfosRequest
	GetDBVersion() *string
	SetOwnerId(v int64) *DescribeDBVersionInfosRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBVersionInfosRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBVersionInfosRequest
	GetResourceGroupId() *string
}

type DescribeDBVersionInfosRequest struct {
	// The resource type of the instance. Valid values:
	//
	// 	- **StorageElastic**: elastic storage mode.
	//
	// 	- **Serverless**: Serverless mode.
	//
	// example:
	//
	// StorageElastic
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// The minor version number that does not include the prefix.
	//
	// example:
	//
	// 6.3.10.20
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBVersionInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBVersionInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBVersionInfosRequest) GetDBInstanceMode() *string {
	return s.DBInstanceMode
}

func (s *DescribeDBVersionInfosRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBVersionInfosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBVersionInfosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBVersionInfosRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBVersionInfosRequest) SetDBInstanceMode(v string) *DescribeDBVersionInfosRequest {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDBVersionInfosRequest) SetDBVersion(v string) *DescribeDBVersionInfosRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBVersionInfosRequest) SetOwnerId(v int64) *DescribeDBVersionInfosRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBVersionInfosRequest) SetRegionId(v string) *DescribeDBVersionInfosRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBVersionInfosRequest) SetResourceGroupId(v string) *DescribeDBVersionInfosRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBVersionInfosRequest) Validate() error {
	return dara.Validate(s)
}
