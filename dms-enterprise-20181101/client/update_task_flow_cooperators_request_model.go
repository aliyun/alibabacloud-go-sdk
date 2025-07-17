// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowCooperatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCooperatorIds(v []*string) *UpdateTaskFlowCooperatorsRequest
	GetCooperatorIds() []*string
	SetDagId(v int64) *UpdateTaskFlowCooperatorsRequest
	GetDagId() *int64
	SetTid(v int64) *UpdateTaskFlowCooperatorsRequest
	GetTid() *int64
}

type UpdateTaskFlowCooperatorsRequest struct {
	// The IDs of the users who are involved in the task flow to be updated.
	CooperatorIds []*string `json:"CooperatorIds,omitempty" xml:"CooperatorIds,omitempty" type:"Repeated"`
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant.
	//
	// > :To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowCooperatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowCooperatorsRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowCooperatorsRequest) GetCooperatorIds() []*string {
	return s.CooperatorIds
}

func (s *UpdateTaskFlowCooperatorsRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowCooperatorsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowCooperatorsRequest) SetCooperatorIds(v []*string) *UpdateTaskFlowCooperatorsRequest {
	s.CooperatorIds = v
	return s
}

func (s *UpdateTaskFlowCooperatorsRequest) SetDagId(v int64) *UpdateTaskFlowCooperatorsRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsRequest) SetTid(v int64) *UpdateTaskFlowCooperatorsRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsRequest) Validate() error {
	return dara.Validate(s)
}
