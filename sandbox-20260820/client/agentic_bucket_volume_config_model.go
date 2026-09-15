// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticBucketVolumeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticBucket(v string) *AgenticBucketVolumeConfig
	GetAgenticBucket() *string
	SetBucketName(v string) *AgenticBucketVolumeConfig
	GetBucketName() *string
	SetBucketPath(v string) *AgenticBucketVolumeConfig
	GetBucketPath() *string
	SetEndpoint(v string) *AgenticBucketVolumeConfig
	GetEndpoint() *string
	SetReadOnly(v bool) *AgenticBucketVolumeConfig
	GetReadOnly() *bool
}

type AgenticBucketVolumeConfig struct {
	AgenticBucket *string `json:"agenticBucket,omitempty" xml:"agenticBucket,omitempty"`
	BucketName    *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath    *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	Endpoint      *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	ReadOnly      *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s AgenticBucketVolumeConfig) String() string {
	return dara.Prettify(s)
}

func (s AgenticBucketVolumeConfig) GoString() string {
	return s.String()
}

func (s *AgenticBucketVolumeConfig) GetAgenticBucket() *string {
	return s.AgenticBucket
}

func (s *AgenticBucketVolumeConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *AgenticBucketVolumeConfig) GetBucketPath() *string {
	return s.BucketPath
}

func (s *AgenticBucketVolumeConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *AgenticBucketVolumeConfig) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *AgenticBucketVolumeConfig) SetAgenticBucket(v string) *AgenticBucketVolumeConfig {
	s.AgenticBucket = &v
	return s
}

func (s *AgenticBucketVolumeConfig) SetBucketName(v string) *AgenticBucketVolumeConfig {
	s.BucketName = &v
	return s
}

func (s *AgenticBucketVolumeConfig) SetBucketPath(v string) *AgenticBucketVolumeConfig {
	s.BucketPath = &v
	return s
}

func (s *AgenticBucketVolumeConfig) SetEndpoint(v string) *AgenticBucketVolumeConfig {
	s.Endpoint = &v
	return s
}

func (s *AgenticBucketVolumeConfig) SetReadOnly(v bool) *AgenticBucketVolumeConfig {
	s.ReadOnly = &v
	return s
}

func (s *AgenticBucketVolumeConfig) Validate() error {
	return dara.Validate(s)
}
