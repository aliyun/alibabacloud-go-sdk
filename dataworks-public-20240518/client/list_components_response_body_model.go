// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListComponentsResponseBodyPagingInfo) *ListComponentsResponseBody
	GetPagingInfo() *ListComponentsResponseBodyPagingInfo
	SetRequestId(v string) *ListComponentsResponseBody
	GetRequestId() *string
}

type ListComponentsResponseBody struct {
	PagingInfo *ListComponentsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 952795279527ab****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) GetPagingInfo() *ListComponentsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComponentsResponseBody) SetPagingInfo(v *ListComponentsResponseBodyPagingInfo) *ListComponentsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListComponentsResponseBody) SetRequestId(v string) *ListComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListComponentsResponseBodyPagingInfo struct {
	Components []*ListComponentsResponseBodyPagingInfoComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComponentsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyPagingInfo) GetComponents() []*ListComponentsResponseBodyPagingInfoComponents {
	return s.Components
}

func (s *ListComponentsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComponentsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComponentsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListComponentsResponseBodyPagingInfo) SetComponents(v []*ListComponentsResponseBodyPagingInfoComponents) *ListComponentsResponseBodyPagingInfo {
	s.Components = v
	return s
}

func (s *ListComponentsResponseBodyPagingInfo) SetPageNumber(v int32) *ListComponentsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfo) SetPageSize(v int32) *ListComponentsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfo) SetTotalCount(v int32) *ListComponentsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListComponentsResponseBodyPagingInfoComponents struct {
	// example:
	//
	// 12312313123
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2023-03-13 16:35:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 对组件的描述
	//
	// example:
	//
	// vpc peering management_staging
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 组件的输入参数列表
	Inputs []*ListComponentsResponseBodyPagingInfoComponentsInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2023-11-30T13:30:58Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// auto_updateAlertRule_test_lJd81f
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 组件的输出参数列表
	Outputs []*ListComponentsResponseBodyPagingInfoComponentsOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// 组件责任人
	//
	// example:
	//
	// 252675537980665607
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 199925
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 工作流的脚本信息
	Script *ListComponentsResponseBodyPagingInfoComponentsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
}

