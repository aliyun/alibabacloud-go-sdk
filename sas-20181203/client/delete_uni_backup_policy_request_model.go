// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUniBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DeleteUniBackupPolicyRequest
	GetPolicyId() *int64
	SetPolicyIds(v string) *DeleteUniBackupPolicyRequest
	GetPolicyIds() *string
}

type DeleteUniBackupPolicyRequest struct {
	// The ID of the anti-ransomware policy.
	//
	// >  You can call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to query the IDs of anti-ransomware policies. You must specify at least one of the PolicyId parameter and the **PolicyIds*	- parameter.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The IDs of anti-ransomware policies.
	//
	// >  You can call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to query the IDs of anti-ransomware policies. You must specify at least one of the **PolicyId*	- parameter and the PolicyIds parameter.
	//
	// example:
	//
	// 123,124
	PolicyIds *string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty"`
}

func (s DeleteUniBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUniBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteUniBackupPolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DeleteUniBackupPolicyRequest) GetPolicyIds() *string {
	return s.PolicyIds
}

func (s *DeleteUniBackupPolicyRequest) SetPolicyId(v int64) *DeleteUniBackupPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteUniBackupPolicyRequest) SetPolicyIds(v string) *DeleteUniBackupPolicyRequest {
	s.PolicyIds = &v
	return s
}

func (s *DeleteUniBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
