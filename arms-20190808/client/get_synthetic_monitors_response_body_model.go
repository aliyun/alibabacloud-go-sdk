// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetSyntheticMonitorsResponseBody
	GetCode() *int64
	SetData(v []*GetSyntheticMonitorsResponseBodyData) *GetSyntheticMonitorsResponseBody
	GetData() []*GetSyntheticMonitorsResponseBodyData
	SetMessage(v string) *GetSyntheticMonitorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSyntheticMonitorsResponseBody
	GetRequestId() *string
}

type GetSyntheticMonitorsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of monitoring points.
	Data []*GetSyntheticMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FD473FF-5398-5A85-9BF6-7AB45561522F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSyntheticMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyntheticMonitorsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetSyntheticMonitorsResponseBody) GetData() []*GetSyntheticMonitorsResponseBodyData {
	return s.Data
}

func (s *GetSyntheticMonitorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSyntheticMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSyntheticMonitorsResponseBody) SetCode(v int64) *GetSyntheticMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBody) SetData(v []*GetSyntheticMonitorsResponseBodyData) *GetSyntheticMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *GetSyntheticMonitorsResponseBody) SetMessage(v string) *GetSyntheticMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBody) SetRequestId(v string) *GetSyntheticMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticMonitorsResponseBodyData struct {
	// Indicates whether the monitoring point is available. Valid values: true and false.
	//
	// example:
	//
	// true
	Available *string `json:"Available,omitempty" xml:"Available,omitempty"`
	// Indicates whether the monitoring point is selected. Valid values: true and false.
	//
	// example:
	//
	// true
	CanBeSelected *bool `json:"CanBeSelected,omitempty" xml:"CanBeSelected,omitempty"`
	// The city.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The city code.
	//
	// example:
	//
	// 1100101
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The client type of the monitoring point. Valid values: 1: data center. 2: Internet. 3: mobile device. 4: ECS instance.
	//
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The country.
	//
	// example:
	//
	// China
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// Indicates whether IPv6 is supported. Valid values: 0: IPv6 is not supported. 1: IPv6 is supported.
	//
	// example:
	//
	// 0
	Ipv6 *int32 `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// The carrier.
	//
	// example:
	//
	// Alibaba Cloud
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The carrier code.
	//
	// example:
	//
	// 1
	OperatorCode *string `json:"OperatorCode,omitempty" xml:"OperatorCode,omitempty"`
	// The region.
	//
	// example:
	//
	// Beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetSyntheticMonitorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSyntheticMonitorsResponseBodyData) GetAvailable() *string {
	return s.Available
}

func (s *GetSyntheticMonitorsResponseBodyData) GetCanBeSelected() *bool {
	return s.CanBeSelected
}

func (s *GetSyntheticMonitorsResponseBodyData) GetCity() *string {
	return s.City
}

func (s *GetSyntheticMonitorsResponseBodyData) GetCityCode() *string {
	return s.CityCode
}

func (s *GetSyntheticMonitorsResponseBodyData) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetSyntheticMonitorsResponseBodyData) GetCountry() *string {
	return s.Country
}

func (s *GetSyntheticMonitorsResponseBodyData) GetIpv6() *int32 {
	return s.Ipv6
}

func (s *GetSyntheticMonitorsResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *GetSyntheticMonitorsResponseBodyData) GetOperatorCode() *string {
	return s.OperatorCode
}

func (s *GetSyntheticMonitorsResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetSyntheticMonitorsResponseBodyData) SetAvailable(v string) *GetSyntheticMonitorsResponseBodyData {
	s.Available = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetCanBeSelected(v bool) *GetSyntheticMonitorsResponseBodyData {
	s.CanBeSelected = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetCity(v string) *GetSyntheticMonitorsResponseBodyData {
	s.City = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetCityCode(v string) *GetSyntheticMonitorsResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetClientType(v int32) *GetSyntheticMonitorsResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetCountry(v string) *GetSyntheticMonitorsResponseBodyData {
	s.Country = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetIpv6(v int32) *GetSyntheticMonitorsResponseBodyData {
	s.Ipv6 = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetOperator(v string) *GetSyntheticMonitorsResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetOperatorCode(v string) *GetSyntheticMonitorsResponseBodyData {
	s.OperatorCode = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) SetRegion(v string) *GetSyntheticMonitorsResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetSyntheticMonitorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
