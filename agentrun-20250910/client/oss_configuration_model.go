// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *OssConfiguration
	GetBucketName() *string
	SetMountPoint(v string) *OssConfiguration
	GetMountPoint() *string
	SetPermission(v string) *OssConfiguration
	GetPermission() *string
	SetPrefix(v string) *OssConfiguration
	GetPrefix() *string
	SetRegion(v string) *OssConfiguration
	GetRegion() *string
}

type OssConfiguration struct {
	// This parameter is required.
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// This parameter is required.
	MountPoint *string `json:"mountPoint,omitempty" xml:"mountPoint,omitempty"`
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// This parameter is required.
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// example:
	//
	// ch-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s OssConfiguration) String() string {
	return dara.Prettify(s)
}

func (s OssConfiguration) GoString() string {
	return s.String()
}

func (s *OssConfiguration) GetBucketName() *string {
	return s.BucketName
}

func (s *OssConfiguration) GetMountPoint() *string {
	return s.MountPoint
}

func (s *OssConfiguration) GetPermission() *string {
	return s.Permission
}

func (s *OssConfiguration) GetPrefix() *string {
	return s.Prefix
}

func (s *OssConfiguration) GetRegion() *string {
	return s.Region
}

func (s *OssConfiguration) SetBucketName(v string) *OssConfiguration {
	s.BucketName = &v
	return s
}

func (s *OssConfiguration) SetMountPoint(v string) *OssConfiguration {
	s.MountPoint = &v
	return s
}

func (s *OssConfiguration) SetPermission(v string) *OssConfiguration {
	s.Permission = &v
	return s
}

func (s *OssConfiguration) SetPrefix(v string) *OssConfiguration {
	s.Prefix = &v
	return s
}

func (s *OssConfiguration) SetRegion(v string) *OssConfiguration {
	s.Region = &v
	return s
}

func (s *OssConfiguration) Validate() error {
	return dara.Validate(s)
}
