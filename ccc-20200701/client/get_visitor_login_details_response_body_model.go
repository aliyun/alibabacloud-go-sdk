// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVisitorLoginDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVisitorLoginDetailsResponseBody
	GetCode() *string
	SetData(v *GetVisitorLoginDetailsResponseBodyData) *GetVisitorLoginDetailsResponseBody
	GetData() *GetVisitorLoginDetailsResponseBodyData
	SetHttpStatusCode(v int32) *GetVisitorLoginDetailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVisitorLoginDetailsResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetVisitorLoginDetailsResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetVisitorLoginDetailsResponseBody
	GetRequestId() *string
}

type GetVisitorLoginDetailsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetVisitorLoginDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 19D09CCC-F298-4124-849A-AFA217819011
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVisitorLoginDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVisitorLoginDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetVisitorLoginDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVisitorLoginDetailsResponseBody) GetData() *GetVisitorLoginDetailsResponseBodyData {
	return s.Data
}

func (s *GetVisitorLoginDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVisitorLoginDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVisitorLoginDetailsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetVisitorLoginDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVisitorLoginDetailsResponseBody) SetCode(v string) *GetVisitorLoginDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBody) SetData(v *GetVisitorLoginDetailsResponseBodyData) *GetVisitorLoginDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetVisitorLoginDetailsResponseBody) SetHttpStatusCode(v int32) *GetVisitorLoginDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBody) SetMessage(v string) *GetVisitorLoginDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBody) SetParams(v []*string) *GetVisitorLoginDetailsResponseBody {
	s.Params = v
	return s
}

func (s *GetVisitorLoginDetailsResponseBody) SetRequestId(v string) *GetVisitorLoginDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVisitorLoginDetailsResponseBodyData struct {
	// example:
	//
	// 7pjxxx
	ChatAppId *string `json:"ChatAppId,omitempty" xml:"ChatAppId,omitempty"`
	// example:
	//
	// 955e4bd7xxxxxxxxxxxxxd7898ba9fa0d0
	ChatAppKey *string `json:"ChatAppKey,omitempty" xml:"ChatAppKey,omitempty"`
	// example:
	//
	// 4c51c9116c36537cb850dc1081d745df
	ChatDeviceId *string `json:"ChatDeviceId,omitempty" xml:"ChatDeviceId,omitempty"`
	// example:
	//
	// {"accessToken":"oauth_cloud_key:+4oJXUAFSWxGD2YuRW4V/oUN0/8qJGNc0I*********n1E3DOr3Q3lX00ZnTpyqRi8Y6hYoLYA7n2ZkWuv485hVtXeSgnIQkKxXPbMgwoLxWaK//lI5Dn/mb4YuDifigv+ZyFzc+07vxm9ZFu/NjA==","accessTokenExpiredTime":86400000,"refreshToken":"oauth_cloud_key:/U+8UueDmpeUszhXC+SWow4pNLZp2C***********U/377BNXF+Mjo1lFgDk6GtEjNNoJpapX2mHH8GcRke2+yKQs/w4gAN9xSMn543Ciung+93pXV6IpQGbEVlu"}
	ChatLoginToken *string `json:"ChatLoginToken,omitempty" xml:"ChatLoginToken,omitempty"`
	// example:
	//
	// wss://wss.im.dingtalk.cn
	ChatServerUrl *string `json:"ChatServerUrl,omitempty" xml:"ChatServerUrl,omitempty"`
	// example:
	//
	// dac9c001****a15684ea91a81317
	ChatUserId *string `json:"ChatUserId,omitempty" xml:"ChatUserId,omitempty"`
}

func (s GetVisitorLoginDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVisitorLoginDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVisitorLoginDetailsResponseBodyData) GetChatAppId() *string {
	return s.ChatAppId
}

func (s *GetVisitorLoginDetailsResponseBodyData) GetChatAppKey() *string {
	return s.ChatAppKey
}

func (s *GetVisitorLoginDetailsResponseBodyData) GetChatDeviceId() *string {
	return s.ChatDeviceId
}

func (s *GetVisitorLoginDetailsResponseBodyData) GetChatLoginToken() *string {
	return s.ChatLoginToken
}

func (s *GetVisitorLoginDetailsResponseBodyData) GetChatServerUrl() *string {
	return s.ChatServerUrl
}

func (s *GetVisitorLoginDetailsResponseBodyData) GetChatUserId() *string {
	return s.ChatUserId
}

func (s *GetVisitorLoginDetailsResponseBodyData) SetChatAppId(v string) *GetVisitorLoginDetailsResponseBodyData {
	s.ChatAppId = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBodyData) SetChatAppKey(v string) *GetVisitorLoginDetailsResponseBodyData {
	s.ChatAppKey = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBodyData) SetChatDeviceId(v string) *GetVisitorLoginDetailsResponseBodyData {
	s.ChatDeviceId = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBodyData) SetChatLoginToken(v string) *GetVisitorLoginDetailsResponseBodyData {
	s.ChatLoginToken = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBodyData) SetChatServerUrl(v string) *GetVisitorLoginDetailsResponseBodyData {
	s.ChatServerUrl = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBodyData) SetChatUserId(v string) *GetVisitorLoginDetailsResponseBodyData {
	s.ChatUserId = &v
	return s
}

func (s *GetVisitorLoginDetailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
