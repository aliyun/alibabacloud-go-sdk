// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaDocumentsResponseBody
	GetCode() *string
	SetData(v *ListLumaDocumentsResponseBodyData) *ListLumaDocumentsResponseBody
	GetData() *ListLumaDocumentsResponseBodyData
	SetMessage(v string) *ListLumaDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaDocumentsResponseBody
	GetSuccess() *bool
}

type ListLumaDocumentsResponseBody struct {
	// The response code. A value of Success indicates that the call succeeds. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The document list result, which contains document entries and pagination information.
	Data *ListLumaDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. You can use this ID for troubleshooting and when you submit a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call succeeds.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaDocumentsResponseBody) GetData() *ListLumaDocumentsResponseBodyData {
	return s.Data
}

func (s *ListLumaDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaDocumentsResponseBody) SetCode(v string) *ListLumaDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaDocumentsResponseBody) SetData(v *ListLumaDocumentsResponseBodyData) *ListLumaDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaDocumentsResponseBody) SetMessage(v string) *ListLumaDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaDocumentsResponseBody) SetRequestId(v string) *ListLumaDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaDocumentsResponseBody) SetSuccess(v bool) *ListLumaDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaDocumentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaDocumentsResponseBodyData struct {
	// The list of document entries.
	//
	// example:
	//
	// [{"DocumentId":"doc-bp1xxxxxxxxxxxx","FileName":"manual.pdf","Status":"COMPLETED"}]
	Documents []*KnowledgeBaseDocument `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The pagination token for the next page. This is an opaque string. Pass this value as the NextToken parameter in the next request to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of documents that match the filter conditions.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLumaDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaDocumentsResponseBodyData) GetDocuments() []*KnowledgeBaseDocument {
	return s.Documents
}

func (s *ListLumaDocumentsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaDocumentsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLumaDocumentsResponseBodyData) SetDocuments(v []*KnowledgeBaseDocument) *ListLumaDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *ListLumaDocumentsResponseBodyData) SetNextToken(v string) *ListLumaDocumentsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLumaDocumentsResponseBodyData) SetTotalCount(v int32) *ListLumaDocumentsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLumaDocumentsResponseBodyData) Validate() error {
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
