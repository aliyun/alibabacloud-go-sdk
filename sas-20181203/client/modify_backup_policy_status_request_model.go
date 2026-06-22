// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyBackupPolicyStatusRequest
	GetId() *int64
	SetPolicyVersion(v string) *ModifyBackupPolicyStatusRequest
	GetPolicyVersion() *string
	SetStatus(v string) *ModifyBackupPolicyStatusRequest
	GetStatus() *string
}

type ModifyBackupPolicyStatusRequest struct {
	// The ID of the anti-ransomware policy that you want to enable or shutdown.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30490
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Deprecated
	//
	// The version of the policy that you want to enable or disable. Set the value to **2.0.0**.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// Specifies whether to enable or disable the anti-ransomware policy. Valid values:
	//
	// - **enabled**: Enables the anti-ransomware policy. After the policy is enabled, anti-ransomware can protect your server data by backing up data on specified servers based on your anti-ransomware policy.
	//
	// - **disabled**: Disables the anti-ransomware policy. After the policy is disabled, running backup tasks are stopped.
	//
	// > Backup tasks consume network bandwidth. Enable the policy during off-peak hours to execute data backup tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyBackupPolicyStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyStatusRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyBackupPolicyStatusRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *ModifyBackupPolicyStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyBackupPolicyStatusRequest) SetId(v int64) *ModifyBackupPolicyStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyBackupPolicyStatusRequest) SetPolicyVersion(v string) *ModifyBackupPolicyStatusRequest {
	s.PolicyVersion = &v
	return s
}

func (s *ModifyBackupPolicyStatusRequest) SetStatus(v string) *ModifyBackupPolicyStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyBackupPolicyStatusRequest) Validate() error {
	return dara.Validate(s)
}
