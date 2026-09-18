// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChunkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetChunkResponseBody
	GetCode() *string
	SetData(v *KnowledgeBaseChunk) *GetChunkResponseBody
	GetData() *KnowledgeBaseChunk
	SetMessage(v string) *GetChunkResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChunkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChunkResponseBody
	GetSuccess() *bool
}

type GetChunkResponseBody struct {
	// The response code. Success indicates a successful call. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The chunk details, including the body content, position, and enabled status.
	Data *KnowledgeBaseChunk `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message. Operation success is returned if the call succeeds. A specific error description is returned if the call fails.
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

func (s GetChunkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChunkResponseBody) GoString() string {
	return s.String()
}

func (s *GetChunkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChunkResponseBody) GetData() *KnowledgeBaseChunk {
	return s.Data
}

func (s *GetChunkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChunkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChunkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChunkResponseBody) SetCode(v string) *GetChunkResponseBody {
	s.Code = &v
	return s
}

func (s *GetChunkResponseBody) SetData(v *KnowledgeBaseChunk) *GetChunkResponseBody {
	s.Data = v
	return s
}

func (s *GetChunkResponseBody) SetMessage(v string) *GetChunkResponseBody {
	s.Message = &v
	return s
}

func (s *GetChunkResponseBody) SetRequestId(v string) *GetChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChunkResponseBody) SetSuccess(v bool) *GetChunkResponseBody {
	s.Success = &v
	return s
}

func (s *GetChunkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
