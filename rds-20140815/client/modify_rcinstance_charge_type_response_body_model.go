// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ModifyRCInstanceChargeTypeResponseBody
	GetChargeType() *string
	SetExpiredTime(v []*string) *ModifyRCInstanceChargeTypeResponseBody
	GetExpiredTime() []*string
	SetFeeOfInstances(v []*ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) *ModifyRCInstanceChargeTypeResponseBody
	GetFeeOfInstances() []*ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances
	SetInstanceIds(v []*string) *ModifyRCInstanceChargeTypeResponseBody
	GetInstanceIds() []*string
	SetOrderId(v string) *ModifyRCInstanceChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyRCInstanceChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyRCInstanceChargeTypeResponseBody struct {
	// example:
	//
	// Prepaid
	ChargeType     *string                                                 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime    []*string                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" type:"Repeated"`
	FeeOfInstances []*ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances `json:"FeeOfInstances,omitempty" xml:"FeeOfInstances,omitempty" type:"Repeated"`
	InstanceIds    []*string                                               `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OrderId        *string                                                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId      *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceChargeTypeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyRCInstanceChargeTypeResponseBody) GetExpiredTime() []*string {
	return s.ExpiredTime
}

func (s *ModifyRCInstanceChargeTypeResponseBody) GetFeeOfInstances() []*ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances {
	return s.FeeOfInstances
}

func (s *ModifyRCInstanceChargeTypeResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyRCInstanceChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyRCInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCInstanceChargeTypeResponseBody) SetChargeType(v string) *ModifyRCInstanceChargeTypeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBody) SetExpiredTime(v []*string) *ModifyRCInstanceChargeTypeResponseBody {
	s.ExpiredTime = v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBody) SetFeeOfInstances(v []*ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) *ModifyRCInstanceChargeTypeResponseBody {
	s.FeeOfInstances = v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBody) SetInstanceIds(v []*string) *ModifyRCInstanceChargeTypeResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBody) SetOrderId(v string) *ModifyRCInstanceChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBody) SetRequestId(v string) *ModifyRCInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBody) Validate() error {
	if s.FeeOfInstances != nil {
		for _, item := range s.FeeOfInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances struct {
	Currency   *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	Fee        *string `json:"Fee,omitempty" xml:"Fee,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) GetCurrency() *string {
	return s.Currency
}

func (s *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) GetFee() *string {
	return s.Fee
}

func (s *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) SetCurrency(v string) *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances {
	s.Currency = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) SetFee(v string) *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances {
	s.Fee = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) SetInstanceId(v string) *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponseBodyFeeOfInstances) Validate() error {
	return dara.Validate(s)
}
