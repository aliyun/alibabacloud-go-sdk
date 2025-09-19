// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssistInfo(v *GetCheckDetailResponseBodyAssistInfo) *GetCheckDetailResponseBody
	GetAssistInfo() *GetCheckDetailResponseBodyAssistInfo
	SetCustomConfigs(v []*GetCheckDetailResponseBodyCustomConfigs) *GetCheckDetailResponseBody
	GetCustomConfigs() []*GetCheckDetailResponseBodyCustomConfigs
	SetDescription(v *GetCheckDetailResponseBodyDescription) *GetCheckDetailResponseBody
	GetDescription() *GetCheckDetailResponseBodyDescription
	SetRepairReset(v string) *GetCheckDetailResponseBody
	GetRepairReset() *string
	SetRepairSetting(v *GetCheckDetailResponseBodyRepairSetting) *GetCheckDetailResponseBody
	GetRepairSetting() *GetCheckDetailResponseBodyRepairSetting
	SetRepairSupportType(v int32) *GetCheckDetailResponseBody
	GetRepairSupportType() *int32
	SetRequestId(v string) *GetCheckDetailResponseBody
	GetRequestId() *string
	SetSolution(v *GetCheckDetailResponseBodySolution) *GetCheckDetailResponseBody
	GetSolution() *GetCheckDetailResponseBodySolution
}

