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
	// The type of asset fingerprint to be queried, with a default value of **sca**. Values:
	//
	// - **lkm**: Kernel module
	//
	// - **autorun**: Startup item
	//
	// - **web_server**: Web site
	//
	// This parameter is required.
	//
	// example:
	//
	// lkm
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// Set the page number from which to start displaying the query results. The default value is **1**, indicating that the display starts from the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The aggregated item name of the asset fingerprint to be queried.
	//
	// > Call the [GetAssetsPropertyItem](~~GetAssetsPropertyItem~~) API to obtain this parameter.
	//
	// example:
	//
	// virtio
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The language type for the request and response. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Used to mark the current read position. Leave it empty to start from the beginning.
	//
	// > Do not fill in for the first call; the response will include the NextToken for the second call. Each subsequent call\\"s response will contain the NextToken for the next call.
	//
	// example:
	//
	// 71640f04f6e7b49764c8d08ae170xxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specify the maximum number of data entries per page in a paginated query. The default number of data entries per page is 20. If the PageSize parameter is empty, 20 data entries will be returned by default.
	//
	// > It is recommended that the PageSize value is not empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Server name or IP.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// A set of conditions for querying asset fingerprint details.
	SearchCriteriaList []*GetAssetsPropertyDetailRequestSearchCriteriaList `json:"SearchCriteriaList,omitempty" xml:"SearchCriteriaList,omitempty" type:"Repeated"`
	// Whether to use the NextToken method to fetch the list of vulnerabilities. If this parameter is used, TotalCount will not be returned. Values:
	//
	// - **true**: Use the NextToken method.
	//
	// - **false**: Do not use the NextToken method.
	//
	// example:
	//
	// true
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The UUID of the asset to be queried.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) API to obtain this parameter.
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
	// The name of the condition to be queried. Values are as follows:
	//
	// - **remarkItemName**: The aggregated item name of the asset fingerprint, supporting fuzzy matching
	//
	//
	// >-   - When **Biz*	- is **web_server**, **remarkItemName*	- represents the domain name as the search condition.
	//
	// >-   - When **Biz*	- is **lkm**, **remarkItemName*	- represents the module name as the search condition.
	//
	// >-   - When **Biz*	- is **autorun**, **remarkItemName*	- represents the startup item path as the search condition.
	//
	// example:
	//
	// remarkItemName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the condition to be queried.
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
