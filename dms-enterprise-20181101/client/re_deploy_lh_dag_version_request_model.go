// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReDeployLhDagVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ReDeployLhDagVersionRequest
	GetDagId() *int64
	SetDagVersion(v int64) *ReDeployLhDagVersionRequest
	GetDagVersion() *int64
	SetTid(v int64) *ReDeployLhDagVersionRequest
	GetTid() *int64
}

type ReDeployLhDagVersionRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the ID of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the task flow version. You can call the [ListDAGVersions](https://help.aliyun.com/document_detail/424682.html) operation to obtain the ID of the task flow version.
	//
	// example:
	//
	// 2****
	DagVersion *int64 `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ReDeployLhDagVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReDeployLhDagVersionRequest) GoString() string {
	return s.String()
}

func (s *ReDeployLhDagVersionRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ReDeployLhDagVersionRequest) GetDagVersion() *int64 {
	return s.DagVersion
}

func (s *ReDeployLhDagVersionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ReDeployLhDagVersionRequest) SetDagId(v int64) *ReDeployLhDagVersionRequest {
	s.DagId = &v
	return s
}

func (s *ReDeployLhDagVersionRequest) SetDagVersion(v int64) *ReDeployLhDagVersionRequest {
	s.DagVersion = &v
	return s
}

func (s *ReDeployLhDagVersionRequest) SetTid(v int64) *ReDeployLhDagVersionRequest {
	s.Tid = &v
	return s
}

func (s *ReDeployLhDagVersionRequest) Validate() error {
	return dara.Validate(s)
}
