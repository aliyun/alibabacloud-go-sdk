// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartExecutionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartExecutionShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *StartExecutionShrinkRequest
	GetDescription() *string
	SetLoopMode(v string) *StartExecutionShrinkRequest
	GetLoopMode() *string
	SetMode(v string) *StartExecutionShrinkRequest
	GetMode() *string
	SetParameters(v string) *StartExecutionShrinkRequest
	GetParameters() *string
	SetParentExecutionId(v string) *StartExecutionShrinkRequest
	GetParentExecutionId() *string
	SetRegionId(v string) *StartExecutionShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StartExecutionShrinkRequest
	GetResourceGroupId() *string
	SetSafetyCheck(v string) *StartExecutionShrinkRequest
	GetSafetyCheck() *string
	SetTagsShrink(v string) *StartExecutionShrinkRequest
	GetTagsShrink() *string
	SetTemplateContent(v string) *StartExecutionShrinkRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *StartExecutionShrinkRequest
	GetTemplateName() *string
	SetTemplateURL(v string) *StartExecutionShrinkRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *StartExecutionShrinkRequest
	GetTemplateVersion() *string
}

type StartExecutionShrinkRequest struct {
	// The access token.
	//
	// example:
	//
	// 123e56767-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// test execution.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The loop mode. Valid values:
	//
	// 	- Automatic: does not suspend the execution of the template. This is the default value.
	//
	// 	- FirstBatchPause: suspends the execution of the template after the first batch is complete.
	//
	// 	- EveryBatchPause: suspends the execution of the template after each batch is complete.
	//
	// example:
	//
	// Automatic
	LoopMode *string `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- Automatic: automatically starts the execution of the template. This is the default value.
	//
	// 	- FailurePause: suspends the execution of the template upon a failure.
	//
	// 	- Debug: manually starts the execution of the template.
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The JSON string that consists of a set of parameters. Default value: {}.
	//
	// example:
	//
	// {"Status":"Running"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The ID of the region in which the execution resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security check mode. Valid values:
	//
	// 	- Skip: specifies that you are aware of the risks. The system performs all actions in the execution without manual confirmation, regardless of the risk level. This parameter is valid only if the `Mode` parameter is set to Automatic.
	//
	// 	- ConfirmEveryHighRiskAction: requires you to confirm each high-risk action. This is the default value. You can call the **NotifyExecution*	- operation to confirm or cancel an action.
	//
	// example:
	//
	// Skip
	SafetyCheck *string `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	// The tags for the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The content of the template in the JSON or YAML format. This parameter is the same as the Content parameter that you can specify when you call the CreateTemplate operation. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// {   "Description": "Example template, describe instances in some status",   "FormatVersion": "OOS-2019-06-01",   "Parameters": {},   "Tasks": [     {       "Name": "describeInstances",       "Action": "ACS::ExecuteAPI",       "Description": "desc-en",       "Properties": {         "Service": "ECS",         "API": "DescribeInstances",         "Parameters": {           "Status": "Running"         }       }     }   ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// vmeixme
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The Object Storage Service (OSS) URL of the object that stores the content of the Operation Orchestration Service (OOS) template. The access control list (ACL) of the object must be public-read. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// http://oos-template.cn-hangzhou.oss.aliyun-inc.com/oos-test-template.json
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version number of the template. If you do not specify this parameter, the system uses the latest version.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s StartExecutionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartExecutionShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *StartExecutionShrinkRequest) GetLoopMode() *string {
	return s.LoopMode
}

func (s *StartExecutionShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *StartExecutionShrinkRequest) GetParameters() *string {
	return s.Parameters
}

func (s *StartExecutionShrinkRequest) GetParentExecutionId() *string {
	return s.ParentExecutionId
}

func (s *StartExecutionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartExecutionShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartExecutionShrinkRequest) GetSafetyCheck() *string {
	return s.SafetyCheck
}

func (s *StartExecutionShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *StartExecutionShrinkRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *StartExecutionShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *StartExecutionShrinkRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *StartExecutionShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *StartExecutionShrinkRequest) SetClientToken(v string) *StartExecutionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetDescription(v string) *StartExecutionShrinkRequest {
	s.Description = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetLoopMode(v string) *StartExecutionShrinkRequest {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetMode(v string) *StartExecutionShrinkRequest {
	s.Mode = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetParameters(v string) *StartExecutionShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetParentExecutionId(v string) *StartExecutionShrinkRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetRegionId(v string) *StartExecutionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetResourceGroupId(v string) *StartExecutionShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetSafetyCheck(v string) *StartExecutionShrinkRequest {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTagsShrink(v string) *StartExecutionShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateContent(v string) *StartExecutionShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateName(v string) *StartExecutionShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateURL(v string) *StartExecutionShrinkRequest {
	s.TemplateURL = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateVersion(v string) *StartExecutionShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *StartExecutionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
