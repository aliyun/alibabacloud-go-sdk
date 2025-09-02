// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkbenchEventResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResult(v string) *UpdateWorkbenchEventResultRequest
	GetCheckResult() *string
	SetCheckResultTip(v string) *UpdateWorkbenchEventResultRequest
	GetCheckResultTip() *string
	SetExtensionCode(v string) *UpdateWorkbenchEventResultRequest
	GetExtensionCode() *string
	SetMessageId(v string) *UpdateWorkbenchEventResultRequest
	GetMessageId() *string
}

type UpdateWorkbenchEventResultRequest struct {
	// The check result of the extension point event. Valid values: OK and Fail.
	//
	// This parameter is required.
	//
	// example:
	//
	// FAIL
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The cause of the check failure.
	//
	// example:
	//
	// SQL is too long
	CheckResultTip *string `json:"CheckResultTip,omitempty" xml:"CheckResultTip,omitempty"`
	// The code of the extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// 58e95e2acd6f408e8707f1bf2591f9e9
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The ID of the message received when the related extension point event is triggered after you enable message subscription by using the OpenEvent module.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03400b03-b721-4c34-8727-2d6884077091
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s UpdateWorkbenchEventResultRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkbenchEventResultRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkbenchEventResultRequest) GetCheckResult() *string {
	return s.CheckResult
}

func (s *UpdateWorkbenchEventResultRequest) GetCheckResultTip() *string {
	return s.CheckResultTip
}

func (s *UpdateWorkbenchEventResultRequest) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *UpdateWorkbenchEventResultRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *UpdateWorkbenchEventResultRequest) SetCheckResult(v string) *UpdateWorkbenchEventResultRequest {
	s.CheckResult = &v
	return s
}

func (s *UpdateWorkbenchEventResultRequest) SetCheckResultTip(v string) *UpdateWorkbenchEventResultRequest {
	s.CheckResultTip = &v
	return s
}

func (s *UpdateWorkbenchEventResultRequest) SetExtensionCode(v string) *UpdateWorkbenchEventResultRequest {
	s.ExtensionCode = &v
	return s
}

func (s *UpdateWorkbenchEventResultRequest) SetMessageId(v string) *UpdateWorkbenchEventResultRequest {
	s.MessageId = &v
	return s
}

func (s *UpdateWorkbenchEventResultRequest) Validate() error {
	return dara.Validate(s)
}
