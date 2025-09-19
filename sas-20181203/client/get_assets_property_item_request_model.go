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
	// The type of the asset fingerprint that you want to query. Default value: **sca**. Valid values:
	//
	// 	- **lkm**: kernel module
	//
	// 	- **autorun**: startup item
	//
	// 	- **web_server**: website
	//
	// This parameter is required.
	//
	// example:
	//
	// lkm
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to forcefully refresh the data that you want to query. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	ForceFlush *bool `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The value of the search condition. You must specify this parameter based on the value of the **SearchItem*	- parameter.
	//
	// 	- If you set **SearchItem*	- to **domain**, you must enter the domain name.
	//
	// 	- If you set **SearchItem*	- to **module_name**, you must enter the module name.
	//
	// 	- If you set **SearchItem*	- to **path**, you must enter the path to the self-starting item.
	//
	// >  You must specify both the **SearchItem*	- and **SearchInfo*	- parameters before you can query the information about asset fingerprints by asset fingerprint name.
	//
	// example:
	//
	// /lib/systemd/s****
	SearchInfo *string `json:"SearchInfo,omitempty" xml:"SearchInfo,omitempty"`
	// The type of the search condition. You must specify this parameter based on the value of the **Biz*	- parameter. Valid values:
	//
	// 	- If you set **Biz*	- to **web_server**, set **SearchItem*	- to the following value:
	//
	//     	- **domain**: the domain name
	//
	// 	- If you set **Biz*	- to **lkm**, set **SearchItem*	- to the following value:
	//
	//     	- **module_name**: the name of the module
	//
	// 	- If you set **Biz*	- to **autorun**, set **SearchItem*	- to the following value:
	//
	//     	- **path**: the path to the self-starting item
	//
	// >  You must specify both the **SearchItem*	- and **SearchInfo*	- parameters before you can query the information about asset fingerprints by asset fingerprint name.
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
