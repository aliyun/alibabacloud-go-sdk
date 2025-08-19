// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInputCodeLocation interface {
	dara.Model
	String() string
	GoString() string
	SetChecksum(v string) *InputCodeLocation
	GetChecksum() *string
	SetOssBucketName(v string) *InputCodeLocation
	GetOssBucketName() *string
	SetOssObjectName(v string) *InputCodeLocation
	GetOssObjectName() *string
	SetZipFile(v string) *InputCodeLocation
	GetZipFile() *string
}

type InputCodeLocation struct {
	// example:
	//
	// 2825179536350****
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// example:
	//
	// demo-bucket
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// example:
	//
	// demo-object
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// example:
	//
	// UEsDBAoAAAAAANF
	ZipFile *string `json:"zipFile,omitempty" xml:"zipFile,omitempty"`
}

func (s InputCodeLocation) String() string {
	return dara.Prettify(s)
}

func (s InputCodeLocation) GoString() string {
	return s.String()
}

func (s *InputCodeLocation) GetChecksum() *string {
	return s.Checksum
}

func (s *InputCodeLocation) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *InputCodeLocation) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *InputCodeLocation) GetZipFile() *string {
	return s.ZipFile
}

func (s *InputCodeLocation) SetChecksum(v string) *InputCodeLocation {
	s.Checksum = &v
	return s
}

func (s *InputCodeLocation) SetOssBucketName(v string) *InputCodeLocation {
	s.OssBucketName = &v
	return s
}

func (s *InputCodeLocation) SetOssObjectName(v string) *InputCodeLocation {
	s.OssObjectName = &v
	return s
}

func (s *InputCodeLocation) SetZipFile(v string) *InputCodeLocation {
	s.ZipFile = &v
	return s
}

func (s *InputCodeLocation) Validate() error {
	return dara.Validate(s)
}
