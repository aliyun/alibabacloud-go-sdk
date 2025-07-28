// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliveryItemDetailSynRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *DeliveryItemDetailSynRequest
	GetChannel() *string
	SetDeliveryItemDetailDTOS(v []*DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) *DeliveryItemDetailSynRequest
	GetDeliveryItemDetailDTOS() []*DeliveryItemDetailSynRequestDeliveryItemDetailDTOS
	SetDeliveryItemId(v string) *DeliveryItemDetailSynRequest
	GetDeliveryItemId() *string
	SetDeliveryPlanId(v string) *DeliveryItemDetailSynRequest
	GetDeliveryPlanId() *string
}

type DeliveryItemDetailSynRequest struct {
	Channel                *string                                               `json:"channel,omitempty" xml:"channel,omitempty"`
	DeliveryItemDetailDTOS []*DeliveryItemDetailSynRequestDeliveryItemDetailDTOS `json:"deliveryItemDetailDTOS,omitempty" xml:"deliveryItemDetailDTOS,omitempty" type:"Repeated"`
	DeliveryItemId         *string                                               `json:"deliveryItemId,omitempty" xml:"deliveryItemId,omitempty"`
	DeliveryPlanId         *string                                               `json:"deliveryPlanId,omitempty" xml:"deliveryPlanId,omitempty"`
}

func (s DeliveryItemDetailSynRequest) String() string {
	return dara.Prettify(s)
}

func (s DeliveryItemDetailSynRequest) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynRequest) GetChannel() *string {
	return s.Channel
}

func (s *DeliveryItemDetailSynRequest) GetDeliveryItemDetailDTOS() []*DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	return s.DeliveryItemDetailDTOS
}

func (s *DeliveryItemDetailSynRequest) GetDeliveryItemId() *string {
	return s.DeliveryItemId
}

func (s *DeliveryItemDetailSynRequest) GetDeliveryPlanId() *string {
	return s.DeliveryPlanId
}

func (s *DeliveryItemDetailSynRequest) SetChannel(v string) *DeliveryItemDetailSynRequest {
	s.Channel = &v
	return s
}

func (s *DeliveryItemDetailSynRequest) SetDeliveryItemDetailDTOS(v []*DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) *DeliveryItemDetailSynRequest {
	s.DeliveryItemDetailDTOS = v
	return s
}

func (s *DeliveryItemDetailSynRequest) SetDeliveryItemId(v string) *DeliveryItemDetailSynRequest {
	s.DeliveryItemId = &v
	return s
}

func (s *DeliveryItemDetailSynRequest) SetDeliveryPlanId(v string) *DeliveryItemDetailSynRequest {
	s.DeliveryPlanId = &v
	return s
}

func (s *DeliveryItemDetailSynRequest) Validate() error {
	return dara.Validate(s)
}

type DeliveryItemDetailSynRequestDeliveryItemDetailDTOS struct {
	ActualSupplyTime             *string `json:"actualSupplyTime,omitempty" xml:"actualSupplyTime,omitempty"`
	Amount                       *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	Comment                      *string `json:"comment,omitempty" xml:"comment,omitempty"`
	DeliveredAmount              *int64  `json:"deliveredAmount,omitempty" xml:"deliveredAmount,omitempty"`
	DeliveryItemId               *string `json:"deliveryItemId,omitempty" xml:"deliveryItemId,omitempty"`
	DeliveryPlanId               *string `json:"deliveryPlanId,omitempty" xml:"deliveryPlanId,omitempty"`
	LastSupplyTime               *string `json:"lastSupplyTime,omitempty" xml:"lastSupplyTime,omitempty"`
	Status                       *string `json:"status,omitempty" xml:"status,omitempty"`
	SubDemandSupplyPerformerName *string `json:"subDemandSupplyPerformerName,omitempty" xml:"subDemandSupplyPerformerName,omitempty"`
	SubDemandSupplyPerformerUid  *string `json:"subDemandSupplyPerformerUid,omitempty" xml:"subDemandSupplyPerformerUid,omitempty"`
	SubDemandSupplyPmName        *string `json:"subDemandSupplyPmName,omitempty" xml:"subDemandSupplyPmName,omitempty"`
	SubDemandSupplyPmUid         *string `json:"subDemandSupplyPmUid,omitempty" xml:"subDemandSupplyPmUid,omitempty"`
	SubOrderId                   *int64  `json:"subOrderId,omitempty" xml:"subOrderId,omitempty"`
	SupplyStatus                 *string `json:"supplyStatus,omitempty" xml:"supplyStatus,omitempty"`
	TotalOrderId                 *int64  `json:"totalOrderId,omitempty" xml:"totalOrderId,omitempty"`
}

func (s DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) String() string {
	return dara.Prettify(s)
}

func (s DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetActualSupplyTime() *string {
	return s.ActualSupplyTime
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetAmount() *int64 {
	return s.Amount
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetComment() *string {
	return s.Comment
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetDeliveredAmount() *int64 {
	return s.DeliveredAmount
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetDeliveryItemId() *string {
	return s.DeliveryItemId
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetDeliveryPlanId() *string {
	return s.DeliveryPlanId
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetLastSupplyTime() *string {
	return s.LastSupplyTime
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetStatus() *string {
	return s.Status
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetSubDemandSupplyPerformerName() *string {
	return s.SubDemandSupplyPerformerName
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetSubDemandSupplyPerformerUid() *string {
	return s.SubDemandSupplyPerformerUid
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetSubDemandSupplyPmName() *string {
	return s.SubDemandSupplyPmName
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetSubDemandSupplyPmUid() *string {
	return s.SubDemandSupplyPmUid
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetSupplyStatus() *string {
	return s.SupplyStatus
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GetTotalOrderId() *int64 {
	return s.TotalOrderId
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetActualSupplyTime(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.ActualSupplyTime = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetAmount(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.Amount = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetComment(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.Comment = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetDeliveredAmount(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.DeliveredAmount = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetDeliveryItemId(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.DeliveryItemId = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetDeliveryPlanId(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.DeliveryPlanId = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetLastSupplyTime(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.LastSupplyTime = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetStatus(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.Status = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPerformerName(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPerformerName = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPerformerUid(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPerformerUid = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPmName(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPmName = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPmUid(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPmUid = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubOrderId(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubOrderId = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSupplyStatus(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SupplyStatus = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetTotalOrderId(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.TotalOrderId = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) Validate() error {
	return dara.Validate(s)
}
