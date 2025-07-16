// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMemberModels(v []*QueryConferenceMembersResponseBodyMemberModels) *QueryConferenceMembersResponseBody
	GetMemberModels() []*QueryConferenceMembersResponseBodyMemberModels
	SetNextToken(v string) *QueryConferenceMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryConferenceMembersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryConferenceMembersResponseBody
	GetTotalCount() *int32
}

type QueryConferenceMembersResponseBody struct {
	MemberModels []*QueryConferenceMembersResponseBodyMemberModels `json:"memberModels,omitempty" xml:"memberModels,omitempty" type:"Repeated"`
	// example:
	//
	// 123000000
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryConferenceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponseBody) GetMemberModels() []*QueryConferenceMembersResponseBodyMemberModels {
	return s.MemberModels
}

func (s *QueryConferenceMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryConferenceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConferenceMembersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryConferenceMembersResponseBody) SetMemberModels(v []*QueryConferenceMembersResponseBodyMemberModels) *QueryConferenceMembersResponseBody {
	s.MemberModels = v
	return s
}

func (s *QueryConferenceMembersResponseBody) SetNextToken(v string) *QueryConferenceMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceMembersResponseBody) SetRequestId(v string) *QueryConferenceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConferenceMembersResponseBody) SetTotalCount(v int32) *QueryConferenceMembersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryConferenceMembersResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryConferenceMembersResponseBodyMemberModels struct {
	// example:
	//
	// 6
	AttendStatus *int32 `json:"AttendStatus,omitempty" xml:"AttendStatus,omitempty"`
	// example:
	//
	// false
	CoHost *bool `json:"CoHost,omitempty" xml:"CoHost,omitempty"`
	// example:
	//
	// 6323dxxxxx
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// example:
	//
	// 10000
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// false
	Host *bool `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// 1663293270000
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// example:
	//
	// 1663293280000
	LeaveTime *int64 `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	// example:
	//
	// true
	OuterOrgMember *bool `json:"OuterOrgMember,omitempty" xml:"OuterOrgMember,omitempty"`
	// example:
	//
	// false
	PstnJoin *bool `json:"PstnJoin,omitempty" xml:"PstnJoin,omitempty"`
	// example:
	//
	// -12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 小钉
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s QueryConferenceMembersResponseBodyMemberModels) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersResponseBodyMemberModels) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetAttendStatus() *int32 {
	return s.AttendStatus
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetCoHost() *bool {
	return s.CoHost
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetHost() *bool {
	return s.Host
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetLeaveTime() *int64 {
	return s.LeaveTime
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetOuterOrgMember() *bool {
	return s.OuterOrgMember
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetPstnJoin() *bool {
	return s.PstnJoin
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetUserId() *string {
	return s.UserId
}

func (s *QueryConferenceMembersResponseBodyMemberModels) GetUserNick() *string {
	return s.UserNick
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetAttendStatus(v int32) *QueryConferenceMembersResponseBodyMemberModels {
	s.AttendStatus = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetCoHost(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.CoHost = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetConferenceId(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetDuration(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.Duration = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetHost(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.Host = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetJoinTime(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.JoinTime = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetLeaveTime(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.LeaveTime = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetOuterOrgMember(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.OuterOrgMember = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetPstnJoin(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.PstnJoin = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetUserId(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.UserId = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetUserNick(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.UserNick = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) Validate() error {
	return dara.Validate(s)
}
