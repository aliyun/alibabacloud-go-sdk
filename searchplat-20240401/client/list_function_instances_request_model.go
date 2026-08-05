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
	// The feature type.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The model type.
	//
	// example:
	//
	// ai_search
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The level of detail in the response. Valid values:
	//
	// - simple: displays only basic information
	//
	// - normal: displays information such as createParameters and cron (default)
	//
	// - detail: returns training task information.
	//
	// example:
	//
	// simple
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
