// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeVodDomainUsageDataResponseBody
	GetArea() *string
	SetDataInterval(v string) *DescribeVodDomainUsageDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainUsageDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainUsageDataResponseBody
	GetStartTime() *string
	SetType(v string) *DescribeVodDomainUsageDataResponseBody
	GetType() *string
	SetUsageDataPerInterval(v *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeVodDomainUsageDataResponseBody
	GetUsageDataPerInterval() *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval
}

type DescribeVodDomainUsageDataResponseBody struct {
	// The usage region.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time interval between records. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name information.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2015-12-10T12:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B955107D-E658-4E77-****-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2015-12-10T10:20:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The data type. Valid values:
	//
	// - **bps**: bandwidth.
	//
	//  - **traf**: traffic.
	//
	// example:
	//
	// bps
	Type                 *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	UsageDataPerInterval *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval `json:"UsageDataPerInterval,omitempty" xml:"UsageDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVodDomainUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponseBody) GetArea() *string {
	return s.Area
}

func (s *DescribeVodDomainUsageDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainUsageDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainUsageDataResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeVodDomainUsageDataResponseBody) GetUsageDataPerInterval() *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval {
	return s.UsageDataPerInterval
}

func (s *DescribeVodDomainUsageDataResponseBody) SetArea(v string) *DescribeVodDomainUsageDataResponseBody {
	s.Area = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetDataInterval(v string) *DescribeVodDomainUsageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetDomainName(v string) *DescribeVodDomainUsageDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetEndTime(v string) *DescribeVodDomainUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetRequestId(v string) *DescribeVodDomainUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetStartTime(v string) *DescribeVodDomainUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetType(v string) *DescribeVodDomainUsageDataResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetUsageDataPerInterval(v *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeVodDomainUsageDataResponseBody {
	s.UsageDataPerInterval = v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) Validate() error {
	if s.UsageDataPerInterval != nil {
		if err := s.UsageDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval struct {
	DataModule []*DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) GetDataModule() []*DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) SetDataModule(v []*DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) Validate() error {
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

type DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
