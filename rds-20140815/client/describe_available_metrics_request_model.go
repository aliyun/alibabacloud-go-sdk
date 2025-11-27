// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeAvailableMetricsRequest
	GetDBInstanceName() *string
	SetResourceGroupId(v string) *DescribeAvailableMetricsRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *DescribeAvailableMetricsRequest
	GetResourceOwnerId() *int64
}

type DescribeAvailableMetricsRequest struct {
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

func (s DescribeAvailableMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMetricsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAvailableMetricsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAvailableMetricsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableMetricsRequest) SetDBInstanceName(v string) *DescribeAvailableMetricsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAvailableMetricsRequest) SetResourceGroupId(v string) *DescribeAvailableMetricsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAvailableMetricsRequest) SetResourceOwnerId(v int64) *DescribeAvailableMetricsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableMetricsRequest) Validate() error {
	return dara.Validate(s)
}
