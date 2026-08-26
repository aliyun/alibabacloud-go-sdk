// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateLiveMessageAppRequest
	GetAppName() *string
	SetAuditType(v int32) *CreateLiveMessageAppRequest
	GetAuditType() *int32
	SetAuditUrl(v string) *CreateLiveMessageAppRequest
	GetAuditUrl() *string
	SetDataCenter(v string) *CreateLiveMessageAppRequest
	GetDataCenter() *string
	SetEventCallbackUrl(v string) *CreateLiveMessageAppRequest
	GetEventCallbackUrl() *string
	SetMsgLifeCycle(v int32) *CreateLiveMessageAppRequest
	GetMsgLifeCycle() *int32
}

type CreateLiveMessageAppRequest struct {
	// The application name. The name must be 2 to 16 characters in length.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The security audit mode. Valid values:
	//
	// - 0: default value. Security audit is disabled.
	//
	// - 1: built-in security audit.
	//
	// - 2: custom security audit.
	//
	// example:
	//
	// 2
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// The URL for custom security audit. This parameter is required when custom security audit is selected (AuditType=2). The URL must start with http:// or https://, must not contain private IP addresses, and must not include port numbers. For the format of custom security audit content, see the following section.
	//
	// example:
	//
	// http://demo.aliyundoc.com/exampleaudit
	AuditUrl *string `json:"AuditUrl,omitempty" xml:"AuditUrl,omitempty"`
	// The data center. Valid values:
	//
	// - cn-shanghai: default value. Shanghai.
	//
	// - ap-southeast-1: Singapore.
	//
	// > When calling other interactive messaging API operations, the data center must be the same as the one specified when creating the interactive messaging application.
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The event callback URL for client logon, logout, join group, and leave group events. If this parameter is empty, event callbacks are disabled. For the callback API operations that are triggered, see [Client access](https://help.aliyun.com/document_detail/2672836.html). The event callback URL must start with http:// or https://, must not contain private IP addresses, and must not include port numbers. For the event callback format and callback authentication logic, see the following section.
	//
	// example:
	//
	// http://demo.aliyundoc.com/examplecallback
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty" xml:"EventCallbackUrl,omitempty"`
	// The storage duration tier for group messages within the application. Valid values:
	//
	// - 0: default value. Messages are stored for 30 days.
	//
	// - 1: messages are stored for 90 days.
	//
	// - 2: messages are stored for 180 days.
	//
	// example:
	//
	// 1
	MsgLifeCycle *int32 `json:"MsgLifeCycle,omitempty" xml:"MsgLifeCycle,omitempty"`
}

func (s CreateLiveMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveMessageAppRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveMessageAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateLiveMessageAppRequest) GetAuditType() *int32 {
	return s.AuditType
}

func (s *CreateLiveMessageAppRequest) GetAuditUrl() *string {
	return s.AuditUrl
}

func (s *CreateLiveMessageAppRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateLiveMessageAppRequest) GetEventCallbackUrl() *string {
	return s.EventCallbackUrl
}

func (s *CreateLiveMessageAppRequest) GetMsgLifeCycle() *int32 {
	return s.MsgLifeCycle
}

func (s *CreateLiveMessageAppRequest) SetAppName(v string) *CreateLiveMessageAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateLiveMessageAppRequest) SetAuditType(v int32) *CreateLiveMessageAppRequest {
	s.AuditType = &v
	return s
}

func (s *CreateLiveMessageAppRequest) SetAuditUrl(v string) *CreateLiveMessageAppRequest {
	s.AuditUrl = &v
	return s
}

func (s *CreateLiveMessageAppRequest) SetDataCenter(v string) *CreateLiveMessageAppRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateLiveMessageAppRequest) SetEventCallbackUrl(v string) *CreateLiveMessageAppRequest {
	s.EventCallbackUrl = &v
	return s
}

func (s *CreateLiveMessageAppRequest) SetMsgLifeCycle(v int32) *CreateLiveMessageAppRequest {
	s.MsgLifeCycle = &v
	return s
}

func (s *CreateLiveMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
