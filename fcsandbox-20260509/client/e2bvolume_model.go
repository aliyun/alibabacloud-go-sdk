// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BVolume interface {
  dara.Model
  String() string
  GoString() string
  SetAgenticFSVolumeConfig(v *AgenticFSVolumeConfig) *E2BVolume
  GetAgenticFSVolumeConfig() *AgenticFSVolumeConfig 
  SetCreatedAt(v string) *E2BVolume
  GetCreatedAt() *string 
  SetMountConfig(v *E2BVolumeMountConfig) *E2BVolume
  GetMountConfig() *E2BVolumeMountConfig 
  SetOssVolumeConfig(v *OSSVolumeConfig) *E2BVolume
  GetOssVolumeConfig() *OSSVolumeConfig 
  SetResourceGroupID(v string) *E2BVolume
  GetResourceGroupID() *string 
  SetStatus(v string) *E2BVolume
  GetStatus() *string 
  SetStatusReason(v string) *E2BVolume
  GetStatusReason() *string 
  SetStorageClass(v string) *E2BVolume
  GetStorageClass() *string 
  SetTeamID(v string) *E2BVolume
  GetTeamID() *string 
  SetUpdatedAt(v string) *E2BVolume
  GetUpdatedAt() *string 
  SetUserID(v string) *E2BVolume
  GetUserID() *string 
  SetVolumeID(v string) *E2BVolume
  GetVolumeID() *string 
  SetVolumeName(v string) *E2BVolume
  GetVolumeName() *string 
}

type E2BVolume struct {
  // The AgenticFS configuration.
  AgenticFSVolumeConfig *AgenticFSVolumeConfig `json:"agenticFSVolumeConfig,omitempty" xml:"agenticFSVolumeConfig,omitempty"`
  // The time when the volume was created.
  // 
  // example:
  // 
  // 2026-07-10T11:05:55Z
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  // The mount configuration.
  MountConfig *E2BVolumeMountConfig `json:"mountConfig,omitempty" xml:"mountConfig,omitempty" type:"Struct"`
  // The OSS configuration.
  OssVolumeConfig *OSSVolumeConfig `json:"ossVolumeConfig,omitempty" xml:"ossVolumeConfig,omitempty"`
  // The resource group ID.
  // 
  // example:
  // 
  // rg-acfmwxqyrgwabcd
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  // The status. Valid values:
  // 
  // - CREATING
  // 
  // - AVAILABLE
  // 
  // - ERROR
  // 
  // - DELETING
  // 
  // example:
  // 
  // CREATING
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // The reason for the status.
  // 
  // example:
  // 
  // OK
  StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
  // The storage type. Valid values:
  // 
  // - OSS
  // 
  // - AGENTIC_FS
  // 
  // example:
  // 
  // OSS
  StorageClass *string `json:"storageClass,omitempty" xml:"storageClass,omitempty"`
  // The unique identifier of the team.
  // 
  // example:
  // 
  // 88a4c762-b0ce-4661-9413-578b2309e60f
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  // The time when the volume was last updated.
  // 
  // example:
  // 
  // 2026-07-10T11:05:55Z
  UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
  // The UID of the creator.
  // 
  // example:
  // 
  // 2190856988355929
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
  // The unique identifier of the volume.
  // 
  // example:
  // 
  // 82c8c42e-cf7a-46d0-8b58-9024409c1579
  VolumeID *string `json:"volumeID,omitempty" xml:"volumeID,omitempty"`
  // The name, which is unique within the team.
  // 
  // example:
  // 
  // workspace
  VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s E2BVolume) String() string {
  return dara.Prettify(s)
}

func (s E2BVolume) GoString() string {
  return s.String()
}

func (s *E2BVolume) GetAgenticFSVolumeConfig() *AgenticFSVolumeConfig  {
  return s.AgenticFSVolumeConfig
}

func (s *E2BVolume) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *E2BVolume) GetMountConfig() *E2BVolumeMountConfig  {
  return s.MountConfig
}

func (s *E2BVolume) GetOssVolumeConfig() *OSSVolumeConfig  {
  return s.OssVolumeConfig
}

func (s *E2BVolume) GetResourceGroupID() *string  {
  return s.ResourceGroupID
}

func (s *E2BVolume) GetStatus() *string  {
  return s.Status
}

func (s *E2BVolume) GetStatusReason() *string  {
  return s.StatusReason
}

func (s *E2BVolume) GetStorageClass() *string  {
  return s.StorageClass
}

func (s *E2BVolume) GetTeamID() *string  {
  return s.TeamID
}

func (s *E2BVolume) GetUpdatedAt() *string  {
  return s.UpdatedAt
}

func (s *E2BVolume) GetUserID() *string  {
  return s.UserID
}

func (s *E2BVolume) GetVolumeID() *string  {
  return s.VolumeID
}

