// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceCredential(v *GetDeviceCredentialResponseBodyDeviceCredential) *GetDeviceCredentialResponseBody
	GetDeviceCredential() *GetDeviceCredentialResponseBodyDeviceCredential
	SetRequestId(v string) *GetDeviceCredentialResponseBody
	GetRequestId() *string
}

type GetDeviceCredentialResponseBody struct {
	// The information about the access credential of the device.
	DeviceCredential *GetDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	// The request ID. This parameter is a common parameter.
	//
	// example:
	//
	// E4581CCF-62AF-44D9-B5B4-D1DQDC0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponseBody) GetDeviceCredential() *GetDeviceCredentialResponseBodyDeviceCredential {
	return s.DeviceCredential
}

func (s *GetDeviceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceCredentialResponseBody) SetDeviceCredential(v *GetDeviceCredentialResponseBodyDeviceCredential) *GetDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

func (s *GetDeviceCredentialResponseBody) SetRequestId(v string) *GetDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceCredentialResponseBody) Validate() error {
	if s.DeviceCredential != nil {
		if err := s.DeviceCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceCredentialResponseBodyDeviceCredential struct {
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
	// The timestamp that indicates when the access credential of the device was last updated. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1605541382000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetDeviceCredentialResponseBodyDeviceCredential) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) GetClientId() *string {
	return s.ClientId
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) GetDeviceAccessKeyId() *string {
	return s.DeviceAccessKeyId
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) GetDeviceAccessKeySecret() *string {
	return s.DeviceAccessKeySecret
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) Validate() error {
	return dara.Validate(s)
}
