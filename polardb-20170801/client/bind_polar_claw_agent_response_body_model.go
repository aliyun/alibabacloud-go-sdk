// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPolarClawAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *BindPolarClawAgentResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *BindPolarClawAgentResponseBody
	GetApplicationId() *string
	SetBinding(v *BindPolarClawAgentResponseBodyBinding) *BindPolarClawAgentResponseBody
	GetBinding() *BindPolarClawAgentResponseBodyBinding
	SetCode(v int32) *BindPolarClawAgentResponseBody
	GetCode() *int32
	SetMessage(v string) *BindPolarClawAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindPolarClawAgentResponseBody
	GetRequestId() *string
	SetTotalBindings(v int32) *BindPolarClawAgentResponseBody
	GetTotalBindings() *int32
}

type BindPolarClawAgentResponseBody struct {
	// The agent ID.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// pa-********************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Details of the newly created binding.
	Binding *BindPolarClawAgentResponseBodyBinding `json:"Binding,omitempty" xml:"Binding,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// A message that indicates the request result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of global bindings after the operation.
	//
	// example:
	//
	// 1
	TotalBindings *int32 `json:"TotalBindings,omitempty" xml:"TotalBindings,omitempty"`
}

func (s BindPolarClawAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindPolarClawAgentResponseBody) GoString() string {
	return s.String()
}

func (s *BindPolarClawAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *BindPolarClawAgentResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *BindPolarClawAgentResponseBody) GetBinding() *BindPolarClawAgentResponseBodyBinding {
	return s.Binding
}

func (s *BindPolarClawAgentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindPolarClawAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindPolarClawAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindPolarClawAgentResponseBody) GetTotalBindings() *int32 {
	return s.TotalBindings
}

func (s *BindPolarClawAgentResponseBody) SetAgentId(v string) *BindPolarClawAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *BindPolarClawAgentResponseBody) SetApplicationId(v string) *BindPolarClawAgentResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *BindPolarClawAgentResponseBody) SetBinding(v *BindPolarClawAgentResponseBodyBinding) *BindPolarClawAgentResponseBody {
	s.Binding = v
	return s
}

func (s *BindPolarClawAgentResponseBody) SetCode(v int32) *BindPolarClawAgentResponseBody {
	s.Code = &v
	return s
}

func (s *BindPolarClawAgentResponseBody) SetMessage(v string) *BindPolarClawAgentResponseBody {
	s.Message = &v
	return s
}

func (s *BindPolarClawAgentResponseBody) SetRequestId(v string) *BindPolarClawAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPolarClawAgentResponseBody) SetTotalBindings(v int32) *BindPolarClawAgentResponseBody {
	s.TotalBindings = &v
	return s
}

func (s *BindPolarClawAgentResponseBody) Validate() error {
	if s.Binding != nil {
		if err := s.Binding.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindPolarClawAgentResponseBodyBinding struct {
	// The account ID.
	//
	// example:
	//
	// default
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The agent ID.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// feishu
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s BindPolarClawAgentResponseBodyBinding) String() string {
	return dara.Prettify(s)
}

func (s BindPolarClawAgentResponseBodyBinding) GoString() string {
	return s.String()
}

func (s *BindPolarClawAgentResponseBodyBinding) GetAccountId() *string {
	return s.AccountId
}

func (s *BindPolarClawAgentResponseBodyBinding) GetAgentId() *string {
	return s.AgentId
}

func (s *BindPolarClawAgentResponseBodyBinding) GetChannel() *string {
	return s.Channel
}

func (s *BindPolarClawAgentResponseBodyBinding) SetAccountId(v string) *BindPolarClawAgentResponseBodyBinding {
	s.AccountId = &v
	return s
}

func (s *BindPolarClawAgentResponseBodyBinding) SetAgentId(v string) *BindPolarClawAgentResponseBodyBinding {
	s.AgentId = &v
	return s
}

func (s *BindPolarClawAgentResponseBodyBinding) SetChannel(v string) *BindPolarClawAgentResponseBodyBinding {
	s.Channel = &v
	return s
}

func (s *BindPolarClawAgentResponseBodyBinding) Validate() error {
	return dara.Validate(s)
}
