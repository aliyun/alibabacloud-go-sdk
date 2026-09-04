// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOutputFilesResponseBody
	GetCode() *string
	SetItems(v []*ListOutputFilesResponseBodyItems) *ListOutputFilesResponseBody
	GetItems() []*ListOutputFilesResponseBodyItems
	SetMessage(v string) *ListOutputFilesResponseBody
	GetMessage() *string
	SetPage(v int64) *ListOutputFilesResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListOutputFilesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListOutputFilesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListOutputFilesResponseBody
	GetTotal() *int64
}

type ListOutputFilesResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The output list.
	Items []*ListOutputFilesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The prompt message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of outputs that match the specified conditions.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOutputFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutputFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOutputFilesResponseBody) GetItems() []*ListOutputFilesResponseBodyItems {
	return s.Items
}

func (s *ListOutputFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOutputFilesResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListOutputFilesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOutputFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOutputFilesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListOutputFilesResponseBody) SetCode(v string) *ListOutputFilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutputFilesResponseBody) SetItems(v []*ListOutputFilesResponseBodyItems) *ListOutputFilesResponseBody {
	s.Items = v
	return s
}

func (s *ListOutputFilesResponseBody) SetMessage(v string) *ListOutputFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutputFilesResponseBody) SetPage(v int64) *ListOutputFilesResponseBody {
	s.Page = &v
	return s
}

func (s *ListOutputFilesResponseBody) SetPageSize(v int64) *ListOutputFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOutputFilesResponseBody) SetRequestId(v string) *ListOutputFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutputFilesResponseBody) SetTotal(v int64) *ListOutputFilesResponseBody {
	s.Total = &v
	return s
}

