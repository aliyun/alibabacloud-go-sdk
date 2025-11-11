// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStyleLearningResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStyleLearningResultResponseBody
	GetCode() *string
	SetData(v *GetStyleLearningResultResponseBodyData) *GetStyleLearningResultResponseBody
	GetData() *GetStyleLearningResultResponseBodyData
	SetHttpStatusCode(v int32) *GetStyleLearningResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStyleLearningResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStyleLearningResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStyleLearningResultResponseBody
	GetSuccess() *bool
}

type GetStyleLearningResultResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetStyleLearningResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStyleLearningResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStyleLearningResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetStyleLearningResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStyleLearningResultResponseBody) GetData() *GetStyleLearningResultResponseBodyData {
	return s.Data
}

func (s *GetStyleLearningResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStyleLearningResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStyleLearningResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStyleLearningResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStyleLearningResultResponseBody) SetCode(v string) *GetStyleLearningResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetStyleLearningResultResponseBody) SetData(v *GetStyleLearningResultResponseBodyData) *GetStyleLearningResultResponseBody {
	s.Data = v
	return s
}

func (s *GetStyleLearningResultResponseBody) SetHttpStatusCode(v int32) *GetStyleLearningResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStyleLearningResultResponseBody) SetMessage(v string) *GetStyleLearningResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetStyleLearningResultResponseBody) SetRequestId(v string) *GetStyleLearningResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStyleLearningResultResponseBody) SetSuccess(v bool) *GetStyleLearningResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetStyleLearningResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStyleLearningResultResponseBodyData struct {
	// example:
	//
	// AIGC 生成的内容
	AigcResult       *string                                              `json:"AigcResult,omitempty" xml:"AigcResult,omitempty"`
	ContentList      []*GetStyleLearningResultResponseBodyDataContentList `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	CustomTextIdList []*int64                                             `json:"CustomTextIdList,omitempty" xml:"CustomTextIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 33
	Id               *int64                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	MaterialIdList   []*int64                                                  `json:"MaterialIdList,omitempty" xml:"MaterialIdList,omitempty" type:"Repeated"`
	MaterialInfoList []*GetStyleLearningResultResponseBodyDataMaterialInfoList `json:"MaterialInfoList,omitempty" xml:"MaterialInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 用户修订后内容
	RewriteResult *string `json:"RewriteResult,omitempty" xml:"RewriteResult,omitempty"`
	// example:
	//
	// 文体风格名称
	StyleName *string `json:"StyleName,omitempty" xml:"StyleName,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetStyleLearningResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStyleLearningResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStyleLearningResultResponseBodyData) GetAigcResult() *string {
	return s.AigcResult
}

func (s *GetStyleLearningResultResponseBodyData) GetContentList() []*GetStyleLearningResultResponseBodyDataContentList {
	return s.ContentList
}

func (s *GetStyleLearningResultResponseBodyData) GetCustomTextIdList() []*int64 {
	return s.CustomTextIdList
}

func (s *GetStyleLearningResultResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetStyleLearningResultResponseBodyData) GetMaterialIdList() []*int64 {
	return s.MaterialIdList
}

func (s *GetStyleLearningResultResponseBodyData) GetMaterialInfoList() []*GetStyleLearningResultResponseBodyDataMaterialInfoList {
	return s.MaterialInfoList
}

func (s *GetStyleLearningResultResponseBodyData) GetRewriteResult() *string {
	return s.RewriteResult
}

func (s *GetStyleLearningResultResponseBodyData) GetStyleName() *string {
	return s.StyleName
}

