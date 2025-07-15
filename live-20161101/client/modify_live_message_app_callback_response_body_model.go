// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageAppCallbackResponseBody
	GetAppId() *string
	SetAppSign(v string) *ModifyLiveMessageAppCallbackResponseBody
	GetAppSign() *string
	SetEventCallbackNeedAuthentication(v bool) *ModifyLiveMessageAppCallbackResponseBody
	GetEventCallbackNeedAuthentication() *bool
	SetEventCallbackUrl(v string) *ModifyLiveMessageAppCallbackResponseBody
	GetEventCallbackUrl() *string
	SetRequestId(v string) *ModifyLiveMessageAppCallbackResponseBody
	GetRequestId() *string
}

type ModifyLiveMessageAppCallbackResponseBody struct {
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The signature of the interactive messaging application. It is required by the interactive messaging SDK.
	//
	// example:
	//
	// **************************************************************************
	AppSign *string `json:"AppSign,omitempty" xml:"AppSign,omitempty"`
	// Indicates whether authentication is required for event callbacks. Default value: true.
	//
	// example:
	//
	// true
	EventCallbackNeedAuthentication *bool `json:"EventCallbackNeedAuthentication,omitempty" xml:"EventCallbackNeedAuthentication,omitempty"`
	// The callback URL for events such as user logon, logoff, joining a group, and leaving a group. This parameter is not returned if it has an empty value.
	//
	// example:
	//
	// http://example.aliyundoc.com/examplecallback
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty" xml:"EventCallbackUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1F0FFEAD-B7D5-1D4A-A6B9-8C63ADF6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveMessageAppCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppCallbackResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageAppCallbackResponseBody) GetAppSign() *string {
	return s.AppSign
}

func (s *ModifyLiveMessageAppCallbackResponseBody) GetEventCallbackNeedAuthentication() *bool {
	return s.EventCallbackNeedAuthentication
}

func (s *ModifyLiveMessageAppCallbackResponseBody) GetEventCallbackUrl() *string {
	return s.EventCallbackUrl
}

func (s *ModifyLiveMessageAppCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveMessageAppCallbackResponseBody) SetAppId(v string) *ModifyLiveMessageAppCallbackResponseBody {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponseBody) SetAppSign(v string) *ModifyLiveMessageAppCallbackResponseBody {
	s.AppSign = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponseBody) SetEventCallbackNeedAuthentication(v bool) *ModifyLiveMessageAppCallbackResponseBody {
	s.EventCallbackNeedAuthentication = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponseBody) SetEventCallbackUrl(v string) *ModifyLiveMessageAppCallbackResponseBody {
	s.EventCallbackUrl = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponseBody) SetRequestId(v string) *ModifyLiveMessageAppCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
