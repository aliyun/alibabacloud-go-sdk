// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlLogTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *CreateSqlLogTaskRequest
	GetEndTime() *int64
	SetFilters(v []*CreateSqlLogTaskRequestFilters) *CreateSqlLogTaskRequest
	GetFilters() []*CreateSqlLogTaskRequestFilters
	SetInstanceId(v string) *CreateSqlLogTaskRequest
	GetInstanceId() *string
	SetName(v string) *CreateSqlLogTaskRequest
	GetName() *string
	SetNodeId(v string) *CreateSqlLogTaskRequest
	GetNodeId() *string
	SetRole(v string) *CreateSqlLogTaskRequest
	GetRole() *string
	SetStartTime(v int64) *CreateSqlLogTaskRequest
	GetStartTime() *int64
	SetType(v string) *CreateSqlLogTaskRequest
	GetType() *string
}

type CreateSqlLogTaskRequest struct {
	// The end time of the task. Specify the value as a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of filter conditions.
	Filters []*CreateSqlLogTaskRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The database instance ID.
	//
	// example:
	//
	// pc-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task name.
	//
	// example:
	//
	// SQL audit export 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node ID.
	//
	// >This parameter is applicable only to cluster instances. You can specify this parameter to query the batch task of a specific node. If you do not specify this parameter, the batch task of the primary node is returned by default.
	//
	// example:
	//
	// pi-uf6k5f6g3912i****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node information of the PolarDB-X 2.0 database instance.
	//
	// - **polarx_cn**: compute node.
	//
	// - **polarx_dn**: data node.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The start time of the task. Specify the value as a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task type.
	//
	// - **Export**: export task.
	//
	// > For the filter parameters and values supported by **Export**, see **Request parameters description**.
	//
	// - **Query**: query task.
	//
	// example:
	//
	// Export
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateSqlLogTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlLogTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlLogTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateSqlLogTaskRequest) GetFilters() []*CreateSqlLogTaskRequestFilters {
	return s.Filters
}

func (s *CreateSqlLogTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSqlLogTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateSqlLogTaskRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateSqlLogTaskRequest) GetRole() *string {
	return s.Role
}

func (s *CreateSqlLogTaskRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateSqlLogTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateSqlLogTaskRequest) SetEndTime(v int64) *CreateSqlLogTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSqlLogTaskRequest) SetFilters(v []*CreateSqlLogTaskRequestFilters) *CreateSqlLogTaskRequest {
	s.Filters = v
	return s
}

func (s *CreateSqlLogTaskRequest) SetInstanceId(v string) *CreateSqlLogTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSqlLogTaskRequest) SetName(v string) *CreateSqlLogTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateSqlLogTaskRequest) SetNodeId(v string) *CreateSqlLogTaskRequest {
	s.NodeId = &v
	return s
}

func (s *CreateSqlLogTaskRequest) SetRole(v string) *CreateSqlLogTaskRequest {
	s.Role = &v
	return s
}

func (s *CreateSqlLogTaskRequest) SetStartTime(v int64) *CreateSqlLogTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSqlLogTaskRequest) SetType(v string) *CreateSqlLogTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateSqlLogTaskRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSqlLogTaskRequestFilters struct {
	// The name of the filter parameter.
	//
	// > For the supported filter parameters and values, see **Request parameters description**.
	//
	// example:
	//
	// KeyWords
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter parameter.
	//
	// example:
	//
	// select
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSqlLogTaskRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlLogTaskRequestFilters) GoString() string {
	return s.String()
}

func (s *CreateSqlLogTaskRequestFilters) GetKey() *string {
	return s.Key
}

func (s *CreateSqlLogTaskRequestFilters) GetValue() *string {
	return s.Value
}

func (s *CreateSqlLogTaskRequestFilters) SetKey(v string) *CreateSqlLogTaskRequestFilters {
	s.Key = &v
	return s
}

func (s *CreateSqlLogTaskRequestFilters) SetValue(v string) *CreateSqlLogTaskRequestFilters {
	s.Value = &v
	return s
}

func (s *CreateSqlLogTaskRequestFilters) Validate() error {
	return dara.Validate(s)
}
