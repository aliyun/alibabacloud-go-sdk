// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDeviceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *RefreshDeviceCredentialRequest
	GetClientId() *string
	SetInstanceId(v string) *RefreshDeviceCredentialRequest
	GetInstanceId() *string
}

type RefreshDeviceCredentialRequest struct {
	// The client ID of the device whose access credential you want to update.
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

func (s RefreshDeviceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RefreshDeviceCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefreshDeviceCredentialRequest) SetClientId(v string) *RefreshDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshDeviceCredentialRequest) SetInstanceId(v string) *RefreshDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *RefreshDeviceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
