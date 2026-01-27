// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeBackupPolicyRequest
	GetInstanceName() *string
	SetRegionCode(v string) *DescribeBackupPolicyRequest
	GetRegionCode() *string
}

type DescribeBackupPolicyRequest struct {
	// The ID of the PolarDB for MySQL cluster.
	//
	// example:
	//
	// pc-2ze3nrr64c5******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region in which the backup set resides.
	//
	// example:
	//
	// cn-beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBackupPolicyRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeBackupPolicyRequest) SetInstanceName(v string) *DescribeBackupPolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetRegionCode(v string) *DescribeBackupPolicyRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
