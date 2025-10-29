// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePushProxyUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeLivePushProxyUsageDataResponseBody
	GetEndTime() *string
	SetPushProxyData(v *DescribeLivePushProxyUsageDataResponseBodyPushProxyData) *DescribeLivePushProxyUsageDataResponseBody
	GetPushProxyData() *DescribeLivePushProxyUsageDataResponseBodyPushProxyData
	SetRequestId(v string) *DescribeLivePushProxyUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLivePushProxyUsageDataResponseBody
	GetStartTime() *string
}

type DescribeLivePushProxyUsageDataResponseBody struct {
	// The end time.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The usage data of live center stream relay.
	PushProxyData *DescribeLivePushProxyUsageDataResponseBodyPushProxyData `json:"PushProxyData,omitempty" xml:"PushProxyData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B460F8B-993C-4F48-B98A-910811DEBFEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLivePushProxyUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePushProxyUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLivePushProxyUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLivePushProxyUsageDataResponseBody) GetPushProxyData() *DescribeLivePushProxyUsageDataResponseBodyPushProxyData {
	return s.PushProxyData
}

func (s *DescribeLivePushProxyUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLivePushProxyUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLivePushProxyUsageDataResponseBody) SetEndTime(v string) *DescribeLivePushProxyUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBody) SetPushProxyData(v *DescribeLivePushProxyUsageDataResponseBodyPushProxyData) *DescribeLivePushProxyUsageDataResponseBody {
	s.PushProxyData = v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBody) SetRequestId(v string) *DescribeLivePushProxyUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBody) SetStartTime(v string) *DescribeLivePushProxyUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBody) Validate() error {
	if s.PushProxyData != nil {
		if err := s.PushProxyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLivePushProxyUsageDataResponseBodyPushProxyData struct {
	PushProxyDataItem []*DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem `json:"PushProxyDataItem,omitempty" xml:"PushProxyDataItem,omitempty" type:"Repeated"`
}

func (s DescribeLivePushProxyUsageDataResponseBodyPushProxyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePushProxyUsageDataResponseBodyPushProxyData) GoString() string {
	return s.String()
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyData) GetPushProxyDataItem() []*DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem {
	return s.PushProxyDataItem
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyData) SetPushProxyDataItem(v []*DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) *DescribeLivePushProxyUsageDataResponseBodyPushProxyData {
	s.PushProxyDataItem = v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyData) Validate() error {
	if s.PushProxyDataItem != nil {
		for _, item := range s.PushProxyDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem struct {
	// The domain name. If the value of SplitBy includes domain, the returned data is grouped by domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the region. If the value of SplitBy includes region, the returned data is grouped by region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The peak number of live center stream relay channels.
	//
	// example:
	//
	// 8
	StreamCount *int64 `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) GoString() string {
	return s.String()
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) GetRegion() *string {
	return s.Region
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) GetStreamCount() *int64 {
	return s.StreamCount
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) SetDomainName(v string) *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) SetRegion(v string) *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem {
	s.Region = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) SetStreamCount(v int64) *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem {
	s.StreamCount = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) SetTimeStamp(v string) *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponseBodyPushProxyDataPushProxyDataItem) Validate() error {
	return dara.Validate(s)
}
