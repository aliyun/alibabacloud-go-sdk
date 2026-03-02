// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *DeleteProductRequest
	GetCompanyId() *int64
}

type DeleteProductRequest struct {
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
}

func (s DeleteProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *DeleteProductRequest) SetCompanyId(v int64) *DeleteProductRequest {
	s.CompanyId = &v
	return s
}

func (s *DeleteProductRequest) Validate() error {
	return dara.Validate(s)
}
