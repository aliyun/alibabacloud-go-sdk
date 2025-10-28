// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetId(v string) *GetChangeSetResponseBody
	GetChangeSetId() *string
	SetChangeSetName(v string) *GetChangeSetResponseBody
	GetChangeSetName() *string
	SetChangeSetType(v string) *GetChangeSetResponseBody
	GetChangeSetType() *string
	SetChanges(v []map[string]interface{}) *GetChangeSetResponseBody
	GetChanges() []map[string]interface{}
	SetCreateTime(v string) *GetChangeSetResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetChangeSetResponseBody
	GetDescription() *string
	SetDisableRollback(v bool) *GetChangeSetResponseBody
	GetDisableRollback() *bool
	SetExecutionStatus(v string) *GetChangeSetResponseBody
	GetExecutionStatus() *string
	SetLog(v *GetChangeSetResponseBodyLog) *GetChangeSetResponseBody
	GetLog() *GetChangeSetResponseBodyLog
	SetParameters(v []*GetChangeSetResponseBodyParameters) *GetChangeSetResponseBody
	GetParameters() []*GetChangeSetResponseBodyParameters
	SetRegionId(v string) *GetChangeSetResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetChangeSetResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetChangeSetResponseBody
	GetResourceGroupId() *string
	SetStackId(v string) *GetChangeSetResponseBody
	GetStackId() *string
	SetStackName(v string) *GetChangeSetResponseBody
	GetStackName() *string
	SetStatus(v string) *GetChangeSetResponseBody
	GetStatus() *string
	SetStatusReason(v string) *GetChangeSetResponseBody
	GetStatusReason() *string
	SetTags(v []*GetChangeSetResponseBodyTags) *GetChangeSetResponseBody
	GetTags() []*GetChangeSetResponseBodyTags
	SetTemplateBody(v string) *GetChangeSetResponseBody
	GetTemplateBody() *string
	SetTimeoutInMinutes(v int32) *GetChangeSetResponseBody
	GetTimeoutInMinutes() *int32
}

type GetChangeSetResponseBody struct {
	// The ID of the change set.
	//
	// example:
	//
	// 4c11658d-bd47-4dd0-ba64-727edc62****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The name of the change set.
	//
	// example:
	//
	// ChangeSet_template
	ChangeSetName *string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	// The type of the change set.
	//
	// example:
	//
	// UPDATE
	ChangeSetType *string `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	// The changes of the change set.
	Changes []map[string]interface{} `json:"Changes,omitempty" xml:"Changes,omitempty" type:"Repeated"`
	// The time when the change set was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-01T02:20:56
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the change set.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether rollback was performed when the stack failed to be created or updated.
	//
	// example:
	//
	// false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The execution status of the change set.
	//
	// example:
	//
	// AVAILABLE
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The output logs of the change set.
	Log *GetChangeSetResponseBodyLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The parameters of the stack.
	Parameters []*GetChangeSetResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID of the change set.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3766EE04-76DD-50F9-9C23-3AF136CD5708
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the stack with which the change set is associated.
	//
	// example:
	//
	// a486fc19-ebb7-4ce9-a70b-554a7c3d****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack with which the change set is associated.
	//
	// example:
	//
	// stack_2021-10-13
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The status of the change set.
	//
	// example:
	//
	// CREATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the change set is in its current state.
	//
	// example:
	//
	// too many changes.
	StatusReason *string                         `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	Tags         []*GetChangeSetResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template body of the change set.
	//
	// > This parameter takes effect only if you set ShowTemplate to true.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion": "2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The timeout period that is specified for the stack creation or update operation.
	//
	// example:
	//
	// 60
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s GetChangeSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBody) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *GetChangeSetResponseBody) GetChangeSetName() *string {
	return s.ChangeSetName
}

func (s *GetChangeSetResponseBody) GetChangeSetType() *string {
	return s.ChangeSetType
}

func (s *GetChangeSetResponseBody) GetChanges() []map[string]interface{} {
	return s.Changes
}

func (s *GetChangeSetResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetChangeSetResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetChangeSetResponseBody) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *GetChangeSetResponseBody) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *GetChangeSetResponseBody) GetLog() *GetChangeSetResponseBodyLog {
	return s.Log
}

func (s *GetChangeSetResponseBody) GetParameters() []*GetChangeSetResponseBodyParameters {
	return s.Parameters
}

func (s *GetChangeSetResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetChangeSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChangeSetResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetChangeSetResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *GetChangeSetResponseBody) GetStackName() *string {
	return s.StackName
}

func (s *GetChangeSetResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetChangeSetResponseBody) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetChangeSetResponseBody) GetTags() []*GetChangeSetResponseBodyTags {
	return s.Tags
}

