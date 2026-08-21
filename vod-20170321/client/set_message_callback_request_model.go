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
	// The application ID. If this parameter is not specified, the ID of the default application is used, which is the fixed value: **app-1000000**.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The authentication key. The key can be up to 32 characters in length and must contain uppercase letters, lowercase letters, and digits. This parameter can be set when the callback method is **HTTP**.
	//
	// example:
	//
	// Dsf346dvet
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The authentication switch for HTTP callbacks. This parameter takes effect only when the callback method is set to **HTTP**. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
	//
	// example:
	//
	// on
	AuthSwitch *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	// The callback method. Valid values:
	//
	// - **HTTP**
	//
	// - **Simple Message Queue (formerly MNS)**
	//
	// example:
	//
	// HTTP
	CallbackType *string `json:"CallbackType,omitempty" xml:"CallbackType,omitempty"`
	// The callback URL. This parameter is required when the callback method is set to **HTTP**.
	//
	// The callback URL cannot exceed 256 bytes in length. Multiple callback URLs are not supported.
	//
	// example:
	//
	// http://developer.aliyundoc.com
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The event types for callbacks. If this parameter is left empty, all notifications are disabled. If this parameter is set to **ALL**, all notifications are enabled. You can also specify specific event types, separated by commas (,). For the valid event types, see [Event types](https://help.aliyun.com/document_detail/55627.html).
	//
	// <props="china">
	//
	// > All AI-related events such as AIMediaAuditComplete and AIMediaDNAComplete use the value **AIComplete**.
	//
	// example:
	//
	// FileUploadComplete
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
	// The public endpoint of Simple Message Queue (formerly MNS). This parameter is required when the callback method is set to **Simple Message Queue (formerly MNS)**. Log on to the [Simple Message Queue (formerly MNS) console](https://account.aliyun.com/login/login.html) and click the **Get Endpoint*	- button in the upper-right corner to obtain the endpoint. For more information, see [Endpoint](https://help.aliyun.com/document_detail/27480.html).
	//
	// example:
	//
	// http://****.mns.cn-shanghai.aliyuncs.com/
	MnsEndpoint *string `json:"MnsEndpoint,omitempty" xml:"MnsEndpoint,omitempty"`
	// The name of the message queue. Log on to the [Simple Message Queue (formerly MNS) console](https://account.aliyun.com/login/login.html) and view the queue in the **Queue List**. This parameter is required when the callback method is set to **Simple Message Queue (formerly MNS)**.
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
