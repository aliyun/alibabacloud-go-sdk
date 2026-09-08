// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgentStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *ResetAgentStateRequest
	GetDeviceId() *string
	SetInstanceId(v string) *ResetAgentStateRequest
	GetInstanceId() *string
	SetUserId(v string) *ResetAgentStateRequest
	GetUserId() *string
}

type ResetAgentStateRequest struct {
	// A string that identifies the device. The value is not processed by the system and can be any string.
	//
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The ID of the Cloud Call Center (CCC) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the agent. This parameter is optional. If omitted, the agent mapped to the current RAM account is reset.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ResetAgentStateRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAgentStateRequest) GoString() string {
	return s.String()
}

func (s *ResetAgentStateRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ResetAgentStateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetAgentStateRequest) GetUserId() *string {
	return s.UserId
}

func (s *ResetAgentStateRequest) SetDeviceId(v string) *ResetAgentStateRequest {
	s.DeviceId = &v
	return s
}

func (s *ResetAgentStateRequest) SetInstanceId(v string) *ResetAgentStateRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetAgentStateRequest) SetUserId(v string) *ResetAgentStateRequest {
	s.UserId = &v
	return s
}

func (s *ResetAgentStateRequest) Validate() error {
	return dara.Validate(s)
}
