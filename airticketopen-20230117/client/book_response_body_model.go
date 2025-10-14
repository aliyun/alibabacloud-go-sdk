// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BookResponseBody
	GetRequestId() *string
	SetData(v *BookResponseBodyData) *BookResponseBody
	GetData() *BookResponseBodyData
	SetErrorCode(v string) *BookResponseBody
	GetErrorCode() *string
	SetErrorData(v *BookResponseBodyErrorData) *BookResponseBody
	GetErrorData() *BookResponseBodyErrorData
	SetErrorMsg(v string) *BookResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *BookResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *BookResponseBody
	GetSuccess() *bool
}

type BookResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *BookResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData *BookResponseBodyErrorData `json:"error_data,omitempty" xml:"error_data,omitempty" type:"Struct"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBody) GoString() string {
	return s.String()
}

func (s *BookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BookResponseBody) GetData() *BookResponseBodyData {
	return s.Data
}

func (s *BookResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BookResponseBody) GetErrorData() *BookResponseBodyErrorData {
	return s.ErrorData
}

func (s *BookResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BookResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *BookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BookResponseBody) SetRequestId(v string) *BookResponseBody {
	s.RequestId = &v
	return s
}

func (s *BookResponseBody) SetData(v *BookResponseBodyData) *BookResponseBody {
	s.Data = v
	return s
}

func (s *BookResponseBody) SetErrorCode(v string) *BookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BookResponseBody) SetErrorData(v *BookResponseBodyErrorData) *BookResponseBody {
	s.ErrorData = v
	return s
}

func (s *BookResponseBody) SetErrorMsg(v string) *BookResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *BookResponseBody) SetStatus(v int32) *BookResponseBody {
	s.Status = &v
	return s
}

func (s *BookResponseBody) SetSuccess(v bool) *BookResponseBody {
	s.Success = &v
	return s
}

func (s *BookResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorData != nil {
		if err := s.ErrorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookResponseBodyData struct {
	// order information list
	OrderList []*BookResponseBodyDataOrderList `json:"order_list,omitempty" xml:"order_list,omitempty" type:"Repeated"`
}

func (s BookResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyData) GoString() string {
	return s.String()
}

func (s *BookResponseBodyData) GetOrderList() []*BookResponseBodyDataOrderList {
	return s.OrderList
}

func (s *BookResponseBodyData) SetOrderList(v []*BookResponseBodyDataOrderList) *BookResponseBodyData {
	s.OrderList = v
	return s
}

func (s *BookResponseBodyData) Validate() error {
	if s.OrderList != nil {
		for _, item := range s.OrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BookResponseBodyDataOrderList struct {
	OrderAttribute *BookResponseBodyDataOrderListOrderAttribute `json:"order_attribute,omitempty" xml:"order_attribute,omitempty" type:"Struct"`
	// order number
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s BookResponseBodyDataOrderList) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyDataOrderList) GoString() string {
	return s.String()
}

func (s *BookResponseBodyDataOrderList) GetOrderAttribute() *BookResponseBodyDataOrderListOrderAttribute {
	return s.OrderAttribute
}

func (s *BookResponseBodyDataOrderList) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *BookResponseBodyDataOrderList) SetOrderAttribute(v *BookResponseBodyDataOrderListOrderAttribute) *BookResponseBodyDataOrderList {
	s.OrderAttribute = v
	return s
}

func (s *BookResponseBodyDataOrderList) SetOrderNum(v int64) *BookResponseBodyDataOrderList {
	s.OrderNum = &v
	return s
}

func (s *BookResponseBodyDataOrderList) Validate() error {
	if s.OrderAttribute != nil {
		if err := s.OrderAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookResponseBodyDataOrderListOrderAttribute struct {
	AbaPayLockRateInfo *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo `json:"aba_pay_lock_rate_info,omitempty" xml:"aba_pay_lock_rate_info,omitempty" type:"Struct"`
}

func (s BookResponseBodyDataOrderListOrderAttribute) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyDataOrderListOrderAttribute) GoString() string {
	return s.String()
}

func (s *BookResponseBodyDataOrderListOrderAttribute) GetAbaPayLockRateInfo() *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo {
	return s.AbaPayLockRateInfo
}

func (s *BookResponseBodyDataOrderListOrderAttribute) SetAbaPayLockRateInfo(v *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) *BookResponseBodyDataOrderListOrderAttribute {
	s.AbaPayLockRateInfo = v
	return s
}

func (s *BookResponseBodyDataOrderListOrderAttribute) Validate() error {
	if s.AbaPayLockRateInfo != nil {
		if err := s.AbaPayLockRateInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo struct {
	// example:
	//
	// 14.2854
	PayIntendedAmount *string `json:"pay_intended_amount,omitempty" xml:"pay_intended_amount,omitempty"`
	// example:
	//
	// CNY
	PayIntendedCurrencyCode *string `json:"pay_intended_currency_code,omitempty" xml:"pay_intended_currency_code,omitempty"`
	// example:
	//
	// USD
	QuotationCurrencyCode *string `json:"quotation_currency_code,omitempty" xml:"quotation_currency_code,omitempty"`
	// example:
	//
	// 7.1427
	ToPayCurrencyRate *string `json:"to_pay_currency_rate,omitempty" xml:"to_pay_currency_rate,omitempty"`
}

func (s BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) GoString() string {
	return s.String()
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) GetPayIntendedAmount() *string {
	return s.PayIntendedAmount
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) GetPayIntendedCurrencyCode() *string {
	return s.PayIntendedCurrencyCode
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) GetQuotationCurrencyCode() *string {
	return s.QuotationCurrencyCode
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) GetToPayCurrencyRate() *string {
	return s.ToPayCurrencyRate
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) SetPayIntendedAmount(v string) *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.PayIntendedAmount = &v
	return s
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) SetPayIntendedCurrencyCode(v string) *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.PayIntendedCurrencyCode = &v
	return s
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) SetQuotationCurrencyCode(v string) *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.QuotationCurrencyCode = &v
	return s
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) SetToPayCurrencyRate(v string) *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.ToPayCurrencyRate = &v
	return s
}

func (s *BookResponseBodyDataOrderListOrderAttributeAbaPayLockRateInfo) Validate() error {
	return dara.Validate(s)
}

type BookResponseBodyErrorData struct {
	// order information list. When the same input parameters are used to repeat a Book, if the booking has already been successful, the order number will be returned.
	OrderList []*BookResponseBodyErrorDataOrderList `json:"order_list,omitempty" xml:"order_list,omitempty" type:"Repeated"`
}

func (s BookResponseBodyErrorData) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyErrorData) GoString() string {
	return s.String()
}

func (s *BookResponseBodyErrorData) GetOrderList() []*BookResponseBodyErrorDataOrderList {
	return s.OrderList
}

func (s *BookResponseBodyErrorData) SetOrderList(v []*BookResponseBodyErrorDataOrderList) *BookResponseBodyErrorData {
	s.OrderList = v
	return s
}

func (s *BookResponseBodyErrorData) Validate() error {
	if s.OrderList != nil {
		for _, item := range s.OrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BookResponseBodyErrorDataOrderList struct {
	OrderAttribute *BookResponseBodyErrorDataOrderListOrderAttribute `json:"order_attribute,omitempty" xml:"order_attribute,omitempty" type:"Struct"`
	// order number
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s BookResponseBodyErrorDataOrderList) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyErrorDataOrderList) GoString() string {
	return s.String()
}

func (s *BookResponseBodyErrorDataOrderList) GetOrderAttribute() *BookResponseBodyErrorDataOrderListOrderAttribute {
	return s.OrderAttribute
}

func (s *BookResponseBodyErrorDataOrderList) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *BookResponseBodyErrorDataOrderList) SetOrderAttribute(v *BookResponseBodyErrorDataOrderListOrderAttribute) *BookResponseBodyErrorDataOrderList {
	s.OrderAttribute = v
	return s
}

func (s *BookResponseBodyErrorDataOrderList) SetOrderNum(v int64) *BookResponseBodyErrorDataOrderList {
	s.OrderNum = &v
	return s
}

func (s *BookResponseBodyErrorDataOrderList) Validate() error {
	if s.OrderAttribute != nil {
		if err := s.OrderAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookResponseBodyErrorDataOrderListOrderAttribute struct {
	AbaPayLockRateInfo *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo `json:"aba_pay_lock_rate_info,omitempty" xml:"aba_pay_lock_rate_info,omitempty" type:"Struct"`
}

func (s BookResponseBodyErrorDataOrderListOrderAttribute) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyErrorDataOrderListOrderAttribute) GoString() string {
	return s.String()
}

func (s *BookResponseBodyErrorDataOrderListOrderAttribute) GetAbaPayLockRateInfo() *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo {
	return s.AbaPayLockRateInfo
}

func (s *BookResponseBodyErrorDataOrderListOrderAttribute) SetAbaPayLockRateInfo(v *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) *BookResponseBodyErrorDataOrderListOrderAttribute {
	s.AbaPayLockRateInfo = v
	return s
}

func (s *BookResponseBodyErrorDataOrderListOrderAttribute) Validate() error {
	if s.AbaPayLockRateInfo != nil {
		if err := s.AbaPayLockRateInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo struct {
	// example:
	//
	// 14.2854
	PayIntendedAmount *string `json:"pay_intended_amount,omitempty" xml:"pay_intended_amount,omitempty"`
	// example:
	//
	// CNY
	PayIntendedCurrencyCode *string `json:"pay_intended_currency_code,omitempty" xml:"pay_intended_currency_code,omitempty"`
	// example:
	//
	// USD
	QuotationCurrencyCode *string `json:"quotation_currency_code,omitempty" xml:"quotation_currency_code,omitempty"`
	// example:
	//
	// 7.1427
	ToPayCurrencyRate *string `json:"to_pay_currency_rate,omitempty" xml:"to_pay_currency_rate,omitempty"`
}

func (s BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) String() string {
	return dara.Prettify(s)
}

func (s BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) GoString() string {
	return s.String()
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) GetPayIntendedAmount() *string {
	return s.PayIntendedAmount
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) GetPayIntendedCurrencyCode() *string {
	return s.PayIntendedCurrencyCode
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) GetQuotationCurrencyCode() *string {
	return s.QuotationCurrencyCode
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) GetToPayCurrencyRate() *string {
	return s.ToPayCurrencyRate
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) SetPayIntendedAmount(v string) *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.PayIntendedAmount = &v
	return s
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) SetPayIntendedCurrencyCode(v string) *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.PayIntendedCurrencyCode = &v
	return s
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) SetQuotationCurrencyCode(v string) *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.QuotationCurrencyCode = &v
	return s
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) SetToPayCurrencyRate(v string) *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo {
	s.ToPayCurrencyRate = &v
	return s
}

func (s *BookResponseBodyErrorDataOrderListOrderAttributeAbaPayLockRateInfo) Validate() error {
	return dara.Validate(s)
}
