// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveTaskFlowToScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *MoveTaskFlowToScenarioRequest
	GetDagId() *int64
	SetScenarioId(v int64) *MoveTaskFlowToScenarioRequest
	GetScenarioId() *int64
	SetTid(v int64) *MoveTaskFlowToScenarioRequest
	GetTid() *int64
}

type MoveTaskFlowToScenarioRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the business scenario to which you want to migrate your task flow. If this parameter is set to the default value or a value that is less than or equal to 0, the task flow is migrated to the default business scenario.
	//
	// example:
	//
	// 1****
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s MoveTaskFlowToScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveTaskFlowToScenarioRequest) GoString() string {
	return s.String()
}

func (s *MoveTaskFlowToScenarioRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *MoveTaskFlowToScenarioRequest) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *MoveTaskFlowToScenarioRequest) GetTid() *int64 {
	return s.Tid
}

func (s *MoveTaskFlowToScenarioRequest) SetDagId(v int64) *MoveTaskFlowToScenarioRequest {
	s.DagId = &v
	return s
}

func (s *MoveTaskFlowToScenarioRequest) SetScenarioId(v int64) *MoveTaskFlowToScenarioRequest {
	s.ScenarioId = &v
	return s
}

func (s *MoveTaskFlowToScenarioRequest) SetTid(v int64) *MoveTaskFlowToScenarioRequest {
	s.Tid = &v
	return s
}

func (s *MoveTaskFlowToScenarioRequest) Validate() error {
	return dara.Validate(s)
}
