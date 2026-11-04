// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtmpKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteRtmpKeyRequest
	GetDeviceName() *string
	SetIotId(v string) *DeleteRtmpKeyRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteRtmpKeyRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *DeleteRtmpKeyRequest
	GetProductKey() *string
	SetType(v int32) *DeleteRtmpKeyRequest
	GetType() *int32
}

type DeleteRtmpKeyRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// This parameter is required.
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteRtmpKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtmpKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteRtmpKeyRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteRtmpKeyRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteRtmpKeyRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteRtmpKeyRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteRtmpKeyRequest) GetType() *int32 {
	return s.Type
}

func (s *DeleteRtmpKeyRequest) SetDeviceName(v string) *DeleteRtmpKeyRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetIotId(v string) *DeleteRtmpKeyRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetIotInstanceId(v string) *DeleteRtmpKeyRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetProductKey(v string) *DeleteRtmpKeyRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetType(v int32) *DeleteRtmpKeyRequest {
	s.Type = &v
	return s
}

func (s *DeleteRtmpKeyRequest) Validate() error {
	return dara.Validate(s)
}
