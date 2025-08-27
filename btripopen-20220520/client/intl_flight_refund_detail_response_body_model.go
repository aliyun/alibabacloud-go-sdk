// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightRefundDetailResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightRefundDetailResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightRefundDetailResponseBodyModule) *IntlFlightRefundDetailResponseBody
	GetModule() *IntlFlightRefundDetailResponseBodyModule
	SetRequestId(v string) *IntlFlightRefundDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightRefundDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightRefundDetailResponseBody
	GetTraceId() *string
}

type IntlFlightRefundDetailResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightRefundDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightRefundDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightRefundDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightRefundDetailResponseBody) GetModule() *IntlFlightRefundDetailResponseBodyModule {
	return s.Module
}

func (s *IntlFlightRefundDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightRefundDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightRefundDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightRefundDetailResponseBody) SetCode(v string) *IntlFlightRefundDetailResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBody) SetMessage(v string) *IntlFlightRefundDetailResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBody) SetModule(v *IntlFlightRefundDetailResponseBodyModule) *IntlFlightRefundDetailResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightRefundDetailResponseBody) SetRequestId(v string) *IntlFlightRefundDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBody) SetSuccess(v bool) *IntlFlightRefundDetailResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBody) SetTraceId(v string) *IntlFlightRefundDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModule struct {
	PassengeRefundFeeDetailList []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList `json:"passenge_refund_fee_detail_list,omitempty" xml:"passenge_refund_fee_detail_list,omitempty" type:"Repeated"`
	PassengerList               []*IntlFlightRefundDetailResponseBodyModulePassengerList               `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	RefundOrderInfo             *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo               `json:"refund_order_info,omitempty" xml:"refund_order_info,omitempty" type:"Struct"`
	SegmentList                 []*IntlFlightRefundDetailResponseBodyModuleSegmentList                 `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s IntlFlightRefundDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModule) GetPassengeRefundFeeDetailList() []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList {
	return s.PassengeRefundFeeDetailList
}

func (s *IntlFlightRefundDetailResponseBodyModule) GetPassengerList() []*IntlFlightRefundDetailResponseBodyModulePassengerList {
	return s.PassengerList
}

func (s *IntlFlightRefundDetailResponseBodyModule) GetRefundOrderInfo() *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	return s.RefundOrderInfo
}

func (s *IntlFlightRefundDetailResponseBodyModule) GetSegmentList() []*IntlFlightRefundDetailResponseBodyModuleSegmentList {
	return s.SegmentList
}

