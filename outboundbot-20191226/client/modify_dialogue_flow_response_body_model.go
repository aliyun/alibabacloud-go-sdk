// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDialogueFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDialogueFlowResponseBody
	GetCode() *string
	SetDialogueFlowDefinition(v string) *ModifyDialogueFlowResponseBody
	GetDialogueFlowDefinition() *string
	SetDialogueFlowId(v string) *ModifyDialogueFlowResponseBody
	GetDialogueFlowId() *string
	SetHttpStatusCode(v int32) *ModifyDialogueFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyDialogueFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDialogueFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDialogueFlowResponseBody
	GetSuccess() *bool
}

type ModifyDialogueFlowResponseBody struct {
	// example:
	//
	// OK
	Code                   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DialogueFlowDefinition *string `json:"DialogueFlowDefinition,omitempty" xml:"DialogueFlowDefinition,omitempty"`
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Succcess
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDialogueFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDialogueFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDialogueFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDialogueFlowResponseBody) GetDialogueFlowDefinition() *string {
	return s.DialogueFlowDefinition
}

func (s *ModifyDialogueFlowResponseBody) GetDialogueFlowId() *string {
	return s.DialogueFlowId
}

func (s *ModifyDialogueFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDialogueFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDialogueFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDialogueFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDialogueFlowResponseBody) SetCode(v string) *ModifyDialogueFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetDialogueFlowDefinition(v string) *ModifyDialogueFlowResponseBody {
	s.DialogueFlowDefinition = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetDialogueFlowId(v string) *ModifyDialogueFlowResponseBody {
	s.DialogueFlowId = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetHttpStatusCode(v int32) *ModifyDialogueFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetMessage(v string) *ModifyDialogueFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetRequestId(v string) *ModifyDialogueFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) SetSuccess(v bool) *ModifyDialogueFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDialogueFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
