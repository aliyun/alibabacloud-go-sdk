// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBack2BackCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalBroker(v string) *StartBack2BackCallRequest
	GetAdditionalBroker() *string
	SetBroker(v string) *StartBack2BackCallRequest
	GetBroker() *string
	SetCallee(v string) *StartBack2BackCallRequest
	GetCallee() *string
	SetCaller(v string) *StartBack2BackCallRequest
	GetCaller() *string
	SetInstanceId(v string) *StartBack2BackCallRequest
	GetInstanceId() *string
	SetTags(v string) *StartBack2BackCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *StartBack2BackCallRequest
	GetTimeoutSeconds() *int32
}

type StartBack2BackCallRequest struct {
	// Additional intermediate number. If this parameter is provided, the intermediate number specified by the Broker parameter is used to call the caller, and the number specified by this parameter is used to call the callee. This parameter is optional and defaults to empty.
	//
	// example:
	//
	// 0102156****
	AdditionalBroker *string `json:"AdditionalBroker,omitempty" xml:"AdditionalBroker,omitempty"`
	// The intermediate number, which must be an active outbound number under the instance. This number is used to sequentially call the caller and the callee in a double-call scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0102157****
	Broker *string `json:"Broker,omitempty" xml:"Broker,omitempty"`
	// The callee number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1372168****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// Caller number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1391814****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Ingest endpoint data. It must not exceed 128 bytes and is primarily used for extension purposes. Ordinary customers do not need to concern themselves with it.
	//
	// example:
	//
	// 无
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The timeout for the dual-call, in seconds. If the call is not answered within the specified time, it will be automatically disconnected. This parameter is optional and defaults to 30 seconds.
	//
	// example:
	//
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s StartBack2BackCallRequest) String() string {
	return dara.Prettify(s)
}

func (s StartBack2BackCallRequest) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallRequest) GetAdditionalBroker() *string {
	return s.AdditionalBroker
}

func (s *StartBack2BackCallRequest) GetBroker() *string {
	return s.Broker
}

func (s *StartBack2BackCallRequest) GetCallee() *string {
	return s.Callee
}

func (s *StartBack2BackCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *StartBack2BackCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartBack2BackCallRequest) GetTags() *string {
	return s.Tags
}

func (s *StartBack2BackCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *StartBack2BackCallRequest) SetAdditionalBroker(v string) *StartBack2BackCallRequest {
	s.AdditionalBroker = &v
	return s
}

func (s *StartBack2BackCallRequest) SetBroker(v string) *StartBack2BackCallRequest {
	s.Broker = &v
	return s
}

func (s *StartBack2BackCallRequest) SetCallee(v string) *StartBack2BackCallRequest {
	s.Callee = &v
	return s
}

func (s *StartBack2BackCallRequest) SetCaller(v string) *StartBack2BackCallRequest {
	s.Caller = &v
	return s
}

func (s *StartBack2BackCallRequest) SetInstanceId(v string) *StartBack2BackCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartBack2BackCallRequest) SetTags(v string) *StartBack2BackCallRequest {
	s.Tags = &v
	return s
}

func (s *StartBack2BackCallRequest) SetTimeoutSeconds(v int32) *StartBack2BackCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *StartBack2BackCallRequest) Validate() error {
	return dara.Validate(s)
}
