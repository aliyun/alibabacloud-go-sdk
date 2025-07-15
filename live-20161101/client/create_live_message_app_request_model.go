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
	// The name of the application. The name must be 2 to 16 characters in length.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The content moderation method. Valid values:
	//
	// 	- 0 (default): disables content moderation.
	//
	// 	- 1: uses built-in content moderation.
	//
	// 	- 2: uses custom content moderation.
	//
	// example:
	//
	// 2
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// The URL for content moderation. If you set AuditType to 2, you must specify this parameter. The URL must start with http:// or https:// and cannot contain a private IP address or a port number. For more information about custom content moderation, see the "Custom content moderation" section of this topic.
	//
	// example:
	//
	// http://demo.aliyundoc.com/exampleaudit
	AuditUrl *string `json:"AuditUrl,omitempty" xml:"AuditUrl,omitempty"`
	// The data center. Valid values:
	//
	// 	- cn-shanghai (default)
	//
	// 	- ap-southeast-1: Singapore
	//
	// >  When you call other operations to manage the interactive messaging application, you must specify the same data center in which the application is created.
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The callback URL for events, such as logon, logoff, and joining and leaving a group. If you leave this parameter empty, event callbacks are disabled. [](~~2672836~~)The callback URL must start with http:// or https:// and cannot contain a private IP address or a port number. For information about the callback message format and authentication logic, see the "Event callbacks" and "Callback authentication" sections of this topic.
	//
	// example:
	//
	// http://demo.aliyundoc.com/examplecallback
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty" xml:"EventCallbackUrl,omitempty"`
	// The retention period of group messages in the application. Valid values:
	//
	// 	- 0 (default): 30 days.
	//
	// 	- 1: 90 days.
	//
	// 	- 2: 180 days.
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
