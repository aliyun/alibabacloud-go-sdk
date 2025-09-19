// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallUniBackupAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *UninstallUniBackupAgentRequest
	GetPolicyId() *int64
}

type UninstallUniBackupAgentRequest struct {
	// The ID of the anti-ransomware policy.
	//
	// > You can call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to query the IDs of anti-ransomware policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s UninstallUniBackupAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallUniBackupAgentRequest) GoString() string {
	return s.String()
}

func (s *UninstallUniBackupAgentRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *UninstallUniBackupAgentRequest) SetPolicyId(v int64) *UninstallUniBackupAgentRequest {
	s.PolicyId = &v
	return s
}

func (s *UninstallUniBackupAgentRequest) Validate() error {
	return dara.Validate(s)
}
