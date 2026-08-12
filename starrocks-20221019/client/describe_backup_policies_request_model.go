// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeBackupPoliciesRequest
	GetInstanceId() *string
	SetPolicyId(v string) *DescribeBackupPoliciesRequest
	GetPolicyId() *string
	SetRegionId(v string) *DescribeBackupPoliciesRequest
	GetRegionId() *string
}

type DescribeBackupPoliciesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// bk-policy-26ec5bc0ea67b5ef
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBackupPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupPoliciesRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeBackupPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupPoliciesRequest) SetInstanceId(v string) *DescribeBackupPoliciesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) SetPolicyId(v string) *DescribeBackupPoliciesRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) SetRegionId(v string) *DescribeBackupPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
