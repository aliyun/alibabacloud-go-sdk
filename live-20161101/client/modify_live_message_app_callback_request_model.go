// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageAppCallbackRequest
	GetAppId() *string
	SetDataCenter(v string) *ModifyLiveMessageAppCallbackRequest
	GetDataCenter() *string
	SetEventCallbackUrl(v string) *ModifyLiveMessageAppCallbackRequest
	GetEventCallbackUrl() *string
}

type ModifyLiveMessageAppCallbackRequest struct {
	// The ID of the interactive messaging application whose callback settings you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The callback URL for events such as user logon, logoff, joining a group, and leaving a group. If you leave this parameter empty, callbacks are disabled. The callback URL must start with http:// or https:// and cannot contain a private IP address or a port number.
	//
	// example:
	//
	// http://example.aliyundoc.com/examplecallback
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty" xml:"EventCallbackUrl,omitempty"`
}

func (s ModifyLiveMessageAppCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppCallbackRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppCallbackRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageAppCallbackRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageAppCallbackRequest) GetEventCallbackUrl() *string {
	return s.EventCallbackUrl
}

func (s *ModifyLiveMessageAppCallbackRequest) SetAppId(v string) *ModifyLiveMessageAppCallbackRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackRequest) SetDataCenter(v string) *ModifyLiveMessageAppCallbackRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackRequest) SetEventCallbackUrl(v string) *ModifyLiveMessageAppCallbackRequest {
	s.EventCallbackUrl = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackRequest) Validate() error {
	return dara.Validate(s)
}
