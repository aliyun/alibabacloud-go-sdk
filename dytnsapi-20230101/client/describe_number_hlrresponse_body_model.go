// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNumberHLRResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeNumberHLRResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeNumberHLRResponseBody
	GetCode() *string
	SetData(v *DescribeNumberHLRResponseBodyData) *DescribeNumberHLRResponseBody
	GetData() *DescribeNumberHLRResponseBodyData
	SetMessage(v string) *DescribeNumberHLRResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeNumberHLRResponseBody
	GetRequestId() *string
}

type DescribeNumberHLRResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeNumberHLRResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1C3B8084-3A7D-570B-BC84-BF945A9CF65E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNumberHLRResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberHLRResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNumberHLRResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeNumberHLRResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeNumberHLRResponseBody) GetData() *DescribeNumberHLRResponseBodyData {
	return s.Data
}

func (s *DescribeNumberHLRResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNumberHLRResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNumberHLRResponseBody) SetAccessDeniedDetail(v string) *DescribeNumberHLRResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeNumberHLRResponseBody) SetCode(v string) *DescribeNumberHLRResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNumberHLRResponseBody) SetData(v *DescribeNumberHLRResponseBodyData) *DescribeNumberHLRResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNumberHLRResponseBody) SetMessage(v string) *DescribeNumberHLRResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNumberHLRResponseBody) SetRequestId(v string) *DescribeNumberHLRResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNumberHLRResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNumberHLRResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	Blocked *string                                `json:"Blocked,omitempty" xml:"Blocked,omitempty"`
	Call    *DescribeNumberHLRResponseBodyDataCall `json:"Call,omitempty" xml:"Call,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 示例值示例值
	CountryIso3 *string                                `json:"CountryIso3,omitempty" xml:"CountryIso3,omitempty"`
	Live        *DescribeNumberHLRResponseBodyDataLive `json:"Live,omitempty" xml:"Live,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	PhoneType *string                               `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	Sms       *DescribeNumberHLRResponseBodyDataSms `json:"Sms,omitempty" xml:"Sms,omitempty" type:"Struct"`
}

func (s DescribeNumberHLRResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberHLRResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNumberHLRResponseBodyData) GetBlocked() *string {
	return s.Blocked
}

func (s *DescribeNumberHLRResponseBodyData) GetCall() *DescribeNumberHLRResponseBodyDataCall {
	return s.Call
}

func (s *DescribeNumberHLRResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribeNumberHLRResponseBodyData) GetCity() *string {
	return s.City
}

func (s *DescribeNumberHLRResponseBodyData) GetCountryIso3() *string {
	return s.CountryIso3
}

func (s *DescribeNumberHLRResponseBodyData) GetLive() *DescribeNumberHLRResponseBodyDataLive {
	return s.Live
}

func (s *DescribeNumberHLRResponseBodyData) GetPhoneType() *string {
	return s.PhoneType
}

func (s *DescribeNumberHLRResponseBodyData) GetSms() *DescribeNumberHLRResponseBodyDataSms {
	return s.Sms
}

func (s *DescribeNumberHLRResponseBodyData) SetBlocked(v string) *DescribeNumberHLRResponseBodyData {
	s.Blocked = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) SetCall(v *DescribeNumberHLRResponseBodyDataCall) *DescribeNumberHLRResponseBodyData {
	s.Call = v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) SetCarrier(v string) *DescribeNumberHLRResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) SetCity(v string) *DescribeNumberHLRResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) SetCountryIso3(v string) *DescribeNumberHLRResponseBodyData {
	s.CountryIso3 = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) SetLive(v *DescribeNumberHLRResponseBodyDataLive) *DescribeNumberHLRResponseBodyData {
	s.Live = v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) SetPhoneType(v string) *DescribeNumberHLRResponseBodyData {
	s.PhoneType = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) SetSms(v *DescribeNumberHLRResponseBodyDataSms) *DescribeNumberHLRResponseBodyData {
	s.Sms = v
	return s
}

