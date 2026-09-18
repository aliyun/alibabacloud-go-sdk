// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *KnowledgeBase) *CreateKnowledgeBaseResponseBody
	GetData() *KnowledgeBase
	SetMessage(v string) *CreateKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type CreateKnowledgeBaseResponseBody struct {
	// The response code. A value of Success indicates a successful operation. An error code is returned if the operation fails.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the created knowledge base, including the name, status, and configuration information.
	Data *KnowledgeBase `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message. A value of Operation success is returned if the operation succeeds. A specific error description is returned if the operation fails.
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
	// Indicates whether the operation is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseResponseBody) GetData() *KnowledgeBase {
	return s.Data
}

func (s *CreateKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKnowledgeBaseResponseBody) SetCode(v string) *CreateKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) SetData(v *KnowledgeBase) *CreateKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) SetMessage(v string) *CreateKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) SetRequestId(v string) *CreateKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) SetSuccess(v bool) *CreateKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
