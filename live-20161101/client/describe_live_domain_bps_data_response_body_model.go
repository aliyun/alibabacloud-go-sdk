// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeLiveDomainBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeLiveDomainBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveDomainBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainBpsDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	// The time granularity of the query. Unit: seconds.
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
	// 2017-12-10T09:00:00Z
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
	// 2017-12-10T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataResponseBody) GetBpsDataPerInterval() *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeLiveDomainBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeLiveDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBody) SetDomainName(v string) *DescribeLiveDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBody) SetEndTime(v string) *DescribeLiveDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBody) SetRequestId(v string) *DescribeLiveDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBody) SetStartTime(v string) *DescribeLiveDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBody) Validate() error {
	if s.BpsDataPerInterval != nil {
		if err := s.BpsDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerInterval) Validate() error {
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

type DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	BpsValue      *string `json:"BpsValue,omitempty" xml:"BpsValue,omitempty"`
	HttpBpsValue  *string `json:"HttpBpsValue,omitempty" xml:"HttpBpsValue,omitempty"`
	HttpsBpsValue *string `json:"HttpsBpsValue,omitempty" xml:"HttpsBpsValue,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetBpsValue() *string {
	return s.BpsValue
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpBpsValue() *string {
	return s.HttpBpsValue
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpsBpsValue() *string {
	return s.HttpsBpsValue
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetBpsValue(v string) *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.BpsValue = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpBpsValue(v string) *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpBpsValue = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsBpsValue(v string) *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsBpsValue = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
