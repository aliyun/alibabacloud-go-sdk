// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionCurrentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunctionCurrentVersionResponseBody
	GetCode() *string
	SetHttpCode(v int64) *GetFunctionCurrentVersionResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *GetFunctionCurrentVersionResponseBody
	GetLatency() *int64
	SetMessage(v string) *GetFunctionCurrentVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFunctionCurrentVersionResponseBody
	GetRequestId() *string
	SetResult(v *GetFunctionCurrentVersionResponseBodyResult) *GetFunctionCurrentVersionResponseBody
	GetResult() *GetFunctionCurrentVersionResponseBodyResult
	SetStatus(v string) *GetFunctionCurrentVersionResponseBody
	GetStatus() *string
}

type GetFunctionCurrentVersionResponseBody struct {
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
	// version not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157479281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *GetFunctionCurrentVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFunctionCurrentVersionResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *GetFunctionCurrentVersionResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *GetFunctionCurrentVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFunctionCurrentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionCurrentVersionResponseBody) GetResult() *GetFunctionCurrentVersionResponseBodyResult {
	return s.Result
}

func (s *GetFunctionCurrentVersionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionCurrentVersionResponseBody) SetCode(v string) *GetFunctionCurrentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetHttpCode(v int64) *GetFunctionCurrentVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetLatency(v int64) *GetFunctionCurrentVersionResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetMessage(v string) *GetFunctionCurrentVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetRequestId(v string) *GetFunctionCurrentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetResult(v *GetFunctionCurrentVersionResponseBodyResult) *GetFunctionCurrentVersionResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetStatus(v string) *GetFunctionCurrentVersionResponseBody {
	s.Status = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionCurrentVersionResponseBodyResult struct {
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
	// The configuration information about the instance.
	VersionConfig *GetFunctionCurrentVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
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

func (s GetFunctionCurrentVersionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionCurrentVersionResponseBodyResult) GetFunctionType() *string {
	return s.FunctionType
}

func (s *GetFunctionCurrentVersionResponseBodyResult) GetModelType() *string {
	return s.ModelType
}

func (s *GetFunctionCurrentVersionResponseBodyResult) GetVersionConfig() *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	return s.VersionConfig
}

func (s *GetFunctionCurrentVersionResponseBodyResult) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetFunctionCurrentVersionResponseBodyResult) GetVersionName() *string {
	return s.VersionName
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetFunctionName(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetFunctionType(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetModelType(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionConfig(v *GetFunctionCurrentVersionResponseBodyResultVersionConfig) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionConfig = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionId(v int64) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionName(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionName = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) Validate() error {
	if s.VersionConfig != nil {
		if err := s.VersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfig struct {
	// The parameters that are used to create the instance.
	//
	// example:
	//
	// [                 {                     "name": "params1",                     "required": "true",                     "formItemProps": "{\\"required\\": true, \\"pattern?\\": \\"/^[a-zA-Z][a-zA-Z0-9_]{0,29}$/\\"}",                     "componentProps": "{\\"component\\": \\"Input\\", \\"attributes\\": {\\"defaultValue\\": \\"value1\\"}}"                 }             ]
	CreateParameters []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The dependencies of the instance.
	Depends []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	// The parameters that are used to use the instance online.
	//
	// example:
	//
	// []
	UsageParameters []*GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfig) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) GetCreateParameters() []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters {
	return s.CreateParameters
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) GetDepends() []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	return s.Depends
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) GetUsageParameters() []*GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters {
	return s.UsageParameters
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetCreateParameters(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetDepends(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.Depends = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetUsageParameters(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.UsageParameters = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) Validate() error {
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

type GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters struct {
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

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) GetName() *string {
	return s.Name
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) GetRequired() *string {
	return s.Required
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) SetName(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) SetRequired(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters {
	s.Required = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) Validate() error {
	return dara.Validate(s)
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends struct {
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

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) GetCondition() *string {
	return s.Condition
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) GetDependency() *string {
	return s.Dependency
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) GetDescription() *string {
	return s.Description
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetCondition(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Condition = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetDependency(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Dependency = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetDescription(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Description = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) Validate() error {
	return dara.Validate(s)
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters struct {
	// The name of the parameter.
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

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) GetName() *string {
	return s.Name
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) GetRequired() *string {
	return s.Required
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) SetName(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) SetRequired(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters {
	s.Required = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) Validate() error {
	return dara.Validate(s)
}
