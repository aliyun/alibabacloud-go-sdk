// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeBackupPolicyRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeBackupPolicyRequest
	GetRegionId() *string
}

type DescribeBackupPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupPolicyRequest) SetDBInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetRegionId(v string) *DescribeBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
