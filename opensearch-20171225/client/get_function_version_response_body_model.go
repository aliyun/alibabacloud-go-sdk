// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunctionVersionResponseBody
	GetCode() *string
	SetHttpCode(v int64) *GetFunctionVersionResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *GetFunctionVersionResponseBody
	GetLatency() *int64
	SetMessage(v string) *GetFunctionVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFunctionVersionResponseBody
	GetRequestId() *string
	SetResult(v *GetFunctionVersionResponseBodyResult) *GetFunctionVersionResponseBody
	GetResult() *GetFunctionVersionResponseBodyResult
	SetStatus(v string) *GetFunctionVersionResponseBody
	GetStatus() *string
}

type GetFunctionVersionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Version.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The maximum duration for which a task can be executed.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// version not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157479281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result body.
	//
	// example:
	//
	// []
	Result *GetFunctionVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFunctionVersionResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *GetFunctionVersionResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *GetFunctionVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFunctionVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionVersionResponseBody) GetResult() *GetFunctionVersionResponseBodyResult {
	return s.Result
}

func (s *GetFunctionVersionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionVersionResponseBody) SetCode(v string) *GetFunctionVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetHttpCode(v int64) *GetFunctionVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetLatency(v int64) *GetFunctionVersionResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetMessage(v string) *GetFunctionVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetRequestId(v string) *GetFunctionVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetResult(v *GetFunctionVersionResponseBodyResult) *GetFunctionVersionResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionVersionResponseBody) SetStatus(v string) *GetFunctionVersionResponseBody {
	s.Status = &v
	return s
}

func (s *GetFunctionVersionResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionVersionResponseBodyResult struct {
	// The name of the feature.
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- PAAS
	//
	// 	- SAAS
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The configuration information.
	VersionConfig *GetFunctionVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
	// The ID of the version.
	//
	// example:
	//
	// 101
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The name of the version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetFunctionVersionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionVersionResponseBodyResult) GetFunctionType() *string {
	return s.FunctionType
}

func (s *GetFunctionVersionResponseBodyResult) GetModelType() *string {
	return s.ModelType
}

func (s *GetFunctionVersionResponseBodyResult) GetVersionConfig() *GetFunctionVersionResponseBodyResultVersionConfig {
	return s.VersionConfig
}

func (s *GetFunctionVersionResponseBodyResult) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetFunctionVersionResponseBodyResult) GetVersionName() *string {
	return s.VersionName
}

func (s *GetFunctionVersionResponseBodyResult) SetFunctionName(v string) *GetFunctionVersionResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetFunctionType(v string) *GetFunctionVersionResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetModelType(v string) *GetFunctionVersionResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionConfig(v *GetFunctionVersionResponseBodyResultVersionConfig) *GetFunctionVersionResponseBodyResult {
	s.VersionConfig = v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionId(v int64) *GetFunctionVersionResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionName(v string) *GetFunctionVersionResponseBodyResult {
	s.VersionName = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) Validate() error {
	if s.VersionConfig != nil {
		if err := s.VersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionVersionResponseBodyResultVersionConfig struct {
	// The parameters that are used to create the instance.
	//
	// example:
	//
	// [                 {                     "name": "params1",                     "required": "true",                     "formItemProps": "{\\"required\\": true, \\"pattern?\\": \\"/^[a-zA-Z][a-zA-Z0-9_]{0,29}$/\\"}",                     "componentProps": "{\\"component\\": \\"Input\\", \\"attributes\\": {\\"defaultValue\\": \\"value1\\"}}"                 }             ]
	CreateParameters []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The dependencies of the instance.
	Depends []*GetFunctionVersionResponseBodyResultVersionConfigDepends `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	// The parameters that are used during online use of the instance.
	//
	// example:
	//
	// []
	UsageParameters []*GetFunctionVersionResponseBodyResultVersionConfigUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfig) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) GetCreateParameters() []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters {
	return s.CreateParameters
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) GetDepends() []*GetFunctionVersionResponseBodyResultVersionConfigDepends {
	return s.Depends
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) GetUsageParameters() []*GetFunctionVersionResponseBodyResultVersionConfigUsageParameters {
	return s.UsageParameters
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetCreateParameters(v []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetDepends(v []*GetFunctionVersionResponseBodyResultVersionConfigDepends) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.Depends = v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetUsageParameters(v []*GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.UsageParameters = v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) Validate() error {
	if s.CreateParameters != nil {
		for _, item := range s.CreateParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Depends != nil {
		for _, item := range s.Depends {
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

type GetFunctionVersionResponseBodyResultVersionConfigCreateParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// params1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// true
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) GetName() *string {
	return s.Name
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) GetRequired() *string {
	return s.Required
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) SetName(v string) *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) SetRequired(v string) *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters {
	s.Required = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) Validate() error {
	return dara.Validate(s)
}

type GetFunctionVersionResponseBodyResultVersionConfigDepends struct {
	// The condition.
	//
	// example:
	//
	// ""
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The dependency.
	//
	// example:
	//
	// ""
	Dependency *string `json:"Dependency,omitempty" xml:"Dependency,omitempty"`
	// The description.
	//
	// example:
	//
	// ""
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigDepends) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigDepends) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) GetCondition() *string {
	return s.Condition
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) GetDependency() *string {
	return s.Dependency
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) GetDescription() *string {
	return s.Description
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetCondition(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Condition = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetDependency(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Dependency = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetDescription(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Description = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) Validate() error {
	return dara.Validate(s)
}

type GetFunctionVersionResponseBodyResultVersionConfigUsageParameters struct {
	// The name of the instance.
	//
	// example:
	//
	// ""
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// ""
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) GetName() *string {
	return s.Name
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) GetRequired() *string {
	return s.Required
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) SetName(v string) *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) SetRequired(v string) *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters {
	s.Required = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) Validate() error {
	return dara.Validate(s)
}
