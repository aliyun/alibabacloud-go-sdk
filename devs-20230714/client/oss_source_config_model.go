// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssSourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *OssSourceConfig
	GetBucket() *string
	SetObject(v string) *OssSourceConfig
	GetObject() *string
}

type OssSourceConfig struct {
	// example:
	//
	// demo-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// demo-object
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
}

func (s OssSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s OssSourceConfig) GoString() string {
	return s.String()
}

func (s *OssSourceConfig) GetBucket() *string {
	return s.Bucket
}

func (s *OssSourceConfig) GetObject() *string {
	return s.Object
}

func (s *OssSourceConfig) SetBucket(v string) *OssSourceConfig {
	s.Bucket = &v
	return s
}

func (s *OssSourceConfig) SetObject(v string) *OssSourceConfig {
	s.Object = &v
	return s
}

func (s *OssSourceConfig) Validate() error {
	return dara.Validate(s)
}
