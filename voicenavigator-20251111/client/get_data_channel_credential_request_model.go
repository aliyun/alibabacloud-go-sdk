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
	// example:
	//
	// d4c38420-****-43bc-b001-56bfc274****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
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
