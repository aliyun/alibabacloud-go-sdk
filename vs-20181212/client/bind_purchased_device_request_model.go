// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPurchasedDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BindPurchasedDeviceRequest
	GetDeviceId() *string
	SetGroupId(v string) *BindPurchasedDeviceRequest
	GetGroupId() *string
	SetOwnerId(v int64) *BindPurchasedDeviceRequest
	GetOwnerId() *int64
	SetRegion(v string) *BindPurchasedDeviceRequest
	GetRegion() *string
}

type BindPurchasedDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3939*****6580539-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3484*****8732174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s BindPurchasedDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s BindPurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *BindPurchasedDeviceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BindPurchasedDeviceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BindPurchasedDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindPurchasedDeviceRequest) GetRegion() *string {
	return s.Region
}

func (s *BindPurchasedDeviceRequest) SetDeviceId(v string) *BindPurchasedDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *BindPurchasedDeviceRequest) SetGroupId(v string) *BindPurchasedDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *BindPurchasedDeviceRequest) SetOwnerId(v int64) *BindPurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *BindPurchasedDeviceRequest) SetRegion(v string) *BindPurchasedDeviceRequest {
	s.Region = &v
	return s
}

func (s *BindPurchasedDeviceRequest) Validate() error {
	return dara.Validate(s)
}
