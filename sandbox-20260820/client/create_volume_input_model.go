// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVolumeInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticBucketVolumeConfig(v *AgenticBucketVolumeConfig) *CreateVolumeInput
	GetAgenticBucketVolumeConfig() *AgenticBucketVolumeConfig
	SetAgenticFSVolumeConfig(v *CreateVolumeInputAgenticFSVolumeConfig) *CreateVolumeInput
	GetAgenticFSVolumeConfig() *CreateVolumeInputAgenticFSVolumeConfig
	SetJuiceFSVolumeConfig(v *JuiceFSVolumeConfig) *CreateVolumeInput
	GetJuiceFSVolumeConfig() *JuiceFSVolumeConfig
	SetMountConfig(v *CreateVolumeInputMountConfig) *CreateVolumeInput
	GetMountConfig() *CreateVolumeInputMountConfig
	SetOssVolumeConfig(v *OSSVolumeConfig) *CreateVolumeInput
	GetOssVolumeConfig() *OSSVolumeConfig
	SetTeamID(v string) *CreateVolumeInput
	GetTeamID() *string
	SetVolumeName(v string) *CreateVolumeInput
	GetVolumeName() *string
}

type CreateVolumeInput struct {
	AgenticBucketVolumeConfig *AgenticBucketVolumeConfig              `json:"agenticBucketVolumeConfig,omitempty" xml:"agenticBucketVolumeConfig,omitempty"`
	AgenticFSVolumeConfig     *CreateVolumeInputAgenticFSVolumeConfig `json:"agenticFSVolumeConfig,omitempty" xml:"agenticFSVolumeConfig,omitempty" type:"Struct"`
	JuiceFSVolumeConfig       *JuiceFSVolumeConfig                    `json:"juiceFSVolumeConfig,omitempty" xml:"juiceFSVolumeConfig,omitempty"`
	MountConfig               *CreateVolumeInputMountConfig           `json:"mountConfig,omitempty" xml:"mountConfig,omitempty" type:"Struct"`
	OssVolumeConfig           *OSSVolumeConfig                        `json:"ossVolumeConfig,omitempty" xml:"ossVolumeConfig,omitempty"`
	TeamID                    *string                                 `json:"teamID,omitempty" xml:"teamID,omitempty"`
	VolumeName                *string                                 `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s CreateVolumeInput) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeInput) GoString() string {
	return s.String()
}

func (s *CreateVolumeInput) GetAgenticBucketVolumeConfig() *AgenticBucketVolumeConfig {
	return s.AgenticBucketVolumeConfig
}

func (s *CreateVolumeInput) GetAgenticFSVolumeConfig() *CreateVolumeInputAgenticFSVolumeConfig {
	return s.AgenticFSVolumeConfig
}

func (s *CreateVolumeInput) GetJuiceFSVolumeConfig() *JuiceFSVolumeConfig {
	return s.JuiceFSVolumeConfig
}

func (s *CreateVolumeInput) GetMountConfig() *CreateVolumeInputMountConfig {
	return s.MountConfig
}

func (s *CreateVolumeInput) GetOssVolumeConfig() *OSSVolumeConfig {
	return s.OssVolumeConfig
}

func (s *CreateVolumeInput) GetTeamID() *string {
	return s.TeamID
}

func (s *CreateVolumeInput) GetVolumeName() *string {
	return s.VolumeName
}

func (s *CreateVolumeInput) SetAgenticBucketVolumeConfig(v *AgenticBucketVolumeConfig) *CreateVolumeInput {
	s.AgenticBucketVolumeConfig = v
	return s
}

func (s *CreateVolumeInput) SetAgenticFSVolumeConfig(v *CreateVolumeInputAgenticFSVolumeConfig) *CreateVolumeInput {
	s.AgenticFSVolumeConfig = v
	return s
}

func (s *CreateVolumeInput) SetJuiceFSVolumeConfig(v *JuiceFSVolumeConfig) *CreateVolumeInput {
	s.JuiceFSVolumeConfig = v
	return s
}

func (s *CreateVolumeInput) SetMountConfig(v *CreateVolumeInputMountConfig) *CreateVolumeInput {
	s.MountConfig = v
	return s
}

func (s *CreateVolumeInput) SetOssVolumeConfig(v *OSSVolumeConfig) *CreateVolumeInput {
	s.OssVolumeConfig = v
	return s
}

func (s *CreateVolumeInput) SetTeamID(v string) *CreateVolumeInput {
	s.TeamID = &v
	return s
}

func (s *CreateVolumeInput) SetVolumeName(v string) *CreateVolumeInput {
	s.VolumeName = &v
	return s
}

func (s *CreateVolumeInput) Validate() error {
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

type CreateVolumeInputAgenticFSVolumeConfig struct {
	GroupID    *int32  `json:"groupID,omitempty" xml:"groupID,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
	UserID     *int32  `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s CreateVolumeInputAgenticFSVolumeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeInputAgenticFSVolumeConfig) GoString() string {
	return s.String()
}

func (s *CreateVolumeInputAgenticFSVolumeConfig) GetGroupID() *int32 {
	return s.GroupID
}

func (s *CreateVolumeInputAgenticFSVolumeConfig) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *CreateVolumeInputAgenticFSVolumeConfig) GetUserID() *int32 {
	return s.UserID
}

func (s *CreateVolumeInputAgenticFSVolumeConfig) SetGroupID(v int32) *CreateVolumeInputAgenticFSVolumeConfig {
	s.GroupID = &v
	return s
}

func (s *CreateVolumeInputAgenticFSVolumeConfig) SetServerAddr(v string) *CreateVolumeInputAgenticFSVolumeConfig {
	s.ServerAddr = &v
	return s
}

func (s *CreateVolumeInputAgenticFSVolumeConfig) SetUserID(v int32) *CreateVolumeInputAgenticFSVolumeConfig {
	s.UserID = &v
	return s
}

func (s *CreateVolumeInputAgenticFSVolumeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateVolumeInputMountConfig struct {
	Role      *string                                `json:"role,omitempty" xml:"role,omitempty"`
	VpcConfig *CreateVolumeInputMountConfigVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s CreateVolumeInputMountConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeInputMountConfig) GoString() string {
	return s.String()
}

func (s *CreateVolumeInputMountConfig) GetRole() *string {
	return s.Role
}

func (s *CreateVolumeInputMountConfig) GetVpcConfig() *CreateVolumeInputMountConfigVpcConfig {
	return s.VpcConfig
}

func (s *CreateVolumeInputMountConfig) SetRole(v string) *CreateVolumeInputMountConfig {
	s.Role = &v
	return s
}

func (s *CreateVolumeInputMountConfig) SetVpcConfig(v *CreateVolumeInputMountConfigVpcConfig) *CreateVolumeInputMountConfig {
	s.VpcConfig = v
	return s
}

func (s *CreateVolumeInputMountConfig) Validate() error {
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVolumeInputMountConfigVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateVolumeInputMountConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeInputMountConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *CreateVolumeInputMountConfigVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateVolumeInputMountConfigVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateVolumeInputMountConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVolumeInputMountConfigVpcConfig) SetSecurityGroupId(v string) *CreateVolumeInputMountConfigVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateVolumeInputMountConfigVpcConfig) SetVSwitchIds(v []*string) *CreateVolumeInputMountConfigVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *CreateVolumeInputMountConfigVpcConfig) SetVpcId(v string) *CreateVolumeInputMountConfigVpcConfig {
	s.VpcId = &v
	return s
}

func (s *CreateVolumeInputMountConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}
