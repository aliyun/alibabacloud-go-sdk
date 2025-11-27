// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceReplicationRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeDBInstanceReplicationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstanceReplicationRequest
	GetResourceGroupId() *string
}

type DescribeDBInstanceReplicationRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can leave this parameter empty.
	//
	// example:
	//
	// rg-acfmx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceReplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceReplicationRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceReplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceReplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceReplicationRequest) SetDBInstanceId(v string) *DescribeDBInstanceReplicationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceReplicationRequest) SetRegionId(v string) *DescribeDBInstanceReplicationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceReplicationRequest) SetResourceGroupId(v string) *DescribeDBInstanceReplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceReplicationRequest) Validate() error {
	return dara.Validate(s)
}
