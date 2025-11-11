// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatasetDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchDatasetDocumentsResponseBody
	GetCode() *string
	SetData(v *SearchDatasetDocumentsResponseBodyData) *SearchDatasetDocumentsResponseBody
	GetData() *SearchDatasetDocumentsResponseBodyData
	SetHttpStatusCode(v int32) *SearchDatasetDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SearchDatasetDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchDatasetDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchDatasetDocumentsResponseBody
	GetSuccess() *bool
}

type SearchDatasetDocumentsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SearchDatasetDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SearchDatasetDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchDatasetDocumentsResponseBody) GetData() *SearchDatasetDocumentsResponseBodyData {
	return s.Data
}

func (s *SearchDatasetDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchDatasetDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchDatasetDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchDatasetDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchDatasetDocumentsResponseBody) SetCode(v string) *SearchDatasetDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetData(v *SearchDatasetDocumentsResponseBodyData) *SearchDatasetDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetHttpStatusCode(v int32) *SearchDatasetDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetMessage(v string) *SearchDatasetDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetRequestId(v string) *SearchDatasetDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetSuccess(v bool) *SearchDatasetDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchDatasetDocumentsResponseBodyData struct {
	Documents []*SearchDatasetDocumentsResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
}

func (s SearchDatasetDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponseBodyData) GetDocuments() []*SearchDatasetDocumentsResponseBodyDataDocuments {
	return s.Documents
}

func (s *SearchDatasetDocumentsResponseBodyData) SetDocuments(v []*SearchDatasetDocumentsResponseBodyDataDocuments) *SearchDatasetDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyData) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchDatasetDocumentsResponseBodyDataDocuments struct {
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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
	// 2024-12-09 17:09:40
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 来源
	SourceFrom *string `json:"SourceFrom,omitempty" xml:"SourceFrom,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchDatasetDocumentsResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetContent() *string {
	return s.Content
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetDocId() *string {
	return s.DocId
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetDocUuid() *string {
	return s.DocUuid
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetPubTime() *string {
	return s.PubTime
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetSourceFrom() *string {
	return s.SourceFrom
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetSummary() *string {
	return s.Summary
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetTitle() *string {
	return s.Title
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetUrl() *string {
	return s.Url
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetContent(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Content = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetDocId(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetDocUuid(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.DocUuid = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetPubTime(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.PubTime = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetSourceFrom(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.SourceFrom = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetSummary(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Summary = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetTitle(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Title = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetUrl(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Url = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) Validate() error {
	return dara.Validate(s)
}
