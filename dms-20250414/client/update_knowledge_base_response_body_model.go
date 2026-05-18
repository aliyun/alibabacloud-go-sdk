// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OneMetaKnowledgeBase) *UpdateKnowledgeBaseResponseBody
	GetData() *OneMetaKnowledgeBase
	SetErrorCode(v string) *UpdateKnowledgeBaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateKnowledgeBaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type UpdateKnowledgeBaseResponseBody struct {
	Data *OneMetaKnowledgeBase `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *UpdateKnowledgeBaseResponseBody) GetData() *OneMetaKnowledgeBase {
	return s.Data
}

func (s *UpdateKnowledgeBaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateKnowledgeBaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateKnowledgeBaseResponseBody) SetData(v *OneMetaKnowledgeBase) *UpdateKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) SetErrorCode(v string) *UpdateKnowledgeBaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) SetErrorMessage(v string) *UpdateKnowledgeBaseResponseBody {
	s.ErrorMessage = &v
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
