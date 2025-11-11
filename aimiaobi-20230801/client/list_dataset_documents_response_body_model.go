// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDatasetDocumentsResponseBody
	GetCode() *string
	SetData(v []*ListDatasetDocumentsResponseBodyData) *ListDatasetDocumentsResponseBody
	GetData() []*ListDatasetDocumentsResponseBodyData
	SetHttpStatusCode(v int32) *ListDatasetDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDatasetDocumentsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListDatasetDocumentsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetDocumentsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDatasetDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatasetDocumentsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDatasetDocumentsResponseBody
	GetTotalCount() *int32
}

type ListDatasetDocumentsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListDatasetDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDatasetDocumentsResponseBody) GetData() []*ListDatasetDocumentsResponseBodyData {
	return s.Data
}

func (s *ListDatasetDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDatasetDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDatasetDocumentsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetDocumentsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatasetDocumentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDatasetDocumentsResponseBody) SetCode(v string) *ListDatasetDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetData(v []*ListDatasetDocumentsResponseBodyData) *ListDatasetDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetHttpStatusCode(v int32) *ListDatasetDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetMessage(v string) *ListDatasetDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetPageNumber(v int32) *ListDatasetDocumentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetPageSize(v int32) *ListDatasetDocumentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetRequestId(v string) *ListDatasetDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetSuccess(v bool) *ListDatasetDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) SetTotalCount(v int32) *ListDatasetDocumentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetDocumentsResponseBody) Validate() error {
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

type ListDatasetDocumentsResponseBodyData struct {
	// example:
	//
	// xx
	CategoryUuid *string `json:"CategoryUuid,omitempty" xml:"CategoryUuid,omitempty"`
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2025-04-14 19:59:53
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	DisableHandleMultimodalMedia *bool `json:"DisableHandleMultimodalMedia,omitempty" xml:"DisableHandleMultimodalMedia,omitempty"`
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// text
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// 内部文档唯一ID
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
	Extend3          *string                                                 `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	MultimodalMedias []*ListDatasetDocumentsResponseBodyDataMultimodalMedias `json:"MultimodalMedias,omitempty" xml:"MultimodalMedias,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-01-01 00:00:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 来源
	SourceFrom *string `json:"SourceFrom,omitempty" xml:"SourceFrom,omitempty"`
	// example:
	//
	// 100
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2025-04-14 19:59:53
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// url
	//
	// example:
	//
	// https://xxx/xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListDatasetDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDatasetDocumentsResponseBodyData) GetCategoryUuid() *string {
	return s.CategoryUuid
}

func (s *ListDatasetDocumentsResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListDatasetDocumentsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDatasetDocumentsResponseBodyData) GetDisableHandleMultimodalMedia() *bool {
	return s.DisableHandleMultimodalMedia
}

func (s *ListDatasetDocumentsResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *ListDatasetDocumentsResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *ListDatasetDocumentsResponseBodyData) GetDocUuid() *string {
	return s.DocUuid
}

func (s *ListDatasetDocumentsResponseBodyData) GetExtend1() *string {
	return s.Extend1
}

func (s *ListDatasetDocumentsResponseBodyData) GetExtend2() *string {
	return s.Extend2
}

func (s *ListDatasetDocumentsResponseBodyData) GetExtend3() *string {
	return s.Extend3
}

func (s *ListDatasetDocumentsResponseBodyData) GetMultimodalMedias() []*ListDatasetDocumentsResponseBodyDataMultimodalMedias {
	return s.MultimodalMedias
}

func (s *ListDatasetDocumentsResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *ListDatasetDocumentsResponseBodyData) GetSourceFrom() *string {
	return s.SourceFrom
}

func (s *ListDatasetDocumentsResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListDatasetDocumentsResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *ListDatasetDocumentsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListDatasetDocumentsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDatasetDocumentsResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListDatasetDocumentsResponseBodyData) SetCategoryUuid(v string) *ListDatasetDocumentsResponseBodyData {
	s.CategoryUuid = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetContent(v string) *ListDatasetDocumentsResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetCreateTime(v string) *ListDatasetDocumentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetDisableHandleMultimodalMedia(v bool) *ListDatasetDocumentsResponseBodyData {
	s.DisableHandleMultimodalMedia = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetDocId(v string) *ListDatasetDocumentsResponseBodyData {
	s.DocId = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetDocType(v string) *ListDatasetDocumentsResponseBodyData {
	s.DocType = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetDocUuid(v string) *ListDatasetDocumentsResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetExtend1(v string) *ListDatasetDocumentsResponseBodyData {
	s.Extend1 = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetExtend2(v string) *ListDatasetDocumentsResponseBodyData {
	s.Extend2 = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetExtend3(v string) *ListDatasetDocumentsResponseBodyData {
	s.Extend3 = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetMultimodalMedias(v []*ListDatasetDocumentsResponseBodyDataMultimodalMedias) *ListDatasetDocumentsResponseBodyData {
	s.MultimodalMedias = v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetPubTime(v string) *ListDatasetDocumentsResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetSourceFrom(v string) *ListDatasetDocumentsResponseBodyData {
	s.SourceFrom = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetStatus(v int32) *ListDatasetDocumentsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetSummary(v string) *ListDatasetDocumentsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetTitle(v string) *ListDatasetDocumentsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetUpdateTime(v string) *ListDatasetDocumentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) SetUrl(v string) *ListDatasetDocumentsResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyData) Validate() error {
	if s.MultimodalMedias != nil {
		for _, item := range s.MultimodalMedias {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDatasetDocumentsResponseBodyDataMultimodalMedias struct {
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

func (s ListDatasetDocumentsResponseBodyDataMultimodalMedias) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetDocumentsResponseBodyDataMultimodalMedias) GoString() string {
	return s.String()
}

func (s *ListDatasetDocumentsResponseBodyDataMultimodalMedias) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListDatasetDocumentsResponseBodyDataMultimodalMedias) GetMediaId() *string {
	return s.MediaId
}

func (s *ListDatasetDocumentsResponseBodyDataMultimodalMedias) GetMediaType() *string {
	return s.MediaType
}

func (s *ListDatasetDocumentsResponseBodyDataMultimodalMedias) SetFileUrl(v string) *ListDatasetDocumentsResponseBodyDataMultimodalMedias {
	s.FileUrl = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyDataMultimodalMedias) SetMediaId(v string) *ListDatasetDocumentsResponseBodyDataMultimodalMedias {
	s.MediaId = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyDataMultimodalMedias) SetMediaType(v string) *ListDatasetDocumentsResponseBodyDataMultimodalMedias {
	s.MediaType = &v
	return s
}

func (s *ListDatasetDocumentsResponseBodyDataMultimodalMedias) Validate() error {
	return dara.Validate(s)
}
