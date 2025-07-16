// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateChatSessionResponseBody
	GetCode() *string
	SetData(v *CreateChatSessionResponseBodyData) *CreateChatSessionResponseBody
	GetData() *CreateChatSessionResponseBodyData
	SetHttpStatusCode(v int64) *CreateChatSessionResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateChatSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateChatSessionResponseBody
	GetSuccess() *bool
}

type CreateChatSessionResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateChatSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateChatSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatSessionResponseBody) GetData() *CreateChatSessionResponseBodyData {
	return s.Data
}

func (s *CreateChatSessionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateChatSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateChatSessionResponseBody) SetCode(v string) *CreateChatSessionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatSessionResponseBody) SetData(v *CreateChatSessionResponseBodyData) *CreateChatSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatSessionResponseBody) SetHttpStatusCode(v int64) *CreateChatSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateChatSessionResponseBody) SetMessage(v string) *CreateChatSessionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatSessionResponseBody) SetRequestId(v string) *CreateChatSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatSessionResponseBody) SetSuccess(v bool) *CreateChatSessionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateChatSessionResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateChatSessionResponseBodyData struct {
	AvatarAssets *CreateChatSessionResponseBodyDataAvatarAssets `json:"avatarAssets,omitempty" xml:"avatarAssets,omitempty" type:"Struct"`
	RtcParams    *CreateChatSessionResponseBodyDataRtcParams    `json:"rtcParams,omitempty" xml:"rtcParams,omitempty" type:"Struct"`
	// example:
	//
	// 9827f4bd-5008-4d34-98fb-62598f3ad3b5
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateChatSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBodyData) GetAvatarAssets() *CreateChatSessionResponseBodyDataAvatarAssets {
	return s.AvatarAssets
}

func (s *CreateChatSessionResponseBodyData) GetRtcParams() *CreateChatSessionResponseBodyDataRtcParams {
	return s.RtcParams
}

func (s *CreateChatSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateChatSessionResponseBodyData) SetAvatarAssets(v *CreateChatSessionResponseBodyDataAvatarAssets) *CreateChatSessionResponseBodyData {
	s.AvatarAssets = v
	return s
}

func (s *CreateChatSessionResponseBodyData) SetRtcParams(v *CreateChatSessionResponseBodyDataRtcParams) *CreateChatSessionResponseBodyData {
	s.RtcParams = v
	return s
}

