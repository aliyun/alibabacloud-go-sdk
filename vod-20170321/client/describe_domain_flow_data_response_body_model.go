// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainFlowDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainFlowDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainFlowDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainFlowDataResponseBody
	GetEndTime() *string
	SetFlowDataPerInterval(v *DescribeDomainFlowDataResponseBodyFlowDataPerInterval) *DescribeDomainFlowDataResponseBody
	GetFlowDataPerInterval() *DescribeDomainFlowDataResponseBodyFlowDataPerInterval
	SetRequestId(v string) *DescribeDomainFlowDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainFlowDataResponseBody
	GetStartTime() *string
}

type DescribeDomainFlowDataResponseBody struct {
	DataInterval        *string                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName          *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime             *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowDataPerInterval *DescribeDomainFlowDataResponseBodyFlowDataPerInterval `json:"FlowDataPerInterval,omitempty" xml:"FlowDataPerInterval,omitempty" type:"Struct"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime           *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainFlowDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainFlowDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainFlowDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainFlowDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainFlowDataResponseBody) GetFlowDataPerInterval() *DescribeDomainFlowDataResponseBodyFlowDataPerInterval {
	return s.FlowDataPerInterval
}

func (s *DescribeDomainFlowDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainFlowDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainFlowDataResponseBody) SetDataInterval(v string) *DescribeDomainFlowDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetDomainName(v string) *DescribeDomainFlowDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetEndTime(v string) *DescribeDomainFlowDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetFlowDataPerInterval(v *DescribeDomainFlowDataResponseBodyFlowDataPerInterval) *DescribeDomainFlowDataResponseBody {
	s.FlowDataPerInterval = v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetRequestId(v string) *DescribeDomainFlowDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetStartTime(v string) *DescribeDomainFlowDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainFlowDataResponseBodyFlowDataPerInterval struct {
	DataModule []*DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainFlowDataResponseBodyFlowDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainFlowDataResponseBodyFlowDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerInterval) GetDataModule() []*DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerInterval) SetDataModule(v []*DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) *DescribeDomainFlowDataResponseBodyFlowDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule struct {
	DomesticValue        *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	DynamicDomesticValue *string `json:"DynamicDomesticValue,omitempty" xml:"DynamicDomesticValue,omitempty"`
	DynamicOverseasValue *string `json:"DynamicOverseasValue,omitempty" xml:"DynamicOverseasValue,omitempty"`
	DynamicValue         *string `json:"DynamicValue,omitempty" xml:"DynamicValue,omitempty"`
	OverseasValue        *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	StaticDomesticValue  *string `json:"StaticDomesticValue,omitempty" xml:"StaticDomesticValue,omitempty"`
	StaticOverseasValue  *string `json:"StaticOverseasValue,omitempty" xml:"StaticOverseasValue,omitempty"`
	StaticValue          *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
	TimeStamp            *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetDynamicDomesticValue() *string {
	return s.DynamicDomesticValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetDynamicOverseasValue() *string {
	return s.DynamicOverseasValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetDynamicValue() *string {
	return s.DynamicValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetStaticDomesticValue() *string {
	return s.StaticDomesticValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetStaticOverseasValue() *string {
	return s.StaticOverseasValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetStaticValue() *string {
	return s.StaticValue
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDynamicDomesticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DynamicDomesticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDynamicOverseasValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DynamicOverseasValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDynamicValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DynamicValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetStaticDomesticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.StaticDomesticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetStaticOverseasValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.StaticOverseasValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetStaticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.StaticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
