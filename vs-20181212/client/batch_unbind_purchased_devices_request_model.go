// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindPurchasedDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BatchUnbindPurchasedDevicesRequest
	GetDeviceId() *string
	SetOwnerId(v int64) *BatchUnbindPurchasedDevicesRequest
	GetOwnerId() *int64
}

type BatchUnbindPurchasedDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchUnbindPurchasedDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchUnbindPurchasedDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchUnbindPurchasedDevicesRequest) SetDeviceId(v string) *BatchUnbindPurchasedDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesRequest) SetOwnerId(v int64) *BatchUnbindPurchasedDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesRequest) Validate() error {
	return dara.Validate(s)
}
