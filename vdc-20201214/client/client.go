// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeAppConfigRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeAppConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppConfigRequest) SetAppId(v string) *DescribeAppConfigRequest {
	s.AppId = &v
	return s
}

type DescribeAppConfigResponseBody struct {
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ThresholdConfig *DescribeAppConfigResponseBodyThresholdConfig `json:"ThresholdConfig,omitempty" xml:"ThresholdConfig,omitempty" type:"Struct"`
}

func (s DescribeAppConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppConfigResponseBody) SetRequestId(v string) *DescribeAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppConfigResponseBody) SetThresholdConfig(v *DescribeAppConfigResponseBodyThresholdConfig) *DescribeAppConfigResponseBody {
	s.ThresholdConfig = v
	return s
}

type DescribeAppConfigResponseBodyThresholdConfig struct {
	JoinSlowTime *int64 `json:"JoinSlowTime,omitempty" xml:"JoinSlowTime,omitempty"`
}

func (s DescribeAppConfigResponseBodyThresholdConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppConfigResponseBodyThresholdConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppConfigResponseBodyThresholdConfig) SetJoinSlowTime(v int64) *DescribeAppConfigResponseBodyThresholdConfig {
	s.JoinSlowTime = &v
	return s
}

type DescribeAppConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppConfigResponse) SetHeaders(v map[string]*string) *DescribeAppConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppConfigResponse) SetStatusCode(v int32) *DescribeAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppConfigResponse) SetBody(v *DescribeAppConfigResponseBody) *DescribeAppConfigResponse {
	s.Body = v
	return s
}

type DescribeCallRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId    *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs    *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs  *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	ExtDataType  *string `json:"ExtDataType,omitempty" xml:"ExtDataType,omitempty"`
	QueryExpInfo *bool   `json:"QueryExpInfo,omitempty" xml:"QueryExpInfo,omitempty"`
}

func (s DescribeCallRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallRequest) SetAppId(v string) *DescribeCallRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallRequest) SetChannelId(v string) *DescribeCallRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallRequest) SetCreatedTs(v int64) *DescribeCallRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallRequest) SetDestroyedTs(v int64) *DescribeCallRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallRequest) SetExtDataType(v string) *DescribeCallRequest {
	s.ExtDataType = &v
	return s
}

func (s *DescribeCallRequest) SetQueryExpInfo(v bool) *DescribeCallRequest {
	s.QueryExpInfo = &v
	return s
}

type DescribeCallResponseBody struct {
	CallInfo       *DescribeCallResponseBodyCallInfo         `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserDetailList []*DescribeCallResponseBodyUserDetailList `json:"UserDetailList,omitempty" xml:"UserDetailList,omitempty" type:"Repeated"`
}

func (s DescribeCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBody) GoString() string {
	return s.String()
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

type DescribeCallResponseBodyCallInfo struct {
	// App ID。
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallStatus  *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeCallResponseBodyCallInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyCallInfo) GoString() string {
	return s.String()
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

type DescribeCallResponseBodyUserDetailList struct {
	CallExp           *string                                                  `json:"CallExp,omitempty" xml:"CallExp,omitempty"`
	CreatedTs         *int64                                                   `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs       *int64                                                   `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	DurMetricStatData *DescribeCallResponseBodyUserDetailListDurMetricStatData `json:"DurMetricStatData,omitempty" xml:"DurMetricStatData,omitempty" type:"Struct"`
	Duration          *int64                                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Location          *string                                                  `json:"Location,omitempty" xml:"Location,omitempty"`
	Network           *string                                                  `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList       []*string                                                `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	OnlineDuration    *int64                                                   `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods     []*DescribeCallResponseBodyUserDetailListOnlinePeriods   `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	Os                *string                                                  `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList            []*string                                                `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles             []*string                                                `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	SdkVersion        *string                                                  `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList    []*string                                                `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	UserId            *string                                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserIdAlias       *string                                                  `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailList) GoString() string {
	return s.String()
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

func (s *DescribeCallResponseBodyUserDetailList) SetUserIdAlias(v string) *DescribeCallResponseBodyUserDetailList {
	s.UserIdAlias = &v
	return s
}

type DescribeCallResponseBodyUserDetailListDurMetricStatData struct {
	PubAudio            *int64 `json:"PubAudio,omitempty" xml:"PubAudio,omitempty"`
	PubVideo1080        *int64 `json:"PubVideo1080,omitempty" xml:"PubVideo1080,omitempty"`
	PubVideo360         *int64 `json:"PubVideo360,omitempty" xml:"PubVideo360,omitempty"`
	PubVideo720         *int64 `json:"PubVideo720,omitempty" xml:"PubVideo720,omitempty"`
	PubVideoScreenShare *int64 `json:"PubVideoScreenShare,omitempty" xml:"PubVideoScreenShare,omitempty"`
	SubAudio            *int64 `json:"SubAudio,omitempty" xml:"SubAudio,omitempty"`
	SubVideo1080        *int64 `json:"SubVideo1080,omitempty" xml:"SubVideo1080,omitempty"`
	SubVideo360         *int64 `json:"SubVideo360,omitempty" xml:"SubVideo360,omitempty"`
	SubVideo720         *int64 `json:"SubVideo720,omitempty" xml:"SubVideo720,omitempty"`
	SubVideoScreenShare *int64 `json:"SubVideoScreenShare,omitempty" xml:"SubVideoScreenShare,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailListDurMetricStatData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailListDurMetricStatData) GoString() string {
	return s.String()
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

type DescribeCallResponseBodyUserDetailListOnlinePeriods struct {
	JoinTs  *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailListOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribeCallResponseBodyUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribeCallResponseBodyUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribeCallResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallResponse) SetHeaders(v map[string]*string) *DescribeCallResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallResponse) SetStatusCode(v int32) *DescribeCallResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallResponse) SetBody(v *DescribeCallResponseBody) *DescribeCallResponse {
	s.Body = v
	return s
}

type DescribeCallInfoRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeCallInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallInfoRequest) SetAppId(v string) *DescribeCallInfoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallInfoRequest) SetChannelId(v string) *DescribeCallInfoRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallInfoRequest) SetCreatedTs(v int64) *DescribeCallInfoRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallInfoRequest) SetDestroyedTs(v int64) *DescribeCallInfoRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeCallInfoResponseBody struct {
	CallInfo  *DescribeCallInfoResponseBodyCallInfo `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCallInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallInfoResponseBody) SetCallInfo(v *DescribeCallInfoResponseBodyCallInfo) *DescribeCallInfoResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeCallInfoResponseBody) SetRequestId(v string) *DescribeCallInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCallInfoResponseBodyCallInfo struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallStatus  *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeCallInfoResponseBodyCallInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallInfoResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeCallInfoResponseBodyCallInfo) SetAppId(v string) *DescribeCallInfoResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeCallInfoResponseBodyCallInfo) SetCallStatus(v string) *DescribeCallInfoResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallInfoResponseBodyCallInfo) SetChannelId(v string) *DescribeCallInfoResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallInfoResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeCallInfoResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallInfoResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeCallInfoResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallInfoResponseBodyCallInfo) SetDuration(v int64) *DescribeCallInfoResponseBodyCallInfo {
	s.Duration = &v
	return s
}

type DescribeCallInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallInfoResponse) SetHeaders(v map[string]*string) *DescribeCallInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallInfoResponse) SetStatusCode(v int32) *DescribeCallInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallInfoResponse) SetBody(v *DescribeCallInfoResponseBody) *DescribeCallInfoResponse {
	s.Body = v
	return s
}

type DescribeCallListRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTs      *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryMode  *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	StartTs    *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeCallListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallListRequest) SetAppId(v string) *DescribeCallListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallListRequest) SetCallStatus(v string) *DescribeCallListRequest {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallListRequest) SetChannelId(v string) *DescribeCallListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallListRequest) SetEndTs(v int64) *DescribeCallListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeCallListRequest) SetOrderBy(v string) *DescribeCallListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeCallListRequest) SetPageNo(v int32) *DescribeCallListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCallListRequest) SetPageSize(v int32) *DescribeCallListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCallListRequest) SetQueryMode(v string) *DescribeCallListRequest {
	s.QueryMode = &v
	return s
}

func (s *DescribeCallListRequest) SetStartTs(v int64) *DescribeCallListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeCallListRequest) SetUserId(v string) *DescribeCallListRequest {
	s.UserId = &v
	return s
}

type DescribeCallListResponseBody struct {
	CallList  []*DescribeCallListResponseBodyCallList `json:"CallList,omitempty" xml:"CallList,omitempty" type:"Repeated"`
	PageNo    *int32                                  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCnt  *int32                                  `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeCallListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponseBody) SetCallList(v []*DescribeCallListResponseBodyCallList) *DescribeCallListResponseBody {
	s.CallList = v
	return s
}

func (s *DescribeCallListResponseBody) SetPageNo(v int32) *DescribeCallListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeCallListResponseBody) SetPageSize(v int32) *DescribeCallListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCallListResponseBody) SetRequestId(v string) *DescribeCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCallListResponseBody) SetTotalCnt(v int32) *DescribeCallListResponseBody {
	s.TotalCnt = &v
	return s
}

type DescribeCallListResponseBodyCallList struct {
	// App ID。
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BadExpUserCnt *int32  `json:"BadExpUserCnt,omitempty" xml:"BadExpUserCnt,omitempty"`
	CallStatus    *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	ChannelId     *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs     *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs   *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	UserCnt       *int32  `json:"UserCnt,omitempty" xml:"UserCnt,omitempty"`
}

