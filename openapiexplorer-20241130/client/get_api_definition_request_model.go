// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *GetApiDefinitionRequest
	GetApi() *string
	SetApiVersion(v string) *GetApiDefinitionRequest
	GetApiVersion() *string
	SetProduct(v string) *GetApiDefinitionRequest
	GetProduct() *string
}

type GetApiDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DescribeRegions
	Api *string `json:"api,omitempty" xml:"api,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetApiDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetApiDefinitionRequest) GetApi() *string {
	return s.Api
}

func (s *GetApiDefinitionRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetApiDefinitionRequest) GetProduct() *string {
	return s.Product
}

func (s *GetApiDefinitionRequest) SetApi(v string) *GetApiDefinitionRequest {
	s.Api = &v
	return s
}

func (s *GetApiDefinitionRequest) SetApiVersion(v string) *GetApiDefinitionRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetApiDefinitionRequest) SetProduct(v string) *GetApiDefinitionRequest {
	s.Product = &v
	return s
}

func (s *GetApiDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
