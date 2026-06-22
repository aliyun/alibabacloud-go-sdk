// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetsPropertyItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBiz(v string) *GetAssetsPropertyItemRequest
	GetBiz() *string
	SetCurrentPage(v int32) *GetAssetsPropertyItemRequest
	GetCurrentPage() *int32
	SetForceFlush(v bool) *GetAssetsPropertyItemRequest
	GetForceFlush() *bool
	SetLang(v string) *GetAssetsPropertyItemRequest
	GetLang() *string
	SetPageSize(v int32) *GetAssetsPropertyItemRequest
	GetPageSize() *int32
	SetSearchInfo(v string) *GetAssetsPropertyItemRequest
	GetSearchInfo() *string
	SetSearchItem(v string) *GetAssetsPropertyItemRequest
	GetSearchItem() *string
}

type GetAssetsPropertyItemRequest struct {
	// The type of Asset Fingerprints to query. Default value: **sca**. Valid values:
	//
	// - **lkm**: kernel module
	//
	// - **autorun**: startup item
	//
	// - **web_server**: website.
	//
	// This parameter is required.
	//
	// example:
	//
	// lkm
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The page number of the page to return. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to forcefully refresh the data to be queried. Valid values:
	//
	// - **true**: Forcefully refresh.
	//
	// - **false**: Do not forcefully refresh.
	//
	// example:
	//
	// false
	ForceFlush *bool `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page when using paging. Default value: 20. If the PageSize parameter is left empty, 20 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The content to query. Specify different query content based on the value of **SearchItem**:
	//
	// - If **SearchItem*	- is set to **domain**, enter the domain name of the Asset Fingerprints entry.
	//
	// - If **SearchItem*	- is set to **module_name**, enter the module name of the Asset Fingerprints entry.
	//
	// - If **SearchItem*	- is set to **path**, enter the startup item path of the Asset Fingerprints entry.
	//
	// > The **SearchItem*	- and **SearchInfo*	- parameters must be used together. Both parameters must be set at the same time for the query to take effect. Setting only one parameter does not take effect. You can use these parameters to query all data of a specific Asset Fingerprints entry by name.
	//
	// example:
	//
	// /lib/systemd/s****
	SearchInfo *string `json:"SearchInfo,omitempty" xml:"SearchInfo,omitempty"`
	// The type of query condition. Set different aggregation search conditions based on the **Biz*	- parameter. Valid values:
	//
	// - If **Biz*	- is set to **web_server**, the following search conditions are supported for **SearchItem**:
	//
	//     - **domain**: domain name
	//
	// - If **Biz*	- is set to **lkm**, the following search conditions are supported for **SearchItem**:
	//
	//     - **module_name**: module name
	//
	// - If **Biz*	- is set to **autorun**, the following search conditions are supported for **SearchItem**:
	//
	//     - **path**: startup item path
	//
	// > The **SearchItem*	- and **SearchInfo*	- parameters must be used together. Both parameters must be set at the same time for the query to take effect. Setting only one parameter does not take effect. You can use these parameters to query all data of a specific Asset Fingerprints entry by name.
	//
	// example:
	//
	// path
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
}

func (s GetAssetsPropertyItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyItemRequest) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyItemRequest) GetBiz() *string {
	return s.Biz
}

func (s *GetAssetsPropertyItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAssetsPropertyItemRequest) GetForceFlush() *bool {
	return s.ForceFlush
}

func (s *GetAssetsPropertyItemRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAssetsPropertyItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAssetsPropertyItemRequest) GetSearchInfo() *string {
	return s.SearchInfo
}

func (s *GetAssetsPropertyItemRequest) GetSearchItem() *string {
	return s.SearchItem
}

func (s *GetAssetsPropertyItemRequest) SetBiz(v string) *GetAssetsPropertyItemRequest {
	s.Biz = &v
	return s
}

func (s *GetAssetsPropertyItemRequest) SetCurrentPage(v int32) *GetAssetsPropertyItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAssetsPropertyItemRequest) SetForceFlush(v bool) *GetAssetsPropertyItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *GetAssetsPropertyItemRequest) SetLang(v string) *GetAssetsPropertyItemRequest {
	s.Lang = &v
	return s
}

func (s *GetAssetsPropertyItemRequest) SetPageSize(v int32) *GetAssetsPropertyItemRequest {
	s.PageSize = &v
	return s
}

func (s *GetAssetsPropertyItemRequest) SetSearchInfo(v string) *GetAssetsPropertyItemRequest {
	s.SearchInfo = &v
	return s
}

func (s *GetAssetsPropertyItemRequest) SetSearchItem(v string) *GetAssetsPropertyItemRequest {
	s.SearchItem = &v
	return s
}

func (s *GetAssetsPropertyItemRequest) Validate() error {
	return dara.Validate(s)
}