func (s DescribeCallListResponseBodyCallList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListResponseBodyCallList) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponseBodyCallList) SetAppId(v string) *DescribeCallListResponseBodyCallList {
	s.AppId = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetBadExpUserCnt(v int32) *DescribeCallListResponseBodyCallList {
	s.BadExpUserCnt = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetCallStatus(v string) *DescribeCallListResponseBodyCallList {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetChannelId(v string) *DescribeCallListResponseBodyCallList {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetCreatedTs(v int64) *DescribeCallListResponseBodyCallList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetDestroyedTs(v int64) *DescribeCallListResponseBodyCallList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetDuration(v int64) *DescribeCallListResponseBodyCallList {
	s.Duration = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetUserCnt(v int32) *DescribeCallListResponseBodyCallList {
	s.UserCnt = &v
	return s
}

type DescribeCallListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponse) SetHeaders(v map[string]*string) *DescribeCallListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallListResponse) SetStatusCode(v int32) *DescribeCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallListResponse) SetBody(v *DescribeCallListResponseBody) *DescribeCallListResponse {
	s.Body = v
	return s
}

type DescribeCallUserExpRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeCallUserExpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserExpRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallUserExpRequest) SetAppId(v string) *DescribeCallUserExpRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallUserExpRequest) SetChannelId(v string) *DescribeCallUserExpRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallUserExpRequest) SetCreatedTs(v int64) *DescribeCallUserExpRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallUserExpRequest) SetDestroyedTs(v int64) *DescribeCallUserExpRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeCallUserExpResponseBody struct {
	ExpInfoList []*DescribeCallUserExpResponseBodyExpInfoList `json:"ExpInfoList,omitempty" xml:"ExpInfoList,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCallUserExpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserExpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallUserExpResponseBody) SetExpInfoList(v []*DescribeCallUserExpResponseBodyExpInfoList) *DescribeCallUserExpResponseBody {
	s.ExpInfoList = v
	return s
}

func (s *DescribeCallUserExpResponseBody) SetRequestId(v string) *DescribeCallUserExpResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCallUserExpResponseBodyExpInfoList struct {
	CallExp *string `json:"CallExp,omitempty" xml:"CallExp,omitempty"`
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeCallUserExpResponseBodyExpInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserExpResponseBodyExpInfoList) GoString() string {
	return s.String()
}

func (s *DescribeCallUserExpResponseBodyExpInfoList) SetCallExp(v string) *DescribeCallUserExpResponseBodyExpInfoList {
	s.CallExp = &v
	return s
}

func (s *DescribeCallUserExpResponseBodyExpInfoList) SetUserId(v string) *DescribeCallUserExpResponseBodyExpInfoList {
	s.UserId = &v
	return s
}

type DescribeCallUserExpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallUserExpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallUserExpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserExpResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallUserExpResponse) SetHeaders(v map[string]*string) *DescribeCallUserExpResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallUserExpResponse) SetStatusCode(v int32) *DescribeCallUserExpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallUserExpResponse) SetBody(v *DescribeCallUserExpResponseBody) *DescribeCallUserExpResponse {
	s.Body = v
	return s
}

type DescribeCallUserListRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId    *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs    *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs  *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	ExtDataType  *string `json:"ExtDataType,omitempty" xml:"ExtDataType,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryExpInfo *bool   `json:"QueryExpInfo,omitempty" xml:"QueryExpInfo,omitempty"`
	RoleType     *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeCallUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallUserListRequest) SetAppId(v string) *DescribeCallUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallUserListRequest) SetChannelId(v string) *DescribeCallUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallUserListRequest) SetCreatedTs(v int64) *DescribeCallUserListRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallUserListRequest) SetDestroyedTs(v int64) *DescribeCallUserListRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallUserListRequest) SetExtDataType(v string) *DescribeCallUserListRequest {
	s.ExtDataType = &v
	return s
}

func (s *DescribeCallUserListRequest) SetPageNo(v int32) *DescribeCallUserListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCallUserListRequest) SetPageSize(v int32) *DescribeCallUserListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCallUserListRequest) SetQueryExpInfo(v bool) *DescribeCallUserListRequest {
	s.QueryExpInfo = &v
	return s
}

func (s *DescribeCallUserListRequest) SetRoleType(v string) *DescribeCallUserListRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCallUserListRequest) SetUserId(v string) *DescribeCallUserListRequest {
	s.UserId = &v
	return s
}

type DescribeCallUserListResponseBody struct {
	PageNo         *int32                                            `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCnt       *int32                                            `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	UserDetailList []*DescribeCallUserListResponseBodyUserDetailList `json:"UserDetailList,omitempty" xml:"UserDetailList,omitempty" type:"Repeated"`
}

func (s DescribeCallUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallUserListResponseBody) SetPageNo(v int32) *DescribeCallUserListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeCallUserListResponseBody) SetPageSize(v int32) *DescribeCallUserListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCallUserListResponseBody) SetRequestId(v string) *DescribeCallUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCallUserListResponseBody) SetTotalCnt(v int32) *DescribeCallUserListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeCallUserListResponseBody) SetUserDetailList(v []*DescribeCallUserListResponseBodyUserDetailList) *DescribeCallUserListResponseBody {
	s.UserDetailList = v
	return s
}

