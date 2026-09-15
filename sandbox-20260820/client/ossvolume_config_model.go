// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOSSVolumeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *OSSVolumeConfig
	GetBucketName() *string
	SetBucketPath(v string) *OSSVolumeConfig
	GetBucketPath() *string
	SetEndpoint(v string) *OSSVolumeConfig
	GetEndpoint() *string
	SetReadOnly(v bool) *OSSVolumeConfig
	GetReadOnly() *bool
}

type OSSVolumeConfig struct {
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s OSSVolumeConfig) String() string {
	return dara.Prettify(s)
}

func (s OSSVolumeConfig) GoString() string {
	return s.String()
}

func (s *OSSVolumeConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *OSSVolumeConfig) GetBucketPath() *string {
	return s.BucketPath
}

func (s *OSSVolumeConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *OSSVolumeConfig) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *OSSVolumeConfig) SetBucketName(v string) *OSSVolumeConfig {
	s.BucketName = &v
	return s
}

func (s *OSSVolumeConfig) SetBucketPath(v string) *OSSVolumeConfig {
	s.BucketPath = &v
	return s
}

func (s *OSSVolumeConfig) SetEndpoint(v string) *OSSVolumeConfig {
	s.Endpoint = &v
	return s
}

func (s *OSSVolumeConfig) SetReadOnly(v bool) *OSSVolumeConfig {
	s.ReadOnly = &v
	return s
}

func (s *OSSVolumeConfig) Validate() error {
	return dara.Validate(s)
}
