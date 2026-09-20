// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallbackExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckMessage(v string) *CallbackExtensionRequest
	GetCheckMessage() *string
	SetCheckResult(v string) *CallbackExtensionRequest
	GetCheckResult() *string
	SetExtensionCode(v string) *CallbackExtensionRequest
	GetExtensionCode() *string
	SetMessageId(v string) *CallbackExtensionRequest
	GetMessageId() *string
}

type CallbackExtensionRequest struct {
	// The reason for the failure when CheckResult is set to FAIL.
	//
	// example:
	//
	// The xxx rule is hit. Modify it and try again.
	CheckMessage *string `json:"CheckMessage,omitempty" xml:"CheckMessage,omitempty"`
	// The check status of the extension program for the extension point event. Valid values:
	//
	// - OK: The extension program check for the extension point event passed.
	//
	// - FAIL: The extension program check for the extension point event failed. View and resolve the error promptly to avoid affecting the normal execution of subsequent programs.
	//
	// - WARN: The extension program check for the extension point event passed, but warnings exist.
	//
	// This parameter is required.
	//
	// example:
	//
	// FAIL
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The unique code of the extension program.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2****
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The message ID of the DataWorks open message. After an extension point event is triggered, you can obtain the message ID from the received event message.
	//
	// <props="china">For more information about the message format, see [Message format](https://help.aliyun.com/document_detail/215367.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 034********091
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s CallbackExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s CallbackExtensionRequest) GoString() string {
	return s.String()
}

func (s *CallbackExtensionRequest) GetCheckMessage() *string {
	return s.CheckMessage
}

func (s *CallbackExtensionRequest) GetCheckResult() *string {
	return s.CheckResult
}

func (s *CallbackExtensionRequest) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *CallbackExtensionRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *CallbackExtensionRequest) SetCheckMessage(v string) *CallbackExtensionRequest {
	s.CheckMessage = &v
	return s
}

func (s *CallbackExtensionRequest) SetCheckResult(v string) *CallbackExtensionRequest {
	s.CheckResult = &v
	return s
}

func (s *CallbackExtensionRequest) SetExtensionCode(v string) *CallbackExtensionRequest {
	s.ExtensionCode = &v
	return s
}

func (s *CallbackExtensionRequest) SetMessageId(v string) *CallbackExtensionRequest {
	s.MessageId = &v
	return s
}

func (s *CallbackExtensionRequest) Validate() error {
	return dara.Validate(s)
}
