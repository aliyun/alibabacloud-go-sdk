// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopGroupSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeDesktopGroupSessionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDesktopGroupSessionsResponseBody
	GetRequestId() *string
	SetSessions(v []*DescribeDesktopGroupSessionsResponseBodySessions) *DescribeDesktopGroupSessionsResponseBody
	GetSessions() []*DescribeDesktopGroupSessionsResponseBodySessions
	SetTotalCount(v int32) *DescribeDesktopGroupSessionsResponseBody
	GetTotalCount() *int32
}

type DescribeDesktopGroupSessionsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0920845-7359-59FC-9899-B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sessions.
	Sessions []*DescribeDesktopGroupSessionsResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The total number of sessions.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDesktopGroupSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupSessionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopGroupSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopGroupSessionsResponseBody) GetSessions() []*DescribeDesktopGroupSessionsResponseBodySessions {
	return s.Sessions
}

func (s *DescribeDesktopGroupSessionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDesktopGroupSessionsResponseBody) SetNextToken(v string) *DescribeDesktopGroupSessionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBody) SetRequestId(v string) *DescribeDesktopGroupSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBody) SetSessions(v []*DescribeDesktopGroupSessionsResponseBodySessions) *DescribeDesktopGroupSessionsResponseBody {
	s.Sessions = v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBody) SetTotalCount(v int32) *DescribeDesktopGroupSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBody) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopGroupSessionsResponseBodySessions struct {
	// The IP address of the client.
	//
	// example:
	//
	// 172.21.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The OS that the client runs.
	//
	// example:
	//
	// windows_\\"Windows10Enterprise\\"10.0(Build22000)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The version of the client.
	//
	// example:
	//
	// 2.0.0-R-20221030.08****
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The ID of the desktop group.
	//
	// example:
	//
	// dg-iaqu3bi2xtie****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The name of the desktop group.
	//
	// example:
	//
	// Test
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// If the session is being established, the value of this parameter indicates the ID of the current cloud desktop. If the session is disconnected, the value of this parameter indicates the ID of the cloud desktop that was most recently connected.
	//
	// example:
	//
	// ecd-g6t1ukbaea****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The point in time when the end user applies for administrator assistance.
	//
	// example:
	//
	// 1678794261000
	EndUserApplyCoordinateTime *int64 `json:"EndUserApplyCoordinateTime,omitempty" xml:"EndUserApplyCoordinateTime,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// xianqiu
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The end time of the most recent connection.
	//
	// example:
	//
	// 2022-08-31 06:56:45
	LastSessionEndTime *string `json:"LastSessionEndTime,omitempty" xml:"LastSessionEndTime,omitempty"`
	// The start time of the most recent connection.
	//
	// example:
	//
	// 2022-08-31 06:56:45
	LastSessionStartTime *string `json:"LastSessionStartTime,omitempty" xml:"LastSessionStartTime,omitempty"`
	// The duration of the most recent session.
	//
	// example:
	//
	// 120
	LatestConnectionTime *int64 `json:"LatestConnectionTime,omitempty" xml:"LatestConnectionTime,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou+dir-8904****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// Test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The OS. Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The type of the session.
	//
	// Valid values:
	//
	// 	- 0: single-session
	//
	// 	- 1: multi-session
	//
	// example:
	//
	// 0
	OwnType *int32 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The type of the protocol.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The duration during which the cloud desktop stays in the Idle state.
	//
	// example:
	//
	// 120
	SessionIdleTime *int64 `json:"SessionIdleTime,omitempty" xml:"SessionIdleTime,omitempty"`
	// The state of the session.
	//
	// Valid values:
	//
	// 	- Connected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Disconnected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Connected
	SessionStatus *string                                                       `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	TerminalInfo  *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo `json:"TerminalInfo,omitempty" xml:"TerminalInfo,omitempty" type:"Struct"`
	// The total duration of the sessions.
	//
	// example:
	//
	// 120
	TotalConnectionDuration *int64 `json:"TotalConnectionDuration,omitempty" xml:"TotalConnectionDuration,omitempty"`
}

func (s DescribeDesktopGroupSessionsResponseBodySessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupSessionsResponseBodySessions) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetClientOS() *string {
	return s.ClientOS
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetEndUserApplyCoordinateTime() *int64 {
	return s.EndUserApplyCoordinateTime
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetLastSessionEndTime() *string {
	return s.LastSessionEndTime
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetLastSessionStartTime() *string {
	return s.LastSessionStartTime
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetLatestConnectionTime() *int64 {
	return s.LatestConnectionTime
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetOsType() *string {
	return s.OsType
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetOwnType() *int32 {
	return s.OwnType
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetSessionIdleTime() *int64 {
	return s.SessionIdleTime
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetTerminalInfo() *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo {
	return s.TerminalInfo
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) GetTotalConnectionDuration() *int64 {
	return s.TotalConnectionDuration
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetClientIp(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.ClientIp = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetClientOS(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.ClientOS = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetClientVersion(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.ClientVersion = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetDesktopGroupId(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetDesktopGroupName(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetDesktopId(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetEndUserApplyCoordinateTime(v int64) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.EndUserApplyCoordinateTime = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetEndUserId(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetLastSessionEndTime(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.LastSessionEndTime = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetLastSessionStartTime(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.LastSessionStartTime = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetLatestConnectionTime(v int64) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.LatestConnectionTime = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetOfficeSiteId(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetOfficeSiteName(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetOsType(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetOwnType(v int32) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.OwnType = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetProtocolType(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetSessionIdleTime(v int64) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.SessionIdleTime = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetSessionStatus(v string) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.SessionStatus = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetTerminalInfo(v *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.TerminalInfo = v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) SetTotalConnectionDuration(v int64) *DescribeDesktopGroupSessionsResponseBodySessions {
	s.TotalConnectionDuration = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessions) Validate() error {
	if s.TerminalInfo != nil {
		if err := s.TerminalInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo struct {
	// example:
	//
	// wuying_mac_x86_64
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

func (s DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) GetModel() *string {
	return s.Model
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) SetModel(v string) *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo {
	s.Model = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) SetProductName(v string) *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo {
	s.ProductName = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) SetSerialNumber(v string) *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) SetUuid(v string) *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo {
	s.Uuid = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponseBodySessionsTerminalInfo) Validate() error {
	return dara.Validate(s)
}
