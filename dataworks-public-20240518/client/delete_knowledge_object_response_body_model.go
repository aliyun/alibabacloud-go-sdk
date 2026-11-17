// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteKnowledgeObjectResponseBody
	GetCode() *string
	SetData(v *DeleteKnowledgeObjectResponseBodyData) *DeleteKnowledgeObjectResponseBody
	GetData() *DeleteKnowledgeObjectResponseBodyData
	SetMessage(v string) *DeleteKnowledgeObjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteKnowledgeObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKnowledgeObjectResponseBody
	GetSuccess() *bool
}

type DeleteKnowledgeObjectResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *DeleteKnowledgeObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The security error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKnowledgeObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteKnowledgeObjectResponseBody) GetData() *DeleteKnowledgeObjectResponseBodyData {
	return s.Data
}

func (s *DeleteKnowledgeObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteKnowledgeObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKnowledgeObjectResponseBody) SetCode(v string) *DeleteKnowledgeObjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKnowledgeObjectResponseBody) SetData(v *DeleteKnowledgeObjectResponseBodyData) *DeleteKnowledgeObjectResponseBody {
	s.Data = v
	return s
}

func (s *DeleteKnowledgeObjectResponseBody) SetMessage(v string) *DeleteKnowledgeObjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteKnowledgeObjectResponseBody) SetRequestId(v string) *DeleteKnowledgeObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeObjectResponseBody) SetSuccess(v bool) *DeleteKnowledgeObjectResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKnowledgeObjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteKnowledgeObjectResponseBodyData struct {
	// The ID of the knowledge base to which the object belongs.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the knowledge object. For semantic models, use the model ID from the semantic layer. For documents, use the stable ID from the local document table. Callers must treat the ID as opaque and must not use a retrieval ID in its place.
	//
	// example:
	//
	// obj-document-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
}

func (s DeleteKnowledgeObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeObjectResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeObjectResponseBodyData) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *DeleteKnowledgeObjectResponseBodyData) SetKnowledgeBaseId(v string) *DeleteKnowledgeObjectResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeObjectResponseBodyData) SetKnowledgeObjectId(v string) *DeleteKnowledgeObjectResponseBodyData {
	s.KnowledgeObjectId = &v
	return s
}

func (s *DeleteKnowledgeObjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
