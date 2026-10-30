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
	SetData(v *KnowledgeBase) *DeleteKnowledgeBaseResponseBody
	GetData() *KnowledgeBase
	SetMessage(v string) *DeleteKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type DeleteKnowledgeBaseResponseBody struct {
	// The response code of the API operation. A value of Success indicates that the call is successful. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the deleted knowledge base.
	Data *KnowledgeBase `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the API operation. If the call is successful, the value is Operation success. If the call fails, a specific error description is returned.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is successful.
	//
	// example:
	//
	// true
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

func (s *DeleteKnowledgeBaseResponseBody) GetData() *KnowledgeBase {
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

func (s *DeleteKnowledgeBaseResponseBody) SetData(v *KnowledgeBase) *DeleteKnowledgeBaseResponseBody {
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
