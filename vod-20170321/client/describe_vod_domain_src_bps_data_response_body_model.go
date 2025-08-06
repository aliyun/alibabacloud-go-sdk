// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainSrcBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainSrcBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainSrcBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainSrcBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainSrcBpsDataResponseBody
	GetRequestId() *string
	SetSrcBpsDataPerInterval(v *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) *DescribeVodDomainSrcBpsDataResponseBody
	GetSrcBpsDataPerInterval() *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval
	SetStartTime(v string) *DescribeVodDomainSrcBpsDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainSrcBpsDataResponseBody struct {
	// The time interval at which data is returned, which is the time granularity. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name whose ICP filing status you want to update.
	//
	// example:
	//
	// sample.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range.
	//
	// example:
	//
	// 2022-08-23T02:02:57Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The origin bandwidth data at each time interval.
	SrcBpsDataPerInterval *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval `json:"SrcBpsDataPerInterval,omitempty" xml:"SrcBpsDataPerInterval,omitempty" type:"Struct"`
	// The beginning of the time range.
	//
	// example:
	//
	// 2022-07-12T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainSrcBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) GetSrcBpsDataPerInterval() *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval {
	return s.SrcBpsDataPerInterval
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeVodDomainSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) SetDomainName(v string) *DescribeVodDomainSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) SetEndTime(v string) *DescribeVodDomainSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) SetSrcBpsDataPerInterval(v *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) *DescribeVodDomainSrcBpsDataResponseBody {
	s.SrcBpsDataPerInterval = v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) SetStartTime(v string) *DescribeVodDomainSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval struct {
	DataModule []*DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) GetDataModule() []*DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) SetDataModule(v []*DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule struct {
	// The bandwidth consumed for fetching resources from the origin over HTTPS.
	//
	// example:
	//
	// 0
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-23T02:02:57Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total origin bandwidth data. Unit: bit/s.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
