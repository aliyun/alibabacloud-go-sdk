// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *ListNodesRequest
	GetBizName() *string
	SetNodeName(v string) *ListNodesRequest
	GetNodeName() *string
	SetOwner(v string) *ListNodesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListNodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodesRequest
	GetPageSize() *int32
	SetProgramType(v string) *ListNodesRequest
	GetProgramType() *string
	SetProjectEnv(v string) *ListNodesRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListNodesRequest
	GetProjectId() *int64
	SetSchedulerType(v string) *ListNodesRequest
	GetSchedulerType() *string
}

type ListNodesRequest struct {
	// The error code returned.
	//
	// example:
	//
	// test_bizName
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The ID of the baseline with which the node is associated.
	//
	// example:
	//
	// liux_test_n****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The description of the node.
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
	// The operation that you want to perform. Set the value to **ListNodes**.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The environment in which the node runs. Valid values: DEV and PROD.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The ID of the owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling type. Valid values:
	//
	// 	- NORMAL: Nodes are scheduled as expected.
	//
	// 	- PAUSE: Nodes are paused.
	//
	// 	- SKIP: Nodes are dry-run. Dry-run nodes are started as scheduled, but the system sets the status of the nodes to successful when it starts to run them.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetBizName() *string {
	return s.BizName
}

func (s *ListNodesRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *ListNodesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListNodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesRequest) GetProgramType() *string {
	return s.ProgramType
}

func (s *ListNodesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListNodesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodesRequest) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *ListNodesRequest) SetBizName(v string) *ListNodesRequest {
	s.BizName = &v
	return s
}

func (s *ListNodesRequest) SetNodeName(v string) *ListNodesRequest {
	s.NodeName = &v
	return s
}

func (s *ListNodesRequest) SetOwner(v string) *ListNodesRequest {
	s.Owner = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetProgramType(v string) *ListNodesRequest {
	s.ProgramType = &v
	return s
}

func (s *ListNodesRequest) SetProjectEnv(v string) *ListNodesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListNodesRequest) SetProjectId(v int64) *ListNodesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNodesRequest) SetSchedulerType(v string) *ListNodesRequest {
	s.SchedulerType = &v
	return s
}

func (s *ListNodesRequest) Validate() error {
	return dara.Validate(s)
}
