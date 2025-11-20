// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelType(v string) *ListModelServicesRequest
	GetModelType() *string
	SetPageNumber(v int32) *ListModelServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelServicesRequest
	GetPageSize() *int32
	SetProvider(v string) *ListModelServicesRequest
	GetProvider() *string
	SetProviderType(v string) *ListModelServicesRequest
	GetProviderType() *string
}

type ListModelServicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// system
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 10
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Aliyun
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// providerType
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
}

func (s ListModelServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesRequest) GoString() string {
	return s.String()
}

func (s *ListModelServicesRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelServicesRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListModelServicesRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListModelServicesRequest) SetModelType(v string) *ListModelServicesRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelServicesRequest) SetPageNumber(v int32) *ListModelServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelServicesRequest) SetPageSize(v int32) *ListModelServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelServicesRequest) SetProvider(v string) *ListModelServicesRequest {
	s.Provider = &v
	return s
}

func (s *ListModelServicesRequest) SetProviderType(v string) *ListModelServicesRequest {
	s.ProviderType = &v
	return s
}

func (s *ListModelServicesRequest) Validate() error {
	return dara.Validate(s)
}
