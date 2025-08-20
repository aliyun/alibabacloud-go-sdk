// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *CreateStackInstancesRequest
	GetAccountIds() []*string
	SetClientToken(v string) *CreateStackInstancesRequest
	GetClientToken() *string
	SetDeploymentOptions(v []*string) *CreateStackInstancesRequest
	GetDeploymentOptions() []*string
	SetDeploymentTargets(v *CreateStackInstancesRequestDeploymentTargets) *CreateStackInstancesRequest
	GetDeploymentTargets() *CreateStackInstancesRequestDeploymentTargets
	SetDisableRollback(v bool) *CreateStackInstancesRequest
	GetDisableRollback() *bool
	SetOperationDescription(v string) *CreateStackInstancesRequest
	GetOperationDescription() *string
	SetOperationPreferences(v map[string]interface{}) *CreateStackInstancesRequest
	GetOperationPreferences() map[string]interface{}
	SetParameterOverrides(v []*CreateStackInstancesRequestParameterOverrides) *CreateStackInstancesRequest
	GetParameterOverrides() []*CreateStackInstancesRequestParameterOverrides
	SetRegionId(v string) *CreateStackInstancesRequest
	GetRegionId() *string
	SetRegionIds(v []*string) *CreateStackInstancesRequest
	GetRegionIds() []*string
	SetStackGroupName(v string) *CreateStackInstancesRequest
	GetStackGroupName() *string
	SetTimeoutInMinutes(v int64) *CreateStackInstancesRequest
	GetTimeoutInMinutes() *int64
}

type CreateStackInstancesRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > You must specify one of the following parameters: `AccountIds` and `DeploymentTargets`.
	//
	// example:
	//
	// ["151266687691****","141261387191****"]
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\\
	//
	// The token can contain letters, digits, hyphens (-), and underscores (_), and cannot exceed 64 characters in length.\\
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeploymentOptions []*string `json:"DeploymentOptions,omitempty" xml:"DeploymentOptions,omitempty" type:"Repeated"`
	// The folders in which ROS deploy stacks in service-managed permission model.
	//
	// > You must specify one of the following parameters: `AccountIds` and `DeploymentTargets`.
	//
	// example:
	//
	// {"RdFolderId": "fd-4PvlVLOL8v"}
	DeploymentTargets *CreateStackInstancesRequestDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
	// Specifies whether to disable rollback when the stack fails to be created.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The description of the stack creation operation.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// Create stack instances in hangzhou and beijing
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the stack creation operation.
	//
	// The following parameters are available:
	//
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, Resource Orchestration Service (ROS) stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
	//
	//     Valid values of N: 0 to 20.
	//
	//     If you do not specify FailureToleranceCount, 0 is used as the default value.
	//
	// -  {"FailureTolerancePercentage": N}
	//
	//     The percentage of the number of accounts within which stack operation failures are allowed to the total number of accounts in each region. If the value of this parameter is exceeded, ROS stops the operation in the region.
	//
	//     Valid values of N: 0 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify FailureTolerancePercentage, 0 is used as the default value.
	//
	// -  {"MaxConcurrentCount": N}
	//
	//    The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//    Valid values of N: 1 to 20.
	//
	//    If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// -  {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which multiple stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the number down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// -  {"RegionConcurrencyType": N}\\
	//
	//     The mode that you want to use to deploy stacks across regions. Valid values:
	//
	//    - SEQUENTIAL (default): deploys stacks in each specified region based on the specified sequence of regions. ROS deploys stacks in one region at a time.
	//
	//    - PARALLEL: deploys stacks in parallel across all specified regions.
	//
	// Separate multiple parameters with commas (,).
	//
	// >-  You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	//
	// >-  You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	//
	// example:
	//
	// {"FailureToleranceCount": 1, "MaxConcurrentCount": 2}
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*CreateStackInstancesRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The region ID of the stack group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to create the stacks. You can specify up to 20 region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cn-hangzhou", "cn-beijing"]
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The name of the stack group. The name must be unique within a region.\\
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The timeout period within which you can create the stack.
	//
	// 	- Default value: 60.
	//
	// 	- Unit: minutes.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s CreateStackInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *CreateStackInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStackInstancesRequest) GetDeploymentOptions() []*string {
	return s.DeploymentOptions
}

