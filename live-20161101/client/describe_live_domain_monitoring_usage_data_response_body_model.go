// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMonitoringUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody
	GetEndTime() *string
	SetInstanceId(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody
	GetInstanceId() *string
	SetMonitoringData(v *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData) *DescribeLiveDomainMonitoringUsageDataResponseBody
	GetMonitoringData() *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData
	SetRegion(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainMonitoringUsageDataResponseBody struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range.
	//
	// example:
	//
	// 2022-12-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the monitoring session.
	//
	// example:
	//
	// e62af24d-a354-3b0c-9f1f-da592c4b****
	InstanceId     *string                                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MonitoringData *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData `json:"MonitoringData,omitempty" xml:"MonitoringData,omitempty" type:"Struct"`
	// The region of the live center.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range.
	//
	// example:
	//
	// 2022-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainMonitoringUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMonitoringUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) GetMonitoringData() *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData {
	return s.MonitoringData
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) SetDomainName(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) SetEndTime(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) SetInstanceId(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) SetMonitoringData(v *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData) *DescribeLiveDomainMonitoringUsageDataResponseBody {
	s.MonitoringData = v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) SetRegion(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) SetRequestId(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) SetStartTime(v string) *DescribeLiveDomainMonitoringUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBody) Validate() error {
	if s.MonitoringData != nil {
		if err := s.MonitoringData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData struct {
	MonitoringDataItem []*DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem `json:"MonitoringDataItem,omitempty" xml:"MonitoringDataItem,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData) GetMonitoringDataItem() []*DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem {
	return s.MonitoringDataItem
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData) SetMonitoringDataItem(v []*DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData {
	s.MonitoringDataItem = v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringData) Validate() error {
	if s.MonitoringDataItem != nil {
		for _, item := range s.MonitoringDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Duration   *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) GetResolution() *string {
	return s.Resolution
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) SetDomainName(v string) *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) SetDuration(v int32) *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem {
	s.Duration = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) SetInstanceId(v string) *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem {
	s.InstanceId = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) SetRegion(v string) *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) SetResolution(v string) *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem {
	s.Resolution = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) SetTimeStamp(v string) *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponseBodyMonitoringDataMonitoringDataItem) Validate() error {
	return dara.Validate(s)
}
