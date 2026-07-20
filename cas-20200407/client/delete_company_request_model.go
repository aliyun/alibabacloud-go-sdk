// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompanyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *DeleteCompanyRequest
	GetCompanyId() *int64
}

type DeleteCompanyRequest struct {
	// The company ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51001
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
}

func (s DeleteCompanyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompanyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCompanyRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *DeleteCompanyRequest) SetCompanyId(v int64) *DeleteCompanyRequest {
	s.CompanyId = &v
	return s
}

func (s *DeleteCompanyRequest) Validate() error {
	return dara.Validate(s)
}
