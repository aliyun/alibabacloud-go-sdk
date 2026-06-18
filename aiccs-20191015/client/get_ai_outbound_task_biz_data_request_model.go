// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskBizDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *GetAiOutboundTaskBizDataRequest
	GetChannelId() *string
	SetInstanceId(v string) *GetAiOutboundTaskBizDataRequest
	GetInstanceId() *string
}

type GetAiOutboundTaskBizDataRequest struct {
	// Session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it in the <b>Instance Management</b> section of the left-side navigation pane in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAiOutboundTaskBizDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskBizDataRequest) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskBizDataRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetAiOutboundTaskBizDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAiOutboundTaskBizDataRequest) SetChannelId(v string) *GetAiOutboundTaskBizDataRequest {
	s.ChannelId = &v
	return s
}

func (s *GetAiOutboundTaskBizDataRequest) SetInstanceId(v string) *GetAiOutboundTaskBizDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAiOutboundTaskBizDataRequest) Validate() error {
	return dara.Validate(s)
}
