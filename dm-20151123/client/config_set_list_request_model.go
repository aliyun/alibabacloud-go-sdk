// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ConfigSetListRequest
	GetAll() *bool
	SetKeyword(v string) *ConfigSetListRequest
	GetKeyword() *string
	SetPageIndex(v string) *ConfigSetListRequest
	GetPageIndex() *string
	SetPageSize(v string) *ConfigSetListRequest
	GetPageSize() *string
}

type ConfigSetListRequest struct {
	// Specifies whether to query all configuration sets. This parameter is optional.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The keyword to search by name.
	//
	// example:
	//
	// xxx
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page index. The value starts from 1. This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page. This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ConfigSetListRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetListRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetListRequest) GetAll() *bool {
	return s.All
}

func (s *ConfigSetListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ConfigSetListRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *ConfigSetListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ConfigSetListRequest) SetAll(v bool) *ConfigSetListRequest {
	s.All = &v
	return s
}

func (s *ConfigSetListRequest) SetKeyword(v string) *ConfigSetListRequest {
	s.Keyword = &v
	return s
}

func (s *ConfigSetListRequest) SetPageIndex(v string) *ConfigSetListRequest {
	s.PageIndex = &v
	return s
}

func (s *ConfigSetListRequest) SetPageSize(v string) *ConfigSetListRequest {
	s.PageSize = &v
	return s
}

func (s *ConfigSetListRequest) Validate() error {
	return dara.Validate(s)
}
