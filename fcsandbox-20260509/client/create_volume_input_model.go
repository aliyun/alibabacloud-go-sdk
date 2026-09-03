// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVolumeInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticFSVolumeConfig(v *AgenticFSVolumeConfig) *CreateVolumeInput
	GetAgenticFSVolumeConfig() *AgenticFSVolumeConfig
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
	// The AgenticFS configuration.
	AgenticFSVolumeConfig *AgenticFSVolumeConfig `json:"agenticFSVolumeConfig,omitempty" xml:"agenticFSVolumeConfig,omitempty"`
	// The mount configuration.
	MountConfig *CreateVolumeInputMountConfig `json:"mountConfig,omitempty" xml:"mountConfig,omitempty" type:"Struct"`
	// The OSS configuration.
	OssVolumeConfig *OSSVolumeConfig `json:"ossVolumeConfig,omitempty" xml:"ossVolumeConfig,omitempty"`
	// The unique identifier of the team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The name, which must be unique within the team.
	//
	// example:
	//
	// workspace
	VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s CreateVolumeInput) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeInput) GoString() string {
	return s.String()
}

func (s *CreateVolumeInput) GetAgenticFSVolumeConfig() *AgenticFSVolumeConfig {
	return s.AgenticFSVolumeConfig
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

func (s *CreateVolumeInput) SetAgenticFSVolumeConfig(v *AgenticFSVolumeConfig) *CreateVolumeInput {
	s.AgenticFSVolumeConfig = v
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
	if s.AgenticFSVolumeConfig != nil {
		if err := s.AgenticFSVolumeConfig.Validate(); err != nil {
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

type CreateVolumeInputMountConfig struct {
	// The RAM role that the user grants to the cloud sandbox. After this role is set, the cloud sandbox assumes the role to generate temporary access credentials. You can use the temporary access credentials of this role to mount storage in the cloud sandbox, such as OSS and AgenticFS.
	//
	// example:
	//
	// acs:ram::1338904783509062:role/aliyunfcdefaultrole
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The VPC configuration.
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
	// The security group ID.
	//
	// example:
	//
	// sg-xxxxxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The list of vSwitches.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-8vbq8hbepimf6lr7uyqub
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
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
