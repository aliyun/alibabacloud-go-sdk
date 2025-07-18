// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIPArrayListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceIPArrayName(v string) *DescribeDBInstanceIPArrayListRequest
	GetDBInstanceIPArrayName() *string
	SetDBInstanceId(v string) *DescribeDBInstanceIPArrayListRequest
	GetDBInstanceId() *string
	SetResourceGroupId(v string) *DescribeDBInstanceIPArrayListRequest
	GetResourceGroupId() *string
}

type DescribeDBInstanceIPArrayListRequest struct {
	// The name of the IP address whitelist. If you do not specify this parameter, the default whitelist is queried.
	//
	// >  Each instance supports up to 50 IP address whitelists.
	//
	// example:
	//
	// Default
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the resource group to which the instance belongs. For information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceIPArrayListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListRequest) GetDBInstanceIPArrayName() *string {
	return s.DBInstanceIPArrayName
}

func (s *DescribeDBInstanceIPArrayListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceIPArrayListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceIPArrayListRequest) SetDBInstanceIPArrayName(v string) *DescribeDBInstanceIPArrayListRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) SetDBInstanceId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) SetResourceGroupId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) Validate() error {
	return dara.Validate(s)
}
