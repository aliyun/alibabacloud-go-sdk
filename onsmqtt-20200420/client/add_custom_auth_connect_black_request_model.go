// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthConnectBlackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *AddCustomAuthConnectBlackRequest
	GetClientId() *string
	SetInstanceId(v string) *AddCustomAuthConnectBlackRequest
	GetInstanceId() *string
}

type AddCustomAuthConnectBlackRequest struct {
	// The client ID of the device to be blocked from connecting.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the Message Queue for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-i7m26mf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddCustomAuthConnectBlackRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthConnectBlackRequest) GoString() string {
	return s.String()
}

func (s *AddCustomAuthConnectBlackRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AddCustomAuthConnectBlackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddCustomAuthConnectBlackRequest) SetClientId(v string) *AddCustomAuthConnectBlackRequest {
	s.ClientId = &v
	return s
}

func (s *AddCustomAuthConnectBlackRequest) SetInstanceId(v string) *AddCustomAuthConnectBlackRequest {
	s.InstanceId = &v
	return s
}

func (s *AddCustomAuthConnectBlackRequest) Validate() error {
	return dara.Validate(s)
}
