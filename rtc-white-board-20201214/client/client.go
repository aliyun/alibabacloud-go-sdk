// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckWhiteBoardHostRequest struct {
	DocKey     *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
}

func (s CheckWhiteBoardHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckWhiteBoardHostRequest) GoString() string {
	return s.String()
}

func (s *CheckWhiteBoardHostRequest) SetDocKey(v string) *CheckWhiteBoardHostRequest {
	s.DocKey = &v
	return s
}

func (s *CheckWhiteBoardHostRequest) SetOriginHost(v string) *CheckWhiteBoardHostRequest {
	s.OriginHost = &v
	return s
}

type CheckWhiteBoardHostResponseBody struct {
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool   `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckWhiteBoardHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckWhiteBoardHostResponseBody) GoString() string {
	return s.String()
}

func (s *CheckWhiteBoardHostResponseBody) SetRequestId(v string) *CheckWhiteBoardHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckWhiteBoardHostResponseBody) SetResponseSuccess(v bool) *CheckWhiteBoardHostResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CheckWhiteBoardHostResponseBody) SetErrorCode(v string) *CheckWhiteBoardHostResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckWhiteBoardHostResponseBody) SetErrorMsg(v string) *CheckWhiteBoardHostResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CheckWhiteBoardHostResponseBody) SetResult(v bool) *CheckWhiteBoardHostResponseBody {
	s.Result = &v
	return s
}

type CheckWhiteBoardHostResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckWhiteBoardHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckWhiteBoardHostResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckWhiteBoardHostResponse) GoString() string {
	return s.String()
}

func (s *CheckWhiteBoardHostResponse) SetHeaders(v map[string]*string) *CheckWhiteBoardHostResponse {
	s.Headers = v
	return s
}

func (s *CheckWhiteBoardHostResponse) SetBody(v *CheckWhiteBoardHostResponseBody) *CheckWhiteBoardHostResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

type CreateAppResponseBody struct {
	// Id of the request
	RequestId       *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool                        `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *CreateAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResponseSuccess(v bool) *CreateAppResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CreateAppResponseBody) SetErrorCode(v string) *CreateAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAppResponseBody) SetErrorMsg(v string) *CreateAppResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponseBodyResult struct {
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) SetAppID(v string) *CreateAppResponseBodyResult {
	s.AppID = &v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateWhiteBoardRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AppID  *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
}

func (s CreateWhiteBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardRequest) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardRequest) SetUserId(v string) *CreateWhiteBoardRequest {
	s.UserId = &v
	return s
}

func (s *CreateWhiteBoardRequest) SetAppID(v string) *CreateWhiteBoardRequest {
	s.AppID = &v
	return s
}

type CreateWhiteBoardResponseBody struct {
	// Id of the request
	RequestId       *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool                               `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *CreateWhiteBoardResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateWhiteBoardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardResponseBody) SetRequestId(v string) *CreateWhiteBoardResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetResponseSuccess(v bool) *CreateWhiteBoardResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetErrorCode(v string) *CreateWhiteBoardResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetErrorMsg(v string) *CreateWhiteBoardResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetResult(v *CreateWhiteBoardResponseBodyResult) *CreateWhiteBoardResponseBody {
	s.Result = v
	return s
}

type CreateWhiteBoardResponseBodyResult struct {
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
}

func (s CreateWhiteBoardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardResponseBodyResult) SetDocKey(v string) *CreateWhiteBoardResponseBodyResult {
	s.DocKey = &v
	return s
}

type CreateWhiteBoardResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWhiteBoardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWhiteBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardResponse) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardResponse) SetHeaders(v map[string]*string) *CreateWhiteBoardResponse {
	s.Headers = v
	return s
}

func (s *CreateWhiteBoardResponse) SetBody(v *CreateWhiteBoardResponseBody) *CreateWhiteBoardResponse {
	s.Body = v
	return s
}

type GetUserPermissionCallbackRequest struct {
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DocKey         *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
}

func (s GetUserPermissionCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserPermissionCallbackRequest) GoString() string {
	return s.String()
}

func (s *GetUserPermissionCallbackRequest) SetUserId(v string) *GetUserPermissionCallbackRequest {
	s.UserId = &v
	return s
}

func (s *GetUserPermissionCallbackRequest) SetDocKey(v string) *GetUserPermissionCallbackRequest {
	s.DocKey = &v
	return s
}

func (s *GetUserPermissionCallbackRequest) SetPermissionType(v string) *GetUserPermissionCallbackRequest {
	s.PermissionType = &v
	return s
}

type GetUserPermissionCallbackResponseBody struct {
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool   `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetUserPermissionCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserPermissionCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPermissionCallbackResponseBody) SetRequestId(v string) *GetUserPermissionCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserPermissionCallbackResponseBody) SetResponseSuccess(v bool) *GetUserPermissionCallbackResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *GetUserPermissionCallbackResponseBody) SetErrorCode(v string) *GetUserPermissionCallbackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserPermissionCallbackResponseBody) SetErrorMsg(v string) *GetUserPermissionCallbackResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserPermissionCallbackResponseBody) SetResult(v bool) *GetUserPermissionCallbackResponseBody {
	s.Result = &v
	return s
}

type GetUserPermissionCallbackResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserPermissionCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserPermissionCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserPermissionCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetUserPermissionCallbackResponse) SetHeaders(v map[string]*string) *GetUserPermissionCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetUserPermissionCallbackResponse) SetBody(v *GetUserPermissionCallbackResponseBody) *GetUserPermissionCallbackResponse {
	s.Body = v
	return s
}

type GetUserProfileCallbackRequest struct {
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	DocKey  *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
}

func (s GetUserProfileCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserProfileCallbackRequest) GoString() string {
	return s.String()
}

func (s *GetUserProfileCallbackRequest) SetUserIds(v string) *GetUserProfileCallbackRequest {
	s.UserIds = &v
	return s
}

func (s *GetUserProfileCallbackRequest) SetDocKey(v string) *GetUserProfileCallbackRequest {
	s.DocKey = &v
	return s
}

type GetUserProfileCallbackResponseBody struct {
	// Id of the request
	RequestId       *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool                                     `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *GetUserProfileCallbackResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetUserProfileCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserProfileCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProfileCallbackResponseBody) SetRequestId(v string) *GetUserProfileCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProfileCallbackResponseBody) SetResponseSuccess(v bool) *GetUserProfileCallbackResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *GetUserProfileCallbackResponseBody) SetErrorCode(v string) *GetUserProfileCallbackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserProfileCallbackResponseBody) SetErrorMsg(v string) *GetUserProfileCallbackResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserProfileCallbackResponseBody) SetResult(v *GetUserProfileCallbackResponseBodyResult) *GetUserProfileCallbackResponseBody {
	s.Result = v
	return s
}

type GetUserProfileCallbackResponseBodyResult struct {
	UserProfileList []*GetUserProfileCallbackResponseBodyResultUserProfileList `json:"UserProfileList,omitempty" xml:"UserProfileList,omitempty" type:"Repeated"`
}

func (s GetUserProfileCallbackResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserProfileCallbackResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserProfileCallbackResponseBodyResult) SetUserProfileList(v []*GetUserProfileCallbackResponseBodyResultUserProfileList) *GetUserProfileCallbackResponseBodyResult {
	s.UserProfileList = v
	return s
}

type GetUserProfileCallbackResponseBodyResultUserProfileList struct {
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AvatarUrl  *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Nick       *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	NickPinyin *string `json:"NickPinyin,omitempty" xml:"NickPinyin,omitempty"`
}

func (s GetUserProfileCallbackResponseBodyResultUserProfileList) String() string {
	return tea.Prettify(s)
}

func (s GetUserProfileCallbackResponseBodyResultUserProfileList) GoString() string {
	return s.String()
}

func (s *GetUserProfileCallbackResponseBodyResultUserProfileList) SetUserId(v string) *GetUserProfileCallbackResponseBodyResultUserProfileList {
	s.UserId = &v
	return s
}

func (s *GetUserProfileCallbackResponseBodyResultUserProfileList) SetAvatarUrl(v string) *GetUserProfileCallbackResponseBodyResultUserProfileList {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserProfileCallbackResponseBodyResultUserProfileList) SetNick(v string) *GetUserProfileCallbackResponseBodyResultUserProfileList {
	s.Nick = &v
	return s
}

func (s *GetUserProfileCallbackResponseBodyResultUserProfileList) SetNickPinyin(v string) *GetUserProfileCallbackResponseBodyResultUserProfileList {
	s.NickPinyin = &v
	return s
}

type GetUserProfileCallbackResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserProfileCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserProfileCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserProfileCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetUserProfileCallbackResponse) SetHeaders(v map[string]*string) *GetUserProfileCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetUserProfileCallbackResponse) SetBody(v *GetUserProfileCallbackResponseBody) *GetUserProfileCallbackResponse {
	s.Body = v
	return s
}

type GetWhiteBoardProfileCallbackRequest struct {
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
}

func (s GetWhiteBoardProfileCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWhiteBoardProfileCallbackRequest) GoString() string {
	return s.String()
}

func (s *GetWhiteBoardProfileCallbackRequest) SetDocKey(v string) *GetWhiteBoardProfileCallbackRequest {
	s.DocKey = &v
	return s
}

type GetWhiteBoardProfileCallbackResponseBody struct {
	// Id of the request
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool                                           `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string                                         `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *GetWhiteBoardProfileCallbackResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetWhiteBoardProfileCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWhiteBoardProfileCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetWhiteBoardProfileCallbackResponseBody) SetRequestId(v string) *GetWhiteBoardProfileCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWhiteBoardProfileCallbackResponseBody) SetResponseSuccess(v bool) *GetWhiteBoardProfileCallbackResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *GetWhiteBoardProfileCallbackResponseBody) SetErrorCode(v string) *GetWhiteBoardProfileCallbackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWhiteBoardProfileCallbackResponseBody) SetErrorMsg(v string) *GetWhiteBoardProfileCallbackResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWhiteBoardProfileCallbackResponseBody) SetResult(v *GetWhiteBoardProfileCallbackResponseBodyResult) *GetWhiteBoardProfileCallbackResponseBody {
	s.Result = v
	return s
}

type GetWhiteBoardProfileCallbackResponseBodyResult struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetWhiteBoardProfileCallbackResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetWhiteBoardProfileCallbackResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWhiteBoardProfileCallbackResponseBodyResult) SetName(v string) *GetWhiteBoardProfileCallbackResponseBodyResult {
	s.Name = &v
	return s
}

type GetWhiteBoardProfileCallbackResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWhiteBoardProfileCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWhiteBoardProfileCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWhiteBoardProfileCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetWhiteBoardProfileCallbackResponse) SetHeaders(v map[string]*string) *GetWhiteBoardProfileCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetWhiteBoardProfileCallbackResponse) SetBody(v *GetWhiteBoardProfileCallbackResponseBody) *GetWhiteBoardProfileCallbackResponse {
	s.Body = v
	return s
}

type OpenWhiteBoardRequest struct {
	AppID  *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
}

func (s OpenWhiteBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardRequest) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardRequest) SetAppID(v string) *OpenWhiteBoardRequest {
	s.AppID = &v
	return s
}

func (s *OpenWhiteBoardRequest) SetUserId(v string) *OpenWhiteBoardRequest {
	s.UserId = &v
	return s
}

func (s *OpenWhiteBoardRequest) SetDocKey(v string) *OpenWhiteBoardRequest {
	s.DocKey = &v
	return s
}

type OpenWhiteBoardResponseBody struct {
	// Id of the request
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool                             `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string                           `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *OpenWhiteBoardResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s OpenWhiteBoardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBody) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBody) SetRequestId(v string) *OpenWhiteBoardResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetResponseSuccess(v bool) *OpenWhiteBoardResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetErrorCode(v string) *OpenWhiteBoardResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetErrorMsg(v string) *OpenWhiteBoardResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetResult(v *OpenWhiteBoardResponseBodyResult) *OpenWhiteBoardResponseBody {
	s.Result = v
	return s
}

type OpenWhiteBoardResponseBodyResult struct {
	DocumentAccessInfo *OpenWhiteBoardResponseBodyResultDocumentAccessInfo `json:"DocumentAccessInfo,omitempty" xml:"DocumentAccessInfo,omitempty" type:"Struct"`
}

func (s OpenWhiteBoardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBodyResult) SetDocumentAccessInfo(v *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) *OpenWhiteBoardResponseBodyResult {
	s.DocumentAccessInfo = v
	return s
}

type OpenWhiteBoardResponseBodyResultDocumentAccessInfo struct {
	AccessToken *string                                                     `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	CollabHost  *string                                                     `json:"CollabHost,omitempty" xml:"CollabHost,omitempty"`
	Permission  *int64                                                      `json:"Permission,omitempty" xml:"Permission,omitempty"`
	UserInfo    *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfo) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfo) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetAccessToken(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.AccessToken = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetCollabHost(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.CollabHost = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetPermission(v int64) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.Permission = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetUserInfo(v *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.UserInfo = v
	return s
}

type OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo struct {
	AvatarUrl  *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Nick       *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	NickPinyin *string `json:"NickPinyin,omitempty" xml:"NickPinyin,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetAvatarUrl(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.AvatarUrl = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetNick(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.Nick = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetNickPinyin(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.NickPinyin = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetUserId(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.UserId = &v
	return s
}

type OpenWhiteBoardResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenWhiteBoardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenWhiteBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponse) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponse) SetHeaders(v map[string]*string) *OpenWhiteBoardResponse {
	s.Headers = v
	return s
}

func (s *OpenWhiteBoardResponse) SetBody(v *OpenWhiteBoardResponseBody) *OpenWhiteBoardResponse {
	s.Body = v
	return s
}

type RefreshUsersPermissionsRequest struct {
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	DocKey  *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	AppID   *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
}

func (s RefreshUsersPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshUsersPermissionsRequest) GoString() string {
	return s.String()
}

func (s *RefreshUsersPermissionsRequest) SetUserIds(v string) *RefreshUsersPermissionsRequest {
	s.UserIds = &v
	return s
}

func (s *RefreshUsersPermissionsRequest) SetDocKey(v string) *RefreshUsersPermissionsRequest {
	s.DocKey = &v
	return s
}

func (s *RefreshUsersPermissionsRequest) SetAppID(v string) *RefreshUsersPermissionsRequest {
	s.AppID = &v
	return s
}

type RefreshUsersPermissionsResponseBody struct {
	// Id of the request
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool                                      `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string                                    `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *RefreshUsersPermissionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s RefreshUsersPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshUsersPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshUsersPermissionsResponseBody) SetRequestId(v string) *RefreshUsersPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetResponseSuccess(v bool) *RefreshUsersPermissionsResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetErrorCode(v string) *RefreshUsersPermissionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetErrorMsg(v string) *RefreshUsersPermissionsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetResult(v *RefreshUsersPermissionsResponseBodyResult) *RefreshUsersPermissionsResponseBody {
	s.Result = v
	return s
}

type RefreshUsersPermissionsResponseBodyResult struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RefreshUsersPermissionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RefreshUsersPermissionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RefreshUsersPermissionsResponseBodyResult) SetCode(v string) *RefreshUsersPermissionsResponseBodyResult {
	s.Code = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBodyResult) SetMessage(v string) *RefreshUsersPermissionsResponseBodyResult {
	s.Message = &v
	return s
}

type RefreshUsersPermissionsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshUsersPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshUsersPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshUsersPermissionsResponse) GoString() string {
	return s.String()
}

func (s *RefreshUsersPermissionsResponse) SetHeaders(v map[string]*string) *RefreshUsersPermissionsResponse {
	s.Headers = v
	return s
}

func (s *RefreshUsersPermissionsResponse) SetBody(v *RefreshUsersPermissionsResponseBody) *RefreshUsersPermissionsResponse {
	s.Body = v
	return s
}

type SetAppCallbackUrlRequest struct {
	AppID              *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	AppCallbackUrl     *string `json:"AppCallbackUrl,omitempty" xml:"AppCallbackUrl,omitempty"`
	AppCallbackAuthKey *string `json:"AppCallbackAuthKey,omitempty" xml:"AppCallbackAuthKey,omitempty"`
}

func (s SetAppCallbackUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppCallbackUrlRequest) GoString() string {
	return s.String()
}

func (s *SetAppCallbackUrlRequest) SetAppID(v string) *SetAppCallbackUrlRequest {
	s.AppID = &v
	return s
}

func (s *SetAppCallbackUrlRequest) SetAppCallbackUrl(v string) *SetAppCallbackUrlRequest {
	s.AppCallbackUrl = &v
	return s
}

func (s *SetAppCallbackUrlRequest) SetAppCallbackAuthKey(v string) *SetAppCallbackUrlRequest {
	s.AppCallbackAuthKey = &v
	return s
}

type SetAppCallbackUrlResponseBody struct {
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool   `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppCallbackUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppCallbackUrlResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppCallbackUrlResponseBody) SetRequestId(v string) *SetAppCallbackUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetResponseSuccess(v bool) *SetAppCallbackUrlResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetErrorCode(v string) *SetAppCallbackUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetErrorMsg(v string) *SetAppCallbackUrlResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetResult(v bool) *SetAppCallbackUrlResponseBody {
	s.Result = &v
	return s
}

type SetAppCallbackUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppCallbackUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppCallbackUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppCallbackUrlResponse) GoString() string {
	return s.String()
}

func (s *SetAppCallbackUrlResponse) SetHeaders(v map[string]*string) *SetAppCallbackUrlResponse {
	s.Headers = v
	return s
}

func (s *SetAppCallbackUrlResponse) SetBody(v *SetAppCallbackUrlResponseBody) *SetAppCallbackUrlResponse {
	s.Body = v
	return s
}

type SetAppDomainNamesRequest struct {
	AppID          *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	AppDomainNames *string `json:"AppDomainNames,omitempty" xml:"AppDomainNames,omitempty"`
}

func (s SetAppDomainNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppDomainNamesRequest) GoString() string {
	return s.String()
}

func (s *SetAppDomainNamesRequest) SetAppID(v string) *SetAppDomainNamesRequest {
	s.AppID = &v
	return s
}

func (s *SetAppDomainNamesRequest) SetAppDomainNames(v string) *SetAppDomainNamesRequest {
	s.AppDomainNames = &v
	return s
}

type SetAppDomainNamesResponseBody struct {
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool   `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppDomainNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppDomainNamesResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppDomainNamesResponseBody) SetRequestId(v string) *SetAppDomainNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetResponseSuccess(v bool) *SetAppDomainNamesResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetErrorCode(v string) *SetAppDomainNamesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetErrorMsg(v string) *SetAppDomainNamesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetResult(v bool) *SetAppDomainNamesResponseBody {
	s.Result = &v
	return s
}

type SetAppDomainNamesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppDomainNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppDomainNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppDomainNamesResponse) GoString() string {
	return s.String()
}

func (s *SetAppDomainNamesResponse) SetHeaders(v map[string]*string) *SetAppDomainNamesResponse {
	s.Headers = v
	return s
}

func (s *SetAppDomainNamesResponse) SetBody(v *SetAppDomainNamesResponseBody) *SetAppDomainNamesResponse {
	s.Body = v
	return s
}

type SetAppNameRequest struct {
	AppID   *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s SetAppNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppNameRequest) GoString() string {
	return s.String()
}

func (s *SetAppNameRequest) SetAppID(v string) *SetAppNameRequest {
	s.AppID = &v
	return s
}

func (s *SetAppNameRequest) SetAppName(v string) *SetAppNameRequest {
	s.AppName = &v
	return s
}

type SetAppNameResponseBody struct {
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool   `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppNameResponseBody) SetRequestId(v string) *SetAppNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppNameResponseBody) SetResponseSuccess(v bool) *SetAppNameResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppNameResponseBody) SetErrorCode(v string) *SetAppNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppNameResponseBody) SetErrorMsg(v string) *SetAppNameResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppNameResponseBody) SetResult(v bool) *SetAppNameResponseBody {
	s.Result = &v
	return s
}

type SetAppNameResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppNameResponse) GoString() string {
	return s.String()
}

func (s *SetAppNameResponse) SetHeaders(v map[string]*string) *SetAppNameResponse {
	s.Headers = v
	return s
}

func (s *SetAppNameResponse) SetBody(v *SetAppNameResponseBody) *SetAppNameResponse {
	s.Body = v
	return s
}

type SetAppStatusRequest struct {
	AppID     *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	AppStatus *int64  `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
}

func (s SetAppStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppStatusRequest) GoString() string {
	return s.String()
}

func (s *SetAppStatusRequest) SetAppID(v string) *SetAppStatusRequest {
	s.AppID = &v
	return s
}

func (s *SetAppStatusRequest) SetAppStatus(v int64) *SetAppStatusRequest {
	s.AppStatus = &v
	return s
}

type SetAppStatusResponseBody struct {
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool   `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg        *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Result          *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppStatusResponseBody) SetRequestId(v string) *SetAppStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppStatusResponseBody) SetResponseSuccess(v bool) *SetAppStatusResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppStatusResponseBody) SetErrorCode(v string) *SetAppStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppStatusResponseBody) SetErrorMsg(v string) *SetAppStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppStatusResponseBody) SetResult(v bool) *SetAppStatusResponseBody {
	s.Result = &v
	return s
}

type SetAppStatusResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppStatusResponse) GoString() string {
	return s.String()
}

func (s *SetAppStatusResponse) SetHeaders(v map[string]*string) *SetAppStatusResponse {
	s.Headers = v
	return s
}

func (s *SetAppStatusResponse) SetBody(v *SetAppStatusResponseBody) *SetAppStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("rtc-white-board"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CheckWhiteBoardHostWithOptions(request *CheckWhiteBoardHostRequest, runtime *util.RuntimeOptions) (_result *CheckWhiteBoardHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckWhiteBoardHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckWhiteBoardHost"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckWhiteBoardHost(request *CheckWhiteBoardHostRequest) (_result *CheckWhiteBoardHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckWhiteBoardHostResponse{}
	_body, _err := client.CheckWhiteBoardHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApp"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWhiteBoardWithOptions(request *CreateWhiteBoardRequest, runtime *util.RuntimeOptions) (_result *CreateWhiteBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWhiteBoardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWhiteBoard"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWhiteBoard(request *CreateWhiteBoardRequest) (_result *CreateWhiteBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWhiteBoardResponse{}
	_body, _err := client.CreateWhiteBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserPermissionCallbackWithOptions(request *GetUserPermissionCallbackRequest, runtime *util.RuntimeOptions) (_result *GetUserPermissionCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserPermissionCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserPermissionCallback"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserPermissionCallback(request *GetUserPermissionCallbackRequest) (_result *GetUserPermissionCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserPermissionCallbackResponse{}
	_body, _err := client.GetUserPermissionCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserProfileCallbackWithOptions(request *GetUserProfileCallbackRequest, runtime *util.RuntimeOptions) (_result *GetUserProfileCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserProfileCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserProfileCallback"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserProfileCallback(request *GetUserProfileCallbackRequest) (_result *GetUserProfileCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserProfileCallbackResponse{}
	_body, _err := client.GetUserProfileCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWhiteBoardProfileCallbackWithOptions(request *GetWhiteBoardProfileCallbackRequest, runtime *util.RuntimeOptions) (_result *GetWhiteBoardProfileCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWhiteBoardProfileCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWhiteBoardProfileCallback"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWhiteBoardProfileCallback(request *GetWhiteBoardProfileCallbackRequest) (_result *GetWhiteBoardProfileCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWhiteBoardProfileCallbackResponse{}
	_body, _err := client.GetWhiteBoardProfileCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenWhiteBoardWithOptions(request *OpenWhiteBoardRequest, runtime *util.RuntimeOptions) (_result *OpenWhiteBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenWhiteBoardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenWhiteBoard"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenWhiteBoard(request *OpenWhiteBoardRequest) (_result *OpenWhiteBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenWhiteBoardResponse{}
	_body, _err := client.OpenWhiteBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshUsersPermissionsWithOptions(request *RefreshUsersPermissionsRequest, runtime *util.RuntimeOptions) (_result *RefreshUsersPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshUsersPermissionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshUsersPermissions"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshUsersPermissions(request *RefreshUsersPermissionsRequest) (_result *RefreshUsersPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshUsersPermissionsResponse{}
	_body, _err := client.RefreshUsersPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppCallbackUrlWithOptions(request *SetAppCallbackUrlRequest, runtime *util.RuntimeOptions) (_result *SetAppCallbackUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppCallbackUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppCallbackUrl"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppCallbackUrl(request *SetAppCallbackUrlRequest) (_result *SetAppCallbackUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppCallbackUrlResponse{}
	_body, _err := client.SetAppCallbackUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppDomainNamesWithOptions(request *SetAppDomainNamesRequest, runtime *util.RuntimeOptions) (_result *SetAppDomainNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppDomainNamesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppDomainNames"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppDomainNames(request *SetAppDomainNamesRequest) (_result *SetAppDomainNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppDomainNamesResponse{}
	_body, _err := client.SetAppDomainNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppNameWithOptions(request *SetAppNameRequest, runtime *util.RuntimeOptions) (_result *SetAppNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppName"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppName(request *SetAppNameRequest) (_result *SetAppNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppNameResponse{}
	_body, _err := client.SetAppNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppStatusWithOptions(request *SetAppStatusRequest, runtime *util.RuntimeOptions) (_result *SetAppStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppStatus"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppStatus(request *SetAppStatusRequest) (_result *SetAppStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppStatusResponse{}
	_body, _err := client.SetAppStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
