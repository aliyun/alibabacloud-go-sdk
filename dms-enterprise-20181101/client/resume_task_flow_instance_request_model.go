// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskFlowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ResumeTaskFlowInstanceRequest
	GetDagId() *int64
	SetDagInstanceId(v int64) *ResumeTaskFlowInstanceRequest
	GetDagInstanceId() *int64
	SetDagVersion(v string) *ResumeTaskFlowInstanceRequest
	GetDagVersion() *string
	SetTid(v int64) *ResumeTaskFlowInstanceRequest
	GetTid() *int64
}

type ResumeTaskFlowInstanceRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the execution record of the task flow. You can call the [ListTaskFlowInstance](https://help.aliyun.com/document_detail/424689.html) operation to query the execution record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3****
	DagInstanceId *int64 `json:"DagInstanceId,omitempty" xml:"DagInstanceId,omitempty"`
	// The version number of the task flow. You can call the [ListDAGVersions](https://help.aliyun.com/document_detail/424682.html) operation to query the version number.
	//
	// example:
	//
	// []
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ResumeTaskFlowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeTaskFlowInstanceRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ResumeTaskFlowInstanceRequest) GetDagInstanceId() *int64 {
	return s.DagInstanceId
}

func (s *ResumeTaskFlowInstanceRequest) GetDagVersion() *string {
	return s.DagVersion
}

func (s *ResumeTaskFlowInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ResumeTaskFlowInstanceRequest) SetDagId(v int64) *ResumeTaskFlowInstanceRequest {
	s.DagId = &v
	return s
}

func (s *ResumeTaskFlowInstanceRequest) SetDagInstanceId(v int64) *ResumeTaskFlowInstanceRequest {
	s.DagInstanceId = &v
	return s
}

func (s *ResumeTaskFlowInstanceRequest) SetDagVersion(v string) *ResumeTaskFlowInstanceRequest {
	s.DagVersion = &v
	return s
}

func (s *ResumeTaskFlowInstanceRequest) SetTid(v int64) *ResumeTaskFlowInstanceRequest {
	s.Tid = &v
	return s
}

func (s *ResumeTaskFlowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
