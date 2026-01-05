// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetTaskDetails(v []*ListTasksResponseBodyTaskDetails) *ListTasksResponseBody
	GetTaskDetails() []*ListTasksResponseBodyTaskDetails
	SetTotalCount(v int32) *ListTasksResponseBody
	GetTotalCount() *int32
}

type ListTasksResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskDetails []*ListTasksResponseBodyTaskDetails `json:"TaskDetails,omitempty" xml:"TaskDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetTaskDetails() []*ListTasksResponseBodyTaskDetails {
	return s.TaskDetails
}

func (s *ListTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTaskDetails(v []*ListTasksResponseBodyTaskDetails) *ListTasksResponseBody {
	s.TaskDetails = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	if s.TaskDetails != nil {
		for _, item := range s.TaskDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTasksResponseBodyTaskDetails struct {
	// example:
	//
	// 2022-05-23T09:46:27Z
	CreateTime *string                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Log        *ListTasksResponseBodyTaskDetailsLog          `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	Outputs    []*ListTasksResponseBodyTaskDetailsOutputs    `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	Parameters []*ListTasksResponseBodyTaskDetailsParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// example:
	//
	// prod-bp18r7q127****
	ProductId   *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// example:
	//
	// 1.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId   *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Resource CREATE failed: terraform stack sc-146611588617****-pp-bp1ddg1n2a***	- failure...
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// task-bp1dmg242c****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// LaunchProduct
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 2022-05-26T03:28:45Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTasksResponseBodyTaskDetails) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetails) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTasksResponseBodyTaskDetails) GetLog() *ListTasksResponseBodyTaskDetailsLog {
	return s.Log
}

func (s *ListTasksResponseBodyTaskDetails) GetOutputs() []*ListTasksResponseBodyTaskDetailsOutputs {
	return s.Outputs
}

func (s *ListTasksResponseBodyTaskDetails) GetParameters() []*ListTasksResponseBodyTaskDetailsParameters {
	return s.Parameters
}

func (s *ListTasksResponseBodyTaskDetails) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *ListTasksResponseBodyTaskDetails) GetProductId() *string {
	return s.ProductId
}

func (s *ListTasksResponseBodyTaskDetails) GetProductName() *string {
	return s.ProductName
}

func (s *ListTasksResponseBodyTaskDetails) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *ListTasksResponseBodyTaskDetails) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *ListTasksResponseBodyTaskDetails) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *ListTasksResponseBodyTaskDetails) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *ListTasksResponseBodyTaskDetails) GetStatus() *string {
	return s.Status
}

func (s *ListTasksResponseBodyTaskDetails) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListTasksResponseBodyTaskDetails) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksResponseBodyTaskDetails) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTasksResponseBodyTaskDetails) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTasksResponseBodyTaskDetails) SetCreateTime(v string) *ListTasksResponseBodyTaskDetails {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetLog(v *ListTasksResponseBodyTaskDetailsLog) *ListTasksResponseBodyTaskDetails {
	s.Log = v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetOutputs(v []*ListTasksResponseBodyTaskDetailsOutputs) *ListTasksResponseBodyTaskDetails {
	s.Outputs = v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetParameters(v []*ListTasksResponseBodyTaskDetailsParameters) *ListTasksResponseBodyTaskDetails {
	s.Parameters = v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetPortfolioId(v string) *ListTasksResponseBodyTaskDetails {
	s.PortfolioId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductId(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductName(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductName = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductVersionId(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductVersionId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductVersionName(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductVersionName = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProvisionedProductId(v string) *ListTasksResponseBodyTaskDetails {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProvisionedProductName(v string) *ListTasksResponseBodyTaskDetails {
	s.ProvisionedProductName = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetStatus(v string) *ListTasksResponseBodyTaskDetails {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetStatusMessage(v string) *ListTasksResponseBodyTaskDetails {
	s.StatusMessage = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetTaskId(v string) *ListTasksResponseBodyTaskDetails {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetTaskType(v string) *ListTasksResponseBodyTaskDetails {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetUpdateTime(v string) *ListTasksResponseBodyTaskDetails {
	s.UpdateTime = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) Validate() error {
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
	return nil
}

type ListTasksResponseBodyTaskDetailsLog struct {
	TerraformLogs []*ListTasksResponseBodyTaskDetailsLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s ListTasksResponseBodyTaskDetailsLog) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsLog) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsLog) GetTerraformLogs() []*ListTasksResponseBodyTaskDetailsLogTerraformLogs {
	return s.TerraformLogs
}

func (s *ListTasksResponseBodyTaskDetailsLog) SetTerraformLogs(v []*ListTasksResponseBodyTaskDetailsLogTerraformLogs) *ListTasksResponseBodyTaskDetailsLog {
	s.TerraformLogs = v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsLog) Validate() error {
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

type ListTasksResponseBodyTaskDetailsLogTerraformLogs struct {
	// example:
	//
	// apply
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// example:
	//
	// Apply complete! Resources: 42 added, 0 changed, 0 destroyed.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// stdout
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s ListTasksResponseBodyTaskDetailsLogTerraformLogs) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) GetCommand() *string {
	return s.Command
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) GetContent() *string {
	return s.Content
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) GetStream() *string {
	return s.Stream
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) SetCommand(v string) *ListTasksResponseBodyTaskDetailsLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) SetContent(v string) *ListTasksResponseBodyTaskDetailsLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) SetStream(v string) *ListTasksResponseBodyTaskDetailsLogTerraformLogs {
	s.Stream = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) Validate() error {
	return dara.Validate(s)
}

type ListTasksResponseBodyTaskDetailsOutputs struct {
	// example:
	//
	// The ECS instance ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// instance_id
	OutputKey *string `json:"OutputKey,omitempty" xml:"OutputKey,omitempty"`
	// example:
	//
	// i-xxxxxx
	OutputValue *string `json:"OutputValue,omitempty" xml:"OutputValue,omitempty"`
}

func (s ListTasksResponseBodyTaskDetailsOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsOutputs) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) GetDescription() *string {
	return s.Description
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) GetOutputKey() *string {
	return s.OutputKey
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) GetOutputValue() *string {
	return s.OutputValue
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) SetDescription(v string) *ListTasksResponseBodyTaskDetailsOutputs {
	s.Description = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) SetOutputKey(v string) *ListTasksResponseBodyTaskDetailsOutputs {
	s.OutputKey = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) SetOutputValue(v string) *ListTasksResponseBodyTaskDetailsOutputs {
	s.OutputValue = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) Validate() error {
	return dara.Validate(s)
}

type ListTasksResponseBodyTaskDetailsParameters struct {
	// example:
	//
	// instance_type
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// example:
	//
	// ecs.s6-c1m1.small
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ListTasksResponseBodyTaskDetailsParameters) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsParameters) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *ListTasksResponseBodyTaskDetailsParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ListTasksResponseBodyTaskDetailsParameters) SetParameterKey(v string) *ListTasksResponseBodyTaskDetailsParameters {
	s.ParameterKey = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsParameters) SetParameterValue(v string) *ListTasksResponseBodyTaskDetailsParameters {
	s.ParameterValue = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsParameters) Validate() error {
	return dara.Validate(s)
}
