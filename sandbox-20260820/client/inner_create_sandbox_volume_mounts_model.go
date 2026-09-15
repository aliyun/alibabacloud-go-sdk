// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInnerCreateSandboxVolumeMounts interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticFs(v *InnerCreateSandboxVolumeMountsAgenticFs) *InnerCreateSandboxVolumeMounts
	GetAgenticFs() *InnerCreateSandboxVolumeMountsAgenticFs
	SetNamed(v *InnerCreateSandboxVolumeMountsNamed) *InnerCreateSandboxVolumeMounts
	GetNamed() *InnerCreateSandboxVolumeMountsNamed
	SetOss(v *InnerCreateSandboxVolumeMountsOss) *InnerCreateSandboxVolumeMounts
	GetOss() *InnerCreateSandboxVolumeMountsOss
}

type InnerCreateSandboxVolumeMounts struct {
	AgenticFs *InnerCreateSandboxVolumeMountsAgenticFs `json:"agenticFs,omitempty" xml:"agenticFs,omitempty" type:"Struct"`
	Named     *InnerCreateSandboxVolumeMountsNamed     `json:"named,omitempty" xml:"named,omitempty" type:"Struct"`
	Oss       *InnerCreateSandboxVolumeMountsOss       `json:"oss,omitempty" xml:"oss,omitempty" type:"Struct"`
}

func (s InnerCreateSandboxVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxVolumeMounts) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxVolumeMounts) GetAgenticFs() *InnerCreateSandboxVolumeMountsAgenticFs {
	return s.AgenticFs
}

func (s *InnerCreateSandboxVolumeMounts) GetNamed() *InnerCreateSandboxVolumeMountsNamed {
	return s.Named
}

func (s *InnerCreateSandboxVolumeMounts) GetOss() *InnerCreateSandboxVolumeMountsOss {
	return s.Oss
}

func (s *InnerCreateSandboxVolumeMounts) SetAgenticFs(v *InnerCreateSandboxVolumeMountsAgenticFs) *InnerCreateSandboxVolumeMounts {
	s.AgenticFs = v
	return s
}

func (s *InnerCreateSandboxVolumeMounts) SetNamed(v *InnerCreateSandboxVolumeMountsNamed) *InnerCreateSandboxVolumeMounts {
	s.Named = v
	return s
}

func (s *InnerCreateSandboxVolumeMounts) SetOss(v *InnerCreateSandboxVolumeMountsOss) *InnerCreateSandboxVolumeMounts {
	s.Oss = v
	return s
}

