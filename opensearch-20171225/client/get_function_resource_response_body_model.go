// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunctionResourceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *GetFunctionResourceResponseBody
	GetHttpCode() *int64
	SetLatency(v float64) *GetFunctionResourceResponseBody
	GetLatency() *float64
	SetMessage(v string) *GetFunctionResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFunctionResourceResponseBody
	GetRequestId() *string
	SetResult(v *GetFunctionResourceResponseBodyResult) *GetFunctionResourceResponseBody
	GetResult() *GetFunctionResourceResponseBodyResult
	SetStatus(v string) *GetFunctionResourceResponseBody
	GetStatus() *string
}

type GetFunctionResourceResponseBody struct {
	// The error code returned. If no error occurs, this value is empty.
	//
	// example:
	//
	// Resource.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the API request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Resource not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E215C843-0795-5293-AC9A-14FE0723E890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Result *GetFunctionResourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The HTTP status code. Valid values:
	//
	// 	- OK
	//
	// 	- FAIL
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFunctionResourceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *GetFunctionResourceResponseBody) GetLatency() *float64 {
	return s.Latency
}

func (s *GetFunctionResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFunctionResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionResourceResponseBody) GetResult() *GetFunctionResourceResponseBodyResult {
	return s.Result
}

func (s *GetFunctionResourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionResourceResponseBody) SetCode(v string) *GetFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetHttpCode(v int64) *GetFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetLatency(v float64) *GetFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetMessage(v string) *GetFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetRequestId(v string) *GetFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetResult(v *GetFunctionResourceResponseBodyResult) *GetFunctionResourceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionResourceResponseBody) SetStatus(v string) *GetFunctionResourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetFunctionResourceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionResourceResponseBodyResult struct {
	// The time when the resource was created. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource data. The data structure varies with the resource type.
	Data *GetFunctionResourceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// ""
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// rank
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The time when the resource was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The algorithm instances that are referenced.
	ReferencedInstances []*string `json:"ReferencedInstances,omitempty" xml:"ReferencedInstances,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// fg_json
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// raw_file
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetFunctionResourceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetFunctionResourceResponseBodyResult) GetData() *GetFunctionResourceResponseBodyResultData {
	return s.Data
}

func (s *GetFunctionResourceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetFunctionResourceResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionResourceResponseBodyResult) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetFunctionResourceResponseBodyResult) GetReferencedInstances() []*string {
	return s.ReferencedInstances
}

func (s *GetFunctionResourceResponseBodyResult) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetFunctionResourceResponseBodyResult) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetFunctionResourceResponseBodyResult) SetCreateTime(v int64) *GetFunctionResourceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetData(v *GetFunctionResourceResponseBodyResultData) *GetFunctionResourceResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetDescription(v string) *GetFunctionResourceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetFunctionName(v string) *GetFunctionResourceResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetModifyTime(v int64) *GetFunctionResourceResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetReferencedInstances(v []*string) *GetFunctionResourceResponseBodyResult {
	s.ReferencedInstances = v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetResourceName(v string) *GetFunctionResourceResponseBodyResult {
	s.ResourceName = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetResourceType(v string) *GetFunctionResourceResponseBodyResult {
	s.ResourceType = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionResourceResponseBodyResultData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// abc
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*GetFunctionResourceResponseBodyResultDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s GetFunctionResourceResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultData) GetContent() *string {
	return s.Content
}

func (s *GetFunctionResourceResponseBodyResultData) GetGenerators() []*GetFunctionResourceResponseBodyResultDataGenerators {
	return s.Generators
}

func (s *GetFunctionResourceResponseBodyResultData) SetContent(v string) *GetFunctionResourceResponseBodyResultData {
	s.Content = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultData) SetGenerators(v []*GetFunctionResourceResponseBodyResultDataGenerators) *GetFunctionResourceResponseBodyResultData {
	s.Generators = v
	return s
}

func (s *GetFunctionResourceResponseBodyResultData) Validate() error {
	if s.Generators != nil {
		for _, item := range s.Generators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFunctionResourceResponseBodyResultDataGenerators struct {
	// The type of the feature generator.
	//
	// example:
	//
	// id
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *GetFunctionResourceResponseBodyResultDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// output_feature_name
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s GetFunctionResourceResponseBodyResultDataGenerators) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultDataGenerators) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) GetGenerator() *string {
	return s.Generator
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) GetInput() *GetFunctionResourceResponseBodyResultDataGeneratorsInput {
	return s.Input
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) GetOutput() *string {
	return s.Output
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) SetGenerator(v string) *GetFunctionResourceResponseBodyResultDataGenerators {
	s.Generator = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) SetInput(v *GetFunctionResourceResponseBodyResultDataGeneratorsInput) *GetFunctionResourceResponseBodyResultDataGenerators {
	s.Input = v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) SetOutput(v string) *GetFunctionResourceResponseBodyResultDataGenerators {
	s.Output = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionResourceResponseBodyResultDataGeneratorsInput struct {
	// The input features.
	Features []*GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInput) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInput) GetFeatures() []*GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures {
	return s.Features
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInput) SetFeatures(v []*GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) *GetFunctionResourceResponseBodyResultDataGeneratorsInput {
	s.Features = v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInput) Validate() error {
	if s.Features != nil {
		for _, item := range s.Features {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) GetName() *string {
	return s.Name
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) GetType() *string {
	return s.Type
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) SetName(v string) *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) SetType(v string) *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) Validate() error {
	return dara.Validate(s)
}
