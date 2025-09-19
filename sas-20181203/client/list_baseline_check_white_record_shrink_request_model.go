// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineCheckWhiteRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIdsShrink(v string) *ListBaselineCheckWhiteRecordShrinkRequest
	GetCheckIdsShrink() *string
	SetCheckItemFuzzy(v string) *ListBaselineCheckWhiteRecordShrinkRequest
	GetCheckItemFuzzy() *string
	SetCurrentPage(v int32) *ListBaselineCheckWhiteRecordShrinkRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListBaselineCheckWhiteRecordShrinkRequest
	GetLang() *string
	SetPageSize(v int32) *ListBaselineCheckWhiteRecordShrinkRequest
	GetPageSize() *int32
	SetRecordIdsShrink(v string) *ListBaselineCheckWhiteRecordShrinkRequest
	GetRecordIdsShrink() *string
	SetSource(v string) *ListBaselineCheckWhiteRecordShrinkRequest
	GetSource() *string
}

type ListBaselineCheckWhiteRecordShrinkRequest struct {
	// The IDs of check items.
	CheckIdsShrink *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	// The name of the check item. Fuzzy match is supported.
	//
	// example:
	//
	// redis
	CheckItemFuzzy *string `json:"CheckItemFuzzy,omitempty" xml:"CheckItemFuzzy,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the whitelist rules.
	RecordIdsShrink *string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty"`
	// The data source. If you leave this parameter empty, the default value is used. Valid values:
	//
	// 	- **default**: server
	//
	// 	- **agentless**: agentless detection
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListBaselineCheckWhiteRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineCheckWhiteRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) GetCheckIdsShrink() *string {
	return s.CheckIdsShrink
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) GetCheckItemFuzzy() *string {
	return s.CheckItemFuzzy
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) GetRecordIdsShrink() *string {
	return s.RecordIdsShrink
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) SetCheckIdsShrink(v string) *ListBaselineCheckWhiteRecordShrinkRequest {
	s.CheckIdsShrink = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) SetCheckItemFuzzy(v string) *ListBaselineCheckWhiteRecordShrinkRequest {
	s.CheckItemFuzzy = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) SetCurrentPage(v int32) *ListBaselineCheckWhiteRecordShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) SetLang(v string) *ListBaselineCheckWhiteRecordShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) SetPageSize(v int32) *ListBaselineCheckWhiteRecordShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) SetRecordIdsShrink(v string) *ListBaselineCheckWhiteRecordShrinkRequest {
	s.RecordIdsShrink = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) SetSource(v string) *ListBaselineCheckWhiteRecordShrinkRequest {
	s.Source = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
