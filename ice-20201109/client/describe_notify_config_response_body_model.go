// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioOssPath(v string) *DescribeNotifyConfigResponseBody
	GetAudioOssPath() *string
	SetCallbackUrl(v string) *DescribeNotifyConfigResponseBody
	GetCallbackUrl() *string
	SetEnableAudioRecording(v bool) *DescribeNotifyConfigResponseBody
	GetEnableAudioRecording() *bool
	SetEnableNotify(v bool) *DescribeNotifyConfigResponseBody
	GetEnableNotify() *bool
	SetEventTypes(v string) *DescribeNotifyConfigResponseBody
	GetEventTypes() *string
	SetRequestId(v string) *DescribeNotifyConfigResponseBody
	GetRequestId() *string
	SetToken(v string) *DescribeNotifyConfigResponseBody
	GetToken() *string
}

type DescribeNotifyConfigResponseBody struct {
	AudioOssPath *string `json:"AudioOssPath,omitempty" xml:"AudioOssPath,omitempty"`
	// example:
	//
	// http://customer.com/callback
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	EnableAudioRecording *bool   `json:"EnableAudioRecording,omitempty" xml:"EnableAudioRecording,omitempty"`
	// example:
	//
	// true
	EnableNotify *bool `json:"EnableNotify,omitempty" xml:"EnableNotify,omitempty"`
	// The event types. If this parameter is empty, all event types are selected.
	//
	// 	- agent_start: The agent is started.
	//
	// 	- agent_stop: The agent is stopped.
	//
	// 	- error: An error occurred.
	//
	// example:
	//
	// agent_start,agent_stop,error
	EventTypes *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// eyJhcHBpZCI6ICIxMjM0MTIzNxxxxx
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotifyConfigResponseBody) GetAudioOssPath() *string {
	return s.AudioOssPath
}

func (s *DescribeNotifyConfigResponseBody) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *DescribeNotifyConfigResponseBody) GetEnableAudioRecording() *bool {
	return s.EnableAudioRecording
}

func (s *DescribeNotifyConfigResponseBody) GetEnableNotify() *bool {
	return s.EnableNotify
}

func (s *DescribeNotifyConfigResponseBody) GetEventTypes() *string {
	return s.EventTypes
}

func (s *DescribeNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNotifyConfigResponseBody) GetToken() *string {
	return s.Token
}

func (s *DescribeNotifyConfigResponseBody) SetAudioOssPath(v string) *DescribeNotifyConfigResponseBody {
	s.AudioOssPath = &v
	return s
}

func (s *DescribeNotifyConfigResponseBody) SetCallbackUrl(v string) *DescribeNotifyConfigResponseBody {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeNotifyConfigResponseBody) SetEnableAudioRecording(v bool) *DescribeNotifyConfigResponseBody {
	s.EnableAudioRecording = &v
	return s
}

func (s *DescribeNotifyConfigResponseBody) SetEnableNotify(v bool) *DescribeNotifyConfigResponseBody {
	s.EnableNotify = &v
	return s
}

func (s *DescribeNotifyConfigResponseBody) SetEventTypes(v string) *DescribeNotifyConfigResponseBody {
	s.EventTypes = &v
	return s
}

func (s *DescribeNotifyConfigResponseBody) SetRequestId(v string) *DescribeNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNotifyConfigResponseBody) SetToken(v string) *DescribeNotifyConfigResponseBody {
	s.Token = &v
	return s
}

func (s *DescribeNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
