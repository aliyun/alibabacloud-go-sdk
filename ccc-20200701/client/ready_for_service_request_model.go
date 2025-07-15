// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadyForServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *ReadyForServiceRequest
	GetDeviceId() *string
	SetInstanceId(v string) *ReadyForServiceRequest
	GetInstanceId() *string
	SetOutboundScenario(v bool) *ReadyForServiceRequest
	GetOutboundScenario() *bool
	SetUserId(v string) *ReadyForServiceRequest
	GetUserId() *string
}

type ReadyForServiceRequest struct {
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// user-test@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ReadyForServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadyForServiceRequest) GoString() string {
	return s.String()
}

func (s *ReadyForServiceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ReadyForServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReadyForServiceRequest) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ReadyForServiceRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReadyForServiceRequest) SetDeviceId(v string) *ReadyForServiceRequest {
	s.DeviceId = &v
	return s
}

func (s *ReadyForServiceRequest) SetInstanceId(v string) *ReadyForServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReadyForServiceRequest) SetOutboundScenario(v bool) *ReadyForServiceRequest {
	s.OutboundScenario = &v
	return s
}

func (s *ReadyForServiceRequest) SetUserId(v string) *ReadyForServiceRequest {
	s.UserId = &v
	return s
}

func (s *ReadyForServiceRequest) Validate() error {
	return dara.Validate(s)
}
