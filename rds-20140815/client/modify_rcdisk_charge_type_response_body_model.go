// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ModifyRCDiskChargeTypeResponseBody
	GetChargeType() *string
	SetExpiredTime(v []*string) *ModifyRCDiskChargeTypeResponseBody
	GetExpiredTime() []*string
	SetInstanceIds(v []*string) *ModifyRCDiskChargeTypeResponseBody
	GetInstanceIds() []*string
	SetOrderId(v string) *ModifyRCDiskChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyRCDiskChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyRCDiskChargeTypeResponseBody struct {
	// example:
	//
	// Prepaid
	ChargeType  *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime []*string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" type:"Repeated"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2133400000****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 6EF82B07-28D2-48D1-B5D6-7E78FED277C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCDiskChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskChargeTypeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyRCDiskChargeTypeResponseBody) GetExpiredTime() []*string {
	return s.ExpiredTime
}

func (s *ModifyRCDiskChargeTypeResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyRCDiskChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyRCDiskChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCDiskChargeTypeResponseBody) SetChargeType(v string) *ModifyRCDiskChargeTypeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *ModifyRCDiskChargeTypeResponseBody) SetExpiredTime(v []*string) *ModifyRCDiskChargeTypeResponseBody {
	s.ExpiredTime = v
	return s
}

func (s *ModifyRCDiskChargeTypeResponseBody) SetInstanceIds(v []*string) *ModifyRCDiskChargeTypeResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ModifyRCDiskChargeTypeResponseBody) SetOrderId(v string) *ModifyRCDiskChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyRCDiskChargeTypeResponseBody) SetRequestId(v string) *ModifyRCDiskChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCDiskChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
