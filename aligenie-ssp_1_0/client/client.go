// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type LoginStateInfo struct {
	SceneCode           *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	ThirdUserIdentifier *string `json:"ThirdUserIdentifier,omitempty" xml:"ThirdUserIdentifier,omitempty"`
	ThirdUserType       *string `json:"ThirdUserType,omitempty" xml:"ThirdUserType,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LoginStateInfo) String() string {
	return tea.Prettify(s)
}

func (s LoginStateInfo) GoString() string {
	return s.String()
}

func (s *LoginStateInfo) SetSceneCode(v string) *LoginStateInfo {
	s.SceneCode = &v
	return s
}

func (s *LoginStateInfo) SetThirdUserIdentifier(v string) *LoginStateInfo {
	s.ThirdUserIdentifier = &v
	return s
}

func (s *LoginStateInfo) SetThirdUserType(v string) *LoginStateInfo {
	s.ThirdUserType = &v
	return s
}

func (s *LoginStateInfo) SetUserId(v string) *LoginStateInfo {
	s.UserId = &v
	return s
}

type AddAndRemoveFavoriteContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddAndRemoveFavoriteContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentHeaders) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentHeaders) SetCommonHeaders(v map[string]*string) *AddAndRemoveFavoriteContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAndRemoveFavoriteContentHeaders) SetXAcsAligenieAccessToken(v string) *AddAndRemoveFavoriteContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddAndRemoveFavoriteContentHeaders) SetAuthorization(v string) *AddAndRemoveFavoriteContentHeaders {
	s.Authorization = &v
	return s
}

type AddAndRemoveFavoriteContentRequest struct {
	DeviceInfo                             *AddAndRemoveFavoriteContentRequestDeviceInfo                             `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenAddAndRemoveFavoriteContentRequest *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest `json:"OpenAddAndRemoveFavoriteContentRequest,omitempty" xml:"OpenAddAndRemoveFavoriteContentRequest,omitempty" type:"Struct"`
	UserInfo                               *AddAndRemoveFavoriteContentRequestUserInfo                               `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s AddAndRemoveFavoriteContentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequest) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequest) SetDeviceInfo(v *AddAndRemoveFavoriteContentRequestDeviceInfo) *AddAndRemoveFavoriteContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequest) SetOpenAddAndRemoveFavoriteContentRequest(v *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) *AddAndRemoveFavoriteContentRequest {
	s.OpenAddAndRemoveFavoriteContentRequest = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequest) SetUserInfo(v *AddAndRemoveFavoriteContentRequestUserInfo) *AddAndRemoveFavoriteContentRequest {
	s.UserInfo = v
	return s
}

type AddAndRemoveFavoriteContentRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddAndRemoveFavoriteContentRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetEncodeKey(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetEncodeType(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetId(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetIdType(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetOrganizationId(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest struct {
	FavoriteCmd         *string                                                                                      `json:"FavoriteCmd,omitempty" xml:"FavoriteCmd,omitempty"`
	OpenSourceRawIdPair *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair `json:"OpenSourceRawIdPair,omitempty" xml:"OpenSourceRawIdPair,omitempty" type:"Struct"`
	PackageType         *string                                                                                      `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) SetFavoriteCmd(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest {
	s.FavoriteCmd = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) SetOpenSourceRawIdPair(v *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest {
	s.OpenSourceRawIdPair = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) SetPackageType(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest {
	s.PackageType = &v
	return s
}

type AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair struct {
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	RawId      *string                `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source     *string                `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) SetExtendInfo(v map[string]interface{}) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair {
	s.ExtendInfo = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) SetRawId(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair {
	s.RawId = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) SetSource(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair {
	s.Source = &v
	return s
}

type AddAndRemoveFavoriteContentRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddAndRemoveFavoriteContentRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetEncodeKey(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetEncodeType(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetId(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetIdType(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetOrganizationId(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type AddAndRemoveFavoriteContentShrinkRequest struct {
	DeviceInfoShrink                             *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenAddAndRemoveFavoriteContentRequestShrink *string `json:"OpenAddAndRemoveFavoriteContentRequest,omitempty" xml:"OpenAddAndRemoveFavoriteContentRequest,omitempty"`
	UserInfoShrink                               *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s AddAndRemoveFavoriteContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) SetDeviceInfoShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) SetOpenAddAndRemoveFavoriteContentRequestShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest {
	s.OpenAddAndRemoveFavoriteContentRequestShrink = &v
	return s
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) SetUserInfoShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type AddAndRemoveFavoriteContentResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAndRemoveFavoriteContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentResponseBody) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetCode(v int32) *AddAndRemoveFavoriteContentResponseBody {
	s.Code = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetMessage(v string) *AddAndRemoveFavoriteContentResponseBody {
	s.Message = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetRequestId(v string) *AddAndRemoveFavoriteContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetResult(v bool) *AddAndRemoveFavoriteContentResponseBody {
	s.Result = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetSuccess(v string) *AddAndRemoveFavoriteContentResponseBody {
	s.Success = &v
	return s
}

type AddAndRemoveFavoriteContentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddAndRemoveFavoriteContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAndRemoveFavoriteContentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAndRemoveFavoriteContentResponse) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentResponse) SetHeaders(v map[string]*string) *AddAndRemoveFavoriteContentResponse {
	s.Headers = v
	return s
}

func (s *AddAndRemoveFavoriteContentResponse) SetStatusCode(v int32) *AddAndRemoveFavoriteContentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponse) SetBody(v *AddAndRemoveFavoriteContentResponseBody) *AddAndRemoveFavoriteContentResponse {
	s.Body = v
	return s
}

type AddSubHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddSubHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddSubHeaders) GoString() string {
	return s.String()
}

func (s *AddSubHeaders) SetCommonHeaders(v map[string]*string) *AddSubHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddSubHeaders) SetXAcsAligenieAccessToken(v string) *AddSubHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddSubHeaders) SetAuthorization(v string) *AddSubHeaders {
	s.Authorization = &v
	return s
}

type AddSubRequest struct {
	AddSubscriptionInfoRequest *AddSubRequestAddSubscriptionInfoRequest `json:"AddSubscriptionInfoRequest,omitempty" xml:"AddSubscriptionInfoRequest,omitempty" type:"Struct"`
	DeviceInfo                 *AddSubRequestDeviceInfo                 `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo                   *AddSubRequestUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s AddSubRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSubRequest) GoString() string {
	return s.String()
}

func (s *AddSubRequest) SetAddSubscriptionInfoRequest(v *AddSubRequestAddSubscriptionInfoRequest) *AddSubRequest {
	s.AddSubscriptionInfoRequest = v
	return s
}

func (s *AddSubRequest) SetDeviceInfo(v *AddSubRequestDeviceInfo) *AddSubRequest {
	s.DeviceInfo = v
	return s
}

func (s *AddSubRequest) SetUserInfo(v *AddSubRequestUserInfo) *AddSubRequest {
	s.UserInfo = v
	return s
}

type AddSubRequestAddSubscriptionInfoRequest struct {
	AlbumId       *string                                              `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	DailyStudyCnt *int32                                               `json:"DailyStudyCnt,omitempty" xml:"DailyStudyCnt,omitempty"`
	PlayMode      *string                                              `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	ScheduleInfo  *AddSubRequestAddSubscriptionInfoRequestScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
}

func (s AddSubRequestAddSubscriptionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSubRequestAddSubscriptionInfoRequest) GoString() string {
	return s.String()
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetAlbumId(v string) *AddSubRequestAddSubscriptionInfoRequest {
	s.AlbumId = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetDailyStudyCnt(v int32) *AddSubRequestAddSubscriptionInfoRequest {
	s.DailyStudyCnt = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetPlayMode(v string) *AddSubRequestAddSubscriptionInfoRequest {
	s.PlayMode = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetScheduleInfo(v *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) *AddSubRequestAddSubscriptionInfoRequest {
	s.ScheduleInfo = v
	return s
}

type AddSubRequestAddSubscriptionInfoRequestScheduleInfo struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s AddSubRequestAddSubscriptionInfoRequestScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s AddSubRequestAddSubscriptionInfoRequestScheduleInfo) GoString() string {
	return s.String()
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) SetDaysOfWeek(v []*int32) *AddSubRequestAddSubscriptionInfoRequestScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) SetHour(v int32) *AddSubRequestAddSubscriptionInfoRequestScheduleInfo {
	s.Hour = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) SetMinute(v int32) *AddSubRequestAddSubscriptionInfoRequestScheduleInfo {
	s.Minute = &v
	return s
}

type AddSubRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddSubRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s AddSubRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *AddSubRequestDeviceInfo) SetEncodeKey(v string) *AddSubRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetEncodeType(v string) *AddSubRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetId(v string) *AddSubRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetIdType(v string) *AddSubRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetOrganizationId(v string) *AddSubRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type AddSubRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddSubRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s AddSubRequestUserInfo) GoString() string {
	return s.String()
}

func (s *AddSubRequestUserInfo) SetEncodeKey(v string) *AddSubRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddSubRequestUserInfo) SetEncodeType(v string) *AddSubRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *AddSubRequestUserInfo) SetId(v string) *AddSubRequestUserInfo {
	s.Id = &v
	return s
}

func (s *AddSubRequestUserInfo) SetIdType(v string) *AddSubRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *AddSubRequestUserInfo) SetOrganizationId(v string) *AddSubRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type AddSubShrinkRequest struct {
	AddSubscriptionInfoRequestShrink *string `json:"AddSubscriptionInfoRequest,omitempty" xml:"AddSubscriptionInfoRequest,omitempty"`
	DeviceInfoShrink                 *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink                   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s AddSubShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSubShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddSubShrinkRequest) SetAddSubscriptionInfoRequestShrink(v string) *AddSubShrinkRequest {
	s.AddSubscriptionInfoRequestShrink = &v
	return s
}

func (s *AddSubShrinkRequest) SetDeviceInfoShrink(v string) *AddSubShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *AddSubShrinkRequest) SetUserInfoShrink(v string) *AddSubShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type AddSubResponseBody struct {
	Code      *int32                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddSubResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s AddSubResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSubResponseBody) GoString() string {
	return s.String()
}

func (s *AddSubResponseBody) SetCode(v int32) *AddSubResponseBody {
	s.Code = &v
	return s
}

func (s *AddSubResponseBody) SetMessage(v string) *AddSubResponseBody {
	s.Message = &v
	return s
}

func (s *AddSubResponseBody) SetRequestId(v string) *AddSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSubResponseBody) SetResult(v *AddSubResponseBodyResult) *AddSubResponseBody {
	s.Result = v
	return s
}

type AddSubResponseBodyResult struct {
	AlbumId       *string                               `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	DailyStudyCnt *int32                                `json:"DailyStudyCnt,omitempty" xml:"DailyStudyCnt,omitempty"`
	DeviceId      *string                               `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Id            *int64                                `json:"Id,omitempty" xml:"Id,omitempty"`
	PlayMode      *string                               `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	ScheduleInfo  *AddSubResponseBodyResultScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	UserId        *string                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddSubResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddSubResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddSubResponseBodyResult) SetAlbumId(v string) *AddSubResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *AddSubResponseBodyResult) SetDailyStudyCnt(v int32) *AddSubResponseBodyResult {
	s.DailyStudyCnt = &v
	return s
}

func (s *AddSubResponseBodyResult) SetDeviceId(v string) *AddSubResponseBodyResult {
	s.DeviceId = &v
	return s
}

func (s *AddSubResponseBodyResult) SetId(v int64) *AddSubResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddSubResponseBodyResult) SetPlayMode(v string) *AddSubResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *AddSubResponseBodyResult) SetScheduleInfo(v *AddSubResponseBodyResultScheduleInfo) *AddSubResponseBodyResult {
	s.ScheduleInfo = v
	return s
}

func (s *AddSubResponseBodyResult) SetUserId(v string) *AddSubResponseBodyResult {
	s.UserId = &v
	return s
}

type AddSubResponseBodyResultScheduleInfo struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s AddSubResponseBodyResultScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s AddSubResponseBodyResultScheduleInfo) GoString() string {
	return s.String()
}

func (s *AddSubResponseBodyResultScheduleInfo) SetDaysOfWeek(v []*int32) *AddSubResponseBodyResultScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *AddSubResponseBodyResultScheduleInfo) SetHour(v int32) *AddSubResponseBodyResultScheduleInfo {
	s.Hour = &v
	return s
}

func (s *AddSubResponseBodyResultScheduleInfo) SetMinute(v int32) *AddSubResponseBodyResultScheduleInfo {
	s.Minute = &v
	return s
}

type AddSubResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSubResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSubResponse) GoString() string {
	return s.String()
}

func (s *AddSubResponse) SetHeaders(v map[string]*string) *AddSubResponse {
	s.Headers = v
	return s
}

func (s *AddSubResponse) SetStatusCode(v int32) *AddSubResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSubResponse) SetBody(v *AddSubResponseBody) *AddSubResponse {
	s.Body = v
	return s
}

type AuthLoginWithAligenieUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoHeaders) SetCommonHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoHeaders) SetXAcsAligenieAccessToken(v string) *AuthLoginWithAligenieUserInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoHeaders) SetAuthorization(v string) *AuthLoginWithAligenieUserInfoHeaders {
	s.Authorization = &v
	return s
}

type AuthLoginWithAligenieUserInfoRequest struct {
	EncryptedAligenieUserIdentifier *string `json:"EncryptedAligenieUserIdentifier,omitempty" xml:"EncryptedAligenieUserIdentifier,omitempty"`
	SessionId                       *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoRequest) SetEncryptedAligenieUserIdentifier(v string) *AuthLoginWithAligenieUserInfoRequest {
	s.EncryptedAligenieUserIdentifier = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoRequest) SetSessionId(v string) *AuthLoginWithAligenieUserInfoRequest {
	s.SessionId = &v
	return s
}

type AuthLoginWithAligenieUserInfoResponseBody struct {
	Code      *int32                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AuthLoginWithAligenieUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetCode(v int32) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetMessage(v string) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetRequestId(v string) *AuthLoginWithAligenieUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetResult(v *AuthLoginWithAligenieUserInfoResponseBodyResult) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetSuccess(v bool) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Success = &v
	return s
}

type AuthLoginWithAligenieUserInfoResponseBodyResult struct {
	ExpiredTimeLong       *int64  `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithAligenieUserInfoResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithAligenieUserInfoResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

type AuthLoginWithAligenieUserInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthLoginWithAligenieUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthLoginWithAligenieUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoResponse) SetHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponse) SetStatusCode(v int32) *AuthLoginWithAligenieUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponse) SetBody(v *AuthLoginWithAligenieUserInfoResponseBody) *AuthLoginWithAligenieUserInfoResponse {
	s.Body = v
	return s
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) SetXAcsAligenieAccessToken(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) SetAuthorization(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest struct {
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) SetSessionId(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest {
	s.SessionId = &v
	return s
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody struct {
	Code      *int32                                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetCode(v int32) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetMessage(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetRequestId(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetResult(v *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetSuccess(v bool) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Success = &v
	return s
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult struct {
	ExpiredTimeLong       *int64  `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) SetHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) SetStatusCode(v int32) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) SetBody(v *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse {
	s.Body = v
	return s
}

type AuthLoginWithTaobaoUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) SetCommonHeaders(v map[string]*string) *AuthLoginWithTaobaoUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) SetXAcsAligenieAccessToken(v string) *AuthLoginWithTaobaoUserInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) SetAuthorization(v string) *AuthLoginWithTaobaoUserInfoHeaders {
	s.Authorization = &v
	return s
}

type AuthLoginWithTaobaoUserInfoRequest struct {
	EncryptedTaobaoUserIdentifier *string `json:"EncryptedTaobaoUserIdentifier,omitempty" xml:"EncryptedTaobaoUserIdentifier,omitempty"`
	SessionId                     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoRequest) SetEncryptedTaobaoUserIdentifier(v string) *AuthLoginWithTaobaoUserInfoRequest {
	s.EncryptedTaobaoUserIdentifier = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoRequest) SetSessionId(v string) *AuthLoginWithTaobaoUserInfoRequest {
	s.SessionId = &v
	return s
}

type AuthLoginWithTaobaoUserInfoResponseBody struct {
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AuthLoginWithTaobaoUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetCode(v int32) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetMessage(v string) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetRequestId(v string) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetResult(v *AuthLoginWithTaobaoUserInfoResponseBodyResult) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetSuccess(v bool) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Success = &v
	return s
}

type AuthLoginWithTaobaoUserInfoResponseBodyResult struct {
	ExpiredTimeLong       *int64  `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithTaobaoUserInfoResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithTaobaoUserInfoResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

type AuthLoginWithTaobaoUserInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthLoginWithTaobaoUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthLoginWithTaobaoUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoResponse) SetHeaders(v map[string]*string) *AuthLoginWithTaobaoUserInfoResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponse) SetStatusCode(v int32) *AuthLoginWithTaobaoUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponse) SetBody(v *AuthLoginWithTaobaoUserInfoResponseBody) *AuthLoginWithTaobaoUserInfoResponse {
	s.Body = v
	return s
}

type AuthLoginWithThirdUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuthLoginWithThirdUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoHeaders) SetCommonHeaders(v map[string]*string) *AuthLoginWithThirdUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthLoginWithThirdUserInfoHeaders) SetXAcsAligenieAccessToken(v string) *AuthLoginWithThirdUserInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoHeaders) SetAuthorization(v string) *AuthLoginWithThirdUserInfoHeaders {
	s.Authorization = &v
	return s
}

type AuthLoginWithThirdUserInfoRequest struct {
	ExtInfo             map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	SceneCode           *string                `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	ThirdUserIdentifier *string                `json:"ThirdUserIdentifier,omitempty" xml:"ThirdUserIdentifier,omitempty"`
	ThirdUserType       *string                `json:"ThirdUserType,omitempty" xml:"ThirdUserType,omitempty"`
}

func (s AuthLoginWithThirdUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoRequest) SetExtInfo(v map[string]interface{}) *AuthLoginWithThirdUserInfoRequest {
	s.ExtInfo = v
	return s
}

func (s *AuthLoginWithThirdUserInfoRequest) SetSceneCode(v string) *AuthLoginWithThirdUserInfoRequest {
	s.SceneCode = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoRequest) SetThirdUserIdentifier(v string) *AuthLoginWithThirdUserInfoRequest {
	s.ThirdUserIdentifier = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoRequest) SetThirdUserType(v string) *AuthLoginWithThirdUserInfoRequest {
	s.ThirdUserType = &v
	return s
}

type AuthLoginWithThirdUserInfoShrinkRequest struct {
	ExtInfoShrink       *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	SceneCode           *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	ThirdUserIdentifier *string `json:"ThirdUserIdentifier,omitempty" xml:"ThirdUserIdentifier,omitempty"`
	ThirdUserType       *string `json:"ThirdUserType,omitempty" xml:"ThirdUserType,omitempty"`
}

func (s AuthLoginWithThirdUserInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetExtInfoShrink(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.ExtInfoShrink = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetSceneCode(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.SceneCode = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetThirdUserIdentifier(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.ThirdUserIdentifier = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetThirdUserType(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.ThirdUserType = &v
	return s
}

type AuthLoginWithThirdUserInfoResponseBody struct {
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DataObj   *AuthLoginWithThirdUserInfoResponseBodyDataObj `json:"DataObj,omitempty" xml:"DataObj,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AuthLoginWithThirdUserInfoResponseBodyResult  `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithThirdUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetCode(v int32) *AuthLoginWithThirdUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetDataObj(v *AuthLoginWithThirdUserInfoResponseBodyDataObj) *AuthLoginWithThirdUserInfoResponseBody {
	s.DataObj = v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetMessage(v string) *AuthLoginWithThirdUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetRequestId(v string) *AuthLoginWithThirdUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetResult(v *AuthLoginWithThirdUserInfoResponseBodyResult) *AuthLoginWithThirdUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetSuccess(v bool) *AuthLoginWithThirdUserInfoResponseBody {
	s.Success = &v
	return s
}

type AuthLoginWithThirdUserInfoResponseBodyDataObj struct {
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithThirdUserInfoResponseBodyDataObj) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponseBodyDataObj) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponseBodyDataObj) SetSessionId(v string) *AuthLoginWithThirdUserInfoResponseBodyDataObj {
	s.SessionId = &v
	return s
}

type AuthLoginWithThirdUserInfoResponseBodyResult struct {
	ExpiredTimeLong       *int64  `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithThirdUserInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithThirdUserInfoResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithThirdUserInfoResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

type AuthLoginWithThirdUserInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthLoginWithThirdUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthLoginWithThirdUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponse) SetHeaders(v map[string]*string) *AuthLoginWithThirdUserInfoResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponse) SetStatusCode(v int32) *AuthLoginWithThirdUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponse) SetBody(v *AuthLoginWithThirdUserInfoResponseBody) *AuthLoginWithThirdUserInfoResponse {
	s.Body = v
	return s
}

type CheckAuthCodeBindForExtHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CheckAuthCodeBindForExtHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtHeaders) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtHeaders) SetCommonHeaders(v map[string]*string) *CheckAuthCodeBindForExtHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckAuthCodeBindForExtHeaders) SetXAcsAligenieAccessToken(v string) *CheckAuthCodeBindForExtHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CheckAuthCodeBindForExtHeaders) SetAuthorization(v string) *CheckAuthCodeBindForExtHeaders {
	s.Authorization = &v
	return s
}

type CheckAuthCodeBindForExtRequest struct {
	AuthCode   *string                                 `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	EncodeKey  *string                                 `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType *string                                 `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	UserInfo   *CheckAuthCodeBindForExtRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CheckAuthCodeBindForExtRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtRequest) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtRequest) SetAuthCode(v string) *CheckAuthCodeBindForExtRequest {
	s.AuthCode = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequest) SetEncodeKey(v string) *CheckAuthCodeBindForExtRequest {
	s.EncodeKey = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequest) SetEncodeType(v string) *CheckAuthCodeBindForExtRequest {
	s.EncodeType = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequest) SetUserInfo(v *CheckAuthCodeBindForExtRequestUserInfo) *CheckAuthCodeBindForExtRequest {
	s.UserInfo = v
	return s
}

type CheckAuthCodeBindForExtRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CheckAuthCodeBindForExtRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetEncodeKey(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetEncodeType(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetId(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetIdType(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetOrganizationId(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type CheckAuthCodeBindForExtShrinkRequest struct {
	AuthCode       *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CheckAuthCodeBindForExtShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetAuthCode(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.AuthCode = &v
	return s
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetEncodeKey(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.EncodeKey = &v
	return s
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetEncodeType(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.EncodeType = &v
	return s
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetUserInfoShrink(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type CheckAuthCodeBindForExtResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckAuthCodeBindForExtResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckAuthCodeBindForExtResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBody) SetCode(v int32) *CheckAuthCodeBindForExtResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBody) SetMessage(v string) *CheckAuthCodeBindForExtResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBody) SetRequestId(v string) *CheckAuthCodeBindForExtResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBody) SetResult(v *CheckAuthCodeBindForExtResponseBodyResult) *CheckAuthCodeBindForExtResponseBody {
	s.Result = v
	return s
}

type CheckAuthCodeBindForExtResponseBodyResult struct {
	DeviceOpenInfo *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo `json:"DeviceOpenInfo,omitempty" xml:"DeviceOpenInfo,omitempty" type:"Struct"`
	UserOpenInfo   *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo   `json:"UserOpenInfo,omitempty" xml:"UserOpenInfo,omitempty" type:"Struct"`
}

func (s CheckAuthCodeBindForExtResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBodyResult) SetDeviceOpenInfo(v *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) *CheckAuthCodeBindForExtResponseBodyResult {
	s.DeviceOpenInfo = v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResult) SetUserOpenInfo(v *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) *CheckAuthCodeBindForExtResponseBodyResult {
	s.UserOpenInfo = v
	return s
}

type CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo struct {
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) SetId(v string) *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo {
	s.Id = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) SetIdType(v string) *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo {
	s.IdType = &v
	return s
}

type CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo struct {
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) SetId(v string) *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo {
	s.Id = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) SetIdType(v string) *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo {
	s.IdType = &v
	return s
}

type CheckAuthCodeBindForExtResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckAuthCodeBindForExtResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckAuthCodeBindForExtResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponse) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponse) SetHeaders(v map[string]*string) *CheckAuthCodeBindForExtResponse {
	s.Headers = v
	return s
}

func (s *CheckAuthCodeBindForExtResponse) SetStatusCode(v int32) *CheckAuthCodeBindForExtResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponse) SetBody(v *CheckAuthCodeBindForExtResponseBody) *CheckAuthCodeBindForExtResponse {
	s.Body = v
	return s
}

type CreateAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateAlarmHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmHeaders) GoString() string {
	return s.String()
}

func (s *CreateAlarmHeaders) SetCommonHeaders(v map[string]*string) *CreateAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAlarmHeaders) SetXAcsAligenieAccessToken(v string) *CreateAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateAlarmHeaders) SetAuthorization(v string) *CreateAlarmHeaders {
	s.Authorization = &v
	return s
}

type CreateAlarmRequest struct {
	DeviceInfo *CreateAlarmRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *CreateAlarmRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *CreateAlarmRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequest) SetDeviceInfo(v *CreateAlarmRequestDeviceInfo) *CreateAlarmRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreateAlarmRequest) SetPayload(v *CreateAlarmRequestPayload) *CreateAlarmRequest {
	s.Payload = v
	return s
}

func (s *CreateAlarmRequest) SetUserInfo(v *CreateAlarmRequestUserInfo) *CreateAlarmRequest {
	s.UserInfo = v
	return s
}

type CreateAlarmRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateAlarmRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestDeviceInfo) SetEncodeKey(v string) *CreateAlarmRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetEncodeType(v string) *CreateAlarmRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetId(v string) *CreateAlarmRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetIdType(v string) *CreateAlarmRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetOrganizationId(v string) *CreateAlarmRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type CreateAlarmRequestPayload struct {
	MusicInfo    *CreateAlarmRequestPayloadMusicInfo    `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	ScheduleInfo *CreateAlarmRequestPayloadScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	Volume       *int32                                 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateAlarmRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestPayload) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayload) SetMusicInfo(v *CreateAlarmRequestPayloadMusicInfo) *CreateAlarmRequestPayload {
	s.MusicInfo = v
	return s
}

func (s *CreateAlarmRequestPayload) SetScheduleInfo(v *CreateAlarmRequestPayloadScheduleInfo) *CreateAlarmRequestPayload {
	s.ScheduleInfo = v
	return s
}

func (s *CreateAlarmRequestPayload) SetVolume(v int32) *CreateAlarmRequestPayload {
	s.Volume = &v
	return s
}

type CreateAlarmRequestPayloadMusicInfo struct {
	MusicId       *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName     *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	MusicUrl      *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s CreateAlarmRequestPayloadMusicInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestPayloadMusicInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicId(v int64) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicId = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicName(v string) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicName = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicType(v int64) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicType = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicTypeName(v string) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicUrl(v string) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicUrl = &v
	return s
}

type CreateAlarmRequestPayloadScheduleInfo struct {
	Once                *CreateAlarmRequestPayloadScheduleInfoOnce                `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	StatutoryWorkingDay *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	Type                *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly              *CreateAlarmRequestPayloadScheduleInfoWeekly              `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s CreateAlarmRequestPayloadScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetOnce(v *CreateAlarmRequestPayloadScheduleInfoOnce) *CreateAlarmRequestPayloadScheduleInfo {
	s.Once = v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetStatutoryWorkingDay(v *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) *CreateAlarmRequestPayloadScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetType(v string) *CreateAlarmRequestPayloadScheduleInfo {
	s.Type = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetWeekly(v *CreateAlarmRequestPayloadScheduleInfoWeekly) *CreateAlarmRequestPayloadScheduleInfo {
	s.Weekly = v
	return s
}

type CreateAlarmRequestPayloadScheduleInfoOnce struct {
	Day    *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month  *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	Year   *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s CreateAlarmRequestPayloadScheduleInfoOnce) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetDay(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetHour(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetMinute(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetMonth(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetYear(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Year = &v
	return s
}

type CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay struct {
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetHour(v int32) *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

type CreateAlarmRequestPayloadScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s CreateAlarmRequestPayloadScheduleInfoWeekly) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *CreateAlarmRequestPayloadScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) SetHour(v int32) *CreateAlarmRequestPayloadScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) SetMinute(v int32) *CreateAlarmRequestPayloadScheduleInfoWeekly {
	s.Minute = &v
	return s
}

type CreateAlarmRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateAlarmRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestUserInfo) SetEncodeKey(v string) *CreateAlarmRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetEncodeType(v string) *CreateAlarmRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetId(v string) *CreateAlarmRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetIdType(v string) *CreateAlarmRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetOrganizationId(v string) *CreateAlarmRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type CreateAlarmShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateAlarmShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmShrinkRequest) SetDeviceInfoShrink(v string) *CreateAlarmShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreateAlarmShrinkRequest) SetPayloadShrink(v string) *CreateAlarmShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *CreateAlarmShrinkRequest) SetUserInfoShrink(v string) *CreateAlarmShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type CreateAlarmResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int64  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponseBody) SetCode(v int32) *CreateAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAlarmResponseBody) SetMessage(v string) *CreateAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAlarmResponseBody) SetRequestId(v string) *CreateAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlarmResponseBody) SetResult(v int64) *CreateAlarmResponseBody {
	s.Result = &v
	return s
}

type CreateAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponse) SetHeaders(v map[string]*string) *CreateAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateAlarmResponse) SetStatusCode(v int32) *CreateAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlarmResponse) SetBody(v *CreateAlarmResponseBody) *CreateAlarmResponse {
	s.Body = v
	return s
}

type CreatePlayingListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreatePlayingListHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListHeaders) GoString() string {
	return s.String()
}

func (s *CreatePlayingListHeaders) SetCommonHeaders(v map[string]*string) *CreatePlayingListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePlayingListHeaders) SetXAcsAligenieAccessToken(v string) *CreatePlayingListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreatePlayingListHeaders) SetAuthorization(v string) *CreatePlayingListHeaders {
	s.Authorization = &v
	return s
}

type CreatePlayingListRequest struct {
	DeviceInfo                   *CreatePlayingListRequestDeviceInfo                   `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenCreatePlayingListRequest *CreatePlayingListRequestOpenCreatePlayingListRequest `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty" type:"Struct"`
	UserInfo                     *CreatePlayingListRequestUserInfo                     `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreatePlayingListRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequest) SetDeviceInfo(v *CreatePlayingListRequestDeviceInfo) *CreatePlayingListRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreatePlayingListRequest) SetOpenCreatePlayingListRequest(v *CreatePlayingListRequestOpenCreatePlayingListRequest) *CreatePlayingListRequest {
	s.OpenCreatePlayingListRequest = v
	return s
}

func (s *CreatePlayingListRequest) SetUserInfo(v *CreatePlayingListRequestUserInfo) *CreatePlayingListRequest {
	s.UserInfo = v
	return s
}

type CreatePlayingListRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreatePlayingListRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestDeviceInfo) SetEncodeKey(v string) *CreatePlayingListRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetEncodeType(v string) *CreatePlayingListRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetId(v string) *CreatePlayingListRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetIdType(v string) *CreatePlayingListRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetOrganizationId(v string) *CreatePlayingListRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type CreatePlayingListRequestOpenCreatePlayingListRequest struct {
	ContentList        []*CreatePlayingListRequestOpenCreatePlayingListRequestContentList `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	ContentType        *string                                                            `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ExtendInfo         map[string]interface{}                                             `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	Index              *int32                                                             `json:"Index,omitempty" xml:"Index,omitempty"`
	NeedAlbumContinued *bool                                                              `json:"NeedAlbumContinued,omitempty" xml:"NeedAlbumContinued,omitempty"`
	PlayFrom           *string                                                            `json:"PlayFrom,omitempty" xml:"PlayFrom,omitempty"`
	PlayMode           *string                                                            `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetContentList(v []*CreatePlayingListRequestOpenCreatePlayingListRequestContentList) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.ContentList = v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetContentType(v string) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.ContentType = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetExtendInfo(v map[string]interface{}) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.ExtendInfo = v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetIndex(v int32) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.Index = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetNeedAlbumContinued(v bool) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.NeedAlbumContinued = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetPlayFrom(v string) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.PlayFrom = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetPlayMode(v string) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.PlayMode = &v
	return s
}

type CreatePlayingListRequestOpenCreatePlayingListRequestContentList struct {
	RawId  *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequestContentList) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequestContentList) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequestContentList) SetRawId(v string) *CreatePlayingListRequestOpenCreatePlayingListRequestContentList {
	s.RawId = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequestContentList) SetSource(v string) *CreatePlayingListRequestOpenCreatePlayingListRequestContentList {
	s.Source = &v
	return s
}

type CreatePlayingListRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreatePlayingListRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestUserInfo) SetEncodeKey(v string) *CreatePlayingListRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetEncodeType(v string) *CreatePlayingListRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetId(v string) *CreatePlayingListRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetIdType(v string) *CreatePlayingListRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetOrganizationId(v string) *CreatePlayingListRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type CreatePlayingListShrinkRequest struct {
	DeviceInfoShrink                   *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenCreatePlayingListRequestShrink *string `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty"`
	UserInfoShrink                     *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreatePlayingListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListShrinkRequest) SetDeviceInfoShrink(v string) *CreatePlayingListShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreatePlayingListShrinkRequest) SetOpenCreatePlayingListRequestShrink(v string) *CreatePlayingListShrinkRequest {
	s.OpenCreatePlayingListRequestShrink = &v
	return s
}

func (s *CreatePlayingListShrinkRequest) SetUserInfoShrink(v string) *CreatePlayingListShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type CreatePlayingListResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreatePlayingListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *string                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePlayingListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBody) SetCode(v int32) *CreatePlayingListResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePlayingListResponseBody) SetMessage(v string) *CreatePlayingListResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePlayingListResponseBody) SetRequestId(v string) *CreatePlayingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePlayingListResponseBody) SetResult(v *CreatePlayingListResponseBodyResult) *CreatePlayingListResponseBody {
	s.Result = v
	return s
}

func (s *CreatePlayingListResponseBody) SetSuccess(v string) *CreatePlayingListResponseBody {
	s.Success = &v
	return s
}

type CreatePlayingListResponseBodyResult struct {
	AlbumName        *string                                   `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	AlbumRawId       *string                                   `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	AudioLength      *int32                                    `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	Copyright        *int32                                    `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover            *CreatePlayingListResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	DefaultPlayOrder *int32                                    `json:"DefaultPlayOrder,omitempty" xml:"DefaultPlayOrder,omitempty"`
	ItemUrl          *string                                   `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Liked            *bool                                     `json:"Liked,omitempty" xml:"Liked,omitempty"`
	LyricUrl         *string                                   `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	PlayMode         *string                                   `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	Pos              *int32                                    `json:"Pos,omitempty" xml:"Pos,omitempty"`
	Progress         *int32                                    `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RawId            *string                                   `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singer           *string                                   `json:"Singer,omitempty" xml:"Singer,omitempty"`
	Source           *string                                   `json:"Source,omitempty" xml:"Source,omitempty"`
	Title            *string                                   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid            *string                                   `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s CreatePlayingListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBodyResult) SetAlbumName(v string) *CreatePlayingListResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetAlbumRawId(v string) *CreatePlayingListResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetAudioLength(v int32) *CreatePlayingListResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetCopyright(v int32) *CreatePlayingListResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetCover(v *CreatePlayingListResponseBodyResultCover) *CreatePlayingListResponseBodyResult {
	s.Cover = v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetDefaultPlayOrder(v int32) *CreatePlayingListResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetItemUrl(v string) *CreatePlayingListResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetLiked(v bool) *CreatePlayingListResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetLyricUrl(v string) *CreatePlayingListResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetPlayMode(v string) *CreatePlayingListResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetPos(v int32) *CreatePlayingListResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetProgress(v int32) *CreatePlayingListResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetRawId(v string) *CreatePlayingListResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetSinger(v string) *CreatePlayingListResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetSource(v string) *CreatePlayingListResponseBodyResult {
	s.Source = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetTitle(v string) *CreatePlayingListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetType(v string) *CreatePlayingListResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetValid(v string) *CreatePlayingListResponseBodyResult {
	s.Valid = &v
	return s
}

type CreatePlayingListResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s CreatePlayingListResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBodyResultCover) SetCanResize(v bool) *CreatePlayingListResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetImg(v string) *CreatePlayingListResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetLarge(v string) *CreatePlayingListResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetMediam(v string) *CreatePlayingListResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetMedium(v string) *CreatePlayingListResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetSmall(v string) *CreatePlayingListResponseBodyResultCover {
	s.Small = &v
	return s
}

type CreatePlayingListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePlayingListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePlayingListResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListResponse) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponse) SetHeaders(v map[string]*string) *CreatePlayingListResponse {
	s.Headers = v
	return s
}

func (s *CreatePlayingListResponse) SetStatusCode(v int32) *CreatePlayingListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePlayingListResponse) SetBody(v *CreatePlayingListResponseBody) *CreatePlayingListResponse {
	s.Body = v
	return s
}

type CreateScheduleTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateScheduleTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateScheduleTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScheduleTaskHeaders) SetXAcsAligenieAccessToken(v string) *CreateScheduleTaskHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateScheduleTaskHeaders) SetAuthorization(v string) *CreateScheduleTaskHeaders {
	s.Authorization = &v
	return s
}

type CreateScheduleTaskRequest struct {
	DeviceInfo *CreateScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *CreateScheduleTaskRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *CreateScheduleTaskRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateScheduleTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequest) SetDeviceInfo(v *CreateScheduleTaskRequestDeviceInfo) *CreateScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreateScheduleTaskRequest) SetPayload(v *CreateScheduleTaskRequestPayload) *CreateScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *CreateScheduleTaskRequest) SetUserInfo(v *CreateScheduleTaskRequestUserInfo) *CreateScheduleTaskRequest {
	s.UserInfo = v
	return s
}

type CreateScheduleTaskRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateScheduleTaskRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetId(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetIdType(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type CreateScheduleTaskRequestPayload struct {
	ActionDTOs   []*CreateScheduleTaskRequestPayloadActionDTOs `json:"ActionDTOs,omitempty" xml:"ActionDTOs,omitempty" type:"Repeated"`
	IdempotentId *string                                       `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	ScheduleDTO  *CreateScheduleTaskRequestPayloadScheduleDTO  `json:"ScheduleDTO,omitempty" xml:"ScheduleDTO,omitempty" type:"Struct"`
}

func (s CreateScheduleTaskRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayload) SetActionDTOs(v []*CreateScheduleTaskRequestPayloadActionDTOs) *CreateScheduleTaskRequestPayload {
	s.ActionDTOs = v
	return s
}

func (s *CreateScheduleTaskRequestPayload) SetIdempotentId(v string) *CreateScheduleTaskRequestPayload {
	s.IdempotentId = &v
	return s
}

func (s *CreateScheduleTaskRequestPayload) SetScheduleDTO(v *CreateScheduleTaskRequestPayloadScheduleDTO) *CreateScheduleTaskRequestPayload {
	s.ScheduleDTO = v
	return s
}

type CreateScheduleTaskRequestPayloadActionDTOs struct {
	CustomAction map[string]interface{} `json:"customAction,omitempty" xml:"customAction,omitempty"`
}

func (s CreateScheduleTaskRequestPayloadActionDTOs) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadActionDTOs) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadActionDTOs) SetCustomAction(v map[string]interface{}) *CreateScheduleTaskRequestPayloadActionDTOs {
	s.CustomAction = v
	return s
}

type CreateScheduleTaskRequestPayloadScheduleDTO struct {
	Once                *CreateScheduleTaskRequestPayloadScheduleDTOOnce                `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	ScheduleEndTime     *int64                                                          `json:"ScheduleEndTime,omitempty" xml:"ScheduleEndTime,omitempty"`
	ScheduleStartTime   *int64                                                          `json:"ScheduleStartTime,omitempty" xml:"ScheduleStartTime,omitempty"`
	ScheduleType        *string                                                         `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	StatutoryWorkingDay *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	Weekly              *CreateScheduleTaskRequestPayloadScheduleDTOWeekly              `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s CreateScheduleTaskRequestPayloadScheduleDTO) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTO) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetOnce(v *CreateScheduleTaskRequestPayloadScheduleDTOOnce) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.Once = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetScheduleEndTime(v int64) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.ScheduleEndTime = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetScheduleStartTime(v int64) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.ScheduleStartTime = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetScheduleType(v string) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetStatutoryWorkingDay(v *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.StatutoryWorkingDay = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetWeekly(v *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.Weekly = v
	return s
}

type CreateScheduleTaskRequestPayloadScheduleDTOOnce struct {
	Day    *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month  *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	Year   *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOOnce) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOOnce) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetDay(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Day = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetHour(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Hour = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetMinute(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Minute = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetMonth(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Month = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetYear(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Year = &v
	return s
}

type CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay struct {
	Hours   []*int32 `json:"Hours,omitempty" xml:"Hours,omitempty" type:"Repeated"`
	Minutes []*int32 `json:"Minutes,omitempty" xml:"Minutes,omitempty" type:"Repeated"`
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) SetHours(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay {
	s.Hours = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) SetMinutes(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay {
	s.Minutes = v
	return s
}

type CreateScheduleTaskRequestPayloadScheduleDTOWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hours      []*int32 `json:"Hours,omitempty" xml:"Hours,omitempty" type:"Repeated"`
	Minutes    []*int32 `json:"Minutes,omitempty" xml:"Minutes,omitempty" type:"Repeated"`
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOWeekly) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOWeekly) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) SetDaysOfWeek(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) SetHours(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOWeekly {
	s.Hours = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) SetMinutes(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOWeekly {
	s.Minutes = v
	return s
}

type CreateScheduleTaskRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateScheduleTaskRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestUserInfo) SetEncodeKey(v string) *CreateScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetEncodeType(v string) *CreateScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetId(v string) *CreateScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetIdType(v string) *CreateScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetOrganizationId(v string) *CreateScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type CreateScheduleTaskShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateScheduleTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskShrinkRequest) SetDeviceInfoShrink(v string) *CreateScheduleTaskShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreateScheduleTaskShrinkRequest) SetPayloadShrink(v string) *CreateScheduleTaskShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *CreateScheduleTaskShrinkRequest) SetUserInfoShrink(v string) *CreateScheduleTaskShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type CreateScheduleTaskResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int64  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateScheduleTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskResponseBody) SetCode(v int32) *CreateScheduleTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScheduleTaskResponseBody) SetMessage(v string) *CreateScheduleTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduleTaskResponseBody) SetRequestId(v string) *CreateScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleTaskResponseBody) SetResult(v int64) *CreateScheduleTaskResponseBody {
	s.Result = &v
	return s
}

type CreateScheduleTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduleTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskResponse) SetHeaders(v map[string]*string) *CreateScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleTaskResponse) SetStatusCode(v int32) *CreateScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleTaskResponse) SetBody(v *CreateScheduleTaskResponseBody) *CreateScheduleTaskResponse {
	s.Body = v
	return s
}

type DeleteAlarmsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteAlarmsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsHeaders) SetCommonHeaders(v map[string]*string) *DeleteAlarmsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAlarmsHeaders) SetXAcsAligenieAccessToken(v string) *DeleteAlarmsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteAlarmsHeaders) SetAuthorization(v string) *DeleteAlarmsHeaders {
	s.Authorization = &v
	return s
}

type DeleteAlarmsRequest struct {
	DeviceInfo *DeleteAlarmsRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *DeleteAlarmsRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *DeleteAlarmsRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeleteAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequest) SetDeviceInfo(v *DeleteAlarmsRequestDeviceInfo) *DeleteAlarmsRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeleteAlarmsRequest) SetPayload(v *DeleteAlarmsRequestPayload) *DeleteAlarmsRequest {
	s.Payload = v
	return s
}

func (s *DeleteAlarmsRequest) SetUserInfo(v *DeleteAlarmsRequestUserInfo) *DeleteAlarmsRequest {
	s.UserInfo = v
	return s
}

type DeleteAlarmsRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteAlarmsRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequestDeviceInfo) SetEncodeKey(v string) *DeleteAlarmsRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetEncodeType(v string) *DeleteAlarmsRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetId(v string) *DeleteAlarmsRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetIdType(v string) *DeleteAlarmsRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetOrganizationId(v string) *DeleteAlarmsRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type DeleteAlarmsRequestPayload struct {
	AlarmIds []*int64 `json:"AlarmIds,omitempty" xml:"AlarmIds,omitempty" type:"Repeated"`
}

func (s DeleteAlarmsRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsRequestPayload) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequestPayload) SetAlarmIds(v []*int64) *DeleteAlarmsRequestPayload {
	s.AlarmIds = v
	return s
}

type DeleteAlarmsRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteAlarmsRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequestUserInfo) SetEncodeKey(v string) *DeleteAlarmsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetEncodeType(v string) *DeleteAlarmsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetId(v string) *DeleteAlarmsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetIdType(v string) *DeleteAlarmsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetOrganizationId(v string) *DeleteAlarmsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type DeleteAlarmsShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeleteAlarmsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsShrinkRequest) SetDeviceInfoShrink(v string) *DeleteAlarmsShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *DeleteAlarmsShrinkRequest) SetPayloadShrink(v string) *DeleteAlarmsShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeleteAlarmsShrinkRequest) SetUserInfoShrink(v string) *DeleteAlarmsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type DeleteAlarmsResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsResponseBody) SetCode(v int32) *DeleteAlarmsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAlarmsResponseBody) SetMessage(v string) *DeleteAlarmsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAlarmsResponseBody) SetRequestId(v string) *DeleteAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlarmsResponseBody) SetResult(v bool) *DeleteAlarmsResponseBody {
	s.Result = &v
	return s
}

type DeleteAlarmsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsResponse) SetHeaders(v map[string]*string) *DeleteAlarmsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlarmsResponse) SetStatusCode(v int32) *DeleteAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlarmsResponse) SetBody(v *DeleteAlarmsResponseBody) *DeleteAlarmsResponse {
	s.Body = v
	return s
}

type DeleteScheduleTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteScheduleTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteScheduleTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteScheduleTaskHeaders) SetXAcsAligenieAccessToken(v string) *DeleteScheduleTaskHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteScheduleTaskHeaders) SetAuthorization(v string) *DeleteScheduleTaskHeaders {
	s.Authorization = &v
	return s
}

type DeleteScheduleTaskRequest struct {
	DeviceInfo *DeleteScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *DeleteScheduleTaskRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *DeleteScheduleTaskRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeleteScheduleTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequest) SetDeviceInfo(v *DeleteScheduleTaskRequestDeviceInfo) *DeleteScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeleteScheduleTaskRequest) SetPayload(v *DeleteScheduleTaskRequestPayload) *DeleteScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *DeleteScheduleTaskRequest) SetUserInfo(v *DeleteScheduleTaskRequestUserInfo) *DeleteScheduleTaskRequest {
	s.UserInfo = v
	return s
}

type DeleteScheduleTaskRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteScheduleTaskRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetId(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetIdType(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type DeleteScheduleTaskRequestPayload struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScheduleTaskRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestPayload) SetId(v int64) *DeleteScheduleTaskRequestPayload {
	s.Id = &v
	return s
}

type DeleteScheduleTaskRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteScheduleTaskRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestUserInfo) SetEncodeKey(v string) *DeleteScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetEncodeType(v string) *DeleteScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetId(v string) *DeleteScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetIdType(v string) *DeleteScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetOrganizationId(v string) *DeleteScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type DeleteScheduleTaskShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeleteScheduleTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskShrinkRequest) SetDeviceInfoShrink(v string) *DeleteScheduleTaskShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *DeleteScheduleTaskShrinkRequest) SetPayloadShrink(v string) *DeleteScheduleTaskShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeleteScheduleTaskShrinkRequest) SetUserInfoShrink(v string) *DeleteScheduleTaskShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type DeleteScheduleTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteScheduleTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskResponseBody) SetCode(v string) *DeleteScheduleTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScheduleTaskResponseBody) SetMessage(v string) *DeleteScheduleTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScheduleTaskResponseBody) SetRequestId(v string) *DeleteScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduleTaskResponseBody) SetResult(v bool) *DeleteScheduleTaskResponseBody {
	s.Result = &v
	return s
}

type DeleteScheduleTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduleTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskResponse) SetHeaders(v map[string]*string) *DeleteScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleTaskResponse) SetStatusCode(v int32) *DeleteScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduleTaskResponse) SetBody(v *DeleteScheduleTaskResponseBody) *DeleteScheduleTaskResponse {
	s.Body = v
	return s
}

type DeleteSubHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteSubHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSubHeaders) SetCommonHeaders(v map[string]*string) *DeleteSubHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSubHeaders) SetXAcsAligenieAccessToken(v string) *DeleteSubHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteSubHeaders) SetAuthorization(v string) *DeleteSubHeaders {
	s.Authorization = &v
	return s
}

type DeleteSubRequest struct {
	SubId *int64 `json:"SubId,omitempty" xml:"SubId,omitempty"`
}

func (s DeleteSubRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubRequest) SetSubId(v int64) *DeleteSubRequest {
	s.SubId = &v
	return s
}

type DeleteSubResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteSubResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubResponseBody) SetCode(v int32) *DeleteSubResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSubResponseBody) SetMessage(v string) *DeleteSubResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSubResponseBody) SetRequestId(v string) *DeleteSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubResponseBody) SetResult(v bool) *DeleteSubResponseBody {
	s.Result = &v
	return s
}

type DeleteSubResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubResponse) SetHeaders(v map[string]*string) *DeleteSubResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubResponse) SetStatusCode(v int32) *DeleteSubResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubResponse) SetBody(v *DeleteSubResponseBody) *DeleteSubResponse {
	s.Body = v
	return s
}

type DeviceControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeviceControlHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlHeaders) GoString() string {
	return s.String()
}

func (s *DeviceControlHeaders) SetCommonHeaders(v map[string]*string) *DeviceControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeviceControlHeaders) SetXAcsAligenieAccessToken(v string) *DeviceControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeviceControlHeaders) SetAuthorization(v string) *DeviceControlHeaders {
	s.Authorization = &v
	return s
}

type DeviceControlRequest struct {
	ControlRequest *DeviceControlRequestControlRequest `json:"ControlRequest,omitempty" xml:"ControlRequest,omitempty" type:"Struct"`
	DeviceInfo     *DeviceControlRequestDeviceInfo     `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s DeviceControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlRequest) SetControlRequest(v *DeviceControlRequestControlRequest) *DeviceControlRequest {
	s.ControlRequest = v
	return s
}

func (s *DeviceControlRequest) SetDeviceInfo(v *DeviceControlRequestDeviceInfo) *DeviceControlRequest {
	s.DeviceInfo = v
	return s
}

type DeviceControlRequestControlRequest struct {
	Muted  *bool  `json:"Muted,omitempty" xml:"Muted,omitempty"`
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DeviceControlRequestControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequestControlRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestControlRequest) SetMuted(v bool) *DeviceControlRequestControlRequest {
	s.Muted = &v
	return s
}

func (s *DeviceControlRequestControlRequest) SetVolume(v int32) *DeviceControlRequestControlRequest {
	s.Volume = &v
	return s
}

type DeviceControlRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeviceControlRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestDeviceInfo) SetEncodeKey(v string) *DeviceControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetEncodeType(v string) *DeviceControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetId(v string) *DeviceControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetIdType(v string) *DeviceControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetOrganizationId(v string) *DeviceControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type DeviceControlShrinkRequest struct {
	ControlRequestShrink *string `json:"ControlRequest,omitempty" xml:"ControlRequest,omitempty"`
	DeviceInfoShrink     *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s DeviceControlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlShrinkRequest) SetControlRequestShrink(v string) *DeviceControlShrinkRequest {
	s.ControlRequestShrink = &v
	return s
}

func (s *DeviceControlShrinkRequest) SetDeviceInfoShrink(v string) *DeviceControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

type DeviceControlResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeviceControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBody) SetCode(v int32) *DeviceControlResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceControlResponseBody) SetMessage(v string) *DeviceControlResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceControlResponseBody) SetRequestId(v string) *DeviceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceControlResponseBody) SetResult(v bool) *DeviceControlResponseBody {
	s.Result = &v
	return s
}

type DeviceControlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeviceControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeviceControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponse) GoString() string {
	return s.String()
}

func (s *DeviceControlResponse) SetHeaders(v map[string]*string) *DeviceControlResponse {
	s.Headers = v
	return s
}

func (s *DeviceControlResponse) SetStatusCode(v int32) *DeviceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceControlResponse) SetBody(v *DeviceControlResponseBody) *DeviceControlResponse {
	s.Body = v
	return s
}

type EcologyOpennessAuthenticateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s EcologyOpennessAuthenticateHeaders) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessAuthenticateHeaders) GoString() string {
	return s.String()
}

func (s *EcologyOpennessAuthenticateHeaders) SetCommonHeaders(v map[string]*string) *EcologyOpennessAuthenticateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EcologyOpennessAuthenticateHeaders) SetXAcsAligenieAccessToken(v string) *EcologyOpennessAuthenticateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *EcologyOpennessAuthenticateHeaders) SetAuthorization(v string) *EcologyOpennessAuthenticateHeaders {
	s.Authorization = &v
	return s
}

type EcologyOpennessAuthenticateRequest struct {
	EncodeKey             *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType            *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s EcologyOpennessAuthenticateRequest) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessAuthenticateRequest) GoString() string {
	return s.String()
}

func (s *EcologyOpennessAuthenticateRequest) SetEncodeKey(v string) *EcologyOpennessAuthenticateRequest {
	s.EncodeKey = &v
	return s
}

func (s *EcologyOpennessAuthenticateRequest) SetEncodeType(v string) *EcologyOpennessAuthenticateRequest {
	s.EncodeType = &v
	return s
}

func (s *EcologyOpennessAuthenticateRequest) SetLoginStateAccessToken(v string) *EcologyOpennessAuthenticateRequest {
	s.LoginStateAccessToken = &v
	return s
}

type EcologyOpennessAuthenticateResponseBody struct {
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EcologyOpennessAuthenticateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EcologyOpennessAuthenticateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessAuthenticateResponseBody) GoString() string {
	return s.String()
}

func (s *EcologyOpennessAuthenticateResponseBody) SetCode(v int32) *EcologyOpennessAuthenticateResponseBody {
	s.Code = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetMessage(v string) *EcologyOpennessAuthenticateResponseBody {
	s.Message = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetRequestId(v string) *EcologyOpennessAuthenticateResponseBody {
	s.RequestId = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetResult(v *EcologyOpennessAuthenticateResponseBodyResult) *EcologyOpennessAuthenticateResponseBody {
	s.Result = v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetSuccess(v bool) *EcologyOpennessAuthenticateResponseBody {
	s.Success = &v
	return s
}

type EcologyOpennessAuthenticateResponseBodyResult struct {
	EncodeKey           *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType          *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	SceneCode           *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	ThirdUserIdentifier *string `json:"ThirdUserIdentifier,omitempty" xml:"ThirdUserIdentifier,omitempty"`
	ThirdUserType       *string `json:"ThirdUserType,omitempty" xml:"ThirdUserType,omitempty"`
	UserOpenId          *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s EcologyOpennessAuthenticateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessAuthenticateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetEncodeKey(v string) *EcologyOpennessAuthenticateResponseBodyResult {
	s.EncodeKey = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetEncodeType(v string) *EcologyOpennessAuthenticateResponseBodyResult {
	s.EncodeType = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetSceneCode(v string) *EcologyOpennessAuthenticateResponseBodyResult {
	s.SceneCode = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetThirdUserIdentifier(v string) *EcologyOpennessAuthenticateResponseBodyResult {
	s.ThirdUserIdentifier = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetThirdUserType(v string) *EcologyOpennessAuthenticateResponseBodyResult {
	s.ThirdUserType = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetUserOpenId(v string) *EcologyOpennessAuthenticateResponseBodyResult {
	s.UserOpenId = &v
	return s
}

type EcologyOpennessAuthenticateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EcologyOpennessAuthenticateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EcologyOpennessAuthenticateResponse) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessAuthenticateResponse) GoString() string {
	return s.String()
}

func (s *EcologyOpennessAuthenticateResponse) SetHeaders(v map[string]*string) *EcologyOpennessAuthenticateResponse {
	s.Headers = v
	return s
}

func (s *EcologyOpennessAuthenticateResponse) SetStatusCode(v int32) *EcologyOpennessAuthenticateResponse {
	s.StatusCode = &v
	return s
}

func (s *EcologyOpennessAuthenticateResponse) SetBody(v *EcologyOpennessAuthenticateResponseBody) *EcologyOpennessAuthenticateResponse {
	s.Body = v
	return s
}

type EcologyOpennessSendVerificationCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeHeaders) GoString() string {
	return s.String()
}

func (s *EcologyOpennessSendVerificationCodeHeaders) SetCommonHeaders(v map[string]*string) *EcologyOpennessSendVerificationCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EcologyOpennessSendVerificationCodeHeaders) SetXAcsAligenieAccessToken(v string) *EcologyOpennessSendVerificationCodeHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeHeaders) SetAuthorization(v string) *EcologyOpennessSendVerificationCodeHeaders {
	s.Authorization = &v
	return s
}

type EcologyOpennessSendVerificationCodeRequest struct {
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeRequest) GoString() string {
	return s.String()
}

func (s *EcologyOpennessSendVerificationCodeRequest) SetPhoneNumber(v string) *EcologyOpennessSendVerificationCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeRequest) SetRegion(v string) *EcologyOpennessSendVerificationCodeRequest {
	s.Region = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeRequest) SetSessionId(v string) *EcologyOpennessSendVerificationCodeRequest {
	s.SessionId = &v
	return s
}

type EcologyOpennessSendVerificationCodeResponseBody struct {
	Code      *int32                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EcologyOpennessSendVerificationCodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetCode(v int32) *EcologyOpennessSendVerificationCodeResponseBody {
	s.Code = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetMessage(v string) *EcologyOpennessSendVerificationCodeResponseBody {
	s.Message = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetRequestId(v string) *EcologyOpennessSendVerificationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetResult(v *EcologyOpennessSendVerificationCodeResponseBodyResult) *EcologyOpennessSendVerificationCodeResponseBody {
	s.Result = v
	return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetSuccess(v bool) *EcologyOpennessSendVerificationCodeResponseBody {
	s.Success = &v
	return s
}

type EcologyOpennessSendVerificationCodeResponseBodyResult struct {
	ExpireIn       *int32 `json:"ExpireIn,omitempty" xml:"ExpireIn,omitempty"`
	RepeatInterval *int32 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EcologyOpennessSendVerificationCodeResponseBodyResult) SetExpireIn(v int32) *EcologyOpennessSendVerificationCodeResponseBodyResult {
	s.ExpireIn = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBodyResult) SetRepeatInterval(v int32) *EcologyOpennessSendVerificationCodeResponseBodyResult {
	s.RepeatInterval = &v
	return s
}

type EcologyOpennessSendVerificationCodeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EcologyOpennessSendVerificationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EcologyOpennessSendVerificationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeResponse) GoString() string {
	return s.String()
}

func (s *EcologyOpennessSendVerificationCodeResponse) SetHeaders(v map[string]*string) *EcologyOpennessSendVerificationCodeResponse {
	s.Headers = v
	return s
}

func (s *EcologyOpennessSendVerificationCodeResponse) SetStatusCode(v int32) *EcologyOpennessSendVerificationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *EcologyOpennessSendVerificationCodeResponse) SetBody(v *EcologyOpennessSendVerificationCodeResponseBody) *EcologyOpennessSendVerificationCodeResponse {
	s.Body = v
	return s
}

type FindUserlistToAuthLoginWithPhoneNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *FindUserlistToAuthLoginWithPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) SetXAcsAligenieAccessToken(v string) *FindUserlistToAuthLoginWithPhoneNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) SetAuthorization(v string) *FindUserlistToAuthLoginWithPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

type FindUserlistToAuthLoginWithPhoneNumberRequest struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetCode(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.Code = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetPhoneNumber(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetRegion(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.Region = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetSessionId(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.SessionId = &v
	return s
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBody struct {
	Code      *int32                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DataObj   *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj `json:"DataObj,omitempty" xml:"DataObj,omitempty" type:"Struct"`
	Message   *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult  `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetCode(v int32) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetDataObj(v *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.DataObj = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetMessage(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetRequestId(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetResult(v *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Result = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetSuccess(v bool) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Success = &v
	return s
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj struct {
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) String() string {
	return tea.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) SetSessionId(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj {
	s.SessionId = &v
	return s
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult struct {
	UserListToAuthLogin []*FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin `json:"UserListToAuthLogin,omitempty" xml:"UserListToAuthLogin,omitempty" type:"Repeated"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) SetUserListToAuthLogin(v []*FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult {
	s.UserListToAuthLogin = v
	return s
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin struct {
	Avatar                  *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	EncryptedUserIdentifier *string `json:"EncryptedUserIdentifier,omitempty" xml:"EncryptedUserIdentifier,omitempty"`
	FindingType             *string `json:"FindingType,omitempty" xml:"FindingType,omitempty"`
	Nickname                *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	UserType                *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) String() string {
	return tea.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetAvatar(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.Avatar = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetEncryptedUserIdentifier(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.EncryptedUserIdentifier = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetFindingType(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.FindingType = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetNickname(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.Nickname = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetUserType(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.UserType = &v
	return s
}

type FindUserlistToAuthLoginWithPhoneNumberResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindUserlistToAuthLoginWithPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) SetHeaders(v map[string]*string) *FindUserlistToAuthLoginWithPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) SetStatusCode(v int32) *FindUserlistToAuthLoginWithPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) SetBody(v *FindUserlistToAuthLoginWithPhoneNumberResponseBody) *FindUserlistToAuthLoginWithPhoneNumberResponse {
	s.Body = v
	return s
}

type GetAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAlarmHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmHeaders) GoString() string {
	return s.String()
}

func (s *GetAlarmHeaders) SetCommonHeaders(v map[string]*string) *GetAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlarmHeaders) SetXAcsAligenieAccessToken(v string) *GetAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAlarmHeaders) SetAuthorization(v string) *GetAlarmHeaders {
	s.Authorization = &v
	return s
}

type GetAlarmRequest struct {
	DeviceInfo *GetAlarmRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetAlarmRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *GetAlarmRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmRequest) SetDeviceInfo(v *GetAlarmRequestDeviceInfo) *GetAlarmRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetAlarmRequest) SetPayload(v *GetAlarmRequestPayload) *GetAlarmRequest {
	s.Payload = v
	return s
}

func (s *GetAlarmRequest) SetUserInfo(v *GetAlarmRequestUserInfo) *GetAlarmRequest {
	s.UserInfo = v
	return s
}

type GetAlarmRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetAlarmRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestDeviceInfo) SetEncodeKey(v string) *GetAlarmRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetEncodeType(v string) *GetAlarmRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetId(v string) *GetAlarmRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetIdType(v string) *GetAlarmRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetOrganizationId(v string) *GetAlarmRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetAlarmRequestPayload struct {
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
}

func (s GetAlarmRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmRequestPayload) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestPayload) SetAlarmId(v int64) *GetAlarmRequestPayload {
	s.AlarmId = &v
	return s
}

type GetAlarmRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetAlarmRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestUserInfo) SetEncodeKey(v string) *GetAlarmRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetEncodeType(v string) *GetAlarmRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetId(v string) *GetAlarmRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetIdType(v string) *GetAlarmRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetOrganizationId(v string) *GetAlarmRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetAlarmShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetAlarmShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmShrinkRequest) SetDeviceInfoShrink(v string) *GetAlarmShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetAlarmShrinkRequest) SetPayloadShrink(v string) *GetAlarmShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetAlarmShrinkRequest) SetUserInfoShrink(v string) *GetAlarmShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetAlarmResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAlarmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBody) SetCode(v int32) *GetAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlarmResponseBody) SetMessage(v string) *GetAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlarmResponseBody) SetRequestId(v string) *GetAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlarmResponseBody) SetResult(v *GetAlarmResponseBodyResult) *GetAlarmResponseBody {
	s.Result = v
	return s
}

type GetAlarmResponseBodyResult struct {
	AlarmId          *int64                                  `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	MusicInfo        *GetAlarmResponseBodyResultMusicInfo    `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	ScheduleInfo     *GetAlarmResponseBodyResultScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	ScheduleTypeDesc *string                                 `json:"ScheduleTypeDesc,omitempty" xml:"ScheduleTypeDesc,omitempty"`
	Status           *int32                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TriggerDateDesc  *string                                 `json:"TriggerDateDesc,omitempty" xml:"TriggerDateDesc,omitempty"`
	TriggerTimeDesc  *string                                 `json:"TriggerTimeDesc,omitempty" xml:"TriggerTimeDesc,omitempty"`
	Volume           *int32                                  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetAlarmResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResult) SetAlarmId(v int64) *GetAlarmResponseBodyResult {
	s.AlarmId = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetMusicInfo(v *GetAlarmResponseBodyResultMusicInfo) *GetAlarmResponseBodyResult {
	s.MusicInfo = v
	return s
}

func (s *GetAlarmResponseBodyResult) SetScheduleInfo(v *GetAlarmResponseBodyResultScheduleInfo) *GetAlarmResponseBodyResult {
	s.ScheduleInfo = v
	return s
}

func (s *GetAlarmResponseBodyResult) SetScheduleTypeDesc(v string) *GetAlarmResponseBodyResult {
	s.ScheduleTypeDesc = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetStatus(v int32) *GetAlarmResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetTriggerDateDesc(v string) *GetAlarmResponseBodyResult {
	s.TriggerDateDesc = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetTriggerTimeDesc(v string) *GetAlarmResponseBodyResult {
	s.TriggerTimeDesc = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetVolume(v int32) *GetAlarmResponseBodyResult {
	s.Volume = &v
	return s
}

type GetAlarmResponseBodyResultMusicInfo struct {
	MusicId       *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName     *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	MusicUrl      *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s GetAlarmResponseBodyResultMusicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponseBodyResultMusicInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicId(v int64) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicId = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicName(v string) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicName = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicType(v int64) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicType = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicTypeName(v string) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicUrl(v string) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicUrl = &v
	return s
}

type GetAlarmResponseBodyResultScheduleInfo struct {
	Once                *GetAlarmResponseBodyResultScheduleInfoOnce                `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	StatutoryWorkingDay *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	Type                *string                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly              *GetAlarmResponseBodyResultScheduleInfoWeekly              `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s GetAlarmResponseBodyResultScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetOnce(v *GetAlarmResponseBodyResultScheduleInfoOnce) *GetAlarmResponseBodyResultScheduleInfo {
	s.Once = v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetStatutoryWorkingDay(v *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) *GetAlarmResponseBodyResultScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetType(v string) *GetAlarmResponseBodyResultScheduleInfo {
	s.Type = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetWeekly(v *GetAlarmResponseBodyResultScheduleInfoWeekly) *GetAlarmResponseBodyResultScheduleInfo {
	s.Weekly = v
	return s
}

type GetAlarmResponseBodyResultScheduleInfoOnce struct {
	Day    *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month  *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	Year   *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s GetAlarmResponseBodyResultScheduleInfoOnce) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetDay(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetHour(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetMinute(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetMonth(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetYear(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Year = &v
	return s
}

type GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay struct {
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) SetHour(v int32) *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

type GetAlarmResponseBodyResultScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s GetAlarmResponseBodyResultScheduleInfoWeekly) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *GetAlarmResponseBodyResultScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) SetHour(v int32) *GetAlarmResponseBodyResultScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) SetMinute(v int32) *GetAlarmResponseBodyResultScheduleInfoWeekly {
	s.Minute = &v
	return s
}

type GetAlarmResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmResponse) GoString() string {
	return s.String()
}

func (s *GetAlarmResponse) SetHeaders(v map[string]*string) *GetAlarmResponse {
	s.Headers = v
	return s
}

func (s *GetAlarmResponse) SetStatusCode(v int32) *GetAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlarmResponse) SetBody(v *GetAlarmResponseBody) *GetAlarmResponse {
	s.Body = v
	return s
}

type GetAlbumHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAlbumHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumHeaders) GoString() string {
	return s.String()
}

func (s *GetAlbumHeaders) SetCommonHeaders(v map[string]*string) *GetAlbumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlbumHeaders) SetXAcsAligenieAccessToken(v string) *GetAlbumHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAlbumHeaders) SetAuthorization(v string) *GetAlbumHeaders {
	s.Authorization = &v
	return s
}

type GetAlbumRequest struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlbumRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumRequest) GoString() string {
	return s.String()
}

func (s *GetAlbumRequest) SetId(v int64) *GetAlbumRequest {
	s.Id = &v
	return s
}

func (s *GetAlbumRequest) SetType(v string) *GetAlbumRequest {
	s.Type = &v
	return s
}

type GetAlbumResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAlbumResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAlbumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBody) SetCode(v int32) *GetAlbumResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlbumResponseBody) SetRequestId(v string) *GetAlbumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlbumResponseBody) SetResult(v *GetAlbumResponseBodyResult) *GetAlbumResponseBody {
	s.Result = v
	return s
}

type GetAlbumResponseBodyResult struct {
	Alias        []*string                            `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	Audition     *bool                                `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors      []*GetAlbumResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	Category     *string                              `json:"Category,omitempty" xml:"Category,omitempty"`
	Charge       *bool                                `json:"Charge,omitempty" xml:"Charge,omitempty"`
	CommCateId   *int64                               `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover        *GetAlbumResponseBodyResultCover     `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description  *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Finished     *string                              `json:"Finished,omitempty" xml:"Finished,omitempty"`
	HotScore     *float64                             `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	Id           *int64                               `json:"Id,omitempty" xml:"Id,omitempty"`
	ItemType     *string                              `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	RawId        *string                              `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source       *string                              `json:"Source,omitempty" xml:"Source,omitempty"`
	Title        *string                              `json:"Title,omitempty" xml:"Title,omitempty"`
	TotalEpisode *int32                               `json:"TotalEpisode,omitempty" xml:"TotalEpisode,omitempty"`
	Type         *string                              `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid        *string                              `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetAlbumResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBodyResult) SetAlias(v []*string) *GetAlbumResponseBodyResult {
	s.Alias = v
	return s
}

func (s *GetAlbumResponseBodyResult) SetAudition(v bool) *GetAlbumResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetAuthors(v []*GetAlbumResponseBodyResultAuthors) *GetAlbumResponseBodyResult {
	s.Authors = v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCategory(v string) *GetAlbumResponseBodyResult {
	s.Category = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCharge(v bool) *GetAlbumResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCommCateId(v int64) *GetAlbumResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCover(v *GetAlbumResponseBodyResultCover) *GetAlbumResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetAlbumResponseBodyResult) SetDescription(v string) *GetAlbumResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetFinished(v string) *GetAlbumResponseBodyResult {
	s.Finished = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetHotScore(v float64) *GetAlbumResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetId(v int64) *GetAlbumResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetItemType(v string) *GetAlbumResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetRawId(v string) *GetAlbumResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetSource(v string) *GetAlbumResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetTitle(v string) *GetAlbumResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetTotalEpisode(v int32) *GetAlbumResponseBodyResult {
	s.TotalEpisode = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetType(v string) *GetAlbumResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetValid(v string) *GetAlbumResponseBodyResult {
	s.Valid = &v
	return s
}

type GetAlbumResponseBodyResultAuthors struct {
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Gender      *string   `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Id          *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Online      *bool     `json:"Online,omitempty" xml:"Online,omitempty"`
	Source      *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetAlbumResponseBodyResultAuthors) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBodyResultAuthors) SetAuthorTypes(v []*string) *GetAlbumResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetGender(v string) *GetAlbumResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetId(v int64) *GetAlbumResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetOnline(v bool) *GetAlbumResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetSource(v string) *GetAlbumResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetTitle(v string) *GetAlbumResponseBodyResultAuthors {
	s.Title = &v
	return s
}

type GetAlbumResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s GetAlbumResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBodyResultCover) SetCanResize(v bool) *GetAlbumResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetImg(v string) *GetAlbumResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetLarge(v string) *GetAlbumResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetMedium(v string) *GetAlbumResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetSmall(v string) *GetAlbumResponseBodyResultCover {
	s.Small = &v
	return s
}

type GetAlbumResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlbumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlbumResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumResponse) GoString() string {
	return s.String()
}

func (s *GetAlbumResponse) SetHeaders(v map[string]*string) *GetAlbumResponse {
	s.Headers = v
	return s
}

func (s *GetAlbumResponse) SetStatusCode(v int32) *GetAlbumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlbumResponse) SetBody(v *GetAlbumResponseBody) *GetAlbumResponse {
	s.Body = v
	return s
}

type GetAlbumDetailByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAlbumDetailByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumDetailByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdHeaders) SetCommonHeaders(v map[string]*string) *GetAlbumDetailByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlbumDetailByIdHeaders) SetXAcsAligenieAccessToken(v string) *GetAlbumDetailByIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAlbumDetailByIdHeaders) SetAuthorization(v string) *GetAlbumDetailByIdHeaders {
	s.Authorization = &v
	return s
}

type GetAlbumDetailByIdRequest struct {
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
}

func (s GetAlbumDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdRequest) SetAlbumId(v string) *GetAlbumDetailByIdRequest {
	s.AlbumId = &v
	return s
}

type GetAlbumDetailByIdResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAlbumDetailByIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAlbumDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponseBody) SetCode(v int32) *GetAlbumDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBody) SetMessage(v string) *GetAlbumDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBody) SetRequestId(v string) *GetAlbumDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBody) SetResult(v *GetAlbumDetailByIdResponseBodyResult) *GetAlbumDetailByIdResponseBody {
	s.Result = v
	return s
}

type GetAlbumDetailByIdResponseBodyResult struct {
	AlbumContentList []*GetAlbumDetailByIdResponseBodyResultAlbumContentList `json:"AlbumContentList,omitempty" xml:"AlbumContentList,omitempty" type:"Repeated"`
	AlbumCoverUrl    *string                                                 `json:"AlbumCoverUrl,omitempty" xml:"AlbumCoverUrl,omitempty"`
	AlbumDescription *string                                                 `json:"AlbumDescription,omitempty" xml:"AlbumDescription,omitempty"`
	AlbumId          *string                                                 `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	AlbumTitle       *string                                                 `json:"AlbumTitle,omitempty" xml:"AlbumTitle,omitempty"`
}

func (s GetAlbumDetailByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumDetailByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumContentList(v []*GetAlbumDetailByIdResponseBodyResultAlbumContentList) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumContentList = v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumCoverUrl(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumCoverUrl = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumDescription(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumDescription = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumId(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumTitle(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumTitle = &v
	return s
}

type GetAlbumDetailByIdResponseBodyResultAlbumContentList struct {
	Duration   *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OrderIndex *string `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetAlbumDetailByIdResponseBodyResultAlbumContentList) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumDetailByIdResponseBodyResultAlbumContentList) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetDuration(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.Duration = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetId(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.Id = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetOrderIndex(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.OrderIndex = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetTitle(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.Title = &v
	return s
}

type GetAlbumDetailByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlbumDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlbumDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponse) SetHeaders(v map[string]*string) *GetAlbumDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAlbumDetailByIdResponse) SetStatusCode(v int32) *GetAlbumDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlbumDetailByIdResponse) SetBody(v *GetAlbumDetailByIdResponseBody) *GetAlbumDetailByIdResponse {
	s.Body = v
	return s
}

type GetAligenieUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAligenieUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAligenieUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAligenieUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAligenieUserInfoHeaders) SetXAcsAligenieAccessToken(v string) *GetAligenieUserInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAligenieUserInfoHeaders) SetAuthorization(v string) *GetAligenieUserInfoHeaders {
	s.Authorization = &v
	return s
}

type GetAligenieUserInfoRequest struct {
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s GetAligenieUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAligenieUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoRequest) SetLoginStateAccessToken(v string) *GetAligenieUserInfoRequest {
	s.LoginStateAccessToken = &v
	return s
}

type GetAligenieUserInfoResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAligenieUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAligenieUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAligenieUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoResponseBody) SetCode(v int32) *GetAligenieUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetMessage(v string) *GetAligenieUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetRequestId(v string) *GetAligenieUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetResult(v *GetAligenieUserInfoResponseBodyResult) *GetAligenieUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetSuccess(v bool) *GetAligenieUserInfoResponseBody {
	s.Success = &v
	return s
}

type GetAligenieUserInfoResponseBodyResult struct {
	AligenieNickname *string `json:"AligenieNickname,omitempty" xml:"AligenieNickname,omitempty"`
	Avatar           *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Deletable        *bool   `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
}

func (s GetAligenieUserInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAligenieUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoResponseBodyResult) SetAligenieNickname(v string) *GetAligenieUserInfoResponseBodyResult {
	s.AligenieNickname = &v
	return s
}

func (s *GetAligenieUserInfoResponseBodyResult) SetAvatar(v string) *GetAligenieUserInfoResponseBodyResult {
	s.Avatar = &v
	return s
}

func (s *GetAligenieUserInfoResponseBodyResult) SetDeletable(v bool) *GetAligenieUserInfoResponseBodyResult {
	s.Deletable = &v
	return s
}

type GetAligenieUserInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAligenieUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAligenieUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAligenieUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoResponse) SetHeaders(v map[string]*string) *GetAligenieUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAligenieUserInfoResponse) SetStatusCode(v int32) *GetAligenieUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAligenieUserInfoResponse) SetBody(v *GetAligenieUserInfoResponseBody) *GetAligenieUserInfoResponse {
	s.Body = v
	return s
}

type GetCodeEnhanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCodeEnhanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCodeEnhanceHeaders) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceHeaders) SetCommonHeaders(v map[string]*string) *GetCodeEnhanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCodeEnhanceHeaders) SetXAcsAligenieAccessToken(v string) *GetCodeEnhanceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCodeEnhanceHeaders) SetAuthorization(v string) *GetCodeEnhanceHeaders {
	s.Authorization = &v
	return s
}

type GetCodeEnhanceRequest struct {
	ChannelInfo *GetCodeEnhanceRequestChannelInfo `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty" type:"Struct"`
	UserInfo    *GetCodeEnhanceRequestUserInfo    `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetCodeEnhanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCodeEnhanceRequest) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceRequest) SetChannelInfo(v *GetCodeEnhanceRequestChannelInfo) *GetCodeEnhanceRequest {
	s.ChannelInfo = v
	return s
}

func (s *GetCodeEnhanceRequest) SetUserInfo(v *GetCodeEnhanceRequestUserInfo) *GetCodeEnhanceRequest {
	s.UserInfo = v
	return s
}

type GetCodeEnhanceRequestChannelInfo struct {
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s GetCodeEnhanceRequestChannelInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCodeEnhanceRequestChannelInfo) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceRequestChannelInfo) SetChannel(v string) *GetCodeEnhanceRequestChannelInfo {
	s.Channel = &v
	return s
}

func (s *GetCodeEnhanceRequestChannelInfo) SetExtInfo(v string) *GetCodeEnhanceRequestChannelInfo {
	s.ExtInfo = &v
	return s
}

type GetCodeEnhanceRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCodeEnhanceRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCodeEnhanceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceRequestUserInfo) SetEncodeKey(v string) *GetCodeEnhanceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetEncodeType(v string) *GetCodeEnhanceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetId(v string) *GetCodeEnhanceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetIdType(v string) *GetCodeEnhanceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetOrganizationId(v string) *GetCodeEnhanceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetCodeEnhanceShrinkRequest struct {
	ChannelInfoShrink *string `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty"`
	UserInfoShrink    *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetCodeEnhanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCodeEnhanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceShrinkRequest) SetChannelInfoShrink(v string) *GetCodeEnhanceShrinkRequest {
	s.ChannelInfoShrink = &v
	return s
}

func (s *GetCodeEnhanceShrinkRequest) SetUserInfoShrink(v string) *GetCodeEnhanceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetCodeEnhanceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetCodeEnhanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeEnhanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceResponseBody) SetCode(v int32) *GetCodeEnhanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetCodeEnhanceResponseBody) SetMessage(v string) *GetCodeEnhanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetCodeEnhanceResponseBody) SetRequestId(v string) *GetCodeEnhanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeEnhanceResponseBody) SetResult(v string) *GetCodeEnhanceResponseBody {
	s.Result = &v
	return s
}

type GetCodeEnhanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCodeEnhanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCodeEnhanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCodeEnhanceResponse) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceResponse) SetHeaders(v map[string]*string) *GetCodeEnhanceResponse {
	s.Headers = v
	return s
}

func (s *GetCodeEnhanceResponse) SetStatusCode(v int32) *GetCodeEnhanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeEnhanceResponse) SetBody(v *GetCodeEnhanceResponseBody) *GetCodeEnhanceResponse {
	s.Body = v
	return s
}

type GetContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetContentHeaders) GoString() string {
	return s.String()
}

func (s *GetContentHeaders) SetCommonHeaders(v map[string]*string) *GetContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetContentHeaders) SetXAcsAligenieAccessToken(v string) *GetContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetContentHeaders) SetAuthorization(v string) *GetContentHeaders {
	s.Authorization = &v
	return s
}

type GetContentRequest struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContentRequest) GoString() string {
	return s.String()
}

func (s *GetContentRequest) SetId(v int64) *GetContentRequest {
	s.Id = &v
	return s
}

func (s *GetContentRequest) SetType(v string) *GetContentRequest {
	s.Type = &v
	return s
}

type GetContentResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetContentResponseBody) SetCode(v int32) *GetContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetContentResponseBody) SetMessage(v string) *GetContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetContentResponseBody) SetRequestId(v string) *GetContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContentResponseBody) SetResult(v *GetContentResponseBodyResult) *GetContentResponseBody {
	s.Result = v
	return s
}

type GetContentResponseBodyResult struct {
	AlbumId     *string                                `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	Alias       []*string                              `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	Audition    *bool                                  `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors     []*GetContentResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	Category    *string                                `json:"Category,omitempty" xml:"Category,omitempty"`
	Charge      *bool                                  `json:"Charge,omitempty" xml:"Charge,omitempty"`
	CommCateId  *int64                                 `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *GetContentResponseBodyResultCover     `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *int64                                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HotScore    *float64                               `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	Id          *int64                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ItemType    *string                                `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	Lyric       *string                                `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
	Source      *string                                `json:"Source,omitempty" xml:"Source,omitempty"`
	Styles      []*string                              `json:"Styles,omitempty" xml:"Styles,omitempty" type:"Repeated"`
	Title       *string                                `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string                                `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid       *string                                `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetContentResponseBodyResult) SetAlbumId(v string) *GetContentResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *GetContentResponseBodyResult) SetAlias(v []*string) *GetContentResponseBodyResult {
	s.Alias = v
	return s
}

func (s *GetContentResponseBodyResult) SetAudition(v bool) *GetContentResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *GetContentResponseBodyResult) SetAuthors(v []*GetContentResponseBodyResultAuthors) *GetContentResponseBodyResult {
	s.Authors = v
	return s
}

func (s *GetContentResponseBodyResult) SetCategory(v string) *GetContentResponseBodyResult {
	s.Category = &v
	return s
}

func (s *GetContentResponseBodyResult) SetCharge(v bool) *GetContentResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *GetContentResponseBodyResult) SetCommCateId(v int64) *GetContentResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *GetContentResponseBodyResult) SetCover(v *GetContentResponseBodyResultCover) *GetContentResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetContentResponseBodyResult) SetDescription(v string) *GetContentResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetContentResponseBodyResult) SetDuration(v int64) *GetContentResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *GetContentResponseBodyResult) SetHotScore(v float64) *GetContentResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *GetContentResponseBodyResult) SetId(v int64) *GetContentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetContentResponseBodyResult) SetItemType(v string) *GetContentResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *GetContentResponseBodyResult) SetLyric(v string) *GetContentResponseBodyResult {
	s.Lyric = &v
	return s
}

func (s *GetContentResponseBodyResult) SetSource(v string) *GetContentResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetContentResponseBodyResult) SetStyles(v []*string) *GetContentResponseBodyResult {
	s.Styles = v
	return s
}

func (s *GetContentResponseBodyResult) SetTitle(v string) *GetContentResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetContentResponseBodyResult) SetType(v string) *GetContentResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetContentResponseBodyResult) SetValid(v string) *GetContentResponseBodyResult {
	s.Valid = &v
	return s
}

type GetContentResponseBodyResultAuthors struct {
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Gender      *string   `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Id          *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Online      *bool     `json:"Online,omitempty" xml:"Online,omitempty"`
	Source      *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetContentResponseBodyResultAuthors) String() string {
	return tea.Prettify(s)
}

func (s GetContentResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *GetContentResponseBodyResultAuthors) SetAuthorTypes(v []*string) *GetContentResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetGender(v string) *GetContentResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetId(v int64) *GetContentResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetOnline(v bool) *GetContentResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetSource(v string) *GetContentResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetTitle(v string) *GetContentResponseBodyResultAuthors {
	s.Title = &v
	return s
}

type GetContentResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s GetContentResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s GetContentResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetContentResponseBodyResultCover) SetCanResize(v bool) *GetContentResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetImg(v string) *GetContentResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetLarge(v string) *GetContentResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetMedium(v string) *GetContentResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetSmall(v string) *GetContentResponseBodyResultCover {
	s.Small = &v
	return s
}

type GetContentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContentResponse) GoString() string {
	return s.String()
}

func (s *GetContentResponse) SetHeaders(v map[string]*string) *GetContentResponse {
	s.Headers = v
	return s
}

func (s *GetContentResponse) SetStatusCode(v int32) *GetContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContentResponse) SetBody(v *GetContentResponseBody) *GetContentResponse {
	s.Body = v
	return s
}

type GetCurrentPlayingItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCurrentPlayingItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemHeaders) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemHeaders) SetCommonHeaders(v map[string]*string) *GetCurrentPlayingItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCurrentPlayingItemHeaders) SetXAcsAligenieAccessToken(v string) *GetCurrentPlayingItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCurrentPlayingItemHeaders) SetAuthorization(v string) *GetCurrentPlayingItemHeaders {
	s.Authorization = &v
	return s
}

type GetCurrentPlayingItemRequest struct {
	DeviceInfo *GetCurrentPlayingItemRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *GetCurrentPlayingItemRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetCurrentPlayingItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequest) SetDeviceInfo(v *GetCurrentPlayingItemRequestDeviceInfo) *GetCurrentPlayingItemRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetCurrentPlayingItemRequest) SetUserInfo(v *GetCurrentPlayingItemRequestUserInfo) *GetCurrentPlayingItemRequest {
	s.UserInfo = v
	return s
}

type GetCurrentPlayingItemRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingItemRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetEncodeKey(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetEncodeType(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetId(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetIdType(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetOrganizationId(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetCurrentPlayingItemRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingItemRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetEncodeKey(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetEncodeType(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetId(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetIdType(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetOrganizationId(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetCurrentPlayingItemShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetCurrentPlayingItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemShrinkRequest) SetDeviceInfoShrink(v string) *GetCurrentPlayingItemShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetCurrentPlayingItemShrinkRequest) SetUserInfoShrink(v string) *GetCurrentPlayingItemShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetCurrentPlayingItemResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetCurrentPlayingItemResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *string                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCurrentPlayingItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponseBody) SetCode(v int32) *GetCurrentPlayingItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetMessage(v string) *GetCurrentPlayingItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetRequestId(v string) *GetCurrentPlayingItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetResult(v *GetCurrentPlayingItemResponseBodyResult) *GetCurrentPlayingItemResponseBody {
	s.Result = v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetSuccess(v string) *GetCurrentPlayingItemResponseBody {
	s.Success = &v
	return s
}

type GetCurrentPlayingItemResponseBodyResult struct {
	AlbumName        *string                                       `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	AlbumRawId       *string                                       `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	AudioLength      *int32                                        `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	Copyright        *int32                                        `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover            *GetCurrentPlayingItemResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	DefaultPlayOrder *int32                                        `json:"DefaultPlayOrder,omitempty" xml:"DefaultPlayOrder,omitempty"`
	ItemUrl          *string                                       `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Liked            *string                                       `json:"Liked,omitempty" xml:"Liked,omitempty"`
	LyricUrl         *string                                       `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	PlayMode         *string                                       `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	Pos              *int32                                        `json:"Pos,omitempty" xml:"Pos,omitempty"`
	Progress         *int32                                        `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RawId            *string                                       `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singer           *string                                       `json:"Singer,omitempty" xml:"Singer,omitempty"`
	Source           *string                                       `json:"Source,omitempty" xml:"Source,omitempty"`
	Title            *string                                       `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid            *string                                       `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetCurrentPlayingItemResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetAlbumName(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetAlbumRawId(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetAudioLength(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetCopyright(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetCover(v *GetCurrentPlayingItemResponseBodyResultCover) *GetCurrentPlayingItemResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetDefaultPlayOrder(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetItemUrl(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetLiked(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetLyricUrl(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetPlayMode(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetPos(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetProgress(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetRawId(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetSinger(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetSource(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetTitle(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetType(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetValid(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Valid = &v
	return s
}

type GetCurrentPlayingItemResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s GetCurrentPlayingItemResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetCanResize(v bool) *GetCurrentPlayingItemResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetImg(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetLarge(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetMediam(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetMedium(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetSmall(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Small = &v
	return s
}

type GetCurrentPlayingItemResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCurrentPlayingItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCurrentPlayingItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingItemResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponse) SetHeaders(v map[string]*string) *GetCurrentPlayingItemResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentPlayingItemResponse) SetStatusCode(v int32) *GetCurrentPlayingItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCurrentPlayingItemResponse) SetBody(v *GetCurrentPlayingItemResponseBody) *GetCurrentPlayingItemResponse {
	s.Body = v
	return s
}

type GetCurrentPlayingListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCurrentPlayingListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListHeaders) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListHeaders) SetCommonHeaders(v map[string]*string) *GetCurrentPlayingListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCurrentPlayingListHeaders) SetXAcsAligenieAccessToken(v string) *GetCurrentPlayingListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCurrentPlayingListHeaders) SetAuthorization(v string) *GetCurrentPlayingListHeaders {
	s.Authorization = &v
	return s
}

type GetCurrentPlayingListRequest struct {
	DeviceInfo               *GetCurrentPlayingListRequestDeviceInfo               `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenQueryPlayListRequest *GetCurrentPlayingListRequestOpenQueryPlayListRequest `json:"OpenQueryPlayListRequest,omitempty" xml:"OpenQueryPlayListRequest,omitempty" type:"Struct"`
	UserInfo                 *GetCurrentPlayingListRequestUserInfo                 `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetCurrentPlayingListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequest) SetDeviceInfo(v *GetCurrentPlayingListRequestDeviceInfo) *GetCurrentPlayingListRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetCurrentPlayingListRequest) SetOpenQueryPlayListRequest(v *GetCurrentPlayingListRequestOpenQueryPlayListRequest) *GetCurrentPlayingListRequest {
	s.OpenQueryPlayListRequest = v
	return s
}

func (s *GetCurrentPlayingListRequest) SetUserInfo(v *GetCurrentPlayingListRequestUserInfo) *GetCurrentPlayingListRequest {
	s.UserInfo = v
	return s
}

type GetCurrentPlayingListRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingListRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetEncodeKey(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetEncodeType(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetId(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetIdType(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetOrganizationId(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetCurrentPlayingListRequestOpenQueryPlayListRequest struct {
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetCurrentPlayingListRequestOpenQueryPlayListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListRequestOpenQueryPlayListRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequestOpenQueryPlayListRequest) SetPageNum(v int32) *GetCurrentPlayingListRequestOpenQueryPlayListRequest {
	s.PageNum = &v
	return s
}

func (s *GetCurrentPlayingListRequestOpenQueryPlayListRequest) SetPageSize(v int32) *GetCurrentPlayingListRequestOpenQueryPlayListRequest {
	s.PageSize = &v
	return s
}

type GetCurrentPlayingListRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingListRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequestUserInfo) SetEncodeKey(v string) *GetCurrentPlayingListRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetEncodeType(v string) *GetCurrentPlayingListRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetId(v string) *GetCurrentPlayingListRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetIdType(v string) *GetCurrentPlayingListRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetOrganizationId(v string) *GetCurrentPlayingListRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetCurrentPlayingListShrinkRequest struct {
	DeviceInfoShrink               *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenQueryPlayListRequestShrink *string `json:"OpenQueryPlayListRequest,omitempty" xml:"OpenQueryPlayListRequest,omitempty"`
	UserInfoShrink                 *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetCurrentPlayingListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListShrinkRequest) SetDeviceInfoShrink(v string) *GetCurrentPlayingListShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetCurrentPlayingListShrinkRequest) SetOpenQueryPlayListRequestShrink(v string) *GetCurrentPlayingListShrinkRequest {
	s.OpenQueryPlayListRequestShrink = &v
	return s
}

func (s *GetCurrentPlayingListShrinkRequest) SetUserInfoShrink(v string) *GetCurrentPlayingListShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetCurrentPlayingListResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetCurrentPlayingListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *string                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCurrentPlayingListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponseBody) SetCode(v int32) *GetCurrentPlayingListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetMessage(v string) *GetCurrentPlayingListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetRequestId(v string) *GetCurrentPlayingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetResult(v []*GetCurrentPlayingListResponseBodyResult) *GetCurrentPlayingListResponseBody {
	s.Result = v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetSuccess(v string) *GetCurrentPlayingListResponseBody {
	s.Success = &v
	return s
}

type GetCurrentPlayingListResponseBodyResult struct {
	AlbumName        *string                                       `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	AlbumRawId       *string                                       `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	AudioLength      *int32                                        `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	Copyright        *int32                                        `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover            *GetCurrentPlayingListResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	DefaultPlayOrder *int32                                        `json:"DefaultPlayOrder,omitempty" xml:"DefaultPlayOrder,omitempty"`
	ItemUrl          *string                                       `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Liked            *bool                                         `json:"Liked,omitempty" xml:"Liked,omitempty"`
	LyricUrl         *string                                       `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	PlayMode         *string                                       `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	Pos              *int32                                        `json:"Pos,omitempty" xml:"Pos,omitempty"`
	Progress         *int32                                        `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RawId            *string                                       `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singer           *string                                       `json:"Singer,omitempty" xml:"Singer,omitempty"`
	Source           *string                                       `json:"Source,omitempty" xml:"Source,omitempty"`
	Title            *string                                       `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid            *string                                       `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetCurrentPlayingListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponseBodyResult) SetAlbumName(v string) *GetCurrentPlayingListResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetAlbumRawId(v string) *GetCurrentPlayingListResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetAudioLength(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetCopyright(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetCover(v *GetCurrentPlayingListResponseBodyResultCover) *GetCurrentPlayingListResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetDefaultPlayOrder(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetItemUrl(v string) *GetCurrentPlayingListResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetLiked(v bool) *GetCurrentPlayingListResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetLyricUrl(v string) *GetCurrentPlayingListResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetPlayMode(v string) *GetCurrentPlayingListResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetPos(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetProgress(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetRawId(v string) *GetCurrentPlayingListResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetSinger(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetSource(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetTitle(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetType(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetValid(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Valid = &v
	return s
}

type GetCurrentPlayingListResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s GetCurrentPlayingListResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetCanResize(v bool) *GetCurrentPlayingListResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetImg(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetLarge(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetMediam(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetMedium(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetSmall(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Small = &v
	return s
}

type GetCurrentPlayingListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCurrentPlayingListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCurrentPlayingListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentPlayingListResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponse) SetHeaders(v map[string]*string) *GetCurrentPlayingListResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentPlayingListResponse) SetStatusCode(v int32) *GetCurrentPlayingListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCurrentPlayingListResponse) SetBody(v *GetCurrentPlayingListResponseBody) *GetCurrentPlayingListResponse {
	s.Body = v
	return s
}

type GetDeviceBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceBasicInfoHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceBasicInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceBasicInfoHeaders) SetAuthorization(v string) *GetDeviceBasicInfoHeaders {
	s.Authorization = &v
	return s
}

type GetDeviceBasicInfoRequest struct {
	DeviceInfo *GetDeviceBasicInfoRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetDeviceBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoRequest) SetDeviceInfo(v *GetDeviceBasicInfoRequestDeviceInfo) *GetDeviceBasicInfoRequest {
	s.DeviceInfo = v
	return s
}

type GetDeviceBasicInfoRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceBasicInfoRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceBasicInfoRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetEncodeType(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetId(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetIdType(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetDeviceBasicInfoShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetDeviceBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceBasicInfoShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

type GetDeviceBasicInfoResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceBasicInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoResponseBody) SetCode(v int32) *GetDeviceBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBody) SetMessage(v string) *GetDeviceBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBody) SetRequestId(v string) *GetDeviceBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBody) SetResult(v *GetDeviceBasicInfoResponseBodyResult) *GetDeviceBasicInfoResponseBody {
	s.Result = v
	return s
}

type GetDeviceBasicInfoResponseBodyResult struct {
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	Mac             *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Sn              *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s GetDeviceBasicInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceBasicInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetFirmwareVersion(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.FirmwareVersion = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetMac(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.Mac = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetName(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetSn(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.Sn = &v
	return s
}

type GetDeviceBasicInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoResponse) SetHeaders(v map[string]*string) *GetDeviceBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceBasicInfoResponse) SetStatusCode(v int32) *GetDeviceBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceBasicInfoResponse) SetBody(v *GetDeviceBasicInfoResponseBody) *GetDeviceBasicInfoResponse {
	s.Body = v
	return s
}

type GetDeviceIdByIdentityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceIdByIdentityHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceIdByIdentityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceIdByIdentityHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceIdByIdentityHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceIdByIdentityHeaders) SetAuthorization(v string) *GetDeviceIdByIdentityHeaders {
	s.Authorization = &v
	return s
}

type GetDeviceIdByIdentityRequest struct {
	EncodeKey    *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType   *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	IdentityId   *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	ProductKey   *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s GetDeviceIdByIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityRequest) SetEncodeKey(v string) *GetDeviceIdByIdentityRequest {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetEncodeType(v string) *GetDeviceIdByIdentityRequest {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetIdentityId(v string) *GetDeviceIdByIdentityRequest {
	s.IdentityId = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetIdentityType(v string) *GetDeviceIdByIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetProductKey(v string) *GetDeviceIdByIdentityRequest {
	s.ProductKey = &v
	return s
}

type GetDeviceIdByIdentityResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceIdByIdentityResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceIdByIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBody) SetCode(v int32) *GetDeviceIdByIdentityResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetMessage(v string) *GetDeviceIdByIdentityResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetRequestId(v string) *GetDeviceIdByIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetResult(v *GetDeviceIdByIdentityResponseBodyResult) *GetDeviceIdByIdentityResponseBody {
	s.Result = v
	return s
}

type GetDeviceIdByIdentityResponseBodyResult struct {
	DeviceOpenId   *string                                                  `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds []*GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
}

func (s GetDeviceIdByIdentityResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBodyResult) SetDeviceOpenId(v string) *GetDeviceIdByIdentityResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResult) SetDeviceUnionIds(v []*GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) *GetDeviceIdByIdentityResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

type GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds struct {
	DeviceUnionId  *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) SetDeviceUnionId(v string) *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) SetOrganizationId(v string) *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

type GetDeviceIdByIdentityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceIdByIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceIdByIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponse) SetHeaders(v map[string]*string) *GetDeviceIdByIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceIdByIdentityResponse) SetStatusCode(v int32) *GetDeviceIdByIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceIdByIdentityResponse) SetBody(v *GetDeviceIdByIdentityResponseBody) *GetDeviceIdByIdentityResponse {
	s.Body = v
	return s
}

type GetDeviceSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceSettingHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceSettingHeaders) SetAuthorization(v string) *GetDeviceSettingHeaders {
	s.Authorization = &v
	return s
}

type GetDeviceSettingRequest struct {
	DeviceInfo *GetDeviceSettingRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Keys       []*string                          `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s GetDeviceSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSettingRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingRequest) SetDeviceInfo(v *GetDeviceSettingRequestDeviceInfo) *GetDeviceSettingRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetDeviceSettingRequest) SetKeys(v []*string) *GetDeviceSettingRequest {
	s.Keys = v
	return s
}

type GetDeviceSettingRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceSettingRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSettingRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceSettingRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetEncodeType(v string) *GetDeviceSettingRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetId(v string) *GetDeviceSettingRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetIdType(v string) *GetDeviceSettingRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceSettingRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetDeviceSettingShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	KeysShrink       *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s GetDeviceSettingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceSettingShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetDeviceSettingShrinkRequest) SetKeysShrink(v string) *GetDeviceSettingShrinkRequest {
	s.KeysShrink = &v
	return s
}

type GetDeviceSettingResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetDeviceSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingResponseBody) SetCode(v int32) *GetDeviceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceSettingResponseBody) SetMessage(v string) *GetDeviceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceSettingResponseBody) SetRequestId(v string) *GetDeviceSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceSettingResponseBody) SetResult(v map[string]interface{}) *GetDeviceSettingResponseBody {
	s.Result = v
	return s
}

type GetDeviceSettingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSettingResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingResponse) SetHeaders(v map[string]*string) *GetDeviceSettingResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceSettingResponse) SetStatusCode(v int32) *GetDeviceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceSettingResponse) SetBody(v *GetDeviceSettingResponseBody) *GetDeviceSettingResponse {
	s.Body = v
	return s
}

type GetDeviceStatusDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceStatusDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceStatusDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceStatusDetailHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceStatusDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceStatusDetailHeaders) SetAuthorization(v string) *GetDeviceStatusDetailHeaders {
	s.Authorization = &v
	return s
}

type GetDeviceStatusDetailRequest struct {
	DeviceInfo *GetDeviceStatusDetailRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Keys       []*string                               `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s GetDeviceStatusDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailRequest) SetDeviceInfo(v *GetDeviceStatusDetailRequestDeviceInfo) *GetDeviceStatusDetailRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetDeviceStatusDetailRequest) SetKeys(v []*string) *GetDeviceStatusDetailRequest {
	s.Keys = v
	return s
}

type GetDeviceStatusDetailRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceStatusDetailRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetEncodeType(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetId(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetIdType(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetDeviceStatusDetailShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	KeysShrink       *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s GetDeviceStatusDetailShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceStatusDetailShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetDeviceStatusDetailShrinkRequest) SetKeysShrink(v string) *GetDeviceStatusDetailShrinkRequest {
	s.KeysShrink = &v
	return s
}

type GetDeviceStatusDetailResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceStatusDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceStatusDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBody) SetCode(v int32) *GetDeviceStatusDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBody) SetMessage(v string) *GetDeviceStatusDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBody) SetRequestId(v string) *GetDeviceStatusDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBody) SetResult(v *GetDeviceStatusDetailResponseBodyResult) *GetDeviceStatusDetailResponseBody {
	s.Result = v
	return s
}

type GetDeviceStatusDetailResponseBodyResult struct {
	Player  *GetDeviceStatusDetailResponseBodyResultPlayer  `json:"Player,omitempty" xml:"Player,omitempty" type:"Struct"`
	Power   *GetDeviceStatusDetailResponseBodyResultPower   `json:"Power,omitempty" xml:"Power,omitempty" type:"Struct"`
	Speaker *GetDeviceStatusDetailResponseBodyResultSpeaker `json:"Speaker,omitempty" xml:"Speaker,omitempty" type:"Struct"`
}

func (s GetDeviceStatusDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResult) SetPlayer(v *GetDeviceStatusDetailResponseBodyResultPlayer) *GetDeviceStatusDetailResponseBodyResult {
	s.Player = v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResult) SetPower(v *GetDeviceStatusDetailResponseBodyResultPower) *GetDeviceStatusDetailResponseBodyResult {
	s.Power = v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResult) SetSpeaker(v *GetDeviceStatusDetailResponseBodyResultSpeaker) *GetDeviceStatusDetailResponseBodyResult {
	s.Speaker = v
	return s
}

type GetDeviceStatusDetailResponseBodyResultPlayer struct {
	AudioAlbum  *string `json:"AudioAlbum,omitempty" xml:"AudioAlbum,omitempty"`
	AudioAnchor *string `json:"AudioAnchor,omitempty" xml:"AudioAnchor,omitempty"`
	AudioExt    *string `json:"AudioExt,omitempty" xml:"AudioExt,omitempty"`
	AudioId     *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	AudioLength *string `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	AudioName   *string `json:"AudioName,omitempty" xml:"AudioName,omitempty"`
	AudioSource *string `json:"AudioSource,omitempty" xml:"AudioSource,omitempty"`
	AudioUrl    *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	Format      *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Progress    *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Timestamp   *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetDeviceStatusDetailResponseBodyResultPlayer) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResultPlayer) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioAlbum(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioAlbum = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioAnchor(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioAnchor = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioExt(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioExt = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioId(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioId = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioLength(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioLength = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioName(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioName = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioSource(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioSource = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioUrl(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioUrl = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetFormat(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Format = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetProgress(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Progress = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetSource(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Source = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetStatus(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Status = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetTimestamp(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Timestamp = &v
	return s
}

type GetDeviceStatusDetailResponseBodyResultPower struct {
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeviceStatusDetailResponseBodyResultPower) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResultPower) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResultPower) SetQuantity(v int32) *GetDeviceStatusDetailResponseBodyResultPower {
	s.Quantity = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPower) SetStatus(v string) *GetDeviceStatusDetailResponseBodyResultPower {
	s.Status = &v
	return s
}

type GetDeviceStatusDetailResponseBodyResultSpeaker struct {
	Muted  *bool  `json:"Muted,omitempty" xml:"Muted,omitempty"`
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetDeviceStatusDetailResponseBodyResultSpeaker) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResultSpeaker) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResultSpeaker) SetMuted(v bool) *GetDeviceStatusDetailResponseBodyResultSpeaker {
	s.Muted = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultSpeaker) SetVolume(v int32) *GetDeviceStatusDetailResponseBodyResultSpeaker {
	s.Volume = &v
	return s
}

type GetDeviceStatusDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceStatusDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceStatusDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponse) SetHeaders(v map[string]*string) *GetDeviceStatusDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceStatusDetailResponse) SetStatusCode(v int32) *GetDeviceStatusDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceStatusDetailResponse) SetBody(v *GetDeviceStatusDetailResponseBody) *GetDeviceStatusDetailResponse {
	s.Body = v
	return s
}

type GetDeviceStatusInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceStatusInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceStatusInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceStatusInfoHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceStatusInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceStatusInfoHeaders) SetAuthorization(v string) *GetDeviceStatusInfoHeaders {
	s.Authorization = &v
	return s
}

type GetDeviceStatusInfoRequest struct {
	DeviceInfo *GetDeviceStatusInfoRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetDeviceStatusInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoRequest) SetDeviceInfo(v *GetDeviceStatusInfoRequestDeviceInfo) *GetDeviceStatusInfoRequest {
	s.DeviceInfo = v
	return s
}

type GetDeviceStatusInfoRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceStatusInfoRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusInfoRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetEncodeType(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetId(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetIdType(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetDeviceStatusInfoShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetDeviceStatusInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceStatusInfoShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

type GetDeviceStatusInfoResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceStatusInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceStatusInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoResponseBody) SetCode(v int32) *GetDeviceStatusInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceStatusInfoResponseBody) SetMessage(v string) *GetDeviceStatusInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceStatusInfoResponseBody) SetRequestId(v string) *GetDeviceStatusInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceStatusInfoResponseBody) SetResult(v *GetDeviceStatusInfoResponseBodyResult) *GetDeviceStatusInfoResponseBody {
	s.Result = v
	return s
}

type GetDeviceStatusInfoResponseBodyResult struct {
	Online *int32 `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s GetDeviceStatusInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoResponseBodyResult) SetOnline(v int32) *GetDeviceStatusInfoResponseBodyResult {
	s.Online = &v
	return s
}

type GetDeviceStatusInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceStatusInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceStatusInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceStatusInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoResponse) SetHeaders(v map[string]*string) *GetDeviceStatusInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceStatusInfoResponse) SetStatusCode(v int32) *GetDeviceStatusInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceStatusInfoResponse) SetBody(v *GetDeviceStatusInfoResponseBody) *GetDeviceStatusInfoResponse {
	s.Body = v
	return s
}

type GetDeviceTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTagHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceTagHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceTagHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceTagHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceTagHeaders) SetAuthorization(v string) *GetDeviceTagHeaders {
	s.Authorization = &v
	return s
}

type GetDeviceTagRequest struct {
	DeviceInfo *GetDeviceTagRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetDeviceTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTagRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceTagRequest) SetDeviceInfo(v *GetDeviceTagRequestDeviceInfo) *GetDeviceTagRequest {
	s.DeviceInfo = v
	return s
}

type GetDeviceTagRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceTagRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTagRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceTagRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceTagRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetEncodeType(v string) *GetDeviceTagRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetId(v string) *GetDeviceTagRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetIdType(v string) *GetDeviceTagRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceTagRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetDeviceTagShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetDeviceTagShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceTagShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceTagShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

type GetDeviceTagResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceTagResponseBody) SetCode(v int32) *GetDeviceTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceTagResponseBody) SetMessage(v string) *GetDeviceTagResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceTagResponseBody) SetRequestId(v string) *GetDeviceTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceTagResponseBody) SetResult(v *GetDeviceTagResponseBodyResult) *GetDeviceTagResponseBody {
	s.Result = v
	return s
}

type GetDeviceTagResponseBodyResult struct {
	DeviceTags map[string]interface{} `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty"`
}

func (s GetDeviceTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceTagResponseBodyResult) SetDeviceTags(v map[string]interface{}) *GetDeviceTagResponseBodyResult {
	s.DeviceTags = v
	return s
}

type GetDeviceTagResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTagResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceTagResponse) SetHeaders(v map[string]*string) *GetDeviceTagResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceTagResponse) SetStatusCode(v int32) *GetDeviceTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceTagResponse) SetBody(v *GetDeviceTagResponseBody) *GetDeviceTagResponse {
	s.Body = v
	return s
}

type GetScheduleTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetScheduleTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskHeaders) SetCommonHeaders(v map[string]*string) *GetScheduleTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScheduleTaskHeaders) SetXAcsAligenieAccessToken(v string) *GetScheduleTaskHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetScheduleTaskHeaders) SetAuthorization(v string) *GetScheduleTaskHeaders {
	s.Authorization = &v
	return s
}

type GetScheduleTaskRequest struct {
	DeviceInfo *GetScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetScheduleTaskRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *GetScheduleTaskRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetScheduleTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequest) SetDeviceInfo(v *GetScheduleTaskRequestDeviceInfo) *GetScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetScheduleTaskRequest) SetPayload(v *GetScheduleTaskRequestPayload) *GetScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *GetScheduleTaskRequest) SetUserInfo(v *GetScheduleTaskRequestUserInfo) *GetScheduleTaskRequest {
	s.UserInfo = v
	return s
}

type GetScheduleTaskRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetScheduleTaskRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *GetScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *GetScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetId(v string) *GetScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetIdType(v string) *GetScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *GetScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetScheduleTaskRequestPayload struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetScheduleTaskRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestPayload) SetId(v int64) *GetScheduleTaskRequestPayload {
	s.Id = &v
	return s
}

type GetScheduleTaskRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetScheduleTaskRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestUserInfo) SetEncodeKey(v string) *GetScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetEncodeType(v string) *GetScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetId(v string) *GetScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetIdType(v string) *GetScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetOrganizationId(v string) *GetScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetScheduleTaskShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetScheduleTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskShrinkRequest) SetDeviceInfoShrink(v string) *GetScheduleTaskShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetScheduleTaskShrinkRequest) SetPayloadShrink(v string) *GetScheduleTaskShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetScheduleTaskShrinkRequest) SetUserInfoShrink(v string) *GetScheduleTaskShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetScheduleTaskResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetScheduleTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetScheduleTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponseBody) SetCode(v int32) *GetScheduleTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduleTaskResponseBody) SetMessage(v string) *GetScheduleTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduleTaskResponseBody) SetRequestId(v string) *GetScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduleTaskResponseBody) SetResult(v *GetScheduleTaskResponseBodyResult) *GetScheduleTaskResponseBody {
	s.Result = v
	return s
}

type GetScheduleTaskResponseBodyResult struct {
	ActionTopicList   []*GetScheduleTaskResponseBodyResultActionTopicList `json:"ActionTopicList,omitempty" xml:"ActionTopicList,omitempty" type:"Repeated"`
	Cron              *string                                             `json:"Cron,omitempty" xml:"Cron,omitempty"`
	ScheduleEndTime   *string                                             `json:"ScheduleEndTime,omitempty" xml:"ScheduleEndTime,omitempty"`
	ScheduleId        *int64                                              `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	ScheduleStartTime *string                                             `json:"ScheduleStartTime,omitempty" xml:"ScheduleStartTime,omitempty"`
	ScheduleType      *string                                             `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
}

func (s GetScheduleTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponseBodyResult) SetActionTopicList(v []*GetScheduleTaskResponseBodyResultActionTopicList) *GetScheduleTaskResponseBodyResult {
	s.ActionTopicList = v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetCron(v string) *GetScheduleTaskResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleEndTime(v string) *GetScheduleTaskResponseBodyResult {
	s.ScheduleEndTime = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleId(v int64) *GetScheduleTaskResponseBodyResult {
	s.ScheduleId = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleStartTime(v string) *GetScheduleTaskResponseBodyResult {
	s.ScheduleStartTime = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleType(v string) *GetScheduleTaskResponseBodyResult {
	s.ScheduleType = &v
	return s
}

type GetScheduleTaskResponseBodyResultActionTopicList struct {
	CustomAction map[string]interface{} `json:"CustomAction,omitempty" xml:"CustomAction,omitempty"`
}

func (s GetScheduleTaskResponseBodyResultActionTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskResponseBodyResultActionTopicList) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponseBodyResultActionTopicList) SetCustomAction(v map[string]interface{}) *GetScheduleTaskResponseBodyResultActionTopicList {
	s.CustomAction = v
	return s
}

type GetScheduleTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScheduleTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponse) SetHeaders(v map[string]*string) *GetScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleTaskResponse) SetStatusCode(v int32) *GetScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduleTaskResponse) SetBody(v *GetScheduleTaskResponseBody) *GetScheduleTaskResponse {
	s.Body = v
	return s
}

type GetUnreadMessageCountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUnreadMessageCountHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUnreadMessageCountHeaders) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountHeaders) SetCommonHeaders(v map[string]*string) *GetUnreadMessageCountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUnreadMessageCountHeaders) SetXAcsAligenieAccessToken(v string) *GetUnreadMessageCountHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetUnreadMessageCountHeaders) SetAuthorization(v string) *GetUnreadMessageCountHeaders {
	s.Authorization = &v
	return s
}

type GetUnreadMessageCountRequest struct {
	UserInfo *GetUnreadMessageCountRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetUnreadMessageCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnreadMessageCountRequest) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountRequest) SetUserInfo(v *GetUnreadMessageCountRequestUserInfo) *GetUnreadMessageCountRequest {
	s.UserInfo = v
	return s
}

type GetUnreadMessageCountRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetUnreadMessageCountRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUnreadMessageCountRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountRequestUserInfo) SetEncodeKey(v string) *GetUnreadMessageCountRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetEncodeType(v string) *GetUnreadMessageCountRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetId(v string) *GetUnreadMessageCountRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetIdType(v string) *GetUnreadMessageCountRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetOrganizationId(v string) *GetUnreadMessageCountRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetUnreadMessageCountShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetUnreadMessageCountShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnreadMessageCountShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountShrinkRequest) SetUserInfoShrink(v string) *GetUnreadMessageCountShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetUnreadMessageCountResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Result  *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetUnreadMessageCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUnreadMessageCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountResponseBody) SetCode(v string) *GetUnreadMessageCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetUnreadMessageCountResponseBody) SetMessage(v string) *GetUnreadMessageCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnreadMessageCountResponseBody) SetResult(v int32) *GetUnreadMessageCountResponseBody {
	s.Result = &v
	return s
}

type GetUnreadMessageCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUnreadMessageCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUnreadMessageCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUnreadMessageCountResponse) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountResponse) SetHeaders(v map[string]*string) *GetUnreadMessageCountResponse {
	s.Headers = v
	return s
}

