// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainBpsDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataInterval(v *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeDcdnDomainBpsDataByLayerResponseBody
	GetBpsDataInterval() *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval
	SetDataInterval(v string) *DescribeDcdnDomainBpsDataByLayerResponseBody
	GetDataInterval() *string
	SetRequestId(v string) *DescribeDcdnDomainBpsDataByLayerResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainBpsDataByLayerResponseBody struct {
	// The bandwidth returned at each time interval.
	BpsDataInterval *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval `json:"BpsDataInterval,omitempty" xml:"BpsDataInterval,omitempty" type:"Struct"`
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainBpsDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBody) GetBpsDataInterval() *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval {
	return s.BpsDataInterval
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBody) SetBpsDataInterval(v *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeDcdnDomainBpsDataByLayerResponseBody {
	s.BpsDataInterval = v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBody) SetDataInterval(v string) *DescribeDcdnDomainBpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBody) SetRequestId(v string) *DescribeDcdnDomainBpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval struct {
	DataModule []*DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval) GetDataModule() []*DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval) SetDataModule(v []*DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule struct {
	// The traffic that is used to deliver dynamic content. Unit: bytes.
	//
	// example:
	//
	// 200
	DynamicTrafficValue *string `json:"DynamicTrafficValue,omitempty" xml:"DynamicTrafficValue,omitempty"`
	// The bandwidth that is used to deliver dynamic content. Unit: bit/s.
	//
	// example:
	//
	// 0.34
	DynamicValue *string `json:"DynamicValue,omitempty" xml:"DynamicValue,omitempty"`
	// The traffic that is used to deliver static content. Unit: bytes.
	//
	// example:
	//
	// 131
	StaticTrafficValue *string `json:"StaticTrafficValue,omitempty" xml:"StaticTrafficValue,omitempty"`
	// The bandwidth that is used to deliver static content. Unit: bit/s.
	//
	// example:
	//
	// 0.22
	StaticValue *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total traffic. Unit: bytes.
	//
	// example:
	//
	// 331
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
	// The total bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 0.56
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetDynamicTrafficValue() *string {
	return s.DynamicTrafficValue
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetDynamicValue() *string {
	return s.DynamicValue
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetStaticTrafficValue() *string {
	return s.StaticTrafficValue
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetStaticValue() *string {
	return s.StaticValue
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTrafficValue() *string {
	return s.TrafficValue
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetDynamicTrafficValue(v string) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.DynamicTrafficValue = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetDynamicValue(v string) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.DynamicValue = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetStaticTrafficValue(v string) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.StaticTrafficValue = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetStaticValue(v string) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.StaticValue = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTrafficValue(v string) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetValue(v string) *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
