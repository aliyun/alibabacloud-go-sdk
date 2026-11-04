// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtmpDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRtmpDeviceRequest
	GetDescription() *string
	SetDeviceName(v string) *CreateRtmpDeviceRequest
	GetDeviceName() *string
	SetIotInstanceId(v string) *CreateRtmpDeviceRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *CreateRtmpDeviceRequest
	GetProductKey() *string
	SetPullAuthKey(v string) *CreateRtmpDeviceRequest
	GetPullAuthKey() *string
	SetPullKeyExpireTime(v int32) *CreateRtmpDeviceRequest
	GetPullKeyExpireTime() *int32
	SetPushAuthKey(v string) *CreateRtmpDeviceRequest
	GetPushAuthKey() *string
	SetPushKeyExpireTime(v int32) *CreateRtmpDeviceRequest
	GetPushKeyExpireTime() *int32
	SetSubStreamName(v string) *CreateRtmpDeviceRequest
	GetSubStreamName() *string
}

type CreateRtmpDeviceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName  *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// This parameter is required.
	ProductKey        *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	SubStreamName     *string `json:"SubStreamName,omitempty" xml:"SubStreamName,omitempty"`
}

func (s CreateRtmpDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRtmpDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRtmpDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateRtmpDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreateRtmpDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *CreateRtmpDeviceRequest) GetPullAuthKey() *string {
	return s.PullAuthKey
}

func (s *CreateRtmpDeviceRequest) GetPullKeyExpireTime() *int32 {
	return s.PullKeyExpireTime
}

func (s *CreateRtmpDeviceRequest) GetPushAuthKey() *string {
	return s.PushAuthKey
}

func (s *CreateRtmpDeviceRequest) GetPushKeyExpireTime() *int32 {
	return s.PushKeyExpireTime
}

func (s *CreateRtmpDeviceRequest) GetSubStreamName() *string {
	return s.SubStreamName
}

func (s *CreateRtmpDeviceRequest) SetDescription(v string) *CreateRtmpDeviceRequest {
	s.Description = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetDeviceName(v string) *CreateRtmpDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetIotInstanceId(v string) *CreateRtmpDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetProductKey(v string) *CreateRtmpDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPullAuthKey(v string) *CreateRtmpDeviceRequest {
	s.PullAuthKey = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPullKeyExpireTime(v int32) *CreateRtmpDeviceRequest {
	s.PullKeyExpireTime = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPushAuthKey(v string) *CreateRtmpDeviceRequest {
	s.PushAuthKey = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPushKeyExpireTime(v int32) *CreateRtmpDeviceRequest {
	s.PushKeyExpireTime = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetSubStreamName(v string) *CreateRtmpDeviceRequest {
	s.SubStreamName = &v
	return s
}

func (s *CreateRtmpDeviceRequest) Validate() error {
	return dara.Validate(s)
}
