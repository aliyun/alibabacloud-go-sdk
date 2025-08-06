// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeSrcTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcTrafficDataPerInterval(v *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody
	GetRealTimeSrcTrafficDataPerInterval() *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval
	SetRequestId(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainRealTimeSrcTrafficDataResponseBody struct {
	DataInterval                      *string                                                                               `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                        *string                                                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                           *string                                                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeSrcTrafficDataPerInterval *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval `json:"RealTimeSrcTrafficDataPerInterval,omitempty" xml:"RealTimeSrcTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                         *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                         *string                                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) GetRealTimeSrcTrafficDataPerInterval() *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	return s.RealTimeSrcTrafficDataPerInterval
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) SetRealTimeSrcTrafficDataPerInterval(v *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody {
	s.RealTimeSrcTrafficDataPerInterval = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval struct {
	DataModule []*DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GetDataModule() []*DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) SetDataModule(v []*DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
