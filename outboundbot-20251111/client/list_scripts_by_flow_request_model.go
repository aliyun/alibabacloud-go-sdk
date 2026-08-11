// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsByFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *ListScriptsByFlowRequest
	GetFlowId() *string
	SetInstanceId(v string) *ListScriptsByFlowRequest
	GetInstanceId() *string
}

type ListScriptsByFlowRequest struct {
	// The flow ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b42
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListScriptsByFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsByFlowRequest) GoString() string {
	return s.String()
}

func (s *ListScriptsByFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *ListScriptsByFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptsByFlowRequest) SetFlowId(v string) *ListScriptsByFlowRequest {
	s.FlowId = &v
	return s
}

func (s *ListScriptsByFlowRequest) SetInstanceId(v string) *ListScriptsByFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptsByFlowRequest) Validate() error {
	return dara.Validate(s)
}
