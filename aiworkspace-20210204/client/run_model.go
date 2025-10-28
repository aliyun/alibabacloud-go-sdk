// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRun interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Run
	GetAccessibility() *string
	SetExperimentId(v string) *Run
	GetExperimentId() *string
	SetGmtCreateTime(v string) *Run
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Run
	GetGmtModifiedTime() *string
	SetLabels(v []*RunLabel) *Run
	GetLabels() []*RunLabel
	SetMetrics(v []*RunMetric) *Run
	GetMetrics() []*RunMetric
	SetName(v string) *Run
	GetName() *string
	SetOwnerId(v string) *Run
	GetOwnerId() *string
	SetParams(v []*RunParam) *Run
	GetParams() []*RunParam
	SetRequestId(v string) *Run
	GetRequestId() *string
	SetRunId(v string) *Run
	GetRunId() *string
	SetSourceId(v string) *Run
	GetSourceId() *string
	SetSourceType(v string) *Run
	GetSourceType() *string
	SetUserId(v string) *Run
	GetUserId() *string
	SetWorkspaceId(v string) *Run
	GetWorkspaceId() *string
}

type Run struct {
	Accessibility   *string      `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ExperimentId    *string      `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string      `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string      `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*RunLabel  `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Metrics         []*RunMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	Name            *string      `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Params          []*RunParam  `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId       *string      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunId           *string      `json:"RunId,omitempty" xml:"RunId,omitempty"`
	SourceId        *string      `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string      `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UserId          *string      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string      `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Run) String() string {
	return dara.Prettify(s)
}

func (s Run) GoString() string {
	return s.String()
}

func (s *Run) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Run) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *Run) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Run) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Run) GetLabels() []*RunLabel {
	return s.Labels
}

func (s *Run) GetMetrics() []*RunMetric {
	return s.Metrics
}

func (s *Run) GetName() *string {
	return s.Name
}

func (s *Run) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Run) GetParams() []*RunParam {
	return s.Params
}

func (s *Run) GetRequestId() *string {
	return s.RequestId
}

func (s *Run) GetRunId() *string {
	return s.RunId
}

func (s *Run) GetSourceId() *string {
	return s.SourceId
}

func (s *Run) GetSourceType() *string {
	return s.SourceType
}

func (s *Run) GetUserId() *string {
	return s.UserId
}

func (s *Run) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Run) SetAccessibility(v string) *Run {
	s.Accessibility = &v
	return s
}

func (s *Run) SetExperimentId(v string) *Run {
	s.ExperimentId = &v
	return s
}

func (s *Run) SetGmtCreateTime(v string) *Run {
	s.GmtCreateTime = &v
	return s
}

func (s *Run) SetGmtModifiedTime(v string) *Run {
	s.GmtModifiedTime = &v
	return s
}

func (s *Run) SetLabels(v []*RunLabel) *Run {
	s.Labels = v
	return s
}

func (s *Run) SetMetrics(v []*RunMetric) *Run {
	s.Metrics = v
	return s
}

func (s *Run) SetName(v string) *Run {
	s.Name = &v
	return s
}

func (s *Run) SetOwnerId(v string) *Run {
	s.OwnerId = &v
	return s
}

func (s *Run) SetParams(v []*RunParam) *Run {
	s.Params = v
	return s
}

func (s *Run) SetRequestId(v string) *Run {
	s.RequestId = &v
	return s
}

func (s *Run) SetRunId(v string) *Run {
	s.RunId = &v
	return s
}

func (s *Run) SetSourceId(v string) *Run {
	s.SourceId = &v
	return s
}

func (s *Run) SetSourceType(v string) *Run {
	s.SourceType = &v
	return s
}

func (s *Run) SetUserId(v string) *Run {
	s.UserId = &v
	return s
}

func (s *Run) SetWorkspaceId(v string) *Run {
	s.WorkspaceId = &v
	return s
}

func (s *Run) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
