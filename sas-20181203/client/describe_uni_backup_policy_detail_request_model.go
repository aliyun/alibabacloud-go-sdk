// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupPolicyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DescribeUniBackupPolicyDetailRequest
	GetPolicyId() *string
}

type DescribeUniBackupPolicyDetailRequest struct {
	// The ID of the anti-ransomware backup policy for databases.
	//
	// >Call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to obtain this parameter.
	//
	// If DescribeUniBackupPolicies returns an empty list, activate the anti-ransomware feature of Security Center first, make sure that the Security Center agent is installed on the ECS instance and the database has been discovered, and then call CreateUniBackupPolicy to create a backup policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeUniBackupPolicyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPolicyDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPolicyDetailRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeUniBackupPolicyDetailRequest) SetPolicyId(v string) *DescribeUniBackupPolicyDetailRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailRequest) Validate() error {
	return dara.Validate(s)
}
