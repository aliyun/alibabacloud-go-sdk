// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompanyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *UpdateCompanyResponseBody
	GetCompanyId() *int64
	SetRequestId(v string) *UpdateCompanyResponseBody
	GetRequestId() *string
}

type UpdateCompanyResponseBody struct {
	// The company ID.
	//
	// example:
	//
	// 51001
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C2BAAB19-BCEB-569B-BE08-7C728344A79C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCompanyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCompanyResponseBody) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *UpdateCompanyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCompanyResponseBody) SetCompanyId(v int64) *UpdateCompanyResponseBody {
	s.CompanyId = &v
	return s
}

func (s *UpdateCompanyResponseBody) SetRequestId(v string) *UpdateCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCompanyResponseBody) Validate() error {
	return dara.Validate(s)
}
