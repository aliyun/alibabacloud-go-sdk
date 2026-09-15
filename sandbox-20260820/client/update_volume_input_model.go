// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVolumeInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticBucketVolumeConfig(v *AgenticBucketVolumeConfig) *UpdateVolumeInput
	GetAgenticBucketVolumeConfig() *AgenticBucketVolumeConfig
	SetAgenticFSVolumeConfig(v *UpdateVolumeInputAgenticFSVolumeConfig) *UpdateVolumeInput
	GetAgenticFSVolumeConfig() *UpdateVolumeInputAgenticFSVolumeConfig
	SetJuiceFSVolumeConfig(v *JuiceFSVolumeConfig) *UpdateVolumeInput
	GetJuiceFSVolumeConfig() *JuiceFSVolumeConfig
	SetMountConfig(v *UpdateVolumeInputMountConfig) *UpdateVolumeInput
	GetMountConfig() *UpdateVolumeInputMountConfig
	SetOssVolumeConfig(v *OSSVolumeConfig) *UpdateVolumeInput
	GetOssVolumeConfig() *OSSVolumeConfig
}

type UpdateVolumeInput struct {
	AgenticBucketVolumeConfig *AgenticBucketVolumeConfig              `json:"agenticBucketVolumeConfig,omitempty" xml:"agenticBucketVolumeConfig,omitempty"`
	AgenticFSVolumeConfig     *UpdateVolumeInputAgenticFSVolumeConfig `json:"agenticFSVolumeConfig,omitempty" xml:"agenticFSVolumeConfig,omitempty" type:"Struct"`
	JuiceFSVolumeConfig       *JuiceFSVolumeConfig                    `json:"juiceFSVolumeConfig,omitempty" xml:"juiceFSVolumeConfig,omitempty"`
	MountConfig               *UpdateVolumeInputMountConfig           `json:"mountConfig,omitempty" xml:"mountConfig,omitempty" type:"Struct"`
	OssVolumeConfig           *OSSVolumeConfig                        `json:"ossVolumeConfig,omitempty" xml:"ossVolumeConfig,omitempty"`
}

func (s UpdateVolumeInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeInput) GoString() string {
	return s.String()
}

func (s *UpdateVolumeInput) GetAgenticBucketVolumeConfig() *AgenticBucketVolumeConfig {
	return s.AgenticBucketVolumeConfig
}

func (s *UpdateVolumeInput) GetAgenticFSVolumeConfig() *UpdateVolumeInputAgenticFSVolumeConfig {
	return s.AgenticFSVolumeConfig
}

func (s *UpdateVolumeInput) GetJuiceFSVolumeConfig() *JuiceFSVolumeConfig {
	return s.JuiceFSVolumeConfig
}

func (s *UpdateVolumeInput) GetMountConfig() *UpdateVolumeInputMountConfig {
	return s.MountConfig
}

func (s *UpdateVolumeInput) GetOssVolumeConfig() *OSSVolumeConfig {
	return s.OssVolumeConfig
}

func (s *UpdateVolumeInput) SetAgenticBucketVolumeConfig(v *AgenticBucketVolumeConfig) *UpdateVolumeInput {
	s.AgenticBucketVolumeConfig = v
	return s
}

func (s *UpdateVolumeInput) SetAgenticFSVolumeConfig(v *UpdateVolumeInputAgenticFSVolumeConfig) *UpdateVolumeInput {
	s.AgenticFSVolumeConfig = v
	return s
}

func (s *UpdateVolumeInput) SetJuiceFSVolumeConfig(v *JuiceFSVolumeConfig) *UpdateVolumeInput {
	s.JuiceFSVolumeConfig = v
	return s
}

func (s *UpdateVolumeInput) SetMountConfig(v *UpdateVolumeInputMountConfig) *UpdateVolumeInput {
	s.MountConfig = v
	return s
}

func (s *UpdateVolumeInput) SetOssVolumeConfig(v *OSSVolumeConfig) *UpdateVolumeInput {
	s.OssVolumeConfig = v
	return s
}

func (s *UpdateVolumeInput) Validate() error {
	if s.AgenticBucketVolumeConfig != nil {
		if err := s.AgenticBucketVolumeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AgenticFSVolumeConfig != nil {
		if err := s.AgenticFSVolumeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.JuiceFSVolumeConfig != nil {
		if err := s.JuiceFSVolumeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MountConfig != nil {
		if err := s.MountConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OssVolumeConfig != nil {
		if err := s.OssVolumeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVolumeInputAgenticFSVolumeConfig struct {
	GroupID    *int32  `json:"groupID,omitempty" xml:"groupID,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
	UserID     *int32  `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s UpdateVolumeInputAgenticFSVolumeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeInputAgenticFSVolumeConfig) GoString() string {
	return s.String()
}

func (s *UpdateVolumeInputAgenticFSVolumeConfig) GetGroupID() *int32 {
	return s.GroupID
}

func (s *UpdateVolumeInputAgenticFSVolumeConfig) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *UpdateVolumeInputAgenticFSVolumeConfig) GetUserID() *int32 {
	return s.UserID
}

func (s *UpdateVolumeInputAgenticFSVolumeConfig) SetGroupID(v int32) *UpdateVolumeInputAgenticFSVolumeConfig {
	s.GroupID = &v
	return s
}

func (s *UpdateVolumeInputAgenticFSVolumeConfig) SetServerAddr(v string) *UpdateVolumeInputAgenticFSVolumeConfig {
	s.ServerAddr = &v
	return s
}

func (s *UpdateVolumeInputAgenticFSVolumeConfig) SetUserID(v int32) *UpdateVolumeInputAgenticFSVolumeConfig {
	s.UserID = &v
	return s
}

func (s *UpdateVolumeInputAgenticFSVolumeConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateVolumeInputMountConfig struct {
	Role      *string                                `json:"role,omitempty" xml:"role,omitempty"`
	VpcConfig *UpdateVolumeInputMountConfigVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s UpdateVolumeInputMountConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeInputMountConfig) GoString() string {
	return s.String()
}

func (s *UpdateVolumeInputMountConfig) GetRole() *string {
	return s.Role
}

func (s *UpdateVolumeInputMountConfig) GetVpcConfig() *UpdateVolumeInputMountConfigVpcConfig {
	return s.VpcConfig
}

func (s *UpdateVolumeInputMountConfig) SetRole(v string) *UpdateVolumeInputMountConfig {
	s.Role = &v
	return s
}

func (s *UpdateVolumeInputMountConfig) SetVpcConfig(v *UpdateVolumeInputMountConfigVpcConfig) *UpdateVolumeInputMountConfig {
	s.VpcConfig = v
	return s
}

func (s *UpdateVolumeInputMountConfig) Validate() error {
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVolumeInputMountConfigVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateVolumeInputMountConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeInputMountConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *UpdateVolumeInputMountConfigVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateVolumeInputMountConfigVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateVolumeInputMountConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateVolumeInputMountConfigVpcConfig) SetSecurityGroupId(v string) *UpdateVolumeInputMountConfigVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateVolumeInputMountConfigVpcConfig) SetVSwitchIds(v []*string) *UpdateVolumeInputMountConfigVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *UpdateVolumeInputMountConfigVpcConfig) SetVpcId(v string) *UpdateVolumeInputMountConfigVpcConfig {
	s.VpcId = &v
	return s
}

func (s *UpdateVolumeInputMountConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}
