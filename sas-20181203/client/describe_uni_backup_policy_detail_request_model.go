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
	// The ID of the anti-ransomware policy.
	//
	// > You can call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to query the IDs of anti-ransomware policies.
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
