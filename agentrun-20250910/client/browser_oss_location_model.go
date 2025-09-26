// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserOssLocation interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *BrowserOssLocation
	GetBucket() *string
	SetPrefix(v string) *BrowserOssLocation
	GetPrefix() *string
}

type BrowserOssLocation struct {
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s BrowserOssLocation) String() string {
	return dara.Prettify(s)
}

func (s BrowserOssLocation) GoString() string {
	return s.String()
}

func (s *BrowserOssLocation) GetBucket() *string {
	return s.Bucket
}

func (s *BrowserOssLocation) GetPrefix() *string {
	return s.Prefix
}

func (s *BrowserOssLocation) SetBucket(v string) *BrowserOssLocation {
	s.Bucket = &v
	return s
}

func (s *BrowserOssLocation) SetPrefix(v string) *BrowserOssLocation {
	s.Prefix = &v
	return s
}

func (s *BrowserOssLocation) Validate() error {
	return dara.Validate(s)
}
