// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateSceneRequest
	GetDescription() *string
	SetFlows(v []*UpdateSceneRequestFlows) *UpdateSceneRequest
	GetFlows() []*UpdateSceneRequestFlows
	SetInstanceId(v string) *UpdateSceneRequest
	GetInstanceId() *string
	SetName(v string) *UpdateSceneRequest
	GetName() *string
}

type UpdateSceneRequest struct {
	// example:
	//
	// This is a test.
	Description *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*UpdateSceneRequestFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSceneRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSceneRequest) GetFlows() []*UpdateSceneRequestFlows {
	return s.Flows
}

func (s *UpdateSceneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSceneRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSceneRequest) SetDescription(v string) *UpdateSceneRequest {
	s.Description = &v
	return s
}

func (s *UpdateSceneRequest) SetFlows(v []*UpdateSceneRequestFlows) *UpdateSceneRequest {
	s.Flows = v
	return s
}

func (s *UpdateSceneRequest) SetInstanceId(v string) *UpdateSceneRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSceneRequest) SetName(v string) *UpdateSceneRequest {
	s.Name = &v
	return s
}

func (s *UpdateSceneRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateSceneRequestFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s UpdateSceneRequestFlows) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneRequestFlows) GoString() string {
	return s.String()
}

func (s *UpdateSceneRequestFlows) GetFlowCode() *string {
	return s.FlowCode
}

func (s *UpdateSceneRequestFlows) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateSceneRequestFlows) SetFlowCode(v string) *UpdateSceneRequestFlows {
	s.FlowCode = &v
	return s
}

func (s *UpdateSceneRequestFlows) SetFlowName(v string) *UpdateSceneRequestFlows {
	s.FlowName = &v
	return s
}

func (s *UpdateSceneRequestFlows) Validate() error {
	return dara.Validate(s)
}
