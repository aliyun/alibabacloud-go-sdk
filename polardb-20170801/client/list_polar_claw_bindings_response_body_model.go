// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolarClawBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListPolarClawBindingsResponseBody
	GetApplicationId() *string
	SetBindings(v []*ListPolarClawBindingsResponseBodyBindings) *ListPolarClawBindingsResponseBody
	GetBindings() []*ListPolarClawBindingsResponseBodyBindings
	SetCode(v int32) *ListPolarClawBindingsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListPolarClawBindingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPolarClawBindingsResponseBody
	GetRequestId() *string
}

type ListPolarClawBindingsResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string                                      `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Bindings      []*ListPolarClawBindingsResponseBodyBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F2007D3-7E74-4ECB-89A8-BF130D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPolarClawBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolarClawBindingsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPolarClawBindingsResponseBody) GetBindings() []*ListPolarClawBindingsResponseBodyBindings {
	return s.Bindings
}

func (s *ListPolarClawBindingsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPolarClawBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPolarClawBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolarClawBindingsResponseBody) SetApplicationId(v string) *ListPolarClawBindingsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ListPolarClawBindingsResponseBody) SetBindings(v []*ListPolarClawBindingsResponseBodyBindings) *ListPolarClawBindingsResponseBody {
	s.Bindings = v
	return s
}

func (s *ListPolarClawBindingsResponseBody) SetCode(v int32) *ListPolarClawBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPolarClawBindingsResponseBody) SetMessage(v string) *ListPolarClawBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPolarClawBindingsResponseBody) SetRequestId(v string) *ListPolarClawBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolarClawBindingsResponseBody) Validate() error {
	if s.Bindings != nil {
		for _, item := range s.Bindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolarClawBindingsResponseBodyBindings struct {
	// Agent ID
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Channel ID
	//
	// example:
	//
	// feishu
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Account ID
	//
	// example:
	//
	// default
	ChannelAccountId *string `json:"ChannelAccountId,omitempty" xml:"ChannelAccountId,omitempty"`
}

func (s ListPolarClawBindingsResponseBodyBindings) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawBindingsResponseBodyBindings) GoString() string {
	return s.String()
}

func (s *ListPolarClawBindingsResponseBodyBindings) GetAgentId() *string {
	return s.AgentId
}

func (s *ListPolarClawBindingsResponseBodyBindings) GetChannel() *string {
	return s.Channel
}

func (s *ListPolarClawBindingsResponseBodyBindings) GetChannelAccountId() *string {
	return s.ChannelAccountId
}

func (s *ListPolarClawBindingsResponseBodyBindings) SetAgentId(v string) *ListPolarClawBindingsResponseBodyBindings {
	s.AgentId = &v
	return s
}

func (s *ListPolarClawBindingsResponseBodyBindings) SetChannel(v string) *ListPolarClawBindingsResponseBodyBindings {
	s.Channel = &v
	return s
}

func (s *ListPolarClawBindingsResponseBodyBindings) SetChannelAccountId(v string) *ListPolarClawBindingsResponseBodyBindings {
	s.ChannelAccountId = &v
	return s
}

func (s *ListPolarClawBindingsResponseBodyBindings) Validate() error {
	return dara.Validate(s)
}
