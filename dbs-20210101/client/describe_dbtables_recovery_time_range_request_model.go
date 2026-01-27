// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryTimeRangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeDBTablesRecoveryTimeRangeRequest
	GetClusterName() *string
	SetInstanceId(v string) *DescribeDBTablesRecoveryTimeRangeRequest
	GetInstanceId() *string
	SetRegionCode(v string) *DescribeDBTablesRecoveryTimeRangeRequest
	GetRegionCode() *string
}

type DescribeDBTablesRecoveryTimeRangeRequest struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode  *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDBTablesRecoveryTimeRangeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryTimeRangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) SetClusterName(v string) *DescribeDBTablesRecoveryTimeRangeRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) SetInstanceId(v string) *DescribeDBTablesRecoveryTimeRangeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) SetRegionCode(v string) *DescribeDBTablesRecoveryTimeRangeRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) Validate() error {
	return dara.Validate(s)
}
