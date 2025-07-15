// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *GetDataChannelCredentialsRequest
	GetDeviceId() *string
	SetInstanceId(v string) *GetDataChannelCredentialsRequest
	GetInstanceId() *string
}

type GetDataChannelCredentialsRequest struct {
	// This parameter is required.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDataChannelCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialsRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDataChannelCredentialsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDataChannelCredentialsRequest) SetDeviceId(v string) *GetDataChannelCredentialsRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDataChannelCredentialsRequest) SetInstanceId(v string) *GetDataChannelCredentialsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDataChannelCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
