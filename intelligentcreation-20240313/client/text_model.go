// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iText interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *Text
	GetAgentId() *string
	SetAgentName(v string) *Text
	GetAgentName() *string
	SetDesc(v string) *Text
	GetDesc() *string
	SetErrMsg(v string) *Text
	GetErrMsg() *string
	SetGmtCreate(v string) *Text
	GetGmtCreate() *string
	SetGmtModified(v string) *Text
	GetGmtModified() *string
	SetIllustrationTaskIdList(v []*int64) *Text
	GetIllustrationTaskIdList() []*int64
	SetPublishStatus(v string) *Text
	GetPublishStatus() *string
	SetTextContent(v string) *Text
	GetTextContent() *string
	SetTextId(v int64) *Text
	GetTextId() *int64
	SetTextIllustrationTag(v bool) *Text
	GetTextIllustrationTag() *bool
	SetTextModeType(v string) *Text
	GetTextModeType() *string
	SetTextStatus(v string) *Text
	GetTextStatus() *string
	SetTextStyleType(v string) *Text
	GetTextStyleType() *string
	SetTextTaskId(v int64) *Text
	GetTextTaskId() *int64
	SetTextThemes(v []*string) *Text
	GetTextThemes() []*string
	SetTitle(v string) *Text
	GetTitle() *string
	SetUserNameCreate(v string) *Text
	GetUserNameCreate() *string
	SetUserNameModified(v string) *Text
	GetUserNameModified() *string
}

type Text struct {
	AgentId   *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Desc                   *string  `json:"desc,omitempty" xml:"desc,omitempty"`
	ErrMsg                 *string  `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	GmtCreate              *string  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *string  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IllustrationTaskIdList []*int64 `json:"illustrationTaskIdList,omitempty" xml:"illustrationTaskIdList,omitempty" type:"Repeated"`
	PublishStatus          *string  `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	TextContent            *string  `json:"textContent,omitempty" xml:"textContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TextId *int64 `json:"textId,omitempty" xml:"textId,omitempty"`
	// example:
	//
	// true
	TextIllustrationTag *bool   `json:"textIllustrationTag,omitempty" xml:"textIllustrationTag,omitempty"`
	TextModeType        *string `json:"textModeType,omitempty" xml:"textModeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Generating
	TextStatus    *string `json:"textStatus,omitempty" xml:"textStatus,omitempty"`
	TextStyleType *string `json:"textStyleType,omitempty" xml:"textStyleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TextTaskId *int64    `json:"textTaskId,omitempty" xml:"textTaskId,omitempty"`
	TextThemes []*string `json:"textThemes,omitempty" xml:"textThemes,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserNameCreate *string `json:"userNameCreate,omitempty" xml:"userNameCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserNameModified *string `json:"userNameModified,omitempty" xml:"userNameModified,omitempty"`
}

func (s Text) String() string {
	return dara.Prettify(s)
}

func (s Text) GoString() string {
	return s.String()
}

func (s *Text) GetAgentId() *string {
	return s.AgentId
}

func (s *Text) GetAgentName() *string {
	return s.AgentName
}

func (s *Text) GetDesc() *string {
	return s.Desc
}

func (s *Text) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *Text) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Text) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Text) GetIllustrationTaskIdList() []*int64 {
	return s.IllustrationTaskIdList
}

func (s *Text) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *Text) GetTextContent() *string {
	return s.TextContent
}

func (s *Text) GetTextId() *int64 {
	return s.TextId
}

func (s *Text) GetTextIllustrationTag() *bool {
	return s.TextIllustrationTag
}

func (s *Text) GetTextModeType() *string {
	return s.TextModeType
}

func (s *Text) GetTextStatus() *string {
	return s.TextStatus
}

func (s *Text) GetTextStyleType() *string {
	return s.TextStyleType
}

func (s *Text) GetTextTaskId() *int64 {
	return s.TextTaskId
}

func (s *Text) GetTextThemes() []*string {
	return s.TextThemes
}

func (s *Text) GetTitle() *string {
	return s.Title
}

func (s *Text) GetUserNameCreate() *string {
	return s.UserNameCreate
}

func (s *Text) GetUserNameModified() *string {
	return s.UserNameModified
}

func (s *Text) SetAgentId(v string) *Text {
	s.AgentId = &v
	return s
}

func (s *Text) SetAgentName(v string) *Text {
	s.AgentName = &v
	return s
}

func (s *Text) SetDesc(v string) *Text {
	s.Desc = &v
	return s
}

func (s *Text) SetErrMsg(v string) *Text {
	s.ErrMsg = &v
	return s
}

func (s *Text) SetGmtCreate(v string) *Text {
	s.GmtCreate = &v
	return s
}

func (s *Text) SetGmtModified(v string) *Text {
	s.GmtModified = &v
	return s
}

func (s *Text) SetIllustrationTaskIdList(v []*int64) *Text {
	s.IllustrationTaskIdList = v
	return s
}

func (s *Text) SetPublishStatus(v string) *Text {
	s.PublishStatus = &v
	return s
}

func (s *Text) SetTextContent(v string) *Text {
	s.TextContent = &v
	return s
}

func (s *Text) SetTextId(v int64) *Text {
	s.TextId = &v
	return s
}

func (s *Text) SetTextIllustrationTag(v bool) *Text {
	s.TextIllustrationTag = &v
	return s
}

func (s *Text) SetTextModeType(v string) *Text {
	s.TextModeType = &v
	return s
}

func (s *Text) SetTextStatus(v string) *Text {
	s.TextStatus = &v
	return s
}

func (s *Text) SetTextStyleType(v string) *Text {
	s.TextStyleType = &v
	return s
}

func (s *Text) SetTextTaskId(v int64) *Text {
	s.TextTaskId = &v
	return s
}

func (s *Text) SetTextThemes(v []*string) *Text {
	s.TextThemes = v
	return s
}

func (s *Text) SetTitle(v string) *Text {
	s.Title = &v
	return s
}

func (s *Text) SetUserNameCreate(v string) *Text {
	s.UserNameCreate = &v
	return s
}

func (s *Text) SetUserNameModified(v string) *Text {
	s.UserNameModified = &v
	return s
}

func (s *Text) Validate() error {
	return dara.Validate(s)
}
