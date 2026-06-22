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
	// The list of check item IDs.
	CheckIdsShrink *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	// The check item name for fuzzy match.
	//
	// example:
	//
	// redis
	CheckItemFuzzy *string `json:"CheckItemFuzzy,omitempty" xml:"CheckItemFuzzy,omitempty"`
	// The page number of the current page when using paging. The value starts from 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for requests and responses. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page when using paging.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of whitelist record IDs.
	RecordIdsShrink *string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty"`
	// The data source. If this parameter is left empty, host data is queried by default. Valid values:
	//
	// - **default**: host
	//
	// - **agentless**: agentless.
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
