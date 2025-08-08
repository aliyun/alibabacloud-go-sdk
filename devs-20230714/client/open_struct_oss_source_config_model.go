// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructOssSourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *OpenStructOssSourceConfig
	GetBucket() *string
	SetObject(v string) *OpenStructOssSourceConfig
	GetObject() *string
}

type OpenStructOssSourceConfig struct {
	// example:
	//
	// demo-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// demo-object
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
}

func (s OpenStructOssSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s OpenStructOssSourceConfig) GoString() string {
	return s.String()
}

func (s *OpenStructOssSourceConfig) GetBucket() *string {
	return s.Bucket
}

func (s *OpenStructOssSourceConfig) GetObject() *string {
	return s.Object
}

func (s *OpenStructOssSourceConfig) SetBucket(v string) *OpenStructOssSourceConfig {
	s.Bucket = &v
	return s
}

func (s *OpenStructOssSourceConfig) SetObject(v string) *OpenStructOssSourceConfig {
	s.Object = &v
	return s
}

func (s *OpenStructOssSourceConfig) Validate() error {
	return dara.Validate(s)
}
