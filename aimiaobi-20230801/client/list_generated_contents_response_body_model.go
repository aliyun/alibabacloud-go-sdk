// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeneratedContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGeneratedContentsResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListGeneratedContentsResponseBody
	GetCurrent() *int32
	SetData(v []*ListGeneratedContentsResponseBodyData) *ListGeneratedContentsResponseBody
	GetData() []*ListGeneratedContentsResponseBodyData
	SetHttpStatusCode(v int32) *ListGeneratedContentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGeneratedContentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGeneratedContentsResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListGeneratedContentsResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListGeneratedContentsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListGeneratedContentsResponseBody
	GetTotal() *int32
}

type ListGeneratedContentsResponseBody struct {
	// Status code
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// Document list
	Data []*ListGeneratedContentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// Number of items per page
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Operation status: true means success. false means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGeneratedContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGeneratedContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGeneratedContentsResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListGeneratedContentsResponseBody) GetData() []*ListGeneratedContentsResponseBodyData {
	return s.Data
}

func (s *ListGeneratedContentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGeneratedContentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGeneratedContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGeneratedContentsResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListGeneratedContentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGeneratedContentsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListGeneratedContentsResponseBody) SetCode(v string) *ListGeneratedContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetCurrent(v int32) *ListGeneratedContentsResponseBody {
	s.Current = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetData(v []*ListGeneratedContentsResponseBodyData) *ListGeneratedContentsResponseBody {
	s.Data = v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetHttpStatusCode(v int32) *ListGeneratedContentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetMessage(v string) *ListGeneratedContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetRequestId(v string) *ListGeneratedContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetSize(v int32) *ListGeneratedContentsResponseBody {
	s.Size = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetSuccess(v bool) *ListGeneratedContentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetTotal(v int32) *ListGeneratedContentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) Validate() error {
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

type ListGeneratedContentsResponseBodyData struct {
	// Content: rich text
	//
	// example:
	//
	// 杭州亚运会
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Content domain (content category)
	//
	// - media: Media writing
	//
	// - government: Government document writing
	//
	// - office: Office writing
	//
	// - market: Marketing writing
	//
	// - custom: Custom writing
	//
	// - commentGenerate: Opinion generation
	//
	// example:
	//
	// media
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	// Content: plain text
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
	// "123"
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Device ID
	//
	// example:
	//
	// xxx
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// File attributes
	FileAttr *ListGeneratedContentsResponseBodyDataFileAttr `json:"FileAttr,omitempty" xml:"FileAttr,omitempty" type:"Struct"`
	// File ID
	//
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// Document ID
	//
	// example:
	//
	// 10
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Keywords
	KeywordList []*string `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	// Keywords (string)
	//
	// example:
	//
	// 观点
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// Last prompt used to generate this document
	//
	// example:
	//
	// 创作xxx文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Conversation task ID
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
	// "1111"
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// UUID traceability ID
	//
	// example:
	//
	// xxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListGeneratedContentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGeneratedContentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListGeneratedContentsResponseBodyData) GetContentDomain() *string {
	return s.ContentDomain
}

func (s *ListGeneratedContentsResponseBodyData) GetContentText() *string {
	return s.ContentText
}

func (s *ListGeneratedContentsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListGeneratedContentsResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListGeneratedContentsResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListGeneratedContentsResponseBodyData) GetFileAttr() *ListGeneratedContentsResponseBodyDataFileAttr {
	return s.FileAttr
}

func (s *ListGeneratedContentsResponseBodyData) GetFileKey() *string {
	return s.FileKey
}

func (s *ListGeneratedContentsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListGeneratedContentsResponseBodyData) GetKeywordList() []*string {
	return s.KeywordList
}

func (s *ListGeneratedContentsResponseBodyData) GetKeywords() *string {
	return s.Keywords
}

func (s *ListGeneratedContentsResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *ListGeneratedContentsResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListGeneratedContentsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListGeneratedContentsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListGeneratedContentsResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *ListGeneratedContentsResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListGeneratedContentsResponseBodyData) SetContent(v string) *ListGeneratedContentsResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetContentDomain(v string) *ListGeneratedContentsResponseBodyData {
	s.ContentDomain = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetContentText(v string) *ListGeneratedContentsResponseBodyData {
	s.ContentText = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetCreateTime(v string) *ListGeneratedContentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetCreateUser(v string) *ListGeneratedContentsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetDeviceId(v string) *ListGeneratedContentsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetFileAttr(v *ListGeneratedContentsResponseBodyDataFileAttr) *ListGeneratedContentsResponseBodyData {
	s.FileAttr = v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetFileKey(v string) *ListGeneratedContentsResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetId(v int64) *ListGeneratedContentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetKeywordList(v []*string) *ListGeneratedContentsResponseBodyData {
	s.KeywordList = v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetKeywords(v string) *ListGeneratedContentsResponseBodyData {
	s.Keywords = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetPrompt(v string) *ListGeneratedContentsResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetTaskId(v string) *ListGeneratedContentsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetTitle(v string) *ListGeneratedContentsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUpdateTime(v string) *ListGeneratedContentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUpdateUser(v string) *ListGeneratedContentsResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUuid(v string) *ListGeneratedContentsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) Validate() error {
	if s.FileAttr != nil {
		if err := s.FileAttr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGeneratedContentsResponseBodyDataFileAttr struct {
	// File name
	//
	// example:
	//
	// homedepothp.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Video height
	//
	// example:
	//
	// 500
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Temporary URL for video file access. Expires in one hour.
	//
	// example:
	//
	// http://www.example.com/xxx.mp4
	TmpUrl *string `json:"TmpUrl,omitempty" xml:"TmpUrl,omitempty"`
	// Video width
	//
	// example:
	//
	// 800
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListGeneratedContentsResponseBodyDataFileAttr) String() string {
	return dara.Prettify(s)
}

func (s ListGeneratedContentsResponseBodyDataFileAttr) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) GetFileName() *string {
	return s.FileName
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) GetHeight() *int32 {
	return s.Height
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) GetTmpUrl() *string {
	return s.TmpUrl
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) GetWidth() *int32 {
	return s.Width
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) SetFileName(v string) *ListGeneratedContentsResponseBodyDataFileAttr {
	s.FileName = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) SetHeight(v int32) *ListGeneratedContentsResponseBodyDataFileAttr {
	s.Height = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) SetTmpUrl(v string) *ListGeneratedContentsResponseBodyDataFileAttr {
	s.TmpUrl = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) SetWidth(v int32) *ListGeneratedContentsResponseBodyDataFileAttr {
	s.Width = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyDataFileAttr) Validate() error {
	return dara.Validate(s)
}
