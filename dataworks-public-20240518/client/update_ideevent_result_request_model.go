// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIDEEventResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResult(v string) *UpdateIDEEventResultRequest
	GetCheckResult() *string
	SetCheckResultTip(v string) *UpdateIDEEventResultRequest
	GetCheckResultTip() *string
	SetExtensionCode(v string) *UpdateIDEEventResultRequest
	GetExtensionCode() *string
	SetMessageId(v string) *UpdateIDEEventResultRequest
	GetMessageId() *string
}

type UpdateIDEEventResultRequest struct {
	// The check status of the extension for this extension point event. Valid values:
	//
	// 	- OK: The extension passed the check for this event.
	//
	// 	- FAIL: The extension failed the check for this event. You need to review and resolve the error promptly to avoid affecting subsequent program execution.
	//
	// 	- WARN: The extension passed the check for this event, but with warnings.
	//
	// example:
	//
	// OK
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// A summary of the check result for this extension point event. This message is displayed on your current development page. When the check fails or has warnings, you can use this summary to quickly identify the cause.
	//
	// example:
	//
	// Succeeded
	CheckResultTip *string `json:"CheckResultTip,omitempty" xml:"CheckResultTip,omitempty"`
	// The unique identifier of the extension. You can obtain the identifier from the Extensions tab on Open Platform in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The OpenEvent message ID from DataWorks. When an extension point event is triggered, you can obtain the message ID from the event message.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s UpdateIDEEventResultRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIDEEventResultRequest) GoString() string {
	return s.String()
}

func (s *UpdateIDEEventResultRequest) GetCheckResult() *string {
	return s.CheckResult
}

func (s *UpdateIDEEventResultRequest) GetCheckResultTip() *string {
	return s.CheckResultTip
}

func (s *UpdateIDEEventResultRequest) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *UpdateIDEEventResultRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *UpdateIDEEventResultRequest) SetCheckResult(v string) *UpdateIDEEventResultRequest {
	s.CheckResult = &v
	return s
}

func (s *UpdateIDEEventResultRequest) SetCheckResultTip(v string) *UpdateIDEEventResultRequest {
	s.CheckResultTip = &v
	return s
}

func (s *UpdateIDEEventResultRequest) SetExtensionCode(v string) *UpdateIDEEventResultRequest {
	s.ExtensionCode = &v
	return s
}

func (s *UpdateIDEEventResultRequest) SetMessageId(v string) *UpdateIDEEventResultRequest {
	s.MessageId = &v
	return s
}

func (s *UpdateIDEEventResultRequest) Validate() error {
	return dara.Validate(s)
}
