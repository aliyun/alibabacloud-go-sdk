// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OneMetaKnowledgeBaseChunk) *ListDocumentChunksResponseBody
	GetData() []*OneMetaKnowledgeBaseChunk
	SetErrorCode(v string) *ListDocumentChunksResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDocumentChunksResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDocumentChunksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDocumentChunksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDocumentChunksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDocumentChunksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListDocumentChunksResponseBody
	GetTotalCount() *int64
}

type ListDocumentChunksResponseBody struct {
	// A list of chunks.
	Data []*OneMetaKnowledgeBaseChunk `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The number of entries returned on this page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page of results.
	//
	// example:
	//
	// zCXSmY0CJbybp6FZV7vo0Wjw64X-*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique request ID. If you encounter an error, provide this ID for troubleshooting.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates if the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of chunks that meet the filter criteria. (This feature is not yet supported, and the value is always 0.)
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDocumentChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentChunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentChunksResponseBody) GetData() []*OneMetaKnowledgeBaseChunk {
	return s.Data
}

func (s *ListDocumentChunksResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDocumentChunksResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDocumentChunksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentChunksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentChunksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDocumentChunksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDocumentChunksResponseBody) SetData(v []*OneMetaKnowledgeBaseChunk) *ListDocumentChunksResponseBody {
	s.Data = v
	return s
}

func (s *ListDocumentChunksResponseBody) SetErrorCode(v string) *ListDocumentChunksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDocumentChunksResponseBody) SetErrorMessage(v string) *ListDocumentChunksResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDocumentChunksResponseBody) SetMaxResults(v int32) *ListDocumentChunksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentChunksResponseBody) SetNextToken(v string) *ListDocumentChunksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDocumentChunksResponseBody) SetRequestId(v string) *ListDocumentChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocumentChunksResponseBody) SetSuccess(v bool) *ListDocumentChunksResponseBody {
	s.Success = &v
	return s
}

func (s *ListDocumentChunksResponseBody) SetTotalCount(v int64) *ListDocumentChunksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDocumentChunksResponseBody) Validate() error {
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
