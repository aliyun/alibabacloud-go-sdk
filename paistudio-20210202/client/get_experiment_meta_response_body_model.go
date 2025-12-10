// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetExperimentMetaResponseBody
	GetAccessibility() *string
	SetCreator(v string) *GetExperimentMetaResponseBody
	GetCreator() *string
	SetDescription(v string) *GetExperimentMetaResponseBody
	GetDescription() *string
	SetExperimentId(v string) *GetExperimentMetaResponseBody
	GetExperimentId() *string
	SetGmtCreateTime(v string) *GetExperimentMetaResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetExperimentMetaResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *GetExperimentMetaResponseBody
	GetName() *string
	SetOptions(v string) *GetExperimentMetaResponseBody
	GetOptions() *string
	SetRequestId(v string) *GetExperimentMetaResponseBody
	GetRequestId() *string
	SetSource(v string) *GetExperimentMetaResponseBody
	GetSource() *string
	SetVersion(v string) *GetExperimentMetaResponseBody
	GetVersion() *string
	SetWorkspaceId(v string) *GetExperimentMetaResponseBody
	GetWorkspaceId() *string
}

type GetExperimentMetaResponseBody struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// 1326****76250
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
	// {"mlflow":{"experimentId":"exp-1"}}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 12
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 23487
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentMetaResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetExperimentMetaResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetExperimentMetaResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetExperimentMetaResponseBody) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *GetExperimentMetaResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetExperimentMetaResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetExperimentMetaResponseBody) GetName() *string {
	return s.Name
}

func (s *GetExperimentMetaResponseBody) GetOptions() *string {
	return s.Options
}

func (s *GetExperimentMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentMetaResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetExperimentMetaResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetExperimentMetaResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetExperimentMetaResponseBody) SetAccessibility(v string) *GetExperimentMetaResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetCreator(v string) *GetExperimentMetaResponseBody {
	s.Creator = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetDescription(v string) *GetExperimentMetaResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetExperimentId(v string) *GetExperimentMetaResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetGmtCreateTime(v string) *GetExperimentMetaResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetGmtModifiedTime(v string) *GetExperimentMetaResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetName(v string) *GetExperimentMetaResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetOptions(v string) *GetExperimentMetaResponseBody {
	s.Options = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetRequestId(v string) *GetExperimentMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetSource(v string) *GetExperimentMetaResponseBody {
	s.Source = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetVersion(v string) *GetExperimentMetaResponseBody {
	s.Version = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetWorkspaceId(v string) *GetExperimentMetaResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetExperimentMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
