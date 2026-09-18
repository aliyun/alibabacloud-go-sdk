// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *SearchKnowledgeBaseResponseBodyData) *SearchKnowledgeBaseResponseBody
	GetData() *SearchKnowledgeBaseResponseBodyData
	SetMessage(v string) *SearchKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type SearchKnowledgeBaseResponseBody struct {
	// The response code. A value of Success indicates a successful call. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The search results, including the list of matched chunks and the time spent.
	Data *SearchKnowledgeBaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. A value of Operation success is returned for a successful call. A specific error description is returned for a failed call.
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
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchKnowledgeBaseResponseBody) GetData() *SearchKnowledgeBaseResponseBodyData {
	return s.Data
}

func (s *SearchKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchKnowledgeBaseResponseBody) SetCode(v string) *SearchKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetData(v *SearchKnowledgeBaseResponseBodyData) *SearchKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetMessage(v string) *SearchKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetRequestId(v string) *SearchKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetSuccess(v bool) *SearchKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchKnowledgeBaseResponseBodyData struct {
	// The list of matched chunks, sorted by relevance.
	//
	// example:
	//
	// [{"DocumentId":"doc-bp1xxxxxxxxxxxx","ChunkSeq":3,"Content":"EventBridge supports event routing"}]
	Chunks []*KnowledgeBaseSearchChunk `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// The server-side processing duration of this search. Unit: milliseconds.
	//
	// example:
	//
	// 128
	TimeSpent *int64 `json:"TimeSpent,omitempty" xml:"TimeSpent,omitempty"`
}

func (s SearchKnowledgeBaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyData) GetChunks() []*KnowledgeBaseSearchChunk {
	return s.Chunks
}

func (s *SearchKnowledgeBaseResponseBodyData) GetTimeSpent() *int64 {
	return s.TimeSpent
}

func (s *SearchKnowledgeBaseResponseBodyData) SetChunks(v []*KnowledgeBaseSearchChunk) *SearchKnowledgeBaseResponseBodyData {
	s.Chunks = v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyData) SetTimeSpent(v int64) *SearchKnowledgeBaseResponseBodyData {
	s.TimeSpent = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyData) Validate() error {
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
