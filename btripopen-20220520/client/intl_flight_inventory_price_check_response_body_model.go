// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightInventoryPriceCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightInventoryPriceCheckResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightInventoryPriceCheckResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightInventoryPriceCheckResponseBodyModule) *IntlFlightInventoryPriceCheckResponseBody
	GetModule() *IntlFlightInventoryPriceCheckResponseBodyModule
	SetRequestId(v string) *IntlFlightInventoryPriceCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightInventoryPriceCheckResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightInventoryPriceCheckResponseBody
	GetTraceId() *string
}

type IntlFlightInventoryPriceCheckResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                          `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightInventoryPriceCheckResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 92359A00-85D8-16C4-AED0-249618DEEC17
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2103ad1516839612078738332dea5c
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightInventoryPriceCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightInventoryPriceCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightInventoryPriceCheckResponseBody) GetModule() *IntlFlightInventoryPriceCheckResponseBodyModule {
	return s.Module
}

func (s *IntlFlightInventoryPriceCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightInventoryPriceCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightInventoryPriceCheckResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightInventoryPriceCheckResponseBody) SetCode(v string) *IntlFlightInventoryPriceCheckResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBody) SetMessage(v string) *IntlFlightInventoryPriceCheckResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBody) SetModule(v *IntlFlightInventoryPriceCheckResponseBodyModule) *IntlFlightInventoryPriceCheckResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBody) SetRequestId(v string) *IntlFlightInventoryPriceCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBody) SetSuccess(v bool) *IntlFlightInventoryPriceCheckResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBody) SetTraceId(v string) *IntlFlightInventoryPriceCheckResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightInventoryPriceCheckResponseBodyModule struct {
	// example:
	//
	// true
	CheckSuccess *bool `json:"check_success,omitempty" xml:"check_success,omitempty"`
	// example:
	//
	// 0
	FailType                      *int32                                                                          `json:"fail_type,omitempty" xml:"fail_type,omitempty"`
	PassengerChangedPriceInfoList []*IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList `json:"passenger_changed_price_info_list,omitempty" xml:"passenger_changed_price_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// fcoid_deb6372db8194f1c94c23bc4fadc508d
	RenderKey *string `json:"render_key,omitempty" xml:"render_key,omitempty"`
}

func (s IntlFlightInventoryPriceCheckResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) GetCheckSuccess() *bool {
	return s.CheckSuccess
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) GetFailType() *int32 {
	return s.FailType
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) GetPassengerChangedPriceInfoList() []*IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList {
	return s.PassengerChangedPriceInfoList
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) GetRenderKey() *string {
	return s.RenderKey
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) SetCheckSuccess(v bool) *IntlFlightInventoryPriceCheckResponseBodyModule {
	s.CheckSuccess = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) SetFailType(v int32) *IntlFlightInventoryPriceCheckResponseBodyModule {
	s.FailType = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) SetPassengerChangedPriceInfoList(v []*IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) *IntlFlightInventoryPriceCheckResponseBodyModule {
	s.PassengerChangedPriceInfoList = v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) SetRenderKey(v string) *IntlFlightInventoryPriceCheckResponseBodyModule {
	s.RenderKey = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList struct {
	// example:
	//
	// false
	Changed       *bool                                                                                      `json:"changed,omitempty" xml:"changed,omitempty"`
	ChangedPrice  *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice  `json:"changed_price,omitempty" xml:"changed_price,omitempty" type:"Struct"`
	OriginalPrice *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice `json:"original_price,omitempty" xml:"original_price,omitempty" type:"Struct"`
	// example:
	//
	// 0
	PassengerType *int32 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}

func (s IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) GetChanged() *bool {
	return s.Changed
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) GetChangedPrice() *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice {
	return s.ChangedPrice
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) GetOriginalPrice() *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice {
	return s.OriginalPrice
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) SetChanged(v bool) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList {
	s.Changed = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) SetChangedPrice(v *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList {
	s.ChangedPrice = v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) SetOriginalPrice(v *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList {
	s.OriginalPrice = v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) SetPassengerType(v int32) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList {
	s.PassengerType = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice struct {
	// example:
	//
	// 12000
	TaxPrice *int64 `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
	// example:
	//
	// 90000
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

func (s IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) GetTaxPrice() *int64 {
	return s.TaxPrice
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) SetTaxPrice(v int64) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice {
	s.TaxPrice = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) SetTicketPrice(v int64) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice {
	s.TicketPrice = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListChangedPrice) Validate() error {
	return dara.Validate(s)
}

type IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice struct {
	// example:
	//
	// 12000
	TaxPrice *int64 `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
	// example:
	//
	// 80000
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

func (s IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) GetTaxPrice() *int64 {
	return s.TaxPrice
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) SetTaxPrice(v int64) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice {
	s.TaxPrice = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) SetTicketPrice(v int64) *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice {
	s.TicketPrice = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponseBodyModulePassengerChangedPriceInfoListOriginalPrice) Validate() error {
	return dara.Validate(s)
}
