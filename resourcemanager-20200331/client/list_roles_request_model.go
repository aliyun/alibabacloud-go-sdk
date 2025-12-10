// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListRolesRequest
	GetLanguage() *string
	SetPageNumber(v int32) *ListRolesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRolesRequest
	GetPageSize() *int32
}

type ListRolesRequest struct {
	// The language in which you want to return the descriptions of the RAM roles. Valid values:
	//
	// 	- en: English
	//
	// 	- zh-CN: Chinese
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListRolesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRolesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRolesRequest) SetLanguage(v string) *ListRolesRequest {
	s.Language = &v
	return s
}

func (s *ListRolesRequest) SetPageNumber(v int32) *ListRolesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRolesRequest) SetPageSize(v int32) *ListRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
