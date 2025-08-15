// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ListOssBucketRequest
	GetBucketName() *string
	SetLang(v string) *ListOssBucketRequest
	GetLang() *string
}

type ListOssBucketRequest struct {
	// example:
	//
	// testBucketName****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListOssBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketRequest) GoString() string {
	return s.String()
}

func (s *ListOssBucketRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ListOssBucketRequest) GetLang() *string {
	return s.Lang
}

func (s *ListOssBucketRequest) SetBucketName(v string) *ListOssBucketRequest {
	s.BucketName = &v
	return s
}

func (s *ListOssBucketRequest) SetLang(v string) *ListOssBucketRequest {
	s.Lang = &v
	return s
}

func (s *ListOssBucketRequest) Validate() error {
	return dara.Validate(s)
}
