// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaterialDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMaterialDocumentsResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListMaterialDocumentsResponseBody
	GetCurrent() *int32
	SetData(v []*ListMaterialDocumentsResponseBodyData) *ListMaterialDocumentsResponseBody
	GetData() []*ListMaterialDocumentsResponseBodyData
	SetHttpStatusCode(v int32) *ListMaterialDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMaterialDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMaterialDocumentsResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListMaterialDocumentsResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListMaterialDocumentsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListMaterialDocumentsResponseBody
	GetTotal() *int32
}

type ListMaterialDocumentsResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                                   `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListMaterialDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMaterialDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMaterialDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMaterialDocumentsResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListMaterialDocumentsResponseBody) GetData() []*ListMaterialDocumentsResponseBodyData {
	return s.Data
}

func (s *ListMaterialDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMaterialDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMaterialDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMaterialDocumentsResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListMaterialDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMaterialDocumentsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListMaterialDocumentsResponseBody) SetCode(v string) *ListMaterialDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetCurrent(v int32) *ListMaterialDocumentsResponseBody {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetData(v []*ListMaterialDocumentsResponseBodyData) *ListMaterialDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetHttpStatusCode(v int32) *ListMaterialDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetMessage(v string) *ListMaterialDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetRequestId(v string) *ListMaterialDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetSize(v int32) *ListMaterialDocumentsResponseBody {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetSuccess(v bool) *ListMaterialDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetTotal(v int32) *ListMaterialDocumentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) Validate() error {
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

type ListMaterialDocumentsResponseBodyData struct {
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser     *string   `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	CreateUserName *string   `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DocKeywords    []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// https://www.example.com
	ExternalUrl *string                                        `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	FileAttr    *ListMaterialDocumentsResponseBodyDataFileAttr `json:"FileAttr,omitempty" xml:"FileAttr,omitempty" type:"Struct"`
	FileKey     *string                                        `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	HtmlContent *string                                        `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// example:
	//
	// 35
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// https://www.example.com
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// user_upload
	SrcFrom           *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary           *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent       *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	ThumbnailInBase64 *string `json:"ThumbnailInBase64,omitempty" xml:"ThumbnailInBase64,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser     *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListMaterialDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMaterialDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponseBodyData) GetAuthor() *string {
	return s.Author
}

func (s *ListMaterialDocumentsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMaterialDocumentsResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListMaterialDocumentsResponseBodyData) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListMaterialDocumentsResponseBodyData) GetDocKeywords() []*string {
	return s.DocKeywords
}

func (s *ListMaterialDocumentsResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *ListMaterialDocumentsResponseBodyData) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *ListMaterialDocumentsResponseBodyData) GetFileAttr() *ListMaterialDocumentsResponseBodyDataFileAttr {
	return s.FileAttr
}

func (s *ListMaterialDocumentsResponseBodyData) GetFileKey() *string {
	return s.FileKey
}

func (s *ListMaterialDocumentsResponseBodyData) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *ListMaterialDocumentsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListMaterialDocumentsResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *ListMaterialDocumentsResponseBodyData) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *ListMaterialDocumentsResponseBodyData) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *ListMaterialDocumentsResponseBodyData) GetSrcFrom() *string {
	return s.SrcFrom
}

func (s *ListMaterialDocumentsResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *ListMaterialDocumentsResponseBodyData) GetTextContent() *string {
	return s.TextContent
}

func (s *ListMaterialDocumentsResponseBodyData) GetThumbnailInBase64() *string {
	return s.ThumbnailInBase64
}

func (s *ListMaterialDocumentsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListMaterialDocumentsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMaterialDocumentsResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *ListMaterialDocumentsResponseBodyData) GetUpdateUserName() *string {
	return s.UpdateUserName
}

func (s *ListMaterialDocumentsResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListMaterialDocumentsResponseBodyData) SetAuthor(v string) *ListMaterialDocumentsResponseBodyData {
	s.Author = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateUser(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateUserName(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetDocKeywords(v []*string) *ListMaterialDocumentsResponseBodyData {
	s.DocKeywords = v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetDocType(v string) *ListMaterialDocumentsResponseBodyData {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetExternalUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.ExternalUrl = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetFileAttr(v *ListMaterialDocumentsResponseBodyDataFileAttr) *ListMaterialDocumentsResponseBodyData {
	s.FileAttr = v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetFileKey(v string) *ListMaterialDocumentsResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetHtmlContent(v string) *ListMaterialDocumentsResponseBodyData {
	s.HtmlContent = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetId(v int64) *ListMaterialDocumentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetPubTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetPublicUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetShareAttr(v int32) *ListMaterialDocumentsResponseBodyData {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetSrcFrom(v string) *ListMaterialDocumentsResponseBodyData {
	s.SrcFrom = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetSummary(v string) *ListMaterialDocumentsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetTextContent(v string) *ListMaterialDocumentsResponseBodyData {
	s.TextContent = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetThumbnailInBase64(v string) *ListMaterialDocumentsResponseBodyData {
	s.ThumbnailInBase64 = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetTitle(v string) *ListMaterialDocumentsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateUser(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateUserName(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) Validate() error {
	if s.FileAttr != nil {
		if err := s.FileAttr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMaterialDocumentsResponseBodyDataFileAttr struct {
	Duration   *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileLength *int64   `json:"FileLength,omitempty" xml:"FileLength,omitempty"`
	FileName   *string  `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Height     *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	MimeType   *string  `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
	Width      *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListMaterialDocumentsResponseBodyDataFileAttr) String() string {
	return dara.Prettify(s)
}

func (s ListMaterialDocumentsResponseBodyDataFileAttr) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) GetDuration() *float64 {
	return s.Duration
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) GetFileLength() *int64 {
	return s.FileLength
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) GetFileName() *string {
	return s.FileName
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) GetHeight() *int32 {
	return s.Height
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) GetMimeType() *string {
	return s.MimeType
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) GetWidth() *int32 {
	return s.Width
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) SetDuration(v float64) *ListMaterialDocumentsResponseBodyDataFileAttr {
	s.Duration = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) SetFileLength(v int64) *ListMaterialDocumentsResponseBodyDataFileAttr {
	s.FileLength = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) SetFileName(v string) *ListMaterialDocumentsResponseBodyDataFileAttr {
	s.FileName = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) SetHeight(v int32) *ListMaterialDocumentsResponseBodyDataFileAttr {
	s.Height = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) SetMimeType(v string) *ListMaterialDocumentsResponseBodyDataFileAttr {
	s.MimeType = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) SetWidth(v int32) *ListMaterialDocumentsResponseBodyDataFileAttr {
	s.Width = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyDataFileAttr) Validate() error {
	return dara.Validate(s)
}