func (s *GetUnreadMessageCountResponse) SetStatusCode(v int32) *GetUnreadMessageCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnreadMessageCountResponse) SetBody(v *GetUnreadMessageCountResponseBody) *GetUnreadMessageCountResponse {
	s.Body = v
	return s
}

type GetUserByDeviceIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserByDeviceIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserByDeviceIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserByDeviceIdHeaders) SetXAcsAligenieAccessToken(v string) *GetUserByDeviceIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetUserByDeviceIdHeaders) SetAuthorization(v string) *GetUserByDeviceIdHeaders {
	s.Authorization = &v
	return s
}

type GetUserByDeviceIdRequest struct {
	DeviceInfo *GetUserByDeviceIdRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetUserByDeviceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdRequest) SetDeviceInfo(v *GetUserByDeviceIdRequestDeviceInfo) *GetUserByDeviceIdRequest {
	s.DeviceInfo = v
	return s
}

type GetUserByDeviceIdRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetUserByDeviceIdRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetEncodeKey(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetEncodeType(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetId(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetIdType(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetOrganizationId(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetUserByDeviceIdShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetUserByDeviceIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdShrinkRequest) SetDeviceInfoShrink(v string) *GetUserByDeviceIdShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

type GetUserByDeviceIdResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetUserByDeviceIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetUserByDeviceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponseBody) SetCode(v int32) *GetUserByDeviceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserByDeviceIdResponseBody) SetMessage(v string) *GetUserByDeviceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserByDeviceIdResponseBody) SetRequestId(v string) *GetUserByDeviceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserByDeviceIdResponseBody) SetResult(v *GetUserByDeviceIdResponseBodyResult) *GetUserByDeviceIdResponseBody {
	s.Result = v
	return s
}

type GetUserByDeviceIdResponseBodyResult struct {
	UserOpenId   *string                                            `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
	UserUnionIds []*GetUserByDeviceIdResponseBodyResultUserUnionIds `json:"UserUnionIds,omitempty" xml:"UserUnionIds,omitempty" type:"Repeated"`
}

func (s GetUserByDeviceIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponseBodyResult) SetUserOpenId(v string) *GetUserByDeviceIdResponseBodyResult {
	s.UserOpenId = &v
	return s
}

func (s *GetUserByDeviceIdResponseBodyResult) SetUserUnionIds(v []*GetUserByDeviceIdResponseBodyResultUserUnionIds) *GetUserByDeviceIdResponseBodyResult {
	s.UserUnionIds = v
	return s
}

type GetUserByDeviceIdResponseBodyResultUserUnionIds struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	UserUnionId    *string `json:"UserUnionId,omitempty" xml:"UserUnionId,omitempty"`
}

func (s GetUserByDeviceIdResponseBodyResultUserUnionIds) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdResponseBodyResultUserUnionIds) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponseBodyResultUserUnionIds) SetOrganizationId(v string) *GetUserByDeviceIdResponseBodyResultUserUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *GetUserByDeviceIdResponseBodyResultUserUnionIds) SetUserUnionId(v string) *GetUserByDeviceIdResponseBodyResultUserUnionIds {
	s.UserUnionId = &v
	return s
}

type GetUserByDeviceIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserByDeviceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserByDeviceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserByDeviceIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponse) SetHeaders(v map[string]*string) *GetUserByDeviceIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserByDeviceIdResponse) SetStatusCode(v int32) *GetUserByDeviceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserByDeviceIdResponse) SetBody(v *GetUserByDeviceIdResponseBody) *GetUserByDeviceIdResponse {
	s.Body = v
	return s
}

type GetWeatherHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetWeatherHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherHeaders) GoString() string {
	return s.String()
}

func (s *GetWeatherHeaders) SetCommonHeaders(v map[string]*string) *GetWeatherHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWeatherHeaders) SetXAcsAligenieAccessToken(v string) *GetWeatherHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetWeatherHeaders) SetAuthorization(v string) *GetWeatherHeaders {
	s.Authorization = &v
	return s
}

type GetWeatherRequest struct {
	DeviceInfo *GetWeatherRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetWeatherRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *GetWeatherRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetWeatherRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherRequest) GoString() string {
	return s.String()
}

func (s *GetWeatherRequest) SetDeviceInfo(v *GetWeatherRequestDeviceInfo) *GetWeatherRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetWeatherRequest) SetPayload(v *GetWeatherRequestPayload) *GetWeatherRequest {
	s.Payload = v
	return s
}

func (s *GetWeatherRequest) SetUserInfo(v *GetWeatherRequestUserInfo) *GetWeatherRequest {
	s.UserInfo = v
	return s
}

type GetWeatherRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetWeatherRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestDeviceInfo) SetEncodeKey(v string) *GetWeatherRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetEncodeType(v string) *GetWeatherRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetId(v string) *GetWeatherRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetIdType(v string) *GetWeatherRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetOrganizationId(v string) *GetWeatherRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetWeatherRequestPayload struct {
}

func (s GetWeatherRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherRequestPayload) GoString() string {
	return s.String()
}

type GetWeatherRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetWeatherRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestUserInfo) SetEncodeKey(v string) *GetWeatherRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetEncodeType(v string) *GetWeatherRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetId(v string) *GetWeatherRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetIdType(v string) *GetWeatherRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetOrganizationId(v string) *GetWeatherRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetWeatherShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetWeatherShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetWeatherShrinkRequest) SetDeviceInfoShrink(v string) *GetWeatherShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetWeatherShrinkRequest) SetPayloadShrink(v string) *GetWeatherShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetWeatherShrinkRequest) SetUserInfoShrink(v string) *GetWeatherShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetWeatherResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetWeatherResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetWeatherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherResponseBody) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBody) SetCode(v int32) *GetWeatherResponseBody {
	s.Code = &v
	return s
}

func (s *GetWeatherResponseBody) SetMessage(v string) *GetWeatherResponseBody {
	s.Message = &v
	return s
}

func (s *GetWeatherResponseBody) SetRequestId(v string) *GetWeatherResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWeatherResponseBody) SetResult(v *GetWeatherResponseBodyResult) *GetWeatherResponseBody {
	s.Result = v
	return s
}

type GetWeatherResponseBodyResult struct {
	CurrentMeteorology *GetWeatherResponseBodyResultCurrentMeteorology `json:"CurrentMeteorology,omitempty" xml:"CurrentMeteorology,omitempty" type:"Struct"`
}

func (s GetWeatherResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResult) SetCurrentMeteorology(v *GetWeatherResponseBodyResultCurrentMeteorology) *GetWeatherResponseBodyResult {
	s.CurrentMeteorology = v
	return s
}

type GetWeatherResponseBodyResultCurrentMeteorology struct {
	Temperature *GetWeatherResponseBodyResultCurrentMeteorologyTemperature `json:"Temperature,omitempty" xml:"Temperature,omitempty" type:"Struct"`
	Weather     *GetWeatherResponseBodyResultCurrentMeteorologyWeather     `json:"Weather,omitempty" xml:"Weather,omitempty" type:"Struct"`
}

func (s GetWeatherResponseBodyResultCurrentMeteorology) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherResponseBodyResultCurrentMeteorology) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResultCurrentMeteorology) SetTemperature(v *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) *GetWeatherResponseBodyResultCurrentMeteorology {
	s.Temperature = v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorology) SetWeather(v *GetWeatherResponseBodyResultCurrentMeteorologyWeather) *GetWeatherResponseBodyResultCurrentMeteorology {
	s.Weather = v
	return s
}

type GetWeatherResponseBodyResultCurrentMeteorologyTemperature struct {
	Current     *string `json:"Current,omitempty" xml:"Current,omitempty"`
	CurrentDesc *string `json:"CurrentDesc,omitempty" xml:"CurrentDesc,omitempty"`
	High        *string `json:"High,omitempty" xml:"High,omitempty"`
	HighDesc    *string `json:"HighDesc,omitempty" xml:"HighDesc,omitempty"`
	Logical     *string `json:"Logical,omitempty" xml:"Logical,omitempty"`
	Low         *string `json:"Low,omitempty" xml:"Low,omitempty"`
	LowDesc     *string `json:"LowDesc,omitempty" xml:"LowDesc,omitempty"`
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyTemperature) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetCurrent(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.Current = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetCurrentDesc(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.CurrentDesc = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetHigh(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.High = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetHighDesc(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.HighDesc = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetLogical(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.Logical = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetLow(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.Low = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetLowDesc(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.LowDesc = &v
	return s
}

type GetWeatherResponseBodyResultCurrentMeteorologyWeather struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyWeather) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyWeather) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyWeather) SetCode(v string) *GetWeatherResponseBodyResultCurrentMeteorologyWeather {
	s.Code = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyWeather) SetName(v string) *GetWeatherResponseBodyResultCurrentMeteorologyWeather {
	s.Name = &v
	return s
}

type GetWeatherResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWeatherResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWeatherResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWeatherResponse) GoString() string {
	return s.String()
}

func (s *GetWeatherResponse) SetHeaders(v map[string]*string) *GetWeatherResponse {
	s.Headers = v
	return s
}

func (s *GetWeatherResponse) SetStatusCode(v int32) *GetWeatherResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWeatherResponse) SetBody(v *GetWeatherResponseBody) *GetWeatherResponse {
	s.Body = v
	return s
}

type IndexControlPlayingListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s IndexControlPlayingListHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListHeaders) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListHeaders) SetCommonHeaders(v map[string]*string) *IndexControlPlayingListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndexControlPlayingListHeaders) SetXAcsAligenieAccessToken(v string) *IndexControlPlayingListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *IndexControlPlayingListHeaders) SetAuthorization(v string) *IndexControlPlayingListHeaders {
	s.Authorization = &v
	return s
}

type IndexControlPlayingListRequest struct {
	DeviceInfo              *IndexControlPlayingListRequestDeviceInfo              `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenIndexControlRequest *IndexControlPlayingListRequestOpenIndexControlRequest `json:"OpenIndexControlRequest,omitempty" xml:"OpenIndexControlRequest,omitempty" type:"Struct"`
	UserInfo                *IndexControlPlayingListRequestUserInfo                `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s IndexControlPlayingListRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequest) SetDeviceInfo(v *IndexControlPlayingListRequestDeviceInfo) *IndexControlPlayingListRequest {
	s.DeviceInfo = v
	return s
}

func (s *IndexControlPlayingListRequest) SetOpenIndexControlRequest(v *IndexControlPlayingListRequestOpenIndexControlRequest) *IndexControlPlayingListRequest {
	s.OpenIndexControlRequest = v
	return s
}

func (s *IndexControlPlayingListRequest) SetUserInfo(v *IndexControlPlayingListRequestUserInfo) *IndexControlPlayingListRequest {
	s.UserInfo = v
	return s
}

type IndexControlPlayingListRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s IndexControlPlayingListRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetEncodeKey(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetEncodeType(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetId(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetIdType(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetOrganizationId(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type IndexControlPlayingListRequestOpenIndexControlRequest struct {
	ExtendInfo           map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	Index                *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	NeedContentContinued *bool                  `json:"NeedContentContinued,omitempty" xml:"NeedContentContinued,omitempty"`
}

func (s IndexControlPlayingListRequestOpenIndexControlRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListRequestOpenIndexControlRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetExtendInfo(v map[string]interface{}) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.ExtendInfo = v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetIndex(v int32) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.Index = &v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetNeedContentContinued(v bool) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.NeedContentContinued = &v
	return s
}

type IndexControlPlayingListRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s IndexControlPlayingListRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListRequestUserInfo) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestUserInfo) SetEncodeKey(v string) *IndexControlPlayingListRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetEncodeType(v string) *IndexControlPlayingListRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetId(v string) *IndexControlPlayingListRequestUserInfo {
	s.Id = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetIdType(v string) *IndexControlPlayingListRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetOrganizationId(v string) *IndexControlPlayingListRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type IndexControlPlayingListShrinkRequest struct {
	DeviceInfoShrink              *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenIndexControlRequestShrink *string `json:"OpenIndexControlRequest,omitempty" xml:"OpenIndexControlRequest,omitempty"`
	UserInfoShrink                *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s IndexControlPlayingListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListShrinkRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListShrinkRequest) SetDeviceInfoShrink(v string) *IndexControlPlayingListShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *IndexControlPlayingListShrinkRequest) SetOpenIndexControlRequestShrink(v string) *IndexControlPlayingListShrinkRequest {
	s.OpenIndexControlRequestShrink = &v
	return s
}

func (s *IndexControlPlayingListShrinkRequest) SetUserInfoShrink(v string) *IndexControlPlayingListShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type IndexControlPlayingListResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *IndexControlPlayingListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *string                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IndexControlPlayingListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponseBody) SetCode(v int32) *IndexControlPlayingListResponseBody {
	s.Code = &v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetMessage(v string) *IndexControlPlayingListResponseBody {
	s.Message = &v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetRequestId(v string) *IndexControlPlayingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetResult(v *IndexControlPlayingListResponseBodyResult) *IndexControlPlayingListResponseBody {
	s.Result = v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetSuccess(v string) *IndexControlPlayingListResponseBody {
	s.Success = &v
	return s
}

type IndexControlPlayingListResponseBodyResult struct {
	AlbumName        *string                                         `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	AlbumRawId       *string                                         `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	AudioLength      *int32                                          `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	Copyright        *int32                                          `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover            *IndexControlPlayingListResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	DefaultPlayOrder *int32                                          `json:"DefaultPlayOrder,omitempty" xml:"DefaultPlayOrder,omitempty"`
	ItemUrl          *string                                         `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Liked            *bool                                           `json:"Liked,omitempty" xml:"Liked,omitempty"`
	LyricUrl         *string                                         `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	PlayMode         *string                                         `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	Pos              *int32                                          `json:"Pos,omitempty" xml:"Pos,omitempty"`
	Progress         *int32                                          `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RawId            *string                                         `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singer           *string                                         `json:"Singer,omitempty" xml:"Singer,omitempty"`
	Source           *string                                         `json:"Source,omitempty" xml:"Source,omitempty"`
	Title            *string                                         `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid            *string                                         `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s IndexControlPlayingListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponseBodyResult) SetAlbumName(v string) *IndexControlPlayingListResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetAlbumRawId(v string) *IndexControlPlayingListResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetAudioLength(v int32) *IndexControlPlayingListResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetCopyright(v int32) *IndexControlPlayingListResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetCover(v *IndexControlPlayingListResponseBodyResultCover) *IndexControlPlayingListResponseBodyResult {
	s.Cover = v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetDefaultPlayOrder(v int32) *IndexControlPlayingListResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetItemUrl(v string) *IndexControlPlayingListResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetLiked(v bool) *IndexControlPlayingListResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetLyricUrl(v string) *IndexControlPlayingListResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetPlayMode(v string) *IndexControlPlayingListResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetPos(v int32) *IndexControlPlayingListResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetProgress(v int32) *IndexControlPlayingListResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetRawId(v string) *IndexControlPlayingListResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetSinger(v string) *IndexControlPlayingListResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetSource(v string) *IndexControlPlayingListResponseBodyResult {
	s.Source = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetTitle(v string) *IndexControlPlayingListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetType(v string) *IndexControlPlayingListResponseBodyResult {
	s.Type = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetValid(v string) *IndexControlPlayingListResponseBodyResult {
	s.Valid = &v
	return s
}

type IndexControlPlayingListResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s IndexControlPlayingListResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetCanResize(v bool) *IndexControlPlayingListResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetImg(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetLarge(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetMediam(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetMedium(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetSmall(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Small = &v
	return s
}

type IndexControlPlayingListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndexControlPlayingListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndexControlPlayingListResponse) String() string {
	return tea.Prettify(s)
}

func (s IndexControlPlayingListResponse) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponse) SetHeaders(v map[string]*string) *IndexControlPlayingListResponse {
	s.Headers = v
	return s
}

func (s *IndexControlPlayingListResponse) SetStatusCode(v int32) *IndexControlPlayingListResponse {
	s.StatusCode = &v
	return s
}

func (s *IndexControlPlayingListResponse) SetBody(v *IndexControlPlayingListResponseBody) *IndexControlPlayingListResponse {
	s.Body = v
	return s
}

type ListAlarmsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAlarmsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsHeaders) GoString() string {
	return s.String()
}

func (s *ListAlarmsHeaders) SetCommonHeaders(v map[string]*string) *ListAlarmsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAlarmsHeaders) SetXAcsAligenieAccessToken(v string) *ListAlarmsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAlarmsHeaders) SetAuthorization(v string) *ListAlarmsHeaders {
	s.Authorization = &v
	return s
}

type ListAlarmsRequest struct {
	DeviceInfo *ListAlarmsRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *ListAlarmsRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *ListAlarmsRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequest) SetDeviceInfo(v *ListAlarmsRequestDeviceInfo) *ListAlarmsRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListAlarmsRequest) SetPayload(v *ListAlarmsRequestPayload) *ListAlarmsRequest {
	s.Payload = v
	return s
}

func (s *ListAlarmsRequest) SetUserInfo(v *ListAlarmsRequestUserInfo) *ListAlarmsRequest {
	s.UserInfo = v
	return s
}

type ListAlarmsRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListAlarmsRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestDeviceInfo) SetEncodeKey(v string) *ListAlarmsRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetEncodeType(v string) *ListAlarmsRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetId(v string) *ListAlarmsRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetIdType(v string) *ListAlarmsRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetOrganizationId(v string) *ListAlarmsRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListAlarmsRequestPayload struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlarmsRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsRequestPayload) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestPayload) SetCurrentPage(v int32) *ListAlarmsRequestPayload {
	s.CurrentPage = &v
	return s
}

func (s *ListAlarmsRequestPayload) SetPageSize(v int32) *ListAlarmsRequestPayload {
	s.PageSize = &v
	return s
}

type ListAlarmsRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListAlarmsRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestUserInfo) SetEncodeKey(v string) *ListAlarmsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetEncodeType(v string) *ListAlarmsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetId(v string) *ListAlarmsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetIdType(v string) *ListAlarmsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetOrganizationId(v string) *ListAlarmsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListAlarmsShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListAlarmsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmsShrinkRequest) SetDeviceInfoShrink(v string) *ListAlarmsShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListAlarmsShrinkRequest) SetPayloadShrink(v string) *ListAlarmsShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListAlarmsShrinkRequest) SetUserInfoShrink(v string) *ListAlarmsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListAlarmsResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListAlarmsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBody) SetCode(v int32) *ListAlarmsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlarmsResponseBody) SetMessage(v string) *ListAlarmsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmsResponseBody) SetRequestId(v string) *ListAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmsResponseBody) SetResult(v *ListAlarmsResponseBodyResult) *ListAlarmsResponseBody {
	s.Result = v
	return s
}

type ListAlarmsResponseBodyResult struct {
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Model       []*ListAlarmsResponseBodyResultModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	PageCount   *int32                               `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize    *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResult) SetCurrentPage(v int32) *ListAlarmsResponseBodyResult {
	s.CurrentPage = &v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetModel(v []*ListAlarmsResponseBodyResultModel) *ListAlarmsResponseBodyResult {
	s.Model = v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetPageCount(v int32) *ListAlarmsResponseBodyResult {
	s.PageCount = &v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetPageSize(v int32) *ListAlarmsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetTotalCount(v int32) *ListAlarmsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListAlarmsResponseBodyResultModel struct {
	AlarmId          *int64                                         `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	MusicInfo        *ListAlarmsResponseBodyResultModelMusicInfo    `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	ScheduleInfo     *ListAlarmsResponseBodyResultModelScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	ScheduleTypeDesc *string                                        `json:"ScheduleTypeDesc,omitempty" xml:"ScheduleTypeDesc,omitempty"`
	Status           *int32                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TriggerDateDesc  *string                                        `json:"TriggerDateDesc,omitempty" xml:"TriggerDateDesc,omitempty"`
	TriggerTimeDesc  *string                                        `json:"TriggerTimeDesc,omitempty" xml:"TriggerTimeDesc,omitempty"`
	Volume           *int32                                         `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ListAlarmsResponseBodyResultModel) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModel) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModel) SetAlarmId(v int64) *ListAlarmsResponseBodyResultModel {
	s.AlarmId = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetMusicInfo(v *ListAlarmsResponseBodyResultModelMusicInfo) *ListAlarmsResponseBodyResultModel {
	s.MusicInfo = v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetScheduleInfo(v *ListAlarmsResponseBodyResultModelScheduleInfo) *ListAlarmsResponseBodyResultModel {
	s.ScheduleInfo = v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetScheduleTypeDesc(v string) *ListAlarmsResponseBodyResultModel {
	s.ScheduleTypeDesc = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetStatus(v int32) *ListAlarmsResponseBodyResultModel {
	s.Status = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetTriggerDateDesc(v string) *ListAlarmsResponseBodyResultModel {
	s.TriggerDateDesc = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetTriggerTimeDesc(v string) *ListAlarmsResponseBodyResultModel {
	s.TriggerTimeDesc = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetVolume(v int32) *ListAlarmsResponseBodyResultModel {
	s.Volume = &v
	return s
}

type ListAlarmsResponseBodyResultModelMusicInfo struct {
	MusicId       *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName     *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	MusicUrl      *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelMusicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelMusicInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicId(v int64) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicId = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicName(v string) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicName = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicType(v int64) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicType = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicTypeName(v string) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicUrl(v string) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicUrl = &v
	return s
}

type ListAlarmsResponseBodyResultModelScheduleInfo struct {
	Once                *ListAlarmsResponseBodyResultModelScheduleInfoOnce                `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	StatutoryWorkingDay *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	Type                *string                                                           `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly              *ListAlarmsResponseBodyResultModelScheduleInfoWeekly              `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetOnce(v *ListAlarmsResponseBodyResultModelScheduleInfoOnce) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.Once = v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetStatutoryWorkingDay(v *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetType(v string) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.Type = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetWeekly(v *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.Weekly = v
	return s
}

type ListAlarmsResponseBodyResultModelScheduleInfoOnce struct {
	Day    *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month  *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	Year   *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoOnce) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetDay(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetHour(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetMinute(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetMonth(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetYear(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Year = &v
	return s
}

type ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay struct {
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) SetHour(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

type ListAlarmsResponseBodyResultModelScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoWeekly) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *ListAlarmsResponseBodyResultModelScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) SetHour(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) SetMinute(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoWeekly {
	s.Minute = &v
	return s
}

type ListAlarmsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmsResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponse) SetHeaders(v map[string]*string) *ListAlarmsResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmsResponse) SetStatusCode(v int32) *ListAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmsResponse) SetBody(v *ListAlarmsResponseBody) *ListAlarmsResponse {
	s.Body = v
	return s
}

type ListAlbumDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAlbumDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailHeaders) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailHeaders) SetCommonHeaders(v map[string]*string) *ListAlbumDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAlbumDetailHeaders) SetXAcsAligenieAccessToken(v string) *ListAlbumDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAlbumDetailHeaders) SetAuthorization(v string) *ListAlbumDetailHeaders {
	s.Authorization = &v
	return s
}

type ListAlbumDetailRequest struct {
	Id       *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlbumDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailRequest) SetId(v int64) *ListAlbumDetailRequest {
	s.Id = &v
	return s
}

func (s *ListAlbumDetailRequest) SetPageNum(v int32) *ListAlbumDetailRequest {
	s.PageNum = &v
	return s
}

func (s *ListAlbumDetailRequest) SetPageSize(v int32) *ListAlbumDetailRequest {
	s.PageSize = &v
	return s
}

type ListAlbumDetailResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListAlbumDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAlbumDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBody) SetCode(v int32) *ListAlbumDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlbumDetailResponseBody) SetMessage(v string) *ListAlbumDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlbumDetailResponseBody) SetRequestId(v string) *ListAlbumDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlbumDetailResponseBody) SetResult(v *ListAlbumDetailResponseBodyResult) *ListAlbumDetailResponseBody {
	s.Result = v
	return s
}

type ListAlbumDetailResponseBodyResult struct {
	CurrentPageNum   *int64                                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	OpenDataItemList []*ListAlbumDetailResponseBodyResultOpenDataItemList `json:"OpenDataItemList,omitempty" xml:"OpenDataItemList,omitempty" type:"Repeated"`
	PageSize         *int64                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize        *int64                                               `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAlbumDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResult) SetCurrentPageNum(v int64) *ListAlbumDetailResponseBodyResult {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResult) SetOpenDataItemList(v []*ListAlbumDetailResponseBodyResultOpenDataItemList) *ListAlbumDetailResponseBodyResult {
	s.OpenDataItemList = v
	return s
}

func (s *ListAlbumDetailResponseBodyResult) SetPageSize(v int64) *ListAlbumDetailResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResult) SetTotalSize(v int64) *ListAlbumDetailResponseBodyResult {
	s.TotalSize = &v
	return s
}

type ListAlbumDetailResponseBodyResultOpenDataItemList struct {
	Alias       []*string                                                   `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	Audition    *bool                                                       `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors     []*ListAlbumDetailResponseBodyResultOpenDataItemListAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	Category    *string                                                     `json:"Category,omitempty" xml:"Category,omitempty"`
	Charge      *bool                                                       `json:"Charge,omitempty" xml:"Charge,omitempty"`
	CommCateId  *int64                                                      `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *ListAlbumDetailResponseBodyResultOpenDataItemListCover     `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *int64                                                      `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HotScore    *float64                                                    `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	Id          *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	ItemType    *string                                                     `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	OrderIndex  *int64                                                      `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	RawId       *string                                                     `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source      *string                                                     `json:"Source,omitempty" xml:"Source,omitempty"`
	Styles      []*string                                                   `json:"Styles,omitempty" xml:"Styles,omitempty" type:"Repeated"`
	Title       *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid       *string                                                     `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemList) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemList) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetAlias(v []*string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Alias = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetAudition(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Audition = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetAuthors(v []*ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Authors = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCategory(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Category = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCharge(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Charge = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCommCateId(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.CommCateId = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCover(v *ListAlbumDetailResponseBodyResultOpenDataItemListCover) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Cover = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetDescription(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Description = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetDuration(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Duration = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetHotScore(v float64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.HotScore = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetId(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Id = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetItemType(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.ItemType = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetOrderIndex(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.OrderIndex = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetRawId(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.RawId = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetSource(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Source = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetStyles(v []*string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Styles = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetTitle(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Title = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetType(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Type = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetValid(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Valid = &v
	return s
}

type ListAlbumDetailResponseBodyResultOpenDataItemListAuthors struct {
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Gender      *string   `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Id          *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Online      *bool     `json:"Online,omitempty" xml:"Online,omitempty"`
	Source      *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetAuthorTypes(v []*string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetGender(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Gender = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetId(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Id = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetOnline(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Online = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetSource(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Source = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetTitle(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Title = &v
	return s
}

type ListAlbumDetailResponseBodyResultOpenDataItemListCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemListCover) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemListCover) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetCanResize(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.CanResize = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetImg(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Img = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetLarge(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Large = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetMedium(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Medium = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetSmall(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Small = &v
	return s
}

type ListAlbumDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlbumDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlbumDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumDetailResponse) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponse) SetHeaders(v map[string]*string) *ListAlbumDetailResponse {
	s.Headers = v
	return s
}

func (s *ListAlbumDetailResponse) SetStatusCode(v int32) *ListAlbumDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlbumDetailResponse) SetBody(v *ListAlbumDetailResponseBody) *ListAlbumDetailResponse {
	s.Body = v
	return s
}

type ListAlbumIsAddedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAlbumIsAddedHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedHeaders) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedHeaders) SetCommonHeaders(v map[string]*string) *ListAlbumIsAddedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAlbumIsAddedHeaders) SetXAcsAligenieAccessToken(v string) *ListAlbumIsAddedHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAlbumIsAddedHeaders) SetAuthorization(v string) *ListAlbumIsAddedHeaders {
	s.Authorization = &v
	return s
}

type ListAlbumIsAddedRequest struct {
	AlbumIdList []*string                          `json:"AlbumIdList,omitempty" xml:"AlbumIdList,omitempty" type:"Repeated"`
	DeviceInfo  *ListAlbumIsAddedRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo    *ListAlbumIsAddedRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListAlbumIsAddedRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedRequest) SetAlbumIdList(v []*string) *ListAlbumIsAddedRequest {
	s.AlbumIdList = v
	return s
}

func (s *ListAlbumIsAddedRequest) SetDeviceInfo(v *ListAlbumIsAddedRequestDeviceInfo) *ListAlbumIsAddedRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListAlbumIsAddedRequest) SetUserInfo(v *ListAlbumIsAddedRequestUserInfo) *ListAlbumIsAddedRequest {
	s.UserInfo = v
	return s
}

type ListAlbumIsAddedRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListAlbumIsAddedRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetEncodeKey(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetEncodeType(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetId(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetIdType(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetOrganizationId(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListAlbumIsAddedRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListAlbumIsAddedRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedRequestUserInfo) SetEncodeKey(v string) *ListAlbumIsAddedRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetEncodeType(v string) *ListAlbumIsAddedRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetId(v string) *ListAlbumIsAddedRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetIdType(v string) *ListAlbumIsAddedRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetOrganizationId(v string) *ListAlbumIsAddedRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListAlbumIsAddedShrinkRequest struct {
	AlbumIdListShrink *string `json:"AlbumIdList,omitempty" xml:"AlbumIdList,omitempty"`
	DeviceInfoShrink  *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink    *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListAlbumIsAddedShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedShrinkRequest) SetAlbumIdListShrink(v string) *ListAlbumIsAddedShrinkRequest {
	s.AlbumIdListShrink = &v
	return s
}

func (s *ListAlbumIsAddedShrinkRequest) SetDeviceInfoShrink(v string) *ListAlbumIsAddedShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListAlbumIsAddedShrinkRequest) SetUserInfoShrink(v string) *ListAlbumIsAddedShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListAlbumIsAddedResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListAlbumIsAddedResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAlbumIsAddedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedResponseBody) SetCode(v int32) *ListAlbumIsAddedResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlbumIsAddedResponseBody) SetMessage(v string) *ListAlbumIsAddedResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlbumIsAddedResponseBody) SetRequestId(v string) *ListAlbumIsAddedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlbumIsAddedResponseBody) SetResult(v []*ListAlbumIsAddedResponseBodyResult) *ListAlbumIsAddedResponseBody {
	s.Result = v
	return s
}

type ListAlbumIsAddedResponseBodyResult struct {
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	IsAdded *string `json:"IsAdded,omitempty" xml:"IsAdded,omitempty"`
}

func (s ListAlbumIsAddedResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedResponseBodyResult) SetAlbumId(v string) *ListAlbumIsAddedResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *ListAlbumIsAddedResponseBodyResult) SetIsAdded(v string) *ListAlbumIsAddedResponseBodyResult {
	s.IsAdded = &v
	return s
}

type ListAlbumIsAddedResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlbumIsAddedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlbumIsAddedResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumIsAddedResponse) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedResponse) SetHeaders(v map[string]*string) *ListAlbumIsAddedResponse {
	s.Headers = v
	return s
}

func (s *ListAlbumIsAddedResponse) SetStatusCode(v int32) *ListAlbumIsAddedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlbumIsAddedResponse) SetBody(v *ListAlbumIsAddedResponseBody) *ListAlbumIsAddedResponse {
	s.Body = v
	return s
}

type ListCateContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCateContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentHeaders) GoString() string {
	return s.String()
}

func (s *ListCateContentHeaders) SetCommonHeaders(v map[string]*string) *ListCateContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCateContentHeaders) SetXAcsAligenieAccessToken(v string) *ListCateContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCateContentHeaders) SetAuthorization(v string) *ListCateContentHeaders {
	s.Authorization = &v
	return s
}

type ListCateContentRequest struct {
	DeviceInfo *ListCateContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Request    *ListCateContentRequestRequest    `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	UserInfo   *ListCateContentRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListCateContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentRequest) GoString() string {
	return s.String()
}

func (s *ListCateContentRequest) SetDeviceInfo(v *ListCateContentRequestDeviceInfo) *ListCateContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListCateContentRequest) SetRequest(v *ListCateContentRequestRequest) *ListCateContentRequest {
	s.Request = v
	return s
}

func (s *ListCateContentRequest) SetUserInfo(v *ListCateContentRequestUserInfo) *ListCateContentRequest {
	s.UserInfo = v
	return s
}

type ListCateContentRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListCateContentRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListCateContentRequestDeviceInfo) SetEncodeKey(v string) *ListCateContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetEncodeType(v string) *ListCateContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetId(v string) *ListCateContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetIdType(v string) *ListCateContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetOrganizationId(v string) *ListCateContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListCateContentRequestRequest struct {
	CateId    *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	IsAlbum   *bool   `json:"IsAlbum,omitempty" xml:"IsAlbum,omitempty"`
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy    *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListCateContentRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentRequestRequest) GoString() string {
	return s.String()
}

func (s *ListCateContentRequestRequest) SetCateId(v int64) *ListCateContentRequestRequest {
	s.CateId = &v
	return s
}

func (s *ListCateContentRequestRequest) SetIsAlbum(v bool) *ListCateContentRequestRequest {
	s.IsAlbum = &v
	return s
}

func (s *ListCateContentRequestRequest) SetPageNum(v int32) *ListCateContentRequestRequest {
	s.PageNum = &v
	return s
}

func (s *ListCateContentRequestRequest) SetPageSize(v int32) *ListCateContentRequestRequest {
	s.PageSize = &v
	return s
}

func (s *ListCateContentRequestRequest) SetSortBy(v string) *ListCateContentRequestRequest {
	s.SortBy = &v
	return s
}

func (s *ListCateContentRequestRequest) SetSortOrder(v string) *ListCateContentRequestRequest {
	s.SortOrder = &v
	return s
}

type ListCateContentRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListCateContentRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListCateContentRequestUserInfo) SetEncodeKey(v string) *ListCateContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetEncodeType(v string) *ListCateContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetId(v string) *ListCateContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetIdType(v string) *ListCateContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetOrganizationId(v string) *ListCateContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListCateContentShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	RequestShrink    *string `json:"Request,omitempty" xml:"Request,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListCateContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCateContentShrinkRequest) SetDeviceInfoShrink(v string) *ListCateContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListCateContentShrinkRequest) SetRequestShrink(v string) *ListCateContentShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *ListCateContentShrinkRequest) SetUserInfoShrink(v string) *ListCateContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListCateContentResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListCateContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListCateContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentResponseBody) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBody) SetCode(v int32) *ListCateContentResponseBody {
	s.Code = &v
	return s
}

func (s *ListCateContentResponseBody) SetMessage(v string) *ListCateContentResponseBody {
	s.Message = &v
	return s
}

func (s *ListCateContentResponseBody) SetRequestId(v string) *ListCateContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCateContentResponseBody) SetResult(v *ListCateContentResponseBodyResult) *ListCateContentResponseBody {
	s.Result = v
	return s
}

type ListCateContentResponseBodyResult struct {
	CurrentPageNum   *int32                                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	OpenDataItemList []*ListCateContentResponseBodyResultOpenDataItemList `json:"OpenDataItemList,omitempty" xml:"OpenDataItemList,omitempty" type:"Repeated"`
	PageSize         *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize        *int64                                               `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListCateContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResult) SetCurrentPageNum(v int32) *ListCateContentResponseBodyResult {
	s.CurrentPageNum = &v
	return s
}

func (s *ListCateContentResponseBodyResult) SetOpenDataItemList(v []*ListCateContentResponseBodyResultOpenDataItemList) *ListCateContentResponseBodyResult {
	s.OpenDataItemList = v
	return s
}

func (s *ListCateContentResponseBodyResult) SetPageSize(v int32) *ListCateContentResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListCateContentResponseBodyResult) SetTotalSize(v int64) *ListCateContentResponseBodyResult {
	s.TotalSize = &v
	return s
}

type ListCateContentResponseBodyResultOpenDataItemList struct {
	Alias       []*string                                                   `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	Audition    *bool                                                       `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors     []*ListCateContentResponseBodyResultOpenDataItemListAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	Category    *string                                                     `json:"Category,omitempty" xml:"Category,omitempty"`
	Charge      *bool                                                       `json:"Charge,omitempty" xml:"Charge,omitempty"`
	CommCateId  *string                                                     `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *ListCateContentResponseBodyResultOpenDataItemListCover     `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	HotScore    *float64                                                    `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	ItemType    *string                                                     `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	RawId       *string                                                     `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source      *string                                                     `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid       *string                                                     `json:"Valid,omitempty" xml:"Valid,omitempty"`
	Id          *int64                                                      `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListCateContentResponseBodyResultOpenDataItemList) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemList) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetAlias(v []*string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Alias = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetAudition(v bool) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Audition = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetAuthors(v []*ListCateContentResponseBodyResultOpenDataItemListAuthors) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Authors = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCategory(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Category = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCharge(v bool) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Charge = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCommCateId(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.CommCateId = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCover(v *ListCateContentResponseBodyResultOpenDataItemListCover) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Cover = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetDescription(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Description = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetHotScore(v float64) *ListCateContentResponseBodyResultOpenDataItemList {
	s.HotScore = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetItemType(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.ItemType = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetRawId(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.RawId = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetSource(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Source = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetTitle(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Title = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetType(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Type = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetValid(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Valid = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetId(v int64) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Id = &v
	return s
}

type ListCateContentResponseBodyResultOpenDataItemListAuthors struct {
	AuthorTypes []*string                                                      `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Cover       *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Gender      *string                                                        `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Id          *int64                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Online      *bool                                                          `json:"Online,omitempty" xml:"Online,omitempty"`
	RawId       *string                                                        `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source      *string                                                        `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string                                                        `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListCateContentResponseBodyResultOpenDataItemListAuthors) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemListAuthors) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetAuthorTypes(v []*string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetCover(v *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Cover = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetDescription(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Description = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetGender(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Gender = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetId(v int64) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Id = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetOnline(v bool) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Online = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetRawId(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.RawId = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetSource(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Source = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetTitle(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Title = &v
	return s
}

type ListCateContentResponseBodyResultOpenDataItemListAuthorsCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetCanResize(v bool) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetImg(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Img = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetLarge(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Large = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetMediam(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Mediam = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetMedium(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Medium = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetSmall(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Small = &v
	return s
}

type ListCateContentResponseBodyResultOpenDataItemListCover struct {
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
	CanResize *bool   `json:"canResize,omitempty" xml:"canResize,omitempty"`
}

func (s ListCateContentResponseBodyResultOpenDataItemListCover) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemListCover) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetImg(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Img = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetLarge(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Large = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetMediam(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Mediam = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetMedium(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Medium = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetSmall(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Small = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetCanResize(v bool) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.CanResize = &v
	return s
}

type ListCateContentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCateContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCateContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCateContentResponse) GoString() string {
	return s.String()
}

func (s *ListCateContentResponse) SetHeaders(v map[string]*string) *ListCateContentResponse {
	s.Headers = v
	return s
}

func (s *ListCateContentResponse) SetStatusCode(v int32) *ListCateContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCateContentResponse) SetBody(v *ListCateContentResponseBody) *ListCateContentResponse {
	s.Body = v
	return s
}

type ListCateInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCateInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCateInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListCateInfoHeaders) SetCommonHeaders(v map[string]*string) *ListCateInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCateInfoHeaders) SetXAcsAligenieAccessToken(v string) *ListCateInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCateInfoHeaders) SetAuthorization(v string) *ListCateInfoHeaders {
	s.Authorization = &v
	return s
}

type ListCateInfoRequest struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCateInfoRequest) GoString() string {
	return s.String()
}

func (s *ListCateInfoRequest) SetType(v string) *ListCateInfoRequest {
	s.Type = &v
	return s
}

type ListCateInfoResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCateInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListCateInfoResponseBody) SetCode(v int32) *ListCateInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListCateInfoResponseBody) SetMessage(v string) *ListCateInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListCateInfoResponseBody) SetRequestId(v string) *ListCateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCateInfoResponseBody) SetResult(v []*ListCateInfoResponseBodyResult) *ListCateInfoResponseBody {
	s.Result = v
	return s
}

type ListCateInfoResponseBodyResult struct {
	CateId       *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName     *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	ParentCateId *int64  `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCateInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCateInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCateInfoResponseBodyResult) SetCateId(v int64) *ListCateInfoResponseBodyResult {
	s.CateId = &v
	return s
}

func (s *ListCateInfoResponseBodyResult) SetCateName(v string) *ListCateInfoResponseBodyResult {
	s.CateName = &v
	return s
}

func (s *ListCateInfoResponseBodyResult) SetParentCateId(v int64) *ListCateInfoResponseBodyResult {
	s.ParentCateId = &v
	return s
}

type ListCateInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCateInfoResponse) GoString() string {
	return s.String()
}

func (s *ListCateInfoResponse) SetHeaders(v map[string]*string) *ListCateInfoResponse {
	s.Headers = v
	return s
}

func (s *ListCateInfoResponse) SetStatusCode(v int32) *ListCateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCateInfoResponse) SetBody(v *ListCateInfoResponseBody) *ListCateInfoResponse {
	s.Body = v
	return s
}

type ListCommonCateFirstFloorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCommonCateFirstFloorHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateFirstFloorHeaders) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorHeaders) SetCommonHeaders(v map[string]*string) *ListCommonCateFirstFloorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCommonCateFirstFloorHeaders) SetXAcsAligenieAccessToken(v string) *ListCommonCateFirstFloorHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCommonCateFirstFloorHeaders) SetAuthorization(v string) *ListCommonCateFirstFloorHeaders {
	s.Authorization = &v
	return s
}

type ListCommonCateFirstFloorRequest struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCommonCateFirstFloorRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateFirstFloorRequest) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorRequest) SetType(v string) *ListCommonCateFirstFloorRequest {
	s.Type = &v
	return s
}

type ListCommonCateFirstFloorResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCommonCateFirstFloorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCommonCateFirstFloorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateFirstFloorResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorResponseBody) SetCode(v int32) *ListCommonCateFirstFloorResponseBody {
	s.Code = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBody) SetMessage(v string) *ListCommonCateFirstFloorResponseBody {
	s.Message = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBody) SetRequestId(v string) *ListCommonCateFirstFloorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBody) SetResult(v []*ListCommonCateFirstFloorResponseBodyResult) *ListCommonCateFirstFloorResponseBody {
	s.Result = v
	return s
}

type ListCommonCateFirstFloorResponseBodyResult struct {
	CateId       *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName     *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	ParentCateId *int64  `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCommonCateFirstFloorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateFirstFloorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorResponseBodyResult) SetCateId(v int64) *ListCommonCateFirstFloorResponseBodyResult {
	s.CateId = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBodyResult) SetCateName(v string) *ListCommonCateFirstFloorResponseBodyResult {
	s.CateName = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBodyResult) SetParentCateId(v int64) *ListCommonCateFirstFloorResponseBodyResult {
	s.ParentCateId = &v
	return s
}

type ListCommonCateFirstFloorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCommonCateFirstFloorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommonCateFirstFloorResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateFirstFloorResponse) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorResponse) SetHeaders(v map[string]*string) *ListCommonCateFirstFloorResponse {
	s.Headers = v
	return s
}

func (s *ListCommonCateFirstFloorResponse) SetStatusCode(v int32) *ListCommonCateFirstFloorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonCateFirstFloorResponse) SetBody(v *ListCommonCateFirstFloorResponseBody) *ListCommonCateFirstFloorResponse {
	s.Body = v
	return s
}

type ListCommonCateSecondFloorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCommonCateSecondFloorHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateSecondFloorHeaders) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorHeaders) SetCommonHeaders(v map[string]*string) *ListCommonCateSecondFloorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCommonCateSecondFloorHeaders) SetXAcsAligenieAccessToken(v string) *ListCommonCateSecondFloorHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCommonCateSecondFloorHeaders) SetAuthorization(v string) *ListCommonCateSecondFloorHeaders {
	s.Authorization = &v
	return s
}

type ListCommonCateSecondFloorRequest struct {
	ParentCateId *int64 `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCommonCateSecondFloorRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateSecondFloorRequest) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorRequest) SetParentCateId(v int64) *ListCommonCateSecondFloorRequest {
	s.ParentCateId = &v
	return s
}

type ListCommonCateSecondFloorResponseBody struct {
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCommonCateSecondFloorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCommonCateSecondFloorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateSecondFloorResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorResponseBody) SetCode(v int32) *ListCommonCateSecondFloorResponseBody {
	s.Code = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBody) SetMessage(v string) *ListCommonCateSecondFloorResponseBody {
	s.Message = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBody) SetRequestId(v string) *ListCommonCateSecondFloorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBody) SetResult(v []*ListCommonCateSecondFloorResponseBodyResult) *ListCommonCateSecondFloorResponseBody {
	s.Result = v
	return s
}

type ListCommonCateSecondFloorResponseBodyResult struct {
	CateId       *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName     *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	ParentCateId *int64  `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCommonCateSecondFloorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateSecondFloorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorResponseBodyResult) SetCateId(v int64) *ListCommonCateSecondFloorResponseBodyResult {
	s.CateId = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBodyResult) SetCateName(v string) *ListCommonCateSecondFloorResponseBodyResult {
	s.CateName = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBodyResult) SetParentCateId(v int64) *ListCommonCateSecondFloorResponseBodyResult {
	s.ParentCateId = &v
	return s
}

type ListCommonCateSecondFloorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCommonCateSecondFloorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommonCateSecondFloorResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommonCateSecondFloorResponse) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorResponse) SetHeaders(v map[string]*string) *ListCommonCateSecondFloorResponse {
	s.Headers = v
	return s
}

func (s *ListCommonCateSecondFloorResponse) SetStatusCode(v int32) *ListCommonCateSecondFloorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonCateSecondFloorResponse) SetBody(v *ListCommonCateSecondFloorResponseBody) *ListCommonCateSecondFloorResponse {
	s.Body = v
	return s
}

type ListDeviceBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceBasicInfoHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceBasicInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceBasicInfoHeaders) SetAuthorization(v string) *ListDeviceBasicInfoHeaders {
	s.Authorization = &v
	return s
}

type ListDeviceBasicInfoRequest struct {
	DeviceInfos *ListDeviceBasicInfoRequestDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Struct"`
}

func (s ListDeviceBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoRequest) SetDeviceInfos(v *ListDeviceBasicInfoRequestDeviceInfos) *ListDeviceBasicInfoRequest {
	s.DeviceInfos = v
	return s
}

type ListDeviceBasicInfoRequestDeviceInfos struct {
	EncodeKey      *string   `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string   `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	IdType         *string   `json:"IdType,omitempty" xml:"IdType,omitempty"`
	Ids            []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	OrganizationId *string   `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceBasicInfoRequestDeviceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBasicInfoRequestDeviceInfos) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetEncodeKey(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetEncodeType(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetIdType(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.IdType = &v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetIds(v []*string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.Ids = v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetOrganizationId(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.OrganizationId = &v
	return s
}

type ListDeviceBasicInfoShrinkRequest struct {
	DeviceInfosShrink *string `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty"`
}

func (s ListDeviceBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoShrinkRequest) SetDeviceInfosShrink(v string) *ListDeviceBasicInfoShrinkRequest {
	s.DeviceInfosShrink = &v
	return s
}

type ListDeviceBasicInfoResponseBody struct {
	Code      *int32                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]*ResultValue `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ListDeviceBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoResponseBody) SetCode(v int32) *ListDeviceBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceBasicInfoResponseBody) SetMessage(v string) *ListDeviceBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceBasicInfoResponseBody) SetRequestId(v string) *ListDeviceBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceBasicInfoResponseBody) SetResult(v map[string]*ResultValue) *ListDeviceBasicInfoResponseBody {
	s.Result = v
	return s
}

type ListDeviceBasicInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoResponse) SetHeaders(v map[string]*string) *ListDeviceBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceBasicInfoResponse) SetStatusCode(v int32) *ListDeviceBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceBasicInfoResponse) SetBody(v *ListDeviceBasicInfoResponseBody) *ListDeviceBasicInfoResponse {
	s.Body = v
	return s
}

type ListDeviceByUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceByUserIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceByUserIdHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceByUserIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceByUserIdHeaders) SetAuthorization(v string) *ListDeviceByUserIdHeaders {
	s.Authorization = &v
	return s
}

type ListDeviceByUserIdRequest struct {
	UserInfo *ListDeviceByUserIdRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListDeviceByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdRequest) SetUserInfo(v *ListDeviceByUserIdRequestUserInfo) *ListDeviceByUserIdRequest {
	s.UserInfo = v
	return s
}

type ListDeviceByUserIdRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdRequestUserInfo) SetEncodeKey(v string) *ListDeviceByUserIdRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetEncodeType(v string) *ListDeviceByUserIdRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetId(v string) *ListDeviceByUserIdRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetIdType(v string) *ListDeviceByUserIdRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetOrganizationId(v string) *ListDeviceByUserIdRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListDeviceByUserIdShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListDeviceByUserIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdShrinkRequest) SetUserInfoShrink(v string) *ListDeviceByUserIdShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListDeviceByUserIdResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDeviceByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponseBody) SetCode(v int32) *ListDeviceByUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceByUserIdResponseBody) SetMessage(v string) *ListDeviceByUserIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceByUserIdResponseBody) SetRequestId(v string) *ListDeviceByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceByUserIdResponseBody) SetResult(v []*ListDeviceByUserIdResponseBodyResult) *ListDeviceByUserIdResponseBody {
	s.Result = v
	return s
}

type ListDeviceByUserIdResponseBodyResult struct {
	DeviceOpenId   *string                                               `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds []*ListDeviceByUserIdResponseBodyResultDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponseBodyResult) SetDeviceOpenId(v string) *ListDeviceByUserIdResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *ListDeviceByUserIdResponseBodyResult) SetDeviceUnionIds(v []*ListDeviceByUserIdResponseBodyResultDeviceUnionIds) *ListDeviceByUserIdResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

type ListDeviceByUserIdResponseBodyResultDeviceUnionIds struct {
	DeviceUnionId  *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdResponseBodyResultDeviceUnionIds) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdResponseBodyResultDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponseBodyResultDeviceUnionIds) SetDeviceUnionId(v string) *ListDeviceByUserIdResponseBodyResultDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

func (s *ListDeviceByUserIdResponseBodyResultDeviceUnionIds) SetOrganizationId(v string) *ListDeviceByUserIdResponseBodyResultDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

type ListDeviceByUserIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponse) SetHeaders(v map[string]*string) *ListDeviceByUserIdResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceByUserIdResponse) SetStatusCode(v int32) *ListDeviceByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceByUserIdResponse) SetBody(v *ListDeviceByUserIdResponseBody) *ListDeviceByUserIdResponse {
	s.Body = v
	return s
}

type ListDeviceByUserIdAndChanelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceByUserIdAndChanelHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceByUserIdAndChanelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceByUserIdAndChanelHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceByUserIdAndChanelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelHeaders) SetAuthorization(v string) *ListDeviceByUserIdAndChanelHeaders {
	s.Authorization = &v
	return s
}

type ListDeviceByUserIdAndChanelRequest struct {
	ChannelInfo *ListDeviceByUserIdAndChanelRequestChannelInfo `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty" type:"Struct"`
	UserInfo    *ListDeviceByUserIdAndChanelRequestUserInfo    `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListDeviceByUserIdAndChanelRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelRequest) SetChannelInfo(v *ListDeviceByUserIdAndChanelRequestChannelInfo) *ListDeviceByUserIdAndChanelRequest {
	s.ChannelInfo = v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequest) SetUserInfo(v *ListDeviceByUserIdAndChanelRequestUserInfo) *ListDeviceByUserIdAndChanelRequest {
	s.UserInfo = v
	return s
}

type ListDeviceByUserIdAndChanelRequestChannelInfo struct {
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s ListDeviceByUserIdAndChanelRequestChannelInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelRequestChannelInfo) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelRequestChannelInfo) SetChannel(v string) *ListDeviceByUserIdAndChanelRequestChannelInfo {
	s.Channel = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestChannelInfo) SetExtInfo(v string) *ListDeviceByUserIdAndChanelRequestChannelInfo {
	s.ExtInfo = &v
	return s
}

type ListDeviceByUserIdAndChanelRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdAndChanelRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetEncodeKey(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetEncodeType(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetId(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetIdType(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetOrganizationId(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListDeviceByUserIdAndChanelShrinkRequest struct {
	ChannelInfoShrink *string `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty"`
	UserInfoShrink    *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListDeviceByUserIdAndChanelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelShrinkRequest) SetChannelInfoShrink(v string) *ListDeviceByUserIdAndChanelShrinkRequest {
	s.ChannelInfoShrink = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelShrinkRequest) SetUserInfoShrink(v string) *ListDeviceByUserIdAndChanelShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListDeviceByUserIdAndChanelResponseBody struct {
	Code      *int32                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDeviceByUserIdAndChanelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdAndChanelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetCode(v int32) *ListDeviceByUserIdAndChanelResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetMessage(v string) *ListDeviceByUserIdAndChanelResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetRequestId(v string) *ListDeviceByUserIdAndChanelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetResult(v []*ListDeviceByUserIdAndChanelResponseBodyResult) *ListDeviceByUserIdAndChanelResponseBody {
	s.Result = v
	return s
}

type ListDeviceByUserIdAndChanelResponseBodyResult struct {
	DeviceOpenId   *string                                                        `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds []*ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdAndChanelResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResult) SetDeviceOpenId(v string) *ListDeviceByUserIdAndChanelResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResult) SetDeviceUnionIds(v []*ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) *ListDeviceByUserIdAndChanelResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

type ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds struct {
	DeviceUnionId  *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) SetDeviceUnionId(v string) *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) SetOrganizationId(v string) *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

type ListDeviceByUserIdAndChanelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceByUserIdAndChanelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceByUserIdAndChanelResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponse) SetHeaders(v map[string]*string) *ListDeviceByUserIdAndChanelResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponse) SetStatusCode(v int32) *ListDeviceByUserIdAndChanelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponse) SetBody(v *ListDeviceByUserIdAndChanelResponseBody) *ListDeviceByUserIdAndChanelResponse {
	s.Body = v
	return s
}

type ListDeviceIdByIdentitiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceIdByIdentitiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIdByIdentitiesHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceIdByIdentitiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceIdByIdentitiesHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceIdByIdentitiesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceIdByIdentitiesHeaders) SetAuthorization(v string) *ListDeviceIdByIdentitiesHeaders {
	s.Authorization = &v
	return s
}

type ListDeviceIdByIdentitiesRequest struct {
	EncodeKey    *string   `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType   *string   `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	IdentityIds  []*string `json:"IdentityIds,omitempty" xml:"IdentityIds,omitempty" type:"Repeated"`
	IdentityType *string   `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	ProductKey   *string   `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ListDeviceIdByIdentitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIdByIdentitiesRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesRequest) SetEncodeKey(v string) *ListDeviceIdByIdentitiesRequest {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetEncodeType(v string) *ListDeviceIdByIdentitiesRequest {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetIdentityIds(v []*string) *ListDeviceIdByIdentitiesRequest {
	s.IdentityIds = v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetIdentityType(v string) *ListDeviceIdByIdentitiesRequest {
	s.IdentityType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetProductKey(v string) *ListDeviceIdByIdentitiesRequest {
	s.ProductKey = &v
	return s
}

type ListDeviceIdByIdentitiesShrinkRequest struct {
	EncodeKey         *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType        *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	IdentityIdsShrink *string `json:"IdentityIds,omitempty" xml:"IdentityIds,omitempty"`
	IdentityType      *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	ProductKey        *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ListDeviceIdByIdentitiesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIdByIdentitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetEncodeKey(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetEncodeType(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetIdentityIdsShrink(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.IdentityIdsShrink = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetIdentityType(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.IdentityType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetProductKey(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.ProductKey = &v
	return s
}

type ListDeviceIdByIdentitiesResponseBody struct {
	Code      *int32                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]*ResultValue `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ListDeviceIdByIdentitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIdByIdentitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetCode(v int32) *ListDeviceIdByIdentitiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetMessage(v string) *ListDeviceIdByIdentitiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetRequestId(v string) *ListDeviceIdByIdentitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetResult(v map[string]*ResultValue) *ListDeviceIdByIdentitiesResponseBody {
	s.Result = v
	return s
}

type ListDeviceIdByIdentitiesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceIdByIdentitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceIdByIdentitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIdByIdentitiesResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesResponse) SetHeaders(v map[string]*string) *ListDeviceIdByIdentitiesResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceIdByIdentitiesResponse) SetStatusCode(v int32) *ListDeviceIdByIdentitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponse) SetBody(v *ListDeviceIdByIdentitiesResponseBody) *ListDeviceIdByIdentitiesResponse {
	s.Body = v
	return s
}

type ListMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListMusicHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMusicHeaders) GoString() string {
	return s.String()
}

func (s *ListMusicHeaders) SetCommonHeaders(v map[string]*string) *ListMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMusicHeaders) SetXAcsAligenieAccessToken(v string) *ListMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListMusicHeaders) SetAuthorization(v string) *ListMusicHeaders {
	s.Authorization = &v
	return s
}

type ListMusicRequest struct {
	DeviceInfo *ListMusicRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *ListMusicRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *ListMusicRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListMusicRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMusicRequest) GoString() string {
	return s.String()
}

func (s *ListMusicRequest) SetDeviceInfo(v *ListMusicRequestDeviceInfo) *ListMusicRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListMusicRequest) SetPayload(v *ListMusicRequestPayload) *ListMusicRequest {
	s.Payload = v
	return s
}

func (s *ListMusicRequest) SetUserInfo(v *ListMusicRequestUserInfo) *ListMusicRequest {
	s.UserInfo = v
	return s
}

type ListMusicRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListMusicRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMusicRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListMusicRequestDeviceInfo) SetEncodeKey(v string) *ListMusicRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetEncodeType(v string) *ListMusicRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetId(v string) *ListMusicRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetIdType(v string) *ListMusicRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetOrganizationId(v string) *ListMusicRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListMusicRequestPayload struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MusicId       *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName     *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMusicRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListMusicRequestPayload) GoString() string {
	return s.String()
}

func (s *ListMusicRequestPayload) SetCurrentPage(v int32) *ListMusicRequestPayload {
	s.CurrentPage = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicId(v int64) *ListMusicRequestPayload {
	s.MusicId = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicName(v string) *ListMusicRequestPayload {
	s.MusicName = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicType(v int64) *ListMusicRequestPayload {
	s.MusicType = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicTypeName(v string) *ListMusicRequestPayload {
	s.MusicTypeName = &v
	return s
}

func (s *ListMusicRequestPayload) SetPageSize(v int32) *ListMusicRequestPayload {
	s.PageSize = &v
	return s
}

type ListMusicRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListMusicRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMusicRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListMusicRequestUserInfo) SetEncodeKey(v string) *ListMusicRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetEncodeType(v string) *ListMusicRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetId(v string) *ListMusicRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetIdType(v string) *ListMusicRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetOrganizationId(v string) *ListMusicRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListMusicShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListMusicShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMusicShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMusicShrinkRequest) SetDeviceInfoShrink(v string) *ListMusicShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListMusicShrinkRequest) SetPayloadShrink(v string) *ListMusicShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListMusicShrinkRequest) SetUserInfoShrink(v string) *ListMusicShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListMusicResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListMusicResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMusicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMusicResponseBody) GoString() string {
	return s.String()
}

func (s *ListMusicResponseBody) SetCode(v int32) *ListMusicResponseBody {
	s.Code = &v
	return s
}

func (s *ListMusicResponseBody) SetMessage(v string) *ListMusicResponseBody {
	s.Message = &v
	return s
}

func (s *ListMusicResponseBody) SetRequestId(v string) *ListMusicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMusicResponseBody) SetResult(v *ListMusicResponseBodyResult) *ListMusicResponseBody {
	s.Result = v
	return s
}

type ListMusicResponseBodyResult struct {
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Model       []*ListMusicResponseBodyResultModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	PageCount   *int32                              `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMusicResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListMusicResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMusicResponseBodyResult) SetCurrentPage(v int32) *ListMusicResponseBodyResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMusicResponseBodyResult) SetModel(v []*ListMusicResponseBodyResultModel) *ListMusicResponseBodyResult {
	s.Model = v
	return s
}

func (s *ListMusicResponseBodyResult) SetPageCount(v int32) *ListMusicResponseBodyResult {
	s.PageCount = &v
	return s
}

func (s *ListMusicResponseBodyResult) SetPageSize(v int32) *ListMusicResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListMusicResponseBodyResult) SetTotalCount(v int32) *ListMusicResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListMusicResponseBodyResultModel struct {
	MusicId       *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName     *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	MusicUrl      *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s ListMusicResponseBodyResultModel) String() string {
	return tea.Prettify(s)
}

func (s ListMusicResponseBodyResultModel) GoString() string {
	return s.String()
}

func (s *ListMusicResponseBodyResultModel) SetMusicId(v int64) *ListMusicResponseBodyResultModel {
	s.MusicId = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicName(v string) *ListMusicResponseBodyResultModel {
	s.MusicName = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicType(v int64) *ListMusicResponseBodyResultModel {
	s.MusicType = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicTypeName(v string) *ListMusicResponseBodyResultModel {
	s.MusicTypeName = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicUrl(v string) *ListMusicResponseBodyResultModel {
	s.MusicUrl = &v
	return s
}

type ListMusicResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMusicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMusicResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMusicResponse) GoString() string {
	return s.String()
}

func (s *ListMusicResponse) SetHeaders(v map[string]*string) *ListMusicResponse {
	s.Headers = v
	return s
}

func (s *ListMusicResponse) SetStatusCode(v int32) *ListMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMusicResponse) SetBody(v *ListMusicResponseBody) *ListMusicResponse {
	s.Body = v
	return s
}

type ListPlayHistoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListPlayHistoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryHeaders) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryHeaders) SetCommonHeaders(v map[string]*string) *ListPlayHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPlayHistoryHeaders) SetXAcsAligenieAccessToken(v string) *ListPlayHistoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListPlayHistoryHeaders) SetAuthorization(v string) *ListPlayHistoryHeaders {
	s.Authorization = &v
	return s
}

type ListPlayHistoryRequest struct {
	DeviceInfo *ListPlayHistoryRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Request    *ListPlayHistoryRequestRequest    `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	UserInfo   *ListPlayHistoryRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListPlayHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequest) SetDeviceInfo(v *ListPlayHistoryRequestDeviceInfo) *ListPlayHistoryRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListPlayHistoryRequest) SetRequest(v *ListPlayHistoryRequestRequest) *ListPlayHistoryRequest {
	s.Request = v
	return s
}

func (s *ListPlayHistoryRequest) SetUserInfo(v *ListPlayHistoryRequestUserInfo) *ListPlayHistoryRequest {
	s.UserInfo = v
	return s
}

type ListPlayHistoryRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListPlayHistoryRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequestDeviceInfo) SetEncodeKey(v string) *ListPlayHistoryRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetEncodeType(v string) *ListPlayHistoryRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetId(v string) *ListPlayHistoryRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetIdType(v string) *ListPlayHistoryRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetOrganizationId(v string) *ListPlayHistoryRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListPlayHistoryRequestRequest struct {
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPlayHistoryRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryRequestRequest) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequestRequest) SetPageNum(v int32) *ListPlayHistoryRequestRequest {
	s.PageNum = &v
	return s
}

func (s *ListPlayHistoryRequestRequest) SetPageSize(v int32) *ListPlayHistoryRequestRequest {
	s.PageSize = &v
	return s
}

func (s *ListPlayHistoryRequestRequest) SetType(v string) *ListPlayHistoryRequestRequest {
	s.Type = &v
	return s
}

type ListPlayHistoryRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListPlayHistoryRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequestUserInfo) SetEncodeKey(v string) *ListPlayHistoryRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetEncodeType(v string) *ListPlayHistoryRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetId(v string) *ListPlayHistoryRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetIdType(v string) *ListPlayHistoryRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetOrganizationId(v string) *ListPlayHistoryRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListPlayHistoryShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	RequestShrink    *string `json:"Request,omitempty" xml:"Request,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListPlayHistoryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryShrinkRequest) SetDeviceInfoShrink(v string) *ListPlayHistoryShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListPlayHistoryShrinkRequest) SetRequestShrink(v string) *ListPlayHistoryShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *ListPlayHistoryShrinkRequest) SetUserInfoShrink(v string) *ListPlayHistoryShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListPlayHistoryResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    []*ListPlayHistoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPlayHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBody) SetCode(v int32) *ListPlayHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListPlayHistoryResponseBody) SetMessage(v string) *ListPlayHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListPlayHistoryResponseBody) SetResult(v []*ListPlayHistoryResponseBodyResult) *ListPlayHistoryResponseBody {
	s.Result = v
	return s
}

func (s *ListPlayHistoryResponseBody) SetRequestId(v string) *ListPlayHistoryResponseBody {
	s.RequestId = &v
	return s
}

type ListPlayHistoryResponseBodyResult struct {
	Alias       []*string                                   `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	Audition    *bool                                       `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors     []*ListPlayHistoryResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	Category    *string                                     `json:"Category,omitempty" xml:"Category,omitempty"`
	Charge      *bool                                       `json:"Charge,omitempty" xml:"Charge,omitempty"`
	CommCateId  *int64                                      `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *ListPlayHistoryResponseBodyResultCover     `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	HotScore    *float64                                    `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	Id          *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	ItemType    *string                                     `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	Source      *string                                     `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid       *string                                     `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s ListPlayHistoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResult) SetAlias(v []*string) *ListPlayHistoryResponseBodyResult {
	s.Alias = v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetAudition(v bool) *ListPlayHistoryResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetAuthors(v []*ListPlayHistoryResponseBodyResultAuthors) *ListPlayHistoryResponseBodyResult {
	s.Authors = v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCategory(v string) *ListPlayHistoryResponseBodyResult {
	s.Category = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCharge(v bool) *ListPlayHistoryResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCommCateId(v int64) *ListPlayHistoryResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCover(v *ListPlayHistoryResponseBodyResultCover) *ListPlayHistoryResponseBodyResult {
	s.Cover = v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetDescription(v string) *ListPlayHistoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetHotScore(v float64) *ListPlayHistoryResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetId(v int64) *ListPlayHistoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetItemType(v string) *ListPlayHistoryResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetSource(v string) *ListPlayHistoryResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetTitle(v string) *ListPlayHistoryResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetType(v string) *ListPlayHistoryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetValid(v string) *ListPlayHistoryResponseBodyResult {
	s.Valid = &v
	return s
}

type ListPlayHistoryResponseBodyResultAuthors struct {
	AuthorTypes []*string                                      `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Cover       *ListPlayHistoryResponseBodyResultAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Gender      *string                                        `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Id          *int64                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Online      *bool                                          `json:"Online,omitempty" xml:"Online,omitempty"`
	RawId       *string                                        `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source      *string                                        `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string                                        `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListPlayHistoryResponseBodyResultAuthors) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetAuthorTypes(v []*string) *ListPlayHistoryResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetCover(v *ListPlayHistoryResponseBodyResultAuthorsCover) *ListPlayHistoryResponseBodyResultAuthors {
	s.Cover = v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetDescription(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Description = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetGender(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetId(v int64) *ListPlayHistoryResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetOnline(v bool) *ListPlayHistoryResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetRawId(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.RawId = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetSource(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetTitle(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Title = &v
	return s
}

type ListPlayHistoryResponseBodyResultAuthorsCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListPlayHistoryResponseBodyResultAuthorsCover) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResultAuthorsCover) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetCanResize(v bool) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetImg(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Img = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetLarge(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Large = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetMedium(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Medium = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetSmall(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Small = &v
	return s
}

type ListPlayHistoryResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListPlayHistoryResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResultCover) SetCanResize(v bool) *ListPlayHistoryResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetImg(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetLarge(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetMediam(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetMedium(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetSmall(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Small = &v
	return s
}

type ListPlayHistoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPlayHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPlayHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPlayHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponse) SetHeaders(v map[string]*string) *ListPlayHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListPlayHistoryResponse) SetStatusCode(v int32) *ListPlayHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlayHistoryResponse) SetBody(v *ListPlayHistoryResponseBody) *ListPlayHistoryResponse {
	s.Body = v
	return s
}

type ListRecommendContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListRecommendContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentHeaders) GoString() string {
	return s.String()
}

func (s *ListRecommendContentHeaders) SetCommonHeaders(v map[string]*string) *ListRecommendContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRecommendContentHeaders) SetXAcsAligenieAccessToken(v string) *ListRecommendContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListRecommendContentHeaders) SetAuthorization(v string) *ListRecommendContentHeaders {
	s.Authorization = &v
	return s
}

type ListRecommendContentRequest struct {
	DeviceInfo *ListRecommendContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Request    *ListRecommendContentRequestRequest    `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	UserInfo   *ListRecommendContentRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListRecommendContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequest) SetDeviceInfo(v *ListRecommendContentRequestDeviceInfo) *ListRecommendContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListRecommendContentRequest) SetRequest(v *ListRecommendContentRequestRequest) *ListRecommendContentRequest {
	s.Request = v
	return s
}

func (s *ListRecommendContentRequest) SetUserInfo(v *ListRecommendContentRequestUserInfo) *ListRecommendContentRequest {
	s.UserInfo = v
	return s
}

type ListRecommendContentRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRecommendContentRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestDeviceInfo) SetEncodeKey(v string) *ListRecommendContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetEncodeType(v string) *ListRecommendContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetId(v string) *ListRecommendContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetIdType(v string) *ListRecommendContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetOrganizationId(v string) *ListRecommendContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListRecommendContentRequestRequest struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecommendContentRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentRequestRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestRequest) SetCount(v int32) *ListRecommendContentRequestRequest {
	s.Count = &v
	return s
}

func (s *ListRecommendContentRequestRequest) SetType(v string) *ListRecommendContentRequestRequest {
	s.Type = &v
	return s
}

type ListRecommendContentRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRecommendContentRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestUserInfo) SetEncodeKey(v string) *ListRecommendContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetEncodeType(v string) *ListRecommendContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetId(v string) *ListRecommendContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetIdType(v string) *ListRecommendContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetOrganizationId(v string) *ListRecommendContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListRecommendContentShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	RequestShrink    *string `json:"Request,omitempty" xml:"Request,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListRecommendContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentShrinkRequest) SetDeviceInfoShrink(v string) *ListRecommendContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListRecommendContentShrinkRequest) SetRequestShrink(v string) *ListRecommendContentShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *ListRecommendContentShrinkRequest) SetUserInfoShrink(v string) *ListRecommendContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListRecommendContentResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListRecommendContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRecommendContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBody) SetCode(v int32) *ListRecommendContentResponseBody {
	s.Code = &v
	return s
}

func (s *ListRecommendContentResponseBody) SetMessage(v string) *ListRecommendContentResponseBody {
	s.Message = &v
	return s
}

func (s *ListRecommendContentResponseBody) SetRequestId(v string) *ListRecommendContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecommendContentResponseBody) SetResult(v []*ListRecommendContentResponseBodyResult) *ListRecommendContentResponseBody {
	s.Result = v
	return s
}

type ListRecommendContentResponseBodyResult struct {
	Alias       []*string                                        `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	Audition    *bool                                            `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors     []*ListRecommendContentResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	Category    *string                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	Charge      *bool                                            `json:"Charge,omitempty" xml:"Charge,omitempty"`
	CommCateId  *int64                                           `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *ListRecommendContentResponseBodyResultCover     `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	HotScore    *float64                                         `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	Id          *int64                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	ItemType    *string                                          `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	RawId       *string                                          `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source      *string                                          `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid       *string                                          `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s ListRecommendContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResult) SetAlias(v []*string) *ListRecommendContentResponseBodyResult {
	s.Alias = v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetAudition(v bool) *ListRecommendContentResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetAuthors(v []*ListRecommendContentResponseBodyResultAuthors) *ListRecommendContentResponseBodyResult {
	s.Authors = v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCategory(v string) *ListRecommendContentResponseBodyResult {
	s.Category = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCharge(v bool) *ListRecommendContentResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCommCateId(v int64) *ListRecommendContentResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCover(v *ListRecommendContentResponseBodyResultCover) *ListRecommendContentResponseBodyResult {
	s.Cover = v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetDescription(v string) *ListRecommendContentResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetHotScore(v float64) *ListRecommendContentResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetId(v int64) *ListRecommendContentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetItemType(v string) *ListRecommendContentResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetRawId(v string) *ListRecommendContentResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetSource(v string) *ListRecommendContentResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetTitle(v string) *ListRecommendContentResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetType(v string) *ListRecommendContentResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetValid(v string) *ListRecommendContentResponseBodyResult {
	s.Valid = &v
	return s
}

type ListRecommendContentResponseBodyResultAuthors struct {
	AuthorTypes []*string                                           `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Cover       *ListRecommendContentResponseBodyResultAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Gender      *string                                             `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Id          *int64                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Online      *bool                                               `json:"Online,omitempty" xml:"Online,omitempty"`
	RawId       *string                                             `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source      *string                                             `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string                                             `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListRecommendContentResponseBodyResultAuthors) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetAuthorTypes(v []*string) *ListRecommendContentResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetCover(v *ListRecommendContentResponseBodyResultAuthorsCover) *ListRecommendContentResponseBodyResultAuthors {
	s.Cover = v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetDescription(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Description = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetGender(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetId(v int64) *ListRecommendContentResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetOnline(v bool) *ListRecommendContentResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetRawId(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.RawId = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetSource(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetTitle(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Title = &v
	return s
}

type ListRecommendContentResponseBodyResultAuthorsCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListRecommendContentResponseBodyResultAuthorsCover) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentResponseBodyResultAuthorsCover) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetCanResize(v bool) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetImg(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Img = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetLarge(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Large = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetMedium(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Medium = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetSmall(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Small = &v
	return s
}

type ListRecommendContentResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListRecommendContentResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResultCover) SetCanResize(v bool) *ListRecommendContentResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetImg(v string) *ListRecommendContentResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetLarge(v string) *ListRecommendContentResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetMediam(v string) *ListRecommendContentResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetMedium(v string) *ListRecommendContentResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetSmall(v string) *ListRecommendContentResponseBodyResultCover {
	s.Small = &v
	return s
}

type ListRecommendContentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecommendContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecommendContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecommendContentResponse) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponse) SetHeaders(v map[string]*string) *ListRecommendContentResponse {
	s.Headers = v
	return s
}

func (s *ListRecommendContentResponse) SetStatusCode(v int32) *ListRecommendContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecommendContentResponse) SetBody(v *ListRecommendContentResponseBody) *ListRecommendContentResponse {
	s.Body = v
	return s
}

type ListSubHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSubHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubHeaders) GoString() string {
	return s.String()
}

func (s *ListSubHeaders) SetCommonHeaders(v map[string]*string) *ListSubHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubHeaders) SetXAcsAligenieAccessToken(v string) *ListSubHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSubHeaders) SetAuthorization(v string) *ListSubHeaders {
	s.Authorization = &v
	return s
}

type ListSubRequest struct {
	DeviceInfo *ListSubRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Page       *ListSubRequestPage       `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	UserInfo   *ListSubRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListSubRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubRequest) GoString() string {
	return s.String()
}

func (s *ListSubRequest) SetDeviceInfo(v *ListSubRequestDeviceInfo) *ListSubRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListSubRequest) SetPage(v *ListSubRequestPage) *ListSubRequest {
	s.Page = v
	return s
}

func (s *ListSubRequest) SetUserInfo(v *ListSubRequestUserInfo) *ListSubRequest {
	s.UserInfo = v
	return s
}

type ListSubRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSubRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListSubRequestDeviceInfo) SetEncodeKey(v string) *ListSubRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetEncodeType(v string) *ListSubRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetId(v string) *ListSubRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetIdType(v string) *ListSubRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetOrganizationId(v string) *ListSubRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListSubRequestPage struct {
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListSubRequestPage) GoString() string {
	return s.String()
}

func (s *ListSubRequestPage) SetPageNum(v int32) *ListSubRequestPage {
	s.PageNum = &v
	return s
}

func (s *ListSubRequestPage) SetPageSize(v int32) *ListSubRequestPage {
	s.PageSize = &v
	return s
}

type ListSubRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSubRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListSubRequestUserInfo) SetEncodeKey(v string) *ListSubRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubRequestUserInfo) SetEncodeType(v string) *ListSubRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubRequestUserInfo) SetId(v string) *ListSubRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListSubRequestUserInfo) SetIdType(v string) *ListSubRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListSubRequestUserInfo) SetOrganizationId(v string) *ListSubRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListSubShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PageShrink       *string `json:"Page,omitempty" xml:"Page,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListSubShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSubShrinkRequest) SetDeviceInfoShrink(v string) *ListSubShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListSubShrinkRequest) SetPageShrink(v string) *ListSubShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListSubShrinkRequest) SetUserInfoShrink(v string) *ListSubShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListSubResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListSubResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListSubResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubResponseBody) SetCode(v int32) *ListSubResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubResponseBody) SetMessage(v string) *ListSubResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubResponseBody) SetRequestId(v string) *ListSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubResponseBody) SetResult(v *ListSubResponseBodyResult) *ListSubResponseBody {
	s.Result = v
	return s
}

type ListSubResponseBodyResult struct {
	DataList       []*ListSubResponseBodyResultDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	HasNext        *bool                                `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	TotalCount     *int64                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32                               `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListSubResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSubResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubResponseBodyResult) SetDataList(v []*ListSubResponseBodyResultDataList) *ListSubResponseBodyResult {
	s.DataList = v
	return s
}

func (s *ListSubResponseBodyResult) SetHasNext(v bool) *ListSubResponseBodyResult {
	s.HasNext = &v
	return s
}

func (s *ListSubResponseBodyResult) SetTotalCount(v int64) *ListSubResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListSubResponseBodyResult) SetTotalPageCount(v int32) *ListSubResponseBodyResult {
	s.TotalPageCount = &v
	return s
}

