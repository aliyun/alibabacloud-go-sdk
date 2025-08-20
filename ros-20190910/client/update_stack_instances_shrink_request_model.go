// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *UpdateStackInstancesShrinkRequest
	GetAccountIdsShrink() *string
	SetClientToken(v string) *UpdateStackInstancesShrinkRequest
	GetClientToken() *string
	SetDeploymentTargetsShrink(v string) *UpdateStackInstancesShrinkRequest
	GetDeploymentTargetsShrink() *string
	SetOperationDescription(v string) *UpdateStackInstancesShrinkRequest
	GetOperationDescription() *string
	SetOperationPreferencesShrink(v string) *UpdateStackInstancesShrinkRequest
	GetOperationPreferencesShrink() *string
	SetParameterOverrides(v []*UpdateStackInstancesShrinkRequestParameterOverrides) *UpdateStackInstancesShrinkRequest
	GetParameterOverrides() []*UpdateStackInstancesShrinkRequestParameterOverrides
	SetRegionId(v string) *UpdateStackInstancesShrinkRequest
	GetRegionId() *string
	SetRegionIdsShrink(v string) *UpdateStackInstancesShrinkRequest
	GetRegionIdsShrink() *string
	SetStackGroupName(v string) *UpdateStackInstancesShrinkRequest
	GetStackGroupName() *string
	SetTimeoutInMinutes(v int64) *UpdateStackInstancesShrinkRequest
	GetTimeoutInMinutes() *int64
}

type UpdateStackInstancesShrinkRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > If you want to update stacks in self-managed permission mode, you must specify this parameter.
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which you want to deploy stacks in service-managed mode.
	//
	// > If you want to update stacks in service-managed permission mode, you must specify this parameter.
	DeploymentTargetsShrink *string `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty"`
	// The description of the update operation.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// Update stack instances in hangzhou and beijing
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the update operation.
	//
	// The following parameters are available:
	//
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, ROS stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
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
	//     The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//     Valid values of N: 1 to 20.
	//
	//     If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// - {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// - {"RegionConcurrencyType": N}
	//
	//   The mode that you want to use to deploy stacks across regions. Valid values:
	//
	//   - SEQUENTIAL (default): deploys stacks in the specified regions one by one in sequence. This way, ROS deploys stacks in only one region at a time.
	//
	//    - PARALLEL: deploys stacks in all the specified regions in parallel.
	//
	// Separate multiple parameters with commas (,).
	//
	// > - You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	//
	// > - You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	//
	// example:
	//
	// {"FailureToleranceCount": 1,"MaxConcurrentCount": 2}
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*UpdateStackInstancesShrinkRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The region ID of the stack group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to update the stacks. You can specify up to 20 region IDs.
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
	// The timeout period for the update operation.
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

func (s UpdateStackInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *UpdateStackInstancesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStackInstancesShrinkRequest) GetDeploymentTargetsShrink() *string {
	return s.DeploymentTargetsShrink
}

func (s *UpdateStackInstancesShrinkRequest) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *UpdateStackInstancesShrinkRequest) GetOperationPreferencesShrink() *string {
	return s.OperationPreferencesShrink
}

func (s *UpdateStackInstancesShrinkRequest) GetParameterOverrides() []*UpdateStackInstancesShrinkRequestParameterOverrides {
	return s.ParameterOverrides
}

func (s *UpdateStackInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStackInstancesShrinkRequest) GetRegionIdsShrink() *string {
	return s.RegionIdsShrink
}

func (s *UpdateStackInstancesShrinkRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *UpdateStackInstancesShrinkRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *UpdateStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetClientToken(v string) *UpdateStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetDeploymentTargetsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.DeploymentTargetsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetOperationDescription(v string) *UpdateStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetParameterOverrides(v []*UpdateStackInstancesShrinkRequestParameterOverrides) *UpdateStackInstancesShrinkRequest {
	s.ParameterOverrides = v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetRegionId(v string) *UpdateStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetStackGroupName(v string) *UpdateStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetTimeoutInMinutes(v int64) *UpdateStackInstancesShrinkRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateStackInstancesShrinkRequestParameterOverrides struct {
	// The key of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the name that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// > -  ParameterOverrides is optional.
	//
	// > - If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the value that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// > -  ParameterOverrides is optional.
	//
	// > - If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackInstancesShrinkRequestParameterOverrides) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackInstancesShrinkRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) SetParameterKey(v string) *UpdateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) SetParameterValue(v string) *UpdateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) Validate() error {
	return dara.Validate(s)
}
