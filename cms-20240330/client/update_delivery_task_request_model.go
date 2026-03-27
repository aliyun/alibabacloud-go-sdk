// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *UpdateDeliveryTaskRequest
	GetDataSourceId() *string
	SetExternalLabels(v map[string]*string) *UpdateDeliveryTaskRequest
	GetExternalLabels() map[string]*string
	SetLabelFilters(v map[string]*string) *UpdateDeliveryTaskRequest
	GetLabelFilters() map[string]*string
	SetLabelFiltersType(v string) *UpdateDeliveryTaskRequest
	GetLabelFiltersType() *string
	SetResourceGroupId(v string) *UpdateDeliveryTaskRequest
	GetResourceGroupId() *string
	SetSinkList(v []*UpdateDeliveryTaskRequestSinkList) *UpdateDeliveryTaskRequest
	GetSinkList() []*UpdateDeliveryTaskRequestSinkList
	SetStatus(v string) *UpdateDeliveryTaskRequest
	GetStatus() *string
	SetTaskDescription(v string) *UpdateDeliveryTaskRequest
	GetTaskDescription() *string
	SetTaskName(v string) *UpdateDeliveryTaskRequest
	GetTaskName() *string
}

type UpdateDeliveryTaskRequest struct {
	// example:
	//
	// rw-5f2b4sc7es4d66
	DataSourceId   *string            `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	ExternalLabels map[string]*string `json:"externalLabels,omitempty" xml:"externalLabels,omitempty"`
	LabelFilters   map[string]*string `json:"labelFilters,omitempty" xml:"labelFilters,omitempty"`
	// example:
	//
	// Deny
	LabelFiltersType *string `json:"labelFiltersType,omitempty" xml:"labelFiltersType,omitempty"`
	// example:
	//
	// rg-aekzoiafjtr7zyq
	ResourceGroupId *string                              `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	SinkList        []*UpdateDeliveryTaskRequestSinkList `json:"sinkList,omitempty" xml:"sinkList,omitempty" type:"Repeated"`
	// example:
	//
	// Enable
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// updated desc
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// new-task-name
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s UpdateDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryTaskRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDeliveryTaskRequest) GetExternalLabels() map[string]*string {
	return s.ExternalLabels
}

func (s *UpdateDeliveryTaskRequest) GetLabelFilters() map[string]*string {
	return s.LabelFilters
}

func (s *UpdateDeliveryTaskRequest) GetLabelFiltersType() *string {
	return s.LabelFiltersType
}

func (s *UpdateDeliveryTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateDeliveryTaskRequest) GetSinkList() []*UpdateDeliveryTaskRequestSinkList {
	return s.SinkList
}

func (s *UpdateDeliveryTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateDeliveryTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *UpdateDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateDeliveryTaskRequest) SetDataSourceId(v string) *UpdateDeliveryTaskRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetExternalLabels(v map[string]*string) *UpdateDeliveryTaskRequest {
	s.ExternalLabels = v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetLabelFilters(v map[string]*string) *UpdateDeliveryTaskRequest {
	s.LabelFilters = v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetLabelFiltersType(v string) *UpdateDeliveryTaskRequest {
	s.LabelFiltersType = &v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetResourceGroupId(v string) *UpdateDeliveryTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetSinkList(v []*UpdateDeliveryTaskRequestSinkList) *UpdateDeliveryTaskRequest {
	s.SinkList = v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetStatus(v string) *UpdateDeliveryTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetTaskDescription(v string) *UpdateDeliveryTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *UpdateDeliveryTaskRequest) SetTaskName(v string) *UpdateDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateDeliveryTaskRequest) Validate() error {
	if s.SinkList != nil {
		for _, item := range s.SinkList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDeliveryTaskRequestSinkList struct {
	SinkConfigs map[string]*string `json:"sinkConfigs,omitempty" xml:"sinkConfigs,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// Prometheus
	SinkType *string `json:"sinkType,omitempty" xml:"sinkType,omitempty"`
}

func (s UpdateDeliveryTaskRequestSinkList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryTaskRequestSinkList) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryTaskRequestSinkList) GetSinkConfigs() map[string]*string {
	return s.SinkConfigs
}

func (s *UpdateDeliveryTaskRequestSinkList) GetSinkType() *string {
	return s.SinkType
}

func (s *UpdateDeliveryTaskRequestSinkList) SetSinkConfigs(v map[string]*string) *UpdateDeliveryTaskRequestSinkList {
	s.SinkConfigs = v
	return s
}

func (s *UpdateDeliveryTaskRequestSinkList) SetSinkType(v string) *UpdateDeliveryTaskRequestSinkList {
	s.SinkType = &v
	return s
}

func (s *UpdateDeliveryTaskRequestSinkList) Validate() error {
	return dara.Validate(s)
}
