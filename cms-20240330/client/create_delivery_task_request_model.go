// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *CreateDeliveryTaskRequest
	GetDataSourceId() *string
	SetExternalLabels(v map[string]*string) *CreateDeliveryTaskRequest
	GetExternalLabels() map[string]*string
	SetLabelFilters(v map[string]*string) *CreateDeliveryTaskRequest
	GetLabelFilters() map[string]*string
	SetLabelFiltersType(v string) *CreateDeliveryTaskRequest
	GetLabelFiltersType() *string
	SetResourceGroupId(v string) *CreateDeliveryTaskRequest
	GetResourceGroupId() *string
	SetSinkList(v []*CreateDeliveryTaskRequestSinkList) *CreateDeliveryTaskRequest
	GetSinkList() []*CreateDeliveryTaskRequestSinkList
	SetTags(v []*CreateDeliveryTaskRequestTags) *CreateDeliveryTaskRequest
	GetTags() []*CreateDeliveryTaskRequestTags
	SetTaskDescription(v string) *CreateDeliveryTaskRequest
	GetTaskDescription() *string
	SetTaskName(v string) *CreateDeliveryTaskRequest
	GetTaskName() *string
}

type CreateDeliveryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rw-xxxxxx
	DataSourceId   *string            `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	ExternalLabels map[string]*string `json:"externalLabels,omitempty" xml:"externalLabels,omitempty"`
	LabelFilters   map[string]*string `json:"labelFilters,omitempty" xml:"labelFilters,omitempty"`
	// example:
	//
	// Allow
	LabelFiltersType *string `json:"labelFiltersType,omitempty" xml:"labelFiltersType,omitempty"`
	// example:
	//
	// rg-ae******ey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// This parameter is required.
	SinkList []*CreateDeliveryTaskRequestSinkList `json:"sinkList,omitempty" xml:"sinkList,omitempty" type:"Repeated"`
	Tags     []*CreateDeliveryTaskRequestTags     `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// my delivery task
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-task
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryTaskRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateDeliveryTaskRequest) GetExternalLabels() map[string]*string {
	return s.ExternalLabels
}

func (s *CreateDeliveryTaskRequest) GetLabelFilters() map[string]*string {
	return s.LabelFilters
}

func (s *CreateDeliveryTaskRequest) GetLabelFiltersType() *string {
	return s.LabelFiltersType
}

func (s *CreateDeliveryTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDeliveryTaskRequest) GetSinkList() []*CreateDeliveryTaskRequestSinkList {
	return s.SinkList
}

func (s *CreateDeliveryTaskRequest) GetTags() []*CreateDeliveryTaskRequestTags {
	return s.Tags
}

func (s *CreateDeliveryTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *CreateDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateDeliveryTaskRequest) SetDataSourceId(v string) *CreateDeliveryTaskRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateDeliveryTaskRequest) SetExternalLabels(v map[string]*string) *CreateDeliveryTaskRequest {
	s.ExternalLabels = v
	return s
}

func (s *CreateDeliveryTaskRequest) SetLabelFilters(v map[string]*string) *CreateDeliveryTaskRequest {
	s.LabelFilters = v
	return s
}

func (s *CreateDeliveryTaskRequest) SetLabelFiltersType(v string) *CreateDeliveryTaskRequest {
	s.LabelFiltersType = &v
	return s
}

func (s *CreateDeliveryTaskRequest) SetResourceGroupId(v string) *CreateDeliveryTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDeliveryTaskRequest) SetSinkList(v []*CreateDeliveryTaskRequestSinkList) *CreateDeliveryTaskRequest {
	s.SinkList = v
	return s
}

func (s *CreateDeliveryTaskRequest) SetTags(v []*CreateDeliveryTaskRequestTags) *CreateDeliveryTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateDeliveryTaskRequest) SetTaskDescription(v string) *CreateDeliveryTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *CreateDeliveryTaskRequest) SetTaskName(v string) *CreateDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateDeliveryTaskRequest) Validate() error {
	if s.SinkList != nil {
		for _, item := range s.SinkList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDeliveryTaskRequestSinkList struct {
	SinkConfigs map[string]*string `json:"sinkConfigs,omitempty" xml:"sinkConfigs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Prometheus
	SinkType *string `json:"sinkType,omitempty" xml:"sinkType,omitempty"`
}

func (s CreateDeliveryTaskRequestSinkList) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryTaskRequestSinkList) GoString() string {
	return s.String()
}

func (s *CreateDeliveryTaskRequestSinkList) GetSinkConfigs() map[string]*string {
	return s.SinkConfigs
}

func (s *CreateDeliveryTaskRequestSinkList) GetSinkType() *string {
	return s.SinkType
}

func (s *CreateDeliveryTaskRequestSinkList) SetSinkConfigs(v map[string]*string) *CreateDeliveryTaskRequestSinkList {
	s.SinkConfigs = v
	return s
}

func (s *CreateDeliveryTaskRequestSinkList) SetSinkType(v string) *CreateDeliveryTaskRequestSinkList {
	s.SinkType = &v
	return s
}

func (s *CreateDeliveryTaskRequestSinkList) Validate() error {
	return dara.Validate(s)
}

type CreateDeliveryTaskRequestTags struct {
	// example:
	//
	// _cms_workspace
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateDeliveryTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryTaskRequestTags) GoString() string {
	return s.String()
}

func (s *CreateDeliveryTaskRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateDeliveryTaskRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateDeliveryTaskRequestTags) SetKey(v string) *CreateDeliveryTaskRequestTags {
	s.Key = &v
	return s
}

func (s *CreateDeliveryTaskRequestTags) SetValue(v string) *CreateDeliveryTaskRequestTags {
	s.Value = &v
	return s
}

func (s *CreateDeliveryTaskRequestTags) Validate() error {
	return dara.Validate(s)
}
