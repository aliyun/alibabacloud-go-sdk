// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLoginDetailsResponseBody
	GetCode() *string
	SetData(v *GetLoginDetailsResponseBodyData) *GetLoginDetailsResponseBody
	GetData() *GetLoginDetailsResponseBodyData
	SetHttpStatusCode(v int32) *GetLoginDetailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLoginDetailsResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetLoginDetailsResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetLoginDetailsResponseBody
	GetRequestId() *string
}

type GetLoginDetailsResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Logon information data.
	Data *GetLoginDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Response parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// F1A4774A-F28B-5C40-AEF6-D88D2DD6C7E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLoginDetailsResponseBody) GetData() *GetLoginDetailsResponseBodyData {
	return s.Data
}

func (s *GetLoginDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLoginDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLoginDetailsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetLoginDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoginDetailsResponseBody) SetCode(v string) *GetLoginDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetLoginDetailsResponseBody) SetData(v *GetLoginDetailsResponseBodyData) *GetLoginDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetLoginDetailsResponseBody) SetHttpStatusCode(v int32) *GetLoginDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLoginDetailsResponseBody) SetMessage(v string) *GetLoginDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetLoginDetailsResponseBody) SetParams(v []*string) *GetLoginDetailsResponseBody {
	s.Params = v
	return s
}

func (s *GetLoginDetailsResponseBody) SetRequestId(v string) *GetLoginDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoginDetailsResponseBodyData struct {
	// Agent data push server URL. After a successful connection, the server pushes agent status data to the agent Workbench.
	//
	// example:
	//
	// sh-wss-ccc.aliyuncs.com
	AgentServerUrl *string `json:"AgentServerUrl,omitempty" xml:"AgentServerUrl,omitempty"`
	// Agent profile picture URL.
	//
	// example:
	//
	// http://abc.com/sam.jpb
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// Chat session application ID.
	//
	// example:
	//
	// 7pjxxx
	ChatAppId *string `json:"ChatAppId,omitempty" xml:"ChatAppId,omitempty"`
	// Chat session application key.
	//
	// example:
	//
	// 955e4bd7xxxxxxxxxxxxxd7898ba9fa0d0
	ChatAppKey *string `json:"ChatAppKey,omitempty" xml:"ChatAppKey,omitempty"`
	// Chat Device ID.
	//
	// example:
	//
	// 4c51c9116c36537cb850dc1081d745df
	ChatDeviceId *string `json:"ChatDeviceId,omitempty" xml:"ChatDeviceId,omitempty"`
	// Chat session logon token.
	//
	// example:
	//
	// {"accessToken":"oauth_cloud_key:4xxx0=","accessTokenExpiredTime":864xxxx,"refreshToken":"oauth_cloud_key:Paxxxx=="}
	ChatLoginToken *string `json:"ChatLoginToken,omitempty" xml:"ChatLoginToken,omitempty"`
	// Chat session server URL.
	//
	// example:
	//
	// wss://wss.im.dingtalk.cn
	ChatServerUrl *string `json:"ChatServerUrl,omitempty" xml:"ChatServerUrl,omitempty"`
	// Chat session Device ID.
	//
	// example:
	//
	// A-8000100x
	ChatUserId *string `json:"ChatUserId,omitempty" xml:"ChatUserId,omitempty"`
	// SIP device extension number. This field is required when using a SIP phone.
	//
	// example:
	//
	// 8033****
	DeviceExt *string `json:"DeviceExt,omitempty" xml:"DeviceExt,omitempty"`
	// If the agent has registered an SIP phone, this parameter is the Device ID of the SIP phone device; otherwise, it is empty.
	//
	// example:
	//
	// Yealink SIP-T23G 44.84.XX.XX
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// SIP device status.
	//
	// example:
	//
	// OFFLINE
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// Agent display name.
	//
	// example:
	//
	// 坐席小王
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Agent extension number.
	//
	// example:
	//
	// 8001****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Agent nickname
	//
	// example:
	//
	// 老王
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// Signature information generated using an asymmetric key encryption algorithm. This information is required for authentication when registering with the SIP server.
	//
	// example:
	//
	// zi31STIMtIfa/UN2l+6lww****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Signature information generated using an asymmetric key encryption algorithm. This information is required for authentication when registering with the SIP server.
	//
	// example:
	//
	// zi31STIMtIfa/UN2l+6lww****
	Signature2 *string `json:"Signature2,omitempty" xml:"Signature2,omitempty"`
	// SIP registration server URL. The agent must register with the server as an SIP user to perform transactional processing (TP) operations.
	//
	// example:
	//
	// sh-sip-ccc.aliyuncs.com:443
	SipServerUrl *string `json:"SipServerUrl,omitempty" xml:"SipServerUrl,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Signature data used to generate signature information. Because this signature mechanism does not support Chinese characters, UserKey2 is currently used for signing instead.
	//
	// example:
	//
	// 坐席@aliyun:802001:1656406628862
	UserKey *string `json:"UserKey,omitempty" xml:"UserKey,omitempty"`
	// Signature data used to generate signature information. UserKey2 was added because UserKey does not support Chinese characters in signatures.
	//
	// example:
	//
	// 802001:1656406628862"
	UserKey2 *string `json:"UserKey2,omitempty" xml:"UserKey2,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s GetLoginDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLoginDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsResponseBodyData) GetAgentServerUrl() *string {
	return s.AgentServerUrl
}

