// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveStatus(v string) *DescribeHistoryEventsStatRequest
	GetArchiveStatus() *string
	SetFromStartTime(v string) *DescribeHistoryEventsStatRequest
	GetFromStartTime() *string
	SetRegionId(v string) *DescribeHistoryEventsStatRequest
	GetRegionId() *string
	SetToStartTime(v string) *DescribeHistoryEventsStatRequest
	GetToStartTime() *string
}

type DescribeHistoryEventsStatRequest struct {
	// The archive status of the events. Valid values:
	//
	// - Archived: The events are archived.
	//
	// - UnArchived: The events are not archived.
	//
	// - All: All events.
	//
	// example:
	//
	// All
	ArchiveStatus *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	// The start time of the time range for task start times to query, in the YYYY-MM-DDThh:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2025-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The ID of the region where the events are located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end time of the time range for task start times to query, in the YYYY-MM-DDThh:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2025-01-03T12:31:03Z
	ToStartTime *string `json:"ToStartTime,omitempty" xml:"ToStartTime,omitempty"`
}

func (s DescribeHistoryEventsStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsStatRequest) GetArchiveStatus() *string {
	return s.ArchiveStatus
}

func (s *DescribeHistoryEventsStatRequest) GetFromStartTime() *string {
	return s.FromStartTime
}

func (s *DescribeHistoryEventsStatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryEventsStatRequest) GetToStartTime() *string {
	return s.ToStartTime
}

func (s *DescribeHistoryEventsStatRequest) SetArchiveStatus(v string) *DescribeHistoryEventsStatRequest {
	s.ArchiveStatus = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetFromStartTime(v string) *DescribeHistoryEventsStatRequest {
	s.FromStartTime = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetRegionId(v string) *DescribeHistoryEventsStatRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetToStartTime(v string) *DescribeHistoryEventsStatRequest {
	s.ToStartTime = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) Validate() error {
	return dara.Validate(s)
}
