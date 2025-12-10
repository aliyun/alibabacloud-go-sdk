// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperiments(v []*ListExperimentsResponseBodyExperiments) *ListExperimentsResponseBody
	GetExperiments() []*ListExperimentsResponseBodyExperiments
	SetRequestId(v string) *ListExperimentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListExperimentsResponseBody
	GetTotalCount() *int64
}

type ListExperimentsResponseBody struct {
	Experiments []*ListExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	// example:
	//
	// 9708FB85-232F-5F9A-9D67-7F9CCCE20E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBody) GetExperiments() []*ListExperimentsResponseBodyExperiments {
	return s.Experiments
}

func (s *ListExperimentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExperimentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExperimentsResponseBody) SetExperiments(v []*ListExperimentsResponseBodyExperiments) *ListExperimentsResponseBody {
	s.Experiments = v
	return s
}

func (s *ListExperimentsResponseBody) SetRequestId(v string) *ListExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentsResponseBody) SetTotalCount(v int64) *ListExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExperimentsResponseBody) Validate() error {
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

type ListExperimentsResponseBodyExperiments struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// 1326*******76250
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// Pipeline draft description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// draft-rbvg5wzljzjhc9ks92
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 12
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 23487
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentsResponseBodyExperiments) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyExperiments) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListExperimentsResponseBodyExperiments) GetCreator() *string {
	return s.Creator
}

func (s *ListExperimentsResponseBodyExperiments) GetDescription() *string {
	return s.Description
}

func (s *ListExperimentsResponseBodyExperiments) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListExperimentsResponseBodyExperiments) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListExperimentsResponseBodyExperiments) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListExperimentsResponseBodyExperiments) GetName() *string {
	return s.Name
}

func (s *ListExperimentsResponseBodyExperiments) GetSource() *string {
	return s.Source
}

func (s *ListExperimentsResponseBodyExperiments) GetVersion() *int64 {
	return s.Version
}

func (s *ListExperimentsResponseBodyExperiments) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListExperimentsResponseBodyExperiments) SetAccessibility(v string) *ListExperimentsResponseBodyExperiments {
	s.Accessibility = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetCreator(v string) *ListExperimentsResponseBodyExperiments {
	s.Creator = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDescription(v string) *ListExperimentsResponseBodyExperiments {
	s.Description = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtCreateTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtCreateTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtModifiedTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetName(v string) *ListExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetSource(v string) *ListExperimentsResponseBodyExperiments {
	s.Source = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetVersion(v int64) *ListExperimentsResponseBodyExperiments {
	s.Version = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetWorkspaceId(v string) *ListExperimentsResponseBodyExperiments {
	s.WorkspaceId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) Validate() error {
	return dara.Validate(s)
}
