// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJuiceFsMountConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *JuiceFsMountConfig
	GetArgs() []*string
	SetMountDir(v string) *JuiceFsMountConfig
	GetMountDir() *string
	SetRemoteDir(v string) *JuiceFsMountConfig
	GetRemoteDir() *string
	SetToken(v string) *JuiceFsMountConfig
	GetToken() *string
	SetVolumeName(v string) *JuiceFsMountConfig
	GetVolumeName() *string
}

type JuiceFsMountConfig struct {
	Args       []*string `json:"args" xml:"args" type:"Repeated"`
	MountDir   *string   `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	RemoteDir  *string   `json:"remoteDir,omitempty" xml:"remoteDir,omitempty"`
	Token      *string   `json:"token,omitempty" xml:"token,omitempty"`
	VolumeName *string   `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s JuiceFsMountConfig) String() string {
	return dara.Prettify(s)
}

func (s JuiceFsMountConfig) GoString() string {
	return s.String()
}

func (s *JuiceFsMountConfig) GetArgs() []*string {
	return s.Args
}

func (s *JuiceFsMountConfig) GetMountDir() *string {
	return s.MountDir
}

func (s *JuiceFsMountConfig) GetRemoteDir() *string {
	return s.RemoteDir
}

func (s *JuiceFsMountConfig) GetToken() *string {
	return s.Token
}

func (s *JuiceFsMountConfig) GetVolumeName() *string {
	return s.VolumeName
}

func (s *JuiceFsMountConfig) SetArgs(v []*string) *JuiceFsMountConfig {
	s.Args = v
	return s
}

func (s *JuiceFsMountConfig) SetMountDir(v string) *JuiceFsMountConfig {
	s.MountDir = &v
	return s
}

func (s *JuiceFsMountConfig) SetRemoteDir(v string) *JuiceFsMountConfig {
	s.RemoteDir = &v
	return s
}

func (s *JuiceFsMountConfig) SetToken(v string) *JuiceFsMountConfig {
	s.Token = &v
	return s
}

func (s *JuiceFsMountConfig) SetVolumeName(v string) *JuiceFsMountConfig {
	s.VolumeName = &v
	return s
}

func (s *JuiceFsMountConfig) Validate() error {
	return dara.Validate(s)
}
