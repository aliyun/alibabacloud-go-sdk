// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchLumaKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchLumaKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *SearchLumaKnowledgeBaseResponseBodyData) *SearchLumaKnowledgeBaseResponseBody
	GetData() *SearchLumaKnowledgeBaseResponseBodyData
	SetMessage(v string) *SearchLumaKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchLumaKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchLumaKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type SearchLumaKnowledgeBaseResponseBody struct {
	// The response code. A value of Success indicates a successful call. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The retrieval result from the knowledge base bound to the Agent.
	Data *SearchLumaKnowledgeBaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket feedback.
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

func (s SearchLumaKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchLumaKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *SearchLumaKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchLumaKnowledgeBaseResponseBody) GetData() *SearchLumaKnowledgeBaseResponseBodyData {
	return s.Data
}

func (s *SearchLumaKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchLumaKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchLumaKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchLumaKnowledgeBaseResponseBody) SetCode(v string) *SearchLumaKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *SearchLumaKnowledgeBaseResponseBody) SetData(v *SearchLumaKnowledgeBaseResponseBodyData) *SearchLumaKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *SearchLumaKnowledgeBaseResponseBody) SetMessage(v string) *SearchLumaKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *SearchLumaKnowledgeBaseResponseBody) SetRequestId(v string) *SearchLumaKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchLumaKnowledgeBaseResponseBody) SetSuccess(v bool) *SearchLumaKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *SearchLumaKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchLumaKnowledgeBaseResponseBodyData struct {
	// The list of matched text chunks, sorted by relevance.
	//
	// example:
	//
	// [{"DocumentId":"doc-bp1xxxxxxxxxxxx","ChunkSeq":3,"Content":"EventBridge supports event routing"}]
	Chunks []*KnowledgeBaseSearchChunk `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// The time spent on the retrieval, in milliseconds.
	//
	// example:
	//
	// 1200
	TimeSpent *int64 `json:"TimeSpent,omitempty" xml:"TimeSpent,omitempty"`
}

func (s SearchLumaKnowledgeBaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchLumaKnowledgeBaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchLumaKnowledgeBaseResponseBodyData) GetChunks() []*KnowledgeBaseSearchChunk {
	return s.Chunks
}

func (s *SearchLumaKnowledgeBaseResponseBodyData) GetTimeSpent() *int64 {
	return s.TimeSpent
}

func (s *SearchLumaKnowledgeBaseResponseBodyData) SetChunks(v []*KnowledgeBaseSearchChunk) *SearchLumaKnowledgeBaseResponseBodyData {
	s.Chunks = v
	return s
}

func (s *SearchLumaKnowledgeBaseResponseBodyData) SetTimeSpent(v int64) *SearchLumaKnowledgeBaseResponseBodyData {
	s.TimeSpent = &v
	return s
}

func (s *SearchLumaKnowledgeBaseResponseBodyData) Validate() error {
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