type ListSubResponseBodyResultDataList struct {
	AlbumId       *string                                        `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	CoverUrl      *string                                        `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	DailyStudyCnt *int32                                         `json:"DailyStudyCnt,omitempty" xml:"DailyStudyCnt,omitempty"`
	DeviceId      *string                                        `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Id            *int64                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	PlayMode      *string                                        `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	ScheduleInfo  *ListSubResponseBodyResultDataListScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	Title         *string                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	UserId        *int64                                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSubResponseBodyResultDataList) String() string {
	return tea.Prettify(s)
}

func (s ListSubResponseBodyResultDataList) GoString() string {
	return s.String()
}

func (s *ListSubResponseBodyResultDataList) SetAlbumId(v string) *ListSubResponseBodyResultDataList {
	s.AlbumId = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetCoverUrl(v string) *ListSubResponseBodyResultDataList {
	s.CoverUrl = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetDailyStudyCnt(v int32) *ListSubResponseBodyResultDataList {
	s.DailyStudyCnt = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetDeviceId(v string) *ListSubResponseBodyResultDataList {
	s.DeviceId = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetId(v int64) *ListSubResponseBodyResultDataList {
	s.Id = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetPlayMode(v string) *ListSubResponseBodyResultDataList {
	s.PlayMode = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetScheduleInfo(v *ListSubResponseBodyResultDataListScheduleInfo) *ListSubResponseBodyResultDataList {
	s.ScheduleInfo = v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetTitle(v string) *ListSubResponseBodyResultDataList {
	s.Title = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetUserId(v int64) *ListSubResponseBodyResultDataList {
	s.UserId = &v
	return s
}

type ListSubResponseBodyResultDataListScheduleInfo struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s ListSubResponseBodyResultDataListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSubResponseBodyResultDataListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) SetDaysOfWeek(v []*int32) *ListSubResponseBodyResultDataListScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) SetHour(v int32) *ListSubResponseBodyResultDataListScheduleInfo {
	s.Hour = &v
	return s
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) SetMinute(v int32) *ListSubResponseBodyResultDataListScheduleInfo {
	s.Minute = &v
	return s
}

type ListSubResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubResponse) GoString() string {
	return s.String()
}

func (s *ListSubResponse) SetHeaders(v map[string]*string) *ListSubResponse {
	s.Headers = v
	return s
}

func (s *ListSubResponse) SetStatusCode(v int32) *ListSubResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubResponse) SetBody(v *ListSubResponseBody) *ListSubResponse {
	s.Body = v
	return s
}

type ListSubAlbumHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSubAlbumHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumHeaders) GoString() string {
	return s.String()
}

func (s *ListSubAlbumHeaders) SetCommonHeaders(v map[string]*string) *ListSubAlbumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubAlbumHeaders) SetXAcsAligenieAccessToken(v string) *ListSubAlbumHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSubAlbumHeaders) SetAuthorization(v string) *ListSubAlbumHeaders {
	s.Authorization = &v
	return s
}

type ListSubAlbumRequest struct {
	DeviceInfo                    *ListSubAlbumRequestDeviceInfo                    `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	QuerySubscriptionAlbumRequest *ListSubAlbumRequestQuerySubscriptionAlbumRequest `json:"QuerySubscriptionAlbumRequest,omitempty" xml:"QuerySubscriptionAlbumRequest,omitempty" type:"Struct"`
	UserInfo                      *ListSubAlbumRequestUserInfo                      `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListSubAlbumRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumRequest) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequest) SetDeviceInfo(v *ListSubAlbumRequestDeviceInfo) *ListSubAlbumRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListSubAlbumRequest) SetQuerySubscriptionAlbumRequest(v *ListSubAlbumRequestQuerySubscriptionAlbumRequest) *ListSubAlbumRequest {
	s.QuerySubscriptionAlbumRequest = v
	return s
}

func (s *ListSubAlbumRequest) SetUserInfo(v *ListSubAlbumRequestUserInfo) *ListSubAlbumRequest {
	s.UserInfo = v
	return s
}

type ListSubAlbumRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubAlbumRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestDeviceInfo) SetEncodeKey(v string) *ListSubAlbumRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetEncodeType(v string) *ListSubAlbumRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetId(v string) *ListSubAlbumRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetIdType(v string) *ListSubAlbumRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetOrganizationId(v string) *ListSubAlbumRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListSubAlbumRequestQuerySubscriptionAlbumRequest struct {
	AlbumId    *string                                               `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	CategoryId *int32                                                `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Page       *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	Title      *string                                               `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequest) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetAlbumId(v string) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.AlbumId = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetCategoryId(v int32) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.CategoryId = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetPage(v *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.Page = v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetTitle(v string) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.Title = &v
	return s
}

type ListSubAlbumRequestQuerySubscriptionAlbumRequestPage struct {
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) SetPageNum(v int32) *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage {
	s.PageNum = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) SetPageSize(v int32) *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage {
	s.PageSize = &v
	return s
}

type ListSubAlbumRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubAlbumRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestUserInfo) SetEncodeKey(v string) *ListSubAlbumRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetEncodeType(v string) *ListSubAlbumRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetId(v string) *ListSubAlbumRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetIdType(v string) *ListSubAlbumRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetOrganizationId(v string) *ListSubAlbumRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListSubAlbumShrinkRequest struct {
	DeviceInfoShrink                    *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	QuerySubscriptionAlbumRequestShrink *string `json:"QuerySubscriptionAlbumRequest,omitempty" xml:"QuerySubscriptionAlbumRequest,omitempty"`
	UserInfoShrink                      *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListSubAlbumShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSubAlbumShrinkRequest) SetDeviceInfoShrink(v string) *ListSubAlbumShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListSubAlbumShrinkRequest) SetQuerySubscriptionAlbumRequestShrink(v string) *ListSubAlbumShrinkRequest {
	s.QuerySubscriptionAlbumRequestShrink = &v
	return s
}

func (s *ListSubAlbumShrinkRequest) SetUserInfoShrink(v string) *ListSubAlbumShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListSubAlbumResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListSubAlbumResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListSubAlbumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBody) SetCode(v int32) *ListSubAlbumResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubAlbumResponseBody) SetMessage(v string) *ListSubAlbumResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubAlbumResponseBody) SetRequestId(v string) *ListSubAlbumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubAlbumResponseBody) SetResult(v *ListSubAlbumResponseBodyResult) *ListSubAlbumResponseBody {
	s.Result = v
	return s
}

type ListSubAlbumResponseBodyResult struct {
	DataList       []*ListSubAlbumResponseBodyResultDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	HasNext        *bool                                     `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	TotalCount     *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32                                    `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListSubAlbumResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBodyResult) SetDataList(v []*ListSubAlbumResponseBodyResultDataList) *ListSubAlbumResponseBodyResult {
	s.DataList = v
	return s
}

func (s *ListSubAlbumResponseBodyResult) SetHasNext(v bool) *ListSubAlbumResponseBodyResult {
	s.HasNext = &v
	return s
}

func (s *ListSubAlbumResponseBodyResult) SetTotalCount(v int32) *ListSubAlbumResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListSubAlbumResponseBodyResult) SetTotalPageCount(v int32) *ListSubAlbumResponseBodyResult {
	s.TotalPageCount = &v
	return s
}

type ListSubAlbumResponseBodyResultDataList struct {
	AlbumId      *string                                             `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	CategoryId   *int32                                              `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CoverUrl     *string                                             `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Id           *int64                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	IsAdded      *bool                                               `json:"IsAdded,omitempty" xml:"IsAdded,omitempty"`
	ScheduleInfo *ListSubAlbumResponseBodyResultDataListScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	Sequence     *int64                                              `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	Title        *string                                             `json:"Title,omitempty" xml:"Title,omitempty"`
	TotalEpisode *int32                                              `json:"TotalEpisode,omitempty" xml:"TotalEpisode,omitempty"`
}

func (s ListSubAlbumResponseBodyResultDataList) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumResponseBodyResultDataList) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBodyResultDataList) SetAlbumId(v string) *ListSubAlbumResponseBodyResultDataList {
	s.AlbumId = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetCategoryId(v int32) *ListSubAlbumResponseBodyResultDataList {
	s.CategoryId = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetCoverUrl(v string) *ListSubAlbumResponseBodyResultDataList {
	s.CoverUrl = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetId(v int64) *ListSubAlbumResponseBodyResultDataList {
	s.Id = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetIsAdded(v bool) *ListSubAlbumResponseBodyResultDataList {
	s.IsAdded = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetScheduleInfo(v *ListSubAlbumResponseBodyResultDataListScheduleInfo) *ListSubAlbumResponseBodyResultDataList {
	s.ScheduleInfo = v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetSequence(v int64) *ListSubAlbumResponseBodyResultDataList {
	s.Sequence = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetTitle(v string) *ListSubAlbumResponseBodyResultDataList {
	s.Title = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetTotalEpisode(v int32) *ListSubAlbumResponseBodyResultDataList {
	s.TotalEpisode = &v
	return s
}

type ListSubAlbumResponseBodyResultDataListScheduleInfo struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
	ScheduleId *int64   `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
}

func (s ListSubAlbumResponseBodyResultDataListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumResponseBodyResultDataListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetDaysOfWeek(v []*int32) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetHour(v int32) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.Hour = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetMinute(v int32) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.Minute = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetScheduleId(v int64) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.ScheduleId = &v
	return s
}

type ListSubAlbumResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubAlbumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubAlbumResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubAlbumResponse) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponse) SetHeaders(v map[string]*string) *ListSubAlbumResponse {
	s.Headers = v
	return s
}

func (s *ListSubAlbumResponse) SetStatusCode(v int32) *ListSubAlbumResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubAlbumResponse) SetBody(v *ListSubAlbumResponseBody) *ListSubAlbumResponse {
	s.Body = v
	return s
}

type ListSubscriptionAlbumCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSubscriptionAlbumCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListSubscriptionAlbumCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubscriptionAlbumCategoryHeaders) SetXAcsAligenieAccessToken(v string) *ListSubscriptionAlbumCategoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryHeaders) SetAuthorization(v string) *ListSubscriptionAlbumCategoryHeaders {
	s.Authorization = &v
	return s
}

type ListSubscriptionAlbumCategoryRequest struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s ListSubscriptionAlbumCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryRequest) SetCategoryName(v string) *ListSubscriptionAlbumCategoryRequest {
	s.CategoryName = &v
	return s
}

type ListSubscriptionAlbumCategoryResponseBody struct {
	Code      *int32                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListSubscriptionAlbumCategoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSubscriptionAlbumCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetCode(v int32) *ListSubscriptionAlbumCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetMessage(v string) *ListSubscriptionAlbumCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetRequestId(v string) *ListSubscriptionAlbumCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetResult(v []*ListSubscriptionAlbumCategoryResponseBodyResult) *ListSubscriptionAlbumCategoryResponseBody {
	s.Result = v
	return s
}

type ListSubscriptionAlbumCategoryResponseBodyResult struct {
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s ListSubscriptionAlbumCategoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryResponseBodyResult) SetCategoryId(v string) *ListSubscriptionAlbumCategoryResponseBodyResult {
	s.CategoryId = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBodyResult) SetCategoryName(v string) *ListSubscriptionAlbumCategoryResponseBodyResult {
	s.CategoryName = &v
	return s
}

type ListSubscriptionAlbumCategoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubscriptionAlbumCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubscriptionAlbumCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryResponse) SetHeaders(v map[string]*string) *ListSubscriptionAlbumCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponse) SetStatusCode(v int32) *ListSubscriptionAlbumCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponse) SetBody(v *ListSubscriptionAlbumCategoryResponseBody) *ListSubscriptionAlbumCategoryResponse {
	s.Body = v
	return s
}

type ListUserMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListUserMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUserMessageHeaders) GoString() string {
	return s.String()
}

func (s *ListUserMessageHeaders) SetCommonHeaders(v map[string]*string) *ListUserMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserMessageHeaders) SetXAcsAligenieAccessToken(v string) *ListUserMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListUserMessageHeaders) SetAuthorization(v string) *ListUserMessageHeaders {
	s.Authorization = &v
	return s
}

type ListUserMessageRequest struct {
	BeforeTime *string                         `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	UserInfo   *ListUserMessageRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	Limit      *int32                          `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s ListUserMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserMessageRequest) GoString() string {
	return s.String()
}

func (s *ListUserMessageRequest) SetBeforeTime(v string) *ListUserMessageRequest {
	s.BeforeTime = &v
	return s
}

func (s *ListUserMessageRequest) SetUserInfo(v *ListUserMessageRequestUserInfo) *ListUserMessageRequest {
	s.UserInfo = v
	return s
}

func (s *ListUserMessageRequest) SetLimit(v int32) *ListUserMessageRequest {
	s.Limit = &v
	return s
}

type ListUserMessageRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListUserMessageRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUserMessageRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListUserMessageRequestUserInfo) SetEncodeKey(v string) *ListUserMessageRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetEncodeType(v string) *ListUserMessageRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetId(v string) *ListUserMessageRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetIdType(v string) *ListUserMessageRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetOrganizationId(v string) *ListUserMessageRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListUserMessageShrinkRequest struct {
	BeforeTime     *string `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	Limit          *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s ListUserMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserMessageShrinkRequest) SetBeforeTime(v string) *ListUserMessageShrinkRequest {
	s.BeforeTime = &v
	return s
}

func (s *ListUserMessageShrinkRequest) SetUserInfoShrink(v string) *ListUserMessageShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListUserMessageShrinkRequest) SetLimit(v int32) *ListUserMessageShrinkRequest {
	s.Limit = &v
	return s
}

type ListUserMessageResponseBody struct {
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Result  []*ListUserMessageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListUserMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserMessageResponseBody) SetCode(v string) *ListUserMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserMessageResponseBody) SetMessage(v string) *ListUserMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserMessageResponseBody) SetResult(v []*ListUserMessageResponseBodyResult) *ListUserMessageResponseBody {
	s.Result = v
	return s
}

type ListUserMessageResponseBodyResult struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	GmtCreate  *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Pic        *string `json:"Pic,omitempty" xml:"Pic,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceUuid *string `json:"SourceUuid,omitempty" xml:"SourceUuid,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListUserMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListUserMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserMessageResponseBodyResult) SetContent(v string) *ListUserMessageResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetDeviceName(v string) *ListUserMessageResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetGmtCreate(v string) *ListUserMessageResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetId(v string) *ListUserMessageResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetPic(v string) *ListUserMessageResponseBodyResult {
	s.Pic = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetSource(v string) *ListUserMessageResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetSourceUuid(v string) *ListUserMessageResponseBodyResult {
	s.SourceUuid = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetStatus(v int32) *ListUserMessageResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetType(v string) *ListUserMessageResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetUrl(v string) *ListUserMessageResponseBodyResult {
	s.Url = &v
	return s
}

type ListUserMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserMessageResponse) GoString() string {
	return s.String()
}

func (s *ListUserMessageResponse) SetHeaders(v map[string]*string) *ListUserMessageResponse {
	s.Headers = v
	return s
}

func (s *ListUserMessageResponse) SetStatusCode(v int32) *ListUserMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserMessageResponse) SetBody(v *ListUserMessageResponseBody) *ListUserMessageResponse {
	s.Body = v
	return s
}

type PlayAndPauseControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PlayAndPauseControlHeaders) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlHeaders) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlHeaders) SetCommonHeaders(v map[string]*string) *PlayAndPauseControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PlayAndPauseControlHeaders) SetXAcsAligenieAccessToken(v string) *PlayAndPauseControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PlayAndPauseControlHeaders) SetAuthorization(v string) *PlayAndPauseControlHeaders {
	s.Authorization = &v
	return s
}

type PlayAndPauseControlRequest struct {
	DeviceInfo                   *PlayAndPauseControlRequestDeviceInfo                   `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenPlayAndPauseControlParam *PlayAndPauseControlRequestOpenPlayAndPauseControlParam `json:"OpenPlayAndPauseControlParam,omitempty" xml:"OpenPlayAndPauseControlParam,omitempty" type:"Struct"`
	UserInfo                     *PlayAndPauseControlRequestUserInfo                     `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PlayAndPauseControlRequest) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlRequest) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequest) SetDeviceInfo(v *PlayAndPauseControlRequestDeviceInfo) *PlayAndPauseControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *PlayAndPauseControlRequest) SetOpenPlayAndPauseControlParam(v *PlayAndPauseControlRequestOpenPlayAndPauseControlParam) *PlayAndPauseControlRequest {
	s.OpenPlayAndPauseControlParam = v
	return s
}

func (s *PlayAndPauseControlRequest) SetUserInfo(v *PlayAndPauseControlRequestUserInfo) *PlayAndPauseControlRequest {
	s.UserInfo = v
	return s
}

type PlayAndPauseControlRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PlayAndPauseControlRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetEncodeKey(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetEncodeType(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetId(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetIdType(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetOrganizationId(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type PlayAndPauseControlRequestOpenPlayAndPauseControlParam struct {
	OpenPlayAndPauseCommand *string `json:"OpenPlayAndPauseCommand,omitempty" xml:"OpenPlayAndPauseCommand,omitempty"`
}

func (s PlayAndPauseControlRequestOpenPlayAndPauseControlParam) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlRequestOpenPlayAndPauseControlParam) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequestOpenPlayAndPauseControlParam) SetOpenPlayAndPauseCommand(v string) *PlayAndPauseControlRequestOpenPlayAndPauseControlParam {
	s.OpenPlayAndPauseCommand = &v
	return s
}

type PlayAndPauseControlRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PlayAndPauseControlRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequestUserInfo) SetEncodeKey(v string) *PlayAndPauseControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetEncodeType(v string) *PlayAndPauseControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetId(v string) *PlayAndPauseControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetIdType(v string) *PlayAndPauseControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetOrganizationId(v string) *PlayAndPauseControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type PlayAndPauseControlShrinkRequest struct {
	DeviceInfoShrink                   *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenPlayAndPauseControlParamShrink *string `json:"OpenPlayAndPauseControlParam,omitempty" xml:"OpenPlayAndPauseControlParam,omitempty"`
	UserInfoShrink                     *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PlayAndPauseControlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlShrinkRequest) SetDeviceInfoShrink(v string) *PlayAndPauseControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PlayAndPauseControlShrinkRequest) SetOpenPlayAndPauseControlParamShrink(v string) *PlayAndPauseControlShrinkRequest {
	s.OpenPlayAndPauseControlParamShrink = &v
	return s
}

func (s *PlayAndPauseControlShrinkRequest) SetUserInfoShrink(v string) *PlayAndPauseControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type PlayAndPauseControlResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PlayAndPauseControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlResponseBody) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlResponseBody) SetCode(v int32) *PlayAndPauseControlResponseBody {
	s.Code = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetMessage(v string) *PlayAndPauseControlResponseBody {
	s.Message = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetRequestId(v string) *PlayAndPauseControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetResult(v bool) *PlayAndPauseControlResponseBody {
	s.Result = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetSuccess(v string) *PlayAndPauseControlResponseBody {
	s.Success = &v
	return s
}

type PlayAndPauseControlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PlayAndPauseControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PlayAndPauseControlResponse) String() string {
	return tea.Prettify(s)
}

func (s PlayAndPauseControlResponse) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlResponse) SetHeaders(v map[string]*string) *PlayAndPauseControlResponse {
	s.Headers = v
	return s
}

func (s *PlayAndPauseControlResponse) SetStatusCode(v int32) *PlayAndPauseControlResponse {
	s.StatusCode = &v
	return s
}

func (s *PlayAndPauseControlResponse) SetBody(v *PlayAndPauseControlResponseBody) *PlayAndPauseControlResponse {
	s.Body = v
	return s
}

type PlayModeControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PlayModeControlHeaders) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlHeaders) GoString() string {
	return s.String()
}

func (s *PlayModeControlHeaders) SetCommonHeaders(v map[string]*string) *PlayModeControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PlayModeControlHeaders) SetXAcsAligenieAccessToken(v string) *PlayModeControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PlayModeControlHeaders) SetAuthorization(v string) *PlayModeControlHeaders {
	s.Authorization = &v
	return s
}

type PlayModeControlRequest struct {
	DeviceInfo                 *PlayModeControlRequestDeviceInfo                 `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenPlayModeControlRequest *PlayModeControlRequestOpenPlayModeControlRequest `json:"OpenPlayModeControlRequest,omitempty" xml:"OpenPlayModeControlRequest,omitempty" type:"Struct"`
	UserInfo                   *PlayModeControlRequestUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PlayModeControlRequest) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequest) SetDeviceInfo(v *PlayModeControlRequestDeviceInfo) *PlayModeControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *PlayModeControlRequest) SetOpenPlayModeControlRequest(v *PlayModeControlRequestOpenPlayModeControlRequest) *PlayModeControlRequest {
	s.OpenPlayModeControlRequest = v
	return s
}

func (s *PlayModeControlRequest) SetUserInfo(v *PlayModeControlRequestUserInfo) *PlayModeControlRequest {
	s.UserInfo = v
	return s
}

type PlayModeControlRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PlayModeControlRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestDeviceInfo) SetEncodeKey(v string) *PlayModeControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetEncodeType(v string) *PlayModeControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetId(v string) *PlayModeControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetIdType(v string) *PlayModeControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetOrganizationId(v string) *PlayModeControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type PlayModeControlRequestOpenPlayModeControlRequest struct {
	OpenPlayMode *string `json:"OpenPlayMode,omitempty" xml:"OpenPlayMode,omitempty"`
}

func (s PlayModeControlRequestOpenPlayModeControlRequest) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlRequestOpenPlayModeControlRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestOpenPlayModeControlRequest) SetOpenPlayMode(v string) *PlayModeControlRequestOpenPlayModeControlRequest {
	s.OpenPlayMode = &v
	return s
}

type PlayModeControlRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PlayModeControlRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestUserInfo) SetEncodeKey(v string) *PlayModeControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetEncodeType(v string) *PlayModeControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetId(v string) *PlayModeControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetIdType(v string) *PlayModeControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetOrganizationId(v string) *PlayModeControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type PlayModeControlShrinkRequest struct {
	DeviceInfoShrink                 *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenPlayModeControlRequestShrink *string `json:"OpenPlayModeControlRequest,omitempty" xml:"OpenPlayModeControlRequest,omitempty"`
	UserInfoShrink                   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PlayModeControlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlShrinkRequest) SetDeviceInfoShrink(v string) *PlayModeControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PlayModeControlShrinkRequest) SetOpenPlayModeControlRequestShrink(v string) *PlayModeControlShrinkRequest {
	s.OpenPlayModeControlRequestShrink = &v
	return s
}

func (s *PlayModeControlShrinkRequest) SetUserInfoShrink(v string) *PlayModeControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type PlayModeControlResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *PlayModeControlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PlayModeControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlResponseBody) GoString() string {
	return s.String()
}

func (s *PlayModeControlResponseBody) SetCode(v int32) *PlayModeControlResponseBody {
	s.Code = &v
	return s
}

func (s *PlayModeControlResponseBody) SetMessage(v string) *PlayModeControlResponseBody {
	s.Message = &v
	return s
}

func (s *PlayModeControlResponseBody) SetRequestId(v string) *PlayModeControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *PlayModeControlResponseBody) SetResult(v *PlayModeControlResponseBodyResult) *PlayModeControlResponseBody {
	s.Result = v
	return s
}

func (s *PlayModeControlResponseBody) SetSuccess(v string) *PlayModeControlResponseBody {
	s.Success = &v
	return s
}

type PlayModeControlResponseBodyResult struct {
	OpenPlayMode *string `json:"OpenPlayMode,omitempty" xml:"OpenPlayMode,omitempty"`
}

func (s PlayModeControlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PlayModeControlResponseBodyResult) SetOpenPlayMode(v string) *PlayModeControlResponseBodyResult {
	s.OpenPlayMode = &v
	return s
}

type PlayModeControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PlayModeControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PlayModeControlResponse) String() string {
	return tea.Prettify(s)
}

func (s PlayModeControlResponse) GoString() string {
	return s.String()
}

func (s *PlayModeControlResponse) SetHeaders(v map[string]*string) *PlayModeControlResponse {
	s.Headers = v
	return s
}

func (s *PlayModeControlResponse) SetStatusCode(v int32) *PlayModeControlResponse {
	s.StatusCode = &v
	return s
}

func (s *PlayModeControlResponse) SetBody(v *PlayModeControlResponseBody) *PlayModeControlResponse {
	s.Body = v
	return s
}

type PreviousAndNextControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PreviousAndNextControlHeaders) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlHeaders) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlHeaders) SetCommonHeaders(v map[string]*string) *PreviousAndNextControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PreviousAndNextControlHeaders) SetXAcsAligenieAccessToken(v string) *PreviousAndNextControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PreviousAndNextControlHeaders) SetAuthorization(v string) *PreviousAndNextControlHeaders {
	s.Authorization = &v
	return s
}

type PreviousAndNextControlRequest struct {
	DeviceInfo                    *PreviousAndNextControlRequestDeviceInfo                    `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenControlPlayingListRequest *PreviousAndNextControlRequestOpenControlPlayingListRequest `json:"OpenControlPlayingListRequest,omitempty" xml:"OpenControlPlayingListRequest,omitempty" type:"Struct"`
	UserInfo                      *PreviousAndNextControlRequestUserInfo                      `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PreviousAndNextControlRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlRequest) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequest) SetDeviceInfo(v *PreviousAndNextControlRequestDeviceInfo) *PreviousAndNextControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *PreviousAndNextControlRequest) SetOpenControlPlayingListRequest(v *PreviousAndNextControlRequestOpenControlPlayingListRequest) *PreviousAndNextControlRequest {
	s.OpenControlPlayingListRequest = v
	return s
}

func (s *PreviousAndNextControlRequest) SetUserInfo(v *PreviousAndNextControlRequestUserInfo) *PreviousAndNextControlRequest {
	s.UserInfo = v
	return s
}

type PreviousAndNextControlRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PreviousAndNextControlRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetEncodeKey(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetEncodeType(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetId(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetIdType(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetOrganizationId(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type PreviousAndNextControlRequestOpenControlPlayingListRequest struct {
	Cmd          *string                `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	ExtendInfo   map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	IsFromDevice *bool                  `json:"IsFromDevice,omitempty" xml:"IsFromDevice,omitempty"`
}

func (s PreviousAndNextControlRequestOpenControlPlayingListRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlRequestOpenControlPlayingListRequest) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) SetCmd(v string) *PreviousAndNextControlRequestOpenControlPlayingListRequest {
	s.Cmd = &v
	return s
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) SetExtendInfo(v map[string]interface{}) *PreviousAndNextControlRequestOpenControlPlayingListRequest {
	s.ExtendInfo = v
	return s
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) SetIsFromDevice(v bool) *PreviousAndNextControlRequestOpenControlPlayingListRequest {
	s.IsFromDevice = &v
	return s
}

type PreviousAndNextControlRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PreviousAndNextControlRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequestUserInfo) SetEncodeKey(v string) *PreviousAndNextControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetEncodeType(v string) *PreviousAndNextControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetId(v string) *PreviousAndNextControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetIdType(v string) *PreviousAndNextControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetOrganizationId(v string) *PreviousAndNextControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type PreviousAndNextControlShrinkRequest struct {
	DeviceInfoShrink                    *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenControlPlayingListRequestShrink *string `json:"OpenControlPlayingListRequest,omitempty" xml:"OpenControlPlayingListRequest,omitempty"`
	UserInfoShrink                      *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PreviousAndNextControlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlShrinkRequest) SetDeviceInfoShrink(v string) *PreviousAndNextControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PreviousAndNextControlShrinkRequest) SetOpenControlPlayingListRequestShrink(v string) *PreviousAndNextControlShrinkRequest {
	s.OpenControlPlayingListRequestShrink = &v
	return s
}

func (s *PreviousAndNextControlShrinkRequest) SetUserInfoShrink(v string) *PreviousAndNextControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type PreviousAndNextControlResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *PreviousAndNextControlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PreviousAndNextControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlResponseBody) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponseBody) SetCode(v int32) *PreviousAndNextControlResponseBody {
	s.Code = &v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetMessage(v string) *PreviousAndNextControlResponseBody {
	s.Message = &v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetRequestId(v string) *PreviousAndNextControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetResult(v *PreviousAndNextControlResponseBodyResult) *PreviousAndNextControlResponseBody {
	s.Result = v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetSuccess(v string) *PreviousAndNextControlResponseBody {
	s.Success = &v
	return s
}

type PreviousAndNextControlResponseBodyResult struct {
	AlbumName        *string                                        `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	AlbumRawId       *string                                        `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	AudioLength      *int32                                         `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	Copyright        *int32                                         `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover            *PreviousAndNextControlResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	DefaultPlayOrder *int32                                         `json:"DefaultPlayOrder,omitempty" xml:"DefaultPlayOrder,omitempty"`
	ItemUrl          *string                                        `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Liked            *bool                                          `json:"Liked,omitempty" xml:"Liked,omitempty"`
	LyricUrl         *string                                        `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	PlayMode         *string                                        `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	Pos              *int32                                         `json:"Pos,omitempty" xml:"Pos,omitempty"`
	Progress         *int32                                         `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RawId            *string                                        `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singer           *string                                        `json:"Singer,omitempty" xml:"Singer,omitempty"`
	Source           *string                                        `json:"Source,omitempty" xml:"Source,omitempty"`
	Title            *string                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid            *string                                        `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s PreviousAndNextControlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponseBodyResult) SetAlbumName(v string) *PreviousAndNextControlResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetAlbumRawId(v string) *PreviousAndNextControlResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetAudioLength(v int32) *PreviousAndNextControlResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetCopyright(v int32) *PreviousAndNextControlResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetCover(v *PreviousAndNextControlResponseBodyResultCover) *PreviousAndNextControlResponseBodyResult {
	s.Cover = v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetDefaultPlayOrder(v int32) *PreviousAndNextControlResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetItemUrl(v string) *PreviousAndNextControlResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetLiked(v bool) *PreviousAndNextControlResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetLyricUrl(v string) *PreviousAndNextControlResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetPlayMode(v string) *PreviousAndNextControlResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetPos(v int32) *PreviousAndNextControlResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetProgress(v int32) *PreviousAndNextControlResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetRawId(v string) *PreviousAndNextControlResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetSinger(v string) *PreviousAndNextControlResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetSource(v string) *PreviousAndNextControlResponseBodyResult {
	s.Source = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetTitle(v string) *PreviousAndNextControlResponseBodyResult {
	s.Title = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetType(v string) *PreviousAndNextControlResponseBodyResult {
	s.Type = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetValid(v string) *PreviousAndNextControlResponseBodyResult {
	s.Valid = &v
	return s
}

type PreviousAndNextControlResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s PreviousAndNextControlResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetCanResize(v bool) *PreviousAndNextControlResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetImg(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetLarge(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetMediam(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetMedium(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetSmall(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Small = &v
	return s
}

type PreviousAndNextControlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PreviousAndNextControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreviousAndNextControlResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviousAndNextControlResponse) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponse) SetHeaders(v map[string]*string) *PreviousAndNextControlResponse {
	s.Headers = v
	return s
}

func (s *PreviousAndNextControlResponse) SetStatusCode(v int32) *PreviousAndNextControlResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviousAndNextControlResponse) SetBody(v *PreviousAndNextControlResponseBody) *PreviousAndNextControlResponse {
	s.Body = v
	return s
}

type ProgressControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ProgressControlHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlHeaders) GoString() string {
	return s.String()
}

func (s *ProgressControlHeaders) SetCommonHeaders(v map[string]*string) *ProgressControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProgressControlHeaders) SetXAcsAligenieAccessToken(v string) *ProgressControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ProgressControlHeaders) SetAuthorization(v string) *ProgressControlHeaders {
	s.Authorization = &v
	return s
}

type ProgressControlRequest struct {
	DeviceInfo                 *ProgressControlRequestDeviceInfo                 `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	OpenProgressControlRequest *ProgressControlRequestOpenProgressControlRequest `json:"OpenProgressControlRequest,omitempty" xml:"OpenProgressControlRequest,omitempty" type:"Struct"`
	UserInfo                   *ProgressControlRequestUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ProgressControlRequest) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlRequest) GoString() string {
	return s.String()
}

func (s *ProgressControlRequest) SetDeviceInfo(v *ProgressControlRequestDeviceInfo) *ProgressControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *ProgressControlRequest) SetOpenProgressControlRequest(v *ProgressControlRequestOpenProgressControlRequest) *ProgressControlRequest {
	s.OpenProgressControlRequest = v
	return s
}

func (s *ProgressControlRequest) SetUserInfo(v *ProgressControlRequestUserInfo) *ProgressControlRequest {
	s.UserInfo = v
	return s
}

type ProgressControlRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ProgressControlRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ProgressControlRequestDeviceInfo) SetEncodeKey(v string) *ProgressControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetEncodeType(v string) *ProgressControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetId(v string) *ProgressControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetIdType(v string) *ProgressControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetOrganizationId(v string) *ProgressControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ProgressControlRequestOpenProgressControlRequest struct {
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	Progress   *int32                 `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s ProgressControlRequestOpenProgressControlRequest) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlRequestOpenProgressControlRequest) GoString() string {
	return s.String()
}

func (s *ProgressControlRequestOpenProgressControlRequest) SetExtendInfo(v map[string]interface{}) *ProgressControlRequestOpenProgressControlRequest {
	s.ExtendInfo = v
	return s
}

func (s *ProgressControlRequestOpenProgressControlRequest) SetProgress(v int32) *ProgressControlRequestOpenProgressControlRequest {
	s.Progress = &v
	return s
}

type ProgressControlRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ProgressControlRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ProgressControlRequestUserInfo) SetEncodeKey(v string) *ProgressControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetEncodeType(v string) *ProgressControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetId(v string) *ProgressControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetIdType(v string) *ProgressControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetOrganizationId(v string) *ProgressControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ProgressControlShrinkRequest struct {
	DeviceInfoShrink                 *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	OpenProgressControlRequestShrink *string `json:"OpenProgressControlRequest,omitempty" xml:"OpenProgressControlRequest,omitempty"`
	UserInfoShrink                   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ProgressControlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *ProgressControlShrinkRequest) SetDeviceInfoShrink(v string) *ProgressControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ProgressControlShrinkRequest) SetOpenProgressControlRequestShrink(v string) *ProgressControlShrinkRequest {
	s.OpenProgressControlRequestShrink = &v
	return s
}

func (s *ProgressControlShrinkRequest) SetUserInfoShrink(v string) *ProgressControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ProgressControlResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ProgressControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlResponseBody) GoString() string {
	return s.String()
}

func (s *ProgressControlResponseBody) SetCode(v int32) *ProgressControlResponseBody {
	s.Code = &v
	return s
}

func (s *ProgressControlResponseBody) SetMessage(v string) *ProgressControlResponseBody {
	s.Message = &v
	return s
}

func (s *ProgressControlResponseBody) SetRequestId(v string) *ProgressControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProgressControlResponseBody) SetResult(v bool) *ProgressControlResponseBody {
	s.Result = &v
	return s
}

func (s *ProgressControlResponseBody) SetSuccess(v string) *ProgressControlResponseBody {
	s.Success = &v
	return s
}

type ProgressControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProgressControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProgressControlResponse) String() string {
	return tea.Prettify(s)
}

func (s ProgressControlResponse) GoString() string {
	return s.String()
}

func (s *ProgressControlResponse) SetHeaders(v map[string]*string) *ProgressControlResponse {
	s.Headers = v
	return s
}

func (s *ProgressControlResponse) SetStatusCode(v int32) *ProgressControlResponse {
	s.StatusCode = &v
	return s
}

func (s *ProgressControlResponse) SetBody(v *ProgressControlResponseBody) *ProgressControlResponse {
	s.Body = v
	return s
}

type QueryMusicTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryMusicTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryMusicTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMusicTypeHeaders) SetXAcsAligenieAccessToken(v string) *QueryMusicTypeHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryMusicTypeHeaders) SetAuthorization(v string) *QueryMusicTypeHeaders {
	s.Authorization = &v
	return s
}

type QueryMusicTypeRequest struct {
	DeviceInfo *QueryMusicTypeRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *QueryMusicTypeRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *QueryMusicTypeRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s QueryMusicTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequest) SetDeviceInfo(v *QueryMusicTypeRequestDeviceInfo) *QueryMusicTypeRequest {
	s.DeviceInfo = v
	return s
}

func (s *QueryMusicTypeRequest) SetPayload(v *QueryMusicTypeRequestPayload) *QueryMusicTypeRequest {
	s.Payload = v
	return s
}

func (s *QueryMusicTypeRequest) SetUserInfo(v *QueryMusicTypeRequestUserInfo) *QueryMusicTypeRequest {
	s.UserInfo = v
	return s
}

type QueryMusicTypeRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s QueryMusicTypeRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestDeviceInfo) SetEncodeKey(v string) *QueryMusicTypeRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetEncodeType(v string) *QueryMusicTypeRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetId(v string) *QueryMusicTypeRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetIdType(v string) *QueryMusicTypeRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetOrganizationId(v string) *QueryMusicTypeRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type QueryMusicTypeRequestPayload struct {
}

func (s QueryMusicTypeRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeRequestPayload) GoString() string {
	return s.String()
}

type QueryMusicTypeRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s QueryMusicTypeRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeRequestUserInfo) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestUserInfo) SetEncodeKey(v string) *QueryMusicTypeRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetEncodeType(v string) *QueryMusicTypeRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetId(v string) *QueryMusicTypeRequestUserInfo {
	s.Id = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetIdType(v string) *QueryMusicTypeRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetOrganizationId(v string) *QueryMusicTypeRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type QueryMusicTypeShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s QueryMusicTypeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeShrinkRequest) SetDeviceInfoShrink(v string) *QueryMusicTypeShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *QueryMusicTypeShrinkRequest) SetPayloadShrink(v string) *QueryMusicTypeShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *QueryMusicTypeShrinkRequest) SetUserInfoShrink(v string) *QueryMusicTypeShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type QueryMusicTypeResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryMusicTypeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryMusicTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeResponseBody) SetCode(v int32) *QueryMusicTypeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMusicTypeResponseBody) SetMessage(v string) *QueryMusicTypeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMusicTypeResponseBody) SetRequestId(v string) *QueryMusicTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMusicTypeResponseBody) SetResult(v []*QueryMusicTypeResponseBodyResult) *QueryMusicTypeResponseBody {
	s.Result = v
	return s
}

type QueryMusicTypeResponseBodyResult struct {
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
}

func (s QueryMusicTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeResponseBodyResult) SetMusicType(v int64) *QueryMusicTypeResponseBodyResult {
	s.MusicType = &v
	return s
}

