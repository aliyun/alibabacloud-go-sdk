// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnRegisterDeviceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *UnRegisterDeviceCredentialRequest
	GetClientId() *string
	SetInstanceId(v string) *UnRegisterDeviceCredentialRequest
	GetInstanceId() *string
}

type UnRegisterDeviceCredentialRequest struct {
	// The client ID of the device whose access credential you want to deregister.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details*	- page that corresponds to the instance in the ApsaraMQ for MQTT console.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UnRegisterDeviceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UnRegisterDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *UnRegisterDeviceCredentialRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UnRegisterDeviceCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnRegisterDeviceCredentialRequest) SetClientId(v string) *UnRegisterDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *UnRegisterDeviceCredentialRequest) SetInstanceId(v string) *UnRegisterDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *UnRegisterDeviceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
