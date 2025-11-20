// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelName(v string) *ListModelProvidersRequest
	GetModelName() *string
	SetModelType(v string) *ListModelProvidersRequest
	GetModelType() *string
	SetPageNumber(v string) *ListModelProvidersRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListModelProvidersRequest
	GetPageSize() *string
	SetProvider(v string) *ListModelProvidersRequest
	GetProvider() *string
}

type ListModelProvidersRequest struct {
	// example:
	//
	// aa
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// pop
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Aliyun
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListModelProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListModelProvidersRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelProvidersRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelProvidersRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListModelProvidersRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListModelProvidersRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListModelProvidersRequest) SetModelName(v string) *ListModelProvidersRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelProvidersRequest) SetModelType(v string) *ListModelProvidersRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelProvidersRequest) SetPageNumber(v string) *ListModelProvidersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelProvidersRequest) SetPageSize(v string) *ListModelProvidersRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelProvidersRequest) SetProvider(v string) *ListModelProvidersRequest {
	s.Provider = &v
	return s
}

func (s *ListModelProvidersRequest) Validate() error {
	return dara.Validate(s)
}