type DescribeCallUserListResponseBodyUserDetailList struct {
	CallExp           *string                                                          `json:"CallExp,omitempty" xml:"CallExp,omitempty"`
	CreatedTs         *int64                                                           `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs       *int64                                                           `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	DurMetricStatData *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData `json:"DurMetricStatData,omitempty" xml:"DurMetricStatData,omitempty" type:"Struct"`
	Duration          *int64                                                           `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Location          *string                                                          `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationCn        *string                                                          `json:"LocationCn,omitempty" xml:"LocationCn,omitempty"`
	LocationEn        *string                                                          `json:"LocationEn,omitempty" xml:"LocationEn,omitempty"`
	Network           *string                                                          `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList       []*string                                                        `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	OnlineDuration    *int64                                                           `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods     []*DescribeCallUserListResponseBodyUserDetailListOnlinePeriods   `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	Os                *string                                                          `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList            []*string                                                        `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles             []*string                                                        `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	SdkVersion        *string                                                          `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList    []*string                                                        `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	UserId            *string                                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserIdAlias       *string                                                          `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribeCallUserListResponseBodyUserDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserListResponseBodyUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetCallExp(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.CallExp = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetCreatedTs(v int64) *DescribeCallUserListResponseBodyUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetDestroyedTs(v int64) *DescribeCallUserListResponseBodyUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetDurMetricStatData(v *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) *DescribeCallUserListResponseBodyUserDetailList {
	s.DurMetricStatData = v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetDuration(v int64) *DescribeCallUserListResponseBodyUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetLocation(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetLocationCn(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.LocationCn = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetLocationEn(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.LocationEn = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetNetwork(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.Network = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetNetworkList(v []*string) *DescribeCallUserListResponseBodyUserDetailList {
	s.NetworkList = v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetOnlineDuration(v int64) *DescribeCallUserListResponseBodyUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetOnlinePeriods(v []*DescribeCallUserListResponseBodyUserDetailListOnlinePeriods) *DescribeCallUserListResponseBodyUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetOs(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.Os = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetOsList(v []*string) *DescribeCallUserListResponseBodyUserDetailList {
	s.OsList = v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetRoles(v []*string) *DescribeCallUserListResponseBodyUserDetailList {
	s.Roles = v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetSdkVersion(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.SdkVersion = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetSdkVersionList(v []*string) *DescribeCallUserListResponseBodyUserDetailList {
	s.SdkVersionList = v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetUserId(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.UserId = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailList) SetUserIdAlias(v string) *DescribeCallUserListResponseBodyUserDetailList {
	s.UserIdAlias = &v
	return s
}

type DescribeCallUserListResponseBodyUserDetailListDurMetricStatData struct {
	PubAudio            *int64 `json:"PubAudio,omitempty" xml:"PubAudio,omitempty"`
	PubVideo360         *int64 `json:"PubVideo360,omitempty" xml:"PubVideo360,omitempty"`
	PubVideo720         *int64 `json:"PubVideo720,omitempty" xml:"PubVideo720,omitempty"`
	PubVideoScreenShare *int64 `json:"PubVideoScreenShare,omitempty" xml:"PubVideoScreenShare,omitempty"`
	SubAudio            *int64 `json:"SubAudio,omitempty" xml:"SubAudio,omitempty"`
	SubVideo1080        *int64 `json:"SubVideo1080,omitempty" xml:"SubVideo1080,omitempty"`
	SubVideo360         *int64 `json:"SubVideo360,omitempty" xml:"SubVideo360,omitempty"`
	SubVideo720         *int64 `json:"SubVideo720,omitempty" xml:"SubVideo720,omitempty"`
	SubVideoScreenShare *int64 `json:"SubVideoScreenShare,omitempty" xml:"SubVideoScreenShare,omitempty"`
}

func (s DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) GoString() string {
	return s.String()
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetPubAudio(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.PubAudio = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetPubVideo360(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo360 = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetPubVideo720(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo720 = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetPubVideoScreenShare(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.PubVideoScreenShare = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetSubAudio(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.SubAudio = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetSubVideo1080(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo1080 = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetSubVideo360(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo360 = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetSubVideo720(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo720 = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData) SetSubVideoScreenShare(v int64) *DescribeCallUserListResponseBodyUserDetailListDurMetricStatData {
	s.SubVideoScreenShare = &v
	return s
}

type DescribeCallUserListResponseBodyUserDetailListOnlinePeriods struct {
	JoinTs  *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeCallUserListResponseBodyUserDetailListOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserListResponseBodyUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeCallUserListResponseBodyUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribeCallUserListResponseBodyUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeCallUserListResponseBodyUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribeCallUserListResponseBodyUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribeCallUserListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallUserListResponse) SetHeaders(v map[string]*string) *DescribeCallUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallUserListResponse) SetStatusCode(v int32) *DescribeCallUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallUserListResponse) SetBody(v *DescribeCallUserListResponseBody) *DescribeCallUserListResponse {
	s.Body = v
	return s
}

type DescribeChannelAreaDistributionStatDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	ParentArea  *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetAppId(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetChannelId(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetCreatedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetDestroyedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

type DescribeChannelAreaDistributionStatDataResponseBody struct {
	AreaStatList []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList `json:"AreaStatList,omitempty" xml:"AreaStatList,omitempty" type:"Repeated"`
	RequestId    *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) SetAreaStatList(v []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) *DescribeChannelAreaDistributionStatDataResponseBody {
	s.AreaStatList = v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeChannelAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList struct {
	AreaName                    *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	CallUserCount               *int32  `json:"CallUserCount,omitempty" xml:"CallUserCount,omitempty"`
	HighQualityTransmissionRate *string `json:"HighQualityTransmissionRate,omitempty" xml:"HighQualityTransmissionRate,omitempty"`
	PubUserCount                *int32  `json:"PubUserCount,omitempty" xml:"PubUserCount,omitempty"`
	SubUserCount                *int32  `json:"SubUserCount,omitempty" xml:"SubUserCount,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetAreaName(v string) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.AreaName = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetCallUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.CallUserCount = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetHighQualityTransmissionRate(v string) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.HighQualityTransmissionRate = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetPubUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.PubUserCount = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetSubUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.SubUserCount = &v
	return s
}

type DescribeChannelAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeChannelAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeChannelAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetBody(v *DescribeChannelAreaDistributionStatDataResponseBody) *DescribeChannelAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeChannelDistributionStatDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	StatDim     *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeChannelDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataRequest) SetAppId(v string) *DescribeChannelDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetChannelId(v string) *DescribeChannelDistributionStatDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetCreatedTs(v int64) *DescribeChannelDistributionStatDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetDestroyedTs(v int64) *DescribeChannelDistributionStatDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetStatDim(v string) *DescribeChannelDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

type DescribeChannelDistributionStatDataResponseBody struct {
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList  []*DescribeChannelDistributionStatDataResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
}

func (s DescribeChannelDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponseBody) SetRequestId(v string) *DescribeChannelDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBody) SetStatList(v []*DescribeChannelDistributionStatDataResponseBodyStatList) *DescribeChannelDistributionStatDataResponseBody {
	s.StatList = v
	return s
}

type DescribeChannelDistributionStatDataResponseBodyStatList struct {
	CallUserCount *int32  `json:"CallUserCount,omitempty" xml:"CallUserCount,omitempty"`
	CallUserRatio *string `json:"CallUserRatio,omitempty" xml:"CallUserRatio,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeChannelDistributionStatDataResponseBodyStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetCallUserCount(v int32) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.CallUserCount = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetCallUserRatio(v string) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.CallUserRatio = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetName(v string) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.Name = &v
	return s
}

type DescribeChannelDistributionStatDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeChannelDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelDistributionStatDataResponse) SetStatusCode(v int32) *DescribeChannelDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponse) SetBody(v *DescribeChannelDistributionStatDataResponseBody) *DescribeChannelDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeChannelJoinInfoRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelJoinInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelJoinInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelJoinInfoRequest) SetAppId(v string) *DescribeChannelJoinInfoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelJoinInfoRequest) SetChannelId(v string) *DescribeChannelJoinInfoRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelJoinInfoRequest) SetCreatedTs(v int64) *DescribeChannelJoinInfoRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelJoinInfoRequest) SetDestroyedTs(v int64) *DescribeChannelJoinInfoRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeChannelJoinInfoResponseBody struct {
	JoinFastSuccessRate *string `json:"JoinFastSuccessRate,omitempty" xml:"JoinFastSuccessRate,omitempty"`
	JoinSlowThreshold   *int64  `json:"JoinSlowThreshold,omitempty" xml:"JoinSlowThreshold,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelJoinInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelJoinInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelJoinInfoResponseBody) SetJoinFastSuccessRate(v string) *DescribeChannelJoinInfoResponseBody {
	s.JoinFastSuccessRate = &v
	return s
}

func (s *DescribeChannelJoinInfoResponseBody) SetJoinSlowThreshold(v int64) *DescribeChannelJoinInfoResponseBody {
	s.JoinSlowThreshold = &v
	return s
}

func (s *DescribeChannelJoinInfoResponseBody) SetRequestId(v string) *DescribeChannelJoinInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelJoinInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelJoinInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelJoinInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelJoinInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelJoinInfoResponse) SetHeaders(v map[string]*string) *DescribeChannelJoinInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelJoinInfoResponse) SetStatusCode(v int32) *DescribeChannelJoinInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelJoinInfoResponse) SetBody(v *DescribeChannelJoinInfoResponseBody) *DescribeChannelJoinInfoResponse {
	s.Body = v
	return s
}

type DescribeChannelOverallDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataRequest) SetAppId(v string) *DescribeChannelOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetChannelId(v string) *DescribeChannelOverallDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetCreatedTs(v int64) *DescribeChannelOverallDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetDestroyedTs(v int64) *DescribeChannelOverallDataRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeChannelOverallDataResponseBody struct {
	CallInfo    *DescribeChannelOverallDataResponseBodyCallInfo      `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	MetricDatas []*DescribeChannelOverallDataResponseBodyMetricDatas `json:"MetricDatas,omitempty" xml:"MetricDatas,omitempty" type:"Repeated"`
	OverallData *DescribeChannelOverallDataResponseBodyOverallData   `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBody) SetCallInfo(v *DescribeChannelOverallDataResponseBodyCallInfo) *DescribeChannelOverallDataResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetMetricDatas(v []*DescribeChannelOverallDataResponseBodyMetricDatas) *DescribeChannelOverallDataResponseBody {
	s.MetricDatas = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetOverallData(v *DescribeChannelOverallDataResponseBodyOverallData) *DescribeChannelOverallDataResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetRequestId(v string) *DescribeChannelOverallDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelOverallDataResponseBodyCallInfo struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallStatus  *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyCallInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetAppId(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetCallStatus(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetChannelId(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetDuration(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.Duration = &v
	return s
}

type DescribeChannelOverallDataResponseBodyMetricDatas struct {
	Nodes []*DescribeChannelOverallDataResponseBodyMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type  *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyMetricDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) SetNodes(v []*DescribeChannelOverallDataResponseBodyMetricDatasNodes) *DescribeChannelOverallDataResponseBodyMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) SetType(v string) *DescribeChannelOverallDataResponseBodyMetricDatas {
	s.Type = &v
	return s
}

type DescribeChannelOverallDataResponseBodyMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	X   *string                `json:"X,omitempty" xml:"X,omitempty"`
	Y   *string                `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyMetricDatasNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetX(v string) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetY(v string) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.Y = &v
	return s
}

type DescribeChannelOverallDataResponseBodyOverallData struct {
	ConnAvgTime         *float32 `json:"ConnAvgTime,omitempty" xml:"ConnAvgTime,omitempty"`
	FiveSecJoinRate     *float32 `json:"FiveSecJoinRate,omitempty" xml:"FiveSecJoinRate,omitempty"`
	TotalAudioStuckRate *float32 `json:"TotalAudioStuckRate,omitempty" xml:"TotalAudioStuckRate,omitempty"`
	TotalVideoStuckRate *float32 `json:"TotalVideoStuckRate,omitempty" xml:"TotalVideoStuckRate,omitempty"`
	TotalVideoVagueRate *float32 `json:"TotalVideoVagueRate,omitempty" xml:"TotalVideoVagueRate,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetConnAvgTime(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.ConnAvgTime = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetFiveSecJoinRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.FiveSecJoinRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalAudioStuckRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalAudioStuckRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalVideoStuckRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalVideoStuckRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalVideoVagueRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalVideoVagueRate = &v
	return s
}

type DescribeChannelOverallDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponse) SetHeaders(v map[string]*string) *DescribeChannelOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelOverallDataResponse) SetStatusCode(v int32) *DescribeChannelOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelOverallDataResponse) SetBody(v *DescribeChannelOverallDataResponseBody) *DescribeChannelOverallDataResponse {
	s.Body = v
	return s
}

type DescribeChannelTopPubUserListRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelTopPubUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListRequest) SetAppId(v string) *DescribeChannelTopPubUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetChannelId(v string) *DescribeChannelTopPubUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetCreatedTs(v int64) *DescribeChannelTopPubUserListRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetDestroyedTs(v int64) *DescribeChannelTopPubUserListRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeChannelTopPubUserListResponseBody struct {
	RequestId            *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TopPubUserDetailList []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList `json:"TopPubUserDetailList,omitempty" xml:"TopPubUserDetailList,omitempty" type:"Repeated"`
}

func (s DescribeChannelTopPubUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBody) SetRequestId(v string) *DescribeChannelTopPubUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBody) SetTopPubUserDetailList(v []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) *DescribeChannelTopPubUserListResponseBody {
	s.TopPubUserDetailList = v
	return s
}

type DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList struct {
	CreatedTs      *int64                                                                        `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs    *int64                                                                        `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration       *int64                                                                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Location       *string                                                                       `json:"Location,omitempty" xml:"Location,omitempty"`
	OnlineDuration *int64                                                                        `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	UserId         *string                                                                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetCreatedTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetDestroyedTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetDuration(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetLocation(v string) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetOnlineDuration(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetOnlinePeriods(v []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetUserId(v string) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.UserId = &v
	return s
}

type DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods struct {
	JoinTs  *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribeChannelTopPubUserListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelTopPubUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelTopPubUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponse) SetHeaders(v map[string]*string) *DescribeChannelTopPubUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelTopPubUserListResponse) SetStatusCode(v int32) *DescribeChannelTopPubUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponse) SetBody(v *DescribeChannelTopPubUserListResponseBody) *DescribeChannelTopPubUserListResponse {
	s.Body = v
	return s
}

type DescribeChannelUserMetricsRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelUserMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsRequest) SetAppId(v string) *DescribeChannelUserMetricsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetChannelId(v string) *DescribeChannelUserMetricsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetCreatedTs(v int64) *DescribeChannelUserMetricsRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetDestroyedTs(v int64) *DescribeChannelUserMetricsRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeChannelUserMetricsResponseBody struct {
	MetricDatas []*DescribeChannelUserMetricsResponseBodyMetricDatas `json:"MetricDatas,omitempty" xml:"MetricDatas,omitempty" type:"Repeated"`
	OverallData *DescribeChannelUserMetricsResponseBodyOverallData   `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBody) SetMetricDatas(v []*DescribeChannelUserMetricsResponseBodyMetricDatas) *DescribeChannelUserMetricsResponseBody {
	s.MetricDatas = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBody) SetOverallData(v *DescribeChannelUserMetricsResponseBodyOverallData) *DescribeChannelUserMetricsResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBody) SetRequestId(v string) *DescribeChannelUserMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelUserMetricsResponseBodyMetricDatas struct {
	Nodes []*DescribeChannelUserMetricsResponseBodyMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type  *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) SetNodes(v []*DescribeChannelUserMetricsResponseBodyMetricDatasNodes) *DescribeChannelUserMetricsResponseBodyMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) SetType(v string) *DescribeChannelUserMetricsResponseBodyMetricDatas {
	s.Type = &v
	return s
}

type DescribeChannelUserMetricsResponseBodyMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	X   *string                `json:"X,omitempty" xml:"X,omitempty"`
	Y   *string                `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatasNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetX(v string) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetY(v string) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.Y = &v
	return s
}

type DescribeChannelUserMetricsResponseBodyOverallData struct {
	TotalBadExpNum   *int64 `json:"TotalBadExpNum,omitempty" xml:"TotalBadExpNum,omitempty"`
	TotalJoinFailNum *int64 `json:"TotalJoinFailNum,omitempty" xml:"TotalJoinFailNum,omitempty"`
	TotalPubUserNum  *int64 `json:"TotalPubUserNum,omitempty" xml:"TotalPubUserNum,omitempty"`
	TotalSubUserNum  *int64 `json:"TotalSubUserNum,omitempty" xml:"TotalSubUserNum,omitempty"`
	TotalUserNum     *int64 `json:"TotalUserNum,omitempty" xml:"TotalUserNum,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalBadExpNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalBadExpNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalJoinFailNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalJoinFailNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalPubUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalPubUserNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalSubUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalSubUserNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalUserNum = &v
	return s
}

type DescribeChannelUserMetricsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelUserMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelUserMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponse) SetHeaders(v map[string]*string) *DescribeChannelUserMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelUserMetricsResponse) SetStatusCode(v int32) *DescribeChannelUserMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelUserMetricsResponse) SetBody(v *DescribeChannelUserMetricsResponseBody) *DescribeChannelUserMetricsResponse {
	s.Body = v
	return s
}

type DescribeEndPointEventListRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	UserIdList  *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s DescribeEndPointEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListRequest) SetAppId(v string) *DescribeEndPointEventListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetChannelId(v string) *DescribeEndPointEventListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetCreatedTs(v int64) *DescribeEndPointEventListRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetDestroyedTs(v int64) *DescribeEndPointEventListRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetUserIdList(v string) *DescribeEndPointEventListRequest {
	s.UserIdList = &v
	return s
}

type DescribeEndPointEventListResponseBody struct {
	Nodes     []*DescribeEndPointEventListResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEndPointEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBody) SetNodes(v []*DescribeEndPointEventListResponseBodyNodes) *DescribeEndPointEventListResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointEventListResponseBody) SetRequestId(v string) *DescribeEndPointEventListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEndPointEventListResponseBodyNodes struct {
	EventDataItems []*DescribeEndPointEventListResponseBodyNodesEventDataItems `json:"EventDataItems,omitempty" xml:"EventDataItems,omitempty" type:"Repeated"`
	UserId         *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodes) SetEventDataItems(v []*DescribeEndPointEventListResponseBodyNodesEventDataItems) *DescribeEndPointEventListResponseBodyNodes {
	s.EventDataItems = v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodes) SetUserId(v string) *DescribeEndPointEventListResponseBodyNodes {
	s.UserId = &v
	return s
}

type DescribeEndPointEventListResponseBodyNodesEventDataItems struct {
	EventList []*DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	Ts        *int64                                                               `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItems) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) SetEventList(v []*DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) *DescribeEndPointEventListResponseBodyNodesEventDataItems {
	s.EventList = v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) SetTs(v int64) *DescribeEndPointEventListResponseBodyNodesEventDataItems {
	s.Ts = &v
	return s
}

type DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList struct {
	Acs        *string `json:"Acs,omitempty" xml:"Acs,omitempty"`
	EventCode  *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	EventName  *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType  *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Os         *string `json:"Os,omitempty" xml:"Os,omitempty"`
	Sdk        *string `json:"Sdk,omitempty" xml:"Sdk,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	TrackCode  *string `json:"TrackCode,omitempty" xml:"TrackCode,omitempty"`
	TrackName  *string `json:"TrackName,omitempty" xml:"TrackName,omitempty"`
	Ts         *int64  `json:"Ts,omitempty" xml:"Ts,omitempty"`
	TsInMs     *string `json:"TsInMs,omitempty" xml:"TsInMs,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetAcs(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.Acs = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetEventCode(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.EventCode = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetEventName(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.EventName = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetEventType(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.EventType = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetOs(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.Os = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetSdk(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.Sdk = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetStreamName(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.StreamName = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetStreamType(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.StreamType = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTrackCode(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.TrackCode = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTrackName(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.TrackName = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTs(v int64) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.Ts = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTsInMs(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.TsInMs = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetUserId(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.UserId = &v
	return s
}

type DescribeEndPointEventListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndPointEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndPointEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponse) SetHeaders(v map[string]*string) *DescribeEndPointEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndPointEventListResponse) SetStatusCode(v int32) *DescribeEndPointEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndPointEventListResponse) SetBody(v *DescribeEndPointEventListResponseBody) *DescribeEndPointEventListResponse {
	s.Body = v
	return s
}

type DescribeEndPointMetricDataRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId     *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs     *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs   *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Metrics       *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	PubCallIdList *string `json:"PubCallIdList,omitempty" xml:"PubCallIdList,omitempty"`
	PubUserId     *string `json:"PubUserId,omitempty" xml:"PubUserId,omitempty"`
	SubUserId     *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeEndPointMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataRequest) SetAppId(v string) *DescribeEndPointMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetChannelId(v string) *DescribeEndPointMetricDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetCreatedTs(v int64) *DescribeEndPointMetricDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetDestroyedTs(v int64) *DescribeEndPointMetricDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetMetrics(v string) *DescribeEndPointMetricDataRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetPubCallIdList(v string) *DescribeEndPointMetricDataRequest {
	s.PubCallIdList = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetPubUserId(v string) *DescribeEndPointMetricDataRequest {
	s.PubUserId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetSubUserId(v string) *DescribeEndPointMetricDataRequest {
	s.SubUserId = &v
	return s
}

type DescribeEndPointMetricDataResponseBody struct {
	PubMetrics []*DescribeEndPointMetricDataResponseBodyPubMetrics `json:"PubMetrics,omitempty" xml:"PubMetrics,omitempty" type:"Repeated"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubMetrics []*DescribeEndPointMetricDataResponseBodySubMetrics `json:"SubMetrics,omitempty" xml:"SubMetrics,omitempty" type:"Repeated"`
}

func (s DescribeEndPointMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBody) SetPubMetrics(v []*DescribeEndPointMetricDataResponseBodyPubMetrics) *DescribeEndPointMetricDataResponseBody {
	s.PubMetrics = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBody) SetRequestId(v string) *DescribeEndPointMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBody) SetSubMetrics(v []*DescribeEndPointMetricDataResponseBodySubMetrics) *DescribeEndPointMetricDataResponseBody {
	s.SubMetrics = v
	return s
}

type DescribeEndPointMetricDataResponseBodyPubMetrics struct {
	Nodes  []*DescribeEndPointMetricDataResponseBodyPubMetricsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type   *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId *string                                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodyPubMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodyPubMetrics) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetNodes(v []*DescribeEndPointMetricDataResponseBodyPubMetricsNodes) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetType(v string) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.Type = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetUserId(v string) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.UserId = &v
	return s
}

type DescribeEndPointMetricDataResponseBodyPubMetricsNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	X   *string                `json:"X,omitempty" xml:"X,omitempty"`
	Y   *string                `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodyPubMetricsNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodyPubMetricsNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetExt(v map[string]interface{}) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.Ext = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetX(v string) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.X = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetY(v string) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.Y = &v
	return s
}

type DescribeEndPointMetricDataResponseBodySubMetrics struct {
	Nodes  []*DescribeEndPointMetricDataResponseBodySubMetricsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type   *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId *string                                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodySubMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodySubMetrics) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetNodes(v []*DescribeEndPointMetricDataResponseBodySubMetricsNodes) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetType(v string) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.Type = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetUserId(v string) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.UserId = &v
	return s
}

type DescribeEndPointMetricDataResponseBodySubMetricsNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	X   *string                `json:"X,omitempty" xml:"X,omitempty"`
	Y   *string                `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodySubMetricsNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodySubMetricsNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetExt(v map[string]interface{}) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.Ext = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetX(v string) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.X = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetY(v string) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.Y = &v
	return s
}

type DescribeEndPointMetricDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndPointMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndPointMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponse) SetHeaders(v map[string]*string) *DescribeEndPointMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndPointMetricDataResponse) SetStatusCode(v int32) *DescribeEndPointMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndPointMetricDataResponse) SetBody(v *DescribeEndPointMetricDataResponseBody) *DescribeEndPointMetricDataResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs   *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	StartTs *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetAppId(v string) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetEndTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetStartTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.StartTs = &v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatResponseBody struct {
	RequestId *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList  []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisFactorDistributionStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) SetStatList(v []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) *DescribeFaultDiagnosisFactorDistributionStatResponseBody {
	s.StatList = v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList struct {
	FactorId  *string  `json:"FactorId,omitempty" xml:"FactorId,omitempty"`
	UserCount *int32   `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	UserRatio *float32 `json:"UserRatio,omitempty" xml:"UserRatio,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetFactorId(v string) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.FactorId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetUserCount(v int32) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.UserCount = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetUserRatio(v float32) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.UserRatio = &v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisFactorDistributionStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetBody(v *DescribeFaultDiagnosisFactorDistributionStatResponseBody) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisOverallDataRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs   *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	StartTs *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetAppId(v string) *DescribeFaultDiagnosisOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetEndTs(v int64) *DescribeFaultDiagnosisOverallDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetStartTs(v int64) *DescribeFaultDiagnosisOverallDataRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetStatDim(v string) *DescribeFaultDiagnosisOverallDataRequest {
	s.StatDim = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBody struct {
	MetricData  *DescribeFaultDiagnosisOverallDataResponseBodyMetricData  `json:"MetricData,omitempty" xml:"MetricData,omitempty" type:"Struct"`
	OverallData *DescribeFaultDiagnosisOverallDataResponseBodyOverallData `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	RequestId   *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetMetricData(v *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.MetricData = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetOverallData(v *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBodyMetricData struct {
	Nodes []*DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) SetNodes(v []*DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) *DescribeFaultDiagnosisOverallDataResponseBodyMetricData {
	s.Nodes = v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	X   *string                `json:"X,omitempty" xml:"X,omitempty"`
	Y   *string                `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetExt(v map[string]interface{}) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.Ext = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetX(v string) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetY(v string) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.Y = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBodyOverallData struct {
	FaultUserCount *int32   `json:"FaultUserCount,omitempty" xml:"FaultUserCount,omitempty"`
	FaultUserRatio *float32 `json:"FaultUserRatio,omitempty" xml:"FaultUserRatio,omitempty"`
	TotalUserCount *int32   `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetFaultUserCount(v int32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.FaultUserCount = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetFaultUserRatio(v float32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.FaultUserRatio = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetTotalUserCount(v int32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.TotalUserCount = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetBody(v *DescribeFaultDiagnosisOverallDataResponseBody) *DescribeFaultDiagnosisOverallDataResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisUserDetailRequest struct {
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId         *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs         *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	FaultType         *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
	QueryCallUserInfo *bool   `json:"QueryCallUserInfo,omitempty" xml:"QueryCallUserInfo,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetAppId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetChannelId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetFaultType(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.FaultType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetQueryCallUserInfo(v bool) *DescribeFaultDiagnosisUserDetailRequest {
	s.QueryCallUserInfo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetUserId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBody struct {
	CallInfo         *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo        `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	FactorList       []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList    `json:"FactorList,omitempty" xml:"FactorList,omitempty" type:"Repeated"`
	FaultMetricData  *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData `json:"FaultMetricData,omitempty" xml:"FaultMetricData,omitempty" type:"Struct"`
	NetworkOperators []*string                                                    `json:"NetworkOperators,omitempty" xml:"NetworkOperators,omitempty" type:"Repeated"`
	RequestId        *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserDetail       *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail      `json:"UserDetail,omitempty" xml:"UserDetail,omitempty" type:"Struct"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetCallInfo(v *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetFactorList(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.FactorList = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetFaultMetricData(v *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.FaultMetricData = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetNetworkOperators(v []*string) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.NetworkOperators = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetUserDetail(v *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.UserDetail = v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyCallInfo struct {
	// App ID。
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallStatus  *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetAppId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetCallStatus(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetChannelId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.Duration = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorList struct {
	FactorId           *string                                                                     `json:"FactorId,omitempty" xml:"FactorId,omitempty"`
	FaultSource        *string                                                                     `json:"FaultSource,omitempty" xml:"FaultSource,omitempty"`
	RelatedEventDatas  []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas  `json:"RelatedEventDatas,omitempty" xml:"RelatedEventDatas,omitempty" type:"Repeated"`
	RelatedMetricDatas []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas `json:"RelatedMetricDatas,omitempty" xml:"RelatedMetricDatas,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetFactorId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.FactorId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetFaultSource(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.FaultSource = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetRelatedEventDatas(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.RelatedEventDatas = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetRelatedMetricDatas(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.RelatedMetricDatas = v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas struct {
	EventDataItems []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems `json:"EventDataItems,omitempty" xml:"EventDataItems,omitempty" type:"Repeated"`
	Role           *string                                                                                  `json:"Role,omitempty" xml:"Role,omitempty"`
	UserId         *string                                                                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetEventDataItems(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.EventDataItems = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetRole(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.Role = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems struct {
	EventList []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	Ts        *int64                                                                                            `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) SetEventList(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems {
	s.EventList = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) SetTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems {
	s.Ts = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList struct {
	Acs        *string `json:"Acs,omitempty" xml:"Acs,omitempty"`
	EventCode  *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	EventName  *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType  *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Os         *string `json:"Os,omitempty" xml:"Os,omitempty"`
	Sdk        *string `json:"Sdk,omitempty" xml:"Sdk,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	TrackCode  *string `json:"TrackCode,omitempty" xml:"TrackCode,omitempty"`
	TrackName  *string `json:"TrackName,omitempty" xml:"TrackName,omitempty"`
	Ts         *int64  `json:"Ts,omitempty" xml:"Ts,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetAcs(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.Acs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetEventCode(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.EventCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetEventName(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.EventName = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetEventType(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.EventType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetOs(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.Os = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetSdk(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.Sdk = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetStreamName(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.StreamName = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetStreamType(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.StreamType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetTrackCode(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.TrackCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetTrackName(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.TrackName = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.Ts = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas struct {
	Nodes  []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Role   *string                                                                          `json:"Role,omitempty" xml:"Role,omitempty"`
	Type   *string                                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId *string                                                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetNodes(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetRole(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Role = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetType(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Type = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	X   *string                `json:"X,omitempty" xml:"X,omitempty"`
	Y   *string                `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetX(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetY(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.Y = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData struct {
	Nodes []*DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) SetNodes(v []*DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData {
	s.Nodes = v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes struct {
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) SetX(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) SetY(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes {
	s.Y = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyUserDetail struct {
	CreatedTs      *int64                                                                 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs    *int64                                                                 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration       *int64                                                                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Location       *string                                                                `json:"Location,omitempty" xml:"Location,omitempty"`
	Network        *string                                                                `json:"Network,omitempty" xml:"Network,omitempty"`
	OnlineDuration *int64                                                                 `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	Os             *string                                                                `json:"Os,omitempty" xml:"Os,omitempty"`
	SdkVersion     *string                                                                `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	UserId         *string                                                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Duration = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetLocation(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Location = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetNetwork(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Network = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOnlineDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOnlinePeriods(v []*DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOs(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Os = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetSdkVersion(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.SdkVersion = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods struct {
	JoinTs  *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) SetJoinTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) SetLeaveTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisUserDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisUserDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetBody(v *DescribeFaultDiagnosisUserDetailResponseBody) *DescribeFaultDiagnosisUserDetailResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisUserListRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTs      *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	FaultTypes *string `json:"FaultTypes,omitempty" xml:"FaultTypes,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTs    *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListRequest) SetAppId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetChannelId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetEndTs(v int64) *DescribeFaultDiagnosisUserListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetFaultTypes(v string) *DescribeFaultDiagnosisUserListRequest {
	s.FaultTypes = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetPageNo(v int32) *DescribeFaultDiagnosisUserListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetPageSize(v int32) *DescribeFaultDiagnosisUserListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetStartTs(v int64) *DescribeFaultDiagnosisUserListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetUserId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserListResponseBody struct {
	PageNo    *int32                                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCnt  *int32                                                `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	UserList  []*DescribeFaultDiagnosisUserListResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetPageNo(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetPageSize(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetTotalCnt(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetUserList(v []*DescribeFaultDiagnosisUserListResponseBodyUserList) *DescribeFaultDiagnosisUserListResponseBody {
	s.UserList = v
	return s
}

type DescribeFaultDiagnosisUserListResponseBodyUserList struct {
	ChannelCreatedTs *int64                                                         `json:"ChannelCreatedTs,omitempty" xml:"ChannelCreatedTs,omitempty"`
	ChannelId        *string                                                        `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs        *int64                                                         `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs      *int64                                                         `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	FaultList        []*DescribeFaultDiagnosisUserListResponseBodyUserListFaultList `json:"FaultList,omitempty" xml:"FaultList,omitempty" type:"Repeated"`
	UserId           *string                                                        `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetChannelCreatedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.ChannelCreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetChannelId(v string) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetFaultList(v []*DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.FaultList = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetUserId(v string) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserListResponseBodyUserListFaultList struct {
	FaultType *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) SetFaultType(v string) *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList {
	s.FaultType = &v
	return s
}

type DescribeFaultDiagnosisUserListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponse) SetBody(v *DescribeFaultDiagnosisUserListResponseBody) *DescribeFaultDiagnosisUserListResponse {
	s.Body = v
	return s
}

type DescribeIceDurPeriodByDaySubTypeRequest struct {
	EndTs    *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	JobType  *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	StartTs  *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeIceDurPeriodByDaySubTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurPeriodByDaySubTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeIceDurPeriodByDaySubTypeRequest) SetEndTs(v int64) *DescribeIceDurPeriodByDaySubTypeRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeRequest) SetJobType(v string) *DescribeIceDurPeriodByDaySubTypeRequest {
	s.JobType = &v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeRequest) SetStartTs(v int64) *DescribeIceDurPeriodByDaySubTypeRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeRequest) SetTimeZone(v string) *DescribeIceDurPeriodByDaySubTypeRequest {
	s.TimeZone = &v
	return s
}

type DescribeIceDurPeriodByDaySubTypeResponseBody struct {
	JobInfoList []*DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList `json:"JobInfoList,omitempty" xml:"JobInfoList,omitempty" type:"Repeated"`
	RequestId   *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIceDurPeriodByDaySubTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurPeriodByDaySubTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIceDurPeriodByDaySubTypeResponseBody) SetJobInfoList(v []*DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList) *DescribeIceDurPeriodByDaySubTypeResponseBody {
	s.JobInfoList = v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeResponseBody) SetRequestId(v string) *DescribeIceDurPeriodByDaySubTypeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList struct {
	DateTs         *int64                                                                   `json:"DateTs,omitempty" xml:"DateTs,omitempty"`
	Duration       *int64                                                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	SubJobInfoList []*DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList `json:"SubJobInfoList,omitempty" xml:"SubJobInfoList,omitempty" type:"Repeated"`
}

func (s DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList) GoString() string {
	return s.String()
}

func (s *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList) SetDateTs(v int64) *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList {
	s.DateTs = &v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList) SetDuration(v int64) *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList {
	s.Duration = &v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList) SetSubJobInfoList(v []*DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList) *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoList {
	s.SubJobInfoList = v
	return s
}

type DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList struct {
	SubJobDuration *int64  `json:"SubJobDuration,omitempty" xml:"SubJobDuration,omitempty"`
	SubJobType     *string `json:"SubJobType,omitempty" xml:"SubJobType,omitempty"`
}

func (s DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList) GoString() string {
	return s.String()
}

func (s *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList) SetSubJobDuration(v int64) *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList {
	s.SubJobDuration = &v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList) SetSubJobType(v string) *DescribeIceDurPeriodByDaySubTypeResponseBodyJobInfoListSubJobInfoList {
	s.SubJobType = &v
	return s
}

type DescribeIceDurPeriodByDaySubTypeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIceDurPeriodByDaySubTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIceDurPeriodByDaySubTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurPeriodByDaySubTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeIceDurPeriodByDaySubTypeResponse) SetHeaders(v map[string]*string) *DescribeIceDurPeriodByDaySubTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeResponse) SetStatusCode(v int32) *DescribeIceDurPeriodByDaySubTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIceDurPeriodByDaySubTypeResponse) SetBody(v *DescribeIceDurPeriodByDaySubTypeResponseBody) *DescribeIceDurPeriodByDaySubTypeResponse {
	s.Body = v
	return s
}

type DescribeIceDurSummaryOverviewRequest struct {
	CurTs    *int64  `json:"CurTs,omitempty" xml:"CurTs,omitempty"`
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeIceDurSummaryOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurSummaryOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeIceDurSummaryOverviewRequest) SetCurTs(v int64) *DescribeIceDurSummaryOverviewRequest {
	s.CurTs = &v
	return s
}

func (s *DescribeIceDurSummaryOverviewRequest) SetTimeZone(v string) *DescribeIceDurSummaryOverviewRequest {
	s.TimeZone = &v
	return s
}

type DescribeIceDurSummaryOverviewResponseBody struct {
	JobInfoList []*DescribeIceDurSummaryOverviewResponseBodyJobInfoList `json:"JobInfoList,omitempty" xml:"JobInfoList,omitempty" type:"Repeated"`
	RequestId   *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIceDurSummaryOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurSummaryOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIceDurSummaryOverviewResponseBody) SetJobInfoList(v []*DescribeIceDurSummaryOverviewResponseBodyJobInfoList) *DescribeIceDurSummaryOverviewResponseBody {
	s.JobInfoList = v
	return s
}

func (s *DescribeIceDurSummaryOverviewResponseBody) SetRequestId(v string) *DescribeIceDurSummaryOverviewResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIceDurSummaryOverviewResponseBodyJobInfoList struct {
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	JobType   *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s DescribeIceDurSummaryOverviewResponseBodyJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurSummaryOverviewResponseBodyJobInfoList) GoString() string {
	return s.String()
}

func (s *DescribeIceDurSummaryOverviewResponseBodyJobInfoList) SetDuration(v int64) *DescribeIceDurSummaryOverviewResponseBodyJobInfoList {
	s.Duration = &v
	return s
}

func (s *DescribeIceDurSummaryOverviewResponseBodyJobInfoList) SetJobType(v string) *DescribeIceDurSummaryOverviewResponseBodyJobInfoList {
	s.JobType = &v
	return s
}

func (s *DescribeIceDurSummaryOverviewResponseBodyJobInfoList) SetTimeRange(v string) *DescribeIceDurSummaryOverviewResponseBodyJobInfoList {
	s.TimeRange = &v
	return s
}

type DescribeIceDurSummaryOverviewResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIceDurSummaryOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIceDurSummaryOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceDurSummaryOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeIceDurSummaryOverviewResponse) SetHeaders(v map[string]*string) *DescribeIceDurSummaryOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeIceDurSummaryOverviewResponse) SetStatusCode(v int32) *DescribeIceDurSummaryOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIceDurSummaryOverviewResponse) SetBody(v *DescribeIceDurSummaryOverviewResponseBody) *DescribeIceDurSummaryOverviewResponse {
	s.Body = v
	return s
}

type DescribePubUserListBySubUserRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	SubUserId   *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribePubUserListBySubUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserRequest) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserRequest) SetAppId(v string) *DescribePubUserListBySubUserRequest {
	s.AppId = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetChannelId(v string) *DescribePubUserListBySubUserRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetCreatedTs(v int64) *DescribePubUserListBySubUserRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetDestroyedTs(v int64) *DescribePubUserListBySubUserRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetSubUserId(v string) *DescribePubUserListBySubUserRequest {
	s.SubUserId = &v
	return s
}

type DescribePubUserListBySubUserResponseBody struct {
	CallStatus        *string                                                      `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	PubUserDetailList []*DescribePubUserListBySubUserResponseBodyPubUserDetailList `json:"PubUserDetailList,omitempty" xml:"PubUserDetailList,omitempty" type:"Repeated"`
	RequestId         *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubUserDetail     *DescribePubUserListBySubUserResponseBodySubUserDetail       `json:"SubUserDetail,omitempty" xml:"SubUserDetail,omitempty" type:"Struct"`
}

func (s DescribePubUserListBySubUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBody) SetCallStatus(v string) *DescribePubUserListBySubUserResponseBody {
	s.CallStatus = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetPubUserDetailList(v []*DescribePubUserListBySubUserResponseBodyPubUserDetailList) *DescribePubUserListBySubUserResponseBody {
	s.PubUserDetailList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetRequestId(v string) *DescribePubUserListBySubUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetSubUserDetail(v *DescribePubUserListBySubUserResponseBodySubUserDetail) *DescribePubUserListBySubUserResponseBody {
	s.SubUserDetail = v
	return s
}

type DescribePubUserListBySubUserResponseBodyPubUserDetailList struct {
	CallIdList     []*string                                                                 `json:"CallIdList,omitempty" xml:"CallIdList,omitempty" type:"Repeated"`
	ClientType     *string                                                                   `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CreatedTs      *int64                                                                    `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs    *int64                                                                    `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration       *int64                                                                    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Location       *string                                                                   `json:"Location,omitempty" xml:"Location,omitempty"`
	Network        *string                                                                   `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList    []*string                                                                 `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	OnlineDuration *int64                                                                    `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	Os             *string                                                                   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList         []*string                                                                 `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles          []*string                                                                 `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	SdkVersion     *string                                                                   `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string                                                                 `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	UserId         *string                                                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserIdAlias    *string                                                                   `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetCallIdList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.CallIdList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetClientType(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.ClientType = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetCreatedTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetDestroyedTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetDuration(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetLocation(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetNetwork(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Network = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetNetworkList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.NetworkList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOnlineDuration(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOnlinePeriods(v []*DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOs(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Os = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOsList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OsList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetRoles(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Roles = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetSdkVersion(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.SdkVersion = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetSdkVersionList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.SdkVersionList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetUserId(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.UserId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetUserIdAlias(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.UserIdAlias = &v
	return s
}

type DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods struct {
	JoinTs  *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribePubUserListBySubUserResponseBodySubUserDetail struct {
	ClientType     *string                                                               `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CreatedTs      *int64                                                                `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs    *int64                                                                `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	Duration       *int64                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Location       *string                                                               `json:"Location,omitempty" xml:"Location,omitempty"`
	Network        *string                                                               `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList    []*string                                                             `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	OnlineDuration *int64                                                                `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	Os             *string                                                               `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList         []*string                                                             `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles          []*string                                                             `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	SdkVersion     *string                                                               `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string                                                             `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	UserId         *string                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserIdAlias    *string                                                               `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetail) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetClientType(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.ClientType = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetCreatedTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetDestroyedTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetDuration(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Duration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetLocation(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Location = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetNetwork(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Network = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetNetworkList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.NetworkList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOnlineDuration(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OnlineDuration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOnlinePeriods(v []*DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OnlinePeriods = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOs(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Os = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOsList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OsList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetRoles(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Roles = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetSdkVersion(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.SdkVersion = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetSdkVersionList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.SdkVersionList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetUserId(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.UserId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetUserIdAlias(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.UserIdAlias = &v
	return s
}

type DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods struct {
	JoinTs  *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) SetJoinTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) SetLeaveTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribePubUserListBySubUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePubUserListBySubUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePubUserListBySubUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponse) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponse) SetHeaders(v map[string]*string) *DescribePubUserListBySubUserResponse {
	s.Headers = v
	return s
}

func (s *DescribePubUserListBySubUserResponse) SetStatusCode(v int32) *DescribePubUserListBySubUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePubUserListBySubUserResponse) SetBody(v *DescribePubUserListBySubUserResponseBody) *DescribePubUserListBySubUserResponse {
	s.Body = v
	return s
}

type DescribeQoeMetricDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreatedTs   *int64  `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataRequest) SetAppId(v string) *DescribeQoeMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetChannelId(v string) *DescribeQoeMetricDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetCreatedTs(v int64) *DescribeQoeMetricDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetDestroyedTs(v int64) *DescribeQoeMetricDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetUserId(v string) *DescribeQoeMetricDataRequest {
	s.UserId = &v
	return s
}

type DescribeQoeMetricDataResponseBody struct {
	AudioData []*DescribeQoeMetricDataResponseBodyAudioData `json:"AudioData,omitempty" xml:"AudioData,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoData []*DescribeQoeMetricDataResponseBodyVideoData `json:"VideoData,omitempty" xml:"VideoData,omitempty" type:"Repeated"`
}

func (s DescribeQoeMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBody) SetAudioData(v []*DescribeQoeMetricDataResponseBodyAudioData) *DescribeQoeMetricDataResponseBody {
	s.AudioData = v
	return s
}

func (s *DescribeQoeMetricDataResponseBody) SetRequestId(v string) *DescribeQoeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBody) SetVideoData(v []*DescribeQoeMetricDataResponseBodyVideoData) *DescribeQoeMetricDataResponseBody {
	s.VideoData = v
	return s
}

type DescribeQoeMetricDataResponseBodyAudioData struct {
	Nodes  []*DescribeQoeMetricDataResponseBodyAudioDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type   *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId *string                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyAudioData) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyAudioData) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetNodes(v []*DescribeQoeMetricDataResponseBodyAudioDataNodes) *DescribeQoeMetricDataResponseBodyAudioData {
	s.Nodes = v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetType(v string) *DescribeQoeMetricDataResponseBodyAudioData {
	s.Type = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetUserId(v string) *DescribeQoeMetricDataResponseBodyAudioData {
	s.UserId = &v
	return s
}

type DescribeQoeMetricDataResponseBodyAudioDataNodes struct {
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyAudioDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyAudioDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) SetX(v string) *DescribeQoeMetricDataResponseBodyAudioDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) SetY(v string) *DescribeQoeMetricDataResponseBodyAudioDataNodes {
	s.Y = &v
	return s
}

type DescribeQoeMetricDataResponseBodyVideoData struct {
	Nodes  []*DescribeQoeMetricDataResponseBodyVideoDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type   *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId *string                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyVideoData) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyVideoData) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetNodes(v []*DescribeQoeMetricDataResponseBodyVideoDataNodes) *DescribeQoeMetricDataResponseBodyVideoData {
	s.Nodes = v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetType(v string) *DescribeQoeMetricDataResponseBodyVideoData {
	s.Type = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetUserId(v string) *DescribeQoeMetricDataResponseBodyVideoData {
	s.UserId = &v
	return s
}

type DescribeQoeMetricDataResponseBodyVideoDataNodes struct {
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyVideoDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyVideoDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) SetX(v string) *DescribeQoeMetricDataResponseBodyVideoDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) SetY(v string) *DescribeQoeMetricDataResponseBodyVideoDataNodes {
	s.Y = &v
	return s
}

type DescribeQoeMetricDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQoeMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQoeMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponse) SetHeaders(v map[string]*string) *DescribeQoeMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQoeMetricDataResponse) SetStatusCode(v int32) *DescribeQoeMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQoeMetricDataResponse) SetBody(v *DescribeQoeMetricDataResponseBody) *DescribeQoeMetricDataResponse {
	s.Body = v
	return s
}

type DescribeQualityAreaDistributionStatDataRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate    *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ParentArea *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
	StartDate  *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetAppId(v string) *DescribeQualityAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityAreaDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeQualityAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityAreaDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeQualityAreaDistributionStatDataResponseBody struct {
	QualityStatDataList []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList `json:"QualityStatDataList,omitempty" xml:"QualityStatDataList,omitempty" type:"Repeated"`
	RequestId           *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) SetQualityStatDataList(v []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityAreaDistributionStatDataResponseBody {
	s.QualityStatDataList = v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList struct {
	AudioDelay                       *int64  `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	AudioSpeakOutDuration            *int64  `json:"AudioSpeakOutDuration,omitempty" xml:"AudioSpeakOutDuration,omitempty"`
	AudioStuckRate                   *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	CallDurationRatio                *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	JoinChannelSucFiveSecRate        *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	JoinChannelSucRate               *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	Name                             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VideoDelay                       *int64  `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	VideoFirstPicDuration            *int64  `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	VideoStuckRate                   *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioDelay(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioSpeakOutDuration(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioSpeakOutDuration = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioStuckRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetCallDurationRatio(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetName(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoDelay(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoStuckRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoStuckRate = &v
	return s
}

type DescribeQualityAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetBody(v *DescribeQualityAreaDistributionStatDataResponseBody) *DescribeQualityAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeQualityDistributionStatDataRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate   *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StatDim   *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeQualityDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataRequest) SetAppId(v string) *DescribeQualityDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetStatDim(v string) *DescribeQualityDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

type DescribeQualityDistributionStatDataResponseBody struct {
	QualityStatDataList []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList `json:"QualityStatDataList,omitempty" xml:"QualityStatDataList,omitempty" type:"Repeated"`
	RequestId           *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponseBody) SetQualityStatDataList(v []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityDistributionStatDataResponseBody {
	s.QualityStatDataList = v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityDistributionStatDataResponseBodyQualityStatDataList struct {
	AudioDelay                       *int64  `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	AudioSpeakOutDuration            *int64  `json:"AudioSpeakOutDuration,omitempty" xml:"AudioSpeakOutDuration,omitempty"`
	AudioStuckRate                   *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	CallDurationRatio                *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	JoinChannelSucFiveSecRate        *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	JoinChannelSucRate               *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	Name                             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VideoDelay                       *int64  `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	VideoFirstPicDuration            *int64  `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	VideoStuckRate                   *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioDelay(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioSpeakOutDuration(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioSpeakOutDuration = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioStuckRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetCallDurationRatio(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetName(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoDelay(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoStuckRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoStuckRate = &v
	return s
}

type DescribeQualityDistributionStatDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponse) SetBody(v *DescribeQualityDistributionStatDataResponseBody) *DescribeQualityDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate   *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetAppId(v string) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataResponseBody struct {
	QualityOsSdkVersionStatDataList []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList `json:"QualityOsSdkVersionStatDataList,omitempty" xml:"QualityOsSdkVersionStatDataList,omitempty" type:"Repeated"`
	RequestId                       *string                                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) SetQualityOsSdkVersionStatDataList(v []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody {
	s.QualityOsSdkVersionStatDataList = v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList struct {
	AudioDelay                       *int64  `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	AudioSpeakOutDuration            *int64  `json:"AudioSpeakOutDuration,omitempty" xml:"AudioSpeakOutDuration,omitempty"`
	AudioStuckRate                   *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	CallDurationRatio                *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	JoinChannelSucFiveSecRate        *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	JoinChannelSucRate               *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	Name                             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Os                               *string `json:"Os,omitempty" xml:"Os,omitempty"`
	VideoDelay                       *int64  `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	VideoFirstPicDuration            *int64  `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	VideoStuckRate                   *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioDelay(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioSpeakOutDuration(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioSpeakOutDuration = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioStuckRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetCallDurationRatio(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetName(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetOs(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.Os = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoDelay(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoStuckRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoStuckRate = &v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityOsSdkVersionDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetBody(v *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeQualityOverallDataRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate   *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Types     *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeQualityOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataRequest) SetAppId(v string) *DescribeQualityOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetEndDate(v int64) *DescribeQualityOverallDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetStartDate(v int64) *DescribeQualityOverallDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetTypes(v string) *DescribeQualityOverallDataRequest {
	s.Types = &v
	return s
}

type DescribeQualityOverallDataResponseBody struct {
	QualityOverallData []*DescribeQualityOverallDataResponseBodyQualityOverallData `json:"QualityOverallData,omitempty" xml:"QualityOverallData,omitempty" type:"Repeated"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBody) SetQualityOverallData(v []*DescribeQualityOverallDataResponseBodyQualityOverallData) *DescribeQualityOverallDataResponseBody {
	s.QualityOverallData = v
	return s
}

func (s *DescribeQualityOverallDataResponseBody) SetRequestId(v string) *DescribeQualityOverallDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityOverallDataResponseBodyQualityOverallData struct {
	Average *string                                                          `json:"Average,omitempty" xml:"Average,omitempty"`
	Nodes   []*DescribeQualityOverallDataResponseBodyQualityOverallDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type    *string                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallData) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetAverage(v string) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Average = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetNodes(v []*DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Nodes = v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetType(v string) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Type = &v
	return s
}

type DescribeQualityOverallDataResponseBodyQualityOverallDataNodes struct {
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) SetX(v string) *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) SetY(v string) *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes {
	s.Y = &v
	return s
}

type DescribeQualityOverallDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponse) SetHeaders(v map[string]*string) *DescribeQualityOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityOverallDataResponse) SetStatusCode(v int32) *DescribeQualityOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityOverallDataResponse) SetBody(v *DescribeQualityOverallDataResponseBody) *DescribeQualityOverallDataResponse {
	s.Body = v
	return s
}

type DescribeUsageAreaDistributionStatDataRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ParentArea *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetAppId(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetEndDate(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetStartDate(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeUsageAreaDistributionStatDataResponseBody struct {
	RequestId         *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageAreaStatList []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList `json:"UsageAreaStatList,omitempty" xml:"UsageAreaStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageAreaDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) SetUsageAreaStatList(v []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) *DescribeUsageAreaDistributionStatDataResponseBody {
	s.UsageAreaStatList = v
	return s
}

type DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList struct {
	AudioCallDuration *int32  `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TotalCallDuration *int32  `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	VideoCallDuration *int32  `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetAudioCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetName(v string) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetTotalCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetVideoCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.VideoCallDuration = &v
	return s
}

type DescribeUsageAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetBody(v *DescribeUsageAreaDistributionStatDataResponseBody) *DescribeUsageAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeUsageDistributionStatDataRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate   *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StatDim   *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeUsageDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataRequest) SetAppId(v string) *DescribeUsageDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetEndDate(v int64) *DescribeUsageDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetStartDate(v int64) *DescribeUsageDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetStatDim(v string) *DescribeUsageDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

type DescribeUsageDistributionStatDataResponseBody struct {
	RequestId     *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageStatList []*DescribeUsageDistributionStatDataResponseBodyUsageStatList `json:"UsageStatList,omitempty" xml:"UsageStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBody) SetUsageStatList(v []*DescribeUsageDistributionStatDataResponseBodyUsageStatList) *DescribeUsageDistributionStatDataResponseBody {
	s.UsageStatList = v
	return s
}

type DescribeUsageDistributionStatDataResponseBodyUsageStatList struct {
	AudioCallDuration *int64  `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TotalCallDuration *int64  `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	VideoCallDuration *int64  `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageDistributionStatDataResponseBodyUsageStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponseBodyUsageStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetAudioCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetCallDurationRatio(v string) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetName(v string) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetTotalCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetVideoCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.VideoCallDuration = &v
	return s
}

type DescribeUsageDistributionStatDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponse) SetBody(v *DescribeUsageDistributionStatDataResponseBody) *DescribeUsageDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate   *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetAppId(v string) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetEndDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetStartDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataResponseBody struct {
	RequestId                 *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageOsSdkVersionStatList []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList `json:"UsageOsSdkVersionStatList,omitempty" xml:"UsageOsSdkVersionStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) SetUsageOsSdkVersionStatList(v []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody {
	s.UsageOsSdkVersionStatList = v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList struct {
	AudioCallDuration *int64  `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Os                *string `json:"Os,omitempty" xml:"Os,omitempty"`
	TotalCallDuration *int64  `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	VideoCallDuration *int64  `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetAudioCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetCallDurationRatio(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetName(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetOs(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.Os = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetTotalCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetVideoCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.VideoCallDuration = &v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageOsSdkVersionDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetBody(v *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeUsageOverallDataRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate   *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Types     *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeUsageOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataRequest) SetAppId(v string) *DescribeUsageOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetEndDate(v int64) *DescribeUsageOverallDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetStartDate(v int64) *DescribeUsageOverallDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetTypes(v string) *DescribeUsageOverallDataRequest {
	s.Types = &v
	return s
}

type DescribeUsageOverallDataResponseBody struct {
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageOverallData []*DescribeUsageOverallDataResponseBodyUsageOverallData `json:"UsageOverallData,omitempty" xml:"UsageOverallData,omitempty" type:"Repeated"`
}

func (s DescribeUsageOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBody) SetRequestId(v string) *DescribeUsageOverallDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBody) SetUsageOverallData(v []*DescribeUsageOverallDataResponseBodyUsageOverallData) *DescribeUsageOverallDataResponseBody {
	s.UsageOverallData = v
	return s
}

type DescribeUsageOverallDataResponseBodyUsageOverallData struct {
	Nodes []*DescribeUsageOverallDataResponseBodyUsageOverallDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Type  *string                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallData) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) SetNodes(v []*DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) *DescribeUsageOverallDataResponseBodyUsageOverallData {
	s.Nodes = v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) SetType(v string) *DescribeUsageOverallDataResponseBodyUsageOverallData {
	s.Type = &v
	return s
}

type DescribeUsageOverallDataResponseBodyUsageOverallDataNodes struct {
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) SetX(v string) *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes {
	s.X = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) SetY(v string) *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes {
	s.Y = &v
	return s
}

type DescribeUsageOverallDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponse) SetHeaders(v map[string]*string) *DescribeUsageOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageOverallDataResponse) SetStatusCode(v int32) *DescribeUsageOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageOverallDataResponse) SetBody(v *DescribeUsageOverallDataResponseBody) *DescribeUsageOverallDataResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("vdc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppConfigWithOptions(request *DescribeAppConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppConfig"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/config/describeAppConfig"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppConfig(request *DescribeAppConfigRequest) (_result *DescribeAppConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppConfigResponse{}
	_body, _err := client.DescribeAppConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCallWithOptions(request *DescribeCallRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.ExtDataType)) {
		query["ExtDataType"] = request.ExtDataType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryExpInfo)) {
		query["QueryExpInfo"] = request.QueryExpInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCall"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeCall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCall(request *DescribeCallRequest) (_result *DescribeCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCallResponse{}
	_body, _err := client.DescribeCallWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCallInfoWithOptions(request *DescribeCallInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCallInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCallInfo"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeCallInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCallInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCallInfo(request *DescribeCallInfoRequest) (_result *DescribeCallInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCallInfoResponse{}
	_body, _err := client.DescribeCallInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCallListWithOptions(request *DescribeCallListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCallListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CallStatus)) {
		query["CallStatus"] = request.CallStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryMode)) {
		query["QueryMode"] = request.QueryMode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCallList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeCallList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCallList(request *DescribeCallListRequest) (_result *DescribeCallListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCallListResponse{}
	_body, _err := client.DescribeCallListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCallUserExpWithOptions(request *DescribeCallUserExpRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCallUserExpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCallUserExp"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeCallUserExp"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCallUserExpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCallUserExp(request *DescribeCallUserExpRequest) (_result *DescribeCallUserExpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCallUserExpResponse{}
	_body, _err := client.DescribeCallUserExpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCallUserListWithOptions(request *DescribeCallUserListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCallUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.ExtDataType)) {
		query["ExtDataType"] = request.ExtDataType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryExpInfo)) {
		query["QueryExpInfo"] = request.QueryExpInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCallUserList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeCallUserList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCallUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCallUserList(request *DescribeCallUserListRequest) (_result *DescribeCallUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCallUserListResponse{}
	_body, _err := client.DescribeCallUserListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelAreaDistributionStatDataWithOptions(request *DescribeChannelAreaDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.ParentArea)) {
		query["ParentArea"] = request.ParentArea
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelAreaDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/channel/describeChannelAreaDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelAreaDistributionStatData(request *DescribeChannelAreaDistributionStatDataRequest) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChannelAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeChannelAreaDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelDistributionStatDataWithOptions(request *DescribeChannelDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/channel/describeChannelDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelDistributionStatData(request *DescribeChannelDistributionStatDataRequest) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChannelDistributionStatDataResponse{}
	_body, _err := client.DescribeChannelDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelJoinInfoWithOptions(request *DescribeChannelJoinInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChannelJoinInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelJoinInfo"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/channel/describeChannelJoinInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelJoinInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelJoinInfo(request *DescribeChannelJoinInfoRequest) (_result *DescribeChannelJoinInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChannelJoinInfoResponse{}
	_body, _err := client.DescribeChannelJoinInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelOverallDataWithOptions(request *DescribeChannelOverallDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChannelOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelOverallData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/channel/describeChannelOverallData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelOverallData(request *DescribeChannelOverallDataRequest) (_result *DescribeChannelOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChannelOverallDataResponse{}
	_body, _err := client.DescribeChannelOverallDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelTopPubUserListWithOptions(request *DescribeChannelTopPubUserListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelTopPubUserList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/channel/describeChannelTopPubUserList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelTopPubUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelTopPubUserList(request *DescribeChannelTopPubUserListRequest) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChannelTopPubUserListResponse{}
	_body, _err := client.DescribeChannelTopPubUserListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelUserMetricsWithOptions(request *DescribeChannelUserMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChannelUserMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelUserMetrics"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/channel/describeChannelUserMetrics"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelUserMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelUserMetrics(request *DescribeChannelUserMetricsRequest) (_result *DescribeChannelUserMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChannelUserMetricsResponse{}
	_body, _err := client.DescribeChannelUserMetricsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEndPointEventListWithOptions(request *DescribeEndPointEventListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeEndPointEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndPointEventList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeEndPointEventList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndPointEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndPointEventList(request *DescribeEndPointEventListRequest) (_result *DescribeEndPointEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEndPointEventListResponse{}
	_body, _err := client.DescribeEndPointEventListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEndPointMetricDataWithOptions(request *DescribeEndPointMetricDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeEndPointMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		query["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.PubCallIdList)) {
		query["PubCallIdList"] = request.PubCallIdList
	}

	if !tea.BoolValue(util.IsUnset(request.PubUserId)) {
		query["PubUserId"] = request.PubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndPointMetricData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeEndPointMetricData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndPointMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndPointMetricData(request *DescribeEndPointMetricDataRequest) (_result *DescribeEndPointMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEndPointMetricDataResponse{}
	_body, _err := client.DescribeEndPointMetricDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisFactorDistributionStatWithOptions(request *DescribeFaultDiagnosisFactorDistributionStatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisFactorDistributionStat"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/diagnosis/describeFaultDiagnosisFactorDistributionStat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisFactorDistributionStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisFactorDistributionStat(request *DescribeFaultDiagnosisFactorDistributionStatRequest) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFaultDiagnosisFactorDistributionStatResponse{}
	_body, _err := client.DescribeFaultDiagnosisFactorDistributionStatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisOverallDataWithOptions(request *DescribeFaultDiagnosisOverallDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisOverallData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/diagnosis/describeFaultDiagnosisOverallData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisOverallData(request *DescribeFaultDiagnosisOverallDataRequest) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFaultDiagnosisOverallDataResponse{}
	_body, _err := client.DescribeFaultDiagnosisOverallDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisUserDetailWithOptions(request *DescribeFaultDiagnosisUserDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.FaultType)) {
		query["FaultType"] = request.FaultType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCallUserInfo)) {
		query["QueryCallUserInfo"] = request.QueryCallUserInfo
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisUserDetail"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/diagnosis/describeFaultDiagnosisUserDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisUserDetail(request *DescribeFaultDiagnosisUserDetailRequest) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFaultDiagnosisUserDetailResponse{}
	_body, _err := client.DescribeFaultDiagnosisUserDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisUserListWithOptions(request *DescribeFaultDiagnosisUserListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.FaultTypes)) {
		query["FaultTypes"] = request.FaultTypes
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisUserList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/diagnosis/describeFaultDiagnosisUserList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaultDiagnosisUserList(request *DescribeFaultDiagnosisUserListRequest) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFaultDiagnosisUserListResponse{}
	_body, _err := client.DescribeFaultDiagnosisUserListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIceDurPeriodByDaySubTypeWithOptions(request *DescribeIceDurPeriodByDaySubTypeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeIceDurPeriodByDaySubTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIceDurPeriodByDaySubType"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/ice/describeIceDurPeriodByDaySubType"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIceDurPeriodByDaySubTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIceDurPeriodByDaySubType(request *DescribeIceDurPeriodByDaySubTypeRequest) (_result *DescribeIceDurPeriodByDaySubTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeIceDurPeriodByDaySubTypeResponse{}
	_body, _err := client.DescribeIceDurPeriodByDaySubTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIceDurSummaryOverviewWithOptions(request *DescribeIceDurSummaryOverviewRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeIceDurSummaryOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurTs)) {
		query["CurTs"] = request.CurTs
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIceDurSummaryOverview"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/ice/describeIceDurSummaryOverview"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIceDurSummaryOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIceDurSummaryOverview(request *DescribeIceDurSummaryOverviewRequest) (_result *DescribeIceDurSummaryOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeIceDurSummaryOverviewResponse{}
	_body, _err := client.DescribeIceDurSummaryOverviewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePubUserListBySubUserWithOptions(request *DescribePubUserListBySubUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribePubUserListBySubUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePubUserListBySubUser"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describePubUserListBySubUser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePubUserListBySubUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePubUserListBySubUser(request *DescribePubUserListBySubUserRequest) (_result *DescribePubUserListBySubUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePubUserListBySubUserResponse{}
	_body, _err := client.DescribePubUserListBySubUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQoeMetricDataWithOptions(request *DescribeQoeMetricDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQoeMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQoeMetricData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/call/describeQoeMetricData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQoeMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQoeMetricData(request *DescribeQoeMetricDataRequest) (_result *DescribeQoeMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQoeMetricDataResponse{}
	_body, _err := client.DescribeQoeMetricDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQualityAreaDistributionStatDataWithOptions(request *DescribeQualityAreaDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ParentArea)) {
		query["ParentArea"] = request.ParentArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityAreaDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/quality/describeQualityAreaDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQualityAreaDistributionStatData(request *DescribeQualityAreaDistributionStatDataRequest) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQualityAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityAreaDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQualityDistributionStatDataWithOptions(request *DescribeQualityDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/quality/describeQualityDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQualityDistributionStatData(request *DescribeQualityDistributionStatDataRequest) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQualityDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQualityOsSdkVersionDistributionStatDataWithOptions(request *DescribeQualityOsSdkVersionDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityOsSdkVersionDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/quality/describeQualityOsSdkVersionDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQualityOsSdkVersionDistributionStatData(request *DescribeQualityOsSdkVersionDistributionStatDataRequest) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQualityOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityOsSdkVersionDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQualityOverallDataWithOptions(request *DescribeQualityOverallDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQualityOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["Types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityOverallData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/quality/describeQualityOverallData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQualityOverallData(request *DescribeQualityOverallDataRequest) (_result *DescribeQualityOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQualityOverallDataResponse{}
	_body, _err := client.DescribeQualityOverallDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsageAreaDistributionStatDataWithOptions(request *DescribeUsageAreaDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ParentArea)) {
		query["ParentArea"] = request.ParentArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageAreaDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/usage/describeUsageAreaDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsageAreaDistributionStatData(request *DescribeUsageAreaDistributionStatDataRequest) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUsageAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageAreaDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsageDistributionStatDataWithOptions(request *DescribeUsageDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/usage/describeUsageDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsageDistributionStatData(request *DescribeUsageDistributionStatDataRequest) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUsageDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsageOsSdkVersionDistributionStatDataWithOptions(request *DescribeUsageOsSdkVersionDistributionStatDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageOsSdkVersionDistributionStatData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/usage/describeUsageOsSdkVersionDistributionStatData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsageOsSdkVersionDistributionStatData(request *DescribeUsageOsSdkVersionDistributionStatDataRequest) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUsageOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageOsSdkVersionDistributionStatDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsageOverallDataWithOptions(request *DescribeUsageOverallDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUsageOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["Types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageOverallData"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/usage/describeUsageOverallData"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsageOverallData(request *DescribeUsageOverallDataRequest) (_result *DescribeUsageOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUsageOverallDataResponse{}
	_body, _err := client.DescribeUsageOverallDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