func (s *CreateStackInstancesRequest) GetDeploymentTargets() *CreateStackInstancesRequestDeploymentTargets {
	return s.DeploymentTargets
}

func (s *CreateStackInstancesRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *CreateStackInstancesRequest) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *CreateStackInstancesRequest) GetOperationPreferences() map[string]interface{} {
	return s.OperationPreferences
}

func (s *CreateStackInstancesRequest) GetParameterOverrides() []*CreateStackInstancesRequestParameterOverrides {
	return s.ParameterOverrides
}

func (s *CreateStackInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStackInstancesRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *CreateStackInstancesRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *CreateStackInstancesRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *CreateStackInstancesRequest) SetAccountIds(v []*string) *CreateStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *CreateStackInstancesRequest) SetClientToken(v string) *CreateStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackInstancesRequest) SetDeploymentOptions(v []*string) *CreateStackInstancesRequest {
	s.DeploymentOptions = v
	return s
}

func (s *CreateStackInstancesRequest) SetDeploymentTargets(v *CreateStackInstancesRequestDeploymentTargets) *CreateStackInstancesRequest {
	s.DeploymentTargets = v
	return s
}

func (s *CreateStackInstancesRequest) SetDisableRollback(v bool) *CreateStackInstancesRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackInstancesRequest) SetOperationDescription(v string) *CreateStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *CreateStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *CreateStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

func (s *CreateStackInstancesRequest) SetParameterOverrides(v []*CreateStackInstancesRequestParameterOverrides) *CreateStackInstancesRequest {
	s.ParameterOverrides = v
	return s
}

func (s *CreateStackInstancesRequest) SetRegionId(v string) *CreateStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackInstancesRequest) SetRegionIds(v []*string) *CreateStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *CreateStackInstancesRequest) SetStackGroupName(v string) *CreateStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackInstancesRequest) SetTimeoutInMinutes(v int64) *CreateStackInstancesRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateStackInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateStackInstancesRequestDeploymentTargets struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The folder IDs of the resource directory. You can add up to five folder IDs.
	//
	// You can create stacks within all the member accounts in the specified folders. If you create stacks in the Root folder, the stacks are created within all member accounts in the resource directory.
	//
	// > To view the folder IDs, go to the **Overview*	- page in the **Resource Management*	- console. For more information, see [View the basic information about a folder](https://help.aliyun.com/document_detail/111223.html).
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s CreateStackInstancesRequestDeploymentTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateStackInstancesRequestDeploymentTargets) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequestDeploymentTargets) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *CreateStackInstancesRequestDeploymentTargets) GetRdFolderIds() []*string {
	return s.RdFolderIds
}

func (s *CreateStackInstancesRequestDeploymentTargets) SetAccountIds(v []*string) *CreateStackInstancesRequestDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *CreateStackInstancesRequestDeploymentTargets) SetRdFolderIds(v []*string) *CreateStackInstancesRequestDeploymentTargets {
	s.RdFolderIds = v
	return s
}

func (s *CreateStackInstancesRequestDeploymentTargets) Validate() error {
	return dara.Validate(s)
}

type CreateStackInstancesRequestParameterOverrides struct {
	// The key of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the name that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// >-   ParameterOverrides is optional.
	//
	// >-   If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the value that you specify when you create the stack group.
	//
	// Maximum value of N: 200.
	//
	// >-  ParameterOverrides is optional.
	//
	// >-  If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackInstancesRequestParameterOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateStackInstancesRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequestParameterOverrides) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateStackInstancesRequestParameterOverrides) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateStackInstancesRequestParameterOverrides) SetParameterKey(v string) *CreateStackInstancesRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackInstancesRequestParameterOverrides) SetParameterValue(v string) *CreateStackInstancesRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

func (s *CreateStackInstancesRequestParameterOverrides) Validate() error {
	return dara.Validate(s)
}
