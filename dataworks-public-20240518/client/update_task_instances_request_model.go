// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateTaskInstancesRequest
	GetComment() *string
	SetTaskInstances(v []*UpdateTaskInstancesRequestTaskInstances) *UpdateTaskInstancesRequest
	GetTaskInstances() []*UpdateTaskInstancesRequestTaskInstances
}

type UpdateTaskInstancesRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The instances.
	TaskInstances []*UpdateTaskInstancesRequestTaskInstances `json:"TaskInstances,omitempty" xml:"TaskInstances,omitempty" type:"Repeated"`
}

func (s UpdateTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateTaskInstancesRequest) GetTaskInstances() []*UpdateTaskInstancesRequestTaskInstances {
	return s.TaskInstances
}

func (s *UpdateTaskInstancesRequest) SetComment(v string) *UpdateTaskInstancesRequest {
	s.Comment = &v
	return s
}

func (s *UpdateTaskInstancesRequest) SetTaskInstances(v []*UpdateTaskInstancesRequestTaskInstances) *UpdateTaskInstancesRequest {
	s.TaskInstances = v
	return s
}

func (s *UpdateTaskInstancesRequest) Validate() error {
	if s.TaskInstances != nil {
		for _, item := range s.TaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTaskInstancesRequestTaskInstances struct {
	// The information about the associated data source.
	DataSource *UpdateTaskInstancesRequestTaskInstancesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The priority of the instance. Valid values: 1, 3, 5, 7, and 8.
	//
	// A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource group information. Set this parameter to the ID of a resource group for scheduling.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	RuntimeResource *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
}

func (s UpdateTaskInstancesRequestTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskInstancesRequestTaskInstances) GoString() string {
	return s.String()
}

func (s *UpdateTaskInstancesRequestTaskInstances) GetDataSource() *UpdateTaskInstancesRequestTaskInstancesDataSource {
	return s.DataSource
}

func (s *UpdateTaskInstancesRequestTaskInstances) GetId() *int64 {
	return s.Id
}

func (s *UpdateTaskInstancesRequestTaskInstances) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateTaskInstancesRequestTaskInstances) GetRuntimeResource() *string {
	return s.RuntimeResource
}

func (s *UpdateTaskInstancesRequestTaskInstances) SetDataSource(v *UpdateTaskInstancesRequestTaskInstancesDataSource) *UpdateTaskInstancesRequestTaskInstances {
	s.DataSource = v
	return s
}

func (s *UpdateTaskInstancesRequestTaskInstances) SetId(v int64) *UpdateTaskInstancesRequestTaskInstances {
	s.Id = &v
	return s
}

func (s *UpdateTaskInstancesRequestTaskInstances) SetPriority(v int32) *UpdateTaskInstancesRequestTaskInstances {
	s.Priority = &v
	return s
}

func (s *UpdateTaskInstancesRequestTaskInstances) SetRuntimeResource(v string) *UpdateTaskInstancesRequestTaskInstances {
	s.RuntimeResource = &v
	return s
}

func (s *UpdateTaskInstancesRequestTaskInstances) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTaskInstancesRequestTaskInstancesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTaskInstancesRequestTaskInstancesDataSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskInstancesRequestTaskInstancesDataSource) GoString() string {
	return s.String()
}

func (s *UpdateTaskInstancesRequestTaskInstancesDataSource) GetName() *string {
	return s.Name
}

func (s *UpdateTaskInstancesRequestTaskInstancesDataSource) SetName(v string) *UpdateTaskInstancesRequestTaskInstancesDataSource {
	s.Name = &v
	return s
}

func (s *UpdateTaskInstancesRequestTaskInstancesDataSource) Validate() error {
	return dara.Validate(s)
}
