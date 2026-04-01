// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceMetricsRequest
	GetDBInstanceName() *string
	SetResourceGroupId(v string) *DescribeDBInstanceMetricsRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceMetricsRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceMetricsRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1s1j103lo6****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMetricsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceMetricsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceMetricsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceMetricsRequest) SetDBInstanceName(v string) *DescribeDBInstanceMetricsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceMetricsRequest) SetResourceGroupId(v string) *DescribeDBInstanceMetricsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceMetricsRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceMetricsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceMetricsRequest) Validate() error {
	return dara.Validate(s)
}
