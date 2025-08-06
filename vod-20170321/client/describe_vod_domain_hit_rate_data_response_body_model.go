// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainHitRateDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainHitRateDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainHitRateDataResponseBody
	GetEndTime() *string
	SetHitRateInterval(v *DescribeVodDomainHitRateDataResponseBodyHitRateInterval) *DescribeVodDomainHitRateDataResponseBody
	GetHitRateInterval() *DescribeVodDomainHitRateDataResponseBodyHitRateInterval
	SetRequestId(v string) *DescribeVodDomainHitRateDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainHitRateDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainHitRateDataResponseBody struct {
	// The time interval at which data is returned, which is the time granularity. Unit: seconds.
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
	// The end of the time range.
	//
	// example:
	//
	// 2024-01-20T14:59:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The byte hit rate data at each time interval.
	HitRateInterval *DescribeVodDomainHitRateDataResponseBodyHitRateInterval `json:"HitRateInterval,omitempty" xml:"HitRateInterval,omitempty" type:"Struct"`
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
	// 2024-01-20T13:59:58Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHitRateDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainHitRateDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainHitRateDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainHitRateDataResponseBody) GetHitRateInterval() *DescribeVodDomainHitRateDataResponseBodyHitRateInterval {
	return s.HitRateInterval
}

func (s *DescribeVodDomainHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainHitRateDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainHitRateDataResponseBody) SetDataInterval(v string) *DescribeVodDomainHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBody) SetDomainName(v string) *DescribeVodDomainHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBody) SetEndTime(v string) *DescribeVodDomainHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBody) SetHitRateInterval(v *DescribeVodDomainHitRateDataResponseBodyHitRateInterval) *DescribeVodDomainHitRateDataResponseBody {
	s.HitRateInterval = v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBody) SetRequestId(v string) *DescribeVodDomainHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBody) SetStartTime(v string) *DescribeVodDomainHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainHitRateDataResponseBodyHitRateInterval struct {
	DataModule []*DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainHitRateDataResponseBodyHitRateInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHitRateDataResponseBodyHitRateInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateInterval) GetDataModule() []*DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateInterval) SetDataModule(v []*DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) *DescribeVodDomainHitRateDataResponseBodyHitRateInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule struct {
	// The HTTPS byte hit rate.
	//
	// example:
	//
	// 50
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-20T13:59:58Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total byte hit rate.
	//
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetValue(v string) *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponseBodyHitRateIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
