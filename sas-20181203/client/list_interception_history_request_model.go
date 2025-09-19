// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInterceptionHistoryRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *ListInterceptionHistoryRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListInterceptionHistoryRequest
	GetEndTime() *int64
	SetHistoryName(v string) *ListInterceptionHistoryRequest
	GetHistoryName() *string
	SetInterceptionTypes(v []*int32) *ListInterceptionHistoryRequest
	GetInterceptionTypes() []*int32
	SetLang(v string) *ListInterceptionHistoryRequest
	GetLang() *string
	SetPageSize(v int32) *ListInterceptionHistoryRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListInterceptionHistoryRequest
	GetStartTime() *int64
}

type ListInterceptionHistoryRequest struct {
	// The ID of the container cluster.
	//
	// example:
	//
	// c7c190a82d9a048be9038d352840f****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1635575219000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// Abnormal access
	HistoryName *string `json:"HistoryName,omitempty" xml:"HistoryName,omitempty"`
	// The types of exceptions.
	InterceptionTypes []*int32 `json:"InterceptionTypes,omitempty" xml:"InterceptionTypes,omitempty" type:"Repeated"`
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
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start of the time range to query. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1651290987000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListInterceptionHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListInterceptionHistoryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInterceptionHistoryRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInterceptionHistoryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListInterceptionHistoryRequest) GetHistoryName() *string {
	return s.HistoryName
}

func (s *ListInterceptionHistoryRequest) GetInterceptionTypes() []*int32 {
	return s.InterceptionTypes
}

func (s *ListInterceptionHistoryRequest) GetLang() *string {
	return s.Lang
}

func (s *ListInterceptionHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterceptionHistoryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListInterceptionHistoryRequest) SetClusterId(v string) *ListInterceptionHistoryRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInterceptionHistoryRequest) SetCurrentPage(v int32) *ListInterceptionHistoryRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInterceptionHistoryRequest) SetEndTime(v int64) *ListInterceptionHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListInterceptionHistoryRequest) SetHistoryName(v string) *ListInterceptionHistoryRequest {
	s.HistoryName = &v
	return s
}

func (s *ListInterceptionHistoryRequest) SetInterceptionTypes(v []*int32) *ListInterceptionHistoryRequest {
	s.InterceptionTypes = v
	return s
}

func (s *ListInterceptionHistoryRequest) SetLang(v string) *ListInterceptionHistoryRequest {
	s.Lang = &v
	return s
}

func (s *ListInterceptionHistoryRequest) SetPageSize(v int32) *ListInterceptionHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterceptionHistoryRequest) SetStartTime(v int64) *ListInterceptionHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *ListInterceptionHistoryRequest) Validate() error {
	return dara.Validate(s)
}
