// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFunctionInstancesResponseBody
	GetCode() *string
	SetHttpCode(v int64) *ListFunctionInstancesResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *ListFunctionInstancesResponseBody
	GetLatency() *int64
	SetMessage(v string) *ListFunctionInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFunctionInstancesResponseBody
	GetRequestId() *string
	SetResult(v []*ListFunctionInstancesResponseBodyResult) *ListFunctionInstancesResponseBody
	GetResult() []*ListFunctionInstancesResponseBodyResult
	SetStatus(v string) *ListFunctionInstancesResponseBody
	GetStatus() *string
	SetTotalCount(v int64) *ListFunctionInstancesResponseBody
	GetTotalCount() *int64
}

type ListFunctionInstancesResponseBody struct {
	// The error code. If no error occurs, the parameter is left empty.
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
	// The error message. If no error occurs, the parameter is left empty.
	//
	// example:
	//
	// instance not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the instances.
	//
	// example:
	//
	// []
	Result []*ListFunctionInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request.
	//
	// example:
	//
	// "OK"
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFunctionInstancesResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ListFunctionInstancesResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *ListFunctionInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFunctionInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunctionInstancesResponseBody) GetResult() []*ListFunctionInstancesResponseBodyResult {
	return s.Result
}

func (s *ListFunctionInstancesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListFunctionInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFunctionInstancesResponseBody) SetCode(v string) *ListFunctionInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetHttpCode(v int64) *ListFunctionInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetLatency(v int64) *ListFunctionInstancesResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetMessage(v string) *ListFunctionInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetRequestId(v string) *ListFunctionInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetResult(v []*ListFunctionInstancesResponseBodyResult) *ListFunctionInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetStatus(v string) *ListFunctionInstancesResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetTotalCount(v int64) *ListFunctionInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFunctionInstancesResponseBodyResult struct {
	// The information about the instance.
	//
	// example:
	//
	// {}
	Belongs *ListFunctionInstancesResponseBodyResultBelongs `json:"Belongs,omitempty" xml:"Belongs,omitempty" type:"Struct"`
	// The parameters of the instance.
	//
	// example:
	//
	// []
	CreateParameters []*ListFunctionInstancesResponseBodyResultCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The time when the instance was created.
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
	// The description.
	//
	// example:
	//
	// " "
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information, which is a JSON string. It includes model evaluation information and error information.
	//
	// example:
	//
	// "{\\"dataReport\\":{},\\"errors\\":{}}"
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// "ctr"
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// "PAAS"
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// "ctr_test"
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// "tf_checkpoint"
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// How the instance is created. Valid values:
	//
	// 	- user: The instance is created by user.
	//
	// 	- builtin: The instance is created by system.
	//
	// example:
	//
	// "user"
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the instance. Valid values:
	//
	// 1.  unavailable: No model is available. Models must be trained before you can use them.
	//
	// 2.  available: Models can be used.
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The parameters that are used.
	UsageParameters []*ListFunctionInstancesResponseBodyResultUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
	// The ID of the version.
	//
	// example:
	//
	// 123
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResult) GetBelongs() *ListFunctionInstancesResponseBodyResultBelongs {
	return s.Belongs
}

func (s *ListFunctionInstancesResponseBodyResult) GetCreateParameters() []*ListFunctionInstancesResponseBodyResultCreateParameters {
	return s.CreateParameters
}

func (s *ListFunctionInstancesResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListFunctionInstancesResponseBodyResult) GetCron() *string {
	return s.Cron
}

func (s *ListFunctionInstancesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListFunctionInstancesResponseBodyResult) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListFunctionInstancesResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListFunctionInstancesResponseBodyResult) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ListFunctionInstancesResponseBodyResult) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListFunctionInstancesResponseBodyResult) GetModelType() *string {
	return s.ModelType
}

func (s *ListFunctionInstancesResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListFunctionInstancesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListFunctionInstancesResponseBodyResult) GetUsageParameters() []*ListFunctionInstancesResponseBodyResultUsageParameters {
	return s.UsageParameters
}

func (s *ListFunctionInstancesResponseBodyResult) GetVersionId() *int64 {
	return s.VersionId
}

func (s *ListFunctionInstancesResponseBodyResult) SetBelongs(v *ListFunctionInstancesResponseBodyResultBelongs) *ListFunctionInstancesResponseBodyResult {
	s.Belongs = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCreateParameters(v []*ListFunctionInstancesResponseBodyResultCreateParameters) *ListFunctionInstancesResponseBodyResult {
	s.CreateParameters = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCreateTime(v int64) *ListFunctionInstancesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCron(v string) *ListFunctionInstancesResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetDescription(v string) *ListFunctionInstancesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetExtendInfo(v string) *ListFunctionInstancesResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetFunctionName(v string) *ListFunctionInstancesResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetFunctionType(v string) *ListFunctionInstancesResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetInstanceName(v string) *ListFunctionInstancesResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetModelType(v string) *ListFunctionInstancesResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetSource(v string) *ListFunctionInstancesResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetStatus(v string) *ListFunctionInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetUsageParameters(v []*ListFunctionInstancesResponseBodyResultUsageParameters) *ListFunctionInstancesResponseBodyResult {
	s.UsageParameters = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetVersionId(v int64) *ListFunctionInstancesResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) Validate() error {
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

type ListFunctionInstancesResponseBodyResultBelongs struct {
	// The category.
	//
	// example:
	//
	// "general"
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The industry.
	//
	// example:
	//
	// "ecommerce"
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The abbreviation of the language that applies.
	//
	// example:
	//
	// "zh"
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResultBelongs) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultBelongs) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) GetCategory() *string {
	return s.Category
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) GetDomain() *string {
	return s.Domain
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) GetLanguage() *string {
	return s.Language
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetCategory(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Category = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetDomain(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Domain = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetLanguage(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Language = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) Validate() error {
	return dara.Validate(s)
}

type ListFunctionInstancesResponseBodyResultCreateParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// "param1"
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// "value1"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResultCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultCreateParameters) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) GetName() *string {
	return s.Name
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) GetValue() *string {
	return s.Value
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) SetName(v string) *ListFunctionInstancesResponseBodyResultCreateParameters {
	s.Name = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) SetValue(v string) *ListFunctionInstancesResponseBodyResultCreateParameters {
	s.Value = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) Validate() error {
	return dara.Validate(s)
}

type ListFunctionInstancesResponseBodyResultUsageParameters struct {
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

func (s ListFunctionInstancesResponseBodyResultUsageParameters) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultUsageParameters) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultUsageParameters) GetName() *string {
	return s.Name
}

func (s *ListFunctionInstancesResponseBodyResultUsageParameters) GetValue() *string {
	return s.Value
}

func (s *ListFunctionInstancesResponseBodyResultUsageParameters) SetName(v string) *ListFunctionInstancesResponseBodyResultUsageParameters {
	s.Name = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultUsageParameters) SetValue(v string) *ListFunctionInstancesResponseBodyResultUsageParameters {
	s.Value = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultUsageParameters) Validate() error {
	return dara.Validate(s)
}
