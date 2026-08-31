// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeSupplementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateNodeSupplementRequestCreateCommand) *CreateNodeSupplementRequest
	GetCreateCommand() *CreateNodeSupplementRequestCreateCommand
	SetEnv(v string) *CreateNodeSupplementRequest
	GetEnv() *string
	SetOpTenantId(v int64) *CreateNodeSupplementRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateNodeSupplementRequest
	GetOpUserId() *string
}

type CreateNodeSupplementRequest struct {
	// The data backfill request.
	//
	// This parameter is required.
	CreateCommand *CreateNodeSupplementRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The environment identifier. Valid values:
	//
	// - DEV: Development environment.
	//
	// - PROD (default): Production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateNodeSupplementRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequest) GetCreateCommand() *CreateNodeSupplementRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateNodeSupplementRequest) GetEnv() *string {
	return s.Env
}

func (s *CreateNodeSupplementRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateNodeSupplementRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateNodeSupplementRequest) SetCreateCommand(v *CreateNodeSupplementRequestCreateCommand) *CreateNodeSupplementRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateNodeSupplementRequest) SetEnv(v string) *CreateNodeSupplementRequest {
	s.Env = &v
	return s
}

func (s *CreateNodeSupplementRequest) SetOpTenantId(v int64) *CreateNodeSupplementRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateNodeSupplementRequest) SetOpUserId(v string) *CreateNodeSupplementRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateNodeSupplementRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNodeSupplementRequestCreateCommand struct {
	// Specifies whether to include all downstream nodes in batch mode. Default value: false.
	//
	// example:
	//
	// false
	ContainAllDownStream *bool `json:"ContainAllDownStream,omitempty" xml:"ContainAllDownStream,omitempty"`
	// The IDs of downstream nodes to run. This parameter takes effect only when ContainAllDownStream is set to false.
	DownStreamNodeIdList []*CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList `json:"DownStreamNodeIdList,omitempty" xml:"DownStreamNodeIdList,omitempty" type:"Repeated"`
	// The end business date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-21
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// The list of filters. This parameter is used together with ContainAllDownStream to include or exclude nodes based on criteria such as project or node. Default value: empty.
	FilterList []*CreateNodeSupplementRequestCreateCommandFilterList `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	// The runtime global parameters.
	GlobalParamList []*CreateNodeSupplementRequestCreateCommandGlobalParamList `json:"GlobalParamList,omitempty" xml:"GlobalParamList,omitempty" type:"Repeated"`
	// The latest trigger time in the HH:MM format. This parameter is applicable only to hourly nodes.
	//
	// example:
	//
	// 20:59
	MaxDueTime *string `json:"MaxDueTime,omitempty" xml:"MaxDueTime,omitempty"`
	// The earliest trigger time in the HH:MM format. This parameter is applicable only to hourly nodes.
	//
	// example:
	//
	// 00:00
	MinDueTime *string `json:"MinDueTime,omitempty" xml:"MinDueTime,omitempty"`
	// The name of the data backfill. If this parameter is not specified, the system automatically generates a name.
	//
	// example:
	//
	// TestBackfill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The root nodes for data backfill. Multiple root nodes are supported.
	//
	// This parameter is required.
	NodeIdList []*CreateNodeSupplementRequestCreateCommandNodeIdList `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// The runtime custom parameters configured by node.
	NodeParamsList []*CreateNodeSupplementRequestCreateCommandNodeParamsList `json:"NodeParamsList,omitempty" xml:"NodeParamsList,omitempty" type:"Repeated"`
	// The concurrency. Default value: 1.
	//
	// example:
	//
	// 1
	Parallelism *int32 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 101121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Corresponds to the "Task Run Time" option in the data backfill dialog box. If this parameter is set to true, the scheduled run time of instances is ignored and all instances run immediately. If this parameter is set to false, instances wait for their scheduled run time before running. Default value: true.
	RunImmediately *bool `json:"RunImmediately,omitempty" xml:"RunImmediately,omitempty"`
	// The start business date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-21
	StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommand) GetContainAllDownStream() *bool {
	return s.ContainAllDownStream
}

