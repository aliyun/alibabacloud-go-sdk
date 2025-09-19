// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetsPropertyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBiz(v string) *GetAssetsPropertyDetailRequest
	GetBiz() *string
	SetCurrentPage(v int32) *GetAssetsPropertyDetailRequest
	GetCurrentPage() *int32
	SetItemName(v string) *GetAssetsPropertyDetailRequest
	GetItemName() *string
	SetLang(v string) *GetAssetsPropertyDetailRequest
	GetLang() *string
	SetPageSize(v int32) *GetAssetsPropertyDetailRequest
	GetPageSize() *int32
	SetRemark(v string) *GetAssetsPropertyDetailRequest
	GetRemark() *string
	SetSearchCriteriaList(v []*GetAssetsPropertyDetailRequestSearchCriteriaList) *GetAssetsPropertyDetailRequest
	GetSearchCriteriaList() []*GetAssetsPropertyDetailRequestSearchCriteriaList
	SetUuid(v string) *GetAssetsPropertyDetailRequest
	GetUuid() *string
}

type GetAssetsPropertyDetailRequest struct {
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
	// The name of the aggregation item for the asset fingerprint that you want to query.
	//
	// > You can call the [GetAssetsPropertyItem](~~GetAssetsPropertyItem~~) operation to query the names of aggregation items.
	//
	// example:
	//
	// virtio
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The conditions that are used to query the details about the asset fingerprint.
	SearchCriteriaList []*GetAssetsPropertyDetailRequestSearchCriteriaList `json:"SearchCriteriaList,omitempty" xml:"SearchCriteriaList,omitempty" type:"Repeated"`
	// The UUID of the server.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 38f72ea4-4c9f-4df1-bc6c-0f267614****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAssetsPropertyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyDetailRequest) GetBiz() *string {
	return s.Biz
}

func (s *GetAssetsPropertyDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAssetsPropertyDetailRequest) GetItemName() *string {
	return s.ItemName
}

func (s *GetAssetsPropertyDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAssetsPropertyDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAssetsPropertyDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *GetAssetsPropertyDetailRequest) GetSearchCriteriaList() []*GetAssetsPropertyDetailRequestSearchCriteriaList {
	return s.SearchCriteriaList
}

func (s *GetAssetsPropertyDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetAssetsPropertyDetailRequest) SetBiz(v string) *GetAssetsPropertyDetailRequest {
	s.Biz = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetCurrentPage(v int32) *GetAssetsPropertyDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetItemName(v string) *GetAssetsPropertyDetailRequest {
	s.ItemName = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetLang(v string) *GetAssetsPropertyDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetPageSize(v int32) *GetAssetsPropertyDetailRequest {
	s.PageSize = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetRemark(v string) *GetAssetsPropertyDetailRequest {
	s.Remark = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetSearchCriteriaList(v []*GetAssetsPropertyDetailRequestSearchCriteriaList) *GetAssetsPropertyDetailRequest {
	s.SearchCriteriaList = v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetUuid(v string) *GetAssetsPropertyDetailRequest {
	s.Uuid = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) Validate() error {
	return dara.Validate(s)
}

type GetAssetsPropertyDetailRequestSearchCriteriaList struct {
	// The name of the condition. Valid values:
	//
	// 	- **remarkItemName**: the aggregation item of the asset fingerprints. Fuzzy match is supported.
	//
	// > 	- If **Biz*	- is set to **web_server**, **remarkItemName*	- specifies a domain name.
	//
	// > 	- If **Biz*	- is set to **lkm**, **remarkItemName*	- specifies a module name.
	//
	// > 	- If **Biz*	- is set to **autorun**, **remarkItemName*	- specifies the path to a startup item.
	//
	// example:
	//
	// remarkItemName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the condition.
	//
	// example:
	//
	// virtio
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAssetsPropertyDetailRequestSearchCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyDetailRequestSearchCriteriaList) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyDetailRequestSearchCriteriaList) GetName() *string {
	return s.Name
}

func (s *GetAssetsPropertyDetailRequestSearchCriteriaList) GetValue() *string {
	return s.Value
}

func (s *GetAssetsPropertyDetailRequestSearchCriteriaList) SetName(v string) *GetAssetsPropertyDetailRequestSearchCriteriaList {
	s.Name = &v
	return s
}

func (s *GetAssetsPropertyDetailRequestSearchCriteriaList) SetValue(v string) *GetAssetsPropertyDetailRequestSearchCriteriaList {
	s.Value = &v
	return s
}

func (s *GetAssetsPropertyDetailRequestSearchCriteriaList) Validate() error {
	return dara.Validate(s)
}
