// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommands(v []*DescribeCommandsResponseBodyCommands) *DescribeCommandsResponseBody
	GetCommands() []*DescribeCommandsResponseBodyCommands
	SetPageNumber(v int32) *DescribeCommandsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCommandsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCommandsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCommandsResponseBody
	GetTotalCount() *int32
}

type DescribeCommandsResponseBody struct {
	// The commands.
	Commands []*DescribeCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The page number of the command list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of commands.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBody) GetCommands() []*DescribeCommandsResponseBodyCommands {
	return s.Commands
}

func (s *DescribeCommandsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCommandsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCommandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommandsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCommandsResponseBody) SetCommands(v []*DescribeCommandsResponseBodyCommands) *DescribeCommandsResponseBody {
	s.Commands = v
	return s
}

func (s *DescribeCommandsResponseBody) SetPageNumber(v int32) *DescribeCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetPageSize(v int32) *DescribeCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetRequestId(v string) *DescribeCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetTotalCount(v int32) *DescribeCommandsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCommandsResponseBody) Validate() error {
	if s.Commands != nil {
		for _, item := range s.Commands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommandsResponseBodyCommands struct {
	// The command content.
	//
	// example:
	//
	// cat /etc/ssh/sshd_config
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command ID.
	//
	// example:
	//
	// c-gov1k1tqwi9****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The time when the command was created. The time is displayed in UTC (Z) in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2023-01-05T06:38:53Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The command description.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether custom parameters are enabled for the command.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The command name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of custom parameter details.
	ParameterDefinitions []*DescribeCommandsResponseBodyCommandsParameterDefinitions `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty" type:"Repeated"`
	// The list of custom parameter names parsed from CommandContent when the command was created. The list is returned in the List format. If the custom parameter feature is not enabled, an empty list is returned.
	ParameterNames []*string `json:"ParameterNames,omitempty" xml:"ParameterNames,omitempty" type:"Repeated"`
	// The provider of the command.
	//
	// example:
	//
	// User
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the command.
	Tags []*DescribeCommandsResponseBodyCommandsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The command type.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The execution path.
	//
	// example:
	//
	// /home
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeCommandsResponseBodyCommands) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommands) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommands) GetCommandContent() *string {
	return s.CommandContent
}

func (s *DescribeCommandsResponseBodyCommands) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeCommandsResponseBodyCommands) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCommandsResponseBodyCommands) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommandsResponseBodyCommands) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *DescribeCommandsResponseBodyCommands) GetName() *string {
	return s.Name
}

func (s *DescribeCommandsResponseBodyCommands) GetParameterDefinitions() []*DescribeCommandsResponseBodyCommandsParameterDefinitions {
	return s.ParameterDefinitions
}

func (s *DescribeCommandsResponseBodyCommands) GetParameterNames() []*string {
	return s.ParameterNames
}

func (s *DescribeCommandsResponseBodyCommands) GetProvider() *string {
	return s.Provider
}

func (s *DescribeCommandsResponseBodyCommands) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCommandsResponseBodyCommands) GetTags() []*DescribeCommandsResponseBodyCommandsTags {
	return s.Tags
}

func (s *DescribeCommandsResponseBodyCommands) GetTimeout() *int64 {
	return s.Timeout
}

func (s *DescribeCommandsResponseBodyCommands) GetType() *string {
	return s.Type
}

func (s *DescribeCommandsResponseBodyCommands) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeCommandsResponseBodyCommands) SetCommandContent(v string) *DescribeCommandsResponseBodyCommands {
	s.CommandContent = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetCommandId(v string) *DescribeCommandsResponseBodyCommands {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetCreationTime(v string) *DescribeCommandsResponseBodyCommands {
	s.CreationTime = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetDescription(v string) *DescribeCommandsResponseBodyCommands {
	s.Description = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetEnableParameter(v bool) *DescribeCommandsResponseBodyCommands {
	s.EnableParameter = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetName(v string) *DescribeCommandsResponseBodyCommands {
	s.Name = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetParameterDefinitions(v []*DescribeCommandsResponseBodyCommandsParameterDefinitions) *DescribeCommandsResponseBodyCommands {
	s.ParameterDefinitions = v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetParameterNames(v []*string) *DescribeCommandsResponseBodyCommands {
	s.ParameterNames = v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetProvider(v string) *DescribeCommandsResponseBodyCommands {
	s.Provider = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetResourceGroupId(v string) *DescribeCommandsResponseBodyCommands {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetTags(v []*DescribeCommandsResponseBodyCommandsTags) *DescribeCommandsResponseBodyCommands {
	s.Tags = v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetTimeout(v int64) *DescribeCommandsResponseBodyCommands {
	s.Timeout = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetType(v string) *DescribeCommandsResponseBodyCommands {
	s.Type = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetWorkingDir(v string) *DescribeCommandsResponseBodyCommands {
	s.WorkingDir = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) Validate() error {
	if s.ParameterDefinitions != nil {
		for _, item := range s.ParameterDefinitions {
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

type DescribeCommandsResponseBodyCommandsParameterDefinitions struct {
	// The default value of the custom parameter.
	//
	// example:
	//
	// https://aliyun-client-assist.oss-accelerate.aliyuncs.com/linux/aliyun_assist_latest.rpm
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the custom parameter.
	//
	// example:
	//
	// Cloud Assistant Agent installation package path
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom parameter.
	//
	// example:
	//
	// DownloadUrl
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The possible values of the enumeration-type custom parameter.
	PossibleValues []*string `json:"PossibleValues,omitempty" xml:"PossibleValues,omitempty" type:"Repeated"`
	// Indicates whether the custom parameter is required. Valid values:
	//
	// - true: Required.
	//
	// - false: Not required.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeCommandsResponseBodyCommandsParameterDefinitions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsParameterDefinitions) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) GetPossibleValues() []*string {
	return s.PossibleValues
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) GetRequired() *bool {
	return s.Required
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetDefaultValue(v string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.DefaultValue = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetDescription(v string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.Description = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetParameterName(v string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.ParameterName = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetPossibleValues(v []*string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.PossibleValues = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetRequired(v bool) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.Required = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) Validate() error {
	return dara.Validate(s)
}

type DescribeCommandsResponseBodyCommandsTags struct {
	// The tag key of the command.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommandsResponseBodyCommandsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsTags) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeCommandsResponseBodyCommandsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeCommandsResponseBodyCommandsTags) SetKey(v string) *DescribeCommandsResponseBodyCommandsTags {
	s.Key = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsTags) SetValue(v string) *DescribeCommandsResponseBodyCommandsTags {
	s.Value = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsTags) Validate() error {
	return dara.Validate(s)
}
