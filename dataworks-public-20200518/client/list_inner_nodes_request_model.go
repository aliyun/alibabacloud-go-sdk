// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInnerNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeName(v string) *ListInnerNodesRequest
	GetNodeName() *string
	SetOuterNodeId(v int64) *ListInnerNodesRequest
	GetOuterNodeId() *int64
	SetPageNumber(v int32) *ListInnerNodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInnerNodesRequest
	GetPageSize() *int32
	SetProgramType(v string) *ListInnerNodesRequest
	GetProgramType() *string
	SetProjectEnv(v string) *ListInnerNodesRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListInnerNodesRequest
	GetProjectId() *int64
}

type ListInnerNodesRequest struct {
	// The name of the node to which the inner nodes belong.
	//
	// example:
	//
	// liux_test_n****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the node group to which the inner nodes belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	OuterNodeId *int64 `json:"OuterNodeId,omitempty" xml:"OuterNodeId,omitempty"`
	// The page number. Valid values: 1 to 100.
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
	// The type of the node to which the inner nodes belong.
	//
	// Valid values: 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 23 (Data Integration), 24 (ODPS Script), 97 (PAI), 98 (node group), 99 (zero load), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (real-time synchronization), 1002 (PAI inner node), 1089 (cross-tenant collaboration), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (assignment), 1106 (for-each), and 1221 (PyODPS 3). You can call the ListNodes operation to query the type of the node.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The environment in which the node is run. Valid values: DEV and PROD. Default value: PROD.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListInnerNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInnerNodesRequest) GoString() string {
	return s.String()
}

func (s *ListInnerNodesRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *ListInnerNodesRequest) GetOuterNodeId() *int64 {
	return s.OuterNodeId
}

func (s *ListInnerNodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInnerNodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInnerNodesRequest) GetProgramType() *string {
	return s.ProgramType
}

func (s *ListInnerNodesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListInnerNodesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListInnerNodesRequest) SetNodeName(v string) *ListInnerNodesRequest {
	s.NodeName = &v
	return s
}

func (s *ListInnerNodesRequest) SetOuterNodeId(v int64) *ListInnerNodesRequest {
	s.OuterNodeId = &v
	return s
}

func (s *ListInnerNodesRequest) SetPageNumber(v int32) *ListInnerNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInnerNodesRequest) SetPageSize(v int32) *ListInnerNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInnerNodesRequest) SetProgramType(v string) *ListInnerNodesRequest {
	s.ProgramType = &v
	return s
}

func (s *ListInnerNodesRequest) SetProjectEnv(v string) *ListInnerNodesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListInnerNodesRequest) SetProjectId(v int64) *ListInnerNodesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListInnerNodesRequest) Validate() error {
	return dara.Validate(s)
}
