// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribePurchasedDeviceRequest
	GetId() *string
	SetOwnerId(v int64) *DescribePurchasedDeviceRequest
	GetOwnerId() *int64
}

type DescribePurchasedDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3614*****66212-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribePurchasedDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDeviceRequest) GetId() *string {
	return s.Id
}

func (s *DescribePurchasedDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePurchasedDeviceRequest) SetId(v string) *DescribePurchasedDeviceRequest {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDeviceRequest) SetOwnerId(v int64) *DescribePurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePurchasedDeviceRequest) Validate() error {
	return dara.Validate(s)
}
