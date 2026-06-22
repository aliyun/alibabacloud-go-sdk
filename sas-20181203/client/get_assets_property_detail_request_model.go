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
	SetNextToken(v string) *GetAssetsPropertyDetailRequest
	GetNextToken() *string
	SetPageSize(v int32) *GetAssetsPropertyDetailRequest
	GetPageSize() *int32
	SetRemark(v string) *GetAssetsPropertyDetailRequest
	GetRemark() *string
	SetSearchCriteriaList(v []*GetAssetsPropertyDetailRequestSearchCriteriaList) *GetAssetsPropertyDetailRequest
	GetSearchCriteriaList() []*GetAssetsPropertyDetailRequestSearchCriteriaList
	SetUseNextToken(v bool) *GetAssetsPropertyDetailRequest
	GetUseNextToken() *bool
	SetUuid(v string) *GetAssetsPropertyDetailRequest
	GetUuid() *string
}

type GetAssetsPropertyDetailRequest struct {
	// The type of Asset Fingerprints to query. Default value: **sca**. Valid values:
	//
	// - **lkm**: kernel module
	//
	// - **autorun**: startup item
	//
	// - **web_server**: web site.
	//
	// This parameter is required.
	//
	// example:
	//
	// lkm
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The page number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The aggregation item name of the Asset Fingerprints to query.
	//
	// > Call the [GetAssetsPropertyItem](~~GetAssetsPropertyItem~~) operation to obtain this parameter.
	//
	// example:
	//
	// virtio
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The language type of the request and response messages. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The token that marks the starting position for the query. Leave this parameter empty to query from the beginning.
	//
	// > Do not specify this parameter for the first call. The response includes the NextToken value for the next call. Each subsequent response contains the NextToken value for the following call.
	//
	// example:
	//
	// 71640f04f6e7b49764c8d08ae170xxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The maximum number of entries per page for a paging query. Default value: 20. If you leave this parameter empty, 20 entries are returned per page by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The server name or IP address.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The collection of search criteria for querying Asset Fingerprints details.
	SearchCriteriaList []*GetAssetsPropertyDetailRequestSearchCriteriaList `json:"SearchCriteriaList,omitempty" xml:"SearchCriteriaList,omitempty" type:"Repeated"`
	// Specifies whether to use the NextToken method to retrieve the vulnerability list. If this parameter is used, TotalCount is no longer returned. Valid values:
	//
	// - **true**: Use the NextToken method.
	//
	// - **false**: Do not use the NextToken method.
	//
	// example:
	//
	// true
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The UUID of the asset to query.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
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

func (s *GetAssetsPropertyDetailRequest) GetNextToken() *string {
	return s.NextToken
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

func (s *GetAssetsPropertyDetailRequest) GetUseNextToken() *bool {
	return s.UseNextToken
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

func (s *GetAssetsPropertyDetailRequest) SetNextToken(v string) *GetAssetsPropertyDetailRequest {
	s.NextToken = &v
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

func (s *GetAssetsPropertyDetailRequest) SetUseNextToken(v bool) *GetAssetsPropertyDetailRequest {
	s.UseNextToken = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) SetUuid(v string) *GetAssetsPropertyDetailRequest {
	s.Uuid = &v
	return s
}

func (s *GetAssetsPropertyDetailRequest) Validate() error {
	if s.SearchCriteriaList != nil {
		for _, item := range s.SearchCriteriaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAssetsPropertyDetailRequestSearchCriteriaList struct {
	// The name of the search criterion. Valid values:
	//
	// - **remarkItemName**: the aggregation item name of Asset Fingerprints. Fuzzy match is supported.
	//
	//
	// >-   - When **Biz*	- is set to **web_server**, **remarkItemName*	- indicates the domain name.
	//
	// >-   - When **Biz*	- is set to **lkm**, **remarkItemName*	- indicates the module name.
	//
	// >-   - When **Biz*	- is set to **autorun**, **remarkItemName*	- indicates the startup item path.
	//
	// example:
	//
	// remarkItemName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the search criterion.
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
