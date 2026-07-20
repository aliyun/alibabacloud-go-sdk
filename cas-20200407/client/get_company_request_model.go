// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompanyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *GetCompanyRequest
	GetCompanyId() *int64
}

type GetCompanyRequest struct {
	// The company ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51001
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
}

func (s GetCompanyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompanyRequest) GoString() string {
	return s.String()
}

func (s *GetCompanyRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *GetCompanyRequest) SetCompanyId(v int64) *GetCompanyRequest {
	s.CompanyId = &v
	return s
}

func (s *GetCompanyRequest) Validate() error {
	return dara.Validate(s)
}
