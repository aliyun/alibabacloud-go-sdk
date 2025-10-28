// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *UpdateStackInstancesRequest
	GetAccountIds() []*string
	SetClientToken(v string) *UpdateStackInstancesRequest
	GetClientToken() *string
	SetDeploymentTargets(v *UpdateStackInstancesRequestDeploymentTargets) *UpdateStackInstancesRequest
	GetDeploymentTargets() *UpdateStackInstancesRequestDeploymentTargets
	SetOperationDescription(v string) *UpdateStackInstancesRequest
	GetOperationDescription() *string
	SetOperationPreferences(v map[string]interface{}) *UpdateStackInstancesRequest
	GetOperationPreferences() map[string]interface{}
	SetParameterOverrides(v []*UpdateStackInstancesRequestParameterOverrides) *UpdateStackInstancesRequest
	GetParameterOverrides() []*UpdateStackInstancesRequestParameterOverrides
	SetRegionId(v string) *UpdateStackInstancesRequest
	GetRegionId() *string
	SetRegionIds(v []*string) *UpdateStackInstancesRequest
	GetRegionIds() []*string
	SetStackGroupName(v string) *UpdateStackInstancesRequest
	GetStackGroupName() *string
	SetTimeoutInMinutes(v int64) *UpdateStackInstancesRequest
	GetTimeoutInMinutes() *int64
}

type UpdateStackInstancesRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > If you want to update stacks in self-managed permission mode, you must specify this parameter.
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which you want to deploy stacks in service-managed mode.
	//
	// > If you want to update stacks in service-managed permission mode, you must specify this parameter.
	DeploymentTargets *UpdateStackInstancesRequestDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
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
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*UpdateStackInstancesRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
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

func (s UpdateStackInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *UpdateStackInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStackInstancesRequest) GetDeploymentTargets() *UpdateStackInstancesRequestDeploymentTargets {
	return s.DeploymentTargets
}

func (s *UpdateStackInstancesRequest) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *UpdateStackInstancesRequest) GetOperationPreferences() map[string]interface{} {
	return s.OperationPreferences
}

func (s *UpdateStackInstancesRequest) GetParameterOverrides() []*UpdateStackInstancesRequestParameterOverrides {
	return s.ParameterOverrides
}

func (s *UpdateStackInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStackInstancesRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *UpdateStackInstancesRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *UpdateStackInstancesRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *UpdateStackInstancesRequest) SetAccountIds(v []*string) *UpdateStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *UpdateStackInstancesRequest) SetClientToken(v string) *UpdateStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetDeploymentTargets(v *UpdateStackInstancesRequestDeploymentTargets) *UpdateStackInstancesRequest {
	s.DeploymentTargets = v
	return s
}

func (s *UpdateStackInstancesRequest) SetOperationDescription(v string) *UpdateStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *UpdateStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

func (s *UpdateStackInstancesRequest) SetParameterOverrides(v []*UpdateStackInstancesRequestParameterOverrides) *UpdateStackInstancesRequest {
	s.ParameterOverrides = v
	return s
}

func (s *UpdateStackInstancesRequest) SetRegionId(v string) *UpdateStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetRegionIds(v []*string) *UpdateStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *UpdateStackInstancesRequest) SetStackGroupName(v string) *UpdateStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetTimeoutInMinutes(v int64) *UpdateStackInstancesRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateStackInstancesRequest) Validate() error {
	if s.DeploymentTargets != nil {
		if err := s.DeploymentTargets.Validate(); err != nil {
			return err
		}
	}
	if s.ParameterOverrides != nil {
		for _, item := range s.ParameterOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateStackInstancesRequestDeploymentTargets struct {
	// The IDs of the member accounts in the resource directory. You can specify up to 20 member account IDs.
	//
	// > To view the member account IDs, go to the **Overview*	- page in the **Resource Management*	- console. For more information, see [View the details of a member](https://help.aliyun.com/document_detail/111624.html).
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The folder IDs of the resource directory.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s UpdateStackInstancesRequestDeploymentTargets) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackInstancesRequestDeploymentTargets) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequestDeploymentTargets) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *UpdateStackInstancesRequestDeploymentTargets) GetRdFolderIds() []*string {
	return s.RdFolderIds
}

func (s *UpdateStackInstancesRequestDeploymentTargets) SetAccountIds(v []*string) *UpdateStackInstancesRequestDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *UpdateStackInstancesRequestDeploymentTargets) SetRdFolderIds(v []*string) *UpdateStackInstancesRequestDeploymentTargets {
	s.RdFolderIds = v
	return s
}

func (s *UpdateStackInstancesRequestDeploymentTargets) Validate() error {
	return dara.Validate(s)
}

type UpdateStackInstancesRequestParameterOverrides struct {
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

func (s UpdateStackInstancesRequestParameterOverrides) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackInstancesRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequestParameterOverrides) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateStackInstancesRequestParameterOverrides) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateStackInstancesRequestParameterOverrides) SetParameterKey(v string) *UpdateStackInstancesRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackInstancesRequestParameterOverrides) SetParameterValue(v string) *UpdateStackInstancesRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

func (s *UpdateStackInstancesRequestParameterOverrides) Validate() error {
	return dara.Validate(s)
}
