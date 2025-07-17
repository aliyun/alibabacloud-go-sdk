// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowOwnerRequest
	GetDagId() *int64
	SetNewOwnerId(v string) *UpdateTaskFlowOwnerRequest
	GetNewOwnerId() *string
	SetTid(v int64) *UpdateTaskFlowOwnerRequest
	GetTid() *int64
}

type UpdateTaskFlowOwnerRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlowInstance](https://help.aliyun.com/document_detail/424689.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The user ID of the new owner. You can call the [GetUser](https://help.aliyun.com/document_detail/147098.html) or [ListUsers](https://help.aliyun.com/document_detail/141938.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51****
	NewOwnerId *string `json:"NewOwnerId,omitempty" xml:"NewOwnerId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowOwnerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowOwnerRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowOwnerRequest) GetNewOwnerId() *string {
	return s.NewOwnerId
}

func (s *UpdateTaskFlowOwnerRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowOwnerRequest) SetDagId(v int64) *UpdateTaskFlowOwnerRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowOwnerRequest) SetNewOwnerId(v string) *UpdateTaskFlowOwnerRequest {
	s.NewOwnerId = &v
	return s
}

func (s *UpdateTaskFlowOwnerRequest) SetTid(v int64) *UpdateTaskFlowOwnerRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowOwnerRequest) Validate() error {
	return dara.Validate(s)
}
