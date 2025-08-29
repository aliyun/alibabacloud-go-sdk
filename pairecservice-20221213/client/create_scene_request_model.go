// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSceneRequest
	GetDescription() *string
	SetFlows(v []*CreateSceneRequestFlows) *CreateSceneRequest
	GetFlows() []*CreateSceneRequestFlows
	SetInstanceId(v string) *CreateSceneRequest
	GetInstanceId() *string
	SetName(v string) *CreateSceneRequest
	GetName() *string
}

type CreateSceneRequest struct {
	// example:
	//
	// This is a test.
	Description *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*CreateSceneRequestFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSceneRequest) GetFlows() []*CreateSceneRequestFlows {
	return s.Flows
}

func (s *CreateSceneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSceneRequest) GetName() *string {
	return s.Name
}

func (s *CreateSceneRequest) SetDescription(v string) *CreateSceneRequest {
	s.Description = &v
	return s
}

func (s *CreateSceneRequest) SetFlows(v []*CreateSceneRequestFlows) *CreateSceneRequest {
	s.Flows = v
	return s
}

func (s *CreateSceneRequest) SetInstanceId(v string) *CreateSceneRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSceneRequest) SetName(v string) *CreateSceneRequest {
	s.Name = &v
	return s
}

func (s *CreateSceneRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSceneRequestFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s CreateSceneRequestFlows) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneRequestFlows) GoString() string {
	return s.String()
}

func (s *CreateSceneRequestFlows) GetFlowCode() *string {
	return s.FlowCode
}

func (s *CreateSceneRequestFlows) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateSceneRequestFlows) SetFlowCode(v string) *CreateSceneRequestFlows {
	s.FlowCode = &v
	return s
}

func (s *CreateSceneRequestFlows) SetFlowName(v string) *CreateSceneRequestFlows {
	s.FlowName = &v
	return s
}

func (s *CreateSceneRequestFlows) Validate() error {
	return dara.Validate(s)
}
