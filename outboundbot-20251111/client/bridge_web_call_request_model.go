// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeWebCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaller(v string) *BridgeWebCallRequest
	GetCaller() *string
	SetDeviceId(v string) *BridgeWebCallRequest
	GetDeviceId() *string
	SetDraftVersion(v bool) *BridgeWebCallRequest
	GetDraftVersion() *bool
	SetInstanceId(v string) *BridgeWebCallRequest
	GetInstanceId() *string
	SetSampleRate(v int32) *BridgeWebCallRequest
	GetSampleRate() *int32
	SetScriptId(v string) *BridgeWebCallRequest
	GetScriptId() *string
	SetTags(v string) *BridgeWebCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *BridgeWebCallRequest
	GetTimeoutSeconds() *int32
}

type BridgeWebCallRequest struct {
	// The caller information used to display and associate with the web call. This is not equivalent to the actual outbound caller phone number. If not provided or left empty, the current web call implementation uses the caller identifier as the caller display information.
	//
	// example:
	//
	// web-user-001
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// Identifies the client device that initiates the web voice call or receives data channel events. The same client must use the same DeviceId when obtaining data channel credentials and initiating a call. The example is only a business device identifier, not a real credential assigned by the platform.
	//
	// example:
	//
	// 10001
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Specifies the scenario version. A value of true indicates that the draft version is used. A value of false indicates that the published version is used. This is not a scenario publishing operation. Ensure that the selected scenario version exists before calling the operation.
	//
	// example:
	//
	// false
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// The instance identity of the associated product. You must use an existing instance that the invoker has permission to access. The instance must be consistent with the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definition.
	//
	// example:
	//
	// instance-example-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The audio sample rate for the web voice, in Hz. The backend default is 8000. The sample rate must match the client audio and gateway capabilities. The current API code does not define a complete list of allowed values, and other sample rates are not documented as guaranteed supported enumerations.
	//
	// example:
	//
	// 8000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The dialogue scenario ID under the instance. Specify an existing scenario. Subsequent dialogue and end requests for the same text session must use the scenario specified when the session is started.
	//
	// example:
	//
	// script-example-001
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The attached data string passed to the call link in key=value format. Separate multiple key-value pairs with semicolons instead of using JSON. For example, the web call service can read the Callee key to obtain the called party information. The example number is a placeholder value. Specify actual values based on your integration agreements. Do not overwrite internal tags that do not belong to the business custom scope.
	//
	// example:
	//
	// Callee=057100000000
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The timeout for waiting for the call to be answered, in seconds. If the call is not answered within this period, it is actively hung up. The default value on the backend is 60 seconds. This does not represent the maximum duration of an answered call. The current request object does not declare additional upper and lower limits, and unverified value ranges are not added to the documentation.
	//
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s BridgeWebCallRequest) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallRequest) GoString() string {
	return s.String()
}

func (s *BridgeWebCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *BridgeWebCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BridgeWebCallRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *BridgeWebCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BridgeWebCallRequest) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *BridgeWebCallRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *BridgeWebCallRequest) GetTags() *string {
	return s.Tags
}

func (s *BridgeWebCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *BridgeWebCallRequest) SetCaller(v string) *BridgeWebCallRequest {
	s.Caller = &v
	return s
}

func (s *BridgeWebCallRequest) SetDeviceId(v string) *BridgeWebCallRequest {
	s.DeviceId = &v
	return s
}

func (s *BridgeWebCallRequest) SetDraftVersion(v bool) *BridgeWebCallRequest {
	s.DraftVersion = &v
	return s
}

func (s *BridgeWebCallRequest) SetInstanceId(v string) *BridgeWebCallRequest {
	s.InstanceId = &v
	return s
}

func (s *BridgeWebCallRequest) SetSampleRate(v int32) *BridgeWebCallRequest {
	s.SampleRate = &v
	return s
}

func (s *BridgeWebCallRequest) SetScriptId(v string) *BridgeWebCallRequest {
	s.ScriptId = &v
	return s
}

func (s *BridgeWebCallRequest) SetTags(v string) *BridgeWebCallRequest {
	s.Tags = &v
	return s
}

func (s *BridgeWebCallRequest) SetTimeoutSeconds(v int32) *BridgeWebCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *BridgeWebCallRequest) Validate() error {
	return dara.Validate(s)
}
