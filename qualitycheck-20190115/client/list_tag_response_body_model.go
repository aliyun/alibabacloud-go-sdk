// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTagResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *ListTagResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListTagResponseBodyData) *ListTagResponseBody
	GetData() []*ListTagResponseBodyData
	SetDataSize(v int32) *ListTagResponseBody
	GetDataSize() *int32
	SetMessage(v string) *ListTagResponseBody
	GetMessage() *string
	SetPageSize(v int32) *ListTagResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTagResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTagResponseBody
	GetTotalCount() *int32
}

type ListTagResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page size.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned data.
	Data []*ListTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The actual number of records returned on the current page.
	//
	// example:
	//
	// 2
	DataSize *int32 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The error message, if any.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records that meet the conditions.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTagResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTagResponseBody) GetData() []*ListTagResponseBodyData {
	return s.Data
}

func (s *ListTagResponseBody) GetDataSize() *int32 {
	return s.DataSize
}

func (s *ListTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTagResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTagResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagResponseBody) SetCode(v string) *ListTagResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResponseBody) SetCurrentPage(v int32) *ListTagResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagResponseBody) SetData(v []*ListTagResponseBodyData) *ListTagResponseBody {
	s.Data = v
	return s
}

func (s *ListTagResponseBody) SetDataSize(v int32) *ListTagResponseBody {
	s.DataSize = &v
	return s
}

func (s *ListTagResponseBody) SetMessage(v string) *ListTagResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagResponseBody) SetPageSize(v int32) *ListTagResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagResponseBody) SetRequestId(v string) *ListTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResponseBody) SetSuccess(v bool) *ListTagResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResponseBody) SetTotalCount(v int32) *ListTagResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResponseBodyData struct {
	// The number of direct child nodes.
	//
	// example:
	//
	// 5
	ChildCount *int32 `json:"ChildCount,omitempty" xml:"ChildCount,omitempty"`
	// The time when the label was created.
	//
	// example:
	//
	// 1748428991000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The label description.
	//
	// example:
	//
	// 用于归集售后服务相关的所有意图与 FAQ
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The level of the current node.
	//
	// example:
	//
	// 2
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The time when the label was last modified.
	//
	// example:
	//
	// 1748428991000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The label name.
	//
	// example:
	//
	// 售后问题
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent label node.
	//
	// example:
	//
	// -1
	ParentTagId *int64 `json:"ParentTagId,omitempty" xml:"ParentTagId,omitempty"`
	// The node path.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// The label ID.
	//
	// example:
	//
	// 128
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagResponseBodyData) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *ListTagResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTagResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListTagResponseBodyData) GetLevel() *int32 {
	return s.Level
}

func (s *ListTagResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListTagResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListTagResponseBodyData) GetParentTagId() *int64 {
	return s.ParentTagId
}

func (s *ListTagResponseBodyData) GetPath() []*string {
	return s.Path
}

func (s *ListTagResponseBodyData) GetTagId() *int64 {
	return s.TagId
}

func (s *ListTagResponseBodyData) SetChildCount(v int32) *ListTagResponseBodyData {
	s.ChildCount = &v
	return s
}

func (s *ListTagResponseBodyData) SetCreateTime(v int64) *ListTagResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListTagResponseBodyData) SetDescription(v string) *ListTagResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListTagResponseBodyData) SetLevel(v int32) *ListTagResponseBodyData {
	s.Level = &v
	return s
}

func (s *ListTagResponseBodyData) SetModifyTime(v int64) *ListTagResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *ListTagResponseBodyData) SetName(v string) *ListTagResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTagResponseBodyData) SetParentTagId(v int64) *ListTagResponseBodyData {
	s.ParentTagId = &v
	return s
}

func (s *ListTagResponseBodyData) SetPath(v []*string) *ListTagResponseBodyData {
	s.Path = v
	return s
}

func (s *ListTagResponseBodyData) SetTagId(v int64) *ListTagResponseBodyData {
	s.TagId = &v
	return s
}

func (s *ListTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