func (s *CreateChatSessionResponseBodyData) SetSessionId(v string) *CreateChatSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateChatSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CreateChatSessionResponseBodyDataAvatarAssets struct {
	// example:
	//
	// 5B83BE2114489274BB88BADE7EBC9BC0
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// example:
	//
	// 0.0.1
	MinRequiredVersion *string `json:"minRequiredVersion,omitempty" xml:"minRequiredVersion,omitempty"`
	// example:
	//
	// J562PNqJBZDhzOQpLBgIcIW8+rHQoM7P6IONGMP7P5vGxrWLxT7VtRenFnMY+wg/zpA2qwpFBmJYO2rVexnlCQ2WE4kvYOH/OKmlTzpQddY34U5jS9KaS3b3ulpq4xnKDjWJ+sLZSRMhuPDdlq8ZPfcfEPhJhF3zPO8Hu4QOSu+D/pAIDJUoixOTo9Q14DXFKGFuuVRQOQ7f/VxJcoSLIWIusV917pLtph/IYBaLd27gzbrTZBEVD8qrucR+WOQPY1g67PGAdafkhJWrs/+coM7+5dc3HEUC+KgI9JN4X4Akelc94aJcy78RZ6tRdr73hBzN83/cMZdzt2hxnjzg==
	Secret *string `json:"secret,omitempty" xml:"secret,omitempty"`
	// example:
	//
	// AVATAR_3D_TRADITIONAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// https://daily-avatar-property.oss-cn-beijing.aliyuncs.com/avatar-share-property/AVATAR_3D_TRADITIONAL/Mt.CNMU6BO4RBYU2/secret_assets_web.zip?Expires=1752637519&OSSAccessKeyId=STS.NZULzwLRx8thHDHQxem94dWdN&Signature=Oni3%2Be8dY8Xrv3iRGDyzn7uReVc%3D&security-token=CAISzAJ1q6Ft5B2yfSjIr5ngB8DDoY1Zj7aDSmL5tXgwYbYYi5LPrDz2IHhMfnloB%2BEcsfU3nmxT6vkZlrp6SJtIXleCZtF94oxN9h2gb4fb4093DEHt08%2FLI3OaLjKm9u2wCryLYbGwU%2FOpbE%2B%2B5U0X6LDmdDKkckW4OJmS8%2FBOZcgWWQ%2FKBlgvRq0hRG1YpdQdKGHaONu0LxfumRCwNkdzvRdmgm4NgsbWgO%2Fks0KA1QSml7ZP%2B9WuesH0M%2FMBZskvD42Hu8VtbbfE3SJq7BxHybx7lqQs%2B02c5onHUwEPsk%2FZYrKOroYzc1RjAbM%2FErRY6fP8nOE9ovbUm5RXHpT05CrMOs62ZPdDoKOscIvBXr6yZaP7JmcGC6iQLG%2FznQkSc081IsK2C7Xq0pe54O3lg9Ab41ZGNYEjq%2BpCIUP%2Fs97dqXEelD2e%2Bh8UezDnKxqAAXuAiYRY7Ox3cf6h2MlmRsK5yywg45O%2FizjiK2k8Z8p6WeOA54W3pfbg6ElV4d8TMWCVZ7tuAbSgRCKBg3q5YYrdS2ENqDu6njIea1pxG4LT4ydGxDBkYpjwcUxutDd0aAhFjsypSK%2Feuk0%2FDCfKMrWzCmkr1AtPpcNfJ8LPj58qIA
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateChatSessionResponseBodyDataAvatarAssets) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponseBodyDataAvatarAssets) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) GetMd5() *string {
	return s.Md5
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) GetMinRequiredVersion() *string {
	return s.MinRequiredVersion
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) GetSecret() *string {
	return s.Secret
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) GetType() *string {
	return s.Type
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) GetUrl() *string {
	return s.Url
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) SetMd5(v string) *CreateChatSessionResponseBodyDataAvatarAssets {
	s.Md5 = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) SetMinRequiredVersion(v string) *CreateChatSessionResponseBodyDataAvatarAssets {
	s.MinRequiredVersion = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) SetSecret(v string) *CreateChatSessionResponseBodyDataAvatarAssets {
	s.Secret = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) SetType(v string) *CreateChatSessionResponseBodyDataAvatarAssets {
	s.Type = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) SetUrl(v string) *CreateChatSessionResponseBodyDataAvatarAssets {
	s.Url = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataAvatarAssets) Validate() error {
	return dara.Validate(s)
}

type CreateChatSessionResponseBodyDataRtcParams struct {
	// example:
	//
	// 895cbf3b
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// E7enIvjUos
	AvatarUserId *string `json:"avatarUserId,omitempty" xml:"avatarUserId,omitempty"`
	// example:
	//
	// pPltqR3FovNCK3hNQc8eHUL3Ztq1wY
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// aw0tqpFlP4
	ClientUserId *string `json:"clientUserId,omitempty" xml:"clientUserId,omitempty"`
	// example:
	//
	// https://gw.rtn.aliyuncs.com
	Gslb *string `json:"gslb,omitempty" xml:"gslb,omitempty"`
	// example:
	//
	// f8b0ef02c5da778f4488e2470c
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// example:
	//
	// YzZtSQP8QX
	ServerUserId *string `json:"serverUserId,omitempty" xml:"serverUserId,omitempty"`
	// example:
	//
	// 1560588594
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// PtGgv2dM9F8tEuAtda50c0VNNFjn0WUbyTDRPa1im4cUBEcxlo
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s CreateChatSessionResponseBodyDataRtcParams) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponseBodyDataRtcParams) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetAppId() *string {
	return s.AppId
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetAvatarUserId() *string {
	return s.AvatarUserId
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetChannel() *string {
	return s.Channel
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetClientUserId() *string {
	return s.ClientUserId
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetGslb() *string {
	return s.Gslb
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetNonce() *string {
	return s.Nonce
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetServerUserId() *string {
	return s.ServerUserId
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CreateChatSessionResponseBodyDataRtcParams) GetToken() *string {
	return s.Token
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetAppId(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.AppId = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetAvatarUserId(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.AvatarUserId = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetChannel(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.Channel = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetClientUserId(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.ClientUserId = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetGslb(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.Gslb = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetNonce(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.Nonce = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetServerUserId(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.ServerUserId = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetTimestamp(v int64) *CreateChatSessionResponseBodyDataRtcParams {
	s.Timestamp = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) SetToken(v string) *CreateChatSessionResponseBodyDataRtcParams {
	s.Token = &v
	return s
}

func (s *CreateChatSessionResponseBodyDataRtcParams) Validate() error {
	return dara.Validate(s)
}
