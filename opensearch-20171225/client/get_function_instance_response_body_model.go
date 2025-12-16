// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunctionInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *GetFunctionInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *GetFunctionInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *GetFunctionInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFunctionInstanceResponseBody
	GetRequestId() *string
	SetResult(v *GetFunctionInstanceResponseBodyResult) *GetFunctionInstanceResponseBody
	GetResult() *GetFunctionInstanceResponseBodyResult
	SetStatus(v string) *GetFunctionInstanceResponseBody
	GetStatus() *string
}

type GetFunctionInstanceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// Instance.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// instance not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 68ED4E1B-92B8-5821-A886-9C90686139EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the instance.
	//
	// example:
	//
	// {}
	Result *GetFunctionInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFunctionInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *GetFunctionInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *GetFunctionInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFunctionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionInstanceResponseBody) GetResult() *GetFunctionInstanceResponseBodyResult {
	return s.Result
}

func (s *GetFunctionInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionInstanceResponseBody) SetCode(v string) *GetFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetHttpCode(v int64) *GetFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetLatency(v int64) *GetFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetMessage(v string) *GetFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetRequestId(v string) *GetFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetResult(v *GetFunctionInstanceResponseBodyResult) *GetFunctionInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetStatus(v string) *GetFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionInstanceResponseBodyResult struct {
	// The information about the instance.
	Belongs *GetFunctionInstanceResponseBodyResultBelongs `json:"Belongs,omitempty" xml:"Belongs,omitempty" type:"Struct"`
	// The parameters that are used to create the instance.
	CreateParameters []*GetFunctionInstanceResponseBodyResultCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The time when the task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cron expression used to schedule training, in the format of (Minutes Hours DayofMonth Month DayofWeek). If the value is empty, it indicates that no periodic training is performed.
	//
	// example:
	//
	// 0 3 ? \\	- 0,1,3,5 (at 3 a.m. on Sunday, Monday, Wednesday, and Friday)
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// instance descriptions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information, which is a JSON string.
	//
	// example:
	//
	// {\\"dataReport\\":{},\\"errors\\":{}}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ctr_test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// How the instance is created. Valid values:
	//
	// 	- user: The instance is created by user.
	//
	// 	- builtin: The instance is created by the system.
	//
	// example:
	//
	// user
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the instance. Valid values:
	//
	// 1.  unavailable: No model is available. Models must be trained before you can use them.
	//
	// 2.  available: Models can be used.
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the training task. This parameter is not displayed if no task is available.
	Task *GetFunctionInstanceResponseBodyResultTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// The parameters that are used.
	UsageParameters []*GetFunctionInstanceResponseBodyResultUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
	// The ID of the version.
	//
	// example:
	//
	// 101
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResult) GetBelongs() *GetFunctionInstanceResponseBodyResultBelongs {
	return s.Belongs
}

func (s *GetFunctionInstanceResponseBodyResult) GetCreateParameters() []*GetFunctionInstanceResponseBodyResultCreateParameters {
	return s.CreateParameters
}

func (s *GetFunctionInstanceResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetFunctionInstanceResponseBodyResult) GetCron() *string {
	return s.Cron
}

func (s *GetFunctionInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetFunctionInstanceResponseBodyResult) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetFunctionInstanceResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionInstanceResponseBodyResult) GetFunctionType() *string {
	return s.FunctionType
}

func (s *GetFunctionInstanceResponseBodyResult) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetFunctionInstanceResponseBodyResult) GetModelType() *string {
	return s.ModelType
}

func (s *GetFunctionInstanceResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *GetFunctionInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionInstanceResponseBodyResult) GetTask() *GetFunctionInstanceResponseBodyResultTask {
	return s.Task
}

func (s *GetFunctionInstanceResponseBodyResult) GetUsageParameters() []*GetFunctionInstanceResponseBodyResultUsageParameters {
	return s.UsageParameters
}

