// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRoleType(v string) *CheckRoleRequest
	GetCheckRoleType() *string
	SetCrossAccountRoleName(v string) *CheckRoleRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountUserId(v int64) *CheckRoleRequest
	GetCrossAccountUserId() *int64
}

type CheckRoleRequest struct {
	// The type of the role. Valid values:
	//
	// 	- EcsRole: a role with the permissions to access Elastic Compute Service (ECS) resources
	//
	// 	- CsgRole: a role with the permissions to perform Cloud Storage Gateway (CSG) backup
	//
	// 	- NasRole: a role with the permissions to perform NAS backup
	//
	// 	- OssRole: a role with the permissions to perform Object Storage Service (OSS) backup
	//
	// 	- UdmRole: a role with the permissions to perform ECS instance backup
	//
	// 	- VMwareLocalRole: a role with the permissions to back up on-premises VMware virtual machines (VMs)
	//
	// 	- VMwareCloudRole: a role with the permissions to back up VMs deployed on Alibaba Cloud VMware Service (ACVS)
	//
	// 	- EcsBackupRole: a role with the permissions to perform ECS backup
	//
	// 	- OtsRole: a role with the permissions to perform Tablestore backup
	//
	// 	- CrossAccountRole: a role with the permissions to perform cross-account backup
	//
	// example:
	//
	// OssRole
	CheckRoleType *string `json:"CheckRoleType,omitempty" xml:"CheckRoleType,omitempty"`
	// The name of the Resource Access Management (RAM) role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s CheckRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckRoleRequest) GetCheckRoleType() *string {
	return s.CheckRoleType
}

func (s *CheckRoleRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CheckRoleRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CheckRoleRequest) SetCheckRoleType(v string) *CheckRoleRequest {
	s.CheckRoleType = &v
	return s
}

func (s *CheckRoleRequest) SetCrossAccountRoleName(v string) *CheckRoleRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CheckRoleRequest) SetCrossAccountUserId(v int64) *CheckRoleRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CheckRoleRequest) Validate() error {
	return dara.Validate(s)
}
