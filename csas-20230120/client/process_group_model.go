// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessGroup interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ProcessGroup
	GetDescription() *string
	SetGmtCreate(v string) *ProcessGroup
	GetGmtCreate() *string
	SetGmtModified(v string) *ProcessGroup
	GetGmtModified() *string
	SetName(v string) *ProcessGroup
	GetName() *string
	SetProcessGroupId(v string) *ProcessGroup
	GetProcessGroupId() *string
	SetProcesses(v []*ProcessItem) *ProcessGroup
	GetProcesses() []*ProcessItem
}

type ProcessGroup struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessGroupId *string `json:"ProcessGroupId,omitempty" xml:"ProcessGroupId,omitempty"`
	// if can be null:
	// false
	Processes []*ProcessItem `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
}

func (s ProcessGroup) String() string {
	return dara.Prettify(s)
}

func (s ProcessGroup) GoString() string {
	return s.String()
}

func (s *ProcessGroup) GetDescription() *string {
	return s.Description
}

func (s *ProcessGroup) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ProcessGroup) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ProcessGroup) GetName() *string {
	return s.Name
}

func (s *ProcessGroup) GetProcessGroupId() *string {
	return s.ProcessGroupId
}

func (s *ProcessGroup) GetProcesses() []*ProcessItem {
	return s.Processes
}

func (s *ProcessGroup) SetDescription(v string) *ProcessGroup {
	s.Description = &v
	return s
}

func (s *ProcessGroup) SetGmtCreate(v string) *ProcessGroup {
	s.GmtCreate = &v
	return s
}

func (s *ProcessGroup) SetGmtModified(v string) *ProcessGroup {
	s.GmtModified = &v
	return s
}

func (s *ProcessGroup) SetName(v string) *ProcessGroup {
	s.Name = &v
	return s
}

func (s *ProcessGroup) SetProcessGroupId(v string) *ProcessGroup {
	s.ProcessGroupId = &v
	return s
}

func (s *ProcessGroup) SetProcesses(v []*ProcessItem) *ProcessGroup {
	s.Processes = v
	return s
}

func (s *ProcessGroup) Validate() error {
	return dara.Validate(s)
}
