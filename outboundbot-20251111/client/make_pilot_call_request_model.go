// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakePilotCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallVariables(v string) *MakePilotCallRequest
	GetCallVariables() *string
	SetDeviceId(v string) *MakePilotCallRequest
	GetDeviceId() *string
	SetDraftVersion(v bool) *MakePilotCallRequest
	GetDraftVersion() *bool
	SetInstanceId(v string) *MakePilotCallRequest
	GetInstanceId() *string
	SetNumber(v string) *MakePilotCallRequest
	GetNumber() *string
	SetScriptId(v string) *MakePilotCallRequest
	GetScriptId() *string
	SetTags(v string) *MakePilotCallRequest
	GetTags() *string
	SetVerificationCode(v string) *MakePilotCallRequest
	GetVerificationCode() *string
}

type MakePilotCallRequest struct {
	// The variables passed to the phone call link. The value is a JSON object string, and variable values are processed as strings. Specific variable names are agreed upon by the call link and business scenario. This parameter is not semicolon-separated Tags. The example uses an empty object when no custom variables are defined.
	//
	// example:
	//
	// {}
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// The ID of the client device that initiates the web voice or receives data channel events. The same client must use the same DeviceId when obtaining data channel credentials and initiating a call. The example is only a business device ID and not a real credential assigned by the platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Specifies whether to use the draft version. Valid values:
	//
	// - true: Uses the draft version.
	//
	// - false: Uses the published version.
	//
	// This parameter does not publish the scenario. Ensure that the selected scenario version exists before you call this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// The instance identity of the product. You must use an instance that the invoker has access to and that actually exists. The instance must be consistent with the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-example-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The called phone number for the current real-number outbound call test. The number must be the same as the one used to obtain the SMS verification code. Specify a number that is actually supported by the business and authorized. The current request object only requires this parameter to be non-empty. This document does not fabricate number formats or international number support scopes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13800000000
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The dialog scenario ID of the instance. Specify an existing scenario. Subsequent dialog and end requests for the same text session must use the scenario specified when the session is started.
	//
	// This parameter is required.
	//
	// example:
	//
	// script-example-001
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The attached data string passed to the call link. Use the key=value format. Separate multiple key-value pairs with semicolons. Do not use JSON. For example, the Callee key can be used by the web call service to read the called party information. The example number is a placeholder value. Specify the actual value based on the access conventions of your business. Do not overwrite internal tags that do not belong to the business custom scope.
	//
	// example:
	//
	// Callee=057100000000
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The SMS verification code sent to the number specified by Number for the verification of the OutboundCallTest scenario. The request must use a valid verification code actually received by the number. The example only demonstrates a string parameter and cannot be used to pass the verification. It also does not declare undefined length rules for the code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s MakePilotCallRequest) String() string {
	return dara.Prettify(s)
}

func (s MakePilotCallRequest) GoString() string {
	return s.String()
}

func (s *MakePilotCallRequest) GetCallVariables() *string {
	return s.CallVariables
}

func (s *MakePilotCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MakePilotCallRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *MakePilotCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MakePilotCallRequest) GetNumber() *string {
	return s.Number
}

func (s *MakePilotCallRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *MakePilotCallRequest) GetTags() *string {
	return s.Tags
}

func (s *MakePilotCallRequest) GetVerificationCode() *string {
	return s.VerificationCode
}

func (s *MakePilotCallRequest) SetCallVariables(v string) *MakePilotCallRequest {
	s.CallVariables = &v
	return s
}

func (s *MakePilotCallRequest) SetDeviceId(v string) *MakePilotCallRequest {
	s.DeviceId = &v
	return s
}

func (s *MakePilotCallRequest) SetDraftVersion(v bool) *MakePilotCallRequest {
	s.DraftVersion = &v
	return s
}

func (s *MakePilotCallRequest) SetInstanceId(v string) *MakePilotCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MakePilotCallRequest) SetNumber(v string) *MakePilotCallRequest {
	s.Number = &v
	return s
}

func (s *MakePilotCallRequest) SetScriptId(v string) *MakePilotCallRequest {
	s.ScriptId = &v
	return s
}

func (s *MakePilotCallRequest) SetTags(v string) *MakePilotCallRequest {
	s.Tags = &v
	return s
}

func (s *MakePilotCallRequest) SetVerificationCode(v string) *MakePilotCallRequest {
	s.VerificationCode = &v
	return s
}

func (s *MakePilotCallRequest) Validate() error {
	return dara.Validate(s)
}
