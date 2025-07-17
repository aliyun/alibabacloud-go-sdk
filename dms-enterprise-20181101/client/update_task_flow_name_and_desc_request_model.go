// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowNameAndDescRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowNameAndDescRequest
	GetDagId() *int64
	SetDagName(v string) *UpdateTaskFlowNameAndDescRequest
	GetDagName() *string
	SetDescription(v string) *UpdateTaskFlowNameAndDescRequest
	GetDescription() *string
	SetTid(v int64) *UpdateTaskFlowNameAndDescRequest
	GetTid() *int64
}

type UpdateTaskFlowNameAndDescRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The new name that you want to specify for the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The description that you want to specify for the task flow.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowNameAndDescRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowNameAndDescRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowNameAndDescRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowNameAndDescRequest) GetDagName() *string {
	return s.DagName
}

func (s *UpdateTaskFlowNameAndDescRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTaskFlowNameAndDescRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowNameAndDescRequest) SetDagId(v int64) *UpdateTaskFlowNameAndDescRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescRequest) SetDagName(v string) *UpdateTaskFlowNameAndDescRequest {
	s.DagName = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescRequest) SetDescription(v string) *UpdateTaskFlowNameAndDescRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescRequest) SetTid(v int64) *UpdateTaskFlowNameAndDescRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescRequest) Validate() error {
	return dara.Validate(s)
}
