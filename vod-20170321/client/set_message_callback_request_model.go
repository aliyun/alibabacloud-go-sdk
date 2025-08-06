// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMessageCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SetMessageCallbackRequest
	GetAppId() *string
	SetAuthKey(v string) *SetMessageCallbackRequest
	GetAuthKey() *string
	SetAuthSwitch(v string) *SetMessageCallbackRequest
	GetAuthSwitch() *string
	SetCallbackType(v string) *SetMessageCallbackRequest
	GetCallbackType() *string
	SetCallbackURL(v string) *SetMessageCallbackRequest
	GetCallbackURL() *string
	SetEventTypeList(v string) *SetMessageCallbackRequest
	GetEventTypeList() *string
	SetMnsEndpoint(v string) *SetMessageCallbackRequest
	GetMnsEndpoint() *string
	SetMnsQueueName(v string) *SetMessageCallbackRequest
	GetMnsQueueName() *string
	SetOwnerAccount(v string) *SetMessageCallbackRequest
	GetOwnerAccount() *string
}

type SetMessageCallbackRequest struct {
	// The ID of the application. If you leave this parameter empty, the default value **app-1000000*	- is used.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The authentication key. The key can be up to 32 characters in length and must contain uppercase letters, lowercase letters, and digits. This parameter takes effect only when you set CallbackType to **HTTP**.
	//
	// example:
	//
	// Dsf346dvet
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// Specifies whether to enable callback authentication. This parameter takes effect only when you set CallbackType to **HTTP**. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	AuthSwitch *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	// The callback method. Valid values:
	//
	// 	- **HTTP**
	//
	// 	- **Simple Message Queue(formerly MNS)**
	//
	// example:
	//
	// HTTP
	CallbackType *string `json:"CallbackType,omitempty" xml:"CallbackType,omitempty"`
	// The callback URL. This parameter is required if you set CallbackType to **HTTP**. The callback URL cannot exceed 256 bytes in length. You can specify only one callback URL.
	//
	// example:
	//
	// http://developer.aliyundoc.com
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The type of the callback event. If you do not set this parameter, notifications for all types of events are disabled. If you set this parameter to **ALL**, notifications for all types of events are enabled. You can specify the event types for which notifications are enabled. Separate multiple event types with commas (,). For more information about the valid values of this parameter, see [Overview](https://help.aliyun.com/document_detail/55627.html).
	//
	// example:
	//
	// FileUploadComplete
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
	// The public endpoint of Message Service (MNS). This parameter only takes effect when the CallbackType parameter is set to **Simple Message Queue(formerly MNS)**. To obtain the public endpoint, log on to the [Simple Message Queue(formerly MNS) console](https://account.aliyun.com/login/login.html) and click **Get Endpoint*	- in the upper-right corner of the Topics page. For more information, see [Endpoint](https://help.aliyun.com/document_detail/27480.html).
	//
	// example:
	//
	// http://****.mns.cn-shanghai.aliyuncs.com/
	MnsEndpoint *string `json:"MnsEndpoint,omitempty" xml:"MnsEndpoint,omitempty"`
	// The name of the Simple Message Queue(formerly MNS). You can obtain the name of the Simple Message Queue(formerly MNS) on the **Queues*	- page in the [Simple Message Queue(formerly MNS) console](https://account.aliyun.com/login/login.html). This parameter is required when you set CallbackType to **Simple Message Queue(formerly MNS)**.
	//
	// example:
	//
	// quene_name
	MnsQueueName *string `json:"MnsQueueName,omitempty" xml:"MnsQueueName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s SetMessageCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s SetMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *SetMessageCallbackRequest) GetAppId() *string {
	return s.AppId
}

func (s *SetMessageCallbackRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *SetMessageCallbackRequest) GetAuthSwitch() *string {
	return s.AuthSwitch
}

func (s *SetMessageCallbackRequest) GetCallbackType() *string {
	return s.CallbackType
}

func (s *SetMessageCallbackRequest) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *SetMessageCallbackRequest) GetEventTypeList() *string {
	return s.EventTypeList
}

func (s *SetMessageCallbackRequest) GetMnsEndpoint() *string {
	return s.MnsEndpoint
}

func (s *SetMessageCallbackRequest) GetMnsQueueName() *string {
	return s.MnsQueueName
}

func (s *SetMessageCallbackRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetMessageCallbackRequest) SetAppId(v string) *SetMessageCallbackRequest {
	s.AppId = &v
	return s
}

func (s *SetMessageCallbackRequest) SetAuthKey(v string) *SetMessageCallbackRequest {
	s.AuthKey = &v
	return s
}

func (s *SetMessageCallbackRequest) SetAuthSwitch(v string) *SetMessageCallbackRequest {
	s.AuthSwitch = &v
	return s
}

func (s *SetMessageCallbackRequest) SetCallbackType(v string) *SetMessageCallbackRequest {
	s.CallbackType = &v
	return s
}

func (s *SetMessageCallbackRequest) SetCallbackURL(v string) *SetMessageCallbackRequest {
	s.CallbackURL = &v
	return s
}

func (s *SetMessageCallbackRequest) SetEventTypeList(v string) *SetMessageCallbackRequest {
	s.EventTypeList = &v
	return s
}

func (s *SetMessageCallbackRequest) SetMnsEndpoint(v string) *SetMessageCallbackRequest {
	s.MnsEndpoint = &v
	return s
}

func (s *SetMessageCallbackRequest) SetMnsQueueName(v string) *SetMessageCallbackRequest {
	s.MnsQueueName = &v
	return s
}

func (s *SetMessageCallbackRequest) SetOwnerAccount(v string) *SetMessageCallbackRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetMessageCallbackRequest) Validate() error {
	return dara.Validate(s)
}
