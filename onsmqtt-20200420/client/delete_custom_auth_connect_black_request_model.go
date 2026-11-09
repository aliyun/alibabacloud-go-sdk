// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthConnectBlackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DeleteCustomAuthConnectBlackRequest
	GetClientId() *string
	SetInstanceId(v string) *DeleteCustomAuthConnectBlackRequest
	GetInstanceId() *string
}

type DeleteCustomAuthConnectBlackRequest struct {
	// The client ID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the MSMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteCustomAuthConnectBlackRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthConnectBlackRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthConnectBlackRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteCustomAuthConnectBlackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomAuthConnectBlackRequest) SetClientId(v string) *DeleteCustomAuthConnectBlackRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteCustomAuthConnectBlackRequest) SetInstanceId(v string) *DeleteCustomAuthConnectBlackRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomAuthConnectBlackRequest) Validate() error {
	return dara.Validate(s)
}
