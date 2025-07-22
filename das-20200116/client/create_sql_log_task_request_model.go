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
	// The end of the time range to query. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter conditions.
	Filters []*CreateSqlLogTaskRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The ID of the database instance.
	//
	// example:
	//
	// pc-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// test01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node ID.
	//
	// >  This parameter is available only for instances that run in a cluster architecture. You can specify this parameter to query the offline tasks of a specific node. By default, if this parameter is not specified, the information about the offline tasks of the primary node is returned.
	//
	// example:
	//
	// pi-uf6k5f6g3912i0dqz
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The role of the node of the PolarDB-X 2.0 database instance. Valid values:
	//
	// 	- **polarx_cn**: compute node
	//
	// 	- **polarx_dn**: data node
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The beginning of the time range to query. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **Export**
	//
	// 	- **Query**
	//
	// 	- **Insight**
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
	return dara.Validate(s)
}

type CreateSqlLogTaskRequestFilters struct {
	// The name of the filter parameter.
	//
	// >  For more information about the supported filter parameters and their valid values, see the following **supplement about the Key parameter**.
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
