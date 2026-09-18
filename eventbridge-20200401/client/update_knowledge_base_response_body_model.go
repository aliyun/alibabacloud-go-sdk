// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *KnowledgeBase) *UpdateKnowledgeBaseResponseBody
	GetData() *KnowledgeBase
	SetMessage(v string) *UpdateKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type UpdateKnowledgeBaseResponseBody struct {
	// The response code. A value of Success indicates that the call was successful. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the knowledge base after the update.
	Data *KnowledgeBase `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. A value of Operation success is returned if the call was successful. If the call fails, a specific error description is returned.
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

func (s UpdateKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKnowledgeBaseResponseBody) GetData() *KnowledgeBase {
	return s.Data
}

func (s *UpdateKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateKnowledgeBaseResponseBody) SetCode(v string) *UpdateKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) SetData(v *KnowledgeBase) *UpdateKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) SetMessage(v string) *UpdateKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) SetSuccess(v bool) *UpdateKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