func (s *CreateNodeSupplementRequestCreateCommand) GetDownStreamNodeIdList() []*CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList {
	return s.DownStreamNodeIdList
}

func (s *CreateNodeSupplementRequestCreateCommand) GetEndBizDate() *string {
	return s.EndBizDate
}

func (s *CreateNodeSupplementRequestCreateCommand) GetFilterList() []*CreateNodeSupplementRequestCreateCommandFilterList {
	return s.FilterList
}

func (s *CreateNodeSupplementRequestCreateCommand) GetGlobalParamList() []*CreateNodeSupplementRequestCreateCommandGlobalParamList {
	return s.GlobalParamList
}

func (s *CreateNodeSupplementRequestCreateCommand) GetMaxDueTime() *string {
	return s.MaxDueTime
}

func (s *CreateNodeSupplementRequestCreateCommand) GetMinDueTime() *string {
	return s.MinDueTime
}

func (s *CreateNodeSupplementRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateNodeSupplementRequestCreateCommand) GetNodeIdList() []*CreateNodeSupplementRequestCreateCommandNodeIdList {
	return s.NodeIdList
}

func (s *CreateNodeSupplementRequestCreateCommand) GetNodeParamsList() []*CreateNodeSupplementRequestCreateCommandNodeParamsList {
	return s.NodeParamsList
}

func (s *CreateNodeSupplementRequestCreateCommand) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *CreateNodeSupplementRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateNodeSupplementRequestCreateCommand) GetRunImmediately() *bool {
	return s.RunImmediately
}

func (s *CreateNodeSupplementRequestCreateCommand) GetStartBizDate() *string {
	return s.StartBizDate
}