func (s *QueryMusicTypeResponseBodyResult) SetMusicTypeName(v string) *QueryMusicTypeResponseBodyResult {
	s.MusicTypeName = &v
	return s
}

type QueryMusicTypeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMusicTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMusicTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMusicTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeResponse) SetHeaders(v map[string]*string) *QueryMusicTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryMusicTypeResponse) SetStatusCode(v int32) *QueryMusicTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMusicTypeResponse) SetBody(v *QueryMusicTypeResponseBody) *QueryMusicTypeResponse {
	s.Body = v
	return s
}

type ReadMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ReadMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageHeaders) GoString() string {
	return s.String()
}

func (s *ReadMessageHeaders) SetCommonHeaders(v map[string]*string) *ReadMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReadMessageHeaders) SetXAcsAligenieAccessToken(v string) *ReadMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ReadMessageHeaders) SetAuthorization(v string) *ReadMessageHeaders {
	s.Authorization = &v
	return s
}

type ReadMessageRequest struct {
	MessageId *int64                      `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	UserInfo  *ReadMessageRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ReadMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageRequest) SetMessageId(v int64) *ReadMessageRequest {
	s.MessageId = &v
	return s
}

func (s *ReadMessageRequest) SetUserInfo(v *ReadMessageRequestUserInfo) *ReadMessageRequest {
	s.UserInfo = v
	return s
}

type ReadMessageRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ReadMessageRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ReadMessageRequestUserInfo) SetEncodeKey(v string) *ReadMessageRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetEncodeType(v string) *ReadMessageRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetId(v string) *ReadMessageRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetIdType(v string) *ReadMessageRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetOrganizationId(v string) *ReadMessageRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ReadMessageShrinkRequest struct {
	MessageId      *int64  `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ReadMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageShrinkRequest) SetMessageId(v int64) *ReadMessageShrinkRequest {
	s.MessageId = &v
	return s
}

func (s *ReadMessageShrinkRequest) SetUserInfoShrink(v string) *ReadMessageShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ReadMessageResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Result  *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ReadMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageResponseBody) SetCode(v string) *ReadMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageResponseBody) SetMessage(v string) *ReadMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageResponseBody) SetResult(v bool) *ReadMessageResponseBody {
	s.Result = &v
	return s
}

type ReadMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReadMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReadMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageResponse) SetHeaders(v map[string]*string) *ReadMessageResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageResponse) SetStatusCode(v int32) *ReadMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageResponse) SetBody(v *ReadMessageResponseBody) *ReadMessageResponse {
	s.Body = v
	return s
}

type ScanCodeBindHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ScanCodeBindHeaders) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindHeaders) GoString() string {
	return s.String()
}

func (s *ScanCodeBindHeaders) SetCommonHeaders(v map[string]*string) *ScanCodeBindHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ScanCodeBindHeaders) SetXAcsAligenieAccessToken(v string) *ScanCodeBindHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ScanCodeBindHeaders) SetAuthorization(v string) *ScanCodeBindHeaders {
	s.Authorization = &v
	return s
}

type ScanCodeBindRequest struct {
	BindReq  *ScanCodeBindRequestBindReq  `json:"BindReq,omitempty" xml:"BindReq,omitempty" type:"Struct"`
	UserInfo *ScanCodeBindRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ScanCodeBindRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindRequest) GoString() string {
	return s.String()
}

func (s *ScanCodeBindRequest) SetBindReq(v *ScanCodeBindRequestBindReq) *ScanCodeBindRequest {
	s.BindReq = v
	return s
}

func (s *ScanCodeBindRequest) SetUserInfo(v *ScanCodeBindRequestUserInfo) *ScanCodeBindRequest {
	s.UserInfo = v
	return s
}

type ScanCodeBindRequestBindReq struct {
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ExtInfo  *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s ScanCodeBindRequestBindReq) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindRequestBindReq) GoString() string {
	return s.String()
}

func (s *ScanCodeBindRequestBindReq) SetClientId(v string) *ScanCodeBindRequestBindReq {
	s.ClientId = &v
	return s
}

func (s *ScanCodeBindRequestBindReq) SetCode(v string) *ScanCodeBindRequestBindReq {
	s.Code = &v
	return s
}

func (s *ScanCodeBindRequestBindReq) SetExtInfo(v string) *ScanCodeBindRequestBindReq {
	s.ExtInfo = &v
	return s
}

type ScanCodeBindRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ScanCodeBindRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ScanCodeBindRequestUserInfo) SetEncodeKey(v string) *ScanCodeBindRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetEncodeType(v string) *ScanCodeBindRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetId(v string) *ScanCodeBindRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetIdType(v string) *ScanCodeBindRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetOrganizationId(v string) *ScanCodeBindRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ScanCodeBindShrinkRequest struct {
	BindReqShrink  *string `json:"BindReq,omitempty" xml:"BindReq,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ScanCodeBindShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScanCodeBindShrinkRequest) SetBindReqShrink(v string) *ScanCodeBindShrinkRequest {
	s.BindReqShrink = &v
	return s
}

func (s *ScanCodeBindShrinkRequest) SetUserInfoShrink(v string) *ScanCodeBindShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ScanCodeBindResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ScanCodeBindResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ScanCodeBindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindResponseBody) GoString() string {
	return s.String()
}

func (s *ScanCodeBindResponseBody) SetCode(v int32) *ScanCodeBindResponseBody {
	s.Code = &v
	return s
}

func (s *ScanCodeBindResponseBody) SetMessage(v string) *ScanCodeBindResponseBody {
	s.Message = &v
	return s
}

func (s *ScanCodeBindResponseBody) SetRequestId(v string) *ScanCodeBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanCodeBindResponseBody) SetResult(v *ScanCodeBindResponseBodyResult) *ScanCodeBindResponseBody {
	s.Result = v
	return s
}

type ScanCodeBindResponseBodyResult struct {
	BizGroup *string `json:"bizGroup,omitempty" xml:"bizGroup,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
}

func (s ScanCodeBindResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ScanCodeBindResponseBodyResult) SetBizGroup(v string) *ScanCodeBindResponseBodyResult {
	s.BizGroup = &v
	return s
}

func (s *ScanCodeBindResponseBodyResult) SetBizType(v string) *ScanCodeBindResponseBodyResult {
	s.BizType = &v
	return s
}

type ScanCodeBindResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScanCodeBindResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScanCodeBindResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanCodeBindResponse) GoString() string {
	return s.String()
}

func (s *ScanCodeBindResponse) SetHeaders(v map[string]*string) *ScanCodeBindResponse {
	s.Headers = v
	return s
}

func (s *ScanCodeBindResponse) SetStatusCode(v int32) *ScanCodeBindResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanCodeBindResponse) SetBody(v *ScanCodeBindResponseBody) *ScanCodeBindResponse {
	s.Body = v
	return s
}

type ScgSearchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ScgSearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchHeaders) GoString() string {
	return s.String()
}

func (s *ScgSearchHeaders) SetCommonHeaders(v map[string]*string) *ScgSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ScgSearchHeaders) SetXAcsAligenieAccessToken(v string) *ScgSearchHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ScgSearchHeaders) SetAuthorization(v string) *ScgSearchHeaders {
	s.Authorization = &v
	return s
}

type ScgSearchRequest struct {
	ScgFilter *ScgSearchRequestScgFilter `json:"ScgFilter,omitempty" xml:"ScgFilter,omitempty" type:"Struct"`
	TopicId   *string                    `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ScgSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchRequest) GoString() string {
	return s.String()
}

func (s *ScgSearchRequest) SetScgFilter(v *ScgSearchRequestScgFilter) *ScgSearchRequest {
	s.ScgFilter = v
	return s
}

func (s *ScgSearchRequest) SetTopicId(v string) *ScgSearchRequest {
	s.TopicId = &v
	return s
}

type ScgSearchRequestScgFilter struct {
	OffSetParam *ScgSearchRequestScgFilterOffSetParam `json:"OffSetParam,omitempty" xml:"OffSetParam,omitempty" type:"Struct"`
	PageParam   *ScgSearchRequestScgFilterPageParam   `json:"PageParam,omitempty" xml:"PageParam,omitempty" type:"Struct"`
	SortParam   *ScgSearchRequestScgFilterSortParam   `json:"SortParam,omitempty" xml:"SortParam,omitempty" type:"Struct"`
	UseOffSet   *bool                                 `json:"UseOffSet,omitempty" xml:"UseOffSet,omitempty"`
}

func (s ScgSearchRequestScgFilter) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchRequestScgFilter) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilter) SetOffSetParam(v *ScgSearchRequestScgFilterOffSetParam) *ScgSearchRequestScgFilter {
	s.OffSetParam = v
	return s
}

func (s *ScgSearchRequestScgFilter) SetPageParam(v *ScgSearchRequestScgFilterPageParam) *ScgSearchRequestScgFilter {
	s.PageParam = v
	return s
}

func (s *ScgSearchRequestScgFilter) SetSortParam(v *ScgSearchRequestScgFilterSortParam) *ScgSearchRequestScgFilter {
	s.SortParam = v
	return s
}

func (s *ScgSearchRequestScgFilter) SetUseOffSet(v bool) *ScgSearchRequestScgFilter {
	s.UseOffSet = &v
	return s
}

type ScgSearchRequestScgFilterOffSetParam struct {
	Limit  *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
}

func (s ScgSearchRequestScgFilterOffSetParam) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchRequestScgFilterOffSetParam) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilterOffSetParam) SetLimit(v int32) *ScgSearchRequestScgFilterOffSetParam {
	s.Limit = &v
	return s
}

func (s *ScgSearchRequestScgFilterOffSetParam) SetOffset(v int32) *ScgSearchRequestScgFilterOffSetParam {
	s.Offset = &v
	return s
}

type ScgSearchRequestScgFilterPageParam struct {
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ScgSearchRequestScgFilterPageParam) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchRequestScgFilterPageParam) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilterPageParam) SetPageNum(v int32) *ScgSearchRequestScgFilterPageParam {
	s.PageNum = &v
	return s
}

func (s *ScgSearchRequestScgFilterPageParam) SetPageSize(v int32) *ScgSearchRequestScgFilterPageParam {
	s.PageSize = &v
	return s
}

type ScgSearchRequestScgFilterSortParam struct {
	SortKey   *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	SortText  *string `json:"SortText,omitempty" xml:"SortText,omitempty"`
}

func (s ScgSearchRequestScgFilterSortParam) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchRequestScgFilterSortParam) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilterSortParam) SetSortKey(v string) *ScgSearchRequestScgFilterSortParam {
	s.SortKey = &v
	return s
}

func (s *ScgSearchRequestScgFilterSortParam) SetSortOrder(v string) *ScgSearchRequestScgFilterSortParam {
	s.SortOrder = &v
	return s
}

func (s *ScgSearchRequestScgFilterSortParam) SetSortText(v string) *ScgSearchRequestScgFilterSortParam {
	s.SortText = &v
	return s
}

type ScgSearchShrinkRequest struct {
	ScgFilterShrink *string `json:"ScgFilter,omitempty" xml:"ScgFilter,omitempty"`
	TopicId         *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ScgSearchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScgSearchShrinkRequest) SetScgFilterShrink(v string) *ScgSearchShrinkRequest {
	s.ScgFilterShrink = &v
	return s
}

func (s *ScgSearchShrinkRequest) SetTopicId(v string) *ScgSearchShrinkRequest {
	s.TopicId = &v
	return s
}

type ScgSearchResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNum   *int32                         `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ScgSearchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ScgSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ScgSearchResponseBody) SetCode(v int32) *ScgSearchResponseBody {
	s.Code = &v
	return s
}

func (s *ScgSearchResponseBody) SetMessage(v string) *ScgSearchResponseBody {
	s.Message = &v
	return s
}

func (s *ScgSearchResponseBody) SetPageNum(v int32) *ScgSearchResponseBody {
	s.PageNum = &v
	return s
}

func (s *ScgSearchResponseBody) SetPageSize(v int32) *ScgSearchResponseBody {
	s.PageSize = &v
	return s
}

func (s *ScgSearchResponseBody) SetRequestId(v string) *ScgSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScgSearchResponseBody) SetResult(v []*ScgSearchResponseBodyResult) *ScgSearchResponseBody {
	s.Result = v
	return s
}

type ScgSearchResponseBodyResult struct {
	Album           *bool                             `json:"Album,omitempty" xml:"Album,omitempty"`
	AlbumRawId      *string                           `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	AlbumType       *int32                            `json:"AlbumType,omitempty" xml:"AlbumType,omitempty"`
	Alias           []*string                         `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	AuthorIds       []*int64                          `json:"AuthorIds,omitempty" xml:"AuthorIds,omitempty" type:"Repeated"`
	AuthorNames     []*string                         `json:"AuthorNames,omitempty" xml:"AuthorNames,omitempty" type:"Repeated"`
	Category        *string                           `json:"Category,omitempty" xml:"Category,omitempty"`
	ContentType     *string                           `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Cover           *ScgSearchResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	IsAudition      *bool                             `json:"IsAudition,omitempty" xml:"IsAudition,omitempty"`
	IsCharge        *string                           `json:"IsCharge,omitempty" xml:"IsCharge,omitempty"`
	NeedCharge      *bool                             `json:"NeedCharge,omitempty" xml:"NeedCharge,omitempty"`
	RawId           *int64                            `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singers         *string                           `json:"Singers,omitempty" xml:"Singers,omitempty"`
	Source          *string                           `json:"Source,omitempty" xml:"Source,omitempty"`
	SupportAudition *bool                             `json:"SupportAudition,omitempty" xml:"SupportAudition,omitempty"`
	Title           *string                           `json:"Title,omitempty" xml:"Title,omitempty"`
	Type            *string                           `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ScgSearchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ScgSearchResponseBodyResult) SetAlbum(v bool) *ScgSearchResponseBodyResult {
	s.Album = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAlbumRawId(v string) *ScgSearchResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAlbumType(v int32) *ScgSearchResponseBodyResult {
	s.AlbumType = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAlias(v []*string) *ScgSearchResponseBodyResult {
	s.Alias = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAuthorIds(v []*int64) *ScgSearchResponseBodyResult {
	s.AuthorIds = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAuthorNames(v []*string) *ScgSearchResponseBodyResult {
	s.AuthorNames = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetCategory(v string) *ScgSearchResponseBodyResult {
	s.Category = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetContentType(v string) *ScgSearchResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetCover(v *ScgSearchResponseBodyResultCover) *ScgSearchResponseBodyResult {
	s.Cover = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetIsAudition(v bool) *ScgSearchResponseBodyResult {
	s.IsAudition = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetIsCharge(v string) *ScgSearchResponseBodyResult {
	s.IsCharge = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetNeedCharge(v bool) *ScgSearchResponseBodyResult {
	s.NeedCharge = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetRawId(v int64) *ScgSearchResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetSingers(v string) *ScgSearchResponseBodyResult {
	s.Singers = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetSource(v string) *ScgSearchResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetSupportAudition(v bool) *ScgSearchResponseBodyResult {
	s.SupportAudition = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetTitle(v string) *ScgSearchResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetType(v string) *ScgSearchResponseBodyResult {
	s.Type = &v
	return s
}

type ScgSearchResponseBodyResultCover struct {
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
	CanResize *bool   `json:"canResize,omitempty" xml:"canResize,omitempty"`
}

func (s ScgSearchResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *ScgSearchResponseBodyResultCover) SetImg(v string) *ScgSearchResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetLarge(v string) *ScgSearchResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetMedium(v string) *ScgSearchResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetSmall(v string) *ScgSearchResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetCanResize(v bool) *ScgSearchResponseBodyResultCover {
	s.CanResize = &v
	return s
}

type ScgSearchResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScgSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScgSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s ScgSearchResponse) GoString() string {
	return s.String()
}

func (s *ScgSearchResponse) SetHeaders(v map[string]*string) *ScgSearchResponse {
	s.Headers = v
	return s
}

func (s *ScgSearchResponse) SetStatusCode(v int32) *ScgSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ScgSearchResponse) SetBody(v *ScgSearchResponseBody) *ScgSearchResponse {
	s.Body = v
	return s
}

type SearchContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SearchContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchContentHeaders) GoString() string {
	return s.String()
}

func (s *SearchContentHeaders) SetCommonHeaders(v map[string]*string) *SearchContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchContentHeaders) SetXAcsAligenieAccessToken(v string) *SearchContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SearchContentHeaders) SetAuthorization(v string) *SearchContentHeaders {
	s.Authorization = &v
	return s
}

type SearchContentRequest struct {
	DeviceInfo *SearchContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Request    *SearchContentRequestRequest    `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	UserInfo   *SearchContentRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SearchContentRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchContentRequest) GoString() string {
	return s.String()
}

func (s *SearchContentRequest) SetDeviceInfo(v *SearchContentRequestDeviceInfo) *SearchContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *SearchContentRequest) SetRequest(v *SearchContentRequestRequest) *SearchContentRequest {
	s.Request = v
	return s
}

func (s *SearchContentRequest) SetUserInfo(v *SearchContentRequestUserInfo) *SearchContentRequest {
	s.UserInfo = v
	return s
}

type SearchContentRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SearchContentRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *SearchContentRequestDeviceInfo) SetEncodeKey(v string) *SearchContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetEncodeType(v string) *SearchContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetId(v string) *SearchContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetIdType(v string) *SearchContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetOrganizationId(v string) *SearchContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type SearchContentRequestRequest struct {
	Cate       *string `json:"Cate,omitempty" xml:"Cate,omitempty"`
	PageNum    *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query      *string `json:"Query,omitempty" xml:"Query,omitempty"`
	QueryAlbum *bool   `json:"QueryAlbum,omitempty" xml:"QueryAlbum,omitempty"`
	SubCate    *string `json:"SubCate,omitempty" xml:"SubCate,omitempty"`
}

func (s SearchContentRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchContentRequestRequest) GoString() string {
	return s.String()
}

func (s *SearchContentRequestRequest) SetCate(v string) *SearchContentRequestRequest {
	s.Cate = &v
	return s
}

func (s *SearchContentRequestRequest) SetPageNum(v int32) *SearchContentRequestRequest {
	s.PageNum = &v
	return s
}

func (s *SearchContentRequestRequest) SetPageSize(v int32) *SearchContentRequestRequest {
	s.PageSize = &v
	return s
}

func (s *SearchContentRequestRequest) SetQuery(v string) *SearchContentRequestRequest {
	s.Query = &v
	return s
}

func (s *SearchContentRequestRequest) SetQueryAlbum(v bool) *SearchContentRequestRequest {
	s.QueryAlbum = &v
	return s
}

func (s *SearchContentRequestRequest) SetSubCate(v string) *SearchContentRequestRequest {
	s.SubCate = &v
	return s
}

type SearchContentRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SearchContentRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SearchContentRequestUserInfo) SetEncodeKey(v string) *SearchContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetEncodeType(v string) *SearchContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetId(v string) *SearchContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetIdType(v string) *SearchContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetOrganizationId(v string) *SearchContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type SearchContentShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	RequestShrink    *string `json:"Request,omitempty" xml:"Request,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SearchContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchContentShrinkRequest) SetDeviceInfoShrink(v string) *SearchContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *SearchContentShrinkRequest) SetRequestShrink(v string) *SearchContentShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *SearchContentShrinkRequest) SetUserInfoShrink(v string) *SearchContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type SearchContentResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*SearchContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s SearchContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchContentResponseBody) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBody) SetCode(v int32) *SearchContentResponseBody {
	s.Code = &v
	return s
}

func (s *SearchContentResponseBody) SetMessage(v string) *SearchContentResponseBody {
	s.Message = &v
	return s
}

func (s *SearchContentResponseBody) SetRequestId(v string) *SearchContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchContentResponseBody) SetResult(v []*SearchContentResponseBodyResult) *SearchContentResponseBody {
	s.Result = v
	return s
}

type SearchContentResponseBodyResult struct {
	AlbumId     *string                                   `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	Alias       []*string                                 `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	Audition    *bool                                     `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors     []*SearchContentResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	Category    *string                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	Charge      *bool                                     `json:"Charge,omitempty" xml:"Charge,omitempty"`
	CommCateId  *int64                                    `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *SearchContentResponseBodyResultCover     `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *int64                                    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HotScore    *float64                                  `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	Id          *int64                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	ItemType    *string                                   `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	Lyric       *string                                   `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
	OrderIndex  *string                                   `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	Source      *string                                   `json:"Source,omitempty" xml:"Source,omitempty"`
	Styles      []*string                                 `json:"Styles,omitempty" xml:"Styles,omitempty" type:"Repeated"`
	Title       *string                                   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid       *string                                   `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s SearchContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResult) SetAlbumId(v string) *SearchContentResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetAlias(v []*string) *SearchContentResponseBodyResult {
	s.Alias = v
	return s
}

func (s *SearchContentResponseBodyResult) SetAudition(v bool) *SearchContentResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetAuthors(v []*SearchContentResponseBodyResultAuthors) *SearchContentResponseBodyResult {
	s.Authors = v
	return s
}

func (s *SearchContentResponseBodyResult) SetCategory(v string) *SearchContentResponseBodyResult {
	s.Category = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetCharge(v bool) *SearchContentResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetCommCateId(v int64) *SearchContentResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetCover(v *SearchContentResponseBodyResultCover) *SearchContentResponseBodyResult {
	s.Cover = v
	return s
}

func (s *SearchContentResponseBodyResult) SetDescription(v string) *SearchContentResponseBodyResult {
	s.Description = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetDuration(v int64) *SearchContentResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetHotScore(v float64) *SearchContentResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetId(v int64) *SearchContentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetItemType(v string) *SearchContentResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetLyric(v string) *SearchContentResponseBodyResult {
	s.Lyric = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetOrderIndex(v string) *SearchContentResponseBodyResult {
	s.OrderIndex = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetSource(v string) *SearchContentResponseBodyResult {
	s.Source = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetStyles(v []*string) *SearchContentResponseBodyResult {
	s.Styles = v
	return s
}

func (s *SearchContentResponseBodyResult) SetTitle(v string) *SearchContentResponseBodyResult {
	s.Title = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetType(v string) *SearchContentResponseBodyResult {
	s.Type = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetValid(v string) *SearchContentResponseBodyResult {
	s.Valid = &v
	return s
}

type SearchContentResponseBodyResultAuthors struct {
	AuthorTypes []*string                                    `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Cover       *SearchContentResponseBodyResultAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Gender      *string                                      `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Id          *int64                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Online      *bool                                        `json:"Online,omitempty" xml:"Online,omitempty"`
	RawId       *string                                      `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Source      *string                                      `json:"Source,omitempty" xml:"Source,omitempty"`
	Title       *string                                      `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchContentResponseBodyResultAuthors) String() string {
	return tea.Prettify(s)
}

func (s SearchContentResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResultAuthors) SetAuthorTypes(v []*string) *SearchContentResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetCover(v *SearchContentResponseBodyResultAuthorsCover) *SearchContentResponseBodyResultAuthors {
	s.Cover = v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetDescription(v string) *SearchContentResponseBodyResultAuthors {
	s.Description = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetGender(v string) *SearchContentResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetId(v int64) *SearchContentResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetOnline(v bool) *SearchContentResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetRawId(v string) *SearchContentResponseBodyResultAuthors {
	s.RawId = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetSource(v string) *SearchContentResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetTitle(v string) *SearchContentResponseBodyResultAuthors {
	s.Title = &v
	return s
}

type SearchContentResponseBodyResultAuthorsCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s SearchContentResponseBodyResultAuthorsCover) String() string {
	return tea.Prettify(s)
}

func (s SearchContentResponseBodyResultAuthorsCover) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetCanResize(v bool) *SearchContentResponseBodyResultAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetImg(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Img = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetLarge(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Large = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetMedium(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Medium = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetSmall(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Small = &v
	return s
}

type SearchContentResponseBodyResultCover struct {
	CanResize *bool   `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	Img       *string `json:"Img,omitempty" xml:"Img,omitempty"`
	Large     *string `json:"Large,omitempty" xml:"Large,omitempty"`
	Mediam    *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Small     *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s SearchContentResponseBodyResultCover) String() string {
	return tea.Prettify(s)
}

func (s SearchContentResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResultCover) SetCanResize(v bool) *SearchContentResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetImg(v string) *SearchContentResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetLarge(v string) *SearchContentResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetMediam(v string) *SearchContentResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetMedium(v string) *SearchContentResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetSmall(v string) *SearchContentResponseBodyResultCover {
	s.Small = &v
	return s
}

type SearchContentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchContentResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchContentResponse) GoString() string {
	return s.String()
}

func (s *SearchContentResponse) SetHeaders(v map[string]*string) *SearchContentResponse {
	s.Headers = v
	return s
}

func (s *SearchContentResponse) SetStatusCode(v int32) *SearchContentResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchContentResponse) SetBody(v *SearchContentResponseBody) *SearchContentResponse {
	s.Body = v
	return s
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetXAcsAligenieAccessToken(v string) *SendMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SendMessageHeaders) SetAuthorization(v string) *SendMessageHeaders {
	s.Authorization = &v
	return s
}

type SendMessageRequest struct {
	Url      *string                     `json:"Url,omitempty" xml:"Url,omitempty"`
	UserInfo *SendMessageRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetUrl(v string) *SendMessageRequest {
	s.Url = &v
	return s
}

func (s *SendMessageRequest) SetUserInfo(v *SendMessageRequestUserInfo) *SendMessageRequest {
	s.UserInfo = v
	return s
}

type SendMessageRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SendMessageRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SendMessageRequestUserInfo) SetEncodeKey(v string) *SendMessageRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetEncodeType(v string) *SendMessageRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetId(v string) *SendMessageRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetIdType(v string) *SendMessageRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetOrganizationId(v string) *SendMessageRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type SendMessageShrinkRequest struct {
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SendMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendMessageShrinkRequest) SetUrl(v string) *SendMessageShrinkRequest {
	s.Url = &v
	return s
}

func (s *SendMessageShrinkRequest) SetUserInfoShrink(v string) *SendMessageShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type SendMessageResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Result  *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetCode(v string) *SendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageResponseBody) SetMessage(v string) *SendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendMessageResponseBody) SetResult(v bool) *SendMessageResponseBody {
	s.Result = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type SetDeviceSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SetDeviceSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceSettingHeaders) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingHeaders) SetCommonHeaders(v map[string]*string) *SetDeviceSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDeviceSettingHeaders) SetXAcsAligenieAccessToken(v string) *SetDeviceSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SetDeviceSettingHeaders) SetAuthorization(v string) *SetDeviceSettingHeaders {
	s.Authorization = &v
	return s
}

type SetDeviceSettingRequest struct {
	DeviceInfo *SetDeviceSettingRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Key        *string                            `json:"Key,omitempty" xml:"Key,omitempty"`
	Value      interface{}                        `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetDeviceSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceSettingRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingRequest) SetDeviceInfo(v *SetDeviceSettingRequestDeviceInfo) *SetDeviceSettingRequest {
	s.DeviceInfo = v
	return s
}

func (s *SetDeviceSettingRequest) SetKey(v string) *SetDeviceSettingRequest {
	s.Key = &v
	return s
}

func (s *SetDeviceSettingRequest) SetValue(v interface{}) *SetDeviceSettingRequest {
	s.Value = v
	return s
}

type SetDeviceSettingRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SetDeviceSettingRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceSettingRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingRequestDeviceInfo) SetEncodeKey(v string) *SetDeviceSettingRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetEncodeType(v string) *SetDeviceSettingRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetId(v string) *SetDeviceSettingRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetIdType(v string) *SetDeviceSettingRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetOrganizationId(v string) *SetDeviceSettingRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type SetDeviceSettingShrinkRequest struct {
	DeviceInfoShrink *string     `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	Key              *string     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value            interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetDeviceSettingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingShrinkRequest) SetDeviceInfoShrink(v string) *SetDeviceSettingShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *SetDeviceSettingShrinkRequest) SetKey(v string) *SetDeviceSettingShrinkRequest {
	s.Key = &v
	return s
}

func (s *SetDeviceSettingShrinkRequest) SetValue(v interface{}) *SetDeviceSettingShrinkRequest {
	s.Value = v
	return s
}

type SetDeviceSettingResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetDeviceSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingResponseBody) SetCode(v int32) *SetDeviceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceSettingResponseBody) SetMessage(v string) *SetDeviceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *SetDeviceSettingResponseBody) SetRequestId(v string) *SetDeviceSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeviceSettingResponseBody) SetResult(v bool) *SetDeviceSettingResponseBody {
	s.Result = &v
	return s
}

type SetDeviceSettingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDeviceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDeviceSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceSettingResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingResponse) SetHeaders(v map[string]*string) *SetDeviceSettingResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceSettingResponse) SetStatusCode(v int32) *SetDeviceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceSettingResponse) SetBody(v *SetDeviceSettingResponseBody) *SetDeviceSettingResponse {
	s.Body = v
	return s
}

type UnbindAligenieUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UnbindAligenieUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnbindAligenieUserHeaders) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserHeaders) SetCommonHeaders(v map[string]*string) *UnbindAligenieUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnbindAligenieUserHeaders) SetXAcsAligenieAccessToken(v string) *UnbindAligenieUserHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UnbindAligenieUserHeaders) SetAuthorization(v string) *UnbindAligenieUserHeaders {
	s.Authorization = &v
	return s
}

type UnbindAligenieUserRequest struct {
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s UnbindAligenieUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindAligenieUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserRequest) SetLoginStateAccessToken(v string) *UnbindAligenieUserRequest {
	s.LoginStateAccessToken = &v
	return s
}

type UnbindAligenieUserResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindAligenieUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindAligenieUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserResponseBody) SetCode(v int32) *UnbindAligenieUserResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindAligenieUserResponseBody) SetMessage(v string) *UnbindAligenieUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindAligenieUserResponseBody) SetRequestId(v string) *UnbindAligenieUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAligenieUserResponseBody) SetSuccess(v bool) *UnbindAligenieUserResponseBody {
	s.Success = &v
	return s
}

type UnbindAligenieUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindAligenieUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindAligenieUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindAligenieUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserResponse) SetHeaders(v map[string]*string) *UnbindAligenieUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindAligenieUserResponse) SetStatusCode(v int32) *UnbindAligenieUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAligenieUserResponse) SetBody(v *UnbindAligenieUserResponseBody) *UnbindAligenieUserResponse {
	s.Body = v
	return s
}

type UnbindDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UnbindDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceHeaders) GoString() string {
	return s.String()
}

func (s *UnbindDeviceHeaders) SetCommonHeaders(v map[string]*string) *UnbindDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnbindDeviceHeaders) SetXAcsAligenieAccessToken(v string) *UnbindDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UnbindDeviceHeaders) SetAuthorization(v string) *UnbindDeviceHeaders {
	s.Authorization = &v
	return s
}

type UnbindDeviceRequest struct {
	DeviceInfo *UnbindDeviceRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *UnbindDeviceRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s UnbindDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequest) SetDeviceInfo(v *UnbindDeviceRequestDeviceInfo) *UnbindDeviceRequest {
	s.DeviceInfo = v
	return s
}

func (s *UnbindDeviceRequest) SetUserInfo(v *UnbindDeviceRequestUserInfo) *UnbindDeviceRequest {
	s.UserInfo = v
	return s
}

type UnbindDeviceRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UnbindDeviceRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequestDeviceInfo) SetEncodeKey(v string) *UnbindDeviceRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetEncodeType(v string) *UnbindDeviceRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetId(v string) *UnbindDeviceRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetIdType(v string) *UnbindDeviceRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetOrganizationId(v string) *UnbindDeviceRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type UnbindDeviceRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UnbindDeviceRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequestUserInfo) SetEncodeKey(v string) *UnbindDeviceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetEncodeType(v string) *UnbindDeviceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetId(v string) *UnbindDeviceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetIdType(v string) *UnbindDeviceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetOrganizationId(v string) *UnbindDeviceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type UnbindDeviceShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s UnbindDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceShrinkRequest) SetDeviceInfoShrink(v string) *UnbindDeviceShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *UnbindDeviceShrinkRequest) SetUserInfoShrink(v string) *UnbindDeviceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type UnbindDeviceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UnbindDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponseBody) SetCode(v int32) *UnbindDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetMessage(v string) *UnbindDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetRequestId(v string) *UnbindDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetResult(v bool) *UnbindDeviceResponseBody {
	s.Result = &v
	return s
}

type UnbindDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponse) SetHeaders(v map[string]*string) *UnbindDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindDeviceResponse) SetStatusCode(v int32) *UnbindDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDeviceResponse) SetBody(v *UnbindDeviceResponseBody) *UnbindDeviceResponse {
	s.Body = v
	return s
}

type UpdateAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateAlarmHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAlarmHeaders) SetCommonHeaders(v map[string]*string) *UpdateAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAlarmHeaders) SetXAcsAligenieAccessToken(v string) *UpdateAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateAlarmHeaders) SetAuthorization(v string) *UpdateAlarmHeaders {
	s.Authorization = &v
	return s
}

type UpdateAlarmRequest struct {
	DeviceInfo *UpdateAlarmRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *UpdateAlarmRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *UpdateAlarmRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s UpdateAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequest) SetDeviceInfo(v *UpdateAlarmRequestDeviceInfo) *UpdateAlarmRequest {
	s.DeviceInfo = v
	return s
}

func (s *UpdateAlarmRequest) SetPayload(v *UpdateAlarmRequestPayload) *UpdateAlarmRequest {
	s.Payload = v
	return s
}

func (s *UpdateAlarmRequest) SetUserInfo(v *UpdateAlarmRequestUserInfo) *UpdateAlarmRequest {
	s.UserInfo = v
	return s
}

type UpdateAlarmRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateAlarmRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestDeviceInfo) SetEncodeKey(v string) *UpdateAlarmRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetEncodeType(v string) *UpdateAlarmRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetId(v string) *UpdateAlarmRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetIdType(v string) *UpdateAlarmRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetOrganizationId(v string) *UpdateAlarmRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type UpdateAlarmRequestPayload struct {
	AlarmId      *int64                                 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	MusicInfo    *UpdateAlarmRequestPayloadMusicInfo    `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	ScheduleInfo *UpdateAlarmRequestPayloadScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	Volume       *int32                                 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s UpdateAlarmRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestPayload) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayload) SetAlarmId(v int64) *UpdateAlarmRequestPayload {
	s.AlarmId = &v
	return s
}

func (s *UpdateAlarmRequestPayload) SetMusicInfo(v *UpdateAlarmRequestPayloadMusicInfo) *UpdateAlarmRequestPayload {
	s.MusicInfo = v
	return s
}

func (s *UpdateAlarmRequestPayload) SetScheduleInfo(v *UpdateAlarmRequestPayloadScheduleInfo) *UpdateAlarmRequestPayload {
	s.ScheduleInfo = v
	return s
}

func (s *UpdateAlarmRequestPayload) SetVolume(v int32) *UpdateAlarmRequestPayload {
	s.Volume = &v
	return s
}

type UpdateAlarmRequestPayloadMusicInfo struct {
	MusicId       *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName     *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	MusicUrl      *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s UpdateAlarmRequestPayloadMusicInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestPayloadMusicInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicId(v int64) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicId = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicName(v string) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicName = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicType(v int64) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicType = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicTypeName(v string) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicUrl(v string) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicUrl = &v
	return s
}

type UpdateAlarmRequestPayloadScheduleInfo struct {
	Once                *UpdateAlarmRequestPayloadScheduleInfoOnce                `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	StatutoryWorkingDay *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	Type                *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly              *UpdateAlarmRequestPayloadScheduleInfoWeekly              `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s UpdateAlarmRequestPayloadScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetOnce(v *UpdateAlarmRequestPayloadScheduleInfoOnce) *UpdateAlarmRequestPayloadScheduleInfo {
	s.Once = v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetStatutoryWorkingDay(v *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) *UpdateAlarmRequestPayloadScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetType(v string) *UpdateAlarmRequestPayloadScheduleInfo {
	s.Type = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetWeekly(v *UpdateAlarmRequestPayloadScheduleInfoWeekly) *UpdateAlarmRequestPayloadScheduleInfo {
	s.Weekly = v
	return s
}

type UpdateAlarmRequestPayloadScheduleInfoOnce struct {
	Day    *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month  *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	Year   *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s UpdateAlarmRequestPayloadScheduleInfoOnce) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetDay(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetHour(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetMinute(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetMonth(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetYear(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Year = &v
	return s
}

type UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay struct {
	Hour   *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetHour(v int32) *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

type UpdateAlarmRequestPayloadScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hour       *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute     *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s UpdateAlarmRequestPayloadScheduleInfoWeekly) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *UpdateAlarmRequestPayloadScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) SetHour(v int32) *UpdateAlarmRequestPayloadScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) SetMinute(v int32) *UpdateAlarmRequestPayloadScheduleInfoWeekly {
	s.Minute = &v
	return s
}

type UpdateAlarmRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateAlarmRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequestUserInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestUserInfo) SetEncodeKey(v string) *UpdateAlarmRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetEncodeType(v string) *UpdateAlarmRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetId(v string) *UpdateAlarmRequestUserInfo {
	s.Id = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetIdType(v string) *UpdateAlarmRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetOrganizationId(v string) *UpdateAlarmRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type UpdateAlarmShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s UpdateAlarmShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlarmShrinkRequest) SetDeviceInfoShrink(v string) *UpdateAlarmShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *UpdateAlarmShrinkRequest) SetPayloadShrink(v string) *UpdateAlarmShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *UpdateAlarmShrinkRequest) SetUserInfoShrink(v string) *UpdateAlarmShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type UpdateAlarmResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlarmResponseBody) SetCode(v int32) *UpdateAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetMessage(v string) *UpdateAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetRequestId(v string) *UpdateAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetResult(v bool) *UpdateAlarmResponseBody {
	s.Result = &v
	return s
}

type UpdateAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlarmResponse) SetHeaders(v map[string]*string) *UpdateAlarmResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlarmResponse) SetStatusCode(v int32) *UpdateAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlarmResponse) SetBody(v *UpdateAlarmResponseBody) *UpdateAlarmResponse {
	s.Body = v
	return s
}

type ResultValue struct {
	DeviceOpenId    *string                      `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds  []*ResultValueDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
	Name            *string                      `json:"Name,omitempty" xml:"Name,omitempty"`
	FirmwareVersion *string                      `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	Mac             *string                      `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Sn              *string                      `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s ResultValue) String() string {
	return tea.Prettify(s)
}

func (s ResultValue) GoString() string {
	return s.String()
}

func (s *ResultValue) SetDeviceOpenId(v string) *ResultValue {
	s.DeviceOpenId = &v
	return s
}

func (s *ResultValue) SetDeviceUnionIds(v []*ResultValueDeviceUnionIds) *ResultValue {
	s.DeviceUnionIds = v
	return s
}

func (s *ResultValue) SetName(v string) *ResultValue {
	s.Name = &v
	return s
}

func (s *ResultValue) SetFirmwareVersion(v string) *ResultValue {
	s.FirmwareVersion = &v
	return s
}

func (s *ResultValue) SetMac(v string) *ResultValue {
	s.Mac = &v
	return s
}

func (s *ResultValue) SetSn(v string) *ResultValue {
	s.Sn = &v
	return s
}

type ResultValueDeviceUnionIds struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	DeviceUnionId  *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
}

func (s ResultValueDeviceUnionIds) String() string {
	return tea.Prettify(s)
}

func (s ResultValueDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *ResultValueDeviceUnionIds) SetOrganizationId(v string) *ResultValueDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *ResultValueDeviceUnionIds) SetDeviceUnionId(v string) *ResultValueDeviceUnionIds {
	s.DeviceUnionId = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddAndRemoveFavoriteContent(request *AddAndRemoveFavoriteContentRequest) (_result *AddAndRemoveFavoriteContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAndRemoveFavoriteContentHeaders{}
	_result = &AddAndRemoveFavoriteContentResponse{}
	_body, _err := client.AddAndRemoveFavoriteContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAndRemoveFavoriteContentWithOptions(tmpReq *AddAndRemoveFavoriteContentRequest, headers *AddAndRemoveFavoriteContentHeaders, runtime *util.RuntimeOptions) (_result *AddAndRemoveFavoriteContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddAndRemoveFavoriteContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenAddAndRemoveFavoriteContentRequest))) {
		request.OpenAddAndRemoveFavoriteContentRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenAddAndRemoveFavoriteContentRequest), tea.String("OpenAddAndRemoveFavoriteContentRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenAddAndRemoveFavoriteContentRequestShrink)) {
		body["OpenAddAndRemoveFavoriteContentRequest"] = request.OpenAddAndRemoveFavoriteContentRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddAndRemoveFavoriteContent"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/AddAndRemoveFavoriteContent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAndRemoveFavoriteContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSub(request *AddSubRequest) (_result *AddSubResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddSubHeaders{}
	_result = &AddSubResponse{}
	_body, _err := client.AddSubWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSubWithOptions(tmpReq *AddSubRequest, headers *AddSubHeaders, runtime *util.RuntimeOptions) (_result *AddSubResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddSubShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.AddSubscriptionInfoRequest))) {
		request.AddSubscriptionInfoRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.AddSubscriptionInfoRequest), tea.String("AddSubscriptionInfoRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddSubscriptionInfoRequestShrink)) {
		query["AddSubscriptionInfoRequest"] = request.AddSubscriptionInfoRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSub"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/addSub"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthLoginWithAligenieUserInfo(request *AuthLoginWithAligenieUserInfoRequest) (_result *AuthLoginWithAligenieUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AuthLoginWithAligenieUserInfoHeaders{}
	_result = &AuthLoginWithAligenieUserInfoResponse{}
	_body, _err := client.AuthLoginWithAligenieUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthLoginWithAligenieUserInfoWithOptions(request *AuthLoginWithAligenieUserInfoRequest, headers *AuthLoginWithAligenieUserInfoHeaders, runtime *util.RuntimeOptions) (_result *AuthLoginWithAligenieUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptedAligenieUserIdentifier)) {
		body["EncryptedAligenieUserIdentifier"] = request.EncryptedAligenieUserIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthLoginWithAligenieUserInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/authLoginWithAligenieUserInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthLoginWithAligenieUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthLoginWithAligenieUserInfoGeneratedByPhoneNumber(request *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) (_result *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders{}
	_result = &AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse{}
	_body, _err := client.AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberWithOptions(request *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest, headers *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders, runtime *util.RuntimeOptions) (_result *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthLoginWithAligenieUserInfoGeneratedByPhoneNumber"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/authLoginWithAligenieUserInfoGeneratedByPhoneNumber"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthLoginWithTaobaoUserInfo(request *AuthLoginWithTaobaoUserInfoRequest) (_result *AuthLoginWithTaobaoUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AuthLoginWithTaobaoUserInfoHeaders{}
	_result = &AuthLoginWithTaobaoUserInfoResponse{}
	_body, _err := client.AuthLoginWithTaobaoUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthLoginWithTaobaoUserInfoWithOptions(request *AuthLoginWithTaobaoUserInfoRequest, headers *AuthLoginWithTaobaoUserInfoHeaders, runtime *util.RuntimeOptions) (_result *AuthLoginWithTaobaoUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptedTaobaoUserIdentifier)) {
		body["EncryptedTaobaoUserIdentifier"] = request.EncryptedTaobaoUserIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthLoginWithTaobaoUserInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/authLoginWithTaobaoUserInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthLoginWithTaobaoUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthLoginWithThirdUserInfo(request *AuthLoginWithThirdUserInfoRequest) (_result *AuthLoginWithThirdUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AuthLoginWithThirdUserInfoHeaders{}
	_result = &AuthLoginWithThirdUserInfoResponse{}
	_body, _err := client.AuthLoginWithThirdUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthLoginWithThirdUserInfoWithOptions(tmpReq *AuthLoginWithThirdUserInfoRequest, headers *AuthLoginWithThirdUserInfoHeaders, runtime *util.RuntimeOptions) (_result *AuthLoginWithThirdUserInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AuthLoginWithThirdUserInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtInfo)) {
		request.ExtInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtInfo, tea.String("ExtInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtInfoShrink)) {
		body["ExtInfo"] = request.ExtInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		body["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdUserIdentifier)) {
		body["ThirdUserIdentifier"] = request.ThirdUserIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdUserType)) {
		body["ThirdUserType"] = request.ThirdUserType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthLoginWithThirdUserInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/authLoginWithThirdUserInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthLoginWithThirdUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAuthCodeBindForExt(request *CheckAuthCodeBindForExtRequest) (_result *CheckAuthCodeBindForExtResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckAuthCodeBindForExtHeaders{}
	_result = &CheckAuthCodeBindForExtResponse{}
	_body, _err := client.CheckAuthCodeBindForExtWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAuthCodeBindForExtWithOptions(tmpReq *CheckAuthCodeBindForExtRequest, headers *CheckAuthCodeBindForExtHeaders, runtime *util.RuntimeOptions) (_result *CheckAuthCodeBindForExtResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CheckAuthCodeBindForExtShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeKey)) {
		query["EncodeKey"] = request.EncodeKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAuthCodeBindForExt"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/checkAuthCodeBindForExt"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAuthCodeBindForExtResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlarm(request *CreateAlarmRequest) (_result *CreateAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAlarmHeaders{}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CreateAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlarmWithOptions(tmpReq *CreateAlarmRequest, headers *CreateAlarmHeaders, runtime *util.RuntimeOptions) (_result *CreateAlarmResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlarm"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/createAlarm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePlayingList(request *CreatePlayingListRequest) (_result *CreatePlayingListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatePlayingListHeaders{}
	_result = &CreatePlayingListResponse{}
	_body, _err := client.CreatePlayingListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePlayingListWithOptions(tmpReq *CreatePlayingListRequest, headers *CreatePlayingListHeaders, runtime *util.RuntimeOptions) (_result *CreatePlayingListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenCreatePlayingListRequest))) {
		request.OpenCreatePlayingListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenCreatePlayingListRequest), tea.String("OpenCreatePlayingListRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenCreatePlayingListRequestShrink)) {
		body["OpenCreatePlayingListRequest"] = request.OpenCreatePlayingListRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePlayingList"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/CreatePlayingList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePlayingListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScheduleTask(request *CreateScheduleTaskRequest) (_result *CreateScheduleTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateScheduleTaskHeaders{}
	_result = &CreateScheduleTaskResponse{}
	_body, _err := client.CreateScheduleTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduleTaskWithOptions(tmpReq *CreateScheduleTaskRequest, headers *CreateScheduleTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateScheduleTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateScheduleTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduleTask"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/CreateScheduleTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlarms(request *DeleteAlarmsRequest) (_result *DeleteAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAlarmsHeaders{}
	_result = &DeleteAlarmsResponse{}
	_body, _err := client.DeleteAlarmsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlarmsWithOptions(tmpReq *DeleteAlarmsRequest, headers *DeleteAlarmsHeaders, runtime *util.RuntimeOptions) (_result *DeleteAlarmsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAlarmsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlarms"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/deleteAlarms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScheduleTask(request *DeleteScheduleTaskRequest) (_result *DeleteScheduleTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteScheduleTaskHeaders{}
	_result = &DeleteScheduleTaskResponse{}
	_body, _err := client.DeleteScheduleTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScheduleTaskWithOptions(tmpReq *DeleteScheduleTaskRequest, headers *DeleteScheduleTaskHeaders, runtime *util.RuntimeOptions) (_result *DeleteScheduleTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteScheduleTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScheduleTask"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/DeleteScheduleTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSub(request *DeleteSubRequest) (_result *DeleteSubResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSubHeaders{}
	_result = &DeleteSubResponse{}
	_body, _err := client.DeleteSubWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubWithOptions(request *DeleteSubRequest, headers *DeleteSubHeaders, runtime *util.RuntimeOptions) (_result *DeleteSubResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSub"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/deleteSub"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeviceControl(request *DeviceControlRequest) (_result *DeviceControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeviceControlHeaders{}
	_result = &DeviceControlResponse{}
	_body, _err := client.DeviceControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeviceControlWithOptions(tmpReq *DeviceControlRequest, headers *DeviceControlHeaders, runtime *util.RuntimeOptions) (_result *DeviceControlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeviceControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ControlRequest))) {
		request.ControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ControlRequest), tea.String("ControlRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ControlRequestShrink)) {
		body["ControlRequest"] = request.ControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceControl"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/control"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EcologyOpennessAuthenticate(request *EcologyOpennessAuthenticateRequest) (_result *EcologyOpennessAuthenticateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EcologyOpennessAuthenticateHeaders{}
	_result = &EcologyOpennessAuthenticateResponse{}
	_body, _err := client.EcologyOpennessAuthenticateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EcologyOpennessAuthenticateWithOptions(request *EcologyOpennessAuthenticateRequest, headers *EcologyOpennessAuthenticateHeaders, runtime *util.RuntimeOptions) (_result *EcologyOpennessAuthenticateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeKey)) {
		body["EncodeKey"] = request.EncodeKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		body["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.LoginStateAccessToken)) {
		body["LoginStateAccessToken"] = request.LoginStateAccessToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EcologyOpennessAuthenticate"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ecologyOpennessAuthenticate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EcologyOpennessAuthenticateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EcologyOpennessSendVerificationCode(request *EcologyOpennessSendVerificationCodeRequest) (_result *EcologyOpennessSendVerificationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EcologyOpennessSendVerificationCodeHeaders{}
	_result = &EcologyOpennessSendVerificationCodeResponse{}
	_body, _err := client.EcologyOpennessSendVerificationCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EcologyOpennessSendVerificationCodeWithOptions(request *EcologyOpennessSendVerificationCodeRequest, headers *EcologyOpennessSendVerificationCodeHeaders, runtime *util.RuntimeOptions) (_result *EcologyOpennessSendVerificationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EcologyOpennessSendVerificationCode"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ecologyOpennessSendVerificationCode"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EcologyOpennessSendVerificationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindUserlistToAuthLoginWithPhoneNumber(request *FindUserlistToAuthLoginWithPhoneNumberRequest) (_result *FindUserlistToAuthLoginWithPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FindUserlistToAuthLoginWithPhoneNumberHeaders{}
	_result = &FindUserlistToAuthLoginWithPhoneNumberResponse{}
	_body, _err := client.FindUserlistToAuthLoginWithPhoneNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindUserlistToAuthLoginWithPhoneNumberWithOptions(request *FindUserlistToAuthLoginWithPhoneNumberRequest, headers *FindUserlistToAuthLoginWithPhoneNumberHeaders, runtime *util.RuntimeOptions) (_result *FindUserlistToAuthLoginWithPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindUserlistToAuthLoginWithPhoneNumber"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/findUserlistToAuthLoginWithPhoneNumber"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FindUserlistToAuthLoginWithPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlarm(request *GetAlarmRequest) (_result *GetAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAlarmHeaders{}
	_result = &GetAlarmResponse{}
	_body, _err := client.GetAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlarmWithOptions(tmpReq *GetAlarmRequest, headers *GetAlarmHeaders, runtime *util.RuntimeOptions) (_result *GetAlarmResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlarm"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getAlarm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlbum(request *GetAlbumRequest) (_result *GetAlbumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAlbumHeaders{}
	_result = &GetAlbumResponse{}
	_body, _err := client.GetAlbumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlbumWithOptions(request *GetAlbumRequest, headers *GetAlbumHeaders, runtime *util.RuntimeOptions) (_result *GetAlbumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlbum"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/GetAlbum"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlbumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlbumDetailById(request *GetAlbumDetailByIdRequest) (_result *GetAlbumDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAlbumDetailByIdHeaders{}
	_result = &GetAlbumDetailByIdResponse{}
	_body, _err := client.GetAlbumDetailByIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlbumDetailByIdWithOptions(request *GetAlbumDetailByIdRequest, headers *GetAlbumDetailByIdHeaders, runtime *util.RuntimeOptions) (_result *GetAlbumDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlbumId)) {
		query["AlbumId"] = request.AlbumId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlbumDetailById"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getAlbumDetailById"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlbumDetailByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAligenieUserInfo(request *GetAligenieUserInfoRequest) (_result *GetAligenieUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAligenieUserInfoHeaders{}
	_result = &GetAligenieUserInfoResponse{}
	_body, _err := client.GetAligenieUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAligenieUserInfoWithOptions(request *GetAligenieUserInfoRequest, headers *GetAligenieUserInfoHeaders, runtime *util.RuntimeOptions) (_result *GetAligenieUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoginStateAccessToken)) {
		query["LoginStateAccessToken"] = request.LoginStateAccessToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAligenieUserInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getAligenieUserInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAligenieUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCodeEnhance(request *GetCodeEnhanceRequest) (_result *GetCodeEnhanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCodeEnhanceHeaders{}
	_result = &GetCodeEnhanceResponse{}
	_body, _err := client.GetCodeEnhanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCodeEnhanceWithOptions(tmpReq *GetCodeEnhanceRequest, headers *GetCodeEnhanceHeaders, runtime *util.RuntimeOptions) (_result *GetCodeEnhanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetCodeEnhanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ChannelInfo))) {
		request.ChannelInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ChannelInfo), tea.String("ChannelInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelInfoShrink)) {
		query["ChannelInfo"] = request.ChannelInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCodeEnhance"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getCodeEnhance"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCodeEnhanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContent(request *GetContentRequest) (_result *GetContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetContentHeaders{}
	_result = &GetContentResponse{}
	_body, _err := client.GetContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContentWithOptions(request *GetContentRequest, headers *GetContentHeaders, runtime *util.RuntimeOptions) (_result *GetContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetContent"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/GetContent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCurrentPlayingItem(request *GetCurrentPlayingItemRequest) (_result *GetCurrentPlayingItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCurrentPlayingItemHeaders{}
	_result = &GetCurrentPlayingItemResponse{}
	_body, _err := client.GetCurrentPlayingItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCurrentPlayingItemWithOptions(tmpReq *GetCurrentPlayingItemRequest, headers *GetCurrentPlayingItemHeaders, runtime *util.RuntimeOptions) (_result *GetCurrentPlayingItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetCurrentPlayingItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCurrentPlayingItem"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/GetCurrentPlayingItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCurrentPlayingItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCurrentPlayingList(request *GetCurrentPlayingListRequest) (_result *GetCurrentPlayingListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCurrentPlayingListHeaders{}
	_result = &GetCurrentPlayingListResponse{}
	_body, _err := client.GetCurrentPlayingListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCurrentPlayingListWithOptions(tmpReq *GetCurrentPlayingListRequest, headers *GetCurrentPlayingListHeaders, runtime *util.RuntimeOptions) (_result *GetCurrentPlayingListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetCurrentPlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenQueryPlayListRequest))) {
		request.OpenQueryPlayListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenQueryPlayListRequest), tea.String("OpenQueryPlayListRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenQueryPlayListRequestShrink)) {
		body["OpenQueryPlayListRequest"] = request.OpenQueryPlayListRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCurrentPlayingList"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/GetCurrentPlayingList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCurrentPlayingListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceBasicInfo(request *GetDeviceBasicInfoRequest) (_result *GetDeviceBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceBasicInfoHeaders{}
	_result = &GetDeviceBasicInfoResponse{}
	_body, _err := client.GetDeviceBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceBasicInfoWithOptions(tmpReq *GetDeviceBasicInfoRequest, headers *GetDeviceBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDeviceBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceBasicInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getDeviceBasicInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceIdByIdentity(request *GetDeviceIdByIdentityRequest) (_result *GetDeviceIdByIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceIdByIdentityHeaders{}
	_result = &GetDeviceIdByIdentityResponse{}
	_body, _err := client.GetDeviceIdByIdentityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceIdByIdentityWithOptions(request *GetDeviceIdByIdentityRequest, headers *GetDeviceIdByIdentityHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceIdByIdentityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeKey)) {
		query["EncodeKey"] = request.EncodeKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityId)) {
		query["IdentityId"] = request.IdentityId
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		query["IdentityType"] = request.IdentityType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceIdByIdentity"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getDeviceIdByIdentity"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceIdByIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceSetting(request *GetDeviceSettingRequest) (_result *GetDeviceSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceSettingHeaders{}
	_result = &GetDeviceSettingResponse{}
	_body, _err := client.GetDeviceSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceSettingWithOptions(tmpReq *GetDeviceSettingRequest, headers *GetDeviceSettingHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceSettingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDeviceSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Keys)) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, tea.String("Keys"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KeysShrink)) {
		query["Keys"] = request.KeysShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceSetting"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getDeviceSetting"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceStatusDetail(request *GetDeviceStatusDetailRequest) (_result *GetDeviceStatusDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceStatusDetailHeaders{}
	_result = &GetDeviceStatusDetailResponse{}
	_body, _err := client.GetDeviceStatusDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceStatusDetailWithOptions(tmpReq *GetDeviceStatusDetailRequest, headers *GetDeviceStatusDetailHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceStatusDetailResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDeviceStatusDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Keys)) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, tea.String("Keys"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KeysShrink)) {
		query["Keys"] = request.KeysShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceStatusDetail"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getDeviceStatusDetail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceStatusDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceStatusInfo(request *GetDeviceStatusInfoRequest) (_result *GetDeviceStatusInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceStatusInfoHeaders{}
	_result = &GetDeviceStatusInfoResponse{}
	_body, _err := client.GetDeviceStatusInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceStatusInfoWithOptions(tmpReq *GetDeviceStatusInfoRequest, headers *GetDeviceStatusInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceStatusInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDeviceStatusInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceStatusInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getDeviceStatusInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceStatusInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceTag(request *GetDeviceTagRequest) (_result *GetDeviceTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceTagHeaders{}
	_result = &GetDeviceTagResponse{}
	_body, _err := client.GetDeviceTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceTagWithOptions(tmpReq *GetDeviceTagRequest, headers *GetDeviceTagHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceTagResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDeviceTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceTag"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getDeviceTag"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScheduleTask(request *GetScheduleTaskRequest) (_result *GetScheduleTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetScheduleTaskHeaders{}
	_result = &GetScheduleTaskResponse{}
	_body, _err := client.GetScheduleTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScheduleTaskWithOptions(tmpReq *GetScheduleTaskRequest, headers *GetScheduleTaskHeaders, runtime *util.RuntimeOptions) (_result *GetScheduleTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetScheduleTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScheduleTask"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/GetScheduleTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUnreadMessageCount(request *GetUnreadMessageCountRequest) (_result *GetUnreadMessageCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUnreadMessageCountHeaders{}
	_result = &GetUnreadMessageCountResponse{}
	_body, _err := client.GetUnreadMessageCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUnreadMessageCountWithOptions(tmpReq *GetUnreadMessageCountRequest, headers *GetUnreadMessageCountHeaders, runtime *util.RuntimeOptions) (_result *GetUnreadMessageCountResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetUnreadMessageCountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUnreadMessageCount"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getUnreadMessageCount"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUnreadMessageCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserByDeviceId(request *GetUserByDeviceIdRequest) (_result *GetUserByDeviceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserByDeviceIdHeaders{}
	_result = &GetUserByDeviceIdResponse{}
	_body, _err := client.GetUserByDeviceIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserByDeviceIdWithOptions(tmpReq *GetUserByDeviceIdRequest, headers *GetUserByDeviceIdHeaders, runtime *util.RuntimeOptions) (_result *GetUserByDeviceIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetUserByDeviceIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserByDeviceId"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getUserByDeviceId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserByDeviceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWeather(request *GetWeatherRequest) (_result *GetWeatherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWeatherHeaders{}
	_result = &GetWeatherResponse{}
	_body, _err := client.GetWeatherWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWeatherWithOptions(tmpReq *GetWeatherRequest, headers *GetWeatherHeaders, runtime *util.RuntimeOptions) (_result *GetWeatherResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetWeatherShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWeather"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/GetWeather"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWeatherResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndexControlPlayingList(request *IndexControlPlayingListRequest) (_result *IndexControlPlayingListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndexControlPlayingListHeaders{}
	_result = &IndexControlPlayingListResponse{}
	_body, _err := client.IndexControlPlayingListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndexControlPlayingListWithOptions(tmpReq *IndexControlPlayingListRequest, headers *IndexControlPlayingListHeaders, runtime *util.RuntimeOptions) (_result *IndexControlPlayingListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &IndexControlPlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenIndexControlRequest))) {
		request.OpenIndexControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenIndexControlRequest), tea.String("OpenIndexControlRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenIndexControlRequestShrink)) {
		body["OpenIndexControlRequest"] = request.OpenIndexControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IndexControlPlayingList"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/IndexControlPlayingList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IndexControlPlayingListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlarms(request *ListAlarmsRequest) (_result *ListAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAlarmsHeaders{}
	_result = &ListAlarmsResponse{}
	_body, _err := client.ListAlarmsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlarmsWithOptions(tmpReq *ListAlarmsRequest, headers *ListAlarmsHeaders, runtime *util.RuntimeOptions) (_result *ListAlarmsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAlarmsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarms"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listAlarm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlbumDetail(request *ListAlbumDetailRequest) (_result *ListAlbumDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAlbumDetailHeaders{}
	_result = &ListAlbumDetailResponse{}
	_body, _err := client.ListAlbumDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlbumDetailWithOptions(request *ListAlbumDetailRequest, headers *ListAlbumDetailHeaders, runtime *util.RuntimeOptions) (_result *ListAlbumDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlbumDetail"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ListAlbumDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlbumDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlbumIsAdded(request *ListAlbumIsAddedRequest) (_result *ListAlbumIsAddedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAlbumIsAddedHeaders{}
	_result = &ListAlbumIsAddedResponse{}
	_body, _err := client.ListAlbumIsAddedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlbumIsAddedWithOptions(tmpReq *ListAlbumIsAddedRequest, headers *ListAlbumIsAddedHeaders, runtime *util.RuntimeOptions) (_result *ListAlbumIsAddedResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAlbumIsAddedShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlbumIdList)) {
		request.AlbumIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlbumIdList, tea.String("AlbumIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlbumIdListShrink)) {
		query["AlbumIdList"] = request.AlbumIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlbumIsAdded"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listAlbumIsAdded"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlbumIsAddedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCateContent(request *ListCateContentRequest) (_result *ListCateContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCateContentHeaders{}
	_result = &ListCateContentResponse{}
	_body, _err := client.ListCateContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCateContentWithOptions(tmpReq *ListCateContentRequest, headers *ListCateContentHeaders, runtime *util.RuntimeOptions) (_result *ListCateContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCateContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Request))) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Request), tea.String("Request"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCateContent"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ListCateContent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCateContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCateInfo(request *ListCateInfoRequest) (_result *ListCateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCateInfoHeaders{}
	_result = &ListCateInfoResponse{}
	_body, _err := client.ListCateInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCateInfoWithOptions(request *ListCateInfoRequest, headers *ListCateInfoHeaders, runtime *util.RuntimeOptions) (_result *ListCateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCateInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ListCateInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommonCateFirstFloor(request *ListCommonCateFirstFloorRequest) (_result *ListCommonCateFirstFloorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCommonCateFirstFloorHeaders{}
	_result = &ListCommonCateFirstFloorResponse{}
	_body, _err := client.ListCommonCateFirstFloorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommonCateFirstFloorWithOptions(request *ListCommonCateFirstFloorRequest, headers *ListCommonCateFirstFloorHeaders, runtime *util.RuntimeOptions) (_result *ListCommonCateFirstFloorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommonCateFirstFloor"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ListCommonCateFirstFloor"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommonCateFirstFloorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommonCateSecondFloor(request *ListCommonCateSecondFloorRequest) (_result *ListCommonCateSecondFloorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCommonCateSecondFloorHeaders{}
	_result = &ListCommonCateSecondFloorResponse{}
	_body, _err := client.ListCommonCateSecondFloorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommonCateSecondFloorWithOptions(request *ListCommonCateSecondFloorRequest, headers *ListCommonCateSecondFloorHeaders, runtime *util.RuntimeOptions) (_result *ListCommonCateSecondFloorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParentCateId)) {
		query["ParentCateId"] = request.ParentCateId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommonCateSecondFloor"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ListCommonCateSecondFloor"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommonCateSecondFloorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceBasicInfo(request *ListDeviceBasicInfoRequest) (_result *ListDeviceBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeviceBasicInfoHeaders{}
	_result = &ListDeviceBasicInfoResponse{}
	_body, _err := client.ListDeviceBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceBasicInfoWithOptions(tmpReq *ListDeviceBasicInfoRequest, headers *ListDeviceBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *ListDeviceBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDeviceBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfos))) {
		request.DeviceInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfos), tea.String("DeviceInfos"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfosShrink)) {
		query["DeviceInfos"] = request.DeviceInfosShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceBasicInfo"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listDeviceBasicInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceByUserId(request *ListDeviceByUserIdRequest) (_result *ListDeviceByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeviceByUserIdHeaders{}
	_result = &ListDeviceByUserIdResponse{}
	_body, _err := client.ListDeviceByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceByUserIdWithOptions(tmpReq *ListDeviceByUserIdRequest, headers *ListDeviceByUserIdHeaders, runtime *util.RuntimeOptions) (_result *ListDeviceByUserIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDeviceByUserIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceByUserId"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listDeviceByUserId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceByUserIdAndChanel(request *ListDeviceByUserIdAndChanelRequest) (_result *ListDeviceByUserIdAndChanelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeviceByUserIdAndChanelHeaders{}
	_result = &ListDeviceByUserIdAndChanelResponse{}
	_body, _err := client.ListDeviceByUserIdAndChanelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceByUserIdAndChanelWithOptions(tmpReq *ListDeviceByUserIdAndChanelRequest, headers *ListDeviceByUserIdAndChanelHeaders, runtime *util.RuntimeOptions) (_result *ListDeviceByUserIdAndChanelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDeviceByUserIdAndChanelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ChannelInfo))) {
		request.ChannelInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ChannelInfo), tea.String("ChannelInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelInfoShrink)) {
		query["ChannelInfo"] = request.ChannelInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceByUserIdAndChanel"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listDeviceByUserIdAndChanel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceByUserIdAndChanelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceIdByIdentities(request *ListDeviceIdByIdentitiesRequest) (_result *ListDeviceIdByIdentitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeviceIdByIdentitiesHeaders{}
	_result = &ListDeviceIdByIdentitiesResponse{}
	_body, _err := client.ListDeviceIdByIdentitiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceIdByIdentitiesWithOptions(tmpReq *ListDeviceIdByIdentitiesRequest, headers *ListDeviceIdByIdentitiesHeaders, runtime *util.RuntimeOptions) (_result *ListDeviceIdByIdentitiesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDeviceIdByIdentitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IdentityIds)) {
		request.IdentityIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IdentityIds, tea.String("IdentityIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeKey)) {
		query["EncodeKey"] = request.EncodeKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityIdsShrink)) {
		query["IdentityIds"] = request.IdentityIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		query["IdentityType"] = request.IdentityType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceIdByIdentities"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listDeviceIdByIdentities"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceIdByIdentitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMusic(request *ListMusicRequest) (_result *ListMusicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMusicHeaders{}
	_result = &ListMusicResponse{}
	_body, _err := client.ListMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMusicWithOptions(tmpReq *ListMusicRequest, headers *ListMusicHeaders, runtime *util.RuntimeOptions) (_result *ListMusicResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMusicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMusic"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listMusic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPlayHistory(request *ListPlayHistoryRequest) (_result *ListPlayHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPlayHistoryHeaders{}
	_result = &ListPlayHistoryResponse{}
	_body, _err := client.ListPlayHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPlayHistoryWithOptions(tmpReq *ListPlayHistoryRequest, headers *ListPlayHistoryHeaders, runtime *util.RuntimeOptions) (_result *ListPlayHistoryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPlayHistoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Request))) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Request), tea.String("Request"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPlayHistory"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ListPlayHistory"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPlayHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecommendContent(request *ListRecommendContentRequest) (_result *ListRecommendContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRecommendContentHeaders{}
	_result = &ListRecommendContentResponse{}
	_body, _err := client.ListRecommendContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecommendContentWithOptions(tmpReq *ListRecommendContentRequest, headers *ListRecommendContentHeaders, runtime *util.RuntimeOptions) (_result *ListRecommendContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListRecommendContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Request))) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Request), tea.String("Request"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecommendContent"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ListRecommendContent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecommendContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSub(request *ListSubRequest) (_result *ListSubResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubHeaders{}
	_result = &ListSubResponse{}
	_body, _err := client.ListSubWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubWithOptions(tmpReq *ListSubRequest, headers *ListSubHeaders, runtime *util.RuntimeOptions) (_result *ListSubResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSubShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Page))) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Page), tea.String("Page"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		query["Page"] = request.PageShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSub"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listSub"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubAlbum(request *ListSubAlbumRequest) (_result *ListSubAlbumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubAlbumHeaders{}
	_result = &ListSubAlbumResponse{}
	_body, _err := client.ListSubAlbumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubAlbumWithOptions(tmpReq *ListSubAlbumRequest, headers *ListSubAlbumHeaders, runtime *util.RuntimeOptions) (_result *ListSubAlbumResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSubAlbumShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.QuerySubscriptionAlbumRequest))) {
		request.QuerySubscriptionAlbumRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.QuerySubscriptionAlbumRequest), tea.String("QuerySubscriptionAlbumRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.QuerySubscriptionAlbumRequestShrink)) {
		query["QuerySubscriptionAlbumRequest"] = request.QuerySubscriptionAlbumRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubAlbum"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listSubAlbum"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubAlbumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubscriptionAlbumCategory(request *ListSubscriptionAlbumCategoryRequest) (_result *ListSubscriptionAlbumCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubscriptionAlbumCategoryHeaders{}
	_result = &ListSubscriptionAlbumCategoryResponse{}
	_body, _err := client.ListSubscriptionAlbumCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubscriptionAlbumCategoryWithOptions(request *ListSubscriptionAlbumCategoryRequest, headers *ListSubscriptionAlbumCategoryHeaders, runtime *util.RuntimeOptions) (_result *ListSubscriptionAlbumCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		query["CategoryName"] = request.CategoryName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubscriptionAlbumCategory"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listSubscriptionAlbumCategory"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubscriptionAlbumCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserMessage(request *ListUserMessageRequest) (_result *ListUserMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUserMessageHeaders{}
	_result = &ListUserMessageResponse{}
	_body, _err := client.ListUserMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserMessageWithOptions(tmpReq *ListUserMessageRequest, headers *ListUserMessageHeaders, runtime *util.RuntimeOptions) (_result *ListUserMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListUserMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeforeTime)) {
		query["BeforeTime"] = request.BeforeTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserMessage"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/listUserMessage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PlayAndPauseControl(request *PlayAndPauseControlRequest) (_result *PlayAndPauseControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PlayAndPauseControlHeaders{}
	_result = &PlayAndPauseControlResponse{}
	_body, _err := client.PlayAndPauseControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PlayAndPauseControlWithOptions(tmpReq *PlayAndPauseControlRequest, headers *PlayAndPauseControlHeaders, runtime *util.RuntimeOptions) (_result *PlayAndPauseControlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PlayAndPauseControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenPlayAndPauseControlParam))) {
		request.OpenPlayAndPauseControlParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenPlayAndPauseControlParam), tea.String("OpenPlayAndPauseControlParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenPlayAndPauseControlParamShrink)) {
		body["OpenPlayAndPauseControlParam"] = request.OpenPlayAndPauseControlParamShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PlayAndPauseControl"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/PlayAndPauseControl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PlayAndPauseControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PlayModeControl(request *PlayModeControlRequest) (_result *PlayModeControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PlayModeControlHeaders{}
	_result = &PlayModeControlResponse{}
	_body, _err := client.PlayModeControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PlayModeControlWithOptions(tmpReq *PlayModeControlRequest, headers *PlayModeControlHeaders, runtime *util.RuntimeOptions) (_result *PlayModeControlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PlayModeControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenPlayModeControlRequest))) {
		request.OpenPlayModeControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenPlayModeControlRequest), tea.String("OpenPlayModeControlRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenPlayModeControlRequestShrink)) {
		body["OpenPlayModeControlRequest"] = request.OpenPlayModeControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PlayModeControl"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/PlayModeControl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PlayModeControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreviousAndNextControl(request *PreviousAndNextControlRequest) (_result *PreviousAndNextControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PreviousAndNextControlHeaders{}
	_result = &PreviousAndNextControlResponse{}
	_body, _err := client.PreviousAndNextControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreviousAndNextControlWithOptions(tmpReq *PreviousAndNextControlRequest, headers *PreviousAndNextControlHeaders, runtime *util.RuntimeOptions) (_result *PreviousAndNextControlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PreviousAndNextControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenControlPlayingListRequest))) {
		request.OpenControlPlayingListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenControlPlayingListRequest), tea.String("OpenControlPlayingListRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenControlPlayingListRequestShrink)) {
		body["OpenControlPlayingListRequest"] = request.OpenControlPlayingListRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PreviousAndNextControl"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/PreviousAndNextControl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PreviousAndNextControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProgressControl(request *ProgressControlRequest) (_result *ProgressControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProgressControlHeaders{}
	_result = &ProgressControlResponse{}
	_body, _err := client.ProgressControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProgressControlWithOptions(tmpReq *ProgressControlRequest, headers *ProgressControlHeaders, runtime *util.RuntimeOptions) (_result *ProgressControlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ProgressControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenProgressControlRequest))) {
		request.OpenProgressControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenProgressControlRequest), tea.String("OpenProgressControlRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenProgressControlRequestShrink)) {
		body["OpenProgressControlRequest"] = request.OpenProgressControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ProgressControl"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/ProgressControl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ProgressControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMusicType(request *QueryMusicTypeRequest) (_result *QueryMusicTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMusicTypeHeaders{}
	_result = &QueryMusicTypeResponse{}
	_body, _err := client.QueryMusicTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMusicTypeWithOptions(tmpReq *QueryMusicTypeRequest, headers *QueryMusicTypeHeaders, runtime *util.RuntimeOptions) (_result *QueryMusicTypeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryMusicTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMusicType"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/queryMusicType"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMusicTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReadMessage(request *ReadMessageRequest) (_result *ReadMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReadMessageHeaders{}
	_result = &ReadMessageResponse{}
	_body, _err := client.ReadMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReadMessageWithOptions(tmpReq *ReadMessageRequest, headers *ReadMessageHeaders, runtime *util.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReadMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadMessage"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/readMessage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScanCodeBind(request *ScanCodeBindRequest) (_result *ScanCodeBindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ScanCodeBindHeaders{}
	_result = &ScanCodeBindResponse{}
	_body, _err := client.ScanCodeBindWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScanCodeBindWithOptions(tmpReq *ScanCodeBindRequest, headers *ScanCodeBindHeaders, runtime *util.RuntimeOptions) (_result *ScanCodeBindResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ScanCodeBindShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.BindReq))) {
		request.BindReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.BindReq), tea.String("BindReq"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindReqShrink)) {
		body["BindReq"] = request.BindReqShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScanCodeBind"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/scanCode"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScanCodeBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScgSearch(request *ScgSearchRequest) (_result *ScgSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ScgSearchHeaders{}
	_result = &ScgSearchResponse{}
	_body, _err := client.ScgSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScgSearchWithOptions(tmpReq *ScgSearchRequest, headers *ScgSearchHeaders, runtime *util.RuntimeOptions) (_result *ScgSearchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ScgSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ScgFilter))) {
		request.ScgFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ScgFilter), tea.String("ScgFilter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScgFilterShrink)) {
		query["ScgFilter"] = request.ScgFilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TopicId)) {
		query["TopicId"] = request.TopicId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScgSearch"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/scgSearch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ScgSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchContent(request *SearchContentRequest) (_result *SearchContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchContentHeaders{}
	_result = &SearchContentResponse{}
	_body, _err := client.SearchContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchContentWithOptions(tmpReq *SearchContentRequest, headers *SearchContentHeaders, runtime *util.RuntimeOptions) (_result *SearchContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Request))) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Request), tea.String("Request"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchContent"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/SearchContent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(tmpReq *SendMessageRequest, headers *SendMessageHeaders, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/sendMessage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDeviceSetting(request *SetDeviceSettingRequest) (_result *SetDeviceSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetDeviceSettingHeaders{}
	_result = &SetDeviceSettingResponse{}
	_body, _err := client.SetDeviceSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDeviceSettingWithOptions(tmpReq *SetDeviceSettingRequest, headers *SetDeviceSettingHeaders, runtime *util.RuntimeOptions) (_result *SetDeviceSettingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetDeviceSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDeviceSetting"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/setDeviceSetting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDeviceSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindAligenieUser(request *UnbindAligenieUserRequest) (_result *UnbindAligenieUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnbindAligenieUserHeaders{}
	_result = &UnbindAligenieUserResponse{}
	_body, _err := client.UnbindAligenieUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindAligenieUserWithOptions(request *UnbindAligenieUserRequest, headers *UnbindAligenieUserHeaders, runtime *util.RuntimeOptions) (_result *UnbindAligenieUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoginStateAccessToken)) {
		body["LoginStateAccessToken"] = request.LoginStateAccessToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindAligenieUser"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/unbindAligenieUser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindAligenieUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDevice(request *UnbindDeviceRequest) (_result *UnbindDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnbindDeviceHeaders{}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.UnbindDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindDeviceWithOptions(tmpReq *UnbindDeviceRequest, headers *UnbindDeviceHeaders, runtime *util.RuntimeOptions) (_result *UnbindDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UnbindDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindDevice"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/unbindDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlarm(request *UpdateAlarmRequest) (_result *UpdateAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAlarmHeaders{}
	_result = &UpdateAlarmResponse{}
	_body, _err := client.UpdateAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlarmWithOptions(tmpReq *UpdateAlarmRequest, headers *UpdateAlarmHeaders, runtime *util.RuntimeOptions) (_result *UpdateAlarmResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlarm"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/updateAlarm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