func (s *GetFunctionInstanceResponseBodyResult) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetFunctionInstanceResponseBodyResult) SetBelongs(v *GetFunctionInstanceResponseBodyResultBelongs) *GetFunctionInstanceResponseBodyResult {
	s.Belongs = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateParameters(v []*GetFunctionInstanceResponseBodyResultCreateParameters) *GetFunctionInstanceResponseBodyResult {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateTime(v int64) *GetFunctionInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCron(v string) *GetFunctionInstanceResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetDescription(v string) *GetFunctionInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetExtendInfo(v string) *GetFunctionInstanceResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionName(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionType(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetInstanceName(v string) *GetFunctionInstanceResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetModelType(v string) *GetFunctionInstanceResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetSource(v string) *GetFunctionInstanceResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetStatus(v string) *GetFunctionInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetTask(v *GetFunctionInstanceResponseBodyResultTask) *GetFunctionInstanceResponseBodyResult {
	s.Task = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetUsageParameters(v []*GetFunctionInstanceResponseBodyResultUsageParameters) *GetFunctionInstanceResponseBodyResult {
	s.UsageParameters = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetVersionId(v int64) *GetFunctionInstanceResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) Validate() error {
	if s.Belongs != nil {
		if err := s.Belongs.Validate(); err != nil {
			return err
		}
	}
	if s.CreateParameters != nil {
		for _, item := range s.CreateParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	if s.UsageParameters != nil {
		for _, item := range s.UsageParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFunctionInstanceResponseBodyResultBelongs struct {
	// The category.
	//
	// example:
	//
	// general
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The industry.
	//
	// example:
	//
	// ecommerce
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The abbreviation of the language that applies.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultBelongs) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultBelongs) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) GetCategory() *string {
	return s.Category
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) GetDomain() *string {
	return s.Domain
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) GetLanguage() *string {
	return s.Language
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetCategory(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Category = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetDomain(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Domain = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetLanguage(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Language = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) Validate() error {
	return dara.Validate(s)
}

type GetFunctionInstanceResponseBodyResultCreateParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// param1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) GetName() *string {
	return s.Name
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) GetValue() *string {
	return s.Value
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetName(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetValue(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Value = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) Validate() error {
	return dara.Validate(s)
}

type GetFunctionInstanceResponseBodyResultTask struct {
	// The status of the task. Valid values:
	//
	// 	- success: succeeded
	//
	// 	- failed: failed
	//
	// 	- untrained: to be trained
	//
	// 	- pending: being scheduled
	//
	// 	- running: being trained
	//
	// example:
	//
	// success
	DagStatus *string `json:"DagStatus,omitempty" xml:"DagStatus,omitempty"`
	// The time consumed for the most recent run, in milliseconds.
	//
	// example:
	//
	// 1234
	LastRunTime *int64 `json:"LastRunTime,omitempty" xml:"LastRunTime,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultTask) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultTask) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultTask) GetDagStatus() *string {
	return s.DagStatus
}

func (s *GetFunctionInstanceResponseBodyResultTask) GetLastRunTime() *int64 {
	return s.LastRunTime
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetDagStatus(v string) *GetFunctionInstanceResponseBodyResultTask {
	s.DagStatus = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetLastRunTime(v int64) *GetFunctionInstanceResponseBodyResultTask {
	s.LastRunTime = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultTask) Validate() error {
	return dara.Validate(s)
}

type GetFunctionInstanceResponseBodyResultUsageParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// use_param1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultUsageParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultUsageParameters) GetName() *string {
	return s.Name
}

func (s *GetFunctionInstanceResponseBodyResultUsageParameters) GetValue() *string {
	return s.Value
}

func (s *GetFunctionInstanceResponseBodyResultUsageParameters) SetName(v string) *GetFunctionInstanceResponseBodyResultUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultUsageParameters) SetValue(v string) *GetFunctionInstanceResponseBodyResultUsageParameters {
	s.Value = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultUsageParameters) Validate() error {
	return dara.Validate(s)
}
