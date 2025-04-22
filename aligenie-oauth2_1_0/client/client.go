// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreatePlayingListRequest struct {
	// This parameter is required.
	DeviceInfo *CreatePlayingListRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	OpenCreatePlayingListRequest *CreatePlayingListRequestOpenCreatePlayingListRequest `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty" type:"Struct"`
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

type CreatePlayingListRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UC_CLIENT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1
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
	// This parameter is required.
	ContentList []*CreatePlayingListRequestOpenCreatePlayingListRequestContentList `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// content
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// {}
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// true
	NeedAlbumContinued *bool `json:"NeedAlbumContinued,omitempty" xml:"NeedAlbumContinued,omitempty"`
	// example:
	//
	// default
	PlayFrom *string `json:"PlayFrom,omitempty" xml:"PlayFrom,omitempty"`
	// example:
	//
	// Normal
	PlayMode *string `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ximalayaH5
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

type CreatePlayingListShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenCreatePlayingListRequestShrink *string `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty"`
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

type CreatePlayingListResponseBody struct {
	// example:
	//
	// 10002398812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePlayingListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBody) SetRequestId(v string) *CreatePlayingListResponseBody {
	s.RequestId = &v
	return s
}

type CreatePlayingListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePlayingListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ExecuteSceneRequest struct {
	// example:
	//
	// a84a55aa410e460a9ac753570c76fecc
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ExecuteSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSceneRequest) GoString() string {
	return s.String()
}

func (s *ExecuteSceneRequest) SetSceneId(v string) *ExecuteSceneRequest {
	s.SceneId = &v
	return s
}

type ExecuteSceneResponseBody struct {
	// example:
	//
	// 191C79AD-F9F9-531E-B8C1-73DF6433B920
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteSceneResponseBody) SetRequestId(v string) *ExecuteSceneResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSceneResponse) GoString() string {
	return s.String()
}

func (s *ExecuteSceneResponse) SetHeaders(v map[string]*string) *ExecuteSceneResponse {
	s.Headers = v
	return s
}

func (s *ExecuteSceneResponse) SetStatusCode(v int32) *ExecuteSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteSceneResponse) SetBody(v *ExecuteSceneResponseBody) *ExecuteSceneResponse {
	s.Body = v
	return s
}

type GetSceneListResponseBody struct {
	// example:
	//
	// 435CF567-58DC-5761-AFA8-650772602E2D
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneList []*GetSceneListResponseBodySceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
}

func (s GetSceneListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneListResponseBody) SetRequestId(v string) *GetSceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSceneListResponseBody) SetSceneList(v []*GetSceneListResponseBodySceneList) *GetSceneListResponseBody {
	s.SceneList = v
	return s
}

type GetSceneListResponseBodySceneList struct {
	// example:
	//
	// 840960b85c3c48e0bd7260c1718295fd
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s GetSceneListResponseBodySceneList) String() string {
	return tea.Prettify(s)
}

func (s GetSceneListResponseBodySceneList) GoString() string {
	return s.String()
}

func (s *GetSceneListResponseBodySceneList) SetSceneId(v string) *GetSceneListResponseBodySceneList {
	s.SceneId = &v
	return s
}

func (s *GetSceneListResponseBodySceneList) SetSceneName(v string) *GetSceneListResponseBodySceneList {
	s.SceneName = &v
	return s
}

type GetSceneListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneListResponse) GoString() string {
	return s.String()
}

func (s *GetSceneListResponse) SetHeaders(v map[string]*string) *GetSceneListResponse {
	s.Headers = v
	return s
}

func (s *GetSceneListResponse) SetStatusCode(v int32) *GetSceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneListResponse) SetBody(v *GetSceneListResponseBody) *GetSceneListResponse {
	s.Body = v
	return s
}

type GetUserBasicInfoResponseBody struct {
	// example:
	//
	// https://xxxxxx
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// xxxxxx
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// 3hPPBd9YuhfJQCzZ/07AAWdoO3K8zCb/KAqW96zPHXPiFkzjB/JfcWuuFHQQDaGZ4wVbNMV6wYuj075p/rhVLg==
	OpenId *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	// example:
	//
	// 4070039E-5822-1F32-9295-1D2883E48BA5
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UnionIds  []*GetUserBasicInfoResponseBodyUnionIds `json:"UnionIds,omitempty" xml:"UnionIds,omitempty" type:"Repeated"`
}

