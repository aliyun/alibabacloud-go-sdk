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
	// The role type. Valid values:
	//
	// - EcsRole: access permissions for ECS resources
	//
	// - CsgRole: permissions to back up Cloud Storage Gateway resources
	//
	// - NasRole: permissions to back up NAS resources
	//
	// - OssRole: permissions to back up OSS resources
	//
	// - UdmRole: permissions to back up entire ECS instances
	//
	// - VMwareLocalRole: permissions to back up on-premises VMware virtual machines
	//
	// - VMwareCloudRole: permissions to back up cloud-based VMware virtual machines
	//
	// - EcsBackupRole: permissions for ECS backup
	//
	// - OtsRole: permissions to back up OTS resources
	//
	// - CrossAccountRole: permissions for cross-account backup
	//
	// example:
	//
	// OssRole
	CheckRoleType *string `json:"CheckRoleType,omitempty" xml:"CheckRoleType,omitempty"`
	// The name of the RAM role created in the source account for cross-account backup managed by the current account.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The ID of the source account for cross-account backup managed by the current account.
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
