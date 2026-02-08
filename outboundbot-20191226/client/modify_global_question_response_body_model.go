// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyGlobalQuestionResponseBody
	GetCode() *string
	SetDialogueQuestionId(v string) *ModifyGlobalQuestionResponseBody
	GetDialogueQuestionId() *string
	SetHttpStatusCode(v int32) *ModifyGlobalQuestionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyGlobalQuestionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyGlobalQuestionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyGlobalQuestionResponseBody
	GetSuccess() *bool
}

type ModifyGlobalQuestionResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Global question ID
	//
	// example:
	//
	// ad80de88-1661-445a-92ec-bf88dc45d581
	DialogueQuestionId *string `json:"DialogueQuestionId,omitempty" xml:"DialogueQuestionId,omitempty"`
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

func (s ModifyGlobalQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalQuestionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyGlobalQuestionResponseBody) GetDialogueQuestionId() *string {
	return s.DialogueQuestionId
}

func (s *ModifyGlobalQuestionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyGlobalQuestionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyGlobalQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalQuestionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyGlobalQuestionResponseBody) SetCode(v string) *ModifyGlobalQuestionResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGlobalQuestionResponseBody) SetDialogueQuestionId(v string) *ModifyGlobalQuestionResponseBody {
	s.DialogueQuestionId = &v
	return s
}

func (s *ModifyGlobalQuestionResponseBody) SetHttpStatusCode(v int32) *ModifyGlobalQuestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyGlobalQuestionResponseBody) SetMessage(v string) *ModifyGlobalQuestionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGlobalQuestionResponseBody) SetRequestId(v string) *ModifyGlobalQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalQuestionResponseBody) SetSuccess(v bool) *ModifyGlobalQuestionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGlobalQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
