// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetsPropertyItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *GetAssetsPropertyItemResponseBodyPageInfo) *GetAssetsPropertyItemResponseBody
	GetPageInfo() *GetAssetsPropertyItemResponseBodyPageInfo
	SetPropertyItems(v []*GetAssetsPropertyItemResponseBodyPropertyItems) *GetAssetsPropertyItemResponseBody
	GetPropertyItems() []*GetAssetsPropertyItemResponseBodyPropertyItems
	SetRequestId(v string) *GetAssetsPropertyItemResponseBody
	GetRequestId() *string
}

type GetAssetsPropertyItemResponseBody struct {
	// The pagination information.
	PageInfo *GetAssetsPropertyItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the aggregation information about asset fingerprints.
	PropertyItems []*GetAssetsPropertyItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 340D7FC4-D575-1661-8ACD-CFA7BE57****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAssetsPropertyItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyItemResponseBody) GetPageInfo() *GetAssetsPropertyItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetAssetsPropertyItemResponseBody) GetPropertyItems() []*GetAssetsPropertyItemResponseBodyPropertyItems {
	return s.PropertyItems
}

func (s *GetAssetsPropertyItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetsPropertyItemResponseBody) SetPageInfo(v *GetAssetsPropertyItemResponseBodyPageInfo) *GetAssetsPropertyItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetAssetsPropertyItemResponseBody) SetPropertyItems(v []*GetAssetsPropertyItemResponseBodyPropertyItems) *GetAssetsPropertyItemResponseBody {
	s.PropertyItems = v
	return s
}

func (s *GetAssetsPropertyItemResponseBody) SetRequestId(v string) *GetAssetsPropertyItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAssetsPropertyItemResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAssetsPropertyItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) SetCount(v int32) *GetAssetsPropertyItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) SetCurrentPage(v int32) *GetAssetsPropertyItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) SetPageSize(v int32) *GetAssetsPropertyItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) SetTotalCount(v int32) *GetAssetsPropertyItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type GetAssetsPropertyItemResponseBodyPropertyItems struct {
	// The number of servers related to the asset fingerprints.
	//
	// example:
	//
	// 23
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The endpoint.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// localhost
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	MiddlewareName *string `json:"MiddlewareName,omitempty" xml:"MiddlewareName,omitempty"`
	ModelName      *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The name of the module.
	//
	// > This parameter is returned only when **Biz*	- is set to **lkm**.
	//
	// example:
	//
	// alihids
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The path to the startup item.
	//
	// > This parameter is returned only when **Biz*	- is set to **autorun**.
	//
	// example:
	//
	// C:/Program Files/****
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetAssetsPropertyItemResponseBodyPropertyItems) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) GetCount() *int32 {
	return s.Count
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) GetDomain() *string {
	return s.Domain
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) GetMiddlewareName() *string {
	return s.MiddlewareName
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) GetModelName() *string {
	return s.ModelName
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) GetPath() *string {
	return s.Path
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) SetCount(v int32) *GetAssetsPropertyItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) SetDomain(v string) *GetAssetsPropertyItemResponseBodyPropertyItems {
	s.Domain = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) SetMiddlewareName(v string) *GetAssetsPropertyItemResponseBodyPropertyItems {
	s.MiddlewareName = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) SetModelName(v string) *GetAssetsPropertyItemResponseBodyPropertyItems {
	s.ModelName = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) SetModuleName(v string) *GetAssetsPropertyItemResponseBodyPropertyItems {
	s.ModuleName = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) SetPath(v string) *GetAssetsPropertyItemResponseBodyPropertyItems {
	s.Path = &v
	return s
}

func (s *GetAssetsPropertyItemResponseBodyPropertyItems) Validate() error {
	return dara.Validate(s)
}
