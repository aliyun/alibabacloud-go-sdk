// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteKnowledgeSourceResponseBody
	GetCode() *string
	SetData(v *DeleteKnowledgeSourceResponseBodyData) *DeleteKnowledgeSourceResponseBody
	GetData() *DeleteKnowledgeSourceResponseBodyData
	SetMessage(v string) *DeleteKnowledgeSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteKnowledgeSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKnowledgeSourceResponseBody
	GetSuccess() *bool
}

type DeleteKnowledgeSourceResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *DeleteKnowledgeSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKnowledgeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteKnowledgeSourceResponseBody) GetData() *DeleteKnowledgeSourceResponseBodyData {
	return s.Data
}

func (s *DeleteKnowledgeSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteKnowledgeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKnowledgeSourceResponseBody) SetCode(v string) *DeleteKnowledgeSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKnowledgeSourceResponseBody) SetData(v *DeleteKnowledgeSourceResponseBodyData) *DeleteKnowledgeSourceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteKnowledgeSourceResponseBody) SetMessage(v string) *DeleteKnowledgeSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteKnowledgeSourceResponseBody) SetRequestId(v string) *DeleteKnowledgeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeSourceResponseBody) SetSuccess(v bool) *DeleteKnowledgeSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKnowledgeSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteKnowledgeSourceResponseBodyData struct {
	// The ID of the knowledge base to which the source belongs.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the deleted source.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s DeleteKnowledgeSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeSourceResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeSourceResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *DeleteKnowledgeSourceResponseBodyData) SetKnowledgeBaseId(v string) *DeleteKnowledgeSourceResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeSourceResponseBodyData) SetSourceId(v string) *DeleteKnowledgeSourceResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *DeleteKnowledgeSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
