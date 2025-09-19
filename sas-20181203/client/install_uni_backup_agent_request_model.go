// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallUniBackupAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *InstallUniBackupAgentRequest
	GetPolicyId() *int64
}

type InstallUniBackupAgentRequest struct {
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

func (s InstallUniBackupAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallUniBackupAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallUniBackupAgentRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *InstallUniBackupAgentRequest) SetPolicyId(v int64) *InstallUniBackupAgentRequest {
	s.PolicyId = &v
	return s
}

func (s *InstallUniBackupAgentRequest) Validate() error {
	return dara.Validate(s)
}
