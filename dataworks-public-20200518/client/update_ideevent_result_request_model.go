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
	// The check status of the extension point event. Valid values:
	//
	// 	- OK: The event passes the check.
	//
	// 	- FAIL: The event fails to pass the check. You must check and handle the reported error at the earliest opportunity to ensure that your program is run as expected.
	//
	// 	- WARN: The event passes the check, but an alert is reported.
	//
	// example:
	//
	// OK
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The summary information of the check result. The information is displayed in DataStudio. If an alert is reported or the event fails the check, you can troubleshoot errors based on the information.
	//
	// example:
	//
	// Succeeded
	CheckResultTip *string `json:"CheckResultTip,omitempty" xml:"CheckResultTip,omitempty"`
	// The unique identifier of the extension. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Extensions tab of the Open Platform page to view the identifier.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The message ID. You can obtain the ID from a received message when an extension point event is triggered.
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
