// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeAttributes interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskEncrypted(v bool) *NodeAttributes
	GetDataDiskEncrypted() *bool
	SetDataDiskKMSKeyId(v string) *NodeAttributes
	GetDataDiskKMSKeyId() *string
	SetKeyPairName(v string) *NodeAttributes
	GetKeyPairName() *string
	SetMasterRootPassword(v string) *NodeAttributes
	GetMasterRootPassword() *string
	SetRamRole(v string) *NodeAttributes
	GetRamRole() *string
	SetSecurityGroupId(v string) *NodeAttributes
	GetSecurityGroupId() *string
	SetSystemDiskEncrypted(v bool) *NodeAttributes
	GetSystemDiskEncrypted() *bool
	SetSystemDiskKMSKeyId(v string) *NodeAttributes
	GetSystemDiskKMSKeyId() *string
	SetVpcId(v string) *NodeAttributes
	GetVpcId() *string
	SetZoneId(v string) *NodeAttributes
	GetZoneId() *string
}

type NodeAttributes struct {
	// Specifies whether to enable encryption for the data disk. For more information, see <props="china">[Encrypt a data disk](https://help.aliyun.com/zh/ecs/user-guide/encrypt-a-data-disk)<props="intl">[Encrypt a data disk](https://www.alibabacloud.com/help/en/ecs/user-guide/encryption-overview). Valid values:
	//
	// - true: Enables encryption.
	//
	// - false (default): Disables encryption.
	//
	// example:
	//
	// false
	DataDiskEncrypted *bool `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	// The ID of the KMS key for the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	DataDiskKMSKeyId *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	// The SSH key pair for logging on to the ECS instances.
	//
	// example:
	//
	// emr_login
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password of the root user for the master node. This parameter is left empty in the response of an API call.
	//
	// example:
	//
	// EAQ#86****
	MasterRootPassword *string `json:"MasterRootPassword,omitempty" xml:"MasterRootPassword,omitempty"`
	// The RAM role that is attached to the ECS instances to access other cloud resources.
	//
	// Default value: AliyunECSInstanceForEMRRole.
	//
	// example:
	//
	// AliyunECSInstanceForEMRRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the security group. EMR supports only basic security groups and does not support enterprise security groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-hp3abbae8lb6lmb1****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Specifies whether to enable disk encryption for the system disk. Valid values:
	//
	// - true: Enables encryption.
	//
	// - false (default): Disables encryption.
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"SystemDiskEncrypted,omitempty" xml:"SystemDiskEncrypted,omitempty"`
	// The ID of the KMS key.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKMSKeyId *string `json:"SystemDiskKMSKeyId,omitempty" xml:"SystemDiskKMSKeyId,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1tgey2p0ytxmdo5****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s NodeAttributes) String() string {
	return dara.Prettify(s)
}

func (s NodeAttributes) GoString() string {
	return s.String()
}

func (s *NodeAttributes) GetDataDiskEncrypted() *bool {
	return s.DataDiskEncrypted
}

func (s *NodeAttributes) GetDataDiskKMSKeyId() *string {
	return s.DataDiskKMSKeyId
}

func (s *NodeAttributes) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *NodeAttributes) GetMasterRootPassword() *string {
	return s.MasterRootPassword
}

func (s *NodeAttributes) GetRamRole() *string {
	return s.RamRole
}

func (s *NodeAttributes) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *NodeAttributes) GetSystemDiskEncrypted() *bool {
	return s.SystemDiskEncrypted
}

func (s *NodeAttributes) GetSystemDiskKMSKeyId() *string {
	return s.SystemDiskKMSKeyId
}

func (s *NodeAttributes) GetVpcId() *string {
	return s.VpcId
}

func (s *NodeAttributes) GetZoneId() *string {
	return s.ZoneId
}

func (s *NodeAttributes) SetDataDiskEncrypted(v bool) *NodeAttributes {
	s.DataDiskEncrypted = &v
	return s
}

func (s *NodeAttributes) SetDataDiskKMSKeyId(v string) *NodeAttributes {
	s.DataDiskKMSKeyId = &v
	return s
}

func (s *NodeAttributes) SetKeyPairName(v string) *NodeAttributes {
	s.KeyPairName = &v
	return s
}

func (s *NodeAttributes) SetMasterRootPassword(v string) *NodeAttributes {
	s.MasterRootPassword = &v
	return s
}

func (s *NodeAttributes) SetRamRole(v string) *NodeAttributes {
	s.RamRole = &v
	return s
}

func (s *NodeAttributes) SetSecurityGroupId(v string) *NodeAttributes {
	s.SecurityGroupId = &v
	return s
}

func (s *NodeAttributes) SetSystemDiskEncrypted(v bool) *NodeAttributes {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *NodeAttributes) SetSystemDiskKMSKeyId(v string) *NodeAttributes {
	s.SystemDiskKMSKeyId = &v
	return s
}

func (s *NodeAttributes) SetVpcId(v string) *NodeAttributes {
	s.VpcId = &v
	return s
}

func (s *NodeAttributes) SetZoneId(v string) *NodeAttributes {
	s.ZoneId = &v
	return s
}

func (s *NodeAttributes) Validate() error {
	return dara.Validate(s)
}
