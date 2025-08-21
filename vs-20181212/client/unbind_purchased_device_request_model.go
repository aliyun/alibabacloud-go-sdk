// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPurchasedDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *UnbindPurchasedDeviceRequest
	GetDeviceId() *string
	SetOwnerId(v int64) *UnbindPurchasedDeviceRequest
	GetOwnerId() *int64
}

type UnbindPurchasedDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3939*****6580539-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnbindPurchasedDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindPurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindPurchasedDeviceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UnbindPurchasedDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindPurchasedDeviceRequest) SetDeviceId(v string) *UnbindPurchasedDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindPurchasedDeviceRequest) SetOwnerId(v int64) *UnbindPurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindPurchasedDeviceRequest) Validate() error {
	return dara.Validate(s)
}
