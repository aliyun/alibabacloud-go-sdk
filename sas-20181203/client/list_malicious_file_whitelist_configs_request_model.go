// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaliciousFileWhitelistConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListMaliciousFileWhitelistConfigsRequest
	GetCurrentPage() *int32
	SetEventName(v string) *ListMaliciousFileWhitelistConfigsRequest
	GetEventName() *string
	SetIdList(v int64) *ListMaliciousFileWhitelistConfigsRequest
	GetIdList() *int64
	SetLang(v string) *ListMaliciousFileWhitelistConfigsRequest
	GetLang() *string
	SetPageSize(v int32) *ListMaliciousFileWhitelistConfigsRequest
	GetPageSize() *int32
	SetSource(v string) *ListMaliciousFileWhitelistConfigsRequest
	GetSource() *string
}

type ListMaliciousFileWhitelistConfigsRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the alert.
	//
	// 	- Set the value to ALL, which indicates all alert types.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	IdList    *int64  `json:"IdList,omitempty" xml:"IdList,omitempty"`
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
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The feature to which this operation belongs. If you leave this parameter empty, the default value agentless is used.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListMaliciousFileWhitelistConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMaliciousFileWhitelistConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListMaliciousFileWhitelistConfigsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMaliciousFileWhitelistConfigsRequest) GetEventName() *string {
	return s.EventName
}

func (s *ListMaliciousFileWhitelistConfigsRequest) GetIdList() *int64 {
	return s.IdList
}

func (s *ListMaliciousFileWhitelistConfigsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListMaliciousFileWhitelistConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMaliciousFileWhitelistConfigsRequest) GetSource() *string {
	return s.Source
}

func (s *ListMaliciousFileWhitelistConfigsRequest) SetCurrentPage(v int32) *ListMaliciousFileWhitelistConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsRequest) SetEventName(v string) *ListMaliciousFileWhitelistConfigsRequest {
	s.EventName = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsRequest) SetIdList(v int64) *ListMaliciousFileWhitelistConfigsRequest {
	s.IdList = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsRequest) SetLang(v string) *ListMaliciousFileWhitelistConfigsRequest {
	s.Lang = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsRequest) SetPageSize(v int32) *ListMaliciousFileWhitelistConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsRequest) SetSource(v string) *ListMaliciousFileWhitelistConfigsRequest {
	s.Source = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsRequest) Validate() error {
	return dara.Validate(s)
}