func (s *CreateNodeSupplementRequestCreateCommand) SetContainAllDownStream(v bool) *CreateNodeSupplementRequestCreateCommand {
	s.ContainAllDownStream = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetDownStreamNodeIdList(v []*CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) *CreateNodeSupplementRequestCreateCommand {
	s.DownStreamNodeIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetEndBizDate(v string) *CreateNodeSupplementRequestCreateCommand {
	s.EndBizDate = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetFilterList(v []*CreateNodeSupplementRequestCreateCommandFilterList) *CreateNodeSupplementRequestCreateCommand {
	s.FilterList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetGlobalParamList(v []*CreateNodeSupplementRequestCreateCommandGlobalParamList) *CreateNodeSupplementRequestCreateCommand {
	s.GlobalParamList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetMaxDueTime(v string) *CreateNodeSupplementRequestCreateCommand {
	s.MaxDueTime = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetMinDueTime(v string) *CreateNodeSupplementRequestCreateCommand {
	s.MinDueTime = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetName(v string) *CreateNodeSupplementRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetNodeIdList(v []*CreateNodeSupplementRequestCreateCommandNodeIdList) *CreateNodeSupplementRequestCreateCommand {
	s.NodeIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetNodeParamsList(v []*CreateNodeSupplementRequestCreateCommandNodeParamsList) *CreateNodeSupplementRequestCreateCommand {
	s.NodeParamsList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetParallelism(v int32) *CreateNodeSupplementRequestCreateCommand {
	s.Parallelism = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetProjectId(v int64) *CreateNodeSupplementRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetRunImmediately(v bool) *CreateNodeSupplementRequestCreateCommand {
	s.RunImmediately = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetStartBizDate(v string) *CreateNodeSupplementRequestCreateCommand {
	s.StartBizDate = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) Validate() error {
	if s.DownStreamNodeIdList != nil {
		for _, item := range s.DownStreamNodeIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterList != nil {
		for _, item := range s.FilterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GlobalParamList != nil {
		for _, item := range s.GlobalParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeIdList != nil {
		for _, item := range s.NodeIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeParamsList != nil {
		for _, item := range s.NodeParamsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList struct {
	// The field ID.
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// 2323232
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) GetFieldIdList() []*string {
	return s.FieldIdList
}

func (s *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) GetId() *string {
	return s.Id
}

func (s *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) SetFieldIdList(v []*string) *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList {
	s.FieldIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) SetId(v string) *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList {
	s.Id = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) Validate() error {
	return dara.Validate(s)
}

type CreateNodeSupplementRequestCreateCommandFilterList struct {
	// Specifies whether to exclude the matched items. Default value: false.
	//
	// example:
	//
	// false
	Exclude *bool `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// The filter key. Valid values:
	//
	// - PROJECT: project
	//
	// - NODE_OUTPUT_NAME: node output name
	//
	// - NODE_NAME: node name
	//
	// - NODE_ID: node ID
	//
	// - TARGETS: specified endpoints
	//
	// - SOURCES: specified start points
	//
	// example:
	//
	// NODE_OUTPUT_NAME
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of filter values.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s CreateNodeSupplementRequestCreateCommandFilterList) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandFilterList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) GetExclude() *bool {
	return s.Exclude
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) GetKey() *string {
	return s.Key
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) GetValueList() []*string {
	return s.ValueList
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) SetExclude(v bool) *CreateNodeSupplementRequestCreateCommandFilterList {
	s.Exclude = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) SetKey(v string) *CreateNodeSupplementRequestCreateCommandFilterList {
	s.Key = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) SetValueList(v []*string) *CreateNodeSupplementRequestCreateCommandFilterList {
	s.ValueList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) Validate() error {
	return dara.Validate(s)
}

type CreateNodeSupplementRequestCreateCommandGlobalParamList struct {
	// The parameter.
	//
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandGlobalParamList) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandGlobalParamList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandGlobalParamList) GetKey() *string {
	return s.Key
}

func (s *CreateNodeSupplementRequestCreateCommandGlobalParamList) GetValue() *string {
	return s.Value
}

func (s *CreateNodeSupplementRequestCreateCommandGlobalParamList) SetKey(v string) *CreateNodeSupplementRequestCreateCommandGlobalParamList {
	s.Key = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandGlobalParamList) SetValue(v string) *CreateNodeSupplementRequestCreateCommandGlobalParamList {
	s.Value = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandGlobalParamList) Validate() error {
	return dara.Validate(s)
}

type CreateNodeSupplementRequestCreateCommandNodeIdList struct {
	// The list of field IDs. This parameter is applicable when the node ID is a logical table node ID. If this parameter is not specified, all fields in the table are used by default.
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n_1232324
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandNodeIdList) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandNodeIdList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandNodeIdList) GetFieldIdList() []*string {
	return s.FieldIdList
}

func (s *CreateNodeSupplementRequestCreateCommandNodeIdList) GetId() *string {
	return s.Id
}

func (s *CreateNodeSupplementRequestCreateCommandNodeIdList) SetFieldIdList(v []*string) *CreateNodeSupplementRequestCreateCommandNodeIdList {
	s.FieldIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeIdList) SetId(v string) *CreateNodeSupplementRequestCreateCommandNodeIdList {
	s.Id = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeIdList) Validate() error {
	return dara.Validate(s)
}

type CreateNodeSupplementRequestCreateCommandNodeParamsList struct {
	// The node ID.
	//
	// example:
	//
	// n_23324
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The parameter list.
	ParamList []*CreateNodeSupplementRequestCreateCommandNodeParamsListParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsList) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsList) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsList) GetParamList() []*CreateNodeSupplementRequestCreateCommandNodeParamsListParamList {
	return s.ParamList
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsList) SetNodeId(v string) *CreateNodeSupplementRequestCreateCommandNodeParamsList {
	s.NodeId = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsList) SetParamList(v []*CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) *CreateNodeSupplementRequestCreateCommandNodeParamsList {
	s.ParamList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsList) Validate() error {
	if s.ParamList != nil {
		for _, item := range s.ParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNodeSupplementRequestCreateCommandNodeParamsListParamList struct {
	// The parameter.
	//
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) GetKey() *string {
	return s.Key
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) GetValue() *string {
	return s.Value
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) SetKey(v string) *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList {
	s.Key = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) SetValue(v string) *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList {
	s.Value = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) Validate() error {
	return dara.Validate(s)
}
