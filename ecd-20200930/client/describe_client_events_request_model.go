// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeClientEventsRequest
	GetDesktopId() *string
	SetDesktopIp(v string) *DescribeClientEventsRequest
	GetDesktopIp() *string
	SetDesktopName(v string) *DescribeClientEventsRequest
	GetDesktopName() *string
	SetDirectoryId(v string) *DescribeClientEventsRequest
	GetDirectoryId() *string
	SetEndTime(v string) *DescribeClientEventsRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeClientEventsRequest
	GetEndUserId() *string
	SetEndUserIds(v []*string) *DescribeClientEventsRequest
	GetEndUserIds() []*string
	SetEventType(v string) *DescribeClientEventsRequest
	GetEventType() *string
	SetEventTypes(v []*string) *DescribeClientEventsRequest
	GetEventTypes() []*string
	SetFillHardwareInfo(v bool) *DescribeClientEventsRequest
	GetFillHardwareInfo() *bool
	SetLanguage(v string) *DescribeClientEventsRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeClientEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeClientEventsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeClientEventsRequest
	GetOfficeSiteId() *string
	SetOfficeSiteName(v string) *DescribeClientEventsRequest
	GetOfficeSiteName() *string
	SetRegionId(v string) *DescribeClientEventsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeClientEventsRequest
	GetStartTime() *string
}

type DescribeClientEventsRequest struct {
	// The cloud computer ID. If you do not specify this parameter, all cloud computers in the region are queried.
	//
	// example:
	//
	// ecd-8fupvkhg0aayu****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The IP address of the cloud computer. If you do not specify this parameter, events of all cloud computers in the region are queried.
	//
	// example:
	//
	// 10.10.*.*
	DesktopIp *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// test
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The end time. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC+0. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 2020-11-31T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The logon user information, which is a Resource Access Management (RAM) user ID or AD username. If you do not specify this parameter, events of all users in the region are queried.
	//
	// example:
	//
	// 28961708130834****
	EndUserId  *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The event type to query. If EventTypes is not empty, the EventTypes combination is used as the query filter condition. If both EventTypes and EventType are empty, all events are queried.
	//
	// example:
	//
	// DESKTOP_DISCONNECT
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The combination of event types to query. You can specify multiple event types. The query results include events of all specified types.
	EventTypes       []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	FillHardwareInfo *bool     `json:"FillHardwareInfo,omitempty" xml:"FillHardwareInfo,omitempty"`
	Language         *string   `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of entries per page for a paged query. Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of NextToken returned in the previous API call.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the office network to which the cloud computer belongs. If you do not specify this parameter, user events in all office networks in the region are queried.
	//
	// example:
	//
	// cn-hangzhou+dir-bh77qa8nmjot4****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The name of the office network.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC+0. If you do not specify this parameter, events are queried backward from the time specified by `EndTime`.
	//
	// example:
	//
	// 2020-11-30T06:32:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClientEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeClientEventsRequest) GetDesktopIp() *string {
	return s.DesktopIp
}

func (s *DescribeClientEventsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeClientEventsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeClientEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeClientEventsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeClientEventsRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeClientEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeClientEventsRequest) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *DescribeClientEventsRequest) GetFillHardwareInfo() *bool {
	return s.FillHardwareInfo
}

func (s *DescribeClientEventsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeClientEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeClientEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeClientEventsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeClientEventsRequest) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeClientEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClientEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeClientEventsRequest) SetDesktopId(v string) *DescribeClientEventsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopIp(v string) *DescribeClientEventsRequest {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopName(v string) *DescribeClientEventsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDirectoryId(v string) *DescribeClientEventsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndTime(v string) *DescribeClientEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndUserId(v string) *DescribeClientEventsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndUserIds(v []*string) *DescribeClientEventsRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeClientEventsRequest) SetEventType(v string) *DescribeClientEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEventTypes(v []*string) *DescribeClientEventsRequest {
	s.EventTypes = v
	return s
}

func (s *DescribeClientEventsRequest) SetFillHardwareInfo(v bool) *DescribeClientEventsRequest {
	s.FillHardwareInfo = &v
	return s
}

func (s *DescribeClientEventsRequest) SetLanguage(v string) *DescribeClientEventsRequest {
	s.Language = &v
	return s
}

func (s *DescribeClientEventsRequest) SetMaxResults(v int32) *DescribeClientEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeClientEventsRequest) SetNextToken(v string) *DescribeClientEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeClientEventsRequest) SetOfficeSiteId(v string) *DescribeClientEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetOfficeSiteName(v string) *DescribeClientEventsRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeClientEventsRequest) SetRegionId(v string) *DescribeClientEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetStartTime(v string) *DescribeClientEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClientEventsRequest) Validate() error {
	return dara.Validate(s)
}
