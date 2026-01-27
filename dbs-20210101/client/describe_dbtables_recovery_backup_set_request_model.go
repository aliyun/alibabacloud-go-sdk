// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryBackupSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeDBTablesRecoveryBackupSetRequest
	GetClusterName() *string
	SetInstanceId(v string) *DescribeDBTablesRecoveryBackupSetRequest
	GetInstanceId() *string
	SetRegionCode(v string) *DescribeDBTablesRecoveryBackupSetRequest
	GetRegionCode() *string
}

type DescribeDBTablesRecoveryBackupSetRequest struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode  *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDBTablesRecoveryBackupSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryBackupSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) SetClusterName(v string) *DescribeDBTablesRecoveryBackupSetRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) SetInstanceId(v string) *DescribeDBTablesRecoveryBackupSetRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) SetRegionCode(v string) *DescribeDBTablesRecoveryBackupSetRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) Validate() error {
	return dara.Validate(s)
}
