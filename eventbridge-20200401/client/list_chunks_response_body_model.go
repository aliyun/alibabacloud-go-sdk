// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChunksResponseBody
	GetCode() *string
	SetData(v *ListChunksResponseBodyData) *ListChunksResponseBody
	GetData() *ListChunksResponseBodyData
	SetMessage(v string) *ListChunksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChunksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChunksResponseBody
	GetSuccess() *bool
}

type ListChunksResponseBody struct {
	// The response code. A value of Success indicates that the call is successful. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result of the chunk list, which includes chunk entries and pagination information.
	Data *ListChunksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The prompt message returned by the API. If the call is successful, the value is Operation success. If the call fails, a specific error description is returned.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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

func (s ListChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChunksResponseBody) GetData() *ListChunksResponseBodyData {
	return s.Data
}

func (s *ListChunksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChunksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChunksResponseBody) SetCode(v string) *ListChunksResponseBody {
	s.Code = &v
	return s
}

func (s *ListChunksResponseBody) SetData(v *ListChunksResponseBodyData) *ListChunksResponseBody {
	s.Data = v
	return s
}

func (s *ListChunksResponseBody) SetMessage(v string) *ListChunksResponseBody {
	s.Message = &v
	return s
}

func (s *ListChunksResponseBody) SetRequestId(v string) *ListChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChunksResponseBody) SetSuccess(v bool) *ListChunksResponseBody {
	s.Success = &v
	return s
}

func (s *ListChunksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChunksResponseBodyData struct {
	// The list of chunk entries.
	//
	// example:
	//
	// [{"DocumentId":"doc-bp1xxxxxxxxxxxx","ChunkSeq":1,"Content":"EventBridge supports event routing"}]
	Chunks []*KnowledgeBaseChunk `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// The actual maximum number of results per page that takes effect for the current request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page, which is an opaque string. Pass this value to the NextToken parameter of the next request to obtain the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of chunks in the specified document. This parameter is returned only when DocumentId is specified. In full-base mode (when DocumentId is not specified), this parameter is not returned, and pagination ends when NextToken is empty.
	//
	// example:
	//
	// 120
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChunksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyData) GetChunks() []*KnowledgeBaseChunk {
	return s.Chunks
}

func (s *ListChunksResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListChunksResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChunksResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListChunksResponseBodyData) SetChunks(v []*KnowledgeBaseChunk) *ListChunksResponseBodyData {
	s.Chunks = v
	return s
}

func (s *ListChunksResponseBodyData) SetMaxResults(v int32) *ListChunksResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListChunksResponseBodyData) SetNextToken(v string) *ListChunksResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListChunksResponseBodyData) SetTotalCount(v int32) *ListChunksResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListChunksResponseBodyData) Validate() error {
	if s.Chunks != nil {
		for _, item := range s.Chunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
