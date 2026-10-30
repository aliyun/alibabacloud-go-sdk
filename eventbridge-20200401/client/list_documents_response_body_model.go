// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDocumentsResponseBody
	GetCode() *string
	SetData(v *ListDocumentsResponseBodyData) *ListDocumentsResponseBody
	GetData() *ListDocumentsResponseBodyData
	SetMessage(v string) *ListDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDocumentsResponseBody
	GetSuccess() *bool
}

type ListDocumentsResponseBody struct {
	// The response code of the API operation. A value of Success indicates that the call is successful. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result of the document list, which includes document entries and pagination information.
	Data *ListDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The prompt message returned by the API operation. If the call is successful, the value is Operation success. If the call fails, a specific error description is returned.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which is used for troubleshooting and ticket submission.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDocumentsResponseBody) GetData() *ListDocumentsResponseBodyData {
	return s.Data
}

func (s *ListDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDocumentsResponseBody) SetCode(v string) *ListDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDocumentsResponseBody) SetData(v *ListDocumentsResponseBodyData) *ListDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDocumentsResponseBody) SetMessage(v string) *ListDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDocumentsResponseBody) SetRequestId(v string) *ListDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocumentsResponseBody) SetSuccess(v bool) *ListDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDocumentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDocumentsResponseBodyData struct {
	// The list of document entries in the knowledge base. Each entry includes information such as the document ID, file name, processing status, size, number of chunks, and metadata.
	//
	// example:
	//
	// [{"DocumentId":"doc-bp1xxxxxxxxxxxx","FileName":"manual.pdf","Status":"COMPLETED","ChunkCount":120}]
	Documents []*KnowledgeBaseDocument `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The actual maximum number of results per page that takes effect for the current request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page, which is an opaque string. You can pass this token as the NextToken parameter in the next request to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of documents that meet the filtering conditions. Pagination ends when NextToken is empty.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBodyData) GetDocuments() []*KnowledgeBaseDocument {
	return s.Documents
}

func (s *ListDocumentsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDocumentsResponseBodyData) SetDocuments(v []*KnowledgeBaseDocument) *ListDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *ListDocumentsResponseBodyData) SetMaxResults(v int32) *ListDocumentsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentsResponseBodyData) SetNextToken(v string) *ListDocumentsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListDocumentsResponseBodyData) SetTotalCount(v int32) *ListDocumentsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDocumentsResponseBodyData) Validate() error {
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
