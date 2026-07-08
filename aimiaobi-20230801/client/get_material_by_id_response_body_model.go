// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaterialByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMaterialByIdResponseBody
	GetCode() *string
	SetData(v *GetMaterialByIdResponseBodyData) *GetMaterialByIdResponseBody
	GetData() *GetMaterialByIdResponseBodyData
	SetHttpStatusCode(v int32) *GetMaterialByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMaterialByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMaterialByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMaterialByIdResponseBody
	GetSuccess() *bool
}

type GetMaterialByIdResponseBody struct {
	// Status code
	//
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business data
	Data *GetMaterialByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error description
	//
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Unique request identifier
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates success: true for success, false for failure
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMaterialByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMaterialByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMaterialByIdResponseBody) GetData() *GetMaterialByIdResponseBodyData {
	return s.Data
}

func (s *GetMaterialByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMaterialByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMaterialByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMaterialByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMaterialByIdResponseBody) SetCode(v string) *GetMaterialByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetData(v *GetMaterialByIdResponseBodyData) *GetMaterialByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetMaterialByIdResponseBody) SetHttpStatusCode(v int32) *GetMaterialByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetMessage(v string) *GetMaterialByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetRequestId(v string) *GetMaterialByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetSuccess(v bool) *GetMaterialByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetMaterialByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMaterialByIdResponseBodyData struct {
	// Author
	//
	// example:
	//
	// 文档作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2023-03-21 11:34:19
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creator user ID
	//
	// example:
	//
	// 1
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Document tags used for classification and other purposes. Separate multiple keywords with commas.
	DocKeywords []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	// Document type, such as pdf, word, url, or image
	//
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// URL uploaded by an external customer. Used only for record keeping.
	//
	// example:
	//
	// https://www.example.com
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// Web page content
	//
	// example:
	//
	// 网页内容
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// Primary key
	//
	// example:
	//
	// 32
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Publication time
	//
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// Temporary public URL
	//
	// example:
	//
	// https://www.example.com
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// Sharing attribute stored as bit flags. The first bit indicates sharing within the workspace, the second bit indicates sharing within the tenant, and the third bit indicates system-wide sharing.
	//
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// Document source, such as user_upload, search, or viewpoint
	//
	// example:
	//
	// user_upload
	SrcFrom *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	// Document summary
	//
	// example:
	//
	// 文档摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Parsed text content. Empty for images.
	//
	// example:
	//
	// 文本内容
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	// Base64-encoded thumbnail for image documents
	//
	// example:
	//
	// Base64编码的缩略图
	ThumbnailInBase64 *string `json:"ThumbnailInBase64,omitempty" xml:"ThumbnailInBase64,omitempty"`
	// Document title
	//
	// example:
	//
	// 文档标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Modification time
	//
	// example:
	//
	// 2022-04-08 19:33:01
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Modifier user ID
	//
	// example:
	//
	// 1
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// Internal document storage URL
	//
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMaterialByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMaterialByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponseBodyData) GetAuthor() *string {
	return s.Author
}

func (s *GetMaterialByIdResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMaterialByIdResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetMaterialByIdResponseBodyData) GetDocKeywords() []*string {
	return s.DocKeywords
}

func (s *GetMaterialByIdResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *GetMaterialByIdResponseBodyData) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *GetMaterialByIdResponseBodyData) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *GetMaterialByIdResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMaterialByIdResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *GetMaterialByIdResponseBodyData) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *GetMaterialByIdResponseBodyData) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *GetMaterialByIdResponseBodyData) GetSrcFrom() *string {
	return s.SrcFrom
}

func (s *GetMaterialByIdResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *GetMaterialByIdResponseBodyData) GetTextContent() *string {
	return s.TextContent
}

func (s *GetMaterialByIdResponseBodyData) GetThumbnailInBase64() *string {
	return s.ThumbnailInBase64
}

func (s *GetMaterialByIdResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetMaterialByIdResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMaterialByIdResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *GetMaterialByIdResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetMaterialByIdResponseBodyData) SetAuthor(v string) *GetMaterialByIdResponseBodyData {
	s.Author = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetCreateTime(v string) *GetMaterialByIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetCreateUser(v string) *GetMaterialByIdResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetDocKeywords(v []*string) *GetMaterialByIdResponseBodyData {
	s.DocKeywords = v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetDocType(v string) *GetMaterialByIdResponseBodyData {
	s.DocType = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetExternalUrl(v string) *GetMaterialByIdResponseBodyData {
	s.ExternalUrl = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetHtmlContent(v string) *GetMaterialByIdResponseBodyData {
	s.HtmlContent = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetId(v int64) *GetMaterialByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetPubTime(v string) *GetMaterialByIdResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetPublicUrl(v string) *GetMaterialByIdResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetShareAttr(v int32) *GetMaterialByIdResponseBodyData {
	s.ShareAttr = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetSrcFrom(v string) *GetMaterialByIdResponseBodyData {
	s.SrcFrom = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetSummary(v string) *GetMaterialByIdResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetTextContent(v string) *GetMaterialByIdResponseBodyData {
	s.TextContent = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetThumbnailInBase64(v string) *GetMaterialByIdResponseBodyData {
	s.ThumbnailInBase64 = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetTitle(v string) *GetMaterialByIdResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUpdateTime(v string) *GetMaterialByIdResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUpdateUser(v string) *GetMaterialByIdResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUrl(v string) *GetMaterialByIdResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
