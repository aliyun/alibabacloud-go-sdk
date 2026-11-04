// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtmpKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *UpdateRtmpKeyRequest
	GetDeviceName() *string
	SetIotId(v string) *UpdateRtmpKeyRequest
	GetIotId() *string
	SetIotInstanceId(v string) *UpdateRtmpKeyRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *UpdateRtmpKeyRequest
	GetProductKey() *string
	SetPullAuthKey(v string) *UpdateRtmpKeyRequest
	GetPullAuthKey() *string
	SetPullKeyExpireTime(v int32) *UpdateRtmpKeyRequest
	GetPullKeyExpireTime() *int32
	SetPushAuthKey(v string) *UpdateRtmpKeyRequest
	GetPushAuthKey() *string
	SetPushKeyExpireTime(v int32) *UpdateRtmpKeyRequest
	GetPushKeyExpireTime() *int32
}

type UpdateRtmpKeyRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId     *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey        *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
}

func (s UpdateRtmpKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtmpKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtmpKeyRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *UpdateRtmpKeyRequest) GetIotId() *string {
	return s.IotId
}

func (s *UpdateRtmpKeyRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *UpdateRtmpKeyRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *UpdateRtmpKeyRequest) GetPullAuthKey() *string {
	return s.PullAuthKey
}

func (s *UpdateRtmpKeyRequest) GetPullKeyExpireTime() *int32 {
	return s.PullKeyExpireTime
}

func (s *UpdateRtmpKeyRequest) GetPushAuthKey() *string {
	return s.PushAuthKey
}

func (s *UpdateRtmpKeyRequest) GetPushKeyExpireTime() *int32 {
	return s.PushKeyExpireTime
}

func (s *UpdateRtmpKeyRequest) SetDeviceName(v string) *UpdateRtmpKeyRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetIotId(v string) *UpdateRtmpKeyRequest {
	s.IotId = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetIotInstanceId(v string) *UpdateRtmpKeyRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetProductKey(v string) *UpdateRtmpKeyRequest {
	s.ProductKey = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPullAuthKey(v string) *UpdateRtmpKeyRequest {
	s.PullAuthKey = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPullKeyExpireTime(v int32) *UpdateRtmpKeyRequest {
	s.PullKeyExpireTime = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPushAuthKey(v string) *UpdateRtmpKeyRequest {
	s.PushAuthKey = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPushKeyExpireTime(v int32) *UpdateRtmpKeyRequest {
	s.PushKeyExpireTime = &v
	return s
}

func (s *UpdateRtmpKeyRequest) Validate() error {
	return dara.Validate(s)
}
