// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeOncallSchedule interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v map[string]interface{}) *MergeOncallSchedule
	GetDetail() map[string]interface{}
	SetGmtCreate(v string) *MergeOncallSchedule
	GetGmtCreate() *string
	SetGmtModified(v string) *MergeOncallSchedule
	GetGmtModified() *string
	SetIdentifier(v string) *MergeOncallSchedule
	GetIdentifier() *string
	SetName(v string) *MergeOncallSchedule
	GetName() *string
	SetSource(v string) *MergeOncallSchedule
	GetSource() *string
	SetWorkspace(v string) *MergeOncallSchedule
	GetWorkspace() *string
}

type MergeOncallSchedule struct {
	Detail      map[string]interface{} `json:"detail,omitempty" xml:"detail,omitempty"`
	GmtCreate   *string                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier  *string                `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	Source      *string                `json:"source,omitempty" xml:"source,omitempty"`
	Workspace   *string                `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s MergeOncallSchedule) String() string {
	return dara.Prettify(s)
}

func (s MergeOncallSchedule) GoString() string {
	return s.String()
}

func (s *MergeOncallSchedule) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *MergeOncallSchedule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MergeOncallSchedule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MergeOncallSchedule) GetIdentifier() *string {
	return s.Identifier
}

func (s *MergeOncallSchedule) GetName() *string {
	return s.Name
}

func (s *MergeOncallSchedule) GetSource() *string {
	return s.Source
}

func (s *MergeOncallSchedule) GetWorkspace() *string {
	return s.Workspace
}

func (s *MergeOncallSchedule) SetDetail(v map[string]interface{}) *MergeOncallSchedule {
	s.Detail = v
	return s
}

func (s *MergeOncallSchedule) SetGmtCreate(v string) *MergeOncallSchedule {
	s.GmtCreate = &v
	return s
}

func (s *MergeOncallSchedule) SetGmtModified(v string) *MergeOncallSchedule {
	s.GmtModified = &v
	return s
}

func (s *MergeOncallSchedule) SetIdentifier(v string) *MergeOncallSchedule {
	s.Identifier = &v
	return s
}

func (s *MergeOncallSchedule) SetName(v string) *MergeOncallSchedule {
	s.Name = &v
	return s
}

func (s *MergeOncallSchedule) SetSource(v string) *MergeOncallSchedule {
	s.Source = &v
	return s
}

func (s *MergeOncallSchedule) SetWorkspace(v string) *MergeOncallSchedule {
	s.Workspace = &v
	return s
}

func (s *MergeOncallSchedule) Validate() error {
	return dara.Validate(s)
}