func (s *E2BVolume) GetVolumeName() *string  {
  return s.VolumeName
}

func (s *E2BVolume) SetAgenticFSVolumeConfig(v *AgenticFSVolumeConfig) *E2BVolume {
  s.AgenticFSVolumeConfig = v
  return s
}

func (s *E2BVolume) SetCreatedAt(v string) *E2BVolume {
  s.CreatedAt = &v
  return s
}

func (s *E2BVolume) SetMountConfig(v *E2BVolumeMountConfig) *E2BVolume {
  s.MountConfig = v
  return s
}

func (s *E2BVolume) SetOssVolumeConfig(v *OSSVolumeConfig) *E2BVolume {
  s.OssVolumeConfig = v
  return s
}

func (s *E2BVolume) SetResourceGroupID(v string) *E2BVolume {
  s.ResourceGroupID = &v
  return s
}

func (s *E2BVolume) SetStatus(v string) *E2BVolume {
  s.Status = &v
  return s
}

func (s *E2BVolume) SetStatusReason(v string) *E2BVolume {
  s.StatusReason = &v
  return s
}

func (s *E2BVolume) SetStorageClass(v string) *E2BVolume {
  s.StorageClass = &v
  return s
}

func (s *E2BVolume) SetTeamID(v string) *E2BVolume {
  s.TeamID = &v
  return s
}

func (s *E2BVolume) SetUpdatedAt(v string) *E2BVolume {
  s.UpdatedAt = &v
  return s
}

func (s *E2BVolume) SetUserID(v string) *E2BVolume {
  s.UserID = &v
  return s
}

func (s *E2BVolume) SetVolumeID(v string) *E2BVolume {
  s.VolumeID = &v
  return s
}

func (s *E2BVolume) SetVolumeName(v string) *E2BVolume {
  s.VolumeName = &v
  return s
}

func (s *E2BVolume) Validate() error {
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

type E2BVolumeMountConfig struct {
  // The RAM role that the user grants to the sandboxed container. After this role is set, the sandboxed container assumes the role to generate temporary access credentials. You can use the temporary access credentials of this role to mount storage in the sandboxed container, such as OSS and AgenticFS.
  // 
  // example:
  // 
  // acs:ram::1673427197867277:role/aliyunfcdefaultrole
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  // The virtual private cloud (VPC) ID.
  VpcConfig *E2BVolumeMountConfigVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s E2BVolumeMountConfig) String() string {
  return dara.Prettify(s)
}

func (s E2BVolumeMountConfig) GoString() string {
  return s.String()
}

func (s *E2BVolumeMountConfig) GetRole() *string  {
  return s.Role
}

func (s *E2BVolumeMountConfig) GetVpcConfig() *E2BVolumeMountConfigVpcConfig  {
  return s.VpcConfig
}

func (s *E2BVolumeMountConfig) SetRole(v string) *E2BVolumeMountConfig {
  s.Role = &v
  return s
}

func (s *E2BVolumeMountConfig) SetVpcConfig(v *E2BVolumeMountConfigVpcConfig) *E2BVolumeMountConfig {
  s.VpcConfig = v
  return s
}

func (s *E2BVolumeMountConfig) Validate() error {
  if s.VpcConfig != nil {
    if err := s.VpcConfig.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type E2BVolumeMountConfigVpcConfig struct {
  // The security group ID.
  // 
  // example:
  // 
  // sg-xxxxxx
  SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
  // The list of vSwitches.
  VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
  // The virtual private cloud (VPC) ID.
  // 
  // example:
  // 
  // vpc-2ze4l2vyhej6a6dwapm4q
  VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s E2BVolumeMountConfigVpcConfig) String() string {
  return dara.Prettify(s)
}

func (s E2BVolumeMountConfigVpcConfig) GoString() string {
  return s.String()
}

func (s *E2BVolumeMountConfigVpcConfig) GetSecurityGroupId() *string  {
  return s.SecurityGroupId
}

func (s *E2BVolumeMountConfigVpcConfig) GetVSwitchIds() []*string  {
  return s.VSwitchIds
}

func (s *E2BVolumeMountConfigVpcConfig) GetVpcId() *string  {
  return s.VpcId
}

func (s *E2BVolumeMountConfigVpcConfig) SetSecurityGroupId(v string) *E2BVolumeMountConfigVpcConfig {
  s.SecurityGroupId = &v
  return s
}

func (s *E2BVolumeMountConfigVpcConfig) SetVSwitchIds(v []*string) *E2BVolumeMountConfigVpcConfig {
  s.VSwitchIds = v
  return s
}

func (s *E2BVolumeMountConfigVpcConfig) SetVpcId(v string) *E2BVolumeMountConfigVpcConfig {
  s.VpcId = &v
  return s
}

func (s *E2BVolumeMountConfigVpcConfig) Validate() error {
  return dara.Validate(s)
}

