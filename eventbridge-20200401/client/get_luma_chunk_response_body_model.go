// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaChunkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLumaChunkResponseBody
	GetCode() *string
	SetData(v *KnowledgeBaseChunk) *GetLumaChunkResponseBody
	GetData() *KnowledgeBaseChunk
	SetMessage(v string) *GetLumaChunkResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLumaChunkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLumaChunkResponseBody
	GetSuccess() *bool
}

type GetLumaChunkResponseBody struct {
	// The response code of the operation. A value of Success indicates that the call succeeds. Otherwise, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The chunk details, including the body content and position.
	Data *KnowledgeBaseChunk `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket submission.
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

func (s GetLumaChunkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLumaChunkResponseBody) GoString() string {
	return s.String()
}

func (s *GetLumaChunkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLumaChunkResponseBody) GetData() *KnowledgeBaseChunk {
	return s.Data
}

func (s *GetLumaChunkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLumaChunkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLumaChunkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLumaChunkResponseBody) SetCode(v string) *GetLumaChunkResponseBody {
	s.Code = &v
	return s
}

func (s *GetLumaChunkResponseBody) SetData(v *KnowledgeBaseChunk) *GetLumaChunkResponseBody {
	s.Data = v
	return s
}

func (s *GetLumaChunkResponseBody) SetMessage(v string) *GetLumaChunkResponseBody {
	s.Message = &v
	return s
}

func (s *GetLumaChunkResponseBody) SetRequestId(v string) *GetLumaChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLumaChunkResponseBody) SetSuccess(v bool) *GetLumaChunkResponseBody {
	s.Success = &v
	return s
}

func (s *GetLumaChunkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
