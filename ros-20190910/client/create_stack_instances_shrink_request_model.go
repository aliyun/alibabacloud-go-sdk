// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *CreateStackInstancesShrinkRequest
	GetAccountIdsShrink() *string
	SetClientToken(v string) *CreateStackInstancesShrinkRequest
	GetClientToken() *string
	SetDeploymentOptions(v []*string) *CreateStackInstancesShrinkRequest
	GetDeploymentOptions() []*string
	SetDeploymentTargetsShrink(v string) *CreateStackInstancesShrinkRequest
	GetDeploymentTargetsShrink() *string
	SetDisableRollback(v bool) *CreateStackInstancesShrinkRequest
	GetDisableRollback() *bool
	SetOperationDescription(v string) *CreateStackInstancesShrinkRequest
	GetOperationDescription() *string
	SetOperationPreferencesShrink(v string) *CreateStackInstancesShrinkRequest
	GetOperationPreferencesShrink() *string
	SetParameterOverrides(v []*CreateStackInstancesShrinkRequestParameterOverrides) *CreateStackInstancesShrinkRequest
	GetParameterOverrides() []*CreateStackInstancesShrinkRequestParameterOverrides
	SetRegionId(v string) *CreateStackInstancesShrinkRequest
	GetRegionId() *string
	SetRegionIdsShrink(v string) *CreateStackInstancesShrinkRequest
	GetRegionIdsShrink() *string
	SetStackGroupName(v string) *CreateStackInstancesShrinkRequest
	GetStackGroupName() *string
	SetTimeoutInMinutes(v int64) *CreateStackInstancesShrinkRequest
	GetTimeoutInMinutes() *int64
}

type CreateStackInstancesShrinkRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > You must specify one of the following parameters: `AccountIds` and `DeploymentTargets`.
	//
	// example:
	//
	// ["151266687691****","141261387191****"]
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
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
	DeploymentTargetsShrink *string `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty"`
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
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*CreateStackInstancesShrinkRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
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
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
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

func (s CreateStackInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *CreateStackInstancesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStackInstancesShrinkRequest) GetDeploymentOptions() []*string {
	return s.DeploymentOptions
}

func (s *CreateStackInstancesShrinkRequest) GetDeploymentTargetsShrink() *string {
	return s.DeploymentTargetsShrink
}

func (s *CreateStackInstancesShrinkRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *CreateStackInstancesShrinkRequest) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *CreateStackInstancesShrinkRequest) GetOperationPreferencesShrink() *string {
	return s.OperationPreferencesShrink
}

func (s *CreateStackInstancesShrinkRequest) GetParameterOverrides() []*CreateStackInstancesShrinkRequestParameterOverrides {
	return s.ParameterOverrides
}

func (s *CreateStackInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStackInstancesShrinkRequest) GetRegionIdsShrink() *string {
	return s.RegionIdsShrink
}

func (s *CreateStackInstancesShrinkRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *CreateStackInstancesShrinkRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *CreateStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetClientToken(v string) *CreateStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetDeploymentOptions(v []*string) *CreateStackInstancesShrinkRequest {
	s.DeploymentOptions = v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetDeploymentTargetsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.DeploymentTargetsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetDisableRollback(v bool) *CreateStackInstancesShrinkRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetOperationDescription(v string) *CreateStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *CreateStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetParameterOverrides(v []*CreateStackInstancesShrinkRequestParameterOverrides) *CreateStackInstancesShrinkRequest {
	s.ParameterOverrides = v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetRegionId(v string) *CreateStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetStackGroupName(v string) *CreateStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetTimeoutInMinutes(v int64) *CreateStackInstancesShrinkRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateStackInstancesShrinkRequestParameterOverrides struct {
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

func (s CreateStackInstancesShrinkRequestParameterOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateStackInstancesShrinkRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) SetParameterKey(v string) *CreateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) SetParameterValue(v string) *CreateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) Validate() error {
	return dara.Validate(s)
}
