// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagName(v string) *CreateTaskFlowRequest
	GetDagName() *string
	SetDescription(v string) *CreateTaskFlowRequest
	GetDescription() *string
	SetScenarioId(v int64) *CreateTaskFlowRequest
	GetScenarioId() *int64
	SetTid(v int64) *CreateTaskFlowRequest
	GetTid() *int64
}

type CreateTaskFlowRequest struct {
	// The name of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dag_Name
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The description of the task flow.
	//
	// example:
	//
	// zht_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the scenario.
	//
	// example:
	//
	// 2**
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateTaskFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskFlowRequest) GetDagName() *string {
	return s.DagName
}

func (s *CreateTaskFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTaskFlowRequest) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *CreateTaskFlowRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateTaskFlowRequest) SetDagName(v string) *CreateTaskFlowRequest {
	s.DagName = &v
	return s
}

func (s *CreateTaskFlowRequest) SetDescription(v string) *CreateTaskFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateTaskFlowRequest) SetScenarioId(v int64) *CreateTaskFlowRequest {
	s.ScenarioId = &v
	return s
}

func (s *CreateTaskFlowRequest) SetTid(v int64) *CreateTaskFlowRequest {
	s.Tid = &v
	return s
}

func (s *CreateTaskFlowRequest) Validate() error {
	return dara.Validate(s)
}
