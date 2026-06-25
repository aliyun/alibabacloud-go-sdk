// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCloudPhoneNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewCloudPhoneNodesRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewCloudPhoneNodesRequest
	GetAutoRenew() *bool
	SetNodeIds(v []*string) *RenewCloudPhoneNodesRequest
	GetNodeIds() []*string
	SetPaidCallBackUrl(v string) *RenewCloudPhoneNodesRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *RenewCloudPhoneNodesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewCloudPhoneNodesRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewCloudPhoneNodesRequest
	GetPromotionId() *string
}

type RenewCloudPhoneNodesRequest struct {
	// Specifies whether to enable automatic payment. The default value is false.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// A list of cloud phone matrix IDs.
	NodeIds         []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	PaidCallBackUrl *string   `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The renewal duration. The `PeriodUnit` parameter specifies the unit.
	//
	// - If `PeriodUnit` is **Year**, the value must be 1.
	//
	// - If `PeriodUnit` is **Month**, the valid values are 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s RenewCloudPhoneNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *RenewCloudPhoneNodesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewCloudPhoneNodesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewCloudPhoneNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *RenewCloudPhoneNodesRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *RenewCloudPhoneNodesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewCloudPhoneNodesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewCloudPhoneNodesRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewCloudPhoneNodesRequest) SetAutoPay(v bool) *RenewCloudPhoneNodesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetAutoRenew(v bool) *RenewCloudPhoneNodesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetNodeIds(v []*string) *RenewCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPaidCallBackUrl(v string) *RenewCloudPhoneNodesRequest {
	s.PaidCallBackUrl = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPeriod(v int32) *RenewCloudPhoneNodesRequest {
	s.Period = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPeriodUnit(v string) *RenewCloudPhoneNodesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPromotionId(v string) *RenewCloudPhoneNodesRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) Validate() error {
	return dara.Validate(s)
}