type GetCheckDetailResponseBody struct {
	// The help information about the check item.
	AssistInfo *GetCheckDetailResponseBodyAssistInfo `json:"AssistInfo,omitempty" xml:"AssistInfo,omitempty" type:"Struct"`
	// The custom configuration items of the check item.
	CustomConfigs []*GetCheckDetailResponseBodyCustomConfigs `json:"CustomConfigs,omitempty" xml:"CustomConfigs,omitempty" type:"Repeated"`
	// The description of the check item.
	Description *GetCheckDetailResponseBodyDescription `json:"Description,omitempty" xml:"Description,omitempty" type:"Struct"`
	// Deprecated
	//
	// >  This parameter is deprecated.
	//
	// example:
	//
	// true
	RepairReset *string `json:"RepairReset,omitempty" xml:"RepairReset,omitempty"`
	// The fixing parameter configurations of the check item.
	RepairSetting *GetCheckDetailResponseBodyRepairSetting `json:"RepairSetting,omitempty" xml:"RepairSetting,omitempty" type:"Struct"`
	// Deprecated
	//
	// >  This parameter is deprecated.
	//
	// example:
	//
	// 1
	RepairSupportType *int32 `json:"RepairSupportType,omitempty" xml:"RepairSupportType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15A6ED6A-DBFE-5255-A248-289907809BEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The solution to handle the risk item.
	Solution *GetCheckDetailResponseBodySolution `json:"Solution,omitempty" xml:"Solution,omitempty" type:"Struct"`
}

func (s GetCheckDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBody) GetAssistInfo() *GetCheckDetailResponseBodyAssistInfo {
	return s.AssistInfo
}

func (s *GetCheckDetailResponseBody) GetCustomConfigs() []*GetCheckDetailResponseBodyCustomConfigs {
	return s.CustomConfigs
}

func (s *GetCheckDetailResponseBody) GetDescription() *GetCheckDetailResponseBodyDescription {
	return s.Description
}

func (s *GetCheckDetailResponseBody) GetRepairReset() *string {
	return s.RepairReset
}

func (s *GetCheckDetailResponseBody) GetRepairSetting() *GetCheckDetailResponseBodyRepairSetting {
	return s.RepairSetting
}

func (s *GetCheckDetailResponseBody) GetRepairSupportType() *int32 {
	return s.RepairSupportType
}

func (s *GetCheckDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckDetailResponseBody) GetSolution() *GetCheckDetailResponseBodySolution {
	return s.Solution
}

func (s *GetCheckDetailResponseBody) SetAssistInfo(v *GetCheckDetailResponseBodyAssistInfo) *GetCheckDetailResponseBody {
	s.AssistInfo = v
	return s
}

func (s *GetCheckDetailResponseBody) SetCustomConfigs(v []*GetCheckDetailResponseBodyCustomConfigs) *GetCheckDetailResponseBody {
	s.CustomConfigs = v
	return s
}

func (s *GetCheckDetailResponseBody) SetDescription(v *GetCheckDetailResponseBodyDescription) *GetCheckDetailResponseBody {
	s.Description = v
	return s
}

func (s *GetCheckDetailResponseBody) SetRepairReset(v string) *GetCheckDetailResponseBody {
	s.RepairReset = &v
	return s
}

func (s *GetCheckDetailResponseBody) SetRepairSetting(v *GetCheckDetailResponseBodyRepairSetting) *GetCheckDetailResponseBody {
	s.RepairSetting = v
	return s
}

func (s *GetCheckDetailResponseBody) SetRepairSupportType(v int32) *GetCheckDetailResponseBody {
	s.RepairSupportType = &v
	return s
}

func (s *GetCheckDetailResponseBody) SetRequestId(v string) *GetCheckDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckDetailResponseBody) SetSolution(v *GetCheckDetailResponseBodySolution) *GetCheckDetailResponseBody {
	s.Solution = v
	return s
}

func (s *GetCheckDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCheckDetailResponseBodyAssistInfo struct {
	// The link to the help information about the risk item when the Type parameter is set to link.
	//
	// example:
	//
	// https://www.alibabacloud.com/help/en/resource-access-management/latest/faq-about-ram-users
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The type of the help information about the risk item. Valid values:
	//
	// 	- **text**
	//
	// 	- **link**
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The content in the help information about the risk item when the Type parameter is set to text.
	//
	// example:
	//
	// Configure an IP address whitelist
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCheckDetailResponseBodyAssistInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBodyAssistInfo) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBodyAssistInfo) GetLink() *string {
	return s.Link
}

func (s *GetCheckDetailResponseBodyAssistInfo) GetType() *string {
	return s.Type
}

func (s *GetCheckDetailResponseBodyAssistInfo) GetValue() *string {
	return s.Value
}

func (s *GetCheckDetailResponseBodyAssistInfo) SetLink(v string) *GetCheckDetailResponseBodyAssistInfo {
	s.Link = &v
	return s
}

func (s *GetCheckDetailResponseBodyAssistInfo) SetType(v string) *GetCheckDetailResponseBodyAssistInfo {
	s.Type = &v
	return s
}

func (s *GetCheckDetailResponseBodyAssistInfo) SetValue(v string) *GetCheckDetailResponseBodyAssistInfo {
	s.Value = &v
	return s
}

func (s *GetCheckDetailResponseBodyAssistInfo) Validate() error {
	return dara.Validate(s)
}

type GetCheckDetailResponseBodyCustomConfigs struct {
	// The default value of the custom configuration item. The value is a string.
	//
	// example:
	//
	// 12
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The name of the custom configuration item, which is unique in a check item.
	//
	// example:
	//
	// SessionTimeMax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the custom configuration item for internationalization.
	//
	// example:
	//
	// Maximum session time
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The type of the custom configuration item. The value is a JSON string.
	//
	// example:
	//
	// {\\"type\\":\\"NUMBER\\",\\"range\\":[1,24]}
	TypeDefine *string `json:"TypeDefine,omitempty" xml:"TypeDefine,omitempty"`
	// The value of the custom configuration item. The value is a string.
	//
	// example:
	//
	// 11
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCheckDetailResponseBodyCustomConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBodyCustomConfigs) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBodyCustomConfigs) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetCheckDetailResponseBodyCustomConfigs) GetName() *string {
	return s.Name
}

func (s *GetCheckDetailResponseBodyCustomConfigs) GetShowName() *string {
	return s.ShowName
}

func (s *GetCheckDetailResponseBodyCustomConfigs) GetTypeDefine() *string {
	return s.TypeDefine
}

func (s *GetCheckDetailResponseBodyCustomConfigs) GetValue() *string {
	return s.Value
}

func (s *GetCheckDetailResponseBodyCustomConfigs) SetDefaultValue(v string) *GetCheckDetailResponseBodyCustomConfigs {
	s.DefaultValue = &v
	return s
}

func (s *GetCheckDetailResponseBodyCustomConfigs) SetName(v string) *GetCheckDetailResponseBodyCustomConfigs {
	s.Name = &v
	return s
}

func (s *GetCheckDetailResponseBodyCustomConfigs) SetShowName(v string) *GetCheckDetailResponseBodyCustomConfigs {
	s.ShowName = &v
	return s
}

func (s *GetCheckDetailResponseBodyCustomConfigs) SetTypeDefine(v string) *GetCheckDetailResponseBodyCustomConfigs {
	s.TypeDefine = &v
	return s
}

func (s *GetCheckDetailResponseBodyCustomConfigs) SetValue(v string) *GetCheckDetailResponseBodyCustomConfigs {
	s.Value = &v
	return s
}

func (s *GetCheckDetailResponseBodyCustomConfigs) Validate() error {
	return dara.Validate(s)
}

type GetCheckDetailResponseBodyDescription struct {
	// The link to the description of the check item.
	//
	// example:
	//
	// https://www.alibabacloud.com/help/en/object-storage-service/latest/tutorial-implement-data-sharing-across-departments-based-on-bucket-policies
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The description type of the check item. The value is fixed as text.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The content in the description of the check item.
	//
	// example:
	//
	// The MSE instance does not enable authentication by default. If public network access is enabled at the same time, the data in the configuration center may be dragged and there is a security risk.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCheckDetailResponseBodyDescription) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBodyDescription) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBodyDescription) GetLink() *string {
	return s.Link
}

func (s *GetCheckDetailResponseBodyDescription) GetType() *string {
	return s.Type
}

func (s *GetCheckDetailResponseBodyDescription) GetValue() *string {
	return s.Value
}

func (s *GetCheckDetailResponseBodyDescription) SetLink(v string) *GetCheckDetailResponseBodyDescription {
	s.Link = &v
	return s
}

func (s *GetCheckDetailResponseBodyDescription) SetType(v string) *GetCheckDetailResponseBodyDescription {
	s.Type = &v
	return s
}

func (s *GetCheckDetailResponseBodyDescription) SetValue(v string) *GetCheckDetailResponseBodyDescription {
	s.Value = &v
	return s
}

func (s *GetCheckDetailResponseBodyDescription) Validate() error {
	return dara.Validate(s)
}

type GetCheckDetailResponseBodyRepairSetting struct {
	// The description of the fixing workflow.
	FlowStep []*GetCheckDetailResponseBodyRepairSettingFlowStep `json:"FlowStep,omitempty" xml:"FlowStep,omitempty" type:"Repeated"`
	// The configurations of the fixing parameters.
	RepairConfigs []*GetCheckDetailResponseBodyRepairSettingRepairConfigs `json:"RepairConfigs,omitempty" xml:"RepairConfigs,omitempty" type:"Repeated"`
	// Indicates whether a restart is required after the fixing. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	RepairReset *bool `json:"RepairReset,omitempty" xml:"RepairReset,omitempty"`
	// Indicates whether the check item supports the quick fix feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	RepairSupport *bool `json:"RepairSupport,omitempty" xml:"RepairSupport,omitempty"`
	// The fixing type that is supported. Valid values:
	//
	// 	- **1**: The fixing and rollback are supported.
	//
	// 	- **2**: The fixing is supported, but the rollback is not supported.
	//
	// 	- **3**: The fixing must be performed on a third-party platform.
	//
	// example:
	//
	// 1
	RepairSupportType *int32 `json:"RepairSupportType,omitempty" xml:"RepairSupportType,omitempty"`
}

func (s GetCheckDetailResponseBodyRepairSetting) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBodyRepairSetting) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBodyRepairSetting) GetFlowStep() []*GetCheckDetailResponseBodyRepairSettingFlowStep {
	return s.FlowStep
}

func (s *GetCheckDetailResponseBodyRepairSetting) GetRepairConfigs() []*GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	return s.RepairConfigs
}

func (s *GetCheckDetailResponseBodyRepairSetting) GetRepairReset() *bool {
	return s.RepairReset
}

func (s *GetCheckDetailResponseBodyRepairSetting) GetRepairSupport() *bool {
	return s.RepairSupport
}

func (s *GetCheckDetailResponseBodyRepairSetting) GetRepairSupportType() *int32 {
	return s.RepairSupportType
}

func (s *GetCheckDetailResponseBodyRepairSetting) SetFlowStep(v []*GetCheckDetailResponseBodyRepairSettingFlowStep) *GetCheckDetailResponseBodyRepairSetting {
	s.FlowStep = v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSetting) SetRepairConfigs(v []*GetCheckDetailResponseBodyRepairSettingRepairConfigs) *GetCheckDetailResponseBodyRepairSetting {
	s.RepairConfigs = v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSetting) SetRepairReset(v bool) *GetCheckDetailResponseBodyRepairSetting {
	s.RepairReset = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSetting) SetRepairSupport(v bool) *GetCheckDetailResponseBodyRepairSetting {
	s.RepairSupport = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSetting) SetRepairSupportType(v int32) *GetCheckDetailResponseBodyRepairSetting {
	s.RepairSupportType = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSetting) Validate() error {
	return dara.Validate(s)
}

type GetCheckDetailResponseBodyRepairSettingFlowStep struct {
	// The text description of the fixing step.
	//
	// example:
	//
	// The first step is to open the calling interface.
	ShowText *string `json:"ShowText,omitempty" xml:"ShowText,omitempty"`
	// The sequence number of the fixing step.
	//
	// example:
	//
	// 1
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s GetCheckDetailResponseBodyRepairSettingFlowStep) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBodyRepairSettingFlowStep) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBodyRepairSettingFlowStep) GetShowText() *string {
	return s.ShowText
}

func (s *GetCheckDetailResponseBodyRepairSettingFlowStep) GetStep() *string {
	return s.Step
}

func (s *GetCheckDetailResponseBodyRepairSettingFlowStep) SetShowText(v string) *GetCheckDetailResponseBodyRepairSettingFlowStep {
	s.ShowText = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingFlowStep) SetStep(v string) *GetCheckDetailResponseBodyRepairSettingFlowStep {
	s.Step = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingFlowStep) Validate() error {
	return dara.Validate(s)
}

type GetCheckDetailResponseBodyRepairSettingRepairConfigs struct {
	// Indicates whether the value of the parameter is displayed in the console. Valid values:
	//
	// 	- 0: The historical value and real-time value of the parameter are displayed.
	//
	// 	- 1: Only the real-time value of the parameter is displayed.
	//
	// 	- 2: The value of the parameter is not displayed in the console.
	//
	// example:
	//
	// 0
	ConsoleParamType *string `json:"ConsoleParamType,omitempty" xml:"ConsoleParamType,omitempty"`
	// Indicates whether custom configurations of the fixing parameters are supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CustomFlag *bool `json:"CustomFlag,omitempty" xml:"CustomFlag,omitempty"`
	// Indicates whether data needs to be encrypted during transmission. Valid values:
	//
	// 	- 0: Data does not need to be encrypted during transmission.
	//
	// 	- 1: Data needs to be encrypted during transmission.
	//
	// 	- 2: Data needs to be encrypted during transmission, and the user must perform secondary confirmation.
	//
	// example:
	//
	// 1
	DataTransformType *string `json:"DataTransformType,omitempty" xml:"DataTransformType,omitempty"`
	// The default value of the parameter. The value is a string.
	//
	// example:
	//
	// 1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Indicates whether this parameter is specified by the user. Valid values:
	//
	// 	- 0: The default value is used.
	//
	// 	- 1: This parameter is required, and no default value is specified.
	//
	// 	- 2: This parameter can be left empty.
	//
	// example:
	//
	// 1
	EmptyParamSwitch *string `json:"EmptyParamSwitch,omitempty" xml:"EmptyParamSwitch,omitempty"`
	// The fixing parameters that are not compatible with this parameter.
	ExclusiveName []*string `json:"ExclusiveName,omitempty" xml:"ExclusiveName,omitempty" type:"Repeated"`
	// The ID of the fixing workflow.
	//
	// example:
	//
	// 64312d3ee19d470a9b54393dab****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the parameter. The name must be unique within the check item.
	//
	// example:
	//
	// navicat
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the parameter.
	//
	// example:
	//
	// port
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The type of the parameter. The value is a JSON string.
	//
	// example:
	//
	// {\\"type\\":\\"NUMBER\\",\\"range\\":[1,24]}
	TypeDefine *string `json:"TypeDefine,omitempty" xml:"TypeDefine,omitempty"`
	// The type of the parameter. Valid values:
	//
	// 	- 1: asset parameters that are required during fixing.
	//
	// 	- 2: user-provided parameters that are required during fixing.
	//
	// 	- 3: parameters that are temporarily provided by the user.
	//
	// example:
	//
	// 1
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	// The user-configured value of the parameter. The value is a string.
	//
	// example:
	//
	// 2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCheckDetailResponseBodyRepairSettingRepairConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBodyRepairSettingRepairConfigs) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetConsoleParamType() *string {
	return s.ConsoleParamType
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetCustomFlag() *bool {
	return s.CustomFlag
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetDataTransformType() *string {
	return s.DataTransformType
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetEmptyParamSwitch() *string {
	return s.EmptyParamSwitch
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetExclusiveName() []*string {
	return s.ExclusiveName
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetFlowId() *string {
	return s.FlowId
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetName() *string {
	return s.Name
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetShowName() *string {
	return s.ShowName
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetTypeDefine() *string {
	return s.TypeDefine
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetUsageType() *string {
	return s.UsageType
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) GetValue() *string {
	return s.Value
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetConsoleParamType(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.ConsoleParamType = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetCustomFlag(v bool) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.CustomFlag = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetDataTransformType(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.DataTransformType = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetDefaultValue(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.DefaultValue = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetEmptyParamSwitch(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.EmptyParamSwitch = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetExclusiveName(v []*string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.ExclusiveName = v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetFlowId(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.FlowId = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetName(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.Name = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetShowName(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.ShowName = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetTypeDefine(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.TypeDefine = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetUsageType(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.UsageType = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) SetValue(v string) *GetCheckDetailResponseBodyRepairSettingRepairConfigs {
	s.Value = &v
	return s
}

func (s *GetCheckDetailResponseBodyRepairSettingRepairConfigs) Validate() error {
	return dara.Validate(s)
}

type GetCheckDetailResponseBodySolution struct {
	// The link to the solution to handle the risk item when the Type parameter is set to link.
	//
	// example:
	//
	// https://www.alibabacloud.com/help/en/object-storage-service/latest/tutorial-implement-data-sharing-across-departments-based-on-bucket-policies
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The type of the solution to handle the risk item. Valid values:
	//
	// 	- **text**
	//
	// 	- **link**
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The content of the solution to handle the risk item when the Type parameter is set to text.
	//
	// example:
	//
	// Enter the MSE product console - registration and configuration center - instance list, click the corresponding instance name to enter the instance details, find the public network whitelist setting option in the basic information, and configure the whitelist according to business needs. It is forbidden to configure 0.0.0.0 or the whitelist as null.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCheckDetailResponseBodySolution) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponseBodySolution) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponseBodySolution) GetLink() *string {
	return s.Link
}

func (s *GetCheckDetailResponseBodySolution) GetType() *string {
	return s.Type
}

func (s *GetCheckDetailResponseBodySolution) GetValue() *string {
	return s.Value
}

func (s *GetCheckDetailResponseBodySolution) SetLink(v string) *GetCheckDetailResponseBodySolution {
	s.Link = &v
	return s
}

func (s *GetCheckDetailResponseBodySolution) SetType(v string) *GetCheckDetailResponseBodySolution {
	s.Type = &v
	return s
}

func (s *GetCheckDetailResponseBodySolution) SetValue(v string) *GetCheckDetailResponseBodySolution {
	s.Value = &v
	return s
}

func (s *GetCheckDetailResponseBodySolution) Validate() error {
	return dara.Validate(s)
}