func (s *GetChangeSetResponseBody) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetChangeSetResponseBody) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *GetChangeSetResponseBody) SetChangeSetId(v string) *GetChangeSetResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChangeSetName(v string) *GetChangeSetResponseBody {
	s.ChangeSetName = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChangeSetType(v string) *GetChangeSetResponseBody {
	s.ChangeSetType = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChanges(v []map[string]interface{}) *GetChangeSetResponseBody {
	s.Changes = v
	return s
}

func (s *GetChangeSetResponseBody) SetCreateTime(v string) *GetChangeSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetChangeSetResponseBody) SetDescription(v string) *GetChangeSetResponseBody {
	s.Description = &v
	return s
}

func (s *GetChangeSetResponseBody) SetDisableRollback(v bool) *GetChangeSetResponseBody {
	s.DisableRollback = &v
	return s
}

func (s *GetChangeSetResponseBody) SetExecutionStatus(v string) *GetChangeSetResponseBody {
	s.ExecutionStatus = &v
	return s
}

func (s *GetChangeSetResponseBody) SetLog(v *GetChangeSetResponseBodyLog) *GetChangeSetResponseBody {
	s.Log = v
	return s
}

func (s *GetChangeSetResponseBody) SetParameters(v []*GetChangeSetResponseBodyParameters) *GetChangeSetResponseBody {
	s.Parameters = v
	return s
}

func (s *GetChangeSetResponseBody) SetRegionId(v string) *GetChangeSetResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetRequestId(v string) *GetChangeSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetResourceGroupId(v string) *GetChangeSetResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStackId(v string) *GetChangeSetResponseBody {
	s.StackId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStackName(v string) *GetChangeSetResponseBody {
	s.StackName = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStatus(v string) *GetChangeSetResponseBody {
	s.Status = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStatusReason(v string) *GetChangeSetResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetChangeSetResponseBody) SetTags(v []*GetChangeSetResponseBodyTags) *GetChangeSetResponseBody {
	s.Tags = v
	return s
}

func (s *GetChangeSetResponseBody) SetTemplateBody(v string) *GetChangeSetResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetChangeSetResponseBody) SetTimeoutInMinutes(v int32) *GetChangeSetResponseBody {
	s.TimeoutInMinutes = &v
	return s
}

func (s *GetChangeSetResponseBody) Validate() error {
	if s.Log != nil {
		if err := s.Log.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
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

type GetChangeSetResponseBodyLog struct {
	// The Terraform logs. This parameter is returned only for change sets of Terraform stacks.
	//
	// > This parameter is not returned for change sets that are in the Creating state. This parameter indicates the logs of the change set creation operation for Terraform stacks.
	TerraformLogs []*GetChangeSetResponseBodyLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s GetChangeSetResponseBodyLog) String() string {
	return dara.Prettify(s)
}

func (s GetChangeSetResponseBodyLog) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyLog) GetTerraformLogs() []*GetChangeSetResponseBodyLogTerraformLogs {
	return s.TerraformLogs
}

func (s *GetChangeSetResponseBodyLog) SetTerraformLogs(v []*GetChangeSetResponseBodyLogTerraformLogs) *GetChangeSetResponseBodyLog {
	s.TerraformLogs = v
	return s
}

func (s *GetChangeSetResponseBodyLog) Validate() error {
	if s.TerraformLogs != nil {
		for _, item := range s.TerraformLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChangeSetResponseBodyLogTerraformLogs struct {
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
	// For more information about Terraform commands, see [Command](https://www.terraform.io/cli/commands).
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

func (s GetChangeSetResponseBodyLogTerraformLogs) String() string {
	return dara.Prettify(s)
}

func (s GetChangeSetResponseBodyLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) GetCommand() *string {
	return s.Command
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) GetContent() *string {
	return s.Content
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) GetStream() *string {
	return s.Stream
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) SetCommand(v string) *GetChangeSetResponseBodyLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) SetContent(v string) *GetChangeSetResponseBodyLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) SetStream(v string) *GetChangeSetResponseBodyLogTerraformLogs {
	s.Stream = &v
	return s
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) Validate() error {
	return dara.Validate(s)
}

type GetChangeSetResponseBodyParameters struct {
	// The key of the parameter.
	//
	// example:
	//
	// ALIYUN::Region
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// cn-hangzhou
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetChangeSetResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetChangeSetResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetChangeSetResponseBodyParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetChangeSetResponseBodyParameters) SetParameterKey(v string) *GetChangeSetResponseBodyParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetChangeSetResponseBodyParameters) SetParameterValue(v string) *GetChangeSetResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetChangeSetResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}

type GetChangeSetResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetChangeSetResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetChangeSetResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetChangeSetResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetChangeSetResponseBodyTags) SetKey(v string) *GetChangeSetResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetChangeSetResponseBodyTags) SetValue(v string) *GetChangeSetResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetChangeSetResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