func (s GetUserBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserBasicInfoResponseBody) SetAvatarUrl(v string) *GetUserBasicInfoResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetNickname(v string) *GetUserBasicInfoResponseBody {
	s.Nickname = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetOpenId(v string) *GetUserBasicInfoResponseBody {
	s.OpenId = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetRequestId(v string) *GetUserBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetUnionIds(v []*GetUserBasicInfoResponseBodyUnionIds) *GetUserBasicInfoResponseBody {
	s.UnionIds = v
	return s
}

type GetUserBasicInfoResponseBodyUnionIds struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	UnionId        *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetUserBasicInfoResponseBodyUnionIds) String() string {
	return tea.Prettify(s)
}

func (s GetUserBasicInfoResponseBodyUnionIds) GoString() string {
	return s.String()
}

func (s *GetUserBasicInfoResponseBodyUnionIds) SetOrganizationId(v string) *GetUserBasicInfoResponseBodyUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *GetUserBasicInfoResponseBodyUnionIds) SetUnionId(v string) *GetUserBasicInfoResponseBodyUnionIds {
	s.UnionId = &v
	return s
}

type GetUserBasicInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserBasicInfoResponse) SetHeaders(v map[string]*string) *GetUserBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserBasicInfoResponse) SetStatusCode(v int32) *GetUserBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserBasicInfoResponse) SetBody(v *GetUserBasicInfoResponseBody) *GetUserBasicInfoResponse {
	s.Body = v
	return s
}

type GetUserPhoneResponseBody struct {
	// example:
	//
	// 18612345678
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CEADB586-51CB-1B6B-95BD-AB85A7A08E97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPhoneResponseBody) SetPhone(v string) *GetUserPhoneResponseBody {
	s.Phone = &v
	return s
}

func (s *GetUserPhoneResponseBody) SetRequestId(v string) *GetUserPhoneResponseBody {
	s.RequestId = &v
	return s
}

type GetUserPhoneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserPhoneResponse) GoString() string {
	return s.String()
}

func (s *GetUserPhoneResponse) SetHeaders(v map[string]*string) *GetUserPhoneResponse {
	s.Headers = v
	return s
}

func (s *GetUserPhoneResponse) SetStatusCode(v int32) *GetUserPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserPhoneResponse) SetBody(v *GetUserPhoneResponseBody) *GetUserPhoneResponse {
	s.Body = v
	return s
}

type OAuth2RevocationEndpointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s OAuth2RevocationEndpointHeaders) String() string {
	return tea.Prettify(s)
}

func (s OAuth2RevocationEndpointHeaders) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointHeaders) SetCommonHeaders(v map[string]*string) *OAuth2RevocationEndpointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OAuth2RevocationEndpointHeaders) SetXAcsAligenieAccessToken(v string) *OAuth2RevocationEndpointHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *OAuth2RevocationEndpointHeaders) SetAuthorization(v string) *OAuth2RevocationEndpointHeaders {
	s.Authorization = &v
	return s
}

type OAuth2RevocationEndpointRequest struct {
	// example:
	//
	// UJMiksSwuMJvwXrJLULMykSw6qZ6VqaxOkN4qd5cW1Q4HhsLxvUR5xVOIv1WB3br5LoP20lPa8xiYLSMbt8JqHACXdSdw7fNkhRTIHnadxWW5jfDg7BELUB0FcFfPiv0
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// refresh_token
	TokenTypeHint *string `json:"TokenTypeHint,omitempty" xml:"TokenTypeHint,omitempty"`
}

func (s OAuth2RevocationEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s OAuth2RevocationEndpointRequest) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointRequest) SetToken(v string) *OAuth2RevocationEndpointRequest {
	s.Token = &v
	return s
}

func (s *OAuth2RevocationEndpointRequest) SetTokenTypeHint(v string) *OAuth2RevocationEndpointRequest {
	s.TokenTypeHint = &v
	return s
}

type OAuth2RevocationEndpointResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 4070039E-5822-1F32-9295-1D2883E48BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OAuth2RevocationEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OAuth2RevocationEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointResponseBody) SetRequestId(v string) *OAuth2RevocationEndpointResponseBody {
	s.RequestId = &v
	return s
}

