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
	SetApiVersion(v string) *ListAddonsRequest
	GetApiVersion() *string
	SetCategory(v string) *ListAddonsRequest
	GetCategory() *string
	SetRegexp(v bool) *ListAddonsRequest
	GetRegexp() *bool
	SetSearch(v string) *ListAddonsRequest
	GetSearch() *string
}

type ListAddonsRequest struct {
	// The locale. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The tag filter.
	//
	// example:
	//
	// database
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Specifies whether to use regular expression matching for the query field. Default value: false.
	//
	// example:
	//
	// false
	Regexp *bool `json:"regexp,omitempty" xml:"regexp,omitempty"`
	// The search query. Supports searching by add-on name, description, or keyword.
	//
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

func (s *ListAddonsRequest) GetApiVersion() *string {
	return s.ApiVersion
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

func (s *ListAddonsRequest) SetApiVersion(v string) *ListAddonsRequest {
	s.ApiVersion = &v
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
