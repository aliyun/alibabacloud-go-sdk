// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInnerSandboxVolumeMount interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticFs(v *InnerSandboxVolumeMountAgenticFs) *InnerSandboxVolumeMount
	GetAgenticFs() *InnerSandboxVolumeMountAgenticFs
	SetNamed(v *InnerSandboxVolumeMountNamed) *InnerSandboxVolumeMount
	GetNamed() *InnerSandboxVolumeMountNamed
	SetOss(v *InnerSandboxVolumeMountOss) *InnerSandboxVolumeMount
	GetOss() *InnerSandboxVolumeMountOss
}

type InnerSandboxVolumeMount struct {
	AgenticFs *InnerSandboxVolumeMountAgenticFs `json:"agenticFs,omitempty" xml:"agenticFs,omitempty" type:"Struct"`
	Named     *InnerSandboxVolumeMountNamed     `json:"named,omitempty" xml:"named,omitempty" type:"Struct"`
	Oss       *InnerSandboxVolumeMountOss       `json:"oss,omitempty" xml:"oss,omitempty" type:"Struct"`
}

func (s InnerSandboxVolumeMount) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxVolumeMount) GoString() string {
	return s.String()
}

func (s *InnerSandboxVolumeMount) GetAgenticFs() *InnerSandboxVolumeMountAgenticFs {
	return s.AgenticFs
}

func (s *InnerSandboxVolumeMount) GetNamed() *InnerSandboxVolumeMountNamed {
	return s.Named
}

func (s *InnerSandboxVolumeMount) GetOss() *InnerSandboxVolumeMountOss {
	return s.Oss
}

func (s *InnerSandboxVolumeMount) SetAgenticFs(v *InnerSandboxVolumeMountAgenticFs) *InnerSandboxVolumeMount {
	s.AgenticFs = v
	return s
}

func (s *InnerSandboxVolumeMount) SetNamed(v *InnerSandboxVolumeMountNamed) *InnerSandboxVolumeMount {
	s.Named = v
	return s
}

func (s *InnerSandboxVolumeMount) SetOss(v *InnerSandboxVolumeMountOss) *InnerSandboxVolumeMount {
	s.Oss = v
	return s
}