type OAuth2RevocationEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OAuth2RevocationEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OAuth2RevocationEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s OAuth2RevocationEndpointResponse) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointResponse) SetHeaders(v map[string]*string) *OAuth2RevocationEndpointResponse {
	s.Headers = v
	return s
}

func (s *OAuth2RevocationEndpointResponse) SetStatusCode(v int32) *OAuth2RevocationEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *OAuth2RevocationEndpointResponse) SetBody(v *OAuth2RevocationEndpointResponseBody) *OAuth2RevocationEndpointResponse {
	s.Body = v
	return s
}

type OAuth2TokenEndpointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s OAuth2TokenEndpointHeaders) String() string {
	return tea.Prettify(s)
}

func (s OAuth2TokenEndpointHeaders) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointHeaders) SetCommonHeaders(v map[string]*string) *OAuth2TokenEndpointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OAuth2TokenEndpointHeaders) SetXAcsAligenieAccessToken(v string) *OAuth2TokenEndpointHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *OAuth2TokenEndpointHeaders) SetAuthorization(v string) *OAuth2TokenEndpointHeaders {
	s.Authorization = &v
	return s
}

type OAuth2TokenEndpointRequest struct {
	// example:
	//
	// rf3mi4JOU-xRIX2zEuRLHi-U9mPnvISeSphbwiBHJ5mEKZtG-xJsbBWrq8RmhQEPRYh0JOd3DaS_VZ90soD_YrsT4OBtgD06DmdIKL2_5KFfI6p_SjXX2-UMJuGfXDkB
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// authorization_code
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// example:
	//
	// https://xxx.xxx.com/xxx
	RedirectUri *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
	// example:
	//
	// zsEcmaUeb8-NZW4IIUDD7qdgBNflrj6fH8BXJYbW9iXihZTgvbcr1_utC9p5HJLn_lXVwhfivBTgUQZBCGvGl5lxqaxFhmFtt-OrBduFQKL9x8p2lpEMKlxuKHZZZJ3A
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s OAuth2TokenEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s OAuth2TokenEndpointRequest) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointRequest) SetCode(v string) *OAuth2TokenEndpointRequest {
	s.Code = &v
	return s
}

func (s *OAuth2TokenEndpointRequest) SetGrantType(v string) *OAuth2TokenEndpointRequest {
	s.GrantType = &v
	return s
}

func (s *OAuth2TokenEndpointRequest) SetRedirectUri(v string) *OAuth2TokenEndpointRequest {
	s.RedirectUri = &v
	return s
}

func (s *OAuth2TokenEndpointRequest) SetRefreshToken(v string) *OAuth2TokenEndpointRequest {
	s.RefreshToken = &v
	return s
}

type OAuth2TokenEndpointResponseBody struct {
	// example:
	//
	// UJMiksSwuMJvwXrJLULMykSw6qZ6VqaxOkN4qd5cW1Q4HhsLxvUR5xVOIv1WB3br5LoP20lPa8xiYLSMbt8JqHACXdSdw7fNkhRTIHnadxWW5jfDg7BELUB0FcFfPiv0
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// 604799
	ExpiresIn *int64 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// example:
	//
	// zsEcmaUeb8-NZW4IIUDD7qdgBNflrj6fH8BXJYbW9iXihZTgvbcr1_utC9p5HJLn_lXVwhfivBTgUQZBCGvGl5lxqaxFhmFtt-OrBduFQKL9x8p2lpEMKlxuKHZZZJ3A
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// example:
	//
	// 4070039E-5822-1F32-9295-1D2883E48BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// aligenie:user:basic:read aligenie:iot:scene:read
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// Bearer
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s OAuth2TokenEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OAuth2TokenEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointResponseBody) SetAccessToken(v string) *OAuth2TokenEndpointResponseBody {
	s.AccessToken = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetExpiresIn(v int64) *OAuth2TokenEndpointResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetRefreshToken(v string) *OAuth2TokenEndpointResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetRequestId(v string) *OAuth2TokenEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetScope(v string) *OAuth2TokenEndpointResponseBody {
	s.Scope = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetTokenType(v string) *OAuth2TokenEndpointResponseBody {
	s.TokenType = &v
	return s
}

type OAuth2TokenEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OAuth2TokenEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OAuth2TokenEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s OAuth2TokenEndpointResponse) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointResponse) SetHeaders(v map[string]*string) *OAuth2TokenEndpointResponse {
	s.Headers = v
	return s
}

