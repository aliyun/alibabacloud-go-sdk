// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLocalMountSpec interface {
	dara.Model
	String() string
	GoString() string
	SetLocalPath(v string) *LocalMountSpec
	GetLocalPath() *string
	SetMountMode(v string) *LocalMountSpec
	GetMountMode() *string
	SetMountPath(v string) *LocalMountSpec
	GetMountPath() *string
}

type LocalMountSpec struct {
	// The local path on the host. This field is invalid when MountMode is set to Block.
	//
	// example:
	//
	// /local/a
	LocalPath *string `json:"LocalPath,omitempty" xml:"LocalPath,omitempty"`
	// The mount mode. Valid values: Block and Filesystem.
	//
	// example:
	//
	// Filesystem
	MountMode *string `json:"MountMode,omitempty" xml:"MountMode,omitempty"`
	// The absolute mount path in the container. The path must start with a forward slash (/). When MountMode is set to Block, the path must start with /dev/.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s LocalMountSpec) String() string {
	return dara.Prettify(s)
}

func (s LocalMountSpec) GoString() string {
	return s.String()
}

func (s *LocalMountSpec) GetLocalPath() *string {
	return s.LocalPath
}

func (s *LocalMountSpec) GetMountMode() *string {
	return s.MountMode
}

func (s *LocalMountSpec) GetMountPath() *string {
	return s.MountPath
}

func (s *LocalMountSpec) SetLocalPath(v string) *LocalMountSpec {
	s.LocalPath = &v
	return s
}

func (s *LocalMountSpec) SetMountMode(v string) *LocalMountSpec {
	s.MountMode = &v
	return s
}

func (s *LocalMountSpec) SetMountPath(v string) *LocalMountSpec {
	s.MountPath = &v
	return s
}

func (s *LocalMountSpec) Validate() error {
	return dara.Validate(s)
}
