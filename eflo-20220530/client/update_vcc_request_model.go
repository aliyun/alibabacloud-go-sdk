// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *UpdateVccRequest
	GetBandwidth() *int32
	SetOrderId(v string) *UpdateVccRequest
	GetOrderId() *string
	SetRegionId(v string) *UpdateVccRequest
	GetRegionId() *string
	SetVccId(v string) *UpdateVccRequest
	GetVccId() *string
	SetVccName(v string) *UpdateVccRequest
	GetVccName() *string
}

type UpdateVccRequest struct {
	// The peak bandwidth of the Lingjun connection instance. Unit: Mbit/s. Valid values: 1000 to 400000
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the order placed on the instance.
	//
	// example:
	//
	// 20006627643
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The name of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-heyuan-backup
	VccName *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
}

func (s UpdateVccRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVccRequest) GoString() string {
	return s.String()
}

func (s *UpdateVccRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateVccRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateVccRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVccRequest) GetVccId() *string {
	return s.VccId
}

func (s *UpdateVccRequest) GetVccName() *string {
	return s.VccName
}

func (s *UpdateVccRequest) SetBandwidth(v int32) *UpdateVccRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateVccRequest) SetOrderId(v string) *UpdateVccRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateVccRequest) SetRegionId(v string) *UpdateVccRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVccRequest) SetVccId(v string) *UpdateVccRequest {
	s.VccId = &v
	return s
}

func (s *UpdateVccRequest) SetVccName(v string) *UpdateVccRequest {
	s.VccName = &v
	return s
}

func (s *UpdateVccRequest) Validate() error {
	return dara.Validate(s)
}
