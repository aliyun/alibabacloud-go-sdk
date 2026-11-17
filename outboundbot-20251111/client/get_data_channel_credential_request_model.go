// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *GetDataChannelCredentialRequest
	GetDeviceId() *string
	SetInstanceId(v string) *GetDataChannelCredentialRequest
	GetInstanceId() *string
}

type GetDataChannelCredentialRequest struct {
	// The client device ID used to apply for data channel credentials. It cannot be empty. It must be consistent with the DeviceId used by the same client for subsequent BridgeWebCall or MakePilotCall operations. The business identifier is provided by the integrator according to the agreement, and is not the ClientId returned by the service or the ChannelId of the voice channel.
	//
	// example:
	//
	// 10001
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The instance identity of the associated product. You must use an instance that the user who invokes the operation has access to and that actually exists. It must be consistent with the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definition.
	//
	// example:
	//
	// instance-example-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDataChannelCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDataChannelCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDataChannelCredentialRequest) SetDeviceId(v string) *GetDataChannelCredentialRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDataChannelCredentialRequest) SetInstanceId(v string) *GetDataChannelCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDataChannelCredentialRequest) Validate() error {
	return dara.Validate(s)
}
