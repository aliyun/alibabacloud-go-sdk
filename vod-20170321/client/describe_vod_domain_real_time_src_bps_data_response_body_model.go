// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeSrcBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcBpsDataPerInterval(v *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeVodDomainRealTimeSrcBpsDataResponseBody
	GetRealTimeSrcBpsDataPerInterval() *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval
	SetRequestId(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainRealTimeSrcBpsDataResponseBody struct {
	DataInterval                  *string                                                                       `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                    *string                                                                       `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                       *string                                                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeSrcBpsDataPerInterval *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval,omitempty" xml:"RealTimeSrcBpsDataPerInterval,omitempty" type:"Struct"`
	RequestId                     *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                     *string                                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) GetRealTimeSrcBpsDataPerInterval() *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	return s.RealTimeSrcBpsDataPerInterval
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) SetDomainName(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) SetEndTime(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) SetRealTimeSrcBpsDataPerInterval(v *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeVodDomainRealTimeSrcBpsDataResponseBody {
	s.RealTimeSrcBpsDataPerInterval = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) SetStartTime(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval struct {
	DataModule []*DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GetDataModule() []*DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) SetDataModule(v []*DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
