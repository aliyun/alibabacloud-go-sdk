// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDesktopSessionsResponseBody
	GetRequestId() *string
	SetSessions(v []*DescribeDesktopSessionsResponseBodySessions) *DescribeDesktopSessionsResponseBody
	GetSessions() []*DescribeDesktopSessionsResponseBodySessions
	SetTotalCount(v int32) *DescribeDesktopSessionsResponseBody
	GetTotalCount() *int32
}

type DescribeDesktopSessionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3EC4A6DB-EC8D-55B0-9038-543DE671****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of sessions.
	Sessions []*DescribeDesktopSessionsResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDesktopSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopSessionsResponseBody) GetSessions() []*DescribeDesktopSessionsResponseBodySessions {
	return s.Sessions
}

func (s *DescribeDesktopSessionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDesktopSessionsResponseBody) SetRequestId(v string) *DescribeDesktopSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBody) SetSessions(v []*DescribeDesktopSessionsResponseBodySessions) *DescribeDesktopSessionsResponseBody {
	s.Sessions = v
	return s
}

func (s *DescribeDesktopSessionsResponseBody) SetTotalCount(v int32) *DescribeDesktopSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopSessionsResponseBodySessions struct {
	// The IP address of the client.
	//
	// example:
	//
	// 172.21.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The client OS.
	//
	// example:
	//
	// windows_\\"Windows10Enterprise\\"10.0(Build22000)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client version.
	//
	// example:
	//
	// 2.0.0-R-20221030.08****
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-g6t1ukbaea****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// testDesktop
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The duration of the remote assistance. Unit: seconds.
	//
	// example:
	//
	// 120
	EndUserApplyCoordinateTime *int64 `json:"EndUserApplyCoordinateTime,omitempty" xml:"EndUserApplyCoordinateTime,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// testUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The duration of the last connection to the cloud computer. Unit: seconds.
	//
	// example:
	//
	// 120
	LatestConnectionTime *int64 `json:"LatestConnectionTime,omitempty" xml:"LatestConnectionTime,omitempty"`
	// The ID of the office network.
	//
	// example:
	//
	// cn-hangzhou+dir-8904****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The name of the office network.
	//
	// example:
	//
	// DemoOfficeSite
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// Indicates whether the switch to check session status of cloud computers is turned on.
	//
	// example:
	//
	// true
	OsSessionStatus *string `json:"OsSessionStatus,omitempty" xml:"OsSessionStatus,omitempty"`
	// The OS.
	//
	// Valid values:
	//
	// 	- Linux
	//
	// 	- Windows
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- HDX
	//
	// 	- ASP
	//
	// example:
	//
	// ASP
	ProtocolType   *string                                                      `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ResourceGroups []*DescribeDesktopSessionsResponseBodySessionsResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// The end time of the session.
	//
	// example:
	//
	// 2023-01-28T02:31:43Z
	SessionEndTime *string `json:"SessionEndTime,omitempty" xml:"SessionEndTime,omitempty"`
	// The idle duration of the session. Unit: seconds.
	//
	// example:
	//
	// 120
	SessionIdleTime *int64 `json:"SessionIdleTime,omitempty" xml:"SessionIdleTime,omitempty"`
	// The start time of the session.
	//
	// example:
	//
	// 2023-01-28T02:31:43Z
	SessionStartTime *string `json:"SessionStartTime,omitempty" xml:"SessionStartTime,omitempty"`
	// The state of the session.
	//
	// Valid values:
	//
	// 	- Connected
	//
	// 	- Disconnected
	//
	// example:
	//
	// Connected
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The billing method of cloud computers.
	//
	// Valid values:
	//
	// 	- duration: hourly plan (available for users in the whitelist)
	//
	// 	- postPaid: pay-as-you-go
	//
	// 	- monthPackage: monthly subscription (120-hour computing plan and 250-hour computing plan)
	//
	// 	- prePaid: monthly subscription (Unlimited computing plan)
	//
	// example:
	//
	// monthPackage
	SubPayType   *string                                                  `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	TerminalInfo *DescribeDesktopSessionsResponseBodySessionsTerminalInfo `json:"TerminalInfo,omitempty" xml:"TerminalInfo,omitempty" type:"Struct"`
	// The total connection duration. Unit: seconds.
	//
	// example:
	//
	// 240
	TotalConnectionTime *int64 `json:"TotalConnectionTime,omitempty" xml:"TotalConnectionTime,omitempty"`
}

func (s DescribeDesktopSessionsResponseBodySessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopSessionsResponseBodySessions) GoString() string {
	return s.String()
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetClientOS() *string {
	return s.ClientOS
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetEndUserApplyCoordinateTime() *int64 {
	return s.EndUserApplyCoordinateTime
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetLatestConnectionTime() *int64 {
	return s.LatestConnectionTime
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetOsSessionStatus() *string {
	return s.OsSessionStatus
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetOsType() *string {
	return s.OsType
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetResourceGroups() []*DescribeDesktopSessionsResponseBodySessionsResourceGroups {
	return s.ResourceGroups
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetSessionEndTime() *string {
	return s.SessionEndTime
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetSessionIdleTime() *int64 {
	return s.SessionIdleTime
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetSessionStartTime() *string {
	return s.SessionStartTime
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetSubPayType() *string {
	return s.SubPayType
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetTerminalInfo() *DescribeDesktopSessionsResponseBodySessionsTerminalInfo {
	return s.TerminalInfo
}

func (s *DescribeDesktopSessionsResponseBodySessions) GetTotalConnectionTime() *int64 {
	return s.TotalConnectionTime
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetClientIp(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.ClientIp = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetClientOS(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.ClientOS = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetClientVersion(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.ClientVersion = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetDesktopId(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetDesktopName(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetEndUserApplyCoordinateTime(v int64) *DescribeDesktopSessionsResponseBodySessions {
	s.EndUserApplyCoordinateTime = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetEndUserId(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetLatestConnectionTime(v int64) *DescribeDesktopSessionsResponseBodySessions {
	s.LatestConnectionTime = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetOfficeSiteId(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetOfficeSiteName(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetOsSessionStatus(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.OsSessionStatus = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetOsType(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetProtocolType(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetResourceGroups(v []*DescribeDesktopSessionsResponseBodySessionsResourceGroups) *DescribeDesktopSessionsResponseBodySessions {
	s.ResourceGroups = v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetSessionEndTime(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.SessionEndTime = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetSessionIdleTime(v int64) *DescribeDesktopSessionsResponseBodySessions {
	s.SessionIdleTime = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetSessionStartTime(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.SessionStartTime = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetSessionStatus(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.SessionStatus = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetSubPayType(v string) *DescribeDesktopSessionsResponseBodySessions {
	s.SubPayType = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetTerminalInfo(v *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) *DescribeDesktopSessionsResponseBodySessions {
	s.TerminalInfo = v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) SetTotalConnectionTime(v int64) *DescribeDesktopSessionsResponseBodySessions {
	s.TotalConnectionTime = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessions) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopSessionsResponseBodySessionsResourceGroups struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDesktopSessionsResponseBodySessionsResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopSessionsResponseBodySessionsResourceGroups) GoString() string {
	return s.String()
}

func (s *DescribeDesktopSessionsResponseBodySessionsResourceGroups) GetId() *string {
	return s.Id
}

func (s *DescribeDesktopSessionsResponseBodySessionsResourceGroups) GetName() *string {
	return s.Name
}

func (s *DescribeDesktopSessionsResponseBodySessionsResourceGroups) SetId(v string) *DescribeDesktopSessionsResponseBodySessionsResourceGroups {
	s.Id = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessionsResourceGroups) SetName(v string) *DescribeDesktopSessionsResponseBodySessionsResourceGroups {
	s.Name = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessionsResourceGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopSessionsResponseBodySessionsTerminalInfo struct {
	// example:
	//
	// Mac
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Mac
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 96c530bc-6095-4014-8bbc-d461b8ac****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// example:
	//
	// EBFDC7773BEBAD418A9F89429652****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeDesktopSessionsResponseBodySessionsTerminalInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopSessionsResponseBodySessionsTerminalInfo) GoString() string {
	return s.String()
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) GetModel() *string {
	return s.Model
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) SetModel(v string) *DescribeDesktopSessionsResponseBodySessionsTerminalInfo {
	s.Model = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) SetProductName(v string) *DescribeDesktopSessionsResponseBodySessionsTerminalInfo {
	s.ProductName = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) SetSerialNumber(v string) *DescribeDesktopSessionsResponseBodySessionsTerminalInfo {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) SetUuid(v string) *DescribeDesktopSessionsResponseBodySessionsTerminalInfo {
	s.Uuid = &v
	return s
}

func (s *DescribeDesktopSessionsResponseBodySessionsTerminalInfo) Validate() error {
	return dara.Validate(s)
}
