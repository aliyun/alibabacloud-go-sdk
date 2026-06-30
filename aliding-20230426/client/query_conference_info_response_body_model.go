// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfInfo(v *QueryConferenceInfoResponseBodyConfInfo) *QueryConferenceInfoResponseBody
	GetConfInfo() *QueryConferenceInfoResponseBodyConfInfo
	SetRequestId(v string) *QueryConferenceInfoResponseBody
	GetRequestId() *string
}

type QueryConferenceInfoResponseBody struct {
	ConfInfo  *QueryConferenceInfoResponseBodyConfInfo `json:"confInfo,omitempty" xml:"confInfo,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryConferenceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponseBody) GetConfInfo() *QueryConferenceInfoResponseBodyConfInfo {
	return s.ConfInfo
}

func (s *QueryConferenceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConferenceInfoResponseBody) SetConfInfo(v *QueryConferenceInfoResponseBodyConfInfo) *QueryConferenceInfoResponseBody {
	s.ConfInfo = v
	return s
}

func (s *QueryConferenceInfoResponseBody) SetRequestId(v string) *QueryConferenceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConferenceInfoResponseBody) Validate() error {
	if s.ConfInfo != nil {
		if err := s.ConfInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConferenceInfoResponseBodyConfInfo struct {
	ActiveNum       *int32  `json:"ActiveNum,omitempty" xml:"ActiveNum,omitempty"`
	AttendNum       *int32  `json:"AttendNum,omitempty" xml:"AttendNum,omitempty"`
	ConfDuration    *int64  `json:"ConfDuration,omitempty" xml:"ConfDuration,omitempty"`
	ConferenceId    *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	CreatorId       *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorNick     *string `json:"CreatorNick,omitempty" xml:"CreatorNick,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExternalLinkUrl *string `json:"ExternalLinkUrl,omitempty" xml:"ExternalLinkUrl,omitempty"`
	InvitedNum      *int32  `json:"InvitedNum,omitempty" xml:"InvitedNum,omitempty"`
	RoomCode        *string `json:"RoomCode,omitempty" xml:"RoomCode,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryConferenceInfoResponseBodyConfInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoResponseBodyConfInfo) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetActiveNum() *int32 {
	return s.ActiveNum
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetAttendNum() *int32 {
	return s.AttendNum
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetConfDuration() *int64 {
	return s.ConfDuration
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetCreatorNick() *string {
	return s.CreatorNick
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetExternalLinkUrl() *string {
	return s.ExternalLinkUrl
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetInvitedNum() *int32 {
	return s.InvitedNum
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetRoomCode() *string {
	return s.RoomCode
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetStatus() *int32 {
	return s.Status
}

func (s *QueryConferenceInfoResponseBodyConfInfo) GetTitle() *string {
	return s.Title
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetActiveNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.ActiveNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetAttendNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.AttendNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetConfDuration(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.ConfDuration = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetConferenceId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCreatorId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.CreatorId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCreatorNick(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.CreatorNick = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetEndTime(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.EndTime = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetExternalLinkUrl(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.ExternalLinkUrl = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetInvitedNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.InvitedNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetRoomCode(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetStartTime(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.StartTime = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetStatus(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.Status = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetTitle(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.Title = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) Validate() error {
	return dara.Validate(s)
}
