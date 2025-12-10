// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v string) *GetJobResponseBody
	GetArguments() *string
	SetCreator(v string) *GetJobResponseBody
	GetCreator() *string
	SetExecuteType(v string) *GetJobResponseBody
	GetExecuteType() *string
	SetExperimentId(v string) *GetJobResponseBody
	GetExperimentId() *string
	SetGmtCreateTime(v string) *GetJobResponseBody
	GetGmtCreateTime() *string
	SetJobId(v string) *GetJobResponseBody
	GetJobId() *string
	SetNodeId(v string) *GetJobResponseBody
	GetNodeId() *string
	SetPaiflowNodeId(v string) *GetJobResponseBody
	GetPaiflowNodeId() *string
	SetRequestId(v string) *GetJobResponseBody
	GetRequestId() *string
	SetRunId(v string) *GetJobResponseBody
	GetRunId() *string
	SetRunInfo(v string) *GetJobResponseBody
	GetRunInfo() *string
	SetSnapshot(v string) *GetJobResponseBody
	GetSnapshot() *string
	SetStatus(v string) *GetJobResponseBody
	GetStatus() *string
	SetWorkspaceId(v string) *GetJobResponseBody
	GetWorkspaceId() *string
}

type GetJobResponseBody struct {
	// example:
	//
	// ---
	//
	// arguments:
	//
	//   parameters:
	//
	//   - name: "execution_maxcompute"
	//
	//     value:
	//
	//       endpoint: "http://service.cn-shanghai.maxcompute.aliyun.com/api"
	//
	//       odpsProject: "xxxxxxx"
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// example:
	//
	// 13266*******76250
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// EXECUTE_TO_HERE
	ExecuteType *string `json:"ExecuteType,omitempty" xml:"ExecuteType,omitempty"`
	// example:
	//
	// draft-rbvg5wzljzjhc9ks92
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// job-6xfhrofqx93y139fg3
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// id-xxxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// node-5dzotrnwyzfxmbwwc8
	PaiflowNodeId *string `json:"PaiflowNodeId,omitempty" xml:"PaiflowNodeId,omitempty"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// flow-lcs8ez11961l6urc3m
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// apiVersion: "core/v1"
	//
	// metadata:
	//
	//   version: "v1"
	//
	//   identifier: "1de8f7c8-a9d8-4433-9e87-d95979ccab14"
	//
	//   annotations: {}
	//
	// spec:
	//
	//   inputs:
	//
	//     artifacts: []
	//
	//     parameters:
	//
	//     - name: "execution"
	//
	//       type: "Map"
	//
	//   arguments:
	//
	//     artifacts: []
	//
	//     parameters: []
	//
	//   dependencies: []
	//
	//   initContainers: []
	//
	//   sideCarContainers: []
	//
	//   pipelines:
	//
	//   - apiVersion: "core/v1"
	//
	//     metadata:
	//
	//       provider: "pai"
	//
	//       name: "id-2d88-1608982098027-91558"
	//
	//       version: "v1"
	//
	//       identifier: "type_transform"
	//
	//       annotations: {}
	RunInfo *string `json:"RunInfo,omitempty" xml:"RunInfo,omitempty"`
	// example:
	//
	// {}
	Snapshot *string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 86995
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetArguments() *string {
	return s.Arguments
}

func (s *GetJobResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetJobResponseBody) GetExecuteType() *string {
	return s.ExecuteType
}

func (s *GetJobResponseBody) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *GetJobResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetJobResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *GetJobResponseBody) GetPaiflowNodeId() *string {
	return s.PaiflowNodeId
}

func (s *GetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *GetJobResponseBody) GetRunInfo() *string {
	return s.RunInfo
}

func (s *GetJobResponseBody) GetSnapshot() *string {
	return s.Snapshot
}

func (s *GetJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetJobResponseBody) SetArguments(v string) *GetJobResponseBody {
	s.Arguments = &v
	return s
}

func (s *GetJobResponseBody) SetCreator(v string) *GetJobResponseBody {
	s.Creator = &v
	return s
}

func (s *GetJobResponseBody) SetExecuteType(v string) *GetJobResponseBody {
	s.ExecuteType = &v
	return s
}

func (s *GetJobResponseBody) SetExperimentId(v string) *GetJobResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetJobResponseBody) SetGmtCreateTime(v string) *GetJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBody) SetJobId(v string) *GetJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBody) SetNodeId(v string) *GetJobResponseBody {
	s.NodeId = &v
	return s
}

func (s *GetJobResponseBody) SetPaiflowNodeId(v string) *GetJobResponseBody {
	s.PaiflowNodeId = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetRunId(v string) *GetJobResponseBody {
	s.RunId = &v
	return s
}

func (s *GetJobResponseBody) SetRunInfo(v string) *GetJobResponseBody {
	s.RunInfo = &v
	return s
}

func (s *GetJobResponseBody) SetSnapshot(v string) *GetJobResponseBody {
	s.Snapshot = &v
	return s
}

func (s *GetJobResponseBody) SetStatus(v string) *GetJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetJobResponseBody) SetWorkspaceId(v string) *GetJobResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetJobResponseBody) Validate() error {
	return dara.Validate(s)
}
