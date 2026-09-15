// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJuiceFSVolumeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *JuiceFSVolumeConfig
	GetArgs() []*string
	SetBaseURL(v string) *JuiceFSVolumeConfig
	GetBaseURL() *string
	SetRemoteDir(v string) *JuiceFSVolumeConfig
	GetRemoteDir() *string
	SetToken(v string) *JuiceFSVolumeConfig
	GetToken() *string
	SetVolumeName(v string) *JuiceFSVolumeConfig
	GetVolumeName() *string
}

type JuiceFSVolumeConfig struct {
	Args       []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
	BaseURL    *string   `json:"baseURL,omitempty" xml:"baseURL,omitempty"`
	RemoteDir  *string   `json:"remoteDir,omitempty" xml:"remoteDir,omitempty"`
	Token      *string   `json:"token,omitempty" xml:"token,omitempty"`
	VolumeName *string   `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s JuiceFSVolumeConfig) String() string {
	return dara.Prettify(s)
}

func (s JuiceFSVolumeConfig) GoString() string {
	return s.String()
}

func (s *JuiceFSVolumeConfig) GetArgs() []*string {
	return s.Args
}

func (s *JuiceFSVolumeConfig) GetBaseURL() *string {
	return s.BaseURL
}

func (s *JuiceFSVolumeConfig) GetRemoteDir() *string {
	return s.RemoteDir
}

func (s *JuiceFSVolumeConfig) GetToken() *string {
	return s.Token
}

func (s *JuiceFSVolumeConfig) GetVolumeName() *string {
	return s.VolumeName
}

func (s *JuiceFSVolumeConfig) SetArgs(v []*string) *JuiceFSVolumeConfig {
	s.Args = v
	return s
}

func (s *JuiceFSVolumeConfig) SetBaseURL(v string) *JuiceFSVolumeConfig {
	s.BaseURL = &v
	return s
}

func (s *JuiceFSVolumeConfig) SetRemoteDir(v string) *JuiceFSVolumeConfig {
	s.RemoteDir = &v
	return s
}

func (s *JuiceFSVolumeConfig) SetToken(v string) *JuiceFSVolumeConfig {
	s.Token = &v
	return s
}

func (s *JuiceFSVolumeConfig) SetVolumeName(v string) *JuiceFSVolumeConfig {
	s.VolumeName = &v
	return s
}

func (s *JuiceFSVolumeConfig) Validate() error {
	return dara.Validate(s)
}
