// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchToConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *SwitchToConferenceRequest
	GetDeviceId() *string
	SetInstanceId(v string) *SwitchToConferenceRequest
	GetInstanceId() *string
	SetJobId(v string) *SwitchToConferenceRequest
	GetJobId() *string
	SetUserId(v string) *SwitchToConferenceRequest
	GetUserId() *string
}

type SwitchToConferenceRequest struct {
	// Invalid field. You can enter any value.
	//
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-24114064019637****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Agent ID. If not specified, the agent mapped to the current Resource Access Management (RAM) user is used by default.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SwitchToConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchToConferenceRequest) GoString() string {
	return s.String()
}

func (s *SwitchToConferenceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SwitchToConferenceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchToConferenceRequest) GetJobId() *string {
	return s.JobId
}

func (s *SwitchToConferenceRequest) GetUserId() *string {
	return s.UserId
}

func (s *SwitchToConferenceRequest) SetDeviceId(v string) *SwitchToConferenceRequest {
	s.DeviceId = &v
	return s
}

func (s *SwitchToConferenceRequest) SetInstanceId(v string) *SwitchToConferenceRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchToConferenceRequest) SetJobId(v string) *SwitchToConferenceRequest {
	s.JobId = &v
	return s
}

func (s *SwitchToConferenceRequest) SetUserId(v string) *SwitchToConferenceRequest {
	s.UserId = &v
	return s
}

func (s *SwitchToConferenceRequest) Validate() error {
	return dara.Validate(s)
}
