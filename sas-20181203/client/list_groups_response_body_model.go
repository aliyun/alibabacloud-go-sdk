// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGroupsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListGroupsResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListGroupsResponseBodyList) *ListGroupsResponseBody
	GetList() []*ListGroupsResponseBodyList
	SetMessage(v string) *ListGroupsResponseBody
	GetMessage() *string
	SetPageInfo(v *ListGroupsResponseBodyPageInfo) *ListGroupsResponseBody
	GetPageInfo() *ListGroupsResponseBodyPageInfo
	SetRequestId(v string) *ListGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGroupsResponseBody
	GetSuccess() *bool
}

type ListGroupsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The groups.
	List []*ListGroupsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListGroupsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CB414DB5-F692-5DAB-9F0F-975C060AF***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether exceptions are handled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGroupsResponseBody) GetList() []*ListGroupsResponseBodyList {
	return s.List
}

func (s *ListGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGroupsResponseBody) GetPageInfo() *ListGroupsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGroupsResponseBody) SetCode(v string) *ListGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupsResponseBody) SetHttpStatusCode(v int32) *ListGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGroupsResponseBody) SetList(v []*ListGroupsResponseBodyList) *ListGroupsResponseBody {
	s.List = v
	return s
}

func (s *ListGroupsResponseBody) SetMessage(v string) *ListGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupsResponseBody) SetPageInfo(v *ListGroupsResponseBodyPageInfo) *ListGroupsResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetSuccess(v bool) *ListGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGroupsResponseBodyList struct {
	// The server group type. Valid values:
	//
	// 	- **0**: default group
	//
	// 	- **1**: other group
	//
	// example:
	//
	// 1
	GroupFlag *int32 `json:"GroupFlag,omitempty" xml:"GroupFlag,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// 11028542
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The server group name.
	//
	// example:
	//
	// cn-shenzhen+dir-1440978***
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListGroupsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyList) GetGroupFlag() *int32 {
	return s.GroupFlag
}

func (s *ListGroupsResponseBodyList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListGroupsResponseBodyList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsResponseBodyList) SetGroupFlag(v int32) *ListGroupsResponseBodyList {
	s.GroupFlag = &v
	return s
}

func (s *ListGroupsResponseBodyList) SetGroupId(v int64) *ListGroupsResponseBodyList {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyList) SetGroupName(v string) *ListGroupsResponseBodyList {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListGroupsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// B604532DEF982B875E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 202
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListGroupsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListGroupsResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGroupsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListGroupsResponseBodyPageInfo) SetCount(v int32) *ListGroupsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListGroupsResponseBodyPageInfo) SetCurrentPage(v int32) *ListGroupsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListGroupsResponseBodyPageInfo) SetNextToken(v string) *ListGroupsResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *ListGroupsResponseBodyPageInfo) SetPageSize(v int32) *ListGroupsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListGroupsResponseBodyPageInfo) SetTotalCount(v int32) *ListGroupsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
