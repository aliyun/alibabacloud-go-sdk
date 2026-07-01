// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCommandRequest
	GetClientToken() *string
	SetCommandContent(v string) *CreateCommandRequest
	GetCommandContent() *string
	SetContentEncoding(v string) *CreateCommandRequest
	GetContentEncoding() *string
	SetDescription(v string) *CreateCommandRequest
	GetDescription() *string
	SetEnableParameter(v bool) *CreateCommandRequest
	GetEnableParameter() *bool
	SetLauncher(v string) *CreateCommandRequest
	GetLauncher() *string
	SetName(v string) *CreateCommandRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateCommandRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCommandRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCommandRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCommandRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateCommandRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCommandRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateCommandRequestTag) *CreateCommandRequest
	GetTag() []*CreateCommandRequestTag
	SetTimeout(v int64) *CreateCommandRequest
	GetTimeout() *int64
	SetType(v string) *CreateCommandRequest
	GetType() *string
	SetWorkingDir(v string) *CreateCommandRequest
	GetWorkingDir() *string
}

type CreateCommandRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The Base64-encoded content of the command.
	//
	// - The value of this parameter must be Base64-encoded, and the script content cannot exceed 18 KB in size after Base64 encoding.
	//
	// - The command content supports custom parameters. To enable the custom parameter feature, set `EnableParameter=true`:
	//
	//     - Custom parameters are defined by enclosing them in `{{}}`. Spaces and line breaks before and after the parameter name within `{{}}` are ignored.
	//
	//     - The number of custom parameters cannot exceed 20.
	//
	//     - Custom parameter names can contain letters (a-z, A-Z), digits (0-9), hyphens (-), and underscores (_). The acs:: prefix for specifying non-built-in environment parameters is not supported. Other characters are not supported. Parameter names are case-insensitive.
	//
	//     - Each parameter name cannot exceed 64 bytes in length.
	//
	// - You can specify built-in environment parameters as custom parameters. When the command is run, Cloud Assistant automatically replaces them with the corresponding values from the environment without requiring manual assignment. The following built-in environment parameters are supported:
	//
	//     - `{{ACS::RegionId}}`: the region ID.
	//
	//     - `{{ACS::AccountId}}`: the UID of the Alibaba Cloud account.
	//
	//     - `{{ACS::InstanceId}}`: the instance ID. When a command is sent to multiple instances and you want to use `{{ACS::InstanceId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::InstanceName}}`: the instance name. When a command is sent to multiple instances and you want to use `{{ACS::InstanceName}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than:
	//
	//         - Linux: 2.2.3.344
	//
	//         - Windows: 2.1.3.344
	//
	//     - `{{ACS::InvokeId}}`: the command execution ID. To use `{{ACS::InvokeId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//
	//
	//     - `{{ACS::CommandId}}`: the command ID. When you run a command by calling the [RunCommand](https://help.aliyun.com/document_detail/141751.html) operation and want to use `{{ACS::CommandId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The encoding mode of the command content (CommandContent). Valid values:
	//
	// - PlainText: no encoding. The content is transmitted in plaintext.
	//
	// - Base64: Base64 encoding.
	//
	// Default value: Base64.
	//
	// > If an invalid value is specified, it is treated as Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The command description. All character sets are supported. The description cannot exceed 512 characters in length.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the command uses custom parameters.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The bootstrap program for script execution. The value cannot exceed 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The command name. All character sets are supported. The name cannot exceed 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the command belongs.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	Tag []*CreateCommandRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The maximum timeout period for the command to run on the ECS instance. Unit: seconds. If the command cannot be run for any reason, a timeout occurs. After the timeout, the command process is forcefully terminated by canceling the PID of the command.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the command. Valid values:
	//
	// - RunBatScript: creates a Bat script to run on Windows instances.
	//
	// - RunPowerShellScript: creates a PowerShell script to run on Windows instances.
	//
	// - RunShellScript: creates a Shell script to run on Linux instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The directory where the command is run on the ECS instance. The value cannot exceed 200 characters in length.
	//
	// Default value:
	//
	// - Linux instances: the home directory of the root user, which is `/root`.
	//
	// - Windows instances: the directory where the Cloud Assistant Agent process is located, such as `C:\\Windows\\System32`.
	//
	// > If you set this parameter to a different directory, make sure that the directory exists on the instance.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequest) GoString() string {
	return s.String()
}

func (s *CreateCommandRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *CreateCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *CreateCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCommandRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *CreateCommandRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *CreateCommandRequest) GetName() *string {
	return s.Name
}

func (s *CreateCommandRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCommandRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCommandRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCommandRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCommandRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCommandRequest) GetTag() []*CreateCommandRequestTag {
	return s.Tag
}

func (s *CreateCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateCommandRequest) GetType() *string {
	return s.Type
}

func (s *CreateCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *CreateCommandRequest) SetClientToken(v string) *CreateCommandRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCommandRequest) SetCommandContent(v string) *CreateCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *CreateCommandRequest) SetContentEncoding(v string) *CreateCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *CreateCommandRequest) SetDescription(v string) *CreateCommandRequest {
	s.Description = &v
	return s
}

func (s *CreateCommandRequest) SetEnableParameter(v bool) *CreateCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *CreateCommandRequest) SetLauncher(v string) *CreateCommandRequest {
	s.Launcher = &v
	return s
}

func (s *CreateCommandRequest) SetName(v string) *CreateCommandRequest {
	s.Name = &v
	return s
}

func (s *CreateCommandRequest) SetOwnerAccount(v string) *CreateCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCommandRequest) SetOwnerId(v int64) *CreateCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCommandRequest) SetRegionId(v string) *CreateCommandRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCommandRequest) SetResourceGroupId(v string) *CreateCommandRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCommandRequest) SetResourceOwnerAccount(v string) *CreateCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCommandRequest) SetResourceOwnerId(v int64) *CreateCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCommandRequest) SetTag(v []*CreateCommandRequestTag) *CreateCommandRequest {
	s.Tag = v
	return s
}

func (s *CreateCommandRequest) SetTimeout(v int64) *CreateCommandRequest {
	s.Timeout = &v
	return s
}

func (s *CreateCommandRequest) SetType(v string) *CreateCommandRequest {
	s.Type = &v
	return s
}

func (s *CreateCommandRequest) SetWorkingDir(v string) *CreateCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *CreateCommandRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCommandRequestTag struct {
	// The tag key of the command. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If you use a single tag to filter resources, the resource count with the specified tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources that have all specified tags attached cannot exceed 1,000. If the resource count exceeds 1,000, use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query resources.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCommandRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCommandRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCommandRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCommandRequestTag) SetKey(v string) *CreateCommandRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCommandRequestTag) SetValue(v string) *CreateCommandRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCommandRequestTag) Validate() error {
	return dara.Validate(s)
}
