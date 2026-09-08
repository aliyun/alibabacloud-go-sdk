// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageListResponseBody
	GetCode() *string
	SetData(v *ReadMessageListResponseBodyData) *ReadMessageListResponseBody
	GetData() *ReadMessageListResponseBodyData
	SetMessage(v string) *ReadMessageListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMessageListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMessageListResponseBody
	GetSuccess() *bool
}

type ReadMessageListResponseBody struct {
	// The error code returned when the call fails. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	Data *ReadMessageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageListResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageListResponseBody) GetData() *ReadMessageListResponseBodyData {
	return s.Data
}

func (s *ReadMessageListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMessageListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMessageListResponseBody) SetCode(v string) *ReadMessageListResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageListResponseBody) SetData(v *ReadMessageListResponseBodyData) *ReadMessageListResponseBody {
	s.Data = v
	return s
}

func (s *ReadMessageListResponseBody) SetMessage(v string) *ReadMessageListResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageListResponseBody) SetRequestId(v string) *ReadMessageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageListResponseBody) SetSuccess(v bool) *ReadMessageListResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMessageListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadMessageListResponseBodyData struct {
	// The number of messages.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// /
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// /
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 24
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned data.
	Rows []*ReadMessageListResponseBodyDataRows `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ReadMessageListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *ReadMessageListResponseBodyData) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ReadMessageListResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadMessageListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ReadMessageListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ReadMessageListResponseBodyData) GetRows() []*ReadMessageListResponseBodyDataRows {
	return s.Rows
}

func (s *ReadMessageListResponseBodyData) SetCount(v int64) *ReadMessageListResponseBodyData {
	s.Count = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetMaxResults(v int64) *ReadMessageListResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetNextToken(v string) *ReadMessageListResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetPage(v int32) *ReadMessageListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetPageSize(v int32) *ReadMessageListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetRows(v []*ReadMessageListResponseBodyDataRows) *ReadMessageListResponseBodyData {
	s.Rows = v
	return s
}

func (s *ReadMessageListResponseBodyData) Validate() error {
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadMessageListResponseBodyDataRows struct {
	// The category code.
	//
	// example:
	//
	// test
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The message category name.
	//
	// example:
	//
	// 活动消息
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// /
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// The message class ID.
	//
	// example:
	//
	// 1
	ClassId *int64 `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The message content.
	//
	// example:
	//
	// "消息内容示例“
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The deletion flag.
	//
	// example:
	//
	// 0
	Deleted *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The time when the message was created.
	//
	// example:
	//
	// 1723772244000
	GmtCreated *int64 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the message was updated.
	//
	// example:
	//
	// 1723772244000
	GmtUpdate *int64 `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// /
	MassId *int64 `json:"MassId,omitempty" xml:"MassId,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// /
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 3727683838
	MsgId *int64 `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The read status. A value of 0 indicates unread. A value of 1 indicates read.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message title.
	//
	// example:
	//
	// "标题示例“
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The highlighted title.
	//
	// example:
	//
	// title
	Titleh *string `json:"Titleh,omitempty" xml:"Titleh,omitempty"`
}

func (s ReadMessageListResponseBodyDataRows) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageListResponseBodyDataRows) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponseBodyDataRows) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *ReadMessageListResponseBodyDataRows) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ReadMessageListResponseBodyDataRows) GetClass() *string {
	return s.Class
}

func (s *ReadMessageListResponseBodyDataRows) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadMessageListResponseBodyDataRows) GetContent() *string {
	return s.Content
}

func (s *ReadMessageListResponseBodyDataRows) GetDeleted() *int32 {
	return s.Deleted
}

func (s *ReadMessageListResponseBodyDataRows) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *ReadMessageListResponseBodyDataRows) GetGmtUpdate() *int64 {
	return s.GmtUpdate
}

func (s *ReadMessageListResponseBodyDataRows) GetMassId() *int64 {
	return s.MassId
}

func (s *ReadMessageListResponseBodyDataRows) GetMemo() *string {
	return s.Memo
}

func (s *ReadMessageListResponseBodyDataRows) GetMsgId() *int64 {
	return s.MsgId
}

func (s *ReadMessageListResponseBodyDataRows) GetStatus() *int64 {
	return s.Status
}

func (s *ReadMessageListResponseBodyDataRows) GetTitle() *string {
	return s.Title
}

func (s *ReadMessageListResponseBodyDataRows) GetTitleh() *string {
	return s.Titleh
}

func (s *ReadMessageListResponseBodyDataRows) SetCategoryCode(v string) *ReadMessageListResponseBodyDataRows {
	s.CategoryCode = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetCategoryName(v string) *ReadMessageListResponseBodyDataRows {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetClass(v string) *ReadMessageListResponseBodyDataRows {
	s.Class = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetClassId(v int64) *ReadMessageListResponseBodyDataRows {
	s.ClassId = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetContent(v string) *ReadMessageListResponseBodyDataRows {
	s.Content = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetDeleted(v int32) *ReadMessageListResponseBodyDataRows {
	s.Deleted = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetGmtCreated(v int64) *ReadMessageListResponseBodyDataRows {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetGmtUpdate(v int64) *ReadMessageListResponseBodyDataRows {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetMassId(v int64) *ReadMessageListResponseBodyDataRows {
	s.MassId = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetMemo(v string) *ReadMessageListResponseBodyDataRows {
	s.Memo = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetMsgId(v int64) *ReadMessageListResponseBodyDataRows {
	s.MsgId = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetStatus(v int64) *ReadMessageListResponseBodyDataRows {
	s.Status = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetTitle(v string) *ReadMessageListResponseBodyDataRows {
	s.Title = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetTitleh(v string) *ReadMessageListResponseBodyDataRows {
	s.Titleh = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) Validate() error {
	return dara.Validate(s)
}
