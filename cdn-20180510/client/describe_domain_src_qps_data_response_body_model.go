// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainSrcQpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainSrcQpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcQpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainSrcQpsDataResponseBody
	GetRequestId() *string
	SetSrcQpsDataPerInterval(v *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) *DescribeDomainSrcQpsDataResponseBody
	GetSrcQpsDataPerInterval() *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval
	SetStartTime(v string) *DescribeDomainSrcQpsDataResponseBody
	GetStartTime() *string
}

type DescribeDomainSrcQpsDataResponseBody struct {
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
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7CBCD6AD-B016-42E5-AE0B-B3731DE8F755
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The back-to-origin bandwidth information at each interval.
	SrcQpsDataPerInterval *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval `json:"SrcQpsDataPerInterval,omitempty" xml:"SrcQpsDataPerInterval,omitempty" type:"Struct"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainSrcQpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcQpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSrcQpsDataResponseBody) GetSrcQpsDataPerInterval() *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval {
	return s.SrcQpsDataPerInterval
}

func (s *DescribeDomainSrcQpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetDomainName(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetEndTime(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetRequestId(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetSrcQpsDataPerInterval(v *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) *DescribeDomainSrcQpsDataResponseBody {
	s.SrcQpsDataPerInterval = v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetStartTime(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval struct {
	DataModule []*DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) GetDataModule() []*DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) SetDataModule(v []*DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The QPS value.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
