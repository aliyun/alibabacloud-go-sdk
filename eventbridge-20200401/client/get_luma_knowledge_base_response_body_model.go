// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLumaKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *KnowledgeBase) *GetLumaKnowledgeBaseResponseBody
	GetData() *KnowledgeBase
	SetMessage(v string) *GetLumaKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLumaKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLumaKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type GetLumaKnowledgeBaseResponseBody struct {
	// The response code returned by the operation. A value of Success indicates a successful call. Otherwise, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the knowledge base bound to the Agent.
	Data *KnowledgeBase `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. The value is Operation success when the call succeeds, or a specific error description when the call fails.
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

func (s GetLumaKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLumaKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetLumaKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLumaKnowledgeBaseResponseBody) GetData() *KnowledgeBase {
	return s.Data
}

func (s *GetLumaKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLumaKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLumaKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLumaKnowledgeBaseResponseBody) SetCode(v string) *GetLumaKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *GetLumaKnowledgeBaseResponseBody) SetData(v *KnowledgeBase) *GetLumaKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *GetLumaKnowledgeBaseResponseBody) SetMessage(v string) *GetLumaKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *GetLumaKnowledgeBaseResponseBody) SetRequestId(v string) *GetLumaKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLumaKnowledgeBaseResponseBody) SetSuccess(v bool) *GetLumaKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetLumaKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