func (s ListComponentsResponseBodyPagingInfoComponents) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyPagingInfoComponents) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetComponentId() *string {
	return s.ComponentId
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetDescription() *string {
	return s.Description
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetInputs() []*ListComponentsResponseBodyPagingInfoComponentsInputs {
	return s.Inputs
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetName() *string {
	return s.Name
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetOutputs() []*ListComponentsResponseBodyPagingInfoComponentsOutputs {
	return s.Outputs
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetOwner() *string {
	return s.Owner
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListComponentsResponseBodyPagingInfoComponents) GetScript() *ListComponentsResponseBodyPagingInfoComponentsScript {
	return s.Script
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetComponentId(v string) *ListComponentsResponseBodyPagingInfoComponents {
	s.ComponentId = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetCreateTime(v string) *ListComponentsResponseBodyPagingInfoComponents {
	s.CreateTime = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetDescription(v string) *ListComponentsResponseBodyPagingInfoComponents {
	s.Description = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetInputs(v []*ListComponentsResponseBodyPagingInfoComponentsInputs) *ListComponentsResponseBodyPagingInfoComponents {
	s.Inputs = v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetModifyTime(v string) *ListComponentsResponseBodyPagingInfoComponents {
	s.ModifyTime = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetName(v string) *ListComponentsResponseBodyPagingInfoComponents {
	s.Name = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetOutputs(v []*ListComponentsResponseBodyPagingInfoComponentsOutputs) *ListComponentsResponseBodyPagingInfoComponents {
	s.Outputs = v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetOwner(v string) *ListComponentsResponseBodyPagingInfoComponents {
	s.Owner = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetProjectId(v int64) *ListComponentsResponseBodyPagingInfoComponents {
	s.ProjectId = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) SetScript(v *ListComponentsResponseBodyPagingInfoComponentsScript) *ListComponentsResponseBodyPagingInfoComponents {
	s.Script = v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponents) Validate() error {
	return dara.Validate(s)
}

type ListComponentsResponseBodyPagingInfoComponentsInputs struct {
	// 输入参数的默认值
	//
	// example:
	//
	// mdb.shard.2x.2xlarge.d
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// 输入参数的描述信息
	//
	// example:
	//
	// None
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 输入参数的名称
	//
	// example:
	//
	// auto_updateAlertRule_test_bULIRo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 输入参数的数据类型
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListComponentsResponseBodyPagingInfoComponentsInputs) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyPagingInfoComponentsInputs) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) GetDescription() *string {
	return s.Description
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) GetName() *string {
	return s.Name
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) GetType() *string {
	return s.Type
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) SetDefaultValue(v string) *ListComponentsResponseBodyPagingInfoComponentsInputs {
	s.DefaultValue = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) SetDescription(v string) *ListComponentsResponseBodyPagingInfoComponentsInputs {
	s.Description = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) SetName(v string) *ListComponentsResponseBodyPagingInfoComponentsInputs {
	s.Name = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) SetType(v string) *ListComponentsResponseBodyPagingInfoComponentsInputs {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsInputs) Validate() error {
	return dara.Validate(s)
}

type ListComponentsResponseBodyPagingInfoComponentsOutputs struct {
	// 输出参数的默认值
	//
	// example:
	//
	// 32000
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// 输出参数的描述信息
	//
	// example:
	//
	// zdy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 输出参数的名称
	//
	// example:
	//
	// auto_updateAlertRule_test_bULIRo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 输出参数的数据类型
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListComponentsResponseBodyPagingInfoComponentsOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyPagingInfoComponentsOutputs) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) GetDescription() *string {
	return s.Description
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) GetName() *string {
	return s.Name
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) GetType() *string {
	return s.Type
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) SetDefaultValue(v string) *ListComponentsResponseBodyPagingInfoComponentsOutputs {
	s.DefaultValue = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) SetDescription(v string) *ListComponentsResponseBodyPagingInfoComponentsOutputs {
	s.Description = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) SetName(v string) *ListComponentsResponseBodyPagingInfoComponentsOutputs {
	s.Name = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) SetType(v string) *ListComponentsResponseBodyPagingInfoComponentsOutputs {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsOutputs) Validate() error {
	return dara.Validate(s)
}

type ListComponentsResponseBodyPagingInfoComponentsScript struct {
	// ID
	//
	// example:
	//
	// 348100
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 脚本路径
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListComponentsResponseBodyPagingInfoComponentsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListComponentsResponseBodyPagingInfoComponentsScript) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyPagingInfoComponentsScript) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScript) GetId() *string {
	return s.Id
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScript) GetPath() *string {
	return s.Path
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScript) GetRuntime() *ListComponentsResponseBodyPagingInfoComponentsScriptRuntime {
	return s.Runtime
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScript) SetId(v string) *ListComponentsResponseBodyPagingInfoComponentsScript {
	s.Id = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScript) SetPath(v string) *ListComponentsResponseBodyPagingInfoComponentsScript {
	s.Path = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScript) SetRuntime(v *ListComponentsResponseBodyPagingInfoComponentsScriptRuntime) *ListComponentsResponseBodyPagingInfoComponentsScript {
	s.Runtime = v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScript) Validate() error {
	return dara.Validate(s)
}

type ListComponentsResponseBodyPagingInfoComponentsScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// SQL_COMPONENT
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListComponentsResponseBodyPagingInfoComponentsScriptRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyPagingInfoComponentsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScriptRuntime) GetCommand() *string {
	return s.Command
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScriptRuntime) SetCommand(v string) *ListComponentsResponseBodyPagingInfoComponentsScriptRuntime {
	s.Command = &v
	return s
}

func (s *ListComponentsResponseBodyPagingInfoComponentsScriptRuntime) Validate() error {
	return dara.Validate(s)
}
