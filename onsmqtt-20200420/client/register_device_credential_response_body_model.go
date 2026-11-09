// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceCredential(v *RegisterDeviceCredentialResponseBodyDeviceCredential) *RegisterDeviceCredentialResponseBody
	GetDeviceCredential() *RegisterDeviceCredentialResponseBodyDeviceCredential
	SetRequestId(v string) *RegisterDeviceCredentialResponseBody
	GetRequestId() *string
}

type RegisterDeviceCredentialResponseBody struct {
	// The access credential of the device.
	DeviceCredential *RegisterDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	// The request ID. This parameter is a common parameter.
	//
	// example:
	//
	// E4581CCF-62AF-44D9-B5B4-D1DBDC0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDeviceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponseBody) GetDeviceCredential() *RegisterDeviceCredentialResponseBodyDeviceCredential {
	return s.DeviceCredential
}

func (s *RegisterDeviceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterDeviceCredentialResponseBody) SetDeviceCredential(v *RegisterDeviceCredentialResponseBodyDeviceCredential) *RegisterDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

func (s *RegisterDeviceCredentialResponseBody) SetRequestId(v string) *RegisterDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBody) Validate() error {
	if s.DeviceCredential != nil {
		if err := s.DeviceCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RegisterDeviceCredentialResponseBodyDeviceCredential struct {
	// The client ID of the device.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The timestamp that indicates when the access credential of the device was created. Unit: milliseconds.
	//
	// example:
	//
	// 1605541382000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The AccessKey ID of the device.
	//
	// example:
	//
	// DC.Z5fXh9sRRVufyLi6wo****
	DeviceAccessKeyId *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	// The AccessKey secret of the device.
	//
	// example:
	//
	// DC.BJMkn4eMQJK2vaApTS****
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp that indicates when the access credential of the device was last updated. Unit: milliseconds.
	//
	// example:
	//
	// 1605541382000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s RegisterDeviceCredentialResponseBodyDeviceCredential) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) GetClientId() *string {
	return s.ClientId
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) GetDeviceAccessKeyId() *string {
	return s.DeviceAccessKeyId
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) GetDeviceAccessKeySecret() *string {
	return s.DeviceAccessKeySecret
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) Validate() error {
	return dara.Validate(s)
}
