// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainBpsDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataInterval(v *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeLiveDomainBpsDataByLayerResponseBody
	GetBpsDataInterval() *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval
	SetDataInterval(v string) *DescribeLiveDomainBpsDataByLayerResponseBody
	GetDataInterval() *string
	SetRequestId(v string) *DescribeLiveDomainBpsDataByLayerResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainBpsDataByLayerResponseBody struct {
	BpsDataInterval *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval `json:"BpsDataInterval,omitempty" xml:"BpsDataInterval,omitempty" type:"Struct"`
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
	// BEA5625F-8FCF-48F4-851B-2A48566EA967
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainBpsDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBody) GetBpsDataInterval() *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval {
	return s.BpsDataInterval
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBody) SetBpsDataInterval(v *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeLiveDomainBpsDataByLayerResponseBody {
	s.BpsDataInterval = v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBody) SetDataInterval(v string) *DescribeLiveDomainBpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBody) SetRequestId(v string) *DescribeLiveDomainBpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBody) Validate() error {
	if s.BpsDataInterval != nil {
		if err := s.BpsDataInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval struct {
	DataModule []*DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval) GetDataModule() []*DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval) SetDataModule(v []*DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataInterval) Validate() error {
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

type DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule struct {
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTrafficValue() *string {
	return s.TrafficValue
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTrafficValue(v string) *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetValue(v string) *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
