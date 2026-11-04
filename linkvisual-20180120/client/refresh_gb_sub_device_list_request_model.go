// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshGbSubDeviceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *RefreshGbSubDeviceListRequest
	GetDeviceName() *string
	SetIotId(v string) *RefreshGbSubDeviceListRequest
	GetIotId() *string
	SetIotInstanceId(v string) *RefreshGbSubDeviceListRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *RefreshGbSubDeviceListRequest
	GetProductKey() *string
}

type RefreshGbSubDeviceListRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s RefreshGbSubDeviceListRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshGbSubDeviceListRequest) GoString() string {
	return s.String()
}

func (s *RefreshGbSubDeviceListRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *RefreshGbSubDeviceListRequest) GetIotId() *string {
	return s.IotId
}

func (s *RefreshGbSubDeviceListRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *RefreshGbSubDeviceListRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *RefreshGbSubDeviceListRequest) SetDeviceName(v string) *RefreshGbSubDeviceListRequest {
	s.DeviceName = &v
	return s
}

func (s *RefreshGbSubDeviceListRequest) SetIotId(v string) *RefreshGbSubDeviceListRequest {
	s.IotId = &v
	return s
}

func (s *RefreshGbSubDeviceListRequest) SetIotInstanceId(v string) *RefreshGbSubDeviceListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *RefreshGbSubDeviceListRequest) SetProductKey(v string) *RefreshGbSubDeviceListRequest {
	s.ProductKey = &v
	return s
}

func (s *RefreshGbSubDeviceListRequest) Validate() error {
	return dara.Validate(s)
}