func (s *GetStyleLearningResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStyleLearningResultResponseBodyData) SetAigcResult(v string) *GetStyleLearningResultResponseBodyData {
	s.AigcResult = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetContentList(v []*GetStyleLearningResultResponseBodyDataContentList) *GetStyleLearningResultResponseBodyData {
	s.ContentList = v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetCustomTextIdList(v []*int64) *GetStyleLearningResultResponseBodyData {
	s.CustomTextIdList = v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetId(v int64) *GetStyleLearningResultResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetMaterialIdList(v []*int64) *GetStyleLearningResultResponseBodyData {
	s.MaterialIdList = v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetMaterialInfoList(v []*GetStyleLearningResultResponseBodyDataMaterialInfoList) *GetStyleLearningResultResponseBodyData {
	s.MaterialInfoList = v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetRewriteResult(v string) *GetStyleLearningResultResponseBodyData {
	s.RewriteResult = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetStyleName(v string) *GetStyleLearningResultResponseBodyData {
	s.StyleName = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) SetTaskId(v string) *GetStyleLearningResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyData) Validate() error {
	if s.ContentList != nil {
		for _, item := range s.ContentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MaterialInfoList != nil {
		for _, item := range s.MaterialInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStyleLearningResultResponseBodyDataContentList struct {
	// example:
	//
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 创建用户
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 修改用户
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s GetStyleLearningResultResponseBodyDataContentList) String() string {
	return dara.Prettify(s)
}

func (s GetStyleLearningResultResponseBodyDataContentList) GoString() string {
	return s.String()
}

func (s *GetStyleLearningResultResponseBodyDataContentList) GetContent() *string {
	return s.Content
}

func (s *GetStyleLearningResultResponseBodyDataContentList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStyleLearningResultResponseBodyDataContentList) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetStyleLearningResultResponseBodyDataContentList) GetId() *int64 {
	return s.Id
}

func (s *GetStyleLearningResultResponseBodyDataContentList) GetTitle() *string {
	return s.Title
}

func (s *GetStyleLearningResultResponseBodyDataContentList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetStyleLearningResultResponseBodyDataContentList) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *GetStyleLearningResultResponseBodyDataContentList) SetContent(v string) *GetStyleLearningResultResponseBodyDataContentList {
	s.Content = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataContentList) SetCreateTime(v string) *GetStyleLearningResultResponseBodyDataContentList {
	s.CreateTime = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataContentList) SetCreateUser(v string) *GetStyleLearningResultResponseBodyDataContentList {
	s.CreateUser = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataContentList) SetId(v int64) *GetStyleLearningResultResponseBodyDataContentList {
	s.Id = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataContentList) SetTitle(v string) *GetStyleLearningResultResponseBodyDataContentList {
	s.Title = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataContentList) SetUpdateTime(v string) *GetStyleLearningResultResponseBodyDataContentList {
	s.UpdateTime = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataContentList) SetUpdateUser(v string) *GetStyleLearningResultResponseBodyDataContentList {
	s.UpdateUser = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataContentList) Validate() error {
	return dara.Validate(s)
}

type GetStyleLearningResultResponseBodyDataMaterialInfoList struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 创建用户ID
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 创建用户姓名
	CreateUserName *string   `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DocKeywords    []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 文档类型，pdf、word、url、image
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// 外部客户上传的URL，仅用作记录保存
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// example:
	//
	// 41
	FileLength *int32 `json:"FileLength,omitempty" xml:"FileLength,omitempty"`
	// example:
	//
	// 解析后的原始html内容
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// example:
	//
	// 50
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 发布时间
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 临时的对外公开的URL
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// 文档来源，user_upload、search、viewpoint
	SrcFrom *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	// example:
	//
	// 文档摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 解析后的文本内容，对于图片来说为空
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	// example:
	//
	// 图片文档类型的Base64缩略图
	ThumbnailInBase64 *string `json:"ThumbnailInBase64,omitempty" xml:"ThumbnailInBase64,omitempty"`
	// example:
	//
	// 文档标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 修改用户ID
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// example:
	//
	// 修改用户姓名
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// example:
	//
	// 内部文档保存的URL，支持多协议，http://,file://,ftp://:客户上传时保存到内部存储的URL、长期保存、到期删除
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetStyleLearningResultResponseBodyDataMaterialInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetStyleLearningResultResponseBodyDataMaterialInfoList) GoString() string {
	return s.String()
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetAuthor() *string {
	return s.Author
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetDocKeywords() []*string {
	return s.DocKeywords
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetDocType() *string {
	return s.DocType
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetFileLength() *int32 {
	return s.FileLength
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetId() *int64 {
	return s.Id
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetPubTime() *string {
	return s.PubTime
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetSrcFrom() *string {
	return s.SrcFrom
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetSummary() *string {
	return s.Summary
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetTextContent() *string {
	return s.TextContent
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetThumbnailInBase64() *string {
	return s.ThumbnailInBase64
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetTitle() *string {
	return s.Title
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetUpdateUserName() *string {
	return s.UpdateUserName
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) GetUrl() *string {
	return s.Url
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetAuthor(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.Author = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetCreateTime(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetCreateUser(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.CreateUser = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetCreateUserName(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.CreateUserName = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetDocKeywords(v []*string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.DocKeywords = v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetDocType(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.DocType = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetExternalUrl(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.ExternalUrl = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetFileLength(v int32) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.FileLength = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetHtmlContent(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.HtmlContent = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetId(v int64) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.Id = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetPubTime(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.PubTime = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetPublicUrl(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.PublicUrl = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetShareAttr(v int32) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.ShareAttr = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetSrcFrom(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.SrcFrom = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetSummary(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.Summary = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetTextContent(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.TextContent = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetThumbnailInBase64(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.ThumbnailInBase64 = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetTitle(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.Title = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetUpdateTime(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.UpdateTime = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetUpdateUser(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.UpdateUser = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetUpdateUserName(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.UpdateUserName = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) SetUrl(v string) *GetStyleLearningResultResponseBodyDataMaterialInfoList {
	s.Url = &v
	return s
}

func (s *GetStyleLearningResultResponseBodyDataMaterialInfoList) Validate() error {
	return dara.Validate(s)
}
