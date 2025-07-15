// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeClientEventsResponseBodyEvents) *DescribeClientEventsResponseBody
	GetEvents() []*DescribeClientEventsResponseBodyEvents
	SetNextToken(v string) *DescribeClientEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeClientEventsResponseBody
	GetRequestId() *string
}

type DescribeClientEventsResponseBody struct {
	// The user events.
	Events []*DescribeClientEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 28A40F12-F340-442B-A35F-46EF6A03227B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseBody) GetEvents() []*DescribeClientEventsResponseBodyEvents {
	return s.Events
}

func (s *DescribeClientEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeClientEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientEventsResponseBody) SetEvents(v []*DescribeClientEventsResponseBodyEvents) *DescribeClientEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeClientEventsResponseBody) SetNextToken(v string) *DescribeClientEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeClientEventsResponseBody) SetRequestId(v string) *DescribeClientEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClientEventsResponseBodyEvents struct {
	// The ID of the Alibaba Cloud account with which the event is associated.
	//
	// example:
	//
	// 112259558861****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The number of bytes that are received.
	//
	// example:
	//
	// 8665
	BytesReceived *string `json:"BytesReceived,omitempty" xml:"BytesReceived,omitempty"`
	// The number of bytes that are sent.
	//
	// example:
	//
	// 2345
	BytesSend *string `json:"BytesSend,omitempty" xml:"BytesSend,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 100.68.*.*
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The OS that the client runs.
	//
	// example:
	//
	// Darwin 17.7.0 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client version.
	//
	// example:
	//
	// 1.0.4 202012021700
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The desktop group ID.
	//
	// example:
	//
	// dg-kadkdfaf****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The desktop group name.
	//
	// example:
	//
	// testName
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// The cloud desktop ID.
	//
	// example:
	//
	// ecd-8fupvkhg0aayu****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The IP address of the cloud desktop.
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
	// The ID of the directory to which the cloud desktop belongs.
	//
	// example:
	//
	// cn-hangzhou+dir-bh77qa8nmjot4****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The directory type.
	//
	// example:
	//
	// RAM
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The information about the end user that connects to the cloud desktop from the EDS client. The information can be a RAM user ID or an AD username.
	//
	// example:
	//
	// 28961708130834****
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 5651188b-3070-d1cc-5311-75753d59****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2020-11-30T06:32:31Z
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The event type. Valid values:
	//
	// example:
	//
	// DESKTOP_DISCONNECT
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the workspace to which the cloud desktop belongs.
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
	// The account type of the workspace.
	//
	// Valid values:
	//
	// 	- SIMPLE: convenience account
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- AD_CONNECTOR: enterprise AD account
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the event. If you set the EventType parameter to `DESKTOP_DISCONNECT` or `GET_CONNECTION_TICKET`, this parameter is returned. Valid values:
	//
	// 	- 200\\. The value indicates that the request is successful.
	//
	// 	- An error message. The value indicates that the request failed. Example: FailedToGetConnectionTicket.
	//
	// example:
	//
	// 200
	Status       *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	TerminalInfo *DescribeClientEventsResponseBodyEventsTerminalInfo `json:"TerminalInfo,omitempty" xml:"TerminalInfo,omitempty" type:"Struct"`
}

func (s DescribeClientEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseBodyEvents) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeClientEventsResponseBodyEvents) GetBytesReceived() *string {
	return s.BytesReceived
}

func (s *DescribeClientEventsResponseBodyEvents) GetBytesSend() *string {
	return s.BytesSend
}

func (s *DescribeClientEventsResponseBodyEvents) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeClientEventsResponseBodyEvents) GetClientOS() *string {
	return s.ClientOS
}

func (s *DescribeClientEventsResponseBodyEvents) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeClientEventsResponseBodyEvents) GetDescription() *string {
	return s.Description
}

func (s *DescribeClientEventsResponseBodyEvents) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeClientEventsResponseBodyEvents) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *DescribeClientEventsResponseBodyEvents) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeClientEventsResponseBodyEvents) GetDesktopIp() *string {
	return s.DesktopIp
}

func (s *DescribeClientEventsResponseBodyEvents) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeClientEventsResponseBodyEvents) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeClientEventsResponseBodyEvents) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *DescribeClientEventsResponseBodyEvents) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeClientEventsResponseBodyEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribeClientEventsResponseBodyEvents) GetEventTime() *string {
	return s.EventTime
}

func (s *DescribeClientEventsResponseBodyEvents) GetEventType() *string {
	return s.EventType
}

func (s *DescribeClientEventsResponseBodyEvents) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeClientEventsResponseBodyEvents) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeClientEventsResponseBodyEvents) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeClientEventsResponseBodyEvents) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClientEventsResponseBodyEvents) GetStatus() *string {
	return s.Status
}

func (s *DescribeClientEventsResponseBodyEvents) GetTerminalInfo() *DescribeClientEventsResponseBodyEventsTerminalInfo {
	return s.TerminalInfo
}

func (s *DescribeClientEventsResponseBodyEvents) SetAliUid(v string) *DescribeClientEventsResponseBodyEvents {
	s.AliUid = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetBytesReceived(v string) *DescribeClientEventsResponseBodyEvents {
	s.BytesReceived = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetBytesSend(v string) *DescribeClientEventsResponseBodyEvents {
	s.BytesSend = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientIp(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientIp = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientOS(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientOS = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientVersion(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDescription(v string) *DescribeClientEventsResponseBodyEvents {
	s.Description = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopGroupId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopGroupName(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopIp(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopName(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopName = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDirectoryId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDirectoryType(v string) *DescribeClientEventsResponseBodyEvents {
	s.DirectoryType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEndUserId(v string) *DescribeClientEventsResponseBodyEvents {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventId(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventTime(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventType(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteId(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteName(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteType(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetRegionId(v string) *DescribeClientEventsResponseBodyEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetStatus(v string) *DescribeClientEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetTerminalInfo(v *DescribeClientEventsResponseBodyEventsTerminalInfo) *DescribeClientEventsResponseBodyEvents {
	s.TerminalInfo = v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}

type DescribeClientEventsResponseBodyEventsTerminalInfo struct {
	Model        *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ProductName  *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DescribeClientEventsResponseBodyEventsTerminalInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientEventsResponseBodyEventsTerminalInfo) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseBodyEventsTerminalInfo) GetModel() *string {
	return s.Model
}

func (s *DescribeClientEventsResponseBodyEventsTerminalInfo) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeClientEventsResponseBodyEventsTerminalInfo) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeClientEventsResponseBodyEventsTerminalInfo) SetModel(v string) *DescribeClientEventsResponseBodyEventsTerminalInfo {
	s.Model = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEventsTerminalInfo) SetProductName(v string) *DescribeClientEventsResponseBodyEventsTerminalInfo {
	s.ProductName = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEventsTerminalInfo) SetSerialNumber(v string) *DescribeClientEventsResponseBodyEventsTerminalInfo {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEventsTerminalInfo) Validate() error {
	return dara.Validate(s)
}
