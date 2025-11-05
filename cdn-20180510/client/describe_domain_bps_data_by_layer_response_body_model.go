// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataInterval(v *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeDomainBpsDataByLayerResponseBody
	GetBpsDataInterval() *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval
	SetDataInterval(v string) *DescribeDomainBpsDataByLayerResponseBody
	GetDataInterval() *string
	SetRequestId(v string) *DescribeDomainBpsDataByLayerResponseBody
	GetRequestId() *string
}

type DescribeDomainBpsDataByLayerResponseBody struct {
	// The data returned at each time interval.
	BpsDataInterval *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval `json:"BpsDataInterval,omitempty" xml:"BpsDataInterval,omitempty" type:"Struct"`
	// The time interval between the data entries. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C565B910-BC3B-467B-9046-2A48566EA967
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainBpsDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponseBody) GetBpsDataInterval() *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval {
	return s.BpsDataInterval
}

func (s *DescribeDomainBpsDataByLayerResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainBpsDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainBpsDataByLayerResponseBody) SetBpsDataInterval(v *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeDomainBpsDataByLayerResponseBody {
	s.BpsDataInterval = v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBody) SetDataInterval(v string) *DescribeDomainBpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainBpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBody) Validate() error {
	if s.BpsDataInterval != nil {
		if err := s.BpsDataInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval struct {
	DataModule []*DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) GetDataModule() []*DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) SetDataModule(v []*DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) Validate() error {
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

type DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2020-05-06T07:10:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// 2838
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
	// The peak bandwidth value. Unit: bit/s.
	//
	// example:
	//
	// 75.68
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTrafficValue() *string {
	return s.TrafficValue
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTrafficValue(v string) *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetValue(v string) *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
