// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupMachineStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DescribeBackupMachineStatusRequest
	GetPolicyId() *int64
	SetPolicyVersion(v string) *DescribeBackupMachineStatusRequest
	GetPolicyVersion() *string
	SetUuid(v string) *DescribeBackupMachineStatusRequest
	GetUuid() *string
}

type DescribeBackupMachineStatusRequest struct {
	// The ID of the anti-ransomware policy.
	//
	// >  You can call the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the IDs of anti-ransomware policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The version of the anti-ransomware policy. Valid values:
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeBackupPolicy](~~DescribeBackupPolicy~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// eb2c782e-64f2-4590-a86c-d90164df****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeBackupMachineStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupMachineStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupMachineStatusRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeBackupMachineStatusRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *DescribeBackupMachineStatusRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeBackupMachineStatusRequest) SetPolicyId(v int64) *DescribeBackupMachineStatusRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeBackupMachineStatusRequest) SetPolicyVersion(v string) *DescribeBackupMachineStatusRequest {
	s.PolicyVersion = &v
	return s
}

func (s *DescribeBackupMachineStatusRequest) SetUuid(v string) *DescribeBackupMachineStatusRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeBackupMachineStatusRequest) Validate() error {
	return dara.Validate(s)
}