func (s *OAuth2TokenEndpointResponse) SetStatusCode(v int32) *OAuth2TokenEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *OAuth2TokenEndpointResponse) SetBody(v *OAuth2TokenEndpointResponseBody) *OAuth2TokenEndpointResponse {
	s.Body = v
	return s
}

type PushDeviceNotificationRequest struct {
	TenantInfo *PushDeviceNotificationRequestTenantInfo `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
	Body       *PushDeviceNotificationRequestBody       `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s PushDeviceNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceNotificationRequest) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequest) SetTenantInfo(v *PushDeviceNotificationRequestTenantInfo) *PushDeviceNotificationRequest {
	s.TenantInfo = v
	return s
}

func (s *PushDeviceNotificationRequest) SetBody(v *PushDeviceNotificationRequestBody) *PushDeviceNotificationRequest {
	s.Body = v
	return s
}

type PushDeviceNotificationRequestTenantInfo struct {
	// example:
	//
	// 12797******304102
	SubjectId *string `json:"SubjectId,omitempty" xml:"SubjectId,omitempty"`
}

func (s PushDeviceNotificationRequestTenantInfo) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceNotificationRequestTenantInfo) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequestTenantInfo) SetSubjectId(v string) *PushDeviceNotificationRequestTenantInfo {
	s.SubjectId = &v
	return s
}

type PushDeviceNotificationRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1923792******8R7392
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// false
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iU81*****G9elJ
	MessageTemplateId *string `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	// example:
	//
	// 29837******2938
	OrganizationId *string            `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	PlaceHolder    map[string]*string `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	// This parameter is required.
	SendTarget *PushDeviceNotificationRequestBodySendTarget `json:"SendTarget,omitempty" xml:"SendTarget,omitempty" type:"Struct"`
}

func (s PushDeviceNotificationRequestBody) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceNotificationRequestBody) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequestBody) SetEncodeKey(v string) *PushDeviceNotificationRequestBody {
	s.EncodeKey = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetEncodeType(v string) *PushDeviceNotificationRequestBody {
	s.EncodeType = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetIsDebug(v bool) *PushDeviceNotificationRequestBody {
	s.IsDebug = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetMessageTemplateId(v string) *PushDeviceNotificationRequestBody {
	s.MessageTemplateId = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetOrganizationId(v string) *PushDeviceNotificationRequestBody {
	s.OrganizationId = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetPlaceHolder(v map[string]*string) *PushDeviceNotificationRequestBody {
	s.PlaceHolder = v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetSendTarget(v *PushDeviceNotificationRequestBodySendTarget) *PushDeviceNotificationRequestBody {
	s.SendTarget = v
	return s
}

type PushDeviceNotificationRequestBodySendTarget struct {
	// example:
	//
	// 2VpiDQ6aMjxz******Eo7r6e08oIVZ3fKrm5TyEfY=
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
	// example:
	//
	// DEVICE_OPEN_ID
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s PushDeviceNotificationRequestBodySendTarget) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceNotificationRequestBodySendTarget) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequestBodySendTarget) SetTargetIdentity(v string) *PushDeviceNotificationRequestBodySendTarget {
	s.TargetIdentity = &v
	return s
}

func (s *PushDeviceNotificationRequestBodySendTarget) SetTargetType(v string) *PushDeviceNotificationRequestBodySendTarget {
	s.TargetType = &v
	return s
}

type PushDeviceNotificationShrinkRequest struct {
	TenantInfoShrink *string `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty"`
	BodyShrink       *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDeviceNotificationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceNotificationShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationShrinkRequest) SetTenantInfoShrink(v string) *PushDeviceNotificationShrinkRequest {
	s.TenantInfoShrink = &v
	return s
}

func (s *PushDeviceNotificationShrinkRequest) SetBodyShrink(v string) *PushDeviceNotificationShrinkRequest {
	s.BodyShrink = &v
	return s
}

type PushDeviceNotificationResponseBody struct {
	// example:
	//
	// 908FA068-529C-0C20-8DB5-63B0EF7CFF1F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PushDeviceNotificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationResponseBody) SetRequestId(v string) *PushDeviceNotificationResponseBody {
	s.RequestId = &v
	return s
}

type PushDeviceNotificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushDeviceNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDeviceNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceNotificationResponse) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationResponse) SetHeaders(v map[string]*string) *PushDeviceNotificationResponse {
	s.Headers = v
	return s
}

func (s *PushDeviceNotificationResponse) SetStatusCode(v int32) *PushDeviceNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *PushDeviceNotificationResponse) SetBody(v *PushDeviceNotificationResponseBody) *PushDeviceNotificationResponse {
	s.Body = v
	return s
}

type QueryDeviceListResponseBody struct {
	DeviceList []*QueryDeviceListResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
	// example:
	//
	// 125****0946
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponseBody) SetDeviceList(v []*QueryDeviceListResponseBodyDeviceList) *QueryDeviceListResponseBody {
	s.DeviceList = v
	return s
}

func (s *QueryDeviceListResponseBody) SetEncodeKey(v string) *QueryDeviceListResponseBody {
	s.EncodeKey = &v
	return s
}

func (s *QueryDeviceListResponseBody) SetEncodeType(v string) *QueryDeviceListResponseBody {
	s.EncodeType = &v
	return s
}

func (s *QueryDeviceListResponseBody) SetRequestId(v string) *QueryDeviceListResponseBody {
	s.RequestId = &v
	return s
}

type QueryDeviceListResponseBodyDeviceList struct {
	// example:
	//
	// https://XXXXXX
	DeviceIconUrl *string `json:"DeviceIconUrl,omitempty" xml:"DeviceIconUrl,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// jMR2********ojVJXk=
	DeviceOpenId   *string                                                `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds []*QueryDeviceListResponseBodyDeviceListDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Online *string `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s QueryDeviceListResponseBodyDeviceList) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListResponseBodyDeviceList) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceIconUrl(v string) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceIconUrl = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceName(v string) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceOpenId(v string) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceOpenId = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceUnionIds(v []*QueryDeviceListResponseBodyDeviceListDeviceUnionIds) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceUnionIds = v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetOnline(v string) *QueryDeviceListResponseBodyDeviceList {
	s.Online = &v
	return s
}

type QueryDeviceListResponseBodyDeviceListDeviceUnionIds struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	UnionId        *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s QueryDeviceListResponseBodyDeviceListDeviceUnionIds) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListResponseBodyDeviceListDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponseBodyDeviceListDeviceUnionIds) SetOrganizationId(v string) *QueryDeviceListResponseBodyDeviceListDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceListDeviceUnionIds) SetUnionId(v string) *QueryDeviceListResponseBodyDeviceListDeviceUnionIds {
	s.UnionId = &v
	return s
}

type QueryDeviceListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponse) SetHeaders(v map[string]*string) *QueryDeviceListResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceListResponse) SetStatusCode(v int32) *QueryDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceListResponse) SetBody(v *QueryDeviceListResponseBody) *QueryDeviceListResponse {
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

// Summary:
//
// 创建播放列表
//
// @param tmpReq - CreatePlayingListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePlayingListResponse
func (client *Client) CreatePlayingListWithOptions(tmpReq *CreatePlayingListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePlayingListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OpenCreatePlayingListRequest)) {
		request.OpenCreatePlayingListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenCreatePlayingListRequest, tea.String("OpenCreatePlayingListRequest"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenCreatePlayingListRequestShrink)) {
		body["OpenCreatePlayingListRequest"] = request.OpenCreatePlayingListRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePlayingList"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/content/playing/create"),
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

// Summary:
//
// 创建播放列表
//
// @param request - CreatePlayingListRequest
//
// @return CreatePlayingListResponse
func (client *Client) CreatePlayingList(request *CreatePlayingListRequest) (_result *CreatePlayingListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePlayingListResponse{}
	_body, _err := client.CreatePlayingListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行场景
//
// @param request - ExecuteSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSceneResponse
func (client *Client) ExecuteSceneWithOptions(request *ExecuteSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteScene"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/iot/scene/execute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行场景
//
// @param request - ExecuteSceneRequest
//
// @return ExecuteSceneResponse
func (client *Client) ExecuteScene(request *ExecuteSceneRequest) (_result *ExecuteSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteSceneResponse{}
	_body, _err := client.ExecuteSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取场景列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneListResponse
func (client *Client) GetSceneListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSceneListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSceneList"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/iot/scene/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取场景列表
//
// @return GetSceneListResponse
func (client *Client) GetSceneList() (_result *GetSceneListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSceneListResponse{}
	_body, _err := client.GetSceneListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBasicInfoResponse
func (client *Client) GetUserBasicInfoWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserBasicInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserBasicInfo"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/users/basic"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取
//
// @return GetUserBasicInfoResponse
func (client *Client) GetUserBasicInfo() (_result *GetUserBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserBasicInfoResponse{}
	_body, _err := client.GetUserBasicInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取天猫精灵用户绑定的手机号
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPhoneResponse
func (client *Client) GetUserPhoneWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserPhoneResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserPhone"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/user/profile/phone"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取天猫精灵用户绑定的手机号
//
// @return GetUserPhoneResponse
func (client *Client) GetUserPhone() (_result *GetUserPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserPhoneResponse{}
	_body, _err := client.GetUserPhoneWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OAuth2令牌撤销端点
//
// @param request - OAuth2RevocationEndpointRequest
//
// @param headers - OAuth2RevocationEndpointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OAuth2RevocationEndpointResponse
func (client *Client) OAuth2RevocationEndpointWithOptions(request *OAuth2RevocationEndpointRequest, headers *OAuth2RevocationEndpointHeaders, runtime *util.RuntimeOptions) (_result *OAuth2RevocationEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.TokenTypeHint)) {
		body["TokenTypeHint"] = request.TokenTypeHint
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
		Action:      tea.String("OAuth2RevocationEndpoint"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OAuth2RevocationEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OAuth2令牌撤销端点
//
// @param request - OAuth2RevocationEndpointRequest
//
// @return OAuth2RevocationEndpointResponse
func (client *Client) OAuth2RevocationEndpoint(request *OAuth2RevocationEndpointRequest) (_result *OAuth2RevocationEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OAuth2RevocationEndpointHeaders{}
	_result = &OAuth2RevocationEndpointResponse{}
	_body, _err := client.OAuth2RevocationEndpointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OAuth2令牌端点
//
// @param request - OAuth2TokenEndpointRequest
//
// @param headers - OAuth2TokenEndpointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OAuth2TokenEndpointResponse
func (client *Client) OAuth2TokenEndpointWithOptions(request *OAuth2TokenEndpointRequest, headers *OAuth2TokenEndpointHeaders, runtime *util.RuntimeOptions) (_result *OAuth2TokenEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		body["GrantType"] = request.GrantType
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUri)) {
		body["RedirectUri"] = request.RedirectUri
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		body["RefreshToken"] = request.RefreshToken
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
		Action:      tea.String("OAuth2TokenEndpoint"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/token"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OAuth2TokenEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OAuth2令牌端点
//
// @param request - OAuth2TokenEndpointRequest
//
// @return OAuth2TokenEndpointResponse
func (client *Client) OAuth2TokenEndpoint(request *OAuth2TokenEndpointRequest) (_result *OAuth2TokenEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OAuth2TokenEndpointHeaders{}
	_result = &OAuth2TokenEndpointResponse{}
	_body, _err := client.OAuth2TokenEndpointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送设备通知
//
// @param tmpReq - PushDeviceNotificationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushDeviceNotificationResponse
func (client *Client) PushDeviceNotificationWithOptions(tmpReq *PushDeviceNotificationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushDeviceNotificationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PushDeviceNotificationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TenantInfo)) {
		request.TenantInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantInfo, tea.String("TenantInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantInfoShrink)) {
		body["TenantInfo"] = request.TenantInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		body["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushDeviceNotification"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/device/notification/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushDeviceNotificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送设备通知
//
// @param request - PushDeviceNotificationRequest
//
// @return PushDeviceNotificationResponse
func (client *Client) PushDeviceNotification(request *PushDeviceNotificationRequest) (_result *PushDeviceNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushDeviceNotificationResponse{}
	_body, _err := client.PushDeviceNotificationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceListResponse
func (client *Client) QueryDeviceListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryDeviceListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceList"),
		Version:     tea.String("oauth2_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/oauth2/device/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备列表
//
// @return QueryDeviceListResponse
func (client *Client) QueryDeviceList() (_result *QueryDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryDeviceListResponse{}
	_body, _err := client.QueryDeviceListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
