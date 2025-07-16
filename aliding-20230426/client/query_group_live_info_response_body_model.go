// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupLiveInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUrl(v string) *QueryGroupLiveInfoResponseBody
	GetCoverUrl() *string
	SetDuration(v int64) *QueryGroupLiveInfoResponseBody
	GetDuration() *int64
	SetEndTime(v int64) *QueryGroupLiveInfoResponseBody
	GetEndTime() *int64
	SetIntroduction(v string) *QueryGroupLiveInfoResponseBody
	GetIntroduction() *string
	SetLiveId(v string) *QueryGroupLiveInfoResponseBody
	GetLiveId() *string
	SetLivePlayUrl(v string) *QueryGroupLiveInfoResponseBody
	GetLivePlayUrl() *string
	SetLiveStatus(v int32) *QueryGroupLiveInfoResponseBody
	GetLiveStatus() *int32
	SetPlaybackDuration(v int64) *QueryGroupLiveInfoResponseBody
	GetPlaybackDuration() *int64
	SetReplayUrl(v string) *QueryGroupLiveInfoResponseBody
	GetReplayUrl() *string
	SetRequestId(v string) *QueryGroupLiveInfoResponseBody
	GetRequestId() *string
	SetStaffId(v string) *QueryGroupLiveInfoResponseBody
	GetStaffId() *string
	SetStartTime(v int64) *QueryGroupLiveInfoResponseBody
	GetStartTime() *int64
	SetSubscribeCount(v int32) *QueryGroupLiveInfoResponseBody
	GetSubscribeCount() *int32
	SetTitle(v string) *QueryGroupLiveInfoResponseBody
	GetTitle() *string
	SetUv(v int32) *QueryGroupLiveInfoResponseBody
	GetUv() *int32
	SetVendorRequestId(v string) *QueryGroupLiveInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryGroupLiveInfoResponseBody
	GetVendorType() *string
}

type QueryGroupLiveInfoResponseBody struct {
	// example:
	//
	// http://xxx/kk.jpg
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 59886
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1687928400000
	EndTime      *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// example:
	//
	// http://ssssss
	LivePlayUrl *string `json:"livePlayUrl,omitempty" xml:"livePlayUrl,omitempty"`
	LiveStatus  *int32  `json:"liveStatus,omitempty" xml:"liveStatus,omitempty"`
	// example:
	//
	// 13414
	PlaybackDuration *int64  `json:"playbackDuration,omitempty" xml:"playbackDuration,omitempty"`
	ReplayUrl        *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	StaffId   *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// example:
	//
	// 1687924800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 0
	SubscribeCount *int32  `json:"subscribeCount,omitempty" xml:"subscribeCount,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 10
	Uv *int32 `json:"uv,omitempty" xml:"uv,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryGroupLiveInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoResponseBody) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *QueryGroupLiveInfoResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryGroupLiveInfoResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryGroupLiveInfoResponseBody) GetIntroduction() *string {
	return s.Introduction
}

func (s *QueryGroupLiveInfoResponseBody) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryGroupLiveInfoResponseBody) GetLivePlayUrl() *string {
	return s.LivePlayUrl
}

func (s *QueryGroupLiveInfoResponseBody) GetLiveStatus() *int32 {
	return s.LiveStatus
}

func (s *QueryGroupLiveInfoResponseBody) GetPlaybackDuration() *int64 {
	return s.PlaybackDuration
}

func (s *QueryGroupLiveInfoResponseBody) GetReplayUrl() *string {
	return s.ReplayUrl
}

func (s *QueryGroupLiveInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGroupLiveInfoResponseBody) GetStaffId() *string {
	return s.StaffId
}

func (s *QueryGroupLiveInfoResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryGroupLiveInfoResponseBody) GetSubscribeCount() *int32 {
	return s.SubscribeCount
}

func (s *QueryGroupLiveInfoResponseBody) GetTitle() *string {
	return s.Title
}

func (s *QueryGroupLiveInfoResponseBody) GetUv() *int32 {
	return s.Uv
}

func (s *QueryGroupLiveInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryGroupLiveInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryGroupLiveInfoResponseBody) SetCoverUrl(v string) *QueryGroupLiveInfoResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetDuration(v int64) *QueryGroupLiveInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetEndTime(v int64) *QueryGroupLiveInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetIntroduction(v string) *QueryGroupLiveInfoResponseBody {
	s.Introduction = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetLiveId(v string) *QueryGroupLiveInfoResponseBody {
	s.LiveId = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetLivePlayUrl(v string) *QueryGroupLiveInfoResponseBody {
	s.LivePlayUrl = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetLiveStatus(v int32) *QueryGroupLiveInfoResponseBody {
	s.LiveStatus = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetPlaybackDuration(v int64) *QueryGroupLiveInfoResponseBody {
	s.PlaybackDuration = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetReplayUrl(v string) *QueryGroupLiveInfoResponseBody {
	s.ReplayUrl = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetRequestId(v string) *QueryGroupLiveInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetStaffId(v string) *QueryGroupLiveInfoResponseBody {
	s.StaffId = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetStartTime(v int64) *QueryGroupLiveInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetSubscribeCount(v int32) *QueryGroupLiveInfoResponseBody {
	s.SubscribeCount = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetTitle(v string) *QueryGroupLiveInfoResponseBody {
	s.Title = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetUv(v int32) *QueryGroupLiveInfoResponseBody {
	s.Uv = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetVendorRequestId(v string) *QueryGroupLiveInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) SetVendorType(v string) *QueryGroupLiveInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryGroupLiveInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
