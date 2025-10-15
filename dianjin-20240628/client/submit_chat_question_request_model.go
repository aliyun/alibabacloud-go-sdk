// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitChatQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGmtService(v string) *SubmitChatQuestionRequest
	GetGmtService() *string
	SetLiveScriptContent(v string) *SubmitChatQuestionRequest
	GetLiveScriptContent() *string
	SetOpenSmallTalk(v bool) *SubmitChatQuestionRequest
	GetOpenSmallTalk() *bool
	SetQuestionList(v []*SubmitChatQuestionRequestQuestionList) *SubmitChatQuestionRequest
	GetQuestionList() []*SubmitChatQuestionRequestQuestionList
	SetRequestId(v string) *SubmitChatQuestionRequest
	GetRequestId() *string
	SetSessionId(v string) *SubmitChatQuestionRequest
	GetSessionId() *string
}

type SubmitChatQuestionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService *string `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
	// This parameter is required.
	LiveScriptContent *string `json:"liveScriptContent,omitempty" xml:"liveScriptContent,omitempty"`
	// example:
	//
	// true
	OpenSmallTalk *bool `json:"openSmallTalk,omitempty" xml:"openSmallTalk,omitempty"`
	// This parameter is required.
	QuestionList []*SubmitChatQuestionRequestQuestionList `json:"questionList,omitempty" xml:"questionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s SubmitChatQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitChatQuestionRequest) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionRequest) GetGmtService() *string {
	return s.GmtService
}

func (s *SubmitChatQuestionRequest) GetLiveScriptContent() *string {
	return s.LiveScriptContent
}

func (s *SubmitChatQuestionRequest) GetOpenSmallTalk() *bool {
	return s.OpenSmallTalk
}

func (s *SubmitChatQuestionRequest) GetQuestionList() []*SubmitChatQuestionRequestQuestionList {
	return s.QuestionList
}

func (s *SubmitChatQuestionRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitChatQuestionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SubmitChatQuestionRequest) SetGmtService(v string) *SubmitChatQuestionRequest {
	s.GmtService = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetLiveScriptContent(v string) *SubmitChatQuestionRequest {
	s.LiveScriptContent = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetOpenSmallTalk(v bool) *SubmitChatQuestionRequest {
	s.OpenSmallTalk = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetQuestionList(v []*SubmitChatQuestionRequestQuestionList) *SubmitChatQuestionRequest {
	s.QuestionList = v
	return s
}

func (s *SubmitChatQuestionRequest) SetRequestId(v string) *SubmitChatQuestionRequest {
	s.RequestId = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetSessionId(v string) *SubmitChatQuestionRequest {
	s.SessionId = &v
	return s
}

func (s *SubmitChatQuestionRequest) Validate() error {
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

type SubmitChatQuestionRequestQuestionList struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-17 10:05:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Reply     *string `json:"reply,omitempty" xml:"reply,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1869300950603128834
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// PRODUCT_QA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 39485783475638465
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s SubmitChatQuestionRequestQuestionList) String() string {
	return dara.Prettify(s)
}

func (s SubmitChatQuestionRequestQuestionList) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionRequestQuestionList) GetContent() *string {
	return s.Content
}

func (s *SubmitChatQuestionRequestQuestionList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *SubmitChatQuestionRequestQuestionList) GetReply() *string {
	return s.Reply
}

func (s *SubmitChatQuestionRequestQuestionList) GetSessionId() *string {
	return s.SessionId
}

func (s *SubmitChatQuestionRequestQuestionList) GetType() *string {
	return s.Type
}

func (s *SubmitChatQuestionRequestQuestionList) GetUserId() *string {
	return s.UserId
}

func (s *SubmitChatQuestionRequestQuestionList) GetUserName() *string {
	return s.UserName
}

func (s *SubmitChatQuestionRequestQuestionList) SetContent(v string) *SubmitChatQuestionRequestQuestionList {
	s.Content = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetGmtCreate(v string) *SubmitChatQuestionRequestQuestionList {
	s.GmtCreate = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetReply(v string) *SubmitChatQuestionRequestQuestionList {
	s.Reply = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetSessionId(v string) *SubmitChatQuestionRequestQuestionList {
	s.SessionId = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetType(v string) *SubmitChatQuestionRequestQuestionList {
	s.Type = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetUserId(v string) *SubmitChatQuestionRequestQuestionList {
	s.UserId = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetUserName(v string) *SubmitChatQuestionRequestQuestionList {
	s.UserName = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) Validate() error {
	return dara.Validate(s)
}
