// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeLiveDomainTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveDomainTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainTrafficDataResponseBody
	GetStartTime() *string
	SetTrafficDataPerInterval(v *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeLiveDomainTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeLiveDomainTrafficDataResponseBody struct {
	// The time granularity of the query.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-10T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-10T14:00:00Z
	StartTime              *string                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeLiveDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeLiveDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeLiveDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeLiveDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeLiveDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeLiveDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBody) Validate() error {
	if s.TrafficDataPerInterval != nil {
		if err := s.TrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
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

type DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	HttpTrafficValue  *string `json:"HttpTrafficValue,omitempty" xml:"HttpTrafficValue,omitempty"`
	HttpsTrafficValue *string `json:"HttpsTrafficValue,omitempty" xml:"HttpsTrafficValue,omitempty"`
	TimeStamp         *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrafficValue      *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpTrafficValue() *string {
	return s.HttpTrafficValue
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetHttpsTrafficValue() *string {
	return s.HttpsTrafficValue
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTrafficValue() *string {
	return s.TrafficValue
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpTrafficValue(v string) *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpTrafficValue = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsTrafficValue(v string) *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsTrafficValue = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
