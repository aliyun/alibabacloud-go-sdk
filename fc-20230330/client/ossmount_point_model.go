// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOSSMountPoint interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *OSSMountPoint
	GetBucketName() *string
	SetBucketPath(v string) *OSSMountPoint
	GetBucketPath() *string
	SetEndpoint(v string) *OSSMountPoint
	GetEndpoint() *string
	SetMountDir(v string) *OSSMountPoint
	GetMountDir() *string
	SetReadOnly(v bool) *OSSMountPoint
	GetReadOnly() *bool
}

type OSSMountPoint struct {
	// example:
	//
	// my-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// example:
	//
	// /my-dir
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// example:
	//
	// http://oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// /mnt/dir
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s OSSMountPoint) String() string {
	return dara.Prettify(s)
}

func (s OSSMountPoint) GoString() string {
	return s.String()
}

func (s *OSSMountPoint) GetBucketName() *string {
	return s.BucketName
}

func (s *OSSMountPoint) GetBucketPath() *string {
	return s.BucketPath
}

func (s *OSSMountPoint) GetEndpoint() *string {
	return s.Endpoint
}

func (s *OSSMountPoint) GetMountDir() *string {
	return s.MountDir
}

func (s *OSSMountPoint) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *OSSMountPoint) SetBucketName(v string) *OSSMountPoint {
	s.BucketName = &v
	return s
}

func (s *OSSMountPoint) SetBucketPath(v string) *OSSMountPoint {
	s.BucketPath = &v
	return s
}

func (s *OSSMountPoint) SetEndpoint(v string) *OSSMountPoint {
	s.Endpoint = &v
	return s
}

func (s *OSSMountPoint) SetMountDir(v string) *OSSMountPoint {
	s.MountDir = &v
	return s
}

func (s *OSSMountPoint) SetReadOnly(v bool) *OSSMountPoint {
	s.ReadOnly = &v
	return s
}

func (s *OSSMountPoint) Validate() error {
	return dara.Validate(s)
}
