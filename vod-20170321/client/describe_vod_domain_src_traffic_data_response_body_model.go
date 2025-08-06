// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainSrcTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainSrcTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainSrcTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainSrcTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainSrcTrafficDataResponseBody
	GetRequestId() *string
	SetSrcTrafficDataPerInterval(v *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) *DescribeVodDomainSrcTrafficDataResponseBody
	GetSrcTrafficDataPerInterval() *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval
	SetStartTime(v string) *DescribeVodDomainSrcTrafficDataResponseBody
	GetStartTime() *string
	SetTotalTraffic(v string) *DescribeVodDomainSrcTrafficDataResponseBody
	GetTotalTraffic() *string
}

type DescribeVodDomainSrcTrafficDataResponseBody struct {
	// The time interval between the entries returned. Unit: seconds.
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
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-23T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the origin traffic returned at each time interval. Unit: bytes.
	SrcTrafficDataPerInterval *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval `json:"SrcTrafficDataPerInterval,omitempty" xml:"SrcTrafficDataPerInterval,omitempty" type:"Struct"`
	// The start of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-29T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total traffic. Unit: bytes.
	//
	// example:
	//
	// 5906662826
	TotalTraffic *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeVodDomainSrcTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) GetSrcTrafficDataPerInterval() *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval {
	return s.SrcTrafficDataPerInterval
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeVodDomainSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeVodDomainSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeVodDomainSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeVodDomainSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) SetSrcTrafficDataPerInterval(v *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) *DescribeVodDomainSrcTrafficDataResponseBody {
	s.SrcTrafficDataPerInterval = v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeVodDomainSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) SetTotalTraffic(v string) *DescribeVodDomainSrcTrafficDataResponseBody {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval struct {
	DataModule []*DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) GetDataModule() []*DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) SetDataModule(v []*DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule struct {
	// The amount of traffic generated by origin HTTPS requests.
	//
	// example:
	//
	// 0
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the returned data. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-23T15:59:59Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The traffic value at each time interval.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
