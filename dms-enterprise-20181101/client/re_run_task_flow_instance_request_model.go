// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReRunTaskFlowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ReRunTaskFlowInstanceRequest
	GetDagId() *int64
	SetDagInstanceId(v int64) *ReRunTaskFlowInstanceRequest
	GetDagInstanceId() *int64
	SetDagVersion(v string) *ReRunTaskFlowInstanceRequest
	GetDagVersion() *string
	SetTid(v int64) *ReRunTaskFlowInstanceRequest
	GetTid() *int64
}

type ReRunTaskFlowInstanceRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the execution record of the task flow. You can call the [ListTaskFlowInstance](https://help.aliyun.com/document_detail/424689.html) operation to query the execution record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47****
	DagInstanceId *int64 `json:"DagInstanceId,omitempty" xml:"DagInstanceId,omitempty"`
	// The version number of the task flow. You can call the ListTaskFlowVersions operation to query the version number of the task flow.
	//
	// example:
	//
	// 2****
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ReRunTaskFlowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReRunTaskFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReRunTaskFlowInstanceRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ReRunTaskFlowInstanceRequest) GetDagInstanceId() *int64 {
	return s.DagInstanceId
}

func (s *ReRunTaskFlowInstanceRequest) GetDagVersion() *string {
	return s.DagVersion
}

func (s *ReRunTaskFlowInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ReRunTaskFlowInstanceRequest) SetDagId(v int64) *ReRunTaskFlowInstanceRequest {
	s.DagId = &v
	return s
}

func (s *ReRunTaskFlowInstanceRequest) SetDagInstanceId(v int64) *ReRunTaskFlowInstanceRequest {
	s.DagInstanceId = &v
	return s
}

func (s *ReRunTaskFlowInstanceRequest) SetDagVersion(v string) *ReRunTaskFlowInstanceRequest {
	s.DagVersion = &v
	return s
}

func (s *ReRunTaskFlowInstanceRequest) SetTid(v int64) *ReRunTaskFlowInstanceRequest {
	s.Tid = &v
	return s
}

func (s *ReRunTaskFlowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
