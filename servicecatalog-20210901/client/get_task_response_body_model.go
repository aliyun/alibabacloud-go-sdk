// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetTaskDetail(v *GetTaskResponseBodyTaskDetail) *GetTaskResponseBody
	GetTaskDetail() *GetTaskResponseBodyTaskDetail
}

type GetTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the task.
	TaskDetail *GetTaskResponseBodyTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetTaskDetail() *GetTaskResponseBodyTaskDetail {
	return s.TaskDetail
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTaskDetail(v *GetTaskResponseBodyTaskDetail) *GetTaskResponseBody {
	s.TaskDetail = v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.TaskDetail != nil {
		if err := s.TaskDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTaskDetail struct {
	// The time when the task was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-23T09:46:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The logs of the instance.
	Log *GetTaskResponseBodyTaskDetailLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The output parameters of the template.
	Outputs []*GetTaskResponseBodyTaskDetailOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The parameters that are specified in the template.
	Parameters []*GetTaskResponseBodyTaskDetailParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the product portfolio.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// DEMO-Create an ECS instance
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the product version.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	//
	// example:
	//
	// 1.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The ID of the product instance.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	//
	// example:
	//
	// DEMO-ECS instance
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- Succeeded: The task was successful.
	//
	// 	- InProgress: The task was in progress.
	//
	// 	- Failed: The task failed.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message that is returned for the status of the task.
	//
	// > This parameter is returned only when Failed is returned for the Status parameter.
	//
	// example:
	//
	// Resource CREATE failed: terraform stack sc-146611588617****-pp-bp1ddg1n2a***	- failure...
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// task-bp1dmg242c****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The custom tags.
	TaskTags []*GetTaskResponseBodyTaskDetailTaskTags `json:"TaskTags,omitempty" xml:"TaskTags,omitempty" type:"Repeated"`
	// The type of the task. Valid values:
	//
	// 	- LaunchProduct: a task that launches the product.
	//
	// 	- UpdateProvisionedProduct: a task that updates the information about the product instance.
	//
	// 	- TerminateProvisionedProduct: a task that terminates the product instance.
	//
	// example:
	//
	// LaunchProduct
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The time when the task was last modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-23T09:47:29Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTaskResponseBodyTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetail) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTaskResponseBodyTaskDetail) GetLog() *GetTaskResponseBodyTaskDetailLog {
	return s.Log
}

func (s *GetTaskResponseBodyTaskDetail) GetOutputs() []*GetTaskResponseBodyTaskDetailOutputs {
	return s.Outputs
}

func (s *GetTaskResponseBodyTaskDetail) GetParameters() []*GetTaskResponseBodyTaskDetailParameters {
	return s.Parameters
}

func (s *GetTaskResponseBodyTaskDetail) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *GetTaskResponseBodyTaskDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetTaskResponseBodyTaskDetail) GetProductName() *string {
	return s.ProductName
}

func (s *GetTaskResponseBodyTaskDetail) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *GetTaskResponseBodyTaskDetail) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *GetTaskResponseBodyTaskDetail) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *GetTaskResponseBodyTaskDetail) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *GetTaskResponseBodyTaskDetail) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBodyTaskDetail) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *GetTaskResponseBodyTaskDetail) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResponseBodyTaskDetail) GetTaskTags() []*GetTaskResponseBodyTaskDetailTaskTags {
	return s.TaskTags
}

