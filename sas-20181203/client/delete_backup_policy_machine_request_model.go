// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DeleteBackupPolicyMachineRequest
	GetPolicyId() *int64
	SetPolicyVersion(v string) *DeleteBackupPolicyMachineRequest
	GetPolicyVersion() *string
	SetUuid(v string) *DeleteBackupPolicyMachineRequest
	GetUuid() *string
	SetUuidList(v []*string) *DeleteBackupPolicyMachineRequest
	GetUuidList() []*string
}

type DeleteBackupPolicyMachineRequest struct {
	// The ID of the anti-ransomware policy.
	//
	// > You can call the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the IDs of anti-ransomware policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The version of the anti-ransomware policy. Valid values:
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// >  You can call the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the versions of anti-ransomware policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The UUID of the server to which the anti-ransomware policy is applied.
	//
	// > You must specify at least one of the `UuidList` and `Uuid` parameters.
	//
	// example:
	//
	// 083036e9-8411-4a9d-83af-9acbd****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The UUIDs of the servers to which the anti-ransomware policy is applied.
	//
	// > You must specify at least one of the `UuidList` and `Uuid` parameters.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s DeleteBackupPolicyMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyMachineRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyMachineRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DeleteBackupPolicyMachineRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *DeleteBackupPolicyMachineRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteBackupPolicyMachineRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *DeleteBackupPolicyMachineRequest) SetPolicyId(v int64) *DeleteBackupPolicyMachineRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteBackupPolicyMachineRequest) SetPolicyVersion(v string) *DeleteBackupPolicyMachineRequest {
	s.PolicyVersion = &v
	return s
}

func (s *DeleteBackupPolicyMachineRequest) SetUuid(v string) *DeleteBackupPolicyMachineRequest {
	s.Uuid = &v
	return s
}

func (s *DeleteBackupPolicyMachineRequest) SetUuidList(v []*string) *DeleteBackupPolicyMachineRequest {
	s.UuidList = v
	return s
}

func (s *DeleteBackupPolicyMachineRequest) Validate() error {
	return dara.Validate(s)
}
