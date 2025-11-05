// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainTrafficDataResponseBody
	GetStartTime() *string
	SetTrafficDataPerInterval(v *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDomainTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeDomainTrafficDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The amount of network traffic at each time interval. Unit: bytes.
	TrafficDataPerInterval *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) Validate() error {
	if s.TrafficDataPerInterval != nil {
		if err := s.TrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	// The amount of network traffic in the Chinese mainland.
	//
	// example:
	//
	// 0
	DomesticValue *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	// The amount of HTTPS traffic on points of presence (POPs) in the Chinese mainland.
	//
	// example:
	//
	// 0
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	// The amount of HTTPS traffic on POPs outside the Chinese mainland.
	//
	// example:
	//
	// 0
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	// The total amount of HTTPS traffic on POPs.
	//
	// example:
	//
	// 423304182
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The amount of network traffic outside the Chinese mainland.
	//
	// example:
	//
	// 0
	OverseasValue *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total volume of traffic.
	//
	// example:
	//
	// 423304182
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpsDomesticValue() *string {
	return s.HttpsDomesticValue
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpsOverseasValue() *string {
	return s.HttpsOverseasValue
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