func (s *GetTaskResponseBodyTaskDetail) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskResponseBodyTaskDetail) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTaskResponseBodyTaskDetail) SetCreateTime(v string) *GetTaskResponseBodyTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetLog(v *GetTaskResponseBodyTaskDetailLog) *GetTaskResponseBodyTaskDetail {
	s.Log = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetOutputs(v []*GetTaskResponseBodyTaskDetailOutputs) *GetTaskResponseBodyTaskDetail {
	s.Outputs = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetParameters(v []*GetTaskResponseBodyTaskDetailParameters) *GetTaskResponseBodyTaskDetail {
	s.Parameters = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetPortfolioId(v string) *GetTaskResponseBodyTaskDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductId(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductName(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductName = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductVersionId(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductVersionId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductVersionName(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductVersionName = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProvisionedProductId(v string) *GetTaskResponseBodyTaskDetail {
	s.ProvisionedProductId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProvisionedProductName(v string) *GetTaskResponseBodyTaskDetail {
	s.ProvisionedProductName = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetStatus(v string) *GetTaskResponseBodyTaskDetail {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetStatusMessage(v string) *GetTaskResponseBodyTaskDetail {
	s.StatusMessage = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetTaskId(v string) *GetTaskResponseBodyTaskDetail {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetTaskTags(v []*GetTaskResponseBodyTaskDetailTaskTags) *GetTaskResponseBodyTaskDetail {
	s.TaskTags = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetTaskType(v string) *GetTaskResponseBodyTaskDetail {
	s.TaskType = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetUpdateTime(v string) *GetTaskResponseBodyTaskDetail {
	s.UpdateTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) Validate() error {
	if s.Log != nil {
		if err := s.Log.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.TaskTags != nil {
		for _, item := range s.TaskTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskResponseBodyTaskDetailLog struct {
	// The Terraform logs.
	TerraformLogs []*GetTaskResponseBodyTaskDetailLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s GetTaskResponseBodyTaskDetailLog) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailLog) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailLog) GetTerraformLogs() []*GetTaskResponseBodyTaskDetailLogTerraformLogs {
	return s.TerraformLogs
}

func (s *GetTaskResponseBodyTaskDetailLog) SetTerraformLogs(v []*GetTaskResponseBodyTaskDetailLogTerraformLogs) *GetTaskResponseBodyTaskDetailLog {
	s.TerraformLogs = v
	return s
}

func (s *GetTaskResponseBodyTaskDetailLog) Validate() error {
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

type GetTaskResponseBodyTaskDetailLogTerraformLogs struct {
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
	// 	- stdout: a standard output stream
	//
	// 	- stderr: a standard error stream
	//
	// example:
	//
	// stdout
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailLogTerraformLogs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) GetCommand() *string {
	return s.Command
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) GetContent() *string {
	return s.Content
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) GetStream() *string {
	return s.Stream
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) SetCommand(v string) *GetTaskResponseBodyTaskDetailLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) SetContent(v string) *GetTaskResponseBodyTaskDetailLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) SetStream(v string) *GetTaskResponseBodyTaskDetailLogTerraformLogs {
	s.Stream = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskDetailOutputs struct {
	// The description of the output parameter for the template.
	//
	// example:
	//
	// The ECS instance ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the output parameter for the template.
	//
	// example:
	//
	// instance_id
	OutputKey *string `json:"OutputKey,omitempty" xml:"OutputKey,omitempty"`
	// The value of the output parameter for the template.
	//
	// example:
	//
	// i-xxxxxx
	OutputValue *string `json:"OutputValue,omitempty" xml:"OutputValue,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailOutputs) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailOutputs) GetDescription() *string {
	return s.Description
}

func (s *GetTaskResponseBodyTaskDetailOutputs) GetOutputKey() *string {
	return s.OutputKey
}

func (s *GetTaskResponseBodyTaskDetailOutputs) GetOutputValue() *string {
	return s.OutputValue
}

func (s *GetTaskResponseBodyTaskDetailOutputs) SetDescription(v string) *GetTaskResponseBodyTaskDetailOutputs {
	s.Description = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailOutputs) SetOutputKey(v string) *GetTaskResponseBodyTaskDetailOutputs {
	s.OutputKey = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailOutputs) SetOutputValue(v string) *GetTaskResponseBodyTaskDetailOutputs {
	s.OutputValue = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailOutputs) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskDetailParameters struct {
	// The name of the input parameter for the template.
	//
	// example:
	//
	// instance_type
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the input parameter for the template.
	//
	// example:
	//
	// ecs.s6-c1m1.small
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailParameters) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailParameters) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTaskResponseBodyTaskDetailParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetTaskResponseBodyTaskDetailParameters) SetParameterKey(v string) *GetTaskResponseBodyTaskDetailParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailParameters) SetParameterValue(v string) *GetTaskResponseBodyTaskDetailParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailParameters) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskDetailTaskTags struct {
	// The custom tag key.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The custom tag value.
	//
	// The value must be 1 to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailTaskTags) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailTaskTags) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailTaskTags) GetKey() *string {
	return s.Key
}

func (s *GetTaskResponseBodyTaskDetailTaskTags) GetValue() *string {
	return s.Value
}

func (s *GetTaskResponseBodyTaskDetailTaskTags) SetKey(v string) *GetTaskResponseBodyTaskDetailTaskTags {
	s.Key = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailTaskTags) SetValue(v string) *GetTaskResponseBodyTaskDetailTaskTags {
	s.Value = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailTaskTags) Validate() error {
	return dara.Validate(s)
}
