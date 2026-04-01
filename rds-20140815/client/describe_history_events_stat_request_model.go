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
	SetSecurityToken(v string) *DescribeHistoryEventsStatRequest
	GetSecurityToken() *string
	SetToStartTime(v string) *DescribeHistoryEventsStatRequest
	GetToStartTime() *string
}

type DescribeHistoryEventsStatRequest struct {
	// The status of the asset instance. Valid values: **starting**, **running**, **stopping**, and **stopped**.
	//
	// example:
	//
	// starting
	ArchiveStatus *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	// The beginning of the time range to query. Only tasks that have a start time later than or equal to the time specified by this parameter are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The start time can be up to 30 days earlier than the current time. If you set this parameter to a time more than 30 days earlier than the current time, this time is automatically converted to a time that is exactly 30 days earlier than the current time.
	//
	// example:
	//
	// 2022-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The end of the time range to query. Only tasks that have a start time earlier than or equal to the time specified by this parameter are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-03-02T11:31:03Z
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

func (s *DescribeHistoryEventsStatRequest) GetSecurityToken() *string {
	return s.SecurityToken
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

func (s *DescribeHistoryEventsStatRequest) SetSecurityToken(v string) *DescribeHistoryEventsStatRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetToStartTime(v string) *DescribeHistoryEventsStatRequest {
	s.ToStartTime = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) Validate() error {
	return dara.Validate(s)
}
