// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *KnowledgeBase) *GetKnowledgeBaseResponseBody
	GetData() *KnowledgeBase
	SetMessage(v string) *GetKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type GetKnowledgeBaseResponseBody struct {
	// The response code. A value of Success indicates that the call was successful. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the knowledge base, including the basic information, metadata schema, chunking configuration, and retrieval configuration.
	Data *KnowledgeBase `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message. If the call is successful, the value is Operation success. If the call fails, a specific error description is returned.
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
	// Indicates whether the call was successful. A value of true indicates that the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseResponseBody) GetData() *KnowledgeBase {
	return s.Data
}

func (s *GetKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeBaseResponseBody) SetCode(v string) *GetKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetData(v *KnowledgeBase) *GetKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetMessage(v string) *GetKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetRequestId(v string) *GetKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetSuccess(v bool) *GetKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
