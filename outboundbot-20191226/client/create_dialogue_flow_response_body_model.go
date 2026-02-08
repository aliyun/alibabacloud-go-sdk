// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogueFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDialogueFlowResponseBody
	GetCode() *string
	SetDialogueFlowId(v string) *CreateDialogueFlowResponseBody
	GetDialogueFlowId() *string
	SetHttpStatusCode(v int32) *CreateDialogueFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDialogueFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDialogueFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDialogueFlowResponseBody
	GetSuccess() *bool
}

type CreateDialogueFlowResponseBody struct {
	// Status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Conversation flow ID
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDialogueFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogueFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDialogueFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDialogueFlowResponseBody) GetDialogueFlowId() *string {
	return s.DialogueFlowId
}

func (s *CreateDialogueFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDialogueFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDialogueFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDialogueFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDialogueFlowResponseBody) SetCode(v string) *CreateDialogueFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDialogueFlowResponseBody) SetDialogueFlowId(v string) *CreateDialogueFlowResponseBody {
	s.DialogueFlowId = &v
	return s
}

func (s *CreateDialogueFlowResponseBody) SetHttpStatusCode(v int32) *CreateDialogueFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDialogueFlowResponseBody) SetMessage(v string) *CreateDialogueFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDialogueFlowResponseBody) SetRequestId(v string) *CreateDialogueFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDialogueFlowResponseBody) SetSuccess(v bool) *CreateDialogueFlowResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDialogueFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
