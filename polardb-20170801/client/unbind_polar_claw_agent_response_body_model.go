// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPolarClawAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UnbindPolarClawAgentResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *UnbindPolarClawAgentResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UnbindPolarClawAgentResponseBody
	GetCode() *int32
	SetMessage(v string) *UnbindPolarClawAgentResponseBody
	GetMessage() *string
	SetRemovedCount(v int32) *UnbindPolarClawAgentResponseBody
	GetRemovedCount() *int32
	SetRequestId(v string) *UnbindPolarClawAgentResponseBody
	GetRequestId() *string
	SetTotalBindings(v int32) *UnbindPolarClawAgentResponseBody
	GetTotalBindings() *int32
}

type UnbindPolarClawAgentResponseBody struct {
	// Agent ID
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// pa-********************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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
	// 1
	RemovedCount *int32 `json:"RemovedCount,omitempty" xml:"RemovedCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	TotalBindings *int32 `json:"TotalBindings,omitempty" xml:"TotalBindings,omitempty"`
}

func (s UnbindPolarClawAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindPolarClawAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPolarClawAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *UnbindPolarClawAgentResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UnbindPolarClawAgentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnbindPolarClawAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindPolarClawAgentResponseBody) GetRemovedCount() *int32 {
	return s.RemovedCount
}

func (s *UnbindPolarClawAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindPolarClawAgentResponseBody) GetTotalBindings() *int32 {
	return s.TotalBindings
}

func (s *UnbindPolarClawAgentResponseBody) SetAgentId(v string) *UnbindPolarClawAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *UnbindPolarClawAgentResponseBody) SetApplicationId(v string) *UnbindPolarClawAgentResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UnbindPolarClawAgentResponseBody) SetCode(v int32) *UnbindPolarClawAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindPolarClawAgentResponseBody) SetMessage(v string) *UnbindPolarClawAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindPolarClawAgentResponseBody) SetRemovedCount(v int32) *UnbindPolarClawAgentResponseBody {
	s.RemovedCount = &v
	return s
}

func (s *UnbindPolarClawAgentResponseBody) SetRequestId(v string) *UnbindPolarClawAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindPolarClawAgentResponseBody) SetTotalBindings(v int32) *UnbindPolarClawAgentResponseBody {
	s.TotalBindings = &v
	return s
}

func (s *UnbindPolarClawAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
