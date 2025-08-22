// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainHitRateDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainHitRateDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainHitRateDataResponseBody
	GetEndTime() *string
	SetHitRatePerInterval(v *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) *DescribeDcdnDomainHitRateDataResponseBody
	GetHitRatePerInterval() *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval
	SetRequestId(v string) *DescribeDcdnDomainHitRateDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainHitRateDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainHitRateDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
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
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2018-03-02T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The byte hit ratio at each time interval. The byte hit ratio is measured in percentage.
	HitRatePerInterval *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval `json:"HitRatePerInterval,omitempty" xml:"HitRatePerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4D07ABFE-4737-4834-B1B9-A661308C47B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2018-03-02T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) GetHitRatePerInterval() *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval {
	return s.HitRatePerInterval
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetHitRatePerInterval(v *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) *DescribeDcdnDomainHitRateDataResponseBody {
	s.HitRatePerInterval = v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval struct {
	DataModule []*DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) GetDataModule() []*DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) SetDataModule(v []*DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule struct {
	// The byte hit ratio.
	//
	// example:
	//
	// 0
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	// The request hit ratio.
	//
	// example:
	//
	// 0
	ReqHitRate *float32 `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2018-03-02T13:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) GetByteHitRate() *float32 {
	return s.ByteHitRate
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) GetReqHitRate() *float32 {
	return s.ReqHitRate
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetByteHitRate(v float32) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetReqHitRate(v float32) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
