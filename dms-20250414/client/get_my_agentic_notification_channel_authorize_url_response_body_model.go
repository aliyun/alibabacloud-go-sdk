// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyAgenticNotificationChannelAuthorizeUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody
	GetData() *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData
	SetErrorCode(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody
	GetSuccess() *bool
}

type GetMyAgenticNotificationChannelAuthorizeUrlResponseBody struct {
	// The authorization URL information.
	Data *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// channel must be FEISHU
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) GetData() *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData {
	return s.Data
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) SetData(v *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) SetErrorCode(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) SetErrorMessage(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) SetRequestId(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) SetSuccess(v bool) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData struct {
	// The URL of the Lark OAuth authorization page. After the user opens the URL in a browser and completes the authorization, the system binds the Lark account as the notification identity for the current user, and the channel switch remains unchanged. The state parameter in the URL is valid for 180 seconds and can be used only once.
	//
	// example:
	//
	// https://open.feishu.cn/open-apis/authen/v1/authorize?app_id=cli_a5f8e2d9c3b1****&redirect_uri=https%3A%2F%2Faig-dms-cn-hangzhou.aliyuncs.com%2Foauth%2Fmobile%2Ffeishu%2Fbind&state=eyJhbGciOiJIUzI1NiJ9****
	AuthorizeUrl *string `json:"AuthorizeUrl,omitempty" xml:"AuthorizeUrl,omitempty"`
	// The callback URL after the authorization is completed. You must add this URL to the redirect URL whitelist of the tenant\\"s Lark application. Otherwise, the authorization fails.
	//
	// example:
	//
	// https://aig-dms-cn-hangzhou.aliyuncs.com/oauth/mobile/feishu/bind
	RedirectUri *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) GetAuthorizeUrl() *string {
	return s.AuthorizeUrl
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) SetAuthorizeUrl(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData {
	s.AuthorizeUrl = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) SetRedirectUri(v string) *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData {
	s.RedirectUri = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
