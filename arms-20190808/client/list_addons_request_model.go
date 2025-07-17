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
	SetRegionId(v string) *ListAddonsRequest
	GetRegionId() *string
	SetSearch(v string) *ListAddonsRequest
	GetSearch() *string
}

type ListAddonsRequest struct {
	// Language,the default language is Chinese.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// Category filter.
	//
	// example:
	//
	// database
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Whether to enable regular matching.
	//
	// example:
	//
	// false
	Regexp *bool `json:"Regexp,omitempty" xml:"Regexp,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A query field can be queried by name or description.
	//
	// example:
	//
	// mysql
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
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

func (s *ListAddonsRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *ListAddonsRequest) SetRegionId(v string) *ListAddonsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAddonsRequest) SetSearch(v string) *ListAddonsRequest {
	s.Search = &v
	return s
}

func (s *ListAddonsRequest) Validate() error {
	return dara.Validate(s)
}
