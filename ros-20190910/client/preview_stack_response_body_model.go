// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PreviewStackResponseBody
	GetRequestId() *string
	SetStack(v *PreviewStackResponseBodyStack) *PreviewStackResponseBody
	GetStack() *PreviewStackResponseBodyStack
}

type PreviewStackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the stack that is previewed.
	Stack *PreviewStackResponseBodyStack `json:"Stack,omitempty" xml:"Stack,omitempty" type:"Struct"`
}

func (s PreviewStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewStackResponseBody) GetStack() *PreviewStackResponseBodyStack {
	return s.Stack
}

func (s *PreviewStackResponseBody) SetRequestId(v string) *PreviewStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewStackResponseBody) SetStack(v *PreviewStackResponseBodyStack) *PreviewStackResponseBody {
	s.Stack = v
	return s
}

func (s *PreviewStackResponseBody) Validate() error {
	return dara.Validate(s)
}

type PreviewStackResponseBodyStack struct {
	// The description of the stack.
	//
	// example:
	//
	// One ECS instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether rollback is disabled for the resources when the stack fails to be created.
	//
	// example:
	//
	// false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The log that is generated when the stack is run.
	Log *PreviewStackResponseBodyStackLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The parameters of the stack.
	Parameters []*PreviewStackResponseBodyStackParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region where the stack resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources in the stack.
	Resources []*PreviewStackResponseBodyStackResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The stack name.
	//
	// example:
	//
	// MyStack
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure that contains the stack policy body.
	//
	// example:
	//
	// {   "Statement": [     {       "Action": "Update:*",       "Resource": "*",       "Effect": "Allow",       "Principal": "*"     },     {       "Action": "Update:*",       "Resource": "LogicalResourceId/apple1",       "Effect": "Deny",       "Principal": "*"     }   ] }
	StackPolicyBody map[string]interface{} `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// One ECS instance.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// The timeout period for creating the stack.
	//
	// Unit: minutes.
	//
	// example:
	//
	// 60
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s PreviewStackResponseBodyStack) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackResponseBodyStack) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStack) GetDescription() *string {
	return s.Description
}

func (s *PreviewStackResponseBodyStack) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *PreviewStackResponseBodyStack) GetLog() *PreviewStackResponseBodyStackLog {
	return s.Log
}

func (s *PreviewStackResponseBodyStack) GetParameters() []*PreviewStackResponseBodyStackParameters {
	return s.Parameters
}

func (s *PreviewStackResponseBodyStack) GetRegionId() *string {
	return s.RegionId
}

func (s *PreviewStackResponseBodyStack) GetResources() []*PreviewStackResponseBodyStackResources {
	return s.Resources
}

func (s *PreviewStackResponseBodyStack) GetStackName() *string {
	return s.StackName
}

func (s *PreviewStackResponseBodyStack) GetStackPolicyBody() map[string]interface{} {
	return s.StackPolicyBody
}

func (s *PreviewStackResponseBodyStack) GetTemplateDescription() *string {
	return s.TemplateDescription
}

func (s *PreviewStackResponseBodyStack) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *PreviewStackResponseBodyStack) SetDescription(v string) *PreviewStackResponseBodyStack {
	s.Description = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetDisableRollback(v bool) *PreviewStackResponseBodyStack {
	s.DisableRollback = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetLog(v *PreviewStackResponseBodyStackLog) *PreviewStackResponseBodyStack {
	s.Log = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetParameters(v []*PreviewStackResponseBodyStackParameters) *PreviewStackResponseBodyStack {
	s.Parameters = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetRegionId(v string) *PreviewStackResponseBodyStack {
	s.RegionId = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetResources(v []*PreviewStackResponseBodyStackResources) *PreviewStackResponseBodyStack {
	s.Resources = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetStackName(v string) *PreviewStackResponseBodyStack {
	s.StackName = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetStackPolicyBody(v map[string]interface{}) *PreviewStackResponseBodyStack {
	s.StackPolicyBody = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetTemplateDescription(v string) *PreviewStackResponseBodyStack {
	s.TemplateDescription = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetTimeoutInMinutes(v int32) *PreviewStackResponseBodyStack {
	s.TimeoutInMinutes = &v
	return s
}

func (s *PreviewStackResponseBodyStack) Validate() error {
	return dara.Validate(s)
}

type PreviewStackResponseBodyStackLog struct {
	// The Terraform logs. This parameter is returned only if the stack is a Terraform stack.
	//
	// > This parameter contains the logs of previewing the stack.
	TerraformLogs []*PreviewStackResponseBodyStackLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s PreviewStackResponseBodyStackLog) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackResponseBodyStackLog) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackLog) GetTerraformLogs() []*PreviewStackResponseBodyStackLogTerraformLogs {
	return s.TerraformLogs
}

func (s *PreviewStackResponseBodyStackLog) SetTerraformLogs(v []*PreviewStackResponseBodyStackLogTerraformLogs) *PreviewStackResponseBodyStackLog {
	s.TerraformLogs = v
	return s
}

func (s *PreviewStackResponseBodyStackLog) Validate() error {
	return dara.Validate(s)
}

type PreviewStackResponseBodyStackLogTerraformLogs struct {
	// The name of the Terraform command that is run. Valid values:
	//
	// 	- apply
	//
	// 	- plan
	//
	// 	- destroy
	//
	// 	- version
	//
	// For more information about Terraform commands, see [Basic CLI Features](https://www.terraform.io/cli/commands).
	//
	// example:
	//
	// apply
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The content of the output stream that is returned after the command is run.
	//
	// example:
	//
	// Apply complete! Resources: 42 added, 0 changed, 0 destroyed.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The output stream. Valid values:
	//
	// 	- stdout: standard output stream
	//
	// 	- stderr: standard error stream
	//
	// example:
	//
	// stdout
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s PreviewStackResponseBodyStackLogTerraformLogs) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackResponseBodyStackLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) GetCommand() *string {
	return s.Command
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) GetContent() *string {
	return s.Content
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) GetStream() *string {
	return s.Stream
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) SetCommand(v string) *PreviewStackResponseBodyStackLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) SetContent(v string) *PreviewStackResponseBodyStackLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) SetStream(v string) *PreviewStackResponseBodyStackLogTerraformLogs {
	s.Stream = &v
	return s
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) Validate() error {
	return dara.Validate(s)
}

type PreviewStackResponseBodyStackParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// ALIYUN::AccountId
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 151266687691****
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s PreviewStackResponseBodyStackParameters) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackResponseBodyStackParameters) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *PreviewStackResponseBodyStackParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *PreviewStackResponseBodyStackParameters) SetParameterKey(v string) *PreviewStackResponseBodyStackParameters {
	s.ParameterKey = &v
	return s
}

func (s *PreviewStackResponseBodyStackParameters) SetParameterValue(v string) *PreviewStackResponseBodyStackParameters {
	s.ParameterValue = &v
	return s
}

func (s *PreviewStackResponseBodyStackParameters) Validate() error {
	return dara.Validate(s)
}

type PreviewStackResponseBodyStackResources struct {
	// The resource type of an Alibaba Cloud service.
	//
	// example:
	//
	// ACS::ECS::Instance
	AcsResourceType *string `json:"AcsResourceType,omitempty" xml:"AcsResourceType,omitempty"`
	// The action that is performed on the resource. Valid values:
	//
	// 	- Add
	//
	// 	- Modify
	//
	// 	- Remove
	//
	// 	- None
	//
	// example:
	//
	// Add
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The description of the resource.
	//
	// example:
	//
	// ECS instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The logical ID of the resource.
	//
	// example:
	//
	// WebServer
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	//
	// This parameter is returned only if Action is set to Modify or Remove.
	//
	// example:
	//
	// i-a1b2c3***
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The resource properties.
	//
	// example:
	//
	// {   "DiskMappings": [     {       "Category": "cloud_ssd",       "Size": "20"     }   ],   "SystemDisk_Category": "cloud_ssd",   "InstanceChargeType": "PostPaid",   "AutoRenew": "False",   "WillReplace": true,   "ImageId": "centos_7",   "InstanceType": "ecs.g6.large",   "AllocatePublicIP": true,   "AutoRenewPeriod": 1,   "IoOptimized": "optimized",   "ZoneId": "cn-beijing-g",   "VSwitchId": "",   "SecurityGroupId": "",   "Period": 1,   "InternetChargeType": "PayByTraffic",   "SystemDiskCategory": "cloud_efficiency",   "InternetMaxBandwidthOut": 1,   "VpcId": "",   "InternetMaxBandwidthIn": 200,   "PeriodUnit": "Month" }
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// Indicates whether a replacement update is performed on the template. Valid values:
	//
	// 	- True: A replacement update is performed on the template.
	//
	// 	- False: A change is made on the template.
	//
	// 	- Conditional: A replacement update may be performed on the template. You can check whether a replacement update is performed when the template is in use.
	//
	// > This parameter is returned only if Action is set to Modify.
	//
	// example:
	//
	// False
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
	// The resources on which the stack depends.
	RequiredBy []*string `json:"RequiredBy,omitempty" xml:"RequiredBy,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The information about the nested stack. The data structure of the value is the same as the data structure of the entire response.
	//
	// example:
	//
	// {}
	Stack map[string]interface{} `json:"Stack,omitempty" xml:"Stack,omitempty"`
}

