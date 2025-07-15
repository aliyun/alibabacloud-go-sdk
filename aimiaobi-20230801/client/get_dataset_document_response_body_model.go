// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDatasetDocumentResponseBody
	GetCode() *string
	SetData(v *GetDatasetDocumentResponseBodyData) *GetDatasetDocumentResponseBody
	GetData() *GetDatasetDocumentResponseBodyData
	SetHttpStatusCode(v int32) *GetDatasetDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDatasetDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDatasetDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatasetDocumentResponseBody
	GetSuccess() *bool
}

type GetDatasetDocumentResponseBody struct {
	// example:
	//
	// NoData
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDatasetDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetDatasetDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDatasetDocumentResponseBody) GetData() *GetDatasetDocumentResponseBodyData {
	return s.Data
}

func (s *GetDatasetDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDatasetDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDatasetDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatasetDocumentResponseBody) SetCode(v string) *GetDatasetDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetData(v *GetDatasetDocumentResponseBodyData) *GetDatasetDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetHttpStatusCode(v int32) *GetDatasetDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetMessage(v string) *GetDatasetDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetRequestId(v string) *GetDatasetDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetSuccess(v bool) *GetDatasetDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDatasetDocumentResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	DisableHandleMultimodalMedia *bool `json:"DisableHandleMultimodalMedia,omitempty" xml:"DisableHandleMultimodalMedia,omitempty"`
	// example:
	//
	// 用户指定的文档唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 内部文档唯一ID
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-05-14 08:54:33
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 来源
	SourceFrom *string `json:"SourceFrom,omitempty" xml:"SourceFrom,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDatasetDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetDatasetDocumentResponseBodyData) GetDisableHandleMultimodalMedia() *bool {
	return s.DisableHandleMultimodalMedia
}

func (s *GetDatasetDocumentResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *GetDatasetDocumentResponseBodyData) GetDocUuid() *string {
	return s.DocUuid
}

func (s *GetDatasetDocumentResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *GetDatasetDocumentResponseBodyData) GetSourceFrom() *string {
	return s.SourceFrom
}

func (s *GetDatasetDocumentResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *GetDatasetDocumentResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetDatasetDocumentResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetDatasetDocumentResponseBodyData) SetContent(v string) *GetDatasetDocumentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetDisableHandleMultimodalMedia(v bool) *GetDatasetDocumentResponseBodyData {
	s.DisableHandleMultimodalMedia = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetDocId(v string) *GetDatasetDocumentResponseBodyData {
	s.DocId = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetDocUuid(v string) *GetDatasetDocumentResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetPubTime(v string) *GetDatasetDocumentResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetSourceFrom(v string) *GetDatasetDocumentResponseBodyData {
	s.SourceFrom = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetSummary(v string) *GetDatasetDocumentResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetTitle(v string) *GetDatasetDocumentResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetUrl(v string) *GetDatasetDocumentResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
