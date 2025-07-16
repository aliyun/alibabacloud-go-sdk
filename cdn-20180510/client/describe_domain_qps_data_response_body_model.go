// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainQpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainQpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainQpsDataResponseBody
	GetEndTime() *string
	SetQpsDataInterval(v *DescribeDomainQpsDataResponseBodyQpsDataInterval) *DescribeDomainQpsDataResponseBody
	GetQpsDataInterval() *DescribeDomainQpsDataResponseBodyQpsDataInterval
	SetRequestId(v string) *DescribeDomainQpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainQpsDataResponseBody
	GetStartTime() *string
}

type DescribeDomainQpsDataResponseBody struct {
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
	// The list of QPS records at each interval.
	QpsDataInterval *DescribeDomainQpsDataResponseBodyQpsDataInterval `json:"QpsDataInterval,omitempty" xml:"QpsDataInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B8333EDB-4595-46E0-AFE9-29BAA290D0E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainQpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainQpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainQpsDataResponseBody) GetQpsDataInterval() *DescribeDomainQpsDataResponseBodyQpsDataInterval {
	return s.QpsDataInterval
}

func (s *DescribeDomainQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainQpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainQpsDataResponseBody) SetDataInterval(v string) *DescribeDomainQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetDomainName(v string) *DescribeDomainQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetEndTime(v string) *DescribeDomainQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetQpsDataInterval(v *DescribeDomainQpsDataResponseBodyQpsDataInterval) *DescribeDomainQpsDataResponseBody {
	s.QpsDataInterval = v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetRequestId(v string) *DescribeDomainQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetStartTime(v string) *DescribeDomainQpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainQpsDataResponseBodyQpsDataInterval struct {
	DataModule []*DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsDataResponseBodyQpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBodyQpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataInterval) GetDataModule() []*DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataInterval) SetDataModule(v []*DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) *DescribeDomainQpsDataResponseBodyQpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule struct {
	// The number of requests in the Chinese mainland.
	//
	// example:
	//
	// 0
	AccDomesticValue *string `json:"AccDomesticValue,omitempty" xml:"AccDomesticValue,omitempty"`
	// The number of requests outside the Chinese mainland.
	//
	// example:
	//
	// 0
	AccOverseasValue *string `json:"AccOverseasValue,omitempty" xml:"AccOverseasValue,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 0
	AccValue *string `json:"AccValue,omitempty" xml:"AccValue,omitempty"`
	// The number of queries per second in the Chinese mainland.
	//
	// example:
	//
	// 0
	DomesticValue *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	// The number of HTTPS requests sent to POPs in the Chinese mainland.
	//
	// example:
	//
	// 1
	HttpsAccDomesticValue *string `json:"HttpsAccDomesticValue,omitempty" xml:"HttpsAccDomesticValue,omitempty"`
	// The number of HTTPS requests sent to POPs outside the Chinese mainland.
	//
	// example:
	//
	// 1
	HttpsAccOverseasValue *string `json:"HttpsAccOverseasValue,omitempty" xml:"HttpsAccOverseasValue,omitempty"`
	// The number of HTTPS requests sent to POPs.
	//
	// example:
	//
	// 1
	HttpsAccValue *string `json:"HttpsAccValue,omitempty" xml:"HttpsAccValue,omitempty"`
	// The number of queries per second that is calculated based on the HTTPS requests sent to POPs in the Chinese mainland.
	//
	// example:
	//
	// 1
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	// The number of queries per second that is calculated based on the HTTPS requests sent to POPs outside the Chinese mainland.
	//
	// example:
	//
	// 1
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	// The number of queries per second that is calculated based on the HTTPS requests sent to points of presence (POPs).
	//
	// example:
	//
	// 1
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The number of queries per second outside the Chinese mainland.
	//
	// example:
	//
	// 0
	OverseasValue *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total QPS.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetAccDomesticValue() *string {
	return s.AccDomesticValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetAccOverseasValue() *string {
	return s.AccOverseasValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetAccValue() *string {
	return s.AccValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsAccDomesticValue() *string {
	return s.HttpsAccDomesticValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsAccOverseasValue() *string {
	return s.HttpsAccOverseasValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsAccValue() *string {
	return s.HttpsAccValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsDomesticValue() *string {
	return s.HttpsDomesticValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsOverseasValue() *string {
	return s.HttpsOverseasValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
