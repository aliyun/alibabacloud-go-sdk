// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFunctionResourcesResponseBody
	GetCode() *string
	SetHttpCode(v int64) *ListFunctionResourcesResponseBody
	GetHttpCode() *int64
	SetLatency(v float64) *ListFunctionResourcesResponseBody
	GetLatency() *float64
	SetMessage(v string) *ListFunctionResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFunctionResourcesResponseBody
	GetRequestId() *string
	SetResult(v []*ListFunctionResourcesResponseBodyResult) *ListFunctionResourcesResponseBody
	GetResult() []*ListFunctionResourcesResponseBodyResult
	SetStatus(v string) *ListFunctionResourcesResponseBody
	GetStatus() *string
	SetTotalCount(v int64) *ListFunctionResourcesResponseBody
	GetTotalCount() *int64
}

type ListFunctionResourcesResponseBody struct {
	// The error code returned. If no error occurs, this value is empty.
	//
	// example:
	//
	// Resource.InvalidResourceName
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The amount of time consumed for the request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Invalid resource name.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// "3A809095-C554-5CF5-8FCE-BE19D4673790"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results returned.
	Result []*ListFunctionResourcesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request. Valid values: OK and FAIL.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of records that meet the requirements.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFunctionResourcesResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ListFunctionResourcesResponseBody) GetLatency() *float64 {
	return s.Latency
}

func (s *ListFunctionResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFunctionResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunctionResourcesResponseBody) GetResult() []*ListFunctionResourcesResponseBodyResult {
	return s.Result
}

func (s *ListFunctionResourcesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListFunctionResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFunctionResourcesResponseBody) SetCode(v string) *ListFunctionResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetHttpCode(v int64) *ListFunctionResourcesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetLatency(v float64) *ListFunctionResourcesResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetMessage(v string) *ListFunctionResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetRequestId(v string) *ListFunctionResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetResult(v []*ListFunctionResourcesResponseBodyResult) *ListFunctionResourcesResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetStatus(v string) *ListFunctionResourcesResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetTotalCount(v int64) *ListFunctionResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) Validate() error {
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

type ListFunctionResourcesResponseBodyResult struct {
	// The time when the resource was created. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource data. The data structure varies with the resource type.
	Data *ListFunctionResourcesResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// resource description
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
	// feature_generator
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListFunctionResourcesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListFunctionResourcesResponseBodyResult) GetData() *ListFunctionResourcesResponseBodyResultData {
	return s.Data
}

func (s *ListFunctionResourcesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListFunctionResourcesResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListFunctionResourcesResponseBodyResult) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListFunctionResourcesResponseBodyResult) GetReferencedInstances() []*string {
	return s.ReferencedInstances
}

func (s *ListFunctionResourcesResponseBodyResult) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListFunctionResourcesResponseBodyResult) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListFunctionResourcesResponseBodyResult) SetCreateTime(v int64) *ListFunctionResourcesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetData(v *ListFunctionResourcesResponseBodyResultData) *ListFunctionResourcesResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetDescription(v string) *ListFunctionResourcesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetFunctionName(v string) *ListFunctionResourcesResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetModifyTime(v int64) *ListFunctionResourcesResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetReferencedInstances(v []*string) *ListFunctionResourcesResponseBodyResult {
	s.ReferencedInstances = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetResourceName(v string) *ListFunctionResourcesResponseBodyResult {
	s.ResourceName = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetResourceType(v string) *ListFunctionResourcesResponseBodyResult {
	s.ResourceType = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFunctionResourcesResponseBodyResultData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// "abc"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*ListFunctionResourcesResponseBodyResultDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s ListFunctionResourcesResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultData) GetContent() *string {
	return s.Content
}

func (s *ListFunctionResourcesResponseBodyResultData) GetGenerators() []*ListFunctionResourcesResponseBodyResultDataGenerators {
	return s.Generators
}

func (s *ListFunctionResourcesResponseBodyResultData) SetContent(v string) *ListFunctionResourcesResponseBodyResultData {
	s.Content = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultData) SetGenerators(v []*ListFunctionResourcesResponseBodyResultDataGenerators) *ListFunctionResourcesResponseBodyResultData {
	s.Generators = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultData) Validate() error {
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

type ListFunctionResourcesResponseBodyResultDataGenerators struct {
	// The type of the feature generator.
	//
	// example:
	//
	// combo
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *ListFunctionResourcesResponseBodyResultDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// feature1
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListFunctionResourcesResponseBodyResultDataGenerators) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultDataGenerators) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) GetGenerator() *string {
	return s.Generator
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) GetInput() *ListFunctionResourcesResponseBodyResultDataGeneratorsInput {
	return s.Input
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) GetOutput() *string {
	return s.Output
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) SetGenerator(v string) *ListFunctionResourcesResponseBodyResultDataGenerators {
	s.Generator = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) SetInput(v *ListFunctionResourcesResponseBodyResultDataGeneratorsInput) *ListFunctionResourcesResponseBodyResultDataGenerators {
	s.Input = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) SetOutput(v string) *ListFunctionResourcesResponseBodyResultDataGenerators {
	s.Output = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFunctionResourcesResponseBodyResultDataGeneratorsInput struct {
	// The input features.
	Features []*ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInput) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInput) GetFeatures() []*ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures {
	return s.Features
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInput) SetFeatures(v []*ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) *ListFunctionResourcesResponseBodyResultDataGeneratorsInput {
	s.Features = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInput) Validate() error {
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

type ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// Valid values:
	//
	// 	- item
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- user
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) GetName() *string {
	return s.Name
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) GetType() *string {
	return s.Type
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) SetName(v string) *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) SetType(v string) *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) Validate() error {
	return dara.Validate(s)
}
