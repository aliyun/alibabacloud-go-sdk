// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectPluginStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListFileProtectPluginStatusRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *ListFileProtectPluginStatusRequest
	GetPageSize() *int64
	SetSwitchId(v string) *ListFileProtectPluginStatusRequest
	GetSwitchId() *string
}

type ListFileProtectPluginStatusRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the core file monitoring rule.
	//
	// example:
	//
	// FILE_PROTECT_RULE_SWITCH_TYPE_1693474122927
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s ListFileProtectPluginStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectPluginStatusRequest) GoString() string {
	return s.String()
}

func (s *ListFileProtectPluginStatusRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListFileProtectPluginStatusRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFileProtectPluginStatusRequest) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListFileProtectPluginStatusRequest) SetCurrentPage(v int64) *ListFileProtectPluginStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectPluginStatusRequest) SetPageSize(v int64) *ListFileProtectPluginStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectPluginStatusRequest) SetSwitchId(v string) *ListFileProtectPluginStatusRequest {
	s.SwitchId = &v
	return s
}

func (s *ListFileProtectPluginStatusRequest) Validate() error {
	return dara.Validate(s)
}
