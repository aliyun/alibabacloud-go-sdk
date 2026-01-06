// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *GetDingtalkMeetingInfoResponseBody
	GetClusterName() *string
	SetConferenceId(v string) *GetDingtalkMeetingInfoResponseBody
	GetConferenceId() *string
	SetCreatorId(v string) *GetDingtalkMeetingInfoResponseBody
	GetCreatorId() *string
	SetCreatorNick(v string) *GetDingtalkMeetingInfoResponseBody
	GetCreatorNick() *string
	SetCreatorWorkNo(v string) *GetDingtalkMeetingInfoResponseBody
	GetCreatorWorkNo() *string
	SetDeptName(v string) *GetDingtalkMeetingInfoResponseBody
	GetDeptName() *string
	SetEnableQualityMonitor(v bool) *GetDingtalkMeetingInfoResponseBody
	GetEnableQualityMonitor() *bool
	SetEndTime(v int64) *GetDingtalkMeetingInfoResponseBody
	GetEndTime() *int64
	SetFreeType(v string) *GetDingtalkMeetingInfoResponseBody
	GetFreeType() *string
	SetRequestId(v string) *GetDingtalkMeetingInfoResponseBody
	GetRequestId() *string
	SetScene(v string) *GetDingtalkMeetingInfoResponseBody
	GetScene() *string
	SetStartTime(v int64) *GetDingtalkMeetingInfoResponseBody
	GetStartTime() *int64
	SetTimeLength(v int64) *GetDingtalkMeetingInfoResponseBody
	GetTimeLength() *int64
	SetTitle(v string) *GetDingtalkMeetingInfoResponseBody
	GetTitle() *string
	SetUserCount(v int32) *GetDingtalkMeetingInfoResponseBody
	GetUserCount() *int32
	SetVendorRequestId(v string) *GetDingtalkMeetingInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDingtalkMeetingInfoResponseBody
	GetVendorType() *string
}

type GetDingtalkMeetingInfoResponseBody struct {
	// example:
	//
	// cluster-1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// uid123456
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 张三
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// example:
	//
	// 123456
	CreatorWorkNo *string `json:"creatorWorkNo,omitempty" xml:"creatorWorkNo,omitempty"`
	// example:
	//
	// 技术部
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// true
	EnableQualityMonitor *bool `json:"enableQualityMonitor,omitempty" xml:"enableQualityMonitor,omitempty"`
	// example:
	//
	// 1638363600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// free
	FreeType *string `json:"freeType,omitempty" xml:"freeType,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// meeting
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 1638360000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 3600
	TimeLength *int64 `json:"timeLength,omitempty" xml:"timeLength,omitempty"`
	// example:
	//
	// 技术分享会
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 10
	UserCount *int32 `json:"userCount,omitempty" xml:"userCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetDingtalkMeetingInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetDingtalkMeetingInfoResponseBody) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingInfoResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetDingtalkMeetingInfoResponseBody) GetCreatorNick() *string {
	return s.CreatorNick
}

func (s *GetDingtalkMeetingInfoResponseBody) GetCreatorWorkNo() *string {
	return s.CreatorWorkNo
}

func (s *GetDingtalkMeetingInfoResponseBody) GetDeptName() *string {
	return s.DeptName
}

func (s *GetDingtalkMeetingInfoResponseBody) GetEnableQualityMonitor() *bool {
	return s.EnableQualityMonitor
}

func (s *GetDingtalkMeetingInfoResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingInfoResponseBody) GetFreeType() *string {
	return s.FreeType
}

func (s *GetDingtalkMeetingInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingtalkMeetingInfoResponseBody) GetScene() *string {
	return s.Scene
}

func (s *GetDingtalkMeetingInfoResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDingtalkMeetingInfoResponseBody) GetTimeLength() *int64 {
	return s.TimeLength
}

func (s *GetDingtalkMeetingInfoResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetDingtalkMeetingInfoResponseBody) GetUserCount() *int32 {
	return s.UserCount
}

func (s *GetDingtalkMeetingInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDingtalkMeetingInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDingtalkMeetingInfoResponseBody) SetClusterName(v string) *GetDingtalkMeetingInfoResponseBody {
	s.ClusterName = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetConferenceId(v string) *GetDingtalkMeetingInfoResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetCreatorId(v string) *GetDingtalkMeetingInfoResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetCreatorNick(v string) *GetDingtalkMeetingInfoResponseBody {
	s.CreatorNick = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetCreatorWorkNo(v string) *GetDingtalkMeetingInfoResponseBody {
	s.CreatorWorkNo = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetDeptName(v string) *GetDingtalkMeetingInfoResponseBody {
	s.DeptName = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetEnableQualityMonitor(v bool) *GetDingtalkMeetingInfoResponseBody {
	s.EnableQualityMonitor = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetEndTime(v int64) *GetDingtalkMeetingInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetFreeType(v string) *GetDingtalkMeetingInfoResponseBody {
	s.FreeType = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetRequestId(v string) *GetDingtalkMeetingInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetScene(v string) *GetDingtalkMeetingInfoResponseBody {
	s.Scene = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetStartTime(v int64) *GetDingtalkMeetingInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetTimeLength(v int64) *GetDingtalkMeetingInfoResponseBody {
	s.TimeLength = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetTitle(v string) *GetDingtalkMeetingInfoResponseBody {
	s.Title = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetUserCount(v int32) *GetDingtalkMeetingInfoResponseBody {
	s.UserCount = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetVendorRequestId(v string) *GetDingtalkMeetingInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) SetVendorType(v string) *GetDingtalkMeetingInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
