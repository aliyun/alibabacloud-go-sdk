// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatQuestionRespResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetChatQuestionRespResponseBody
	GetCost() *int64
	SetData(v *GetChatQuestionRespResponseBodyData) *GetChatQuestionRespResponseBody
	GetData() *GetChatQuestionRespResponseBodyData
	SetDataType(v string) *GetChatQuestionRespResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetChatQuestionRespResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetChatQuestionRespResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatQuestionRespResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChatQuestionRespResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetChatQuestionRespResponseBody
	GetTime() *string
}

type GetChatQuestionRespResponseBody struct {
	// example:
	//
	// null
	Cost *int64                               `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetChatQuestionRespResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetChatQuestionRespResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatQuestionRespResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetChatQuestionRespResponseBody) GetData() *GetChatQuestionRespResponseBodyData {
	return s.Data
}

func (s *GetChatQuestionRespResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetChatQuestionRespResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetChatQuestionRespResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatQuestionRespResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatQuestionRespResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChatQuestionRespResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetChatQuestionRespResponseBody) SetCost(v int64) *GetChatQuestionRespResponseBody {
	s.Cost = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetData(v *GetChatQuestionRespResponseBodyData) *GetChatQuestionRespResponseBody {
	s.Data = v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetDataType(v string) *GetChatQuestionRespResponseBody {
	s.DataType = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetErrCode(v string) *GetChatQuestionRespResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetMessage(v string) *GetChatQuestionRespResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetRequestId(v string) *GetChatQuestionRespResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetSuccess(v bool) *GetChatQuestionRespResponseBody {
	s.Success = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetTime(v string) *GetChatQuestionRespResponseBody {
	s.Time = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatQuestionRespResponseBodyData struct {
	// example:
	//
	// PROCESSING
	CurrentState *string                                            `json:"currentState,omitempty" xml:"currentState,omitempty"`
	QuestionList []*GetChatQuestionRespResponseBodyDataQuestionList `json:"questionList,omitempty" xml:"questionList,omitempty" type:"Repeated"`
}

func (s GetChatQuestionRespResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatQuestionRespResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponseBodyData) GetCurrentState() *string {
	return s.CurrentState
}

func (s *GetChatQuestionRespResponseBodyData) GetQuestionList() []*GetChatQuestionRespResponseBodyDataQuestionList {
	return s.QuestionList
}

func (s *GetChatQuestionRespResponseBodyData) SetCurrentState(v string) *GetChatQuestionRespResponseBodyData {
	s.CurrentState = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyData) SetQuestionList(v []*GetChatQuestionRespResponseBodyDataQuestionList) *GetChatQuestionRespResponseBodyData {
	s.QuestionList = v
	return s
}

func (s *GetChatQuestionRespResponseBodyData) Validate() error {
	if s.QuestionList != nil {
		for _, item := range s.QuestionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChatQuestionRespResponseBodyDataQuestionList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-11-17 10:05:00
	GmtCreate  *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	OriContent *string `json:"oriContent,omitempty" xml:"oriContent,omitempty"`
	Reply      *string `json:"reply,omitempty" xml:"reply,omitempty"`
	// example:
	//
	// 1732846760323001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// PRODUCT_QA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 39847834568436
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetChatQuestionRespResponseBodyDataQuestionList) String() string {
	return dara.Prettify(s)
}

func (s GetChatQuestionRespResponseBodyDataQuestionList) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetContent() *string {
	return s.Content
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetOriContent() *string {
	return s.OriContent
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetReply() *string {
	return s.Reply
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetType() *string {
	return s.Type
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetUserId() *string {
	return s.UserId
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) GetUserName() *string {
	return s.UserName
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetContent(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.Content = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetGmtCreate(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.GmtCreate = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetOriContent(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.OriContent = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetReply(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.Reply = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetSessionId(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.SessionId = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetType(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.Type = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetUserId(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.UserId = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetUserName(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.UserName = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) Validate() error {
	return dara.Validate(s)
}
