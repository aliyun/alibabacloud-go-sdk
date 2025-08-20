// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructRefreshJobModel interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *OpenStructRefreshJobModel
	GetEndTime() *string
	SetName(v string) *OpenStructRefreshJobModel
	GetName() *string
	SetProcessid(v string) *OpenStructRefreshJobModel
	GetProcessid() *string
	SetRefreshInterval(v string) *OpenStructRefreshJobModel
	GetRefreshInterval() *string
	SetRefreshModel(v string) *OpenStructRefreshJobModel
	GetRefreshModel() *string
	SetResourceGroup(v string) *OpenStructRefreshJobModel
	GetResourceGroup() *string
	SetScheduledStartTime(v string) *OpenStructRefreshJobModel
	GetScheduledStartTime() *string
	SetSchemaName(v string) *OpenStructRefreshJobModel
	GetSchemaName() *string
	SetStartTime(v string) *OpenStructRefreshJobModel
	GetStartTime() *string
	SetStatus(v string) *OpenStructRefreshJobModel
	GetStatus() *string
}

type OpenStructRefreshJobModel struct {
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Processid          *string `json:"Processid,omitempty" xml:"Processid,omitempty"`
	RefreshInterval    *string `json:"RefreshInterval,omitempty" xml:"RefreshInterval,omitempty"`
	RefreshModel       *string `json:"RefreshModel,omitempty" xml:"RefreshModel,omitempty"`
	ResourceGroup      *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	SchemaName         *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OpenStructRefreshJobModel) String() string {
	return dara.Prettify(s)
}

func (s OpenStructRefreshJobModel) GoString() string {
	return s.String()
}

func (s *OpenStructRefreshJobModel) GetEndTime() *string {
	return s.EndTime
}

func (s *OpenStructRefreshJobModel) GetName() *string {
	return s.Name
}

func (s *OpenStructRefreshJobModel) GetProcessid() *string {
	return s.Processid
}

func (s *OpenStructRefreshJobModel) GetRefreshInterval() *string {
	return s.RefreshInterval
}

func (s *OpenStructRefreshJobModel) GetRefreshModel() *string {
	return s.RefreshModel
}

func (s *OpenStructRefreshJobModel) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *OpenStructRefreshJobModel) GetScheduledStartTime() *string {
	return s.ScheduledStartTime
}

func (s *OpenStructRefreshJobModel) GetSchemaName() *string {
	return s.SchemaName
}

func (s *OpenStructRefreshJobModel) GetStartTime() *string {
	return s.StartTime
}

func (s *OpenStructRefreshJobModel) GetStatus() *string {
	return s.Status
}

func (s *OpenStructRefreshJobModel) SetEndTime(v string) *OpenStructRefreshJobModel {
	s.EndTime = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetName(v string) *OpenStructRefreshJobModel {
	s.Name = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetProcessid(v string) *OpenStructRefreshJobModel {
	s.Processid = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetRefreshInterval(v string) *OpenStructRefreshJobModel {
	s.RefreshInterval = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetRefreshModel(v string) *OpenStructRefreshJobModel {
	s.RefreshModel = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetResourceGroup(v string) *OpenStructRefreshJobModel {
	s.ResourceGroup = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetScheduledStartTime(v string) *OpenStructRefreshJobModel {
	s.ScheduledStartTime = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetSchemaName(v string) *OpenStructRefreshJobModel {
	s.SchemaName = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetStartTime(v string) *OpenStructRefreshJobModel {
	s.StartTime = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetStatus(v string) *OpenStructRefreshJobModel {
	s.Status = &v
	return s
}

func (s *OpenStructRefreshJobModel) Validate() error {
	return dara.Validate(s)
}