func (s *GetLoginDetailsResponseBodyData) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetLoginDetailsResponseBodyData) GetChatAppId() *string {
	return s.ChatAppId
}

func (s *GetLoginDetailsResponseBodyData) GetChatAppKey() *string {
	return s.ChatAppKey
}

func (s *GetLoginDetailsResponseBodyData) GetChatDeviceId() *string {
	return s.ChatDeviceId
}

func (s *GetLoginDetailsResponseBodyData) GetChatLoginToken() *string {
	return s.ChatLoginToken
}

func (s *GetLoginDetailsResponseBodyData) GetChatServerUrl() *string {
	return s.ChatServerUrl
}

func (s *GetLoginDetailsResponseBodyData) GetChatUserId() *string {
	return s.ChatUserId
}

func (s *GetLoginDetailsResponseBodyData) GetDeviceExt() *string {
	return s.DeviceExt
}

func (s *GetLoginDetailsResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetLoginDetailsResponseBodyData) GetDeviceState() *string {
	return s.DeviceState
}

func (s *GetLoginDetailsResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetLoginDetailsResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *GetLoginDetailsResponseBodyData) GetNickname() *string {
	return s.Nickname
}

func (s *GetLoginDetailsResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetLoginDetailsResponseBodyData) GetSignature2() *string {
	return s.Signature2
}

func (s *GetLoginDetailsResponseBodyData) GetSipServerUrl() *string {
	return s.SipServerUrl
}

func (s *GetLoginDetailsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetLoginDetailsResponseBodyData) GetUserKey() *string {
	return s.UserKey
}

func (s *GetLoginDetailsResponseBodyData) GetUserKey2() *string {
	return s.UserKey2
}

func (s *GetLoginDetailsResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *GetLoginDetailsResponseBodyData) SetAgentServerUrl(v string) *GetLoginDetailsResponseBodyData {
	s.AgentServerUrl = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetAvatarUrl(v string) *GetLoginDetailsResponseBodyData {
	s.AvatarUrl = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetChatAppId(v string) *GetLoginDetailsResponseBodyData {
	s.ChatAppId = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetChatAppKey(v string) *GetLoginDetailsResponseBodyData {
	s.ChatAppKey = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetChatDeviceId(v string) *GetLoginDetailsResponseBodyData {
	s.ChatDeviceId = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetChatLoginToken(v string) *GetLoginDetailsResponseBodyData {
	s.ChatLoginToken = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetChatServerUrl(v string) *GetLoginDetailsResponseBodyData {
	s.ChatServerUrl = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetChatUserId(v string) *GetLoginDetailsResponseBodyData {
	s.ChatUserId = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetDeviceExt(v string) *GetLoginDetailsResponseBodyData {
	s.DeviceExt = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetDeviceId(v string) *GetLoginDetailsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetDeviceState(v string) *GetLoginDetailsResponseBodyData {
	s.DeviceState = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetDisplayName(v string) *GetLoginDetailsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetExtension(v string) *GetLoginDetailsResponseBodyData {
	s.Extension = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetNickname(v string) *GetLoginDetailsResponseBodyData {
	s.Nickname = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetSignature(v string) *GetLoginDetailsResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetSignature2(v string) *GetLoginDetailsResponseBodyData {
	s.Signature2 = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetSipServerUrl(v string) *GetLoginDetailsResponseBodyData {
	s.SipServerUrl = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetUserId(v string) *GetLoginDetailsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetUserKey(v string) *GetLoginDetailsResponseBodyData {
	s.UserKey = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetUserKey2(v string) *GetLoginDetailsResponseBodyData {
	s.UserKey2 = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetWorkMode(v string) *GetLoginDetailsResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
