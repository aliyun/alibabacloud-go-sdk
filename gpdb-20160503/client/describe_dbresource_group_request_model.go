// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBResourceGroupRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBResourceGroupRequest
	GetOwnerId() *int64
	SetResourceGroupName(v string) *DescribeDBResourceGroupRequest
	GetResourceGroupName() *string
}

type DescribeDBResourceGroupRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// > If this parameter is omitted, details for all resource groups are returned. Otherwise, only details for the specified resource group are returned.
	//
	// example:
	//
	// testgroup
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBResourceGroupRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeDBResourceGroupRequest) SetDBInstanceId(v string) *DescribeDBResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetOwnerId(v int64) *DescribeDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetResourceGroupName(v string) *DescribeDBResourceGroupRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
