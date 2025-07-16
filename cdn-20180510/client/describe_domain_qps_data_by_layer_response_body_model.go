// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainQpsDataByLayerResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainQpsDataByLayerResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainQpsDataByLayerResponseBody
	GetEndTime() *string
	SetLayer(v string) *DescribeDomainQpsDataByLayerResponseBody
	GetLayer() *string
	SetQpsDataInterval(v *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) *DescribeDomainQpsDataByLayerResponseBody
	GetQpsDataInterval() *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval
	SetRequestId(v string) *DescribeDomainQpsDataByLayerResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainQpsDataByLayerResponseBody
	GetStartTime() *string
}

type DescribeDomainQpsDataByLayerResponseBody struct {
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
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The layer at which the data was collected.
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The number of queries per second at each interval.
	QpsDataInterval *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval `json:"QpsDataInterval,omitempty" xml:"QpsDataInterval,omitempty" type:"Struct"`
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
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainQpsDataByLayerResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainQpsDataByLayerResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainQpsDataByLayerResponseBody) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDomainQpsDataByLayerResponseBody) GetQpsDataInterval() *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval {
	return s.QpsDataInterval
}

func (s *DescribeDomainQpsDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainQpsDataByLayerResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetDataInterval(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetDomainName(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetEndTime(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetLayer(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.Layer = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetQpsDataInterval(v *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) *DescribeDomainQpsDataByLayerResponseBody {
	s.QpsDataInterval = v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetStartTime(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval struct {
	DataModule []*DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) GetDataModule() []*DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) SetDataModule(v []*DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule struct {
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

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetAccDomesticValue() *string {
	return s.AccDomesticValue
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetAccOverseasValue() *string {
	return s.AccOverseasValue
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetAccValue() *string {
	return s.AccValue
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccDomesticValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccOverseasValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetDomesticValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetOverseasValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