func (s *InnerCreateSandboxVolumeMounts) Validate() error {
	if s.AgenticFs != nil {
		if err := s.AgenticFs.Validate(); err != nil {
			return err
		}
	}
	if s.Named != nil {
		if err := s.Named.Validate(); err != nil {
			return err
		}
	}
	if s.Oss != nil {
		if err := s.Oss.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InnerCreateSandboxVolumeMountsAgenticFs struct {
	GroupID     *int32                                                `json:"groupID,omitempty" xml:"groupID,omitempty"`
	MountPoints []*InnerCreateSandboxVolumeMountsAgenticFsMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserID      *int32                                                `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s InnerCreateSandboxVolumeMountsAgenticFs) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxVolumeMountsAgenticFs) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxVolumeMountsAgenticFs) GetGroupID() *int32 {
	return s.GroupID
}

func (s *InnerCreateSandboxVolumeMountsAgenticFs) GetMountPoints() []*InnerCreateSandboxVolumeMountsAgenticFsMountPoints {
	return s.MountPoints
}

func (s *InnerCreateSandboxVolumeMountsAgenticFs) GetUserID() *int32 {
	return s.UserID
}

func (s *InnerCreateSandboxVolumeMountsAgenticFs) SetGroupID(v int32) *InnerCreateSandboxVolumeMountsAgenticFs {
	s.GroupID = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsAgenticFs) SetMountPoints(v []*InnerCreateSandboxVolumeMountsAgenticFsMountPoints) *InnerCreateSandboxVolumeMountsAgenticFs {
	s.MountPoints = v
	return s
}

func (s *InnerCreateSandboxVolumeMountsAgenticFs) SetUserID(v int32) *InnerCreateSandboxVolumeMountsAgenticFs {
	s.UserID = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsAgenticFs) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InnerCreateSandboxVolumeMountsAgenticFsMountPoints struct {
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s InnerCreateSandboxVolumeMountsAgenticFsMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxVolumeMountsAgenticFsMountPoints) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxVolumeMountsAgenticFsMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InnerCreateSandboxVolumeMountsAgenticFsMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *InnerCreateSandboxVolumeMountsAgenticFsMountPoints) SetMountDir(v string) *InnerCreateSandboxVolumeMountsAgenticFsMountPoints {
	s.MountDir = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsAgenticFsMountPoints) SetServerAddr(v string) *InnerCreateSandboxVolumeMountsAgenticFsMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsAgenticFsMountPoints) Validate() error {
	return dara.Validate(s)
}

type InnerCreateSandboxVolumeMountsNamed struct {
	MountPoints []*InnerCreateSandboxVolumeMountsNamedMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s InnerCreateSandboxVolumeMountsNamed) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxVolumeMountsNamed) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxVolumeMountsNamed) GetMountPoints() []*InnerCreateSandboxVolumeMountsNamedMountPoints {
	return s.MountPoints
}

func (s *InnerCreateSandboxVolumeMountsNamed) SetMountPoints(v []*InnerCreateSandboxVolumeMountsNamedMountPoints) *InnerCreateSandboxVolumeMountsNamed {
	s.MountPoints = v
	return s
}

func (s *InnerCreateSandboxVolumeMountsNamed) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InnerCreateSandboxVolumeMountsNamedMountPoints struct {
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s InnerCreateSandboxVolumeMountsNamedMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxVolumeMountsNamedMountPoints) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxVolumeMountsNamedMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InnerCreateSandboxVolumeMountsNamedMountPoints) GetVolumeName() *string {
	return s.VolumeName
}

func (s *InnerCreateSandboxVolumeMountsNamedMountPoints) SetMountDir(v string) *InnerCreateSandboxVolumeMountsNamedMountPoints {
	s.MountDir = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsNamedMountPoints) SetVolumeName(v string) *InnerCreateSandboxVolumeMountsNamedMountPoints {
	s.VolumeName = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsNamedMountPoints) Validate() error {
	return dara.Validate(s)
}

type InnerCreateSandboxVolumeMountsOss struct {
	MountPoints []*InnerCreateSandboxVolumeMountsOssMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s InnerCreateSandboxVolumeMountsOss) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxVolumeMountsOss) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxVolumeMountsOss) GetMountPoints() []*InnerCreateSandboxVolumeMountsOssMountPoints {
	return s.MountPoints
}

func (s *InnerCreateSandboxVolumeMountsOss) SetMountPoints(v []*InnerCreateSandboxVolumeMountsOssMountPoints) *InnerCreateSandboxVolumeMountsOss {
	s.MountPoints = v
	return s
}

func (s *InnerCreateSandboxVolumeMountsOss) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InnerCreateSandboxVolumeMountsOssMountPoints struct {
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s InnerCreateSandboxVolumeMountsOssMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InnerCreateSandboxVolumeMountsOssMountPoints) GoString() string {
	return s.String()
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) SetBucketName(v string) *InnerCreateSandboxVolumeMountsOssMountPoints {
	s.BucketName = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) SetBucketPath(v string) *InnerCreateSandboxVolumeMountsOssMountPoints {
	s.BucketPath = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) SetEndpoint(v string) *InnerCreateSandboxVolumeMountsOssMountPoints {
	s.Endpoint = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) SetMountDir(v string) *InnerCreateSandboxVolumeMountsOssMountPoints {
	s.MountDir = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) SetReadOnly(v bool) *InnerCreateSandboxVolumeMountsOssMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *InnerCreateSandboxVolumeMountsOssMountPoints) Validate() error {
	return dara.Validate(s)
}