func (s *DescribeNumberHLRResponseBodyData) Validate() error {
	if s.Call != nil {
		if err := s.Call.Validate(); err != nil {
			return err
		}
	}
	if s.Live != nil {
		if err := s.Live.Validate(); err != nil {
			return err
		}
	}
	if s.Sms != nil {
		if err := s.Sms.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNumberHLRResponseBodyDataCall struct {
	// call
	//
	// example:
	//
	// 示例值
	CleansedCode *string `json:"CleansedCode,omitempty" xml:"CleansedCode,omitempty"`
	// example:
	//
	// 22
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// example:
	//
	// 21
	MinLength *int64 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
}

func (s DescribeNumberHLRResponseBodyDataCall) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberHLRResponseBodyDataCall) GoString() string {
	return s.String()
}

func (s *DescribeNumberHLRResponseBodyDataCall) GetCleansedCode() *string {
	return s.CleansedCode
}

func (s *DescribeNumberHLRResponseBodyDataCall) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *DescribeNumberHLRResponseBodyDataCall) GetMinLength() *int64 {
	return s.MinLength
}

func (s *DescribeNumberHLRResponseBodyDataCall) SetCleansedCode(v string) *DescribeNumberHLRResponseBodyDataCall {
	s.CleansedCode = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataCall) SetMaxLength(v int64) *DescribeNumberHLRResponseBodyDataCall {
	s.MaxLength = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataCall) SetMinLength(v int64) *DescribeNumberHLRResponseBodyDataCall {
	s.MinLength = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataCall) Validate() error {
	return dara.Validate(s)
}

type DescribeNumberHLRResponseBodyDataLive struct {
	// example:
	//
	// 示例值
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 示例值示例值
	Roaming *string `json:"Roaming,omitempty" xml:"Roaming,omitempty"`
	// example:
	//
	// 示例值
	RoamingCountry *string `json:"RoamingCountry,omitempty" xml:"RoamingCountry,omitempty"`
	// example:
	//
	// 示例值
	SubscriberStatus *string `json:"SubscriberStatus,omitempty" xml:"SubscriberStatus,omitempty"`
}

func (s DescribeNumberHLRResponseBodyDataLive) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberHLRResponseBodyDataLive) GoString() string {
	return s.String()
}

func (s *DescribeNumberHLRResponseBodyDataLive) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *DescribeNumberHLRResponseBodyDataLive) GetRoaming() *string {
	return s.Roaming
}

func (s *DescribeNumberHLRResponseBodyDataLive) GetRoamingCountry() *string {
	return s.RoamingCountry
}

func (s *DescribeNumberHLRResponseBodyDataLive) GetSubscriberStatus() *string {
	return s.SubscriberStatus
}

func (s *DescribeNumberHLRResponseBodyDataLive) SetDeviceStatus(v string) *DescribeNumberHLRResponseBodyDataLive {
	s.DeviceStatus = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataLive) SetRoaming(v string) *DescribeNumberHLRResponseBodyDataLive {
	s.Roaming = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataLive) SetRoamingCountry(v string) *DescribeNumberHLRResponseBodyDataLive {
	s.RoamingCountry = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataLive) SetSubscriberStatus(v string) *DescribeNumberHLRResponseBodyDataLive {
	s.SubscriberStatus = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataLive) Validate() error {
	return dara.Validate(s)
}

type DescribeNumberHLRResponseBodyDataSms struct {
	// sms
	//
	// example:
	//
	// 58
	CleansedCode *int64 `json:"CleansedCode,omitempty" xml:"CleansedCode,omitempty"`
	// example:
	//
	// 59
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// example:
	//
	// 2
	MinLength *int64 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
}

func (s DescribeNumberHLRResponseBodyDataSms) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberHLRResponseBodyDataSms) GoString() string {
	return s.String()
}

func (s *DescribeNumberHLRResponseBodyDataSms) GetCleansedCode() *int64 {
	return s.CleansedCode
}

func (s *DescribeNumberHLRResponseBodyDataSms) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *DescribeNumberHLRResponseBodyDataSms) GetMinLength() *int64 {
	return s.MinLength
}

func (s *DescribeNumberHLRResponseBodyDataSms) SetCleansedCode(v int64) *DescribeNumberHLRResponseBodyDataSms {
	s.CleansedCode = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataSms) SetMaxLength(v int64) *DescribeNumberHLRResponseBodyDataSms {
	s.MaxLength = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataSms) SetMinLength(v int64) *DescribeNumberHLRResponseBodyDataSms {
	s.MinLength = &v
	return s
}

func (s *DescribeNumberHLRResponseBodyDataSms) Validate() error {
	return dara.Validate(s)
}
