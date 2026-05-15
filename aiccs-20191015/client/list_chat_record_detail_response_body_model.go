// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatRecordDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChatRecordDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListChatRecordDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListChatRecordDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatRecordDetailResponseBody
	GetRequestId() *string
	SetResultData(v *ListChatRecordDetailResponseBodyResultData) *ListChatRecordDetailResponseBody
	GetResultData() *ListChatRecordDetailResponseBodyResultData
	SetSuccess(v bool) *ListChatRecordDetailResponseBody
	GetSuccess() *bool
}

type ListChatRecordDetailResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultData *ListChatRecordDetailResponseBodyResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatRecordDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatRecordDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatRecordDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListChatRecordDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatRecordDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatRecordDetailResponseBody) GetResultData() *ListChatRecordDetailResponseBodyResultData {
	return s.ResultData
}

func (s *ListChatRecordDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatRecordDetailResponseBody) SetCode(v string) *ListChatRecordDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetHttpStatusCode(v int32) *ListChatRecordDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetMessage(v string) *ListChatRecordDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetRequestId(v string) *ListChatRecordDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetResultData(v *ListChatRecordDetailResponseBodyResultData) *ListChatRecordDetailResponseBody {
	s.ResultData = v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetSuccess(v bool) *ListChatRecordDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) Validate() error {
	if s.ResultData != nil {
		if err := s.ResultData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChatRecordDetailResponseBodyResultData struct {
	// example:
	//
	// 1
	CurrentPage *int64                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListChatRecordDetailResponseBodyResultDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	OnePageSize *int64 `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	// example:
	//
	// 10
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// example:
	//
	// 100
	TotalResults *int64 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s ListChatRecordDetailResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListChatRecordDetailResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBodyResultData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListChatRecordDetailResponseBodyResultData) GetData() []*ListChatRecordDetailResponseBodyResultDataData {
	return s.Data
}

func (s *ListChatRecordDetailResponseBodyResultData) GetOnePageSize() *int64 {
	return s.OnePageSize
}

func (s *ListChatRecordDetailResponseBodyResultData) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListChatRecordDetailResponseBodyResultData) GetTotalResults() *int64 {
	return s.TotalResults
}

func (s *ListChatRecordDetailResponseBodyResultData) SetCurrentPage(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.CurrentPage = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetData(v []*ListChatRecordDetailResponseBodyResultDataData) *ListChatRecordDetailResponseBodyResultData {
	s.Data = v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetOnePageSize(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.OnePageSize = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetTotalPage(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.TotalPage = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetTotalResults(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.TotalResults = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) Validate() error {
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

type ListChatRecordDetailResponseBodyResultDataData struct {
	// example:
	//
	// 1614578410000
	EndTime     *int64                                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MessageList []*ListChatRecordDetailResponseBodyResultDataDataMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
	// example:
	//
	// 123@123.com
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	// example:
	//
	// 1614578400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListChatRecordDetailResponseBodyResultDataData) String() string {
	return dara.Prettify(s)
}

func (s ListChatRecordDetailResponseBodyResultDataData) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBodyResultDataData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListChatRecordDetailResponseBodyResultDataData) GetMessageList() []*ListChatRecordDetailResponseBodyResultDataDataMessageList {
	return s.MessageList
}

func (s *ListChatRecordDetailResponseBodyResultDataData) GetServicerName() *string {
	return s.ServicerName
}

func (s *ListChatRecordDetailResponseBodyResultDataData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetEndTime(v int64) *ListChatRecordDetailResponseBodyResultDataData {
	s.EndTime = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetMessageList(v []*ListChatRecordDetailResponseBodyResultDataDataMessageList) *ListChatRecordDetailResponseBodyResultDataData {
	s.MessageList = v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetServicerName(v string) *ListChatRecordDetailResponseBodyResultDataData {
	s.ServicerName = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetStartTime(v int64) *ListChatRecordDetailResponseBodyResultDataData {
	s.StartTime = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataData) Validate() error {
	if s.MessageList != nil {
		for _, item := range s.MessageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatRecordDetailResponseBodyResultDataDataMessageList struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1614578400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// MSG
	MsgType *string `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	// example:
	//
	// account1
	SenderName *string `json:"SenderName,omitempty" xml:"SenderName,omitempty"`
	// example:
	//
	// 2
	SenderType *int64 `json:"SenderType,omitempty" xml:"SenderType,omitempty"`
}

func (s ListChatRecordDetailResponseBodyResultDataDataMessageList) String() string {
	return dara.Prettify(s)
}

func (s ListChatRecordDetailResponseBodyResultDataDataMessageList) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) GetContent() *string {
	return s.Content
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) GetMsgType() *string {
	return s.MsgType
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) GetSenderName() *string {
	return s.SenderName
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) GetSenderType() *int64 {
	return s.SenderType
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetContent(v string) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.Content = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetCreateTime(v int64) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.CreateTime = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetMsgType(v string) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.MsgType = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetSenderName(v string) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.SenderName = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetSenderType(v int64) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.SenderType = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) Validate() error {
	return dara.Validate(s)
}
