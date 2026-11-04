// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRtmpKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryRtmpKeyRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryRtmpKeyRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRtmpKeyRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryRtmpKeyRequest
	GetProductKey() *string
}

type QueryRtmpKeyRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryRtmpKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRtmpKeyRequest) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRtmpKeyRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRtmpKeyRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRtmpKeyRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRtmpKeyRequest) SetDeviceName(v string) *QueryRtmpKeyRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRtmpKeyRequest) SetIotId(v string) *QueryRtmpKeyRequest {
	s.IotId = &v
	return s
}

func (s *QueryRtmpKeyRequest) SetIotInstanceId(v string) *QueryRtmpKeyRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRtmpKeyRequest) SetProductKey(v string) *QueryRtmpKeyRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRtmpKeyRequest) Validate() error {
	return dara.Validate(s)
}
