// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductEnvsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListProductEnvsRequest
	GetCompanyId() *int64
	SetProductId(v int64) *ListProductEnvsRequest
	GetProductId() *int64
}

type ListProductEnvsRequest struct {
	// example:
	//
	// 42
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// 12
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
}

func (s ListProductEnvsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductEnvsRequest) GoString() string {
	return s.String()
}

func (s *ListProductEnvsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListProductEnvsRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListProductEnvsRequest) SetCompanyId(v int64) *ListProductEnvsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListProductEnvsRequest) SetProductId(v int64) *ListProductEnvsRequest {
	s.ProductId = &v
	return s
}

func (s *ListProductEnvsRequest) Validate() error {
	return dara.Validate(s)
}
