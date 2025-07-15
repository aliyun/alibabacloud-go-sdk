// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTaskDialogueDatasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArticles(v []*ListSearchTaskDialogueDatasResponseBodyArticles) *ListSearchTaskDialogueDatasResponseBody
	GetArticles() []*ListSearchTaskDialogueDatasResponseBodyArticles
	SetCode(v string) *ListSearchTaskDialogueDatasResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSearchTaskDialogueDatasResponseBody
	GetHttpStatusCode() *int32
	SetImages(v []*ListSearchTaskDialogueDatasResponseBodyImages) *ListSearchTaskDialogueDatasResponseBody
	GetImages() []*ListSearchTaskDialogueDatasResponseBodyImages
	SetMessage(v string) *ListSearchTaskDialogueDatasResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListSearchTaskDialogueDatasResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSearchTaskDialogueDatasResponseBody
	GetPageSize() *int32
	SetRealtimeSearch(v bool) *ListSearchTaskDialogueDatasResponseBody
	GetRealtimeSearch() *bool
	SetRequestId(v string) *ListSearchTaskDialogueDatasResponseBody
	GetRequestId() *string
	SetSearchType(v string) *ListSearchTaskDialogueDatasResponseBody
	GetSearchType() *string
	SetSuccess(v bool) *ListSearchTaskDialogueDatasResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSearchTaskDialogueDatasResponseBody
	GetTotalCount() *int32
	SetVideos(v []*ListSearchTaskDialogueDatasResponseBodyVideos) *ListSearchTaskDialogueDatasResponseBody
	GetVideos() []*ListSearchTaskDialogueDatasResponseBodyVideos
}

