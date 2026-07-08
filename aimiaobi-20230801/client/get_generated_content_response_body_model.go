// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGeneratedContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGeneratedContentResponseBody
	GetCode() *string
	SetData(v *GetGeneratedContentResponseBodyData) *GetGeneratedContentResponseBody
	GetData() *GetGeneratedContentResponseBodyData
	SetHttpStatusCode(v int32) *GetGeneratedContentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGeneratedContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGeneratedContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGeneratedContentResponseBody
	GetSuccess() *bool
}

type GetGeneratedContentResponseBody struct {
	// Status code
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business data
	Data *GetGeneratedContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. true indicates success. false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGeneratedContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGeneratedContentResponseBody) GetData() *GetGeneratedContentResponseBodyData {
	return s.Data
}

func (s *GetGeneratedContentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGeneratedContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGeneratedContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGeneratedContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGeneratedContentResponseBody) SetCode(v string) *GetGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetData(v *GetGeneratedContentResponseBodyData) *GetGeneratedContentResponseBody {
	s.Data = v
	return s
}

func (s *GetGeneratedContentResponseBody) SetHttpStatusCode(v int32) *GetGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetMessage(v string) *GetGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetRequestId(v string) *GetGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetSuccess(v bool) *GetGeneratedContentResponseBody {
	s.Success = &v
	return s
}

func (s *GetGeneratedContentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGeneratedContentResponseBodyData struct {
	// Rich text content
	//
	// example:
	//
	// 杭州亚运会
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Content generation domain. Valid values: media (news) and government (government services)
	//
	// example:
	//
	// media
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	// Plain text content
	//
	// example:
	//
	// 杭州亚运会
	ContentText *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2024-01-04 11:46:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creator
	//
	// example:
	//
	// "1"
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Device ID
	//
	// example:
	//
	// xxx
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Document ID
	//
	// example:
	//
	// 86
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// List of words to ignore during content audit (JSON string)
	//
	// example:
	//
	// "[{}]"
	IgnoreContentAuditWords *string `json:"IgnoreContentAuditWords,omitempty" xml:"IgnoreContentAuditWords,omitempty"`
	// Keywords
	KeywordList []*string `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	// Keywords (string format)
	//
	// example:
	//
	// [\\"教师\\",\\"乡村\\"]
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// Last prompt used to generate this content
	//
	// example:
	//
	// 创作xxx文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Session task ID
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Title
	//
	// example:
	//
	// 杭州亚运会
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Update time
	//
	// example:
	//
	// 2024-01-04 11:46:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Updater
	//
	// example:
	//
	// "1"
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// UUID for traceability
	//
	// example:
	//
	// 0961a514-2e26-4aa6-b22b-f592d145fe47
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetGeneratedContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGeneratedContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetGeneratedContentResponseBodyData) GetContentDomain() *string {
	return s.ContentDomain
}

func (s *GetGeneratedContentResponseBodyData) GetContentText() *string {
	return s.ContentText
}

func (s *GetGeneratedContentResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetGeneratedContentResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetGeneratedContentResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetGeneratedContentResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGeneratedContentResponseBodyData) GetIgnoreContentAuditWords() *string {
	return s.IgnoreContentAuditWords
}

func (s *GetGeneratedContentResponseBodyData) GetKeywordList() []*string {
	return s.KeywordList
}

func (s *GetGeneratedContentResponseBodyData) GetKeywords() *string {
	return s.Keywords
}

func (s *GetGeneratedContentResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *GetGeneratedContentResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetGeneratedContentResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetGeneratedContentResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetGeneratedContentResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *GetGeneratedContentResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetGeneratedContentResponseBodyData) SetContent(v string) *GetGeneratedContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetContentDomain(v string) *GetGeneratedContentResponseBodyData {
	s.ContentDomain = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetContentText(v string) *GetGeneratedContentResponseBodyData {
	s.ContentText = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetCreateTime(v string) *GetGeneratedContentResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetCreateUser(v string) *GetGeneratedContentResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetDeviceId(v string) *GetGeneratedContentResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetId(v int64) *GetGeneratedContentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetIgnoreContentAuditWords(v string) *GetGeneratedContentResponseBodyData {
	s.IgnoreContentAuditWords = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetKeywordList(v []*string) *GetGeneratedContentResponseBodyData {
	s.KeywordList = v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetKeywords(v string) *GetGeneratedContentResponseBodyData {
	s.Keywords = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetPrompt(v string) *GetGeneratedContentResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetTaskId(v string) *GetGeneratedContentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetTitle(v string) *GetGeneratedContentResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUpdateTime(v string) *GetGeneratedContentResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUpdateUser(v string) *GetGeneratedContentResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUuid(v string) *GetGeneratedContentResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
