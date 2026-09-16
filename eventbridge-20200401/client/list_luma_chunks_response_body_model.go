// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaChunksResponseBody
	GetCode() *string
	SetData(v *ListLumaChunksResponseBodyData) *ListLumaChunksResponseBody
	GetData() *ListLumaChunksResponseBodyData
	SetMessage(v string) *ListLumaChunksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaChunksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaChunksResponseBody
	GetSuccess() *bool
}

type ListLumaChunksResponseBody struct {
	// The response code. A value of Success indicates that the call succeeded. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The chunk list result, which contains chunk entries and pagination information.
	Data *ListLumaChunksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. Use this ID for troubleshooting and when submitting a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaChunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaChunksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaChunksResponseBody) GetData() *ListLumaChunksResponseBodyData {
	return s.Data
}

func (s *ListLumaChunksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaChunksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaChunksResponseBody) SetCode(v string) *ListLumaChunksResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaChunksResponseBody) SetData(v *ListLumaChunksResponseBodyData) *ListLumaChunksResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaChunksResponseBody) SetMessage(v string) *ListLumaChunksResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaChunksResponseBody) SetRequestId(v string) *ListLumaChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaChunksResponseBody) SetSuccess(v bool) *ListLumaChunksResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaChunksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaChunksResponseBodyData struct {
	// The list of chunk entries.
	//
	// example:
	//
	// [{"DocumentId":"doc-bp1xxxxxxxxxxxx","ChunkSeq":1}]
	Chunks []*KnowledgeBaseChunk `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// The pagination token for the next page (an opaque string). Pass this value as the NextToken parameter in the next request to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of chunks in the document.
	//
	// example:
	//
	// 120
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLumaChunksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaChunksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaChunksResponseBodyData) GetChunks() []*KnowledgeBaseChunk {
	return s.Chunks
}

func (s *ListLumaChunksResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaChunksResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLumaChunksResponseBodyData) SetChunks(v []*KnowledgeBaseChunk) *ListLumaChunksResponseBodyData {
	s.Chunks = v
	return s
}

func (s *ListLumaChunksResponseBodyData) SetNextToken(v string) *ListLumaChunksResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLumaChunksResponseBodyData) SetTotalCount(v int32) *ListLumaChunksResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLumaChunksResponseBodyData) Validate() error {
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
