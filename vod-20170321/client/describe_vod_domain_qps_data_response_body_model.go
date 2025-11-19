// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainQpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainQpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainQpsDataResponseBody
	GetEndTime() *string
	SetQpsDataInterval(v *DescribeVodDomainQpsDataResponseBodyQpsDataInterval) *DescribeVodDomainQpsDataResponseBody
	GetQpsDataInterval() *DescribeVodDomainQpsDataResponseBodyQpsDataInterval
	SetRequestId(v string) *DescribeVodDomainQpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainQpsDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainQpsDataResponseBody struct {
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
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the *YYYY-MM-DD**Thh:mm:ss	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-05-02T15:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of QPS records at each interval.
	QpsDataInterval *DescribeVodDomainQpsDataResponseBodyQpsDataInterval `json:"QpsDataInterval,omitempty" xml:"QpsDataInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried. The time follows the ISO 8601 standard in the *YYYY-MM-DD**Thh:mm:ss	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-05-02T15:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainQpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainQpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainQpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainQpsDataResponseBody) GetQpsDataInterval() *DescribeVodDomainQpsDataResponseBodyQpsDataInterval {
	return s.QpsDataInterval
}

func (s *DescribeVodDomainQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainQpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainQpsDataResponseBody) SetDataInterval(v string) *DescribeVodDomainQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBody) SetDomainName(v string) *DescribeVodDomainQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBody) SetEndTime(v string) *DescribeVodDomainQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBody) SetQpsDataInterval(v *DescribeVodDomainQpsDataResponseBodyQpsDataInterval) *DescribeVodDomainQpsDataResponseBody {
	s.QpsDataInterval = v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBody) SetStartTime(v string) *DescribeVodDomainQpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBody) Validate() error {
	if s.QpsDataInterval != nil {
		if err := s.QpsDataInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainQpsDataResponseBodyQpsDataInterval struct {
	DataModule []*DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainQpsDataResponseBodyQpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainQpsDataResponseBodyQpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataInterval) GetDataModule() []*DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataInterval) SetDataModule(v []*DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) *DescribeVodDomainQpsDataResponseBodyQpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataInterval) Validate() error {
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

type DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule struct {
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
	// The QPS data in the Chinese mainland.
	//
	// example:
	//
	// 0
	DomesticValue *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	// The QPS that is calculated based on the HTTPS requests sent to POPs in the Chinese mainland.
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
	// The QPS that is calculated based on the HTTPS requests sent to POPs in the Chinese mainland.
	//
	// example:
	//
	// 1
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	// The QPS that is calculated based on the HTTPS requests sent to POPs outside the Chinese mainland.
	//
	// example:
	//
	// 1
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	// The QPS that is calculated based on the HTTPS requests sent to points of presence (POPs).
	//
	// example:
	//
	// 1
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The QPS data outside the Chinese mainland.
	//
	// example:
	//
	// 0
	OverseasValue *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-27 10:10:58
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total QPS.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetAccDomesticValue() *string {
	return s.AccDomesticValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetAccOverseasValue() *string {
	return s.AccOverseasValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetAccValue() *string {
	return s.AccValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsAccDomesticValue() *string {
	return s.HttpsAccDomesticValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsAccOverseasValue() *string {
	return s.HttpsAccOverseasValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsAccValue() *string {
	return s.HttpsAccValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsDomesticValue() *string {
	return s.HttpsDomesticValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsOverseasValue() *string {
	return s.HttpsOverseasValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccDomesticValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccDomesticValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccOverseasValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccOverseasValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetDomesticValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccDomesticValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccDomesticValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccOverseasValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccOverseasValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetOverseasValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetValue(v string) *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponseBodyQpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