type ListSearchTaskDialogueDatasResponseBody struct {
	Articles []*ListSearchTaskDialogueDatasResponseBodyArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Images         []*ListSearchTaskDialogueDatasResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	RealtimeSearch *bool `json:"RealtimeSearch,omitempty" xml:"RealtimeSearch,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// realtime
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Videos     []*ListSearchTaskDialogueDatasResponseBodyVideos `json:"Videos,omitempty" xml:"Videos,omitempty" type:"Repeated"`
}

func (s ListSearchTaskDialogueDatasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialogueDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetArticles() []*ListSearchTaskDialogueDatasResponseBodyArticles {
	return s.Articles
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetImages() []*ListSearchTaskDialogueDatasResponseBodyImages {
	return s.Images
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetRealtimeSearch() *bool {
	return s.RealtimeSearch
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetSearchType() *string {
	return s.SearchType
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSearchTaskDialogueDatasResponseBody) GetVideos() []*ListSearchTaskDialogueDatasResponseBodyVideos {
	return s.Videos
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetArticles(v []*ListSearchTaskDialogueDatasResponseBodyArticles) *ListSearchTaskDialogueDatasResponseBody {
	s.Articles = v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetCode(v string) *ListSearchTaskDialogueDatasResponseBody {
	s.Code = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetHttpStatusCode(v int32) *ListSearchTaskDialogueDatasResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetImages(v []*ListSearchTaskDialogueDatasResponseBodyImages) *ListSearchTaskDialogueDatasResponseBody {
	s.Images = v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetMessage(v string) *ListSearchTaskDialogueDatasResponseBody {
	s.Message = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetPageNumber(v int32) *ListSearchTaskDialogueDatasResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetPageSize(v int32) *ListSearchTaskDialogueDatasResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetRealtimeSearch(v bool) *ListSearchTaskDialogueDatasResponseBody {
	s.RealtimeSearch = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetRequestId(v string) *ListSearchTaskDialogueDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetSearchType(v string) *ListSearchTaskDialogueDatasResponseBody {
	s.SearchType = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetSuccess(v bool) *ListSearchTaskDialogueDatasResponseBody {
	s.Success = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetTotalCount(v int32) *ListSearchTaskDialogueDatasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) SetVideos(v []*ListSearchTaskDialogueDatasResponseBodyVideos) *ListSearchTaskDialogueDatasResponseBody {
	s.Videos = v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSearchTaskDialogueDatasResponseBodyArticles struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// xx
	CategoryUuid *string `json:"CategoryUuid,omitempty" xml:"CategoryUuid,omitempty"`
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文档-自定义的唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// text
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// xxx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xx
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	// example:
	//
	// xx
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	// example:
	//
	// xx
	Extend3          *string                                                            `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	MultimodalMedias []*ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias `json:"MultimodalMedias,omitempty" xml:"MultimodalMedias,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-11-25 14:25:59
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 新华社
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSearchTaskDialogueDatasResponseBodyArticles) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialogueDatasResponseBodyArticles) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetAuthor() *string {
	return s.Author
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetCategoryUuid() *string {
	return s.CategoryUuid
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetContent() *string {
	return s.Content
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetDocId() *string {
	return s.DocId
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetDocType() *string {
	return s.DocType
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetExtend1() *string {
	return s.Extend1
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetExtend2() *string {
	return s.Extend2
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetExtend3() *string {
	return s.Extend3
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetMultimodalMedias() []*ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias {
	return s.MultimodalMedias
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetSource() *string {
	return s.Source
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetSummary() *string {
	return s.Summary
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetTitle() *string {
	return s.Title
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) GetUrl() *string {
	return s.Url
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetAuthor(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Author = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetCategoryUuid(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.CategoryUuid = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetContent(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Content = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetDocId(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.DocId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetDocType(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.DocType = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetDocUuid(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.DocUuid = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetExtend1(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Extend1 = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetExtend2(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Extend2 = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetExtend3(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Extend3 = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetMultimodalMedias(v []*ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.MultimodalMedias = v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetPubTime(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.PubTime = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetSource(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Source = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetSummary(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Summary = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetTitle(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Title = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) SetUrl(v string) *ListSearchTaskDialogueDatasResponseBodyArticles {
	s.Url = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticles) Validate() error {
	return dara.Validate(s)
}

type ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias struct {
	// example:
	//
	// 图片或视频文件地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 多模态数据唯一标识
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 多模态数据类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) GetMediaId() *string {
	return s.MediaId
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) GetMediaType() *string {
	return s.MediaType
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) SetFileUrl(v string) *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias {
	s.FileUrl = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) SetMediaId(v string) *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias {
	s.MediaId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) SetMediaType(v string) *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias {
	s.MediaType = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyArticlesMultimodalMedias) Validate() error {
	return dara.Validate(s)
}

type ListSearchTaskDialogueDatasResponseBodyImages struct {
	// example:
	//
	// 图片或视频文件地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 多模态数据唯一标识
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 多模态数据类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s ListSearchTaskDialogueDatasResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialogueDatasResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialogueDatasResponseBodyImages) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListSearchTaskDialogueDatasResponseBodyImages) GetMediaId() *string {
	return s.MediaId
}

func (s *ListSearchTaskDialogueDatasResponseBodyImages) GetMediaType() *string {
	return s.MediaType
}

func (s *ListSearchTaskDialogueDatasResponseBodyImages) SetFileUrl(v string) *ListSearchTaskDialogueDatasResponseBodyImages {
	s.FileUrl = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyImages) SetMediaId(v string) *ListSearchTaskDialogueDatasResponseBodyImages {
	s.MediaId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyImages) SetMediaType(v string) *ListSearchTaskDialogueDatasResponseBodyImages {
	s.MediaType = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyImages) Validate() error {
	return dara.Validate(s)
}

type ListSearchTaskDialogueDatasResponseBodyVideos struct {
	// example:
	//
	// 图片或视频文件地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 多模态数据唯一标识
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 多模态数据类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s ListSearchTaskDialogueDatasResponseBodyVideos) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialogueDatasResponseBodyVideos) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialogueDatasResponseBodyVideos) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListSearchTaskDialogueDatasResponseBodyVideos) GetMediaId() *string {
	return s.MediaId
}

func (s *ListSearchTaskDialogueDatasResponseBodyVideos) GetMediaType() *string {
	return s.MediaType
}

func (s *ListSearchTaskDialogueDatasResponseBodyVideos) SetFileUrl(v string) *ListSearchTaskDialogueDatasResponseBodyVideos {
	s.FileUrl = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyVideos) SetMediaId(v string) *ListSearchTaskDialogueDatasResponseBodyVideos {
	s.MediaId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyVideos) SetMediaType(v string) *ListSearchTaskDialogueDatasResponseBodyVideos {
	s.MediaType = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponseBodyVideos) Validate() error {
	return dara.Validate(s)
}
