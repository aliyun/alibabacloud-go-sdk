// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginBizdate(v string) *ListInstancesRequest
	GetBeginBizdate() *string
	SetBizName(v string) *ListInstancesRequest
	GetBizName() *string
	SetBizdate(v string) *ListInstancesRequest
	GetBizdate() *string
	SetDagId(v int64) *ListInstancesRequest
	GetDagId() *int64
	SetEndBizdate(v string) *ListInstancesRequest
	GetEndBizdate() *string
	SetNodeId(v int64) *ListInstancesRequest
	GetNodeId() *int64
	SetNodeName(v string) *ListInstancesRequest
	GetNodeName() *string
	SetOrderBy(v string) *ListInstancesRequest
	GetOrderBy() *string
	SetOwner(v string) *ListInstancesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetProgramType(v string) *ListInstancesRequest
	GetProgramType() *string
	SetProjectEnv(v string) *ListInstancesRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListInstancesRequest
	GetProjectId() *int64
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
}

type ListInstancesRequest struct {
	// The start date for which to retrieve the instance list. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-02-02 00:00:00
	BeginBizdate *string `json:"BeginBizdate,omitempty" xml:"BeginBizdate,omitempty"`
	// The name of the workflow. You can call [ListBusiness](https://help.aliyun.com/document_detail/173945.html) to query workflow information.
	//
	// example:
	//
	// test_bizName
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The date for which to retrieve the instance list. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-02-02 00:00:00
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The DAG ID. The DagId can be the DagId returned by operations such as [RunCycleDagNodes](https://help.aliyun.com/document_detail/212961.html) for data backfill, [RunSmokeTest](https://help.aliyun.com/document_detail/212949.html) for smoke testing, and [RunManualDagNodes](https://help.aliyun.com/document_detail/212830.html) for manual workflows.
	//
	// example:
	//
	// 11111
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The end date for which to retrieve the instance list. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-02-03 00:00:00
	EndBizdate *string `json:"EndBizdate,omitempty" xml:"EndBizdate,omitempty"`
	// The node ID. You can call [ListNodes](https://help.aliyun.com/document_detail/173979.html) to query the node ID.
	//
	// example:
	//
	// 100000000000
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node name. You can call [ListNodes](https://help.aliyun.com/document_detail/173979.html) to query the node name.
	//
	// example:
	//
	// openmr_8****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The sorting rule for the returned results. Valid values:
	//
	// - CREATE_TIME_DESC: sorted by creation time in descending order.
	//
	// - INSTANCE_ID_DESC: default value. Sorted by instance ID in descending order.
	//
	// example:
	//
	// INSTANCE_ID_DESC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The ID of the owner, which is the UID of the workspace administrator. You can logon to the Alibaba Cloud Management Console and view the UID in the Security Settings section of the storage management page.
	//
	// example:
	//
	// 193379****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The node type. You can call [ListNodes](https://help.aliyun.com/document_detail/173979.html) to query the node type.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The runtime environment. Valid values:
	//
	// - PROD: production environment.
	//
	// - DEV: development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID. You can call [ListProjects](https://help.aliyun.com/document_detail/178393.html) to query the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the node. Valid values:
	//
	// - NOT_RUN: The node is not run.
	//
	// - WAIT_TIME: The node is waiting for the scheduled time (DueTime or CycTime) to arrive.
	//
	// - WAIT_RESOURCE: The node is waiting for resources.
	//
	// - RUNNING: The node is running.
	//
	// - CHECKING: The node has been sent to Data Quality for data validation.
	//
	// - CHECKING_CONDITION: The node is undergoing branch condition verification.
	//
	// - FAILURE: Failed to execute.
	//
	// - SUCCESS: Execute successfully.
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetBeginBizdate() *string {
	return s.BeginBizdate
}

func (s *ListInstancesRequest) GetBizName() *string {
	return s.BizName
}

func (s *ListInstancesRequest) GetBizdate() *string {
	return s.Bizdate
}

func (s *ListInstancesRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ListInstancesRequest) GetEndBizdate() *string {
	return s.EndBizdate
}

func (s *ListInstancesRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListInstancesRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *ListInstancesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListInstancesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetProgramType() *string {
	return s.ProgramType
}

func (s *ListInstancesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListInstancesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) SetBeginBizdate(v string) *ListInstancesRequest {
	s.BeginBizdate = &v
	return s
}

func (s *ListInstancesRequest) SetBizName(v string) *ListInstancesRequest {
	s.BizName = &v
	return s
}

func (s *ListInstancesRequest) SetBizdate(v string) *ListInstancesRequest {
	s.Bizdate = &v
	return s
}

func (s *ListInstancesRequest) SetDagId(v int64) *ListInstancesRequest {
	s.DagId = &v
	return s
}

func (s *ListInstancesRequest) SetEndBizdate(v string) *ListInstancesRequest {
	s.EndBizdate = &v
	return s
}

func (s *ListInstancesRequest) SetNodeId(v int64) *ListInstancesRequest {
	s.NodeId = &v
	return s
}

func (s *ListInstancesRequest) SetNodeName(v string) *ListInstancesRequest {
	s.NodeName = &v
	return s
}

func (s *ListInstancesRequest) SetOrderBy(v string) *ListInstancesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListInstancesRequest) SetOwner(v string) *ListInstancesRequest {
	s.Owner = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetProgramType(v string) *ListInstancesRequest {
	s.ProgramType = &v
	return s
}

func (s *ListInstancesRequest) SetProjectEnv(v string) *ListInstancesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListInstancesRequest) SetProjectId(v int64) *ListInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
