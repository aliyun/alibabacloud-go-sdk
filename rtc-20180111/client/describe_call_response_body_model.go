// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallInfo(v *DescribeCallResponseBodyCallInfo) *DescribeCallResponseBody
	GetCallInfo() *DescribeCallResponseBodyCallInfo
	SetRequestId(v string) *DescribeCallResponseBody
	GetRequestId() *string
	SetUserDetailList(v []*DescribeCallResponseBodyUserDetailList) *DescribeCallResponseBody
	GetUserDetailList() []*DescribeCallResponseBodyUserDetailList
}

type DescribeCallResponseBody struct {
	CallInfo *DescribeCallResponseBodyCallInfo `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserDetailList []*DescribeCallResponseBodyUserDetailList `json:"UserDetailList,omitempty" xml:"UserDetailList,omitempty" type:"Repeated"`
}

func (s DescribeCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBody) GetCallInfo() *DescribeCallResponseBodyCallInfo {
	return s.CallInfo
}

func (s *DescribeCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCallResponseBody) GetUserDetailList() []*DescribeCallResponseBodyUserDetailList {
	return s.UserDetailList
}

func (s *DescribeCallResponseBody) SetCallInfo(v *DescribeCallResponseBodyCallInfo) *DescribeCallResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeCallResponseBody) SetRequestId(v string) *DescribeCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCallResponseBody) SetUserDetailList(v []*DescribeCallResponseBodyUserDetailList) *DescribeCallResponseBody {
	s.UserDetailList = v
	return s
}

func (s *DescribeCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCallResponseBodyCallInfo struct {
	// App ID。
	//
	// example:
	//
	// xxxxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// IN
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 123456
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1615860711
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615860811
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeCallResponseBodyCallInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyCallInfo) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCallResponseBodyCallInfo) GetCallStatus() *string {
	return s.CallStatus
}

func (s *DescribeCallResponseBodyCallInfo) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCallResponseBodyCallInfo) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeCallResponseBodyCallInfo) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeCallResponseBodyCallInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeCallResponseBodyCallInfo) SetAppId(v string) *DescribeCallResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetCallStatus(v string) *DescribeCallResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetChannelId(v string) *DescribeCallResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeCallResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeCallResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetDuration(v int64) *DescribeCallResponseBodyCallInfo {
	s.Duration = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCallResponseBodyUserDetailList struct {
	// example:
	//
	// GOOD
	CallExp *string `json:"CallExp,omitempty" xml:"CallExp,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs       *int64                                                   `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	DurMetricStatData *DescribeCallResponseBodyUserDetailListDurMetricStatData `json:"DurMetricStatData,omitempty" xml:"DurMetricStatData,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network     *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList []*string `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                 `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeCallResponseBodyUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os     *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList []*string `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0.0
	SdkVersion     *string   `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyUserDetailList) GetCallExp() *string {
	return s.CallExp
}

func (s *DescribeCallResponseBodyUserDetailList) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeCallResponseBodyUserDetailList) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeCallResponseBodyUserDetailList) GetDurMetricStatData() *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	return s.DurMetricStatData
}

func (s *DescribeCallResponseBodyUserDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeCallResponseBodyUserDetailList) GetLocation() *string {
	return s.Location
}

func (s *DescribeCallResponseBodyUserDetailList) GetNetwork() *string {
	return s.Network
}

func (s *DescribeCallResponseBodyUserDetailList) GetNetworkList() []*string {
	return s.NetworkList
}

func (s *DescribeCallResponseBodyUserDetailList) GetOnlineDuration() *int64 {
	return s.OnlineDuration
}

func (s *DescribeCallResponseBodyUserDetailList) GetOnlinePeriods() []*DescribeCallResponseBodyUserDetailListOnlinePeriods {
	return s.OnlinePeriods
}

func (s *DescribeCallResponseBodyUserDetailList) GetOs() *string {
	return s.Os
}

func (s *DescribeCallResponseBodyUserDetailList) GetOsList() []*string {
	return s.OsList
}

func (s *DescribeCallResponseBodyUserDetailList) GetRoles() []*string {
	return s.Roles
}

func (s *DescribeCallResponseBodyUserDetailList) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *DescribeCallResponseBodyUserDetailList) GetSdkVersionList() []*string {
	return s.SdkVersionList
}

func (s *DescribeCallResponseBodyUserDetailList) GetUserId() *string {
	return s.UserId
}

