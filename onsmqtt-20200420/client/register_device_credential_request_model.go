// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *RegisterDeviceCredentialRequest
	GetClientId() *string
	SetInstanceId(v string) *RegisterDeviceCredentialRequest
	GetInstanceId() *string
}

type RegisterDeviceCredentialRequest struct {
	// The client ID of the device for which you want to create an access credential.
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

func (s RegisterDeviceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RegisterDeviceCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RegisterDeviceCredentialRequest) SetClientId(v string) *RegisterDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *RegisterDeviceCredentialRequest) SetInstanceId(v string) *RegisterDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *RegisterDeviceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
