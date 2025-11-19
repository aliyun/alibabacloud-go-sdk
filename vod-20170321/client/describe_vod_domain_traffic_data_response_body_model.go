// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainTrafficDataResponseBody
	GetStartTime() *string
	SetTotalTraffic(v string) *DescribeVodDomainTrafficDataResponseBody
	GetTotalTraffic() *string
	SetTrafficDataPerInterval(v *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeVodDomainTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeVodDomainTrafficDataResponseBody struct {
	// The time interval at which data is returned, which is the time granularity. Unit: seconds.
	//
	// example:
	//
	// 3600
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range.
	//
	// example:
	//
	// 2019-01-20T14:59:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D94E471F-1A27-442E-552D-D4D2000C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range.
	//
	// example:
	//
	// 2019-01-20T13:59:58Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total amount of network traffic.
	//
	// example:
	//
	// 5906662826
	TotalTraffic *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	// The amount of network traffic at each time interval.
	TrafficDataPerInterval *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVodDomainTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainTrafficDataResponseBody) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeVodDomainTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetTotalTraffic(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeVodDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) Validate() error {
	if s.TrafficDataPerInterval != nil {
		if err := s.TrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
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

type DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	// The amount of network traffic in the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 0
	DomesticValue *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	// The amount of HTTPS network traffic on points of presence (POPs) in the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 0
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	// The amount of HTTPS network traffic on POPs outside the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 0
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	// The total amount of HTTPS network traffic on POPs. Unit: bytes.
	//
	// example:
	//
	// 0
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The amount of network traffic outside the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 0
	OverseasValue *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-15T19:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total traffic. Unit: bytes.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpsDomesticValue() *string {
	return s.HttpsDomesticValue
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpsOverseasValue() *string {
	return s.HttpsOverseasValue
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
