// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainQpsDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody
	GetEndTime() *string
	SetLayer(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody
	GetLayer() *string
	SetQpsDataInterval(v *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval) *DescribeDcdnDomainQpsDataByLayerResponseBody
	GetQpsDataInterval() *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval
	SetRequestId(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainQpsDataByLayerResponseBody struct {
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
	// The layer at which the data was collected.
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The QPS returned at each time interval.
	QpsDataInterval *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval `json:"QpsDataInterval,omitempty" xml:"QpsDataInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainQpsDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) GetQpsDataInterval() *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval {
	return s.QpsDataInterval
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) SetDataInterval(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) SetDomainName(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) SetEndTime(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) SetLayer(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody {
	s.Layer = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) SetQpsDataInterval(v *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval) *DescribeDcdnDomainQpsDataByLayerResponseBody {
	s.QpsDataInterval = v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) SetRequestId(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) SetStartTime(v string) *DescribeDcdnDomainQpsDataByLayerResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval struct {
	DataModule []*DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval) GetDataModule() []*DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval) SetDataModule(v []*DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule struct {
	// The number of requests in the Chinese mainland.
	//
	// example:
	//
	// 12
	AccDomesticValue *string `json:"AccDomesticValue,omitempty" xml:"AccDomesticValue,omitempty"`
	// The number of requests outside the Chinese mainland.
	//
	// example:
	//
	// 44
	AccOverseasValue *string `json:"AccOverseasValue,omitempty" xml:"AccOverseasValue,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 56
	AccValue *string `json:"AccValue,omitempty" xml:"AccValue,omitempty"`
	// The number of queries per second in the Chinese mainland.
	//
	// example:
	//
	// 0.12
	DomesticValue *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	// The number of queries per second outside the Chinese mainland.
	//
	// example:
	//
	// 0.44
	OverseasValue *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total number of queries per second.
	//
	// example:
	//
	// 0.56
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetAccDomesticValue() *string {
	return s.AccDomesticValue
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetAccOverseasValue() *string {
	return s.AccOverseasValue
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetAccValue() *string {
	return s.AccValue
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccDomesticValue(v string) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccDomesticValue = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccOverseasValue(v string) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccOverseasValue = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccValue(v string) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetDomesticValue(v string) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetOverseasValue(v string) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetValue(v string) *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
