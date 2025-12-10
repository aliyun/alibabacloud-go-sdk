// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody
	GetJobs() []*ListJobsResponseBodyJobs
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
}

type ListJobsResponseBody struct {
	Jobs []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetJobs() []*ListJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsResponseBodyJobs struct {
	// example:
	//
	// 17677*******89598
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// EXECUTE_ALL
	ExecuteType *string `json:"ExecuteType,omitempty" xml:"ExecuteType,omitempty"`
	// example:
	//
	// draft-4x4iv3a9enuxw4vgka
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// job-mewqhd72nsrqujn1px
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// nodeId1
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// node-xdfasf8ewxfdaddl
	PaiflowNodeId *string `json:"PaiflowNodeId,omitempty" xml:"PaiflowNodeId,omitempty"`
	// example:
	//
	// flow-aayfouai80i980ncvz
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 94436
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) GetCreator() *string {
	return s.Creator
}

func (s *ListJobsResponseBodyJobs) GetExecuteType() *string {
	return s.ExecuteType
}

func (s *ListJobsResponseBodyJobs) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListJobsResponseBodyJobs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListJobsResponseBodyJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobs) GetNodeId() *string {
	return s.NodeId
}

func (s *ListJobsResponseBodyJobs) GetPaiflowNodeId() *string {
	return s.PaiflowNodeId
}

func (s *ListJobsResponseBodyJobs) GetRunId() *string {
	return s.RunId
}

func (s *ListJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListJobsResponseBodyJobs) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobsResponseBodyJobs) SetCreator(v string) *ListJobsResponseBodyJobs {
	s.Creator = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExecuteType(v string) *ListJobsResponseBodyJobs {
	s.ExecuteType = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExperimentId(v string) *ListJobsResponseBodyJobs {
	s.ExperimentId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetGmtCreateTime(v string) *ListJobsResponseBodyJobs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobId(v string) *ListJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetNodeId(v string) *ListJobsResponseBodyJobs {
	s.NodeId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetPaiflowNodeId(v string) *ListJobsResponseBodyJobs {
	s.PaiflowNodeId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetRunId(v string) *ListJobsResponseBodyJobs {
	s.RunId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatus(v string) *ListJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetWorkspaceId(v string) *ListJobsResponseBodyJobs {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}
