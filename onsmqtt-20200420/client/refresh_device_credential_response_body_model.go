// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDeviceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceCredential(v *RefreshDeviceCredentialResponseBodyDeviceCredential) *RefreshDeviceCredentialResponseBody
	GetDeviceCredential() *RefreshDeviceCredentialResponseBodyDeviceCredential
	SetRequestId(v string) *RefreshDeviceCredentialResponseBody
	GetRequestId() *string
}

type RefreshDeviceCredentialResponseBody struct {
	// The access credential of the device.
	DeviceCredential *RefreshDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	// The request ID. This parameter is a common parameter.
	//
	// example:
	//
	// E4581CCF-62AF-44D9-B5B4-D1DBDC0F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshDeviceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponseBody) GetDeviceCredential() *RefreshDeviceCredentialResponseBodyDeviceCredential {
	return s.DeviceCredential
}

func (s *RefreshDeviceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshDeviceCredentialResponseBody) SetDeviceCredential(v *RefreshDeviceCredentialResponseBodyDeviceCredential) *RefreshDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

func (s *RefreshDeviceCredentialResponseBody) SetRequestId(v string) *RefreshDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBody) Validate() error {
	if s.DeviceCredential != nil {
		if err := s.DeviceCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefreshDeviceCredentialResponseBodyDeviceCredential struct {
	// The client ID of the device.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The timestamp that indicates when the access credential of the device was created. The value of this parameter is a UNIX timestamp in milliseconds.
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

func (s RefreshDeviceCredentialResponseBodyDeviceCredential) String() string {
	return dara.Prettify(s)
}

func (s RefreshDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) GetClientId() *string {
	return s.ClientId
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) GetDeviceAccessKeyId() *string {
	return s.DeviceAccessKeyId
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) GetDeviceAccessKeySecret() *string {
	return s.DeviceAccessKeySecret
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) Validate() error {
	return dara.Validate(s)
}
