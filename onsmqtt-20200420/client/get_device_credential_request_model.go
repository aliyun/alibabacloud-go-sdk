// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *GetDeviceCredentialRequest
	GetClientId() *string
	SetInstanceId(v string) *GetDeviceCredentialRequest
	GetInstanceId() *string
}

type GetDeviceCredentialRequest struct {
	// The client ID of the device whose access credential you want to query.
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

func (s GetDeviceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetDeviceCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDeviceCredentialRequest) SetClientId(v string) *GetDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *GetDeviceCredentialRequest) SetInstanceId(v string) *GetDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDeviceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
