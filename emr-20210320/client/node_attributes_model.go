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
	// 是否启用云盘加密。取值范围：
	//
	// - true：启用加密。
	//
	// - false：不加密。
	//
	// 默认值：false，不加密
	//
	// example:
	//
	// false
	DataDiskEncrypted *bool `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	// KMS加密秘钥ID。
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	DataDiskKMSKeyId *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	// ECS ssh登录秘钥。
	//
	// example:
	//
	// emr_login
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// MASTER节点root密码。
	//
	// example:
	//
	// Adxefswfd****
	MasterRootPassword *string `json:"MasterRootPassword,omitempty" xml:"MasterRootPassword,omitempty"`
	// ECS访问资源绑定的角色。
	//
	// example:
	//
	// AliyunECSInstanceForEMRRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// 安全组ID。EMR只支持普通安全组，不支持企业安全组。
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-hp3abbae8lb6lmb1****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 是否启用云盘加密。取值范围：
	//
	// - true：启用加密。
	//
	// - false：不加密。
	//
	// 默认值：false，不加密
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"SystemDiskEncrypted,omitempty" xml:"SystemDiskEncrypted,omitempty"`
	// KMS加密秘钥ID。
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKMSKeyId *string `json:"SystemDiskKMSKeyId,omitempty" xml:"SystemDiskKMSKeyId,omitempty"`
	// 专有网络ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1tgey2p0ytxmdo5****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 可用区ID。
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
