// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHttpCodeDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainHttpCodeDataByLayerResponseBody
	GetDataInterval() *string
	SetHttpCodeDataInterval(v *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) *DescribeDomainHttpCodeDataByLayerResponseBody
	GetHttpCodeDataInterval() *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval
	SetRequestId(v string) *DescribeDomainHttpCodeDataByLayerResponseBody
	GetRequestId() *string
}

type DescribeDomainHttpCodeDataByLayerResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The statistics of HTTP status codes returned at each time interval.
	HttpCodeDataInterval *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval `json:"HttpCodeDataInterval,omitempty" xml:"HttpCodeDataInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainHttpCodeDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) GetHttpCodeDataInterval() *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval {
	return s.HttpCodeDataInterval
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) SetDataInterval(v string) *DescribeDomainHttpCodeDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) SetHttpCodeDataInterval(v *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) *DescribeDomainHttpCodeDataByLayerResponseBody {
	s.HttpCodeDataInterval = v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainHttpCodeDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) Validate() error {
	if s.HttpCodeDataInterval != nil {
		if err := s.HttpCodeDataInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval struct {
	DataModule []*DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) GetDataModule() []*DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) SetDataModule(v []*DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) Validate() error {
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

type DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:35:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total number of times that HTTP status codes were returned.
	//
	// example:
	//
	// 110
	TotalValue *string `json:"TotalValue,omitempty" xml:"TotalValue,omitempty"`
	// The number of times that the HTTP status code was returned.
	//
	// example:
	//
	// { "200": 10, "206": 100}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GetTotalValue() *string {
	return s.TotalValue
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetTotalValue(v string) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.TotalValue = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetValue(v string) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
