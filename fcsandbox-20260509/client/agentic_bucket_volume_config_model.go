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
	// example:
	//
	// absx9f61c8a-1311999136518149-cn-hangzhou-ab-apsr
	AgenticBucket *string `json:"agenticBucket,omitempty" xml:"agenticBucket,omitempty"`
	// example:
	//
	// bs429pop1-1311999136518149-cn-hangzhou-bs-apsr
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// /test/
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// example:
	//
	// https://oss-cn-hangzhou-internal.aliyuncs.com"
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
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
