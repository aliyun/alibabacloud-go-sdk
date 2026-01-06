// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetDingtalkMeetingMemberListResponseBody
	GetCurrentPage() *int32
	SetData(v []*GetDingtalkMeetingMemberListResponseBodyData) *GetDingtalkMeetingMemberListResponseBody
	GetData() []*GetDingtalkMeetingMemberListResponseBodyData
	SetRequestId(v string) *GetDingtalkMeetingMemberListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetDingtalkMeetingMemberListResponseBody
	GetTotalCount() *int32
	SetVendorRequestId(v string) *GetDingtalkMeetingMemberListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDingtalkMeetingMemberListResponseBody
	GetVendorType() *string
}

type GetDingtalkMeetingMemberListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// []
	Data []*GetDingtalkMeetingMemberListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetDingtalkMeetingMemberListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkMeetingMemberListResponseBody) GetData() []*GetDingtalkMeetingMemberListResponseBodyData {
	return s.Data
}

func (s *GetDingtalkMeetingMemberListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingtalkMeetingMemberListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetDingtalkMeetingMemberListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDingtalkMeetingMemberListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDingtalkMeetingMemberListResponseBody) SetCurrentPage(v int32) *GetDingtalkMeetingMemberListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBody) SetData(v []*GetDingtalkMeetingMemberListResponseBodyData) *GetDingtalkMeetingMemberListResponseBody {
	s.Data = v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBody) SetRequestId(v string) *GetDingtalkMeetingMemberListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBody) SetTotalCount(v int32) *GetDingtalkMeetingMemberListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBody) SetVendorRequestId(v string) *GetDingtalkMeetingMemberListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBody) SetVendorType(v string) *GetDingtalkMeetingMemberListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDingtalkMeetingMemberListResponseBodyData struct {
	// example:
	//
	// channel-1
	ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// PC
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// example:
	//
	// 3600
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1638360000000
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// example:
	//
	// 1638363600000
	LeaveTime *int64 `json:"leaveTime,omitempty" xml:"leaveTime,omitempty"`
	// example:
	//
	// good
	NetworkQuality *string `json:"networkQuality,omitempty" xml:"networkQuality,omitempty"`
	// example:
	//
	// 张三
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// example:
	//
	// host
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// session123
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// uid123456
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 432432
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetDingtalkMeetingMemberListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetDeviceType() *string {
	return s.DeviceType
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetLeaveTime() *int64 {
	return s.LeaveTime
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetNetworkQuality() *string {
	return s.NetworkQuality
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetNick() *string {
	return s.Nick
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetUnionId() *string {
	return s.UnionId
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) GetWorkNo() *string {
	return s.WorkNo
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetChannelName(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetConferenceId(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetDeviceType(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetDuration(v int64) *GetDingtalkMeetingMemberListResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetJoinTime(v int64) *GetDingtalkMeetingMemberListResponseBodyData {
	s.JoinTime = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetLeaveTime(v int64) *GetDingtalkMeetingMemberListResponseBodyData {
	s.LeaveTime = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetNetworkQuality(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.NetworkQuality = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetNick(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.Nick = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetRole(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.Role = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetSessionId(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetStatus(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetUnionId(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.UnionId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetVersion(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) SetWorkNo(v string) *GetDingtalkMeetingMemberListResponseBodyData {
	s.WorkNo = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
