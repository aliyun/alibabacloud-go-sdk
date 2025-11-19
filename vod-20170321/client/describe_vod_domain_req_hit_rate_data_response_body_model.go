// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainReqHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeVodDomainReqHitRateDataResponseBodyData) *DescribeVodDomainReqHitRateDataResponseBody
	GetData() *DescribeVodDomainReqHitRateDataResponseBodyData
	SetDataInterval(v string) *DescribeVodDomainReqHitRateDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainReqHitRateDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainReqHitRateDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainReqHitRateDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainReqHitRateDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainReqHitRateDataResponseBody struct {
	// The request hit rate data at each time interval.
	Data *DescribeVodDomainReqHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 2023-12-22T08:00:00Z
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
	// 2023-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainReqHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) GetData() *DescribeVodDomainReqHitRateDataResponseBodyData {
	return s.Data
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) SetData(v *DescribeVodDomainReqHitRateDataResponseBodyData) *DescribeVodDomainReqHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) SetDataInterval(v string) *DescribeVodDomainReqHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) SetDomainName(v string) *DescribeVodDomainReqHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) SetEndTime(v string) *DescribeVodDomainReqHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) SetRequestId(v string) *DescribeVodDomainReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) SetStartTime(v string) *DescribeVodDomainReqHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainReqHitRateDataResponseBodyData struct {
	DataModule []*DescribeVodDomainReqHitRateDataResponseBodyDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainReqHitRateDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainReqHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyData) GetDataModule() []*DescribeVodDomainReqHitRateDataResponseBodyDataDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyData) SetDataModule(v []*DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) *DescribeVodDomainReqHitRateDataResponseBodyData {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyData) Validate() error {
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

type DescribeVodDomainReqHitRateDataResponseBodyDataDataModule struct {
	// The HTTPS request hit rate.
	//
	// example:
	//
	// 50
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-21T08:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total request hit rate.
	//
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) SetHttpsValue(v string) *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) SetTimeStamp(v string) *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) SetValue(v string) *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponseBodyDataDataModule) Validate() error {
	return dara.Validate(s)
}
