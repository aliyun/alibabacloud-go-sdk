// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrial interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Trial
	GetAccessibility() *string
	SetExperimentId(v string) *Trial
	GetExperimentId() *string
	SetGmtCreateTime(v string) *Trial
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Trial
	GetGmtModifiedTime() *string
	SetLabels(v []map[string]interface{}) *Trial
	GetLabels() []map[string]interface{}
	SetName(v string) *Trial
	GetName() *string
	SetOwnerId(v string) *Trial
	GetOwnerId() *string
	SetSourceId(v string) *Trial
	GetSourceId() *string
	SetSourceType(v string) *Trial
	GetSourceType() *string
	SetTrialId(v string) *Trial
	GetTrialId() *string
	SetUserId(v string) *Trial
	GetUserId() *string
	SetWorkspaceId(v string) *Trial
	GetWorkspaceId() *string
}

type Trial struct {
	Accessibility   *string                  `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ExperimentId    *string                  `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string                  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceId        *string                  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrialId         *string                  `json:"TrialId,omitempty" xml:"TrialId,omitempty"`
	UserId          *string                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Trial) String() string {
	return dara.Prettify(s)
}

func (s Trial) GoString() string {
	return s.String()
}

func (s *Trial) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Trial) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *Trial) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Trial) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Trial) GetLabels() []map[string]interface{} {
	return s.Labels
}

func (s *Trial) GetName() *string {
	return s.Name
}

func (s *Trial) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Trial) GetSourceId() *string {
	return s.SourceId
}

func (s *Trial) GetSourceType() *string {
	return s.SourceType
}

func (s *Trial) GetTrialId() *string {
	return s.TrialId
}

func (s *Trial) GetUserId() *string {
	return s.UserId
}

func (s *Trial) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Trial) SetAccessibility(v string) *Trial {
	s.Accessibility = &v
	return s
}

func (s *Trial) SetExperimentId(v string) *Trial {
	s.ExperimentId = &v
	return s
}

func (s *Trial) SetGmtCreateTime(v string) *Trial {
	s.GmtCreateTime = &v
	return s
}

func (s *Trial) SetGmtModifiedTime(v string) *Trial {
	s.GmtModifiedTime = &v
	return s
}

func (s *Trial) SetLabels(v []map[string]interface{}) *Trial {
	s.Labels = v
	return s
}

func (s *Trial) SetName(v string) *Trial {
	s.Name = &v
	return s
}

func (s *Trial) SetOwnerId(v string) *Trial {
	s.OwnerId = &v
	return s
}

func (s *Trial) SetSourceId(v string) *Trial {
	s.SourceId = &v
	return s
}

func (s *Trial) SetSourceType(v string) *Trial {
	s.SourceType = &v
	return s
}

func (s *Trial) SetTrialId(v string) *Trial {
	s.TrialId = &v
	return s
}

func (s *Trial) SetUserId(v string) *Trial {
	s.UserId = &v
	return s
}

func (s *Trial) SetWorkspaceId(v string) *Trial {
	s.WorkspaceId = &v
	return s
}

func (s *Trial) Validate() error {
	return dara.Validate(s)
}
