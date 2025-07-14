// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebNASMountPoint interface {
	dara.Model
	String() string
	GoString() string
	SetMountDir(v string) *WebNASMountPoint
	GetMountDir() *string
	SetNasAddr(v string) *WebNASMountPoint
	GetNasAddr() *string
	SetNasPath(v string) *WebNASMountPoint
	GetNasPath() *string
}

type WebNASMountPoint struct {
	MountDir *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	NasAddr  *string `json:"NasAddr,omitempty" xml:"NasAddr,omitempty"`
	NasPath  *string `json:"NasPath,omitempty" xml:"NasPath,omitempty"`
}

func (s WebNASMountPoint) String() string {
	return dara.Prettify(s)
}

func (s WebNASMountPoint) GoString() string {
	return s.String()
}

func (s *WebNASMountPoint) GetMountDir() *string {
	return s.MountDir
}

func (s *WebNASMountPoint) GetNasAddr() *string {
	return s.NasAddr
}

func (s *WebNASMountPoint) GetNasPath() *string {
	return s.NasPath
}

func (s *WebNASMountPoint) SetMountDir(v string) *WebNASMountPoint {
	s.MountDir = &v
	return s
}

func (s *WebNASMountPoint) SetNasAddr(v string) *WebNASMountPoint {
	s.NasAddr = &v
	return s
}

func (s *WebNASMountPoint) SetNasPath(v string) *WebNASMountPoint {
	s.NasPath = &v
	return s
}

func (s *WebNASMountPoint) Validate() error {
	return dara.Validate(s)
}
