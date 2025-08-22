// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHttpCodeDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainHttpCodeDataByLayerResponseBody
	GetDataInterval() *string
	SetHttpCodeDataInterval(v *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) *DescribeDcdnDomainHttpCodeDataByLayerResponseBody
	GetHttpCodeDataInterval() *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval
	SetRequestId(v string) *DescribeDcdnDomainHttpCodeDataByLayerResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainHttpCodeDataByLayerResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The distribution of HTTP status codes at each time interval.
	HttpCodeDataInterval *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval `json:"HttpCodeDataInterval,omitempty" xml:"HttpCodeDataInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) GetHttpCodeDataInterval() *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval {
	return s.HttpCodeDataInterval
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) SetDataInterval(v string) *DescribeDcdnDomainHttpCodeDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) SetHttpCodeDataInterval(v *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) *DescribeDcdnDomainHttpCodeDataByLayerResponseBody {
	s.HttpCodeDataInterval = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) SetRequestId(v string) *DescribeDcdnDomainHttpCodeDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval struct {
	DataModule []*DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) GetDataModule() []*DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) SetDataModule(v []*DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
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
	// {"200": 10,"206": 100}
	Value map[string]interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GetTotalValue() *string {
	return s.TotalValue
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GetValue() map[string]interface{} {
	return s.Value
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetTotalValue(v string) *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.TotalValue = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetValue(v map[string]interface{}) *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.Value = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
