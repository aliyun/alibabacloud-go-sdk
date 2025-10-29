// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDelayedStreamingUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDelayData(v *DescribeLiveDelayedStreamingUsageResponseBodyDelayData) *DescribeLiveDelayedStreamingUsageResponseBody
	GetDelayData() *DescribeLiveDelayedStreamingUsageResponseBodyDelayData
	SetEndTime(v string) *DescribeLiveDelayedStreamingUsageResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveDelayedStreamingUsageResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDelayedStreamingUsageResponseBody
	GetStartTime() *string
}

type DescribeLiveDelayedStreamingUsageResponseBody struct {
	// The details about the stream delay usage data.
	DelayData *DescribeLiveDelayedStreamingUsageResponseBodyDelayData `json:"DelayData,omitempty" xml:"DelayData,omitempty" type:"Struct"`
	// The end of the time range during which the data was queried.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B460F8B-993C-4F48-B98A-910811DEBFEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which the data was queried.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDelayedStreamingUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayedStreamingUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) GetDelayData() *DescribeLiveDelayedStreamingUsageResponseBodyDelayData {
	return s.DelayData
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) SetDelayData(v *DescribeLiveDelayedStreamingUsageResponseBodyDelayData) *DescribeLiveDelayedStreamingUsageResponseBody {
	s.DelayData = v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) SetEndTime(v string) *DescribeLiveDelayedStreamingUsageResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) SetRequestId(v string) *DescribeLiveDelayedStreamingUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) SetStartTime(v string) *DescribeLiveDelayedStreamingUsageResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBody) Validate() error {
	if s.DelayData != nil {
		if err := s.DelayData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDelayedStreamingUsageResponseBodyDelayData struct {
	DelayDataItem []*DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem `json:"DelayDataItem,omitempty" xml:"DelayDataItem,omitempty" type:"Repeated"`
}

func (s DescribeLiveDelayedStreamingUsageResponseBodyDelayData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayedStreamingUsageResponseBodyDelayData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayData) GetDelayDataItem() []*DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem {
	return s.DelayDataItem
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayData) SetDelayDataItem(v []*DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) *DescribeLiveDelayedStreamingUsageResponseBodyDelayData {
	s.DelayDataItem = v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayData) Validate() error {
	if s.DelayDataItem != nil {
		for _, item := range s.DelayDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem struct {
	// The domain name. If SplitBy is set to domain, the data returned is grouped by domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The duration of stream delay.
	//
	// example:
	//
	// 84
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the region. If SplitBy is set to region, the data returned is grouped by region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the stream. If SplitBy is set to stream, the data returned is grouped by stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) SetDomainName(v string) *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) SetDuration(v int64) *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem {
	s.Duration = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) SetRegion(v string) *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem {
	s.Region = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) SetStreamName(v string) *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) SetTimeStamp(v string) *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponseBodyDelayDataDelayDataItem) Validate() error {
	return dara.Validate(s)
}
