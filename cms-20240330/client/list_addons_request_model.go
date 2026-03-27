// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *ListAddonsRequest
	GetAliyunLang() *string
	SetCategory(v string) *ListAddonsRequest
	GetCategory() *string
	SetRegexp(v bool) *ListAddonsRequest
	GetRegexp() *bool
	SetSearch(v string) *ListAddonsRequest
	GetSearch() *string
}

type ListAddonsRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// example:
	//
	// database
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	Regexp   *bool   `json:"regexp,omitempty" xml:"regexp,omitempty"`
	// example:
	//
	// 105095
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
}

func (s ListAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsRequest) GoString() string {
	return s.String()
}

func (s *ListAddonsRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *ListAddonsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListAddonsRequest) GetRegexp() *bool {
	return s.Regexp
}

func (s *ListAddonsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListAddonsRequest) SetAliyunLang(v string) *ListAddonsRequest {
	s.AliyunLang = &v
	return s
}

func (s *ListAddonsRequest) SetCategory(v string) *ListAddonsRequest {
	s.Category = &v
	return s
}

func (s *ListAddonsRequest) SetRegexp(v bool) *ListAddonsRequest {
	s.Regexp = &v
	return s
}

func (s *ListAddonsRequest) SetSearch(v string) *ListAddonsRequest {
	s.Search = &v
	return s
}

func (s *ListAddonsRequest) Validate() error {
	return dara.Validate(s)
}
