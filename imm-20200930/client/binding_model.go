// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBinding interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *Binding
	GetCreateTime() *string
	SetDatasetName(v string) *Binding
	GetDatasetName() *string
	SetPhase(v string) *Binding
	GetPhase() *string
	SetProjectName(v string) *Binding
	GetProjectName() *string
	SetReason(v string) *Binding
	GetReason() *string
	SetState(v string) *Binding
	GetState() *string
	SetURI(v string) *Binding
	GetURI() *string
	SetUpdateTime(v string) *Binding
	GetUpdateTime() *string
}

type Binding struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Phase       *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Binding) String() string {
	return dara.Prettify(s)
}

func (s Binding) GoString() string {
	return s.String()
}

func (s *Binding) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Binding) GetDatasetName() *string {
	return s.DatasetName
}

func (s *Binding) GetPhase() *string {
	return s.Phase
}

func (s *Binding) GetProjectName() *string {
	return s.ProjectName
}

func (s *Binding) GetReason() *string {
	return s.Reason
}

func (s *Binding) GetState() *string {
	return s.State
}

func (s *Binding) GetURI() *string {
	return s.URI
}

func (s *Binding) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Binding) SetCreateTime(v string) *Binding {
	s.CreateTime = &v
	return s
}

func (s *Binding) SetDatasetName(v string) *Binding {
	s.DatasetName = &v
	return s
}

func (s *Binding) SetPhase(v string) *Binding {
	s.Phase = &v
	return s
}

func (s *Binding) SetProjectName(v string) *Binding {
	s.ProjectName = &v
	return s
}

func (s *Binding) SetReason(v string) *Binding {
	s.Reason = &v
	return s
}

func (s *Binding) SetState(v string) *Binding {
	s.State = &v
	return s
}

func (s *Binding) SetURI(v string) *Binding {
	s.URI = &v
	return s
}

func (s *Binding) SetUpdateTime(v string) *Binding {
	s.UpdateTime = &v
	return s
}

func (s *Binding) Validate() error {
	return dara.Validate(s)
}