func (s *InnerSandboxVolumeMount) Validate() error {
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

type InnerSandboxVolumeMountAgenticFs struct {
	// example:
	//
	// 1000
	GroupID     *int32                                         `json:"groupID,omitempty" xml:"groupID,omitempty"`
	MountPoints []*InnerSandboxVolumeMountAgenticFsMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	UserID *int32 `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s InnerSandboxVolumeMountAgenticFs) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxVolumeMountAgenticFs) GoString() string {
	return s.String()
}

func (s *InnerSandboxVolumeMountAgenticFs) GetGroupID() *int32 {
	return s.GroupID
}

func (s *InnerSandboxVolumeMountAgenticFs) GetMountPoints() []*InnerSandboxVolumeMountAgenticFsMountPoints {
	return s.MountPoints
}

func (s *InnerSandboxVolumeMountAgenticFs) GetUserID() *int32 {
	return s.UserID
}

func (s *InnerSandboxVolumeMountAgenticFs) SetGroupID(v int32) *InnerSandboxVolumeMountAgenticFs {
	s.GroupID = &v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFs) SetMountPoints(v []*InnerSandboxVolumeMountAgenticFsMountPoints) *InnerSandboxVolumeMountAgenticFs {
	s.MountPoints = v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFs) SetUserID(v int32) *InnerSandboxVolumeMountAgenticFs {
	s.UserID = &v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFs) Validate() error {
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

type InnerSandboxVolumeMountAgenticFsMountPoints struct {
	// example:
	//
	// ap-xxxx
	AccessPointID *string `json:"accessPointID,omitempty" xml:"accessPointID,omitempty"`
	// example:
	//
	// agentic-xxxx
	AgenticSpaceID *string `json:"agenticSpaceID,omitempty" xml:"agenticSpaceID,omitempty"`
	// example:
	//
	// 03204sl2qjiax4oxxxx
	FileSystemID *string `json:"fileSystemID,omitempty" xml:"fileSystemID,omitempty"`
	// example:
	//
	// /mnt/agenticfs
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// example:
	//
	// ap-xxxx.xxxx-ljs60.cn-shanghai.nas.aliyuncs.com
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s InnerSandboxVolumeMountAgenticFsMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxVolumeMountAgenticFsMountPoints) GoString() string {
	return s.String()
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) GetAccessPointID() *string {
	return s.AccessPointID
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) GetAgenticSpaceID() *string {
	return s.AgenticSpaceID
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) GetFileSystemID() *string {
	return s.FileSystemID
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) SetAccessPointID(v string) *InnerSandboxVolumeMountAgenticFsMountPoints {
	s.AccessPointID = &v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) SetAgenticSpaceID(v string) *InnerSandboxVolumeMountAgenticFsMountPoints {
	s.AgenticSpaceID = &v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) SetFileSystemID(v string) *InnerSandboxVolumeMountAgenticFsMountPoints {
	s.FileSystemID = &v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) SetMountDir(v string) *InnerSandboxVolumeMountAgenticFsMountPoints {
	s.MountDir = &v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) SetServerAddr(v string) *InnerSandboxVolumeMountAgenticFsMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *InnerSandboxVolumeMountAgenticFsMountPoints) Validate() error {
	return dara.Validate(s)
}

type InnerSandboxVolumeMountNamed struct {
	MountPoints []*InnerSandboxVolumeMountNamedMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s InnerSandboxVolumeMountNamed) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxVolumeMountNamed) GoString() string {
	return s.String()
}

func (s *InnerSandboxVolumeMountNamed) GetMountPoints() []*InnerSandboxVolumeMountNamedMountPoints {
	return s.MountPoints
}

func (s *InnerSandboxVolumeMountNamed) SetMountPoints(v []*InnerSandboxVolumeMountNamedMountPoints) *InnerSandboxVolumeMountNamed {
	s.MountPoints = v
	return s
}

func (s *InnerSandboxVolumeMountNamed) Validate() error {
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

type InnerSandboxVolumeMountNamedMountPoints struct {
	// example:
	//
	// /mnt/named
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// example:
	//
	// workspace
	VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s InnerSandboxVolumeMountNamedMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxVolumeMountNamedMountPoints) GoString() string {
	return s.String()
}

func (s *InnerSandboxVolumeMountNamedMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InnerSandboxVolumeMountNamedMountPoints) GetVolumeName() *string {
	return s.VolumeName
}

func (s *InnerSandboxVolumeMountNamedMountPoints) SetMountDir(v string) *InnerSandboxVolumeMountNamedMountPoints {
	s.MountDir = &v
	return s
}

func (s *InnerSandboxVolumeMountNamedMountPoints) SetVolumeName(v string) *InnerSandboxVolumeMountNamedMountPoints {
	s.VolumeName = &v
	return s
}

func (s *InnerSandboxVolumeMountNamedMountPoints) Validate() error {
	return dara.Validate(s)
}

type InnerSandboxVolumeMountOss struct {
	MountPoints []*InnerSandboxVolumeMountOssMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s InnerSandboxVolumeMountOss) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxVolumeMountOss) GoString() string {
	return s.String()
}

func (s *InnerSandboxVolumeMountOss) GetMountPoints() []*InnerSandboxVolumeMountOssMountPoints {
	return s.MountPoints
}

func (s *InnerSandboxVolumeMountOss) SetMountPoints(v []*InnerSandboxVolumeMountOssMountPoints) *InnerSandboxVolumeMountOss {
	s.MountPoints = v
	return s
}

func (s *InnerSandboxVolumeMountOss) Validate() error {
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

type InnerSandboxVolumeMountOssMountPoints struct {
	// example:
	//
	// oss-bucket-test
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// example:
	//
	// /
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// example:
	//
	// oss-cn-shenzhen-internal.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// /mnt/oss
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s InnerSandboxVolumeMountOssMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxVolumeMountOssMountPoints) GoString() string {
	return s.String()
}

func (s *InnerSandboxVolumeMountOssMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *InnerSandboxVolumeMountOssMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *InnerSandboxVolumeMountOssMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *InnerSandboxVolumeMountOssMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InnerSandboxVolumeMountOssMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *InnerSandboxVolumeMountOssMountPoints) SetBucketName(v string) *InnerSandboxVolumeMountOssMountPoints {
	s.BucketName = &v
	return s
}

func (s *InnerSandboxVolumeMountOssMountPoints) SetBucketPath(v string) *InnerSandboxVolumeMountOssMountPoints {
	s.BucketPath = &v
	return s
}

func (s *InnerSandboxVolumeMountOssMountPoints) SetEndpoint(v string) *InnerSandboxVolumeMountOssMountPoints {
	s.Endpoint = &v
	return s
}

func (s *InnerSandboxVolumeMountOssMountPoints) SetMountDir(v string) *InnerSandboxVolumeMountOssMountPoints {
	s.MountDir = &v
	return s
}

func (s *InnerSandboxVolumeMountOssMountPoints) SetReadOnly(v bool) *InnerSandboxVolumeMountOssMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *InnerSandboxVolumeMountOssMountPoints) Validate() error {
	return dara.Validate(s)
}
