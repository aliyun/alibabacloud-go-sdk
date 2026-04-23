// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmInventoryKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ConfirmInventoryKnowledgeResponseBody
	GetData() *bool
	SetErrorCode(v string) *ConfirmInventoryKnowledgeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ConfirmInventoryKnowledgeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ConfirmInventoryKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfirmInventoryKnowledgeResponseBody
	GetSuccess() *bool
}

type ConfirmInventoryKnowledgeResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmInventoryKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmInventoryKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmInventoryKnowledgeResponseBody) GetData() *bool {
	return s.Data
}

func (s *ConfirmInventoryKnowledgeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ConfirmInventoryKnowledgeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ConfirmInventoryKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmInventoryKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmInventoryKnowledgeResponseBody) SetData(v bool) *ConfirmInventoryKnowledgeResponseBody {
	s.Data = &v
	return s
}

func (s *ConfirmInventoryKnowledgeResponseBody) SetErrorCode(v string) *ConfirmInventoryKnowledgeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfirmInventoryKnowledgeResponseBody) SetErrorMessage(v string) *ConfirmInventoryKnowledgeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ConfirmInventoryKnowledgeResponseBody) SetRequestId(v string) *ConfirmInventoryKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmInventoryKnowledgeResponseBody) SetSuccess(v bool) *ConfirmInventoryKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmInventoryKnowledgeResponseBody) Validate() error {
	return dara.Validate(s)
}
