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
  AgenticFSVolumeConfig *AgenticFSVolumeConfig `json:"agenticFSVolumeConfig,omitempty" xml:"agenticFSVolumeConfig,omitempty"`
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  OssVolumeConfig *OSSVolumeConfig `json:"ossVolumeConfig,omitempty" xml:"ossVolumeConfig,omitempty"`
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
  StorageClass *string `json:"storageClass,omitempty" xml:"storageClass,omitempty"`
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
  VolumeID *string `json:"volumeID,omitempty" xml:"volumeID,omitempty"`
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
  if s.OssVolumeConfig != nil {
    if err := s.OssVolumeConfig.Validate(); err != nil {
      return err
    }
  }
  return nil
}