func (s *IntlFlightRefundDetailResponseBodyModule) SetPassengeRefundFeeDetailList(v []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) *IntlFlightRefundDetailResponseBodyModule {
	s.PassengeRefundFeeDetailList = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModule) SetPassengerList(v []*IntlFlightRefundDetailResponseBodyModulePassengerList) *IntlFlightRefundDetailResponseBodyModule {
	s.PassengerList = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModule) SetRefundOrderInfo(v *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) *IntlFlightRefundDetailResponseBodyModule {
	s.RefundOrderInfo = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModule) SetSegmentList(v []*IntlFlightRefundDetailResponseBodyModuleSegmentList) *IntlFlightRefundDetailResponseBodyModule {
	s.SegmentList = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList struct {
	// example:
	//
	// 100001
	PassengerId     *int64                                                                              `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	RefundFeeDetail *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail `json:"refund_fee_detail,omitempty" xml:"refund_fee_detail,omitempty" type:"Struct"`
	TicketList      []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList    `json:"ticket_list,omitempty" xml:"ticket_list,omitempty" type:"Repeated"`
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) GetRefundFeeDetail() *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	return s.RefundFeeDetail
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) GetTicketList() []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList {
	return s.TicketList
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) SetPassengerId(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) SetRefundFeeDetail(v *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList {
	s.RefundFeeDetail = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) SetTicketList(v []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList {
	s.TicketList = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail struct {
	// example:
	//
	// 12400
	AlreadyUsedTotalAmount *int64 `json:"already_used_total_amount,omitempty" xml:"already_used_total_amount,omitempty"`
	// example:
	//
	// 21000
	NonRefundableReShopHandlingFee *int64 `json:"non_refundable_re_shop_handling_fee,omitempty" xml:"non_refundable_re_shop_handling_fee,omitempty"`
	// example:
	//
	// 0
	NonRefundableReShopUpgradeFee *int64 `json:"non_refundable_re_shop_upgrade_fee,omitempty" xml:"non_refundable_re_shop_upgrade_fee,omitempty"`
	// example:
	//
	// 0
	NonRefundableTaxDiffFee *int64 `json:"non_refundable_tax_diff_fee,omitempty" xml:"non_refundable_tax_diff_fee,omitempty"`
	// example:
	//
	// 14000
	ReShopRefundAmount *int64 `json:"re_shop_refund_amount,omitempty" xml:"re_shop_refund_amount,omitempty"`
	// example:
	//
	// 14000
	ReShopServiceRefundAmount *int64 `json:"re_shop_service_refund_amount,omitempty" xml:"re_shop_service_refund_amount,omitempty"`
	// example:
	//
	// 0
	ReShopUpgradeRefundAmount *int64                                                                                                         `json:"re_shop_upgrade_refund_amount,omitempty" xml:"re_shop_upgrade_refund_amount,omitempty"`
	RefundReShopFeeDetailList []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList `json:"refund_re_shop_fee_detail_list,omitempty" xml:"refund_re_shop_fee_detail_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RefundTaxDiffAmount *int64 `json:"refund_tax_diff_amount,omitempty" xml:"refund_tax_diff_amount,omitempty"`
	// example:
	//
	// 45000
	RefundTaxFee *int64 `json:"refund_tax_fee,omitempty" xml:"refund_tax_fee,omitempty"`
	// example:
	//
	// 2000
	RefundTicketFee    *int64 `json:"refund_ticket_fee,omitempty" xml:"refund_ticket_fee,omitempty"`
	TaxRefundAmount    *int64 `json:"tax_refund_amount,omitempty" xml:"tax_refund_amount,omitempty"`
	TicketRefundAmount *int64 `json:"ticket_refund_amount,omitempty" xml:"ticket_refund_amount,omitempty"`
	TotalRefundAmount  *int64 `json:"total_refund_amount,omitempty" xml:"total_refund_amount,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetAlreadyUsedTotalAmount() *int64 {
	return s.AlreadyUsedTotalAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetNonRefundableReShopHandlingFee() *int64 {
	return s.NonRefundableReShopHandlingFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetNonRefundableReShopUpgradeFee() *int64 {
	return s.NonRefundableReShopUpgradeFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetNonRefundableTaxDiffFee() *int64 {
	return s.NonRefundableTaxDiffFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetReShopRefundAmount() *int64 {
	return s.ReShopRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetReShopServiceRefundAmount() *int64 {
	return s.ReShopServiceRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetReShopUpgradeRefundAmount() *int64 {
	return s.ReShopUpgradeRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetRefundReShopFeeDetailList() []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	return s.RefundReShopFeeDetailList
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetRefundTaxDiffAmount() *int64 {
	return s.RefundTaxDiffAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetRefundTaxFee() *int64 {
	return s.RefundTaxFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetRefundTicketFee() *int64 {
	return s.RefundTicketFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetTaxRefundAmount() *int64 {
	return s.TaxRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetTicketRefundAmount() *int64 {
	return s.TicketRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) GetTotalRefundAmount() *int64 {
	return s.TotalRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetAlreadyUsedTotalAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.AlreadyUsedTotalAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetNonRefundableReShopHandlingFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.NonRefundableReShopHandlingFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetNonRefundableReShopUpgradeFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.NonRefundableReShopUpgradeFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetNonRefundableTaxDiffFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.NonRefundableTaxDiffFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetReShopRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.ReShopRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetReShopServiceRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.ReShopServiceRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetReShopUpgradeRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.ReShopUpgradeRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetRefundReShopFeeDetailList(v []*IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.RefundReShopFeeDetailList = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetRefundTaxDiffAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.RefundTaxDiffAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetRefundTaxFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.RefundTaxFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetRefundTicketFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.RefundTicketFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetTaxRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.TaxRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetTicketRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.TicketRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) SetTotalRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail {
	s.TotalRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetail) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList struct {
	// example:
	//
	// 21000
	NonRefundableReShopHandlingFee *int64 `json:"non_refundable_re_shop_handling_fee,omitempty" xml:"non_refundable_re_shop_handling_fee,omitempty"`
	// example:
	//
	// 0
	NonRefundableReShopUpgradeFee *int64 `json:"non_refundable_re_shop_upgrade_fee,omitempty" xml:"non_refundable_re_shop_upgrade_fee,omitempty"`
	// example:
	//
	// 0
	NonRefundableTaxDiffFee *int64 `json:"non_refundable_tax_diff_fee,omitempty" xml:"non_refundable_tax_diff_fee,omitempty"`
	// example:
	//
	// 10002340021
	ReShopApplyId *string `json:"re_shop_apply_id,omitempty" xml:"re_shop_apply_id,omitempty"`
	// example:
	//
	// 14000
	ReShopRefundAmount *int64 `json:"re_shop_refund_amount,omitempty" xml:"re_shop_refund_amount,omitempty"`
	// example:
	//
	// 14000
	ReShopServiceRefundAmount *int64 `json:"re_shop_service_refund_amount,omitempty" xml:"re_shop_service_refund_amount,omitempty"`
	// example:
	//
	// 0
	ReShopUpgradeRefundAmount *int64 `json:"re_shop_upgrade_refund_amount,omitempty" xml:"re_shop_upgrade_refund_amount,omitempty"`
	// example:
	//
	// 0
	RefundTaxDiffAmount *int64 `json:"refund_tax_diff_amount,omitempty" xml:"refund_tax_diff_amount,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetNonRefundableReShopHandlingFee() *int64 {
	return s.NonRefundableReShopHandlingFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetNonRefundableReShopUpgradeFee() *int64 {
	return s.NonRefundableReShopUpgradeFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetNonRefundableTaxDiffFee() *int64 {
	return s.NonRefundableTaxDiffFee
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetReShopApplyId() *string {
	return s.ReShopApplyId
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetReShopRefundAmount() *int64 {
	return s.ReShopRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetReShopServiceRefundAmount() *int64 {
	return s.ReShopServiceRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetReShopUpgradeRefundAmount() *int64 {
	return s.ReShopUpgradeRefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) GetRefundTaxDiffAmount() *int64 {
	return s.RefundTaxDiffAmount
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetNonRefundableReShopHandlingFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.NonRefundableReShopHandlingFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetNonRefundableReShopUpgradeFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.NonRefundableReShopUpgradeFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetNonRefundableTaxDiffFee(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.NonRefundableTaxDiffFee = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetReShopApplyId(v string) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.ReShopApplyId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetReShopRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.ReShopRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetReShopServiceRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.ReShopServiceRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetReShopUpgradeRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.ReShopUpgradeRefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) SetRefundTaxDiffAmount(v int64) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList {
	s.RefundTaxDiffAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListRefundFeeDetailRefundReShopFeeDetailList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList struct {
	SegmentKeyList []*string `json:"segment_key_list,omitempty" xml:"segment_key_list,omitempty" type:"Repeated"`
	// example:
	//
	// 784-3553845201
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) GetSegmentKeyList() []*string {
	return s.SegmentKeyList
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) SetSegmentKeyList(v []*string) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList {
	s.SegmentKeyList = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) SetTicketNo(v string) *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList {
	s.TicketNo = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengeRefundFeeDetailListTicketList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModulePassengerList struct {
	// example:
	//
	// 1996-09-13
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// example:
	//
	// 1
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 1001101
	JobNo       *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// example:
	//
	// CN
	NationalityCode *string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// example:
	//
	// 8432002
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// btrip8432002
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModulePassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModulePassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetJobNo() *string {
	return s.JobNo
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetUserId() *string {
	return s.UserId
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) GetUserType() *int32 {
	return s.UserType
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetBirthday(v string) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.Birthday = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetFullName(v string) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetGender(v int32) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.Gender = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetJobNo(v string) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.JobNo = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetNationality(v string) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.Nationality = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetNationalityCode(v string) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.NationalityCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetPassengerId(v int64) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetType(v int32) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.Type = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetUserId(v string) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.UserId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) SetUserType(v int32) *IntlFlightRefundDetailResponseBodyModulePassengerList {
	s.UserType = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModulePassengerList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo struct {
	// example:
	//
	// 2025-06-16 19:20:00
	ApplyTime   *string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	CloseReason *string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// example:
	//
	// 4000
	HandingAmount *int64 `json:"handing_amount,omitempty" xml:"handing_amount,omitempty"`
	// example:
	//
	// 2025011317110900006
	OutRefundApplyId *string `json:"out_refund_apply_id,omitempty" xml:"out_refund_apply_id,omitempty"`
	// example:
	//
	// 0
	ReasonCode *string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	ReasonDesc *string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
	// example:
	//
	// 10200
	RefundAmount *int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// example:
	//
	// 1000000003437017
	RefundApplyId *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	// example:
	//
	// 1000000003437020
	RelationRefundApplyId *int64 `json:"relation_refund_apply_id,omitempty" xml:"relation_refund_apply_id,omitempty"`
	// example:
	//
	// 9
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2025-06-16 20:20:00
	SuccessTime *string `json:"success_time,omitempty" xml:"success_time,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetCloseReason() *string {
	return s.CloseReason
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetHandingAmount() *int64 {
	return s.HandingAmount
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetOutRefundApplyId() *string {
	return s.OutRefundApplyId
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetReasonDesc() *string {
	return s.ReasonDesc
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetRefundAmount() *int64 {
	return s.RefundAmount
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetRelationRefundApplyId() *int64 {
	return s.RelationRefundApplyId
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetStatus() *int32 {
	return s.Status
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetSuccessTime() *string {
	return s.SuccessTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetApplyTime(v string) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.ApplyTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetCloseReason(v string) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.CloseReason = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetHandingAmount(v int64) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.HandingAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetOutRefundApplyId(v string) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.OutRefundApplyId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetReasonCode(v string) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.ReasonCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetReasonDesc(v string) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.ReasonDesc = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetRefundAmount(v int64) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.RefundAmount = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetRefundApplyId(v string) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.RefundApplyId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetRelationRefundApplyId(v int64) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.RelationRefundApplyId = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetStatus(v int32) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.Status = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetSuccessTime(v string) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.SuccessTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) SetVoluntary(v bool) *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo {
	s.Voluntary = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleRefundOrderInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentList struct {
	AirlineInfo    *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:25
	ArrTime        *string                                                            `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepAirportInfo *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// 370100
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 120
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// NS8210
	FlightNo           *string                                                                  `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo      `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                  `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList []*IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 787
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// 0
	JourneyIndex      *int32                                                                `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	LuggageDirectInfo *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	Manufacturer      *string                                                               `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc          *string                                                               `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 1
	OneMore     *int32  `json:"one_more,omitempty" xml:"one_more,omitempty"`
	OneMoreShow *string `json:"one_more_show,omitempty" xml:"one_more_show,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// CZ5009PKXHKG0616
	SegmentKey        *string                                                               `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
	// example:
	//
	// true
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	// example:
	//
	// true
	Stop      *bool   `json:"stop,omitempty" xml:"stop,omitempty"`
	TotalTime *string `json:"total_time,omitempty" xml:"total_time,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetAirlineInfo() *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetArrAirportInfo() *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetDepAirportInfo() *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetFlightShareInfo() *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetFlightStopInfoList() []*IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetLuggageDirectInfo() *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetSegmentVisaRemark() *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetAirlineInfo(v *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetArrAirportInfo(v *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetArrCityCode(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetArrCityName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetArrTime(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetDepAirportInfo(v *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetDepCityCode(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetDepCityName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetDepTime(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.DepTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetDuration(v int32) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.Duration = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetFlightNo(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetFlightShareInfo(v *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetFlightSize(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetFlightStopInfoList(v []*IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetFlightType(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.FlightType = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetJourneyIndex(v int32) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetLuggageDirectInfo(v *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetManufacturer(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetMealDesc(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetOneMore(v int32) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.OneMore = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetOneMoreShow(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetSegmentIndex(v int32) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetSegmentKey(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetSegmentVisaRemark(v *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetShare(v bool) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.Share = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetShortFlightSize(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetStop(v bool) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.Stop = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) SetTotalTime(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentList {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo struct {
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) SetAirlineCode(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) SetAirlineName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) SetShortName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo struct {
	// example:
	//
	// HKG
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) SetAirportCode(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) SetAirportName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) SetAirportShortName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) SetTerminal(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo struct {
	// example:
	//
	// PEK
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T1
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) SetAirportCode(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) SetAirportName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) SetAirportShortName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) SetTerminal(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CA0001
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList struct {
	// example:
	//
	// HGH
	StopAirport     *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportName *string `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	// example:
	//
	// T1
	StopArrTerm *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// HGH
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	// example:
	//
	// T1
	StopDepTerm *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	// example:
	//
	// 2023-08-13 07:45
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// example:
	//
	// 20
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopAirport(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopAirportName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopArrTime(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopCityCode(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopCityName(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopDepTime(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) SetStopTime(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListFlightStopInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo struct {
	// example:
	//
	// 1
	DepCityLuggageDirect *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	// example:
	//
	// 1
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark struct {
	DepCityVisaRemark *string `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	// example:
	//
	// 1
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightRefundDetailResponseBodyModuleSegmentListSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}
