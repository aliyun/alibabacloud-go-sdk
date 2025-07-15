// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartExecutionRequest
	GetClientToken() *string
	SetDescription(v string) *StartExecutionRequest
	GetDescription() *string
	SetLoopMode(v string) *StartExecutionRequest
	GetLoopMode() *string
	SetMode(v string) *StartExecutionRequest
	GetMode() *string
	SetParameters(v string) *StartExecutionRequest
	GetParameters() *string
	SetParentExecutionId(v string) *StartExecutionRequest
	GetParentExecutionId() *string
	SetRegionId(v string) *StartExecutionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StartExecutionRequest
	GetResourceGroupId() *string
	SetSafetyCheck(v string) *StartExecutionRequest
	GetSafetyCheck() *string
	SetTags(v map[string]interface{}) *StartExecutionRequest
	GetTags() map[string]interface{}
	SetTemplateContent(v string) *StartExecutionRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *StartExecutionRequest
	GetTemplateName() *string
	SetTemplateURL(v string) *StartExecutionRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *StartExecutionRequest
	GetTemplateVersion() *string
}

type StartExecutionRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s StartExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartExecutionRequest) GetDescription() *string {
	return s.Description
}

func (s *StartExecutionRequest) GetLoopMode() *string {
	return s.LoopMode
}

func (s *StartExecutionRequest) GetMode() *string {
	return s.Mode
}

func (s *StartExecutionRequest) GetParameters() *string {
	return s.Parameters
}

func (s *StartExecutionRequest) GetParentExecutionId() *string {
	return s.ParentExecutionId
}

func (s *StartExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartExecutionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartExecutionRequest) GetSafetyCheck() *string {
	return s.SafetyCheck
}

func (s *StartExecutionRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *StartExecutionRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *StartExecutionRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *StartExecutionRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *StartExecutionRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *StartExecutionRequest) SetClientToken(v string) *StartExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *StartExecutionRequest) SetDescription(v string) *StartExecutionRequest {
	s.Description = &v
	return s
}

func (s *StartExecutionRequest) SetLoopMode(v string) *StartExecutionRequest {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionRequest) SetMode(v string) *StartExecutionRequest {
	s.Mode = &v
	return s
}

func (s *StartExecutionRequest) SetParameters(v string) *StartExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *StartExecutionRequest) SetParentExecutionId(v string) *StartExecutionRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionRequest) SetRegionId(v string) *StartExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *StartExecutionRequest) SetResourceGroupId(v string) *StartExecutionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartExecutionRequest) SetSafetyCheck(v string) *StartExecutionRequest {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionRequest) SetTags(v map[string]interface{}) *StartExecutionRequest {
	s.Tags = v
	return s
}

func (s *StartExecutionRequest) SetTemplateContent(v string) *StartExecutionRequest {
	s.TemplateContent = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateName(v string) *StartExecutionRequest {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateURL(v string) *StartExecutionRequest {
	s.TemplateURL = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateVersion(v string) *StartExecutionRequest {
	s.TemplateVersion = &v
	return s
}

func (s *StartExecutionRequest) Validate() error {
	return dara.Validate(s)
}