func (s *ListOutputFilesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOutputFilesResponseBodyItems struct {
	// The conversation ID.
	//
	// example:
	//
	// exampleConversationId
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The update time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The output name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The output ID.
	//
	// example:
	//
	// exampleOutputId
	OutputId *string `json:"outputId,omitempty" xml:"outputId,omitempty"`
	// The output detail list.
	OutputItems []*ListOutputFilesResponseBodyItemsOutputItems `json:"outputItems,omitempty" xml:"outputItems,omitempty" type:"Repeated"`
	// The output type: `conversation/skill/task`.
	//
	// example:
	//
	// conversation
	OutputType *string `json:"outputType,omitempty" xml:"outputType,omitempty"`
	// The internationalized display name of the output type.
	//
	// example:
	//
	// string_value
	OutputTypeDisplayName *string `json:"outputTypeDisplayName,omitempty" xml:"outputTypeDisplayName,omitempty"`
	// The skill output ID.
	//
	// example:
	//
	// exampleSkillOutputId
	SkillOutputId *string `json:"skillOutputId,omitempty" xml:"skillOutputId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// string_example_value
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListOutputFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListOutputFilesResponseBodyItems) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListOutputFilesResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListOutputFilesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListOutputFilesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListOutputFilesResponseBodyItems) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOutputFilesResponseBodyItems) GetOutputId() *string {
	return s.OutputId
}

func (s *ListOutputFilesResponseBodyItems) GetOutputItems() []*ListOutputFilesResponseBodyItemsOutputItems {
	return s.OutputItems
}

func (s *ListOutputFilesResponseBodyItems) GetOutputType() *string {
	return s.OutputType
}

func (s *ListOutputFilesResponseBodyItems) GetOutputTypeDisplayName() *string {
	return s.OutputTypeDisplayName
}

func (s *ListOutputFilesResponseBodyItems) GetSkillOutputId() *string {
	return s.SkillOutputId
}

func (s *ListOutputFilesResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListOutputFilesResponseBodyItems) SetConversationId(v string) *ListOutputFilesResponseBodyItems {
	s.ConversationId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetGmtCreate(v string) *ListOutputFilesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetGmtModified(v string) *ListOutputFilesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetName(v string) *ListOutputFilesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetOperatingObjectName(v string) *ListOutputFilesResponseBodyItems {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetOutputId(v string) *ListOutputFilesResponseBodyItems {
	s.OutputId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetOutputItems(v []*ListOutputFilesResponseBodyItemsOutputItems) *ListOutputFilesResponseBodyItems {
	s.OutputItems = v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetOutputType(v string) *ListOutputFilesResponseBodyItems {
	s.OutputType = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetOutputTypeDisplayName(v string) *ListOutputFilesResponseBodyItems {
	s.OutputTypeDisplayName = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetSkillOutputId(v string) *ListOutputFilesResponseBodyItems {
	s.SkillOutputId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) SetTaskId(v string) *ListOutputFilesResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItems) Validate() error {
	if s.OutputItems != nil {
		for _, item := range s.OutputItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOutputFilesResponseBodyItemsOutputItems struct {
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The email information. This field is present when the output type is email.
	EmailInfo *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo `json:"emailInfo,omitempty" xml:"emailInfo,omitempty" type:"Struct"`
	// The file information. This field is present when the output type is file.
	FileInfo *ListOutputFilesResponseBodyItemsOutputItemsFileInfo `json:"fileInfo,omitempty" xml:"fileInfo,omitempty" type:"Struct"`
	// The database creation time in ISO 8601 format.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The database update time in ISO 8601 format.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The output name.
	//
	// example:
	//
	// exampleItemName
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// The type of the output item. Valid values: ppt, html, document, picture, slides, video, audio, email, and others.
	//
	// example:
	//
	// ppt
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The internationalized display name of the output detail type.
	//
	// example:
	//
	// string_value
	ItemTypeDisplayName *string `json:"itemTypeDisplayName,omitempty" xml:"itemTypeDisplayName,omitempty"`
	// The message ID.
	//
	// example:
	//
	// exampleMessageId
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The output detail ID.
	//
	// example:
	//
	// string_example_value
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// Indicates whether sharing is enabled.
	//
	// example:
	//
	// true
	ShareEnabled *bool `json:"shareEnabled,omitempty" xml:"shareEnabled,omitempty"`
	// The share token that is present when sharing is enabled. You can use this token to access the public share preview API.
	//
	// example:
	//
	// example_share_token
	ShareToken *string `json:"shareToken,omitempty" xml:"shareToken,omitempty"`
	// The skill output ID.
	//
	// example:
	//
	// exampleSkillOutputId
	SkillOutputId *string `json:"skillOutputId,omitempty" xml:"skillOutputId,omitempty"`
	// The slides information. This field is present when the output type is slides.
	SlidesInfo *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo `json:"slidesInfo,omitempty" xml:"slidesInfo,omitempty" type:"Struct"`
	// The task execution ID.
	//
	// example:
	//
	// exampleTaskExecutionId
	TaskExecutionId *string `json:"taskExecutionId,omitempty" xml:"taskExecutionId,omitempty"`
}

func (s ListOutputFilesResponseBodyItemsOutputItems) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesResponseBodyItemsOutputItems) GoString() string {
	return s.String()
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetEmailInfo() *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo {
	return s.EmailInfo
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetFileInfo() *ListOutputFilesResponseBodyItemsOutputItemsFileInfo {
	return s.FileInfo
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetItemName() *string {
	return s.ItemName
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetItemTypeDisplayName() *string {
	return s.ItemTypeDisplayName
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetMessageId() *string {
	return s.MessageId
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetShareEnabled() *bool {
	return s.ShareEnabled
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetShareToken() *string {
	return s.ShareToken
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetSkillOutputId() *string {
	return s.SkillOutputId
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetSlidesInfo() *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo {
	return s.SlidesInfo
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetCreateTime(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.CreateTime = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetEmailInfo(v *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) *ListOutputFilesResponseBodyItemsOutputItems {
	s.EmailInfo = v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetFileInfo(v *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) *ListOutputFilesResponseBodyItemsOutputItems {
	s.FileInfo = v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetGmtCreate(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.GmtCreate = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetGmtModified(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.GmtModified = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetItemName(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.ItemName = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetItemType(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.ItemType = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetItemTypeDisplayName(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.ItemTypeDisplayName = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetMessageId(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.MessageId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetOutputItemId(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.OutputItemId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetShareEnabled(v bool) *ListOutputFilesResponseBodyItemsOutputItems {
	s.ShareEnabled = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetShareToken(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.ShareToken = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetSkillOutputId(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.SkillOutputId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetSlidesInfo(v *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) *ListOutputFilesResponseBodyItemsOutputItems {
	s.SlidesInfo = v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) SetTaskExecutionId(v string) *ListOutputFilesResponseBodyItemsOutputItems {
	s.TaskExecutionId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItems) Validate() error {
	if s.EmailInfo != nil {
		if err := s.EmailInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FileInfo != nil {
		if err := s.FileInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SlidesInfo != nil {
		if err := s.SlidesInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOutputFilesResponseBodyItemsOutputItemsEmailInfo struct {
	// The email body.
	//
	// example:
	//
	// string_value
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The content type, such as MARKDOWN/JSONML/HTML.
	//
	// example:
	//
	// string_value
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The recipient list.
	//
	// example:
	//
	// string_value
	Recipients []*string `json:"recipients,omitempty" xml:"recipients,omitempty" type:"Repeated"`
	// The email subject.
	//
	// example:
	//
	// string_value
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) GoString() string {
	return s.String()
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) GetBody() *string {
	return s.Body
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) GetContentType() *string {
	return s.ContentType
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) GetRecipients() []*string {
	return s.Recipients
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) GetSubject() *string {
	return s.Subject
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) SetBody(v string) *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo {
	s.Body = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) SetContentType(v string) *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo {
	s.ContentType = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) SetRecipients(v []*string) *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo {
	s.Recipients = v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) SetSubject(v string) *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo {
	s.Subject = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsEmailInfo) Validate() error {
	return dara.Validate(s)
}

type ListOutputFilesResponseBodyItemsOutputItemsFileInfo struct {
	// The file description.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The file path (OSS object key).
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The file type, such as .pdf or .md.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListOutputFilesResponseBodyItemsOutputItemsFileInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesResponseBodyItemsOutputItemsFileInfo) GoString() string {
	return s.String()
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) GetDescription() *string {
	return s.Description
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) GetName() *string {
	return s.Name
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) GetPath() *string {
	return s.Path
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) GetType() *string {
	return s.Type
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) SetDescription(v string) *ListOutputFilesResponseBodyItemsOutputItemsFileInfo {
	s.Description = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) SetName(v string) *ListOutputFilesResponseBodyItemsOutputItemsFileInfo {
	s.Name = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) SetPath(v string) *ListOutputFilesResponseBodyItemsOutputItemsFileInfo {
	s.Path = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) SetType(v string) *ListOutputFilesResponseBodyItemsOutputItemsFileInfo {
	s.Type = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsFileInfo) Validate() error {
	return dara.Validate(s)
}

type ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo struct {
	// The number of completed slides.
	//
	// example:
	//
	// 1
	CompletedSlides *int64 `json:"completedSlides,omitempty" xml:"completedSlides,omitempty"`
	// The PPT file ID.
	//
	// example:
	//
	// examplePptId
	PptId *string `json:"pptId,omitempty" xml:"pptId,omitempty"`
	// The PPT name.
	//
	// example:
	//
	// string_value
	PptName *string `json:"pptName,omitempty" xml:"pptName,omitempty"`
	// The total number of slides.
	//
	// example:
	//
	// 1
	TotalSlides *int64 `json:"totalSlides,omitempty" xml:"totalSlides,omitempty"`
}

func (s ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) GoString() string {
	return s.String()
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) GetCompletedSlides() *int64 {
	return s.CompletedSlides
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) GetPptId() *string {
	return s.PptId
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) GetPptName() *string {
	return s.PptName
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) GetTotalSlides() *int64 {
	return s.TotalSlides
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) SetCompletedSlides(v int64) *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo {
	s.CompletedSlides = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) SetPptId(v string) *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo {
	s.PptId = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) SetPptName(v string) *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo {
	s.PptName = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) SetTotalSlides(v int64) *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo {
	s.TotalSlides = &v
	return s
}

func (s *ListOutputFilesResponseBodyItemsOutputItemsSlidesInfo) Validate() error {
	return dara.Validate(s)
}
