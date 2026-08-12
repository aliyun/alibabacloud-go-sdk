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
	SetOssVolumeConfig(v *OSSVolumeConfig) *CreateVolumeInput
	GetOssVolumeConfig() *OSSVolumeConfig
	SetTeamID(v string) *CreateVolumeInput
	GetTeamID() *string
	SetVolumeName(v string) *CreateVolumeInput
	GetVolumeName() *string
}

type CreateVolumeInput struct {
	AgenticFSVolumeConfig *AgenticFSVolumeConfig `json:"agenticFSVolumeConfig,omitempty" xml:"agenticFSVolumeConfig,omitempty"`
	OssVolumeConfig       *OSSVolumeConfig       `json:"ossVolumeConfig,omitempty" xml:"ossVolumeConfig,omitempty"`
	TeamID                *string                `json:"teamID,omitempty" xml:"teamID,omitempty"`
	VolumeName            *string                `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
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
	if s.OssVolumeConfig != nil {
		if err := s.OssVolumeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
