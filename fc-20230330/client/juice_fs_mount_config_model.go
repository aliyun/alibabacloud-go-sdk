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
	// An array of strings containing additional command-line arguments for the mount command. For example, use these arguments to set cache sizes or other performance-tuning options.
	Args []*string `json:"args" xml:"args" type:"Repeated"`
	// The path within the function\\"s local filesystem to mount the volume. For example, /mnt/data. This parameter is required.
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// The subdirectory within the JuiceFS volume to mount. If not specified, the root of the volume is mounted.
	RemoteDir *string `json:"remoteDir,omitempty" xml:"remoteDir,omitempty"`
	// The authentication token to access the JuiceFS volume.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The name of the JuiceFS volume to mount. This parameter is required.
	VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
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
