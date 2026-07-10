// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelAskingPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelAskingPriceResponseBody
	GetCode() *string
	SetMessage(v string) *HotelAskingPriceResponseBody
	GetMessage() *string
	SetModule(v *HotelAskingPriceResponseBodyModule) *HotelAskingPriceResponseBody
	GetModule() *HotelAskingPriceResponseBodyModule
	SetRequestId(v string) *HotelAskingPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelAskingPriceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelAskingPriceResponseBody
	GetTraceId() *string
}

type HotelAskingPriceResponseBody struct {
	Code      *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module    *HotelAskingPriceResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                             `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelAskingPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelAskingPriceResponseBody) GoString() string {
	return s.String()
}

func (s *HotelAskingPriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelAskingPriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelAskingPriceResponseBody) GetModule() *HotelAskingPriceResponseBodyModule {
	return s.Module
}

func (s *HotelAskingPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelAskingPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelAskingPriceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelAskingPriceResponseBody) SetCode(v string) *HotelAskingPriceResponseBody {
	s.Code = &v
	return s
}

func (s *HotelAskingPriceResponseBody) SetMessage(v string) *HotelAskingPriceResponseBody {
	s.Message = &v
	return s
}

func (s *HotelAskingPriceResponseBody) SetModule(v *HotelAskingPriceResponseBodyModule) *HotelAskingPriceResponseBody {
	s.Module = v
	return s
}

func (s *HotelAskingPriceResponseBody) SetRequestId(v string) *HotelAskingPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelAskingPriceResponseBody) SetSuccess(v bool) *HotelAskingPriceResponseBody {
	s.Success = &v
	return s
}

func (s *HotelAskingPriceResponseBody) SetTraceId(v string) *HotelAskingPriceResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelAskingPriceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HotelAskingPriceResponseBodyModule struct {
	HotelAskingPriceDetails []*HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails `json:"hotel_asking_price_details,omitempty" xml:"hotel_asking_price_details,omitempty" type:"Repeated"`
}

func (s HotelAskingPriceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelAskingPriceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelAskingPriceResponseBodyModule) GetHotelAskingPriceDetails() []*HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	return s.HotelAskingPriceDetails
}

func (s *HotelAskingPriceResponseBodyModule) SetHotelAskingPriceDetails(v []*HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) *HotelAskingPriceResponseBodyModule {
	s.HotelAskingPriceDetails = v
	return s
}

func (s *HotelAskingPriceResponseBodyModule) Validate() error {
	if s.HotelAskingPriceDetails != nil {
		for _, item := range s.HotelAskingPriceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails struct {
	CityCode         *string  `json:"city_code,omitempty" xml:"city_code,omitempty"`
	HotelAddress     *string  `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	HotelCode        *string  `json:"hotel_code,omitempty" xml:"hotel_code,omitempty"`
	HotelName        *string  `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	IsProtocol       *bool    `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
	MinPrice         *float64 `json:"min_price,omitempty" xml:"min_price,omitempty"`
	OriginalMinPrice *float64 `json:"original_min_price,omitempty" xml:"original_min_price,omitempty"`
}

func (s HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) String() string {
	return dara.Prettify(s)
}

func (s HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GoString() string {
	return s.String()
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GetHotelCode() *string {
	return s.HotelCode
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GetMinPrice() *float64 {
	return s.MinPrice
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) GetOriginalMinPrice() *float64 {
	return s.OriginalMinPrice
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) SetCityCode(v string) *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	s.CityCode = &v
	return s
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) SetHotelAddress(v string) *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	s.HotelAddress = &v
	return s
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) SetHotelCode(v string) *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	s.HotelCode = &v
	return s
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) SetHotelName(v string) *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	s.HotelName = &v
	return s
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) SetIsProtocol(v bool) *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	s.IsProtocol = &v
	return s
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) SetMinPrice(v float64) *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	s.MinPrice = &v
	return s
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) SetOriginalMinPrice(v float64) *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails {
	s.OriginalMinPrice = &v
	return s
}

func (s *HotelAskingPriceResponseBodyModuleHotelAskingPriceDetails) Validate() error {
	return dara.Validate(s)
}
