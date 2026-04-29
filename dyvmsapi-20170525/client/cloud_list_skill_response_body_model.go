// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListSkillResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListSkillResponseBody
	GetCode() *string
	SetData(v *CloudListSkillResponseBodyData) *CloudListSkillResponseBody
	GetData() *CloudListSkillResponseBodyData
	SetMessage(v string) *CloudListSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListSkillResponseBody
	GetRequestId() *string
}

type CloudListSkillResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListSkillResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListSkillResponseBody) GetData() *CloudListSkillResponseBodyData {
	return s.Data
}

func (s *CloudListSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListSkillResponseBody) SetAccessDeniedDetail(v string) *CloudListSkillResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListSkillResponseBody) SetCode(v string) *CloudListSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListSkillResponseBody) SetData(v *CloudListSkillResponseBodyData) *CloudListSkillResponseBody {
	s.Data = v
	return s
}

func (s *CloudListSkillResponseBody) SetMessage(v string) *CloudListSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListSkillResponseBody) SetRequestId(v string) *CloudListSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListSkillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListSkillResponseBodyData struct {
	// example:
	//
	// 4
	EndRow *string `json:"EndRow,omitempty" xml:"EndRow,omitempty"`
	// example:
	//
	// 1
	FirstPage *string `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	// example:
	//
	// true
	HasNextPage *bool `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	// example:
	//
	// false
	HasPreviousPage *bool `json:"HasPreviousPage,omitempty" xml:"HasPreviousPage,omitempty"`
	// example:
	//
	// true
	IsFirstPage *bool `json:"IsFirstPage,omitempty" xml:"IsFirstPage,omitempty"`
	// example:
	//
	// true
	IsLastPage *bool `json:"IsLastPage,omitempty" xml:"IsLastPage,omitempty"`
	// example:
	//
	// 1
	LastPage *string                               `json:"LastPage,omitempty" xml:"LastPage,omitempty"`
	List     []*CloudListSkillResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	NavigatePages    *string   `json:"NavigatePages,omitempty" xml:"NavigatePages,omitempty"`
	NavigatepageNums []*string `json:"NavigatepageNums,omitempty" xml:"NavigatepageNums,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	NextPage *string `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// ""
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// 0
	PrePage *string `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// 4
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	StartRow *string `json:"StartRow,omitempty" xml:"StartRow,omitempty"`
	// example:
	//
	// 4
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CloudListSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListSkillResponseBodyData) GetEndRow() *string {
	return s.EndRow
}

func (s *CloudListSkillResponseBodyData) GetFirstPage() *string {
	return s.FirstPage
}

func (s *CloudListSkillResponseBodyData) GetHasNextPage() *bool {
	return s.HasNextPage
}

func (s *CloudListSkillResponseBodyData) GetHasPreviousPage() *bool {
	return s.HasPreviousPage
}

func (s *CloudListSkillResponseBodyData) GetIsFirstPage() *bool {
	return s.IsFirstPage
}

func (s *CloudListSkillResponseBodyData) GetIsLastPage() *bool {
	return s.IsLastPage
}

func (s *CloudListSkillResponseBodyData) GetLastPage() *string {
	return s.LastPage
}

func (s *CloudListSkillResponseBodyData) GetList() []*CloudListSkillResponseBodyDataList {
	return s.List
}

func (s *CloudListSkillResponseBodyData) GetNavigatePages() *string {
	return s.NavigatePages
}

func (s *CloudListSkillResponseBodyData) GetNavigatepageNums() []*string {
	return s.NavigatepageNums
}

func (s *CloudListSkillResponseBodyData) GetNextPage() *string {
	return s.NextPage
}

func (s *CloudListSkillResponseBodyData) GetOrderBy() *string {
	return s.OrderBy
}

func (s *CloudListSkillResponseBodyData) GetPageNum() *string {
	return s.PageNum
}

func (s *CloudListSkillResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *CloudListSkillResponseBodyData) GetPages() *string {
	return s.Pages
}

func (s *CloudListSkillResponseBodyData) GetPrePage() *string {
	return s.PrePage
}

func (s *CloudListSkillResponseBodyData) GetSize() *string {
	return s.Size
}

func (s *CloudListSkillResponseBodyData) GetStartRow() *string {
	return s.StartRow
}

func (s *CloudListSkillResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *CloudListSkillResponseBodyData) SetEndRow(v string) *CloudListSkillResponseBodyData {
	s.EndRow = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetFirstPage(v string) *CloudListSkillResponseBodyData {
	s.FirstPage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetHasNextPage(v bool) *CloudListSkillResponseBodyData {
	s.HasNextPage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetHasPreviousPage(v bool) *CloudListSkillResponseBodyData {
	s.HasPreviousPage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetIsFirstPage(v bool) *CloudListSkillResponseBodyData {
	s.IsFirstPage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetIsLastPage(v bool) *CloudListSkillResponseBodyData {
	s.IsLastPage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetLastPage(v string) *CloudListSkillResponseBodyData {
	s.LastPage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetList(v []*CloudListSkillResponseBodyDataList) *CloudListSkillResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListSkillResponseBodyData) SetNavigatePages(v string) *CloudListSkillResponseBodyData {
	s.NavigatePages = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetNavigatepageNums(v []*string) *CloudListSkillResponseBodyData {
	s.NavigatepageNums = v
	return s
}

func (s *CloudListSkillResponseBodyData) SetNextPage(v string) *CloudListSkillResponseBodyData {
	s.NextPage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetOrderBy(v string) *CloudListSkillResponseBodyData {
	s.OrderBy = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetPageNum(v string) *CloudListSkillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetPageSize(v string) *CloudListSkillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetPages(v string) *CloudListSkillResponseBodyData {
	s.Pages = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetPrePage(v string) *CloudListSkillResponseBodyData {
	s.PrePage = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetSize(v string) *CloudListSkillResponseBodyData {
	s.Size = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetStartRow(v string) *CloudListSkillResponseBodyData {
	s.StartRow = &v
	return s
}

func (s *CloudListSkillResponseBodyData) SetTotal(v string) *CloudListSkillResponseBodyData {
	s.Total = &v
	return s
}

func (s *CloudListSkillResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudListSkillResponseBodyDataList struct {
	// 描述
	//
	// example:
	//
	// 示例值示例值示例值
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 创建时间,精确到秒
	//
	// example:
	//
	// 2026-03-30 06:12:41
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7122600
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 技能id
	//
	// example:
	//
	// 52593
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 技能名称
	//
	// example:
	//
	// 示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CloudListSkillResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListSkillResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListSkillResponseBodyDataList) GetComment() *string {
	return s.Comment
}

func (s *CloudListSkillResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListSkillResponseBodyDataList) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListSkillResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *CloudListSkillResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *CloudListSkillResponseBodyDataList) SetComment(v string) *CloudListSkillResponseBodyDataList {
	s.Comment = &v
	return s
}

func (s *CloudListSkillResponseBodyDataList) SetCreateTime(v string) *CloudListSkillResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListSkillResponseBodyDataList) SetEnterpriseId(v int64) *CloudListSkillResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListSkillResponseBodyDataList) SetId(v int64) *CloudListSkillResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudListSkillResponseBodyDataList) SetName(v string) *CloudListSkillResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *CloudListSkillResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
