// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainReqHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainReqHitRateDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainReqHitRateDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainReqHitRateDataResponseBody
	GetEndTime() *string
	SetReqHitRateInterval(v *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) *DescribeDomainReqHitRateDataResponseBody
	GetReqHitRateInterval() *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval
	SetRequestId(v string) *DescribeDomainReqHitRateDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainReqHitRateDataResponseBody
	GetStartTime() *string
}

type DescribeDomainReqHitRateDataResponseBody struct {
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
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request hit ratio data at each time interval. The hit ratio is measured in percentage.
	ReqHitRateInterval *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval `json:"ReqHitRateInterval,omitempty" xml:"ReqHitRateInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainReqHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainReqHitRateDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainReqHitRateDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainReqHitRateDataResponseBody) GetReqHitRateInterval() *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval {
	return s.ReqHitRateInterval
}

func (s *DescribeDomainReqHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainReqHitRateDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetDataInterval(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetDomainName(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetEndTime(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetReqHitRateInterval(v *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) *DescribeDomainReqHitRateDataResponseBody {
	s.ReqHitRateInterval = v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetStartTime(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) Validate() error {
	if s.ReqHitRateInterval != nil {
		if err := s.ReqHitRateInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval struct {
	DataModule []*DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) GetDataModule() []*DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) SetDataModule(v []*DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) Validate() error {
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

type DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule struct {
	// The hit ratio of HTTPS requests.
	//
	// example:
	//
	// 50.0
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2017-12-22T08:00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The request hit ratio.
	//
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetHttpsValue(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetTimeStamp(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetValue(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
