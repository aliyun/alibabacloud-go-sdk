// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OneMetaKnowledgeBaseDocument) *ListDocumentsResponseBody
	GetData() []*OneMetaKnowledgeBaseDocument
	SetErrorCode(v string) *ListDocumentsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDocumentsResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDocumentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDocumentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDocumentsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListDocumentsResponseBody
	GetTotalCount() *int64
}

type ListDocumentsResponseBody struct {
	// A list of documents.
	Data []*OneMetaKnowledgeBaseDocument `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code, returned only when the request fails.
	//
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message, returned only when the request fails.
	//
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// zCXSmY0CJbybp6FZV7vo0Wjw64X-*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique request ID. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of documents that meet the specified criteria. This parameter is not currently supported and always returns 0.
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBody) GetData() []*OneMetaKnowledgeBaseDocument {
	return s.Data
}

func (s *ListDocumentsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDocumentsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDocumentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDocumentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDocumentsResponseBody) SetData(v []*OneMetaKnowledgeBaseDocument) *ListDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDocumentsResponseBody) SetErrorCode(v string) *ListDocumentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDocumentsResponseBody) SetErrorMessage(v string) *ListDocumentsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDocumentsResponseBody) SetMaxResults(v int32) *ListDocumentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentsResponseBody) SetNextToken(v string) *ListDocumentsResponseBody {
	s.NextToken = &v
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

func (s *ListDocumentsResponseBody) SetTotalCount(v int64) *ListDocumentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDocumentsResponseBody) Validate() error {
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
