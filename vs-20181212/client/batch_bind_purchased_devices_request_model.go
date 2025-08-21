// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindPurchasedDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BatchBindPurchasedDevicesRequest
	GetDeviceId() *string
	SetGroupId(v string) *BatchBindPurchasedDevicesRequest
	GetGroupId() *string
	SetOwnerId(v int64) *BatchBindPurchasedDevicesRequest
	GetOwnerId() *int64
	SetRegion(v string) *BatchBindPurchasedDevicesRequest
	GetRegion() *string
}

type BatchBindPurchasedDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s BatchBindPurchasedDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchBindPurchasedDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchBindPurchasedDevicesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchBindPurchasedDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchBindPurchasedDevicesRequest) GetRegion() *string {
	return s.Region
}

func (s *BatchBindPurchasedDevicesRequest) SetDeviceId(v string) *BatchBindPurchasedDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchBindPurchasedDevicesRequest) SetGroupId(v string) *BatchBindPurchasedDevicesRequest {
	s.GroupId = &v
	return s
}

func (s *BatchBindPurchasedDevicesRequest) SetOwnerId(v int64) *BatchBindPurchasedDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindPurchasedDevicesRequest) SetRegion(v string) *BatchBindPurchasedDevicesRequest {
	s.Region = &v
	return s
}

func (s *BatchBindPurchasedDevicesRequest) Validate() error {
	return dara.Validate(s)
}
