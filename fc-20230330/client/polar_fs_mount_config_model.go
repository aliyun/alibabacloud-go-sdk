// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolarFsMountConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PolarFsMountConfig
	GetInstanceId() *string
	SetMountDir(v string) *PolarFsMountConfig
	GetMountDir() *string
	SetReadOnly(v bool) *PolarFsMountConfig
	GetReadOnly() *bool
	SetRemoteDir(v string) *PolarFsMountConfig
	GetRemoteDir() *string
}

type PolarFsMountConfig struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	RemoteDir  *string `json:"remoteDir,omitempty" xml:"remoteDir,omitempty"`
}

func (s PolarFsMountConfig) String() string {
	return dara.Prettify(s)
}

func (s PolarFsMountConfig) GoString() string {
	return s.String()
}

func (s *PolarFsMountConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PolarFsMountConfig) GetMountDir() *string {
	return s.MountDir
}

func (s *PolarFsMountConfig) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *PolarFsMountConfig) GetRemoteDir() *string {
	return s.RemoteDir
}

func (s *PolarFsMountConfig) SetInstanceId(v string) *PolarFsMountConfig {
	s.InstanceId = &v
	return s
}

func (s *PolarFsMountConfig) SetMountDir(v string) *PolarFsMountConfig {
	s.MountDir = &v
	return s
}

func (s *PolarFsMountConfig) SetReadOnly(v bool) *PolarFsMountConfig {
	s.ReadOnly = &v
	return s
}

func (s *PolarFsMountConfig) SetRemoteDir(v string) *PolarFsMountConfig {
	s.RemoteDir = &v
	return s
}

func (s *PolarFsMountConfig) Validate() error {
	return dara.Validate(s)
}
