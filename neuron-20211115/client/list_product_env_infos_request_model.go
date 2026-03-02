// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductEnvInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListProductEnvInfosRequest
	GetCompanyId() *int64
	SetProductId(v int64) *ListProductEnvInfosRequest
	GetProductId() *int64
	SetType(v string) *ListProductEnvInfosRequest
	GetType() *string
}

type ListProductEnvInfosRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// SYSTEM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProductEnvInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductEnvInfosRequest) GoString() string {
	return s.String()
}

func (s *ListProductEnvInfosRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListProductEnvInfosRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListProductEnvInfosRequest) GetType() *string {
	return s.Type
}

func (s *ListProductEnvInfosRequest) SetCompanyId(v int64) *ListProductEnvInfosRequest {
	s.CompanyId = &v
	return s
}

func (s *ListProductEnvInfosRequest) SetProductId(v int64) *ListProductEnvInfosRequest {
	s.ProductId = &v
	return s
}

func (s *ListProductEnvInfosRequest) SetType(v string) *ListProductEnvInfosRequest {
	s.Type = &v
	return s
}

func (s *ListProductEnvInfosRequest) Validate() error {
	return dara.Validate(s)
}
