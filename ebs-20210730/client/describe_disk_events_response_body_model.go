// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskEvents(v []*DescribeDiskEventsResponseBodyDiskEvents) *DescribeDiskEventsResponseBody
	GetDiskEvents() []*DescribeDiskEventsResponseBodyDiskEvents
	SetNextToken(v string) *DescribeDiskEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDiskEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDiskEventsResponseBody
	GetTotalCount() *int64
}

type DescribeDiskEventsResponseBody struct {
	// The risk events of the disk.
	DiskEvents []*DescribeDiskEventsResponseBodyDiskEvents `json:"DiskEvents,omitempty" xml:"DiskEvents,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsResponseBody) GetDiskEvents() []*DescribeDiskEventsResponseBodyDiskEvents {
	return s.DiskEvents
}

func (s *DescribeDiskEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskEventsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDiskEventsResponseBody) SetDiskEvents(v []*DescribeDiskEventsResponseBodyDiskEvents) *DescribeDiskEventsResponseBody {
	s.DiskEvents = v
	return s
}

func (s *DescribeDiskEventsResponseBody) SetNextToken(v string) *DescribeDiskEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskEventsResponseBody) SetRequestId(v string) *DescribeDiskEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskEventsResponseBody) SetTotalCount(v int64) *DescribeDiskEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiskEventsResponseBody) Validate() error {
	if s.DiskEvents != nil {
		for _, item := range s.DiskEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskEventsResponseBodyDiskEvents struct {
	// The description of the event.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp1bq5g3dxxo1x4o****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The recommended action after the event occurred. Valid values:
	//
	// 	- Resize: resizes the disk.
	//
	// 	- ModifyDiskSpec: changes the category of the disk.
	//
	// 	- NoAction: performs no operation.
	//
	// example:
	//
	// NoAction
	RecommendAction *string `json:"RecommendAction,omitempty" xml:"RecommendAction,omitempty"`
	// The region ID of the disk.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the event. Valid values:
	//
	// 	- Solved
	//
	// 	- UnSolved
	//
	// example:
	//
	// Solved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the event occurred. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-06-01T08:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The type of the event. Only DataNeedProtect can be returned.
	//
	// example:
	//
	// DataNeedProtect
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskEventsResponseBodyDiskEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskEventsResponseBodyDiskEvents) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) GetDescription() *string {
	return s.Description
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) GetRecommendAction() *string {
	return s.RecommendAction
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) GetType() *string {
	return s.Type
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetDescription(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Description = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetDiskId(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetRecommendAction(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.RecommendAction = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetRegionId(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetStatus(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Status = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetTimestamp(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Timestamp = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetType(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Type = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) Validate() error {
	return dara.Validate(s)
}
