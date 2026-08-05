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
	// The error code.
	//
	// example:
	//
	// not found
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The elapsed time.
	//
	// example:
	//
	// 39.108
	Latency *int64 `json:"latency,omitempty" xml:"latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// "xx not found"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 33E4F0CA-F766-5803-B11C-70DC57A5A6E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*ListFunctionInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// The ownership information.
	Belongs *ListFunctionInstancesResponseBodyResultBelongs `json:"belongs,omitempty" xml:"belongs,omitempty" type:"Struct"`
	// The creation parameter body.
	CreateParameters []*ListFunctionInstancesResponseBodyResultCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 1713352442039
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The cron expression for the timed scheduling node.
	//
	// example:
	//
	// ""
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The extended information.
	//
	// example:
	//
	// ""
	ExtendInfo *string `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// The configuration item.
	//
	// example:
	//
	// nl2sql
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The configuration type.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The configuration name.
	//
	// example:
	//
	// a_test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The model type.
	//
	// example:
	//
	// ops-query-analyze-001
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The instance source. Valid values:
	//
	// - builtin: system instance
	//
	// - user: user instance (default)
	//
	// - all: all instances.
	//
	// example:
	//
	// all
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The status. Valid values:
	//
	// - available
	//
	// - unavailable.
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// usageParameters
	UsageParameters []map[string]interface{} `json:"usageParameters,omitempty" xml:"usageParameters,omitempty" type:"Repeated"`
	// The version ID.
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"versionId,omitempty" xml:"versionId,omitempty"`
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

func (s *ListFunctionInstancesResponseBodyResult) GetUsageParameters() []map[string]interface{} {
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

func (s *ListFunctionInstancesResponseBodyResult) SetUsageParameters(v []map[string]interface{}) *ListFunctionInstancesResponseBodyResult {
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
	return nil
}

type ListFunctionInstancesResponseBodyResultBelongs struct {
	// The category.
	//
	// example:
	//
	// ""
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The industry type.
	//
	// example:
	//
	// ""
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The language. Valid values:
	//
	// - zh_CN: Chinese (default)
	//
	// - en_US: English.
	//
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
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
	// The parameter name.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
