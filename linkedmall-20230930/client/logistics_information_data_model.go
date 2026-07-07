// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogisticsInformationData interface {
	dara.Model
	String() string
	GoString() string
	SetLogisticsStatus(v string) *LogisticsInformationData
	GetLogisticsStatus() *string
	SetModifiedTime(v string) *LogisticsInformationData
	GetModifiedTime() *string
	SetOrderId(v string) *LogisticsInformationData
	GetOrderId() *string
	SetOrderLineId(v string) *LogisticsInformationData
	GetOrderLineId() *string
	SetOuterPurchaseOrderId(v string) *LogisticsInformationData
	GetOuterPurchaseOrderId() *string
	SetPurchaserId(v string) *LogisticsInformationData
	GetPurchaserId() *string
	SetTrackingCompanyCode(v string) *LogisticsInformationData
	GetTrackingCompanyCode() *string
	SetTrackingCompanyName(v string) *LogisticsInformationData
	GetTrackingCompanyName() *string
	SetTrackingNumber(v string) *LogisticsInformationData
	GetTrackingNumber() *string
}

type LogisticsInformationData struct {
	// The logistics status. A value of 2 indicates that the order is shipped and awaits the buyer\\"s confirmation of receipt.
	//
	// example:
	//
	// 1
	LogisticsStatus *string `json:"logisticsStatus,omitempty" xml:"logisticsStatus,omitempty"`
	// The time when the information was last updated, in yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 1713407100321
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 6696070566****8593
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The ID of the order line.
	//
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// The external ID of the purchase order, provided upon creation.
	//
	// example:
	//
	// 233111
	OuterPurchaseOrderId *string `json:"outerPurchaseOrderId,omitempty" xml:"outerPurchaseOrderId,omitempty"`
	// The ID of the purchaser.
	//
	// example:
	//
	// PID22000009
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// The code of the logistics company.
	//
	// example:
	//
	// SF
	TrackingCompanyCode *string `json:"trackingCompanyCode,omitempty" xml:"trackingCompanyCode,omitempty"`
	// The name of the logistics company.
	//
	// example:
	//
	// 顺丰快递
	TrackingCompanyName *string `json:"trackingCompanyName,omitempty" xml:"trackingCompanyName,omitempty"`
	// The tracking number.
	//
	// example:
	//
	// SF198913131
	TrackingNumber *string `json:"trackingNumber,omitempty" xml:"trackingNumber,omitempty"`
}

func (s LogisticsInformationData) String() string {
	return dara.Prettify(s)
}

func (s LogisticsInformationData) GoString() string {
	return s.String()
}

func (s *LogisticsInformationData) GetLogisticsStatus() *string {
	return s.LogisticsStatus
}

func (s *LogisticsInformationData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *LogisticsInformationData) GetOrderId() *string {
	return s.OrderId
}

func (s *LogisticsInformationData) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *LogisticsInformationData) GetOuterPurchaseOrderId() *string {
	return s.OuterPurchaseOrderId
}

func (s *LogisticsInformationData) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *LogisticsInformationData) GetTrackingCompanyCode() *string {
	return s.TrackingCompanyCode
}

func (s *LogisticsInformationData) GetTrackingCompanyName() *string {
	return s.TrackingCompanyName
}

func (s *LogisticsInformationData) GetTrackingNumber() *string {
	return s.TrackingNumber
}

func (s *LogisticsInformationData) SetLogisticsStatus(v string) *LogisticsInformationData {
	s.LogisticsStatus = &v
	return s
}

func (s *LogisticsInformationData) SetModifiedTime(v string) *LogisticsInformationData {
	s.ModifiedTime = &v
	return s
}

func (s *LogisticsInformationData) SetOrderId(v string) *LogisticsInformationData {
	s.OrderId = &v
	return s
}

func (s *LogisticsInformationData) SetOrderLineId(v string) *LogisticsInformationData {
	s.OrderLineId = &v
	return s
}

func (s *LogisticsInformationData) SetOuterPurchaseOrderId(v string) *LogisticsInformationData {
	s.OuterPurchaseOrderId = &v
	return s
}

func (s *LogisticsInformationData) SetPurchaserId(v string) *LogisticsInformationData {
	s.PurchaserId = &v
	return s
}

func (s *LogisticsInformationData) SetTrackingCompanyCode(v string) *LogisticsInformationData {
	s.TrackingCompanyCode = &v
	return s
}

func (s *LogisticsInformationData) SetTrackingCompanyName(v string) *LogisticsInformationData {
	s.TrackingCompanyName = &v
	return s
}

func (s *LogisticsInformationData) SetTrackingNumber(v string) *LogisticsInformationData {
	s.TrackingNumber = &v
	return s
}

func (s *LogisticsInformationData) Validate() error {
	return dara.Validate(s)
}
