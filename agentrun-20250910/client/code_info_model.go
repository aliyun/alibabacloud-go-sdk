// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeInfo interface {
	dara.Model
	String() string
	GoString() string
	SetOssBucketName(v string) *CodeInfo
	GetOssBucketName() *string
	SetOssObjectName(v string) *CodeInfo
	GetOssObjectName() *string
}

type CodeInfo struct {
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
}

func (s CodeInfo) String() string {
	return dara.Prettify(s)
}

func (s CodeInfo) GoString() string {
	return s.String()
}

func (s *CodeInfo) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CodeInfo) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *CodeInfo) SetOssBucketName(v string) *CodeInfo {
	s.OssBucketName = &v
	return s
}

func (s *CodeInfo) SetOssObjectName(v string) *CodeInfo {
	s.OssObjectName = &v
	return s
}

func (s *CodeInfo) Validate() error {
	return dara.Validate(s)
}
