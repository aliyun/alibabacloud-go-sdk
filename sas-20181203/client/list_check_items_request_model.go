// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *ListCheckItemsRequest
	GetCheckId() *int64
	SetCheckShowName(v string) *ListCheckItemsRequest
	GetCheckShowName() *string
	SetCheckTypes(v []*string) *ListCheckItemsRequest
	GetCheckTypes() []*string
	SetCurrentPage(v int32) *ListCheckItemsRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListCheckItemsRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckItemsRequest
	GetPageSize() *int32
	SetStatuses(v []*string) *ListCheckItemsRequest
	GetStatuses() []*string
}

type ListCheckItemsRequest struct {
	// The ID of the check item.
	//
	// example:
	//
	// 100000000001
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the custom check item.
	//
	// example:
	//
	// testCheckItemName
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
	// The source type of the situational awareness check item.
	CheckTypes []*string `json:"CheckTypes,omitempty" xml:"CheckTypes,omitempty" type:"Repeated"`
	// Specifies the page number to display when performing a paginated query. The starting value is **1**, and the default value is **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for requests and responses. The default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies the maximum number of data entries to display per page when performing a paginated query. The default number of data entries displayed per page is 20, and if the PageSize parameter is empty, it will default to returning 20 data entries.
	//
	// > It is recommended that the PageSize value is not left empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the check item.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListCheckItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemsRequest) GoString() string {
	return s.String()
}

func (s *ListCheckItemsRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckItemsRequest) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *ListCheckItemsRequest) GetCheckTypes() []*string {
	return s.CheckTypes
}

func (s *ListCheckItemsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckItemsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckItemsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListCheckItemsRequest) SetCheckId(v int64) *ListCheckItemsRequest {
	s.CheckId = &v
	return s
}

func (s *ListCheckItemsRequest) SetCheckShowName(v string) *ListCheckItemsRequest {
	s.CheckShowName = &v
	return s
}

func (s *ListCheckItemsRequest) SetCheckTypes(v []*string) *ListCheckItemsRequest {
	s.CheckTypes = v
	return s
}

func (s *ListCheckItemsRequest) SetCurrentPage(v int32) *ListCheckItemsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckItemsRequest) SetLang(v string) *ListCheckItemsRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckItemsRequest) SetPageSize(v int32) *ListCheckItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckItemsRequest) SetStatuses(v []*string) *ListCheckItemsRequest {
	s.Statuses = v
	return s
}

func (s *ListCheckItemsRequest) Validate() error {
	return dara.Validate(s)
}