func (s PreviewStackResponseBodyStackResources) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackResponseBodyStackResources) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackResources) GetAcsResourceType() *string {
	return s.AcsResourceType
}

func (s *PreviewStackResponseBodyStackResources) GetAction() *string {
	return s.Action
}

func (s *PreviewStackResponseBodyStackResources) GetDescription() *string {
	return s.Description
}

func (s *PreviewStackResponseBodyStackResources) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *PreviewStackResponseBodyStackResources) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *PreviewStackResponseBodyStackResources) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *PreviewStackResponseBodyStackResources) GetReplacement() *string {
	return s.Replacement
}

func (s *PreviewStackResponseBodyStackResources) GetRequiredBy() []*string {
	return s.RequiredBy
}

func (s *PreviewStackResponseBodyStackResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *PreviewStackResponseBodyStackResources) GetStack() map[string]interface{} {
	return s.Stack
}

func (s *PreviewStackResponseBodyStackResources) SetAcsResourceType(v string) *PreviewStackResponseBodyStackResources {
	s.AcsResourceType = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetAction(v string) *PreviewStackResponseBodyStackResources {
	s.Action = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetDescription(v string) *PreviewStackResponseBodyStackResources {
	s.Description = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetLogicalResourceId(v string) *PreviewStackResponseBodyStackResources {
	s.LogicalResourceId = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetPhysicalResourceId(v string) *PreviewStackResponseBodyStackResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetProperties(v map[string]interface{}) *PreviewStackResponseBodyStackResources {
	s.Properties = v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetReplacement(v string) *PreviewStackResponseBodyStackResources {
	s.Replacement = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetRequiredBy(v []*string) *PreviewStackResponseBodyStackResources {
	s.RequiredBy = v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetResourceType(v string) *PreviewStackResponseBodyStackResources {
	s.ResourceType = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetStack(v map[string]interface{}) *PreviewStackResponseBodyStackResources {
	s.Stack = v
	return s
}

func (s *PreviewStackResponseBodyStackResources) Validate() error {
	return dara.Validate(s)
}
