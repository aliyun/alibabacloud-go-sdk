// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeletePolarClawAgentResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *DeletePolarClawAgentResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DeletePolarClawAgentResponseBody
	GetCode() *int32
	SetMessage(v string) *DeletePolarClawAgentResponseBody
	GetMessage() *string
	SetRemovedBindings(v int32) *DeletePolarClawAgentResponseBody
	GetRemovedBindings() *int32
	SetRequestId(v string) *DeletePolarClawAgentResponseBody
	GetRequestId() *string
}

type DeletePolarClawAgentResponseBody struct {
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// pa-**************
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
	RemovedBindings *int32 `json:"RemovedBindings,omitempty" xml:"RemovedBindings,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolarClawAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarClawAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *DeletePolarClawAgentResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeletePolarClawAgentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeletePolarClawAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolarClawAgentResponseBody) GetRemovedBindings() *int32 {
	return s.RemovedBindings
}

func (s *DeletePolarClawAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarClawAgentResponseBody) SetAgentId(v string) *DeletePolarClawAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *DeletePolarClawAgentResponseBody) SetApplicationId(v string) *DeletePolarClawAgentResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DeletePolarClawAgentResponseBody) SetCode(v int32) *DeletePolarClawAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolarClawAgentResponseBody) SetMessage(v string) *DeletePolarClawAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolarClawAgentResponseBody) SetRemovedBindings(v int32) *DeletePolarClawAgentResponseBody {
	s.RemovedBindings = &v
	return s
}

func (s *DeletePolarClawAgentResponseBody) SetRequestId(v string) *DeletePolarClawAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarClawAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
