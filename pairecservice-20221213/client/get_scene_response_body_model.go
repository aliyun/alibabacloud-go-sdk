// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetSceneResponseBody
	GetDescription() *string
	SetFlows(v []*GetSceneResponseBodyFlows) *GetSceneResponseBody
	GetFlows() []*GetSceneResponseBodyFlows
	SetName(v string) *GetSceneResponseBody
	GetName() *string
	SetRequestId(v string) *GetSceneResponseBody
	GetRequestId() *string
}

type GetSceneResponseBody struct {
	// example:
	//
	// This is a test.
	Description *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*GetSceneResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B8987BF7-6028-5B17-80E0-251B7BD67BBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSceneResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSceneResponseBody) GetFlows() []*GetSceneResponseBodyFlows {
	return s.Flows
}

func (s *GetSceneResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSceneResponseBody) SetDescription(v string) *GetSceneResponseBody {
	s.Description = &v
	return s
}

func (s *GetSceneResponseBody) SetFlows(v []*GetSceneResponseBodyFlows) *GetSceneResponseBody {
	s.Flows = v
	return s
}

func (s *GetSceneResponseBody) SetName(v string) *GetSceneResponseBody {
	s.Name = &v
	return s
}

func (s *GetSceneResponseBody) SetRequestId(v string) *GetSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSceneResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSceneResponseBodyFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s GetSceneResponseBodyFlows) String() string {
	return dara.Prettify(s)
}

func (s GetSceneResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *GetSceneResponseBodyFlows) GetFlowCode() *string {
	return s.FlowCode
}

func (s *GetSceneResponseBodyFlows) GetFlowName() *string {
	return s.FlowName
}

func (s *GetSceneResponseBodyFlows) SetFlowCode(v string) *GetSceneResponseBodyFlows {
	s.FlowCode = &v
	return s
}

func (s *GetSceneResponseBodyFlows) SetFlowName(v string) *GetSceneResponseBodyFlows {
	s.FlowName = &v
	return s
}

func (s *GetSceneResponseBodyFlows) Validate() error {
	return dara.Validate(s)
}
