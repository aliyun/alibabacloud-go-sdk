// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionType(v string) *ListFunctionInstancesRequest
	GetFunctionType() *string
	SetModelType(v string) *ListFunctionInstancesRequest
	GetModelType() *string
	SetOutput(v string) *ListFunctionInstancesRequest
	GetOutput() *string
	SetPageNumber(v int32) *ListFunctionInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFunctionInstancesRequest
	GetPageSize() *int32
	SetSource(v string) *ListFunctionInstancesRequest
	GetSource() *string
}

type ListFunctionInstancesRequest struct {
	// The type of the feature.
	//
	// example:
	//
	// "PAAS"
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The level of detail for the returned information. Valid values:
	//
	// - normal: Displays information such as createParameters and cron. This is the default value.
	//
	// - simple: Displays only basic information.
	//
	// - detail: Returns the details of the training task.
	//
	// example:
	//
	// normal
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// The page number. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The source of the instance. Valid values:
	//
	// - builtin: The instance is created by the system.
	//
	// - user: The instance is created by the user. This is the default value.
	//
	// - all: All instances.
	//
	// example:
	//
	// user
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListFunctionInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ListFunctionInstancesRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListFunctionInstancesRequest) GetOutput() *string {
	return s.Output
}

func (s *ListFunctionInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFunctionInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFunctionInstancesRequest) GetSource() *string {
	return s.Source
}

func (s *ListFunctionInstancesRequest) SetFunctionType(v string) *ListFunctionInstancesRequest {
	s.FunctionType = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetModelType(v string) *ListFunctionInstancesRequest {
	s.ModelType = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetOutput(v string) *ListFunctionInstancesRequest {
	s.Output = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetPageNumber(v int32) *ListFunctionInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetPageSize(v int32) *ListFunctionInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetSource(v string) *ListFunctionInstancesRequest {
	s.Source = &v
	return s
}

func (s *ListFunctionInstancesRequest) Validate() error {
	return dara.Validate(s)
}
