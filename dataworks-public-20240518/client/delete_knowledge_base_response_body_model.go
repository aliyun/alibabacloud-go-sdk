// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *DeleteKnowledgeBaseResponseBodyData) *DeleteKnowledgeBaseResponseBody
	GetData() *DeleteKnowledgeBaseResponseBodyData
	SetMessage(v string) *DeleteKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type DeleteKnowledgeBaseResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned upon success.
	Data *DeleteKnowledgeBaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The security error message returned when the request fails.
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

func (s DeleteKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteKnowledgeBaseResponseBody) GetData() *DeleteKnowledgeBaseResponseBodyData {
	return s.Data
}

func (s *DeleteKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKnowledgeBaseResponseBody) SetCode(v string) *DeleteKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) SetData(v *DeleteKnowledgeBaseResponseBodyData) *DeleteKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) SetMessage(v string) *DeleteKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) SetRequestId(v string) *DeleteKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) SetSuccess(v bool) *DeleteKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteKnowledgeBaseResponseBodyData struct {
	// The associated knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The status. The value is fixed to DELETED upon success.
	//
	// example:
	//
	// DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteKnowledgeBaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeBaseResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteKnowledgeBaseResponseBodyData) SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBodyData) SetStatus(v string) *DeleteKnowledgeBaseResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
