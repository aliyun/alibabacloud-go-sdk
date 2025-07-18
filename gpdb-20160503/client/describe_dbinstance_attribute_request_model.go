// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest
	GetResourceGroupId() *string
}

type DescribeDBInstanceAttributeRequest struct {
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) interface to view the instance IDs of all AnalyticDB for PostgreSQL instances in the target region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp13ue79qk8y1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is deprecated and should not be passed.
	//
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceAttributeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
