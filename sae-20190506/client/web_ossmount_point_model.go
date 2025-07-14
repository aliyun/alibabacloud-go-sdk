// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebOSSMountPoint interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *WebOSSMountPoint
	GetBucketName() *string
	SetBucketPath(v string) *WebOSSMountPoint
	GetBucketPath() *string
	SetMountDir(v string) *WebOSSMountPoint
	GetMountDir() *string
	SetReadOnly(v bool) *WebOSSMountPoint
	GetReadOnly() *bool
}

type WebOSSMountPoint struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketPath *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	MountDir   *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	ReadOnly   *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s WebOSSMountPoint) String() string {
	return dara.Prettify(s)
}

func (s WebOSSMountPoint) GoString() string {
	return s.String()
}

func (s *WebOSSMountPoint) GetBucketName() *string {
	return s.BucketName
}

func (s *WebOSSMountPoint) GetBucketPath() *string {
	return s.BucketPath
}

func (s *WebOSSMountPoint) GetMountDir() *string {
	return s.MountDir
}

func (s *WebOSSMountPoint) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *WebOSSMountPoint) SetBucketName(v string) *WebOSSMountPoint {
	s.BucketName = &v
	return s
}

func (s *WebOSSMountPoint) SetBucketPath(v string) *WebOSSMountPoint {
	s.BucketPath = &v
	return s
}

func (s *WebOSSMountPoint) SetMountDir(v string) *WebOSSMountPoint {
	s.MountDir = &v
	return s
}

func (s *WebOSSMountPoint) SetReadOnly(v bool) *WebOSSMountPoint {
	s.ReadOnly = &v
	return s
}

func (s *WebOSSMountPoint) Validate() error {
	return dara.Validate(s)
}
