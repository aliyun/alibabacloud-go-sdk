// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentExperimentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperiments(v []*ListRecentExperimentsResponseBodyExperiments) *ListRecentExperimentsResponseBody
	GetExperiments() []*ListRecentExperimentsResponseBodyExperiments
	SetRequestId(v string) *ListRecentExperimentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRecentExperimentsResponseBody
	GetTotalCount() *int64
}

type ListRecentExperimentsResponseBody struct {
	Experiments []*ListRecentExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecentExperimentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecentExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsResponseBody) GetExperiments() []*ListRecentExperimentsResponseBodyExperiments {
	return s.Experiments
}

func (s *ListRecentExperimentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecentExperimentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRecentExperimentsResponseBody) SetExperiments(v []*ListRecentExperimentsResponseBodyExperiments) *ListRecentExperimentsResponseBody {
	s.Experiments = v
	return s
}

func (s *ListRecentExperimentsResponseBody) SetRequestId(v string) *ListRecentExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentExperimentsResponseBody) SetTotalCount(v int64) *ListRecentExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecentExperimentsResponseBody) Validate() error {
	if s.Experiments != nil {
		for _, item := range s.Experiments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecentExperimentsResponseBodyExperiments struct {
	// example:
	//
	// Pipeline draft description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// draft-76p70ye0gwv3vbur55
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 10
	ModifyCnt *int64 `json:"ModifyCnt,omitempty" xml:"ModifyCnt,omitempty"`
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	RecentGmtModifiedTime *string `json:"RecentGmtModifiedTime,omitempty" xml:"RecentGmtModifiedTime,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 15821
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// Workspace name
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListRecentExperimentsResponseBodyExperiments) String() string {
	return dara.Prettify(s)
}

func (s ListRecentExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetDescription() *string {
	return s.Description
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetModifyCnt() *int64 {
	return s.ModifyCnt
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetName() *string {
	return s.Name
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetRecentGmtModifiedTime() *string {
	return s.RecentGmtModifiedTime
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetSource() *string {
	return s.Source
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListRecentExperimentsResponseBodyExperiments) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetDescription(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.Description = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetModifyCnt(v int64) *ListRecentExperimentsResponseBodyExperiments {
	s.ModifyCnt = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetName(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetRecentGmtModifiedTime(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.RecentGmtModifiedTime = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetSource(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.Source = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetWorkspaceId(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.WorkspaceId = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetWorkspaceName(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.WorkspaceName = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) Validate() error {
	return dara.Validate(s)
}
