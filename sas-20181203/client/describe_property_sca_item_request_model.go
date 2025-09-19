// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBiz(v string) *DescribePropertyScaItemRequest
	GetBiz() *string
	SetCurrentPage(v int32) *DescribePropertyScaItemRequest
	GetCurrentPage() *int32
	SetForceFlush(v bool) *DescribePropertyScaItemRequest
	GetForceFlush() *bool
	SetLang(v string) *DescribePropertyScaItemRequest
	GetLang() *string
	SetPageSize(v int32) *DescribePropertyScaItemRequest
	GetPageSize() *int32
	SetSearchInfo(v string) *DescribePropertyScaItemRequest
	GetSearchInfo() *string
	SetSearchItem(v string) *DescribePropertyScaItemRequest
	GetSearchItem() *string
}

type DescribePropertyScaItemRequest struct {
	// The type of the asset fingerprint that you want to query. Default value: **sca**. Valid values:
	//
	// 	- **sca**: middleware
	//
	// 	- **sca_database**: database
	//
	// 	- **sca_web**: web service
	//
	// > If you do not specify this parameter, the default value **sca*	- is used, which indicates that middleware fingerprints are queried.
	//
	// example:
	//
	// sca
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
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword. You must specify this parameter based on the value of the **SearchItem*	- parameter.
	//
	// 	- If the **SearchItem*	- parameter is set to **name**, you must enter the name of an asset fingerprint.
	//
	// 	- If the **SearchItem*	- parameter is set to **type**, you must enter the type of an asset fingerprint. Valid values:
	//
	//     	- **system_service**: system service
	//
	//     	- **software_library**: software library
	//
	//     	- **docker_component**: container component
	//
	//     	- **database**: database
	//
	//     	- **web_container**: web container
	//
	//     	- **jar**: JAR package
	//
	//     	- **web_framework**: web framework
	//
	// > You must specify both the **SearchItem*	- and **SearchInfo*	- parameters before you can query the asset fingerprints based on the specified name or type.
	//
	// example:
	//
	// system_service
	SearchInfo *string `json:"SearchInfo,omitempty" xml:"SearchInfo,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **name**: the name of a database, middleware, or web service
	//
	// 	- **type**: the type of a database, middleware, or web service
	//
	// > You must specify both the **SearchItem*	- and **SearchInfo*	- parameters before you can query the asset fingerprints based on the specified name or type.
	//
	// example:
	//
	// type
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
}

func (s DescribePropertyScaItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaItemRequest) GetBiz() *string {
	return s.Biz
}

func (s *DescribePropertyScaItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyScaItemRequest) GetForceFlush() *bool {
	return s.ForceFlush
}

func (s *DescribePropertyScaItemRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePropertyScaItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyScaItemRequest) GetSearchInfo() *string {
	return s.SearchInfo
}

func (s *DescribePropertyScaItemRequest) GetSearchItem() *string {
	return s.SearchItem
}

func (s *DescribePropertyScaItemRequest) SetBiz(v string) *DescribePropertyScaItemRequest {
	s.Biz = &v
	return s
}

func (s *DescribePropertyScaItemRequest) SetCurrentPage(v int32) *DescribePropertyScaItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaItemRequest) SetForceFlush(v bool) *DescribePropertyScaItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyScaItemRequest) SetLang(v string) *DescribePropertyScaItemRequest {
	s.Lang = &v
	return s
}

func (s *DescribePropertyScaItemRequest) SetPageSize(v int32) *DescribePropertyScaItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyScaItemRequest) SetSearchInfo(v string) *DescribePropertyScaItemRequest {
	s.SearchInfo = &v
	return s
}

func (s *DescribePropertyScaItemRequest) SetSearchItem(v string) *DescribePropertyScaItemRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribePropertyScaItemRequest) Validate() error {
	return dara.Validate(s)
}
