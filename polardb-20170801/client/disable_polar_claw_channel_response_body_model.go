// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisablePolarClawChannelResponseBody
	GetApplicationId() *string
	SetChannelId(v string) *DisablePolarClawChannelResponseBody
	GetChannelId() *string
	SetCode(v int32) *DisablePolarClawChannelResponseBody
	GetCode() *int32
	SetMessage(v string) *DisablePolarClawChannelResponseBody
	GetMessage() *string
	SetOk(v bool) *DisablePolarClawChannelResponseBody
	GetOk() *bool
	SetRequestId(v string) *DisablePolarClawChannelResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *DisablePolarClawChannelResponseBody
	GetRestarted() *bool
}

type DisablePolarClawChannelResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// feishu
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s DisablePolarClawChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DisablePolarClawChannelResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisablePolarClawChannelResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *DisablePolarClawChannelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisablePolarClawChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisablePolarClawChannelResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *DisablePolarClawChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisablePolarClawChannelResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *DisablePolarClawChannelResponseBody) SetApplicationId(v string) *DisablePolarClawChannelResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DisablePolarClawChannelResponseBody) SetChannelId(v string) *DisablePolarClawChannelResponseBody {
	s.ChannelId = &v
	return s
}

func (s *DisablePolarClawChannelResponseBody) SetCode(v int32) *DisablePolarClawChannelResponseBody {
	s.Code = &v
	return s
}

func (s *DisablePolarClawChannelResponseBody) SetMessage(v string) *DisablePolarClawChannelResponseBody {
	s.Message = &v
	return s
}

func (s *DisablePolarClawChannelResponseBody) SetOk(v bool) *DisablePolarClawChannelResponseBody {
	s.Ok = &v
	return s
}

func (s *DisablePolarClawChannelResponseBody) SetRequestId(v string) *DisablePolarClawChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisablePolarClawChannelResponseBody) SetRestarted(v bool) *DisablePolarClawChannelResponseBody {
	s.Restarted = &v
	return s
}

func (s *DisablePolarClawChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
