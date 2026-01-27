// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeDBTablesRecoveryStateRequest
	GetClusterName() *string
	SetInstanceId(v string) *DescribeDBTablesRecoveryStateRequest
	GetInstanceId() *string
	SetRegionCode(v string) *DescribeDBTablesRecoveryStateRequest
	GetRegionCode() *string
}

type DescribeDBTablesRecoveryStateRequest struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode  *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDBTablesRecoveryStateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryStateRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryStateRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDBTablesRecoveryStateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDBTablesRecoveryStateRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDBTablesRecoveryStateRequest) SetClusterName(v string) *DescribeDBTablesRecoveryStateRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateRequest) SetInstanceId(v string) *DescribeDBTablesRecoveryStateRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateRequest) SetRegionCode(v string) *DescribeDBTablesRecoveryStateRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateRequest) Validate() error {
	return dara.Validate(s)
}