func (s *DescribeCallResponseBodyUserDetailList) SetCallExp(v string) *DescribeCallResponseBodyUserDetailList {
	s.CallExp = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetCreatedTs(v int64) *DescribeCallResponseBodyUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetDestroyedTs(v int64) *DescribeCallResponseBodyUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetDurMetricStatData(v *DescribeCallResponseBodyUserDetailListDurMetricStatData) *DescribeCallResponseBodyUserDetailList {
	s.DurMetricStatData = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetDuration(v int64) *DescribeCallResponseBodyUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetLocation(v string) *DescribeCallResponseBodyUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetNetwork(v string) *DescribeCallResponseBodyUserDetailList {
	s.Network = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetNetworkList(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.NetworkList = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOnlineDuration(v int64) *DescribeCallResponseBodyUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOnlinePeriods(v []*DescribeCallResponseBodyUserDetailListOnlinePeriods) *DescribeCallResponseBodyUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOs(v string) *DescribeCallResponseBodyUserDetailList {
	s.Os = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOsList(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.OsList = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetRoles(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.Roles = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetSdkVersion(v string) *DescribeCallResponseBodyUserDetailList {
	s.SdkVersion = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetSdkVersionList(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.SdkVersionList = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetUserId(v string) *DescribeCallResponseBodyUserDetailList {
	s.UserId = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) Validate() error {
	return dara.Validate(s)
}

type DescribeCallResponseBodyUserDetailListDurMetricStatData struct {
	// example:
	//
	// 0
	PubAudio *int64 `json:"PubAudio,omitempty" xml:"PubAudio,omitempty"`
	// example:
	//
	// 0
	PubVideo1080 *int64 `json:"PubVideo1080,omitempty" xml:"PubVideo1080,omitempty"`
	// example:
	//
	// 0
	PubVideo360 *int64 `json:"PubVideo360,omitempty" xml:"PubVideo360,omitempty"`
	// example:
	//
	// 0
	PubVideo720 *int64 `json:"PubVideo720,omitempty" xml:"PubVideo720,omitempty"`
	// example:
	//
	// 0
	PubVideoScreenShare *int64 `json:"PubVideoScreenShare,omitempty" xml:"PubVideoScreenShare,omitempty"`
	// example:
	//
	// 0
	SubAudio *int64 `json:"SubAudio,omitempty" xml:"SubAudio,omitempty"`
	// example:
	//
	// 0
	SubVideo1080 *int64 `json:"SubVideo1080,omitempty" xml:"SubVideo1080,omitempty"`
	// example:
	//
	// 0
	SubVideo360 *int64 `json:"SubVideo360,omitempty" xml:"SubVideo360,omitempty"`
	// example:
	//
	// 0
	SubVideo720 *int64 `json:"SubVideo720,omitempty" xml:"SubVideo720,omitempty"`
	// example:
	//
	// 0
	SubVideoScreenShare *int64 `json:"SubVideoScreenShare,omitempty" xml:"SubVideoScreenShare,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailListDurMetricStatData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailListDurMetricStatData) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetPubAudio() *int64 {
	return s.PubAudio
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetPubVideo1080() *int64 {
	return s.PubVideo1080
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetPubVideo360() *int64 {
	return s.PubVideo360
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetPubVideo720() *int64 {
	return s.PubVideo720
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetPubVideoScreenShare() *int64 {
	return s.PubVideoScreenShare
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetSubAudio() *int64 {
	return s.SubAudio
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetSubVideo1080() *int64 {
	return s.SubVideo1080
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetSubVideo360() *int64 {
	return s.SubVideo360
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetSubVideo720() *int64 {
	return s.SubVideo720
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) GetSubVideoScreenShare() *int64 {
	return s.SubVideoScreenShare
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubAudio(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubAudio = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideo1080(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo1080 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideo360(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo360 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideo720(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo720 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideoScreenShare(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideoScreenShare = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubAudio(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubAudio = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideo1080(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo1080 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideo360(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo360 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideo720(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo720 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideoScreenShare(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideoScreenShare = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) Validate() error {
	return dara.Validate(s)
}

type DescribeCallResponseBodyUserDetailListOnlinePeriods struct {
	// example:
	//
	// 1614936817
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1614936817
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailListOnlinePeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) GetJoinTs() *int64 {
	return s.JoinTs
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) GetLeaveTs() *int64 {
	return s.LeaveTs
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribeCallResponseBodyUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribeCallResponseBodyUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) Validate() error {
	return dara.Validate(s)
}
