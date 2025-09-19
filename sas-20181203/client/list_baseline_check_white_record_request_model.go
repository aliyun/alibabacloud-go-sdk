// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineCheckWhiteRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *ListBaselineCheckWhiteRecordRequest
	GetCheckIds() []*int64
	SetCheckItemFuzzy(v string) *ListBaselineCheckWhiteRecordRequest
	GetCheckItemFuzzy() *string
	SetCurrentPage(v int32) *ListBaselineCheckWhiteRecordRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListBaselineCheckWhiteRecordRequest
	GetLang() *string
	SetPageSize(v int32) *ListBaselineCheckWhiteRecordRequest
	GetPageSize() *int32
	SetRecordIds(v []*int64) *ListBaselineCheckWhiteRecordRequest
	GetRecordIds() []*int64
	SetSource(v string) *ListBaselineCheckWhiteRecordRequest
	GetSource() *string
}

type ListBaselineCheckWhiteRecordRequest struct {
	// The IDs of check items.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
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
	RecordIds []*int64 `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
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

func (s ListBaselineCheckWhiteRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineCheckWhiteRecordRequest) GoString() string {
	return s.String()
}

func (s *ListBaselineCheckWhiteRecordRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *ListBaselineCheckWhiteRecordRequest) GetCheckItemFuzzy() *string {
	return s.CheckItemFuzzy
}

func (s *ListBaselineCheckWhiteRecordRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListBaselineCheckWhiteRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *ListBaselineCheckWhiteRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselineCheckWhiteRecordRequest) GetRecordIds() []*int64 {
	return s.RecordIds
}

func (s *ListBaselineCheckWhiteRecordRequest) GetSource() *string {
	return s.Source
}

func (s *ListBaselineCheckWhiteRecordRequest) SetCheckIds(v []*int64) *ListBaselineCheckWhiteRecordRequest {
	s.CheckIds = v
	return s
}

func (s *ListBaselineCheckWhiteRecordRequest) SetCheckItemFuzzy(v string) *ListBaselineCheckWhiteRecordRequest {
	s.CheckItemFuzzy = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordRequest) SetCurrentPage(v int32) *ListBaselineCheckWhiteRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordRequest) SetLang(v string) *ListBaselineCheckWhiteRecordRequest {
	s.Lang = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordRequest) SetPageSize(v int32) *ListBaselineCheckWhiteRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordRequest) SetRecordIds(v []*int64) *ListBaselineCheckWhiteRecordRequest {
	s.RecordIds = v
	return s
}

func (s *ListBaselineCheckWhiteRecordRequest) SetSource(v string) *ListBaselineCheckWhiteRecordRequest {
	s.Source = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordRequest) Validate() error {
	return dara.Validate(s)
}
