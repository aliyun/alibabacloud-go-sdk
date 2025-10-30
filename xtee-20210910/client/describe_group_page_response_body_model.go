// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGroupPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeGroupPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeGroupPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeGroupPageResponseBodyResultObject) *DescribeGroupPageResponseBody
	GetResultObject() []*DescribeGroupPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeGroupPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeGroupPageResponseBody
	GetTotalPage() *int32
}

type DescribeGroupPageResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeGroupPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeGroupPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupPageResponseBody) GetResultObject() []*DescribeGroupPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeGroupPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeGroupPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeGroupPageResponseBody) SetRequestId(v string) *DescribeGroupPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupPageResponseBody) SetCurrentPage(v int32) *DescribeGroupPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupPageResponseBody) SetPageSize(v int32) *DescribeGroupPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupPageResponseBody) SetResultObject(v []*DescribeGroupPageResponseBodyResultObject) *DescribeGroupPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeGroupPageResponseBody) SetTotalItem(v int32) *DescribeGroupPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeGroupPageResponseBody) SetTotalPage(v int32) *DescribeGroupPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeGroupPageResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupPageResponseBodyResultObject struct {
	// Community number.
	//
	// example:
	//
	// 129838420210118141502RaMMIgVG
	CommunityNo *string `json:"communityNo,omitempty" xml:"communityNo,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1699450018265
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Group risk concentration.
	//
	// example:
	//
	// 80
	GroupRisk *string `json:"groupRisk,omitempty" xml:"groupRisk,omitempty"`
	// Group scale.
	//
	// example:
	//
	// 2
	GroupScale *string `json:"groupScale,omitempty" xml:"groupScale,omitempty"`
	// Primary key ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Scene name.
	//
	// example:
	//
	// 通用图计算
	SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 6770764
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// User UID.
	//
	// example:
	//
	// 1519714049632764
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DescribeGroupPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeGroupPageResponseBodyResultObject) GetCommunityNo() *string {
	return s.CommunityNo
}

func (s *DescribeGroupPageResponseBodyResultObject) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGroupPageResponseBodyResultObject) GetGroupRisk() *string {
	return s.GroupRisk
}

func (s *DescribeGroupPageResponseBodyResultObject) GetGroupScale() *string {
	return s.GroupScale
}

func (s *DescribeGroupPageResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeGroupPageResponseBodyResultObject) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeGroupPageResponseBodyResultObject) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeGroupPageResponseBodyResultObject) GetUserId() *string {
	return s.UserId
}

func (s *DescribeGroupPageResponseBodyResultObject) SetCommunityNo(v string) *DescribeGroupPageResponseBodyResultObject {
	s.CommunityNo = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) SetCreateTime(v int64) *DescribeGroupPageResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) SetGroupRisk(v string) *DescribeGroupPageResponseBodyResultObject {
	s.GroupRisk = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) SetGroupScale(v string) *DescribeGroupPageResponseBodyResultObject {
	s.GroupScale = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) SetId(v int64) *DescribeGroupPageResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) SetSceneName(v string) *DescribeGroupPageResponseBodyResultObject {
	s.SceneName = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) SetTaskId(v int64) *DescribeGroupPageResponseBodyResultObject {
	s.TaskId = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) SetUserId(v string) *DescribeGroupPageResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeGroupPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
