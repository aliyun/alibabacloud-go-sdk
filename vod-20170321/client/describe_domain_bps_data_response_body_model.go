// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDomainBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeDomainBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeDomainBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainBpsDataResponseBody
	GetStartTime() *string
	SetSupplyBpsDatas(v *DescribeDomainBpsDataResponseBodySupplyBpsDatas) *DescribeDomainBpsDataResponseBody
	GetSupplyBpsDatas() *DescribeDomainBpsDataResponseBodySupplyBpsDatas
}

type DescribeDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SupplyBpsDatas     *DescribeDomainBpsDataResponseBodySupplyBpsDatas     `json:"SupplyBpsDatas,omitempty" xml:"SupplyBpsDatas,omitempty" type:"Struct"`
}

func (s DescribeDomainBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBody) GetBpsDataPerInterval() *DescribeDomainBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeDomainBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainBpsDataResponseBody) GetSupplyBpsDatas() *DescribeDomainBpsDataResponseBodySupplyBpsDatas {
	return s.SupplyBpsDatas
}

func (s *DescribeDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDomainName(v string) *DescribeDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetEndTime(v string) *DescribeDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetRequestId(v string) *DescribeDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetStartTime(v string) *DescribeDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetSupplyBpsDatas(v *DescribeDomainBpsDataResponseBodySupplyBpsDatas) *DescribeDomainBpsDataResponseBody {
	s.SupplyBpsDatas = v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	DomesticL2Value      *string `json:"DomesticL2Value,omitempty" xml:"DomesticL2Value,omitempty"`
	DomesticValue        *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	DynamicDomesticValue *string `json:"DynamicDomesticValue,omitempty" xml:"DynamicDomesticValue,omitempty"`
	DynamicOverseasValue *string `json:"DynamicOverseasValue,omitempty" xml:"DynamicOverseasValue,omitempty"`
	DynamicValue         *int64  `json:"DynamicValue,omitempty" xml:"DynamicValue,omitempty"`
	L2Value              *string `json:"L2Value,omitempty" xml:"L2Value,omitempty"`
	OverseasL2Value      *string `json:"OverseasL2Value,omitempty" xml:"OverseasL2Value,omitempty"`
	OverseasValue        *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	StaticDomesticValue  *string `json:"StaticDomesticValue,omitempty" xml:"StaticDomesticValue,omitempty"`
	StaticOverseasValue  *string `json:"StaticOverseasValue,omitempty" xml:"StaticOverseasValue,omitempty"`
	StaticValue          *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
	TimeStamp            *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDomesticL2Value() *string {
	return s.DomesticL2Value
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDynamicDomesticValue() *string {
	return s.DynamicDomesticValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDynamicOverseasValue() *string {
	return s.DynamicOverseasValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDynamicValue() *int64 {
	return s.DynamicValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetL2Value() *string {
	return s.L2Value
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetOverseasL2Value() *string {
	return s.OverseasL2Value
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetStaticDomesticValue() *string {
	return s.StaticDomesticValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetStaticOverseasValue() *string {
	return s.StaticOverseasValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetStaticValue() *string {
	return s.StaticValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticL2Value(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticL2Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicDomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicOverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicValue(v int64) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetL2Value(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.L2Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasL2Value(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasL2Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticDomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticOverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainBpsDataResponseBodySupplyBpsDatas struct {
	DataModule []*DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataResponseBodySupplyBpsDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodySupplyBpsDatas) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatas) GetDataModule() []*DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule {
	return s.DataModule
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatas) SetDataModule(v []*DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) *DescribeDomainBpsDataResponseBodySupplyBpsDatas {
	s.DataModule = v
	return s
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatas) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) SetTimeStamp(v string) *DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) SetValue(v string) *DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodySupplyBpsDatasDataModule) Validate() error {
	return dara.Validate(s)
}
