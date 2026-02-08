// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDialogueFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDialogueFlowResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDialogueFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDialogueFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDialogueFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDialogueFlowResponseBody
	GetSuccess() *bool
}

type DeleteDialogueFlowResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
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
	// Indicates whether the invocation succeeded.
	//
	// - **true**: The invocation succeeded.
	//
	// - **false**: Failed to invoke.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDialogueFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDialogueFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDialogueFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDialogueFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDialogueFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDialogueFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDialogueFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDialogueFlowResponseBody) SetCode(v string) *DeleteDialogueFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDialogueFlowResponseBody) SetHttpStatusCode(v int32) *DeleteDialogueFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDialogueFlowResponseBody) SetMessage(v string) *DeleteDialogueFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDialogueFlowResponseBody) SetRequestId(v string) *DeleteDialogueFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDialogueFlowResponseBody) SetSuccess(v bool) *DeleteDialogueFlowResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDialogueFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
