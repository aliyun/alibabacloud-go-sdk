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
	// The cloud desktop ID. If you do not specify a value for this parameter, events of all cloud desktops in the specified region are queried.
	//
	// example:
	//
	// ecd-8fupvkhg0aayu****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The IP address of the cloud desktop. If you do not specify a value for this parameter, the events of all cloud desktops in the specified region are queried.
	//
	// example:
	//
	// 10.10.*.*
	DesktopIp *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	// The cloud desktop name.
	//
	// example:
	//
	// test
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// This parameter is not available to the public.
	//
	// example:
	//
	// cn-hangzhou+dir-bh77qa8nmjot4****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.\\
	//
	// If you do not specify a value for this parameter, the current time is used.
	//
	// example:
	//
	// 2020-11-31T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the end user that connects to the cloud desktop from the Elastic Desktop Service (EDS) client. The information can be a Resource Access Management (RAM) user ID or an Active Directory (AD) username. If you do not specify a value for this parameter, the events of all end users in the specified region are queried.
	//
	// example:
	//
	// 28961708130834****
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The type of the events that you want to query. If you specify multiple values for the EventTypes parameter, the events of all specified types are returned. If you do not specify values for the EventTypes and EventType parameters, all events of end users in the specified region are returned.
	//
	// Valid values:
	//
	// 	- DESKTOP_STOP: End users stop the cloud desktop.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- GET_LITE_CONNECTION_TICKET: End users obtain the credential for reconnecting to the cloud desktop upon disconnection.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DESKTOP_DISCONNECT: End users disconnect desktop sessions.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- GET_CONNECTION_TICKET: End users request to connect to the cloud desktop.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CLIENT_LOGIN: End users log on to the cloud desktop.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DESKTOP_REBOOT: End users restart the cloud desktop.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DESKTOP_CONNECT: End users establish desktop sessions.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DESKTOP_START: End users start the cloud desktop.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DESKTOP_DISCONNECT
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The array of event types that you want to query. You can specify multiple event types. The response contains all or specified types of events.
	EventTypes       []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	FillHardwareInfo *bool     `json:"FillHardwareInfo,omitempty" xml:"FillHardwareInfo,omitempty"`
	Language         *string   `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of entries per page.\\
	//
	// Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the workspace to which the cloud desktop belongs. If you do not specify a value for this parameter, the events of all workspaces in the specified region are queried.
	//
	// example:
	//
	// cn-hangzhou+dir-bh77qa8nmjot4****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.\\
	//
	// If you do not specify a value for this parameter, all events that occurred before the point in time that you specify for `EndTime` are queried.
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
