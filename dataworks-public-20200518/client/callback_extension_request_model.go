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
	// The check message of the extension point event. If CheckResult is set to FAIL, you must provide the failure cause.
	//
	// example:
	//
	// The xxx rule is hit. Modify it and try again.
	CheckMessage *string `json:"CheckMessage,omitempty" xml:"CheckMessage,omitempty"`
	// The check status of the extension point event. Valid values:
	//
	// 	- OK: The event passes the check.
	//
	// 	- FAIL: The event fails to pass the check. You must check and handle the reported error at the earliest opportunity to ensure that your program is run as expected.
	//
	// 	- WARN: The event passes the check, but an alert is reported.
	//
	// This parameter is required.
	//
	// example:
	//
	// FAIL
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The unique code of the extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The message ID in DataWorks OpenEvent. You can obtain the ID from a received message when an extension point event is triggered.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03400b03-b721-4c34-8727-2d6884077091
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
