// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompanyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *CreateCompanyResponseBody
	GetCompanyId() *int64
	SetRequestId(v string) *CreateCompanyResponseBody
	GetRequestId() *string
}

type CreateCompanyResponseBody struct {
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
	// 838603C0-72A1-5070-A2E6-16E43861DB71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCompanyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCompanyResponseBody) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *CreateCompanyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCompanyResponseBody) SetCompanyId(v int64) *CreateCompanyResponseBody {
	s.CompanyId = &v
	return s
}

func (s *CreateCompanyResponseBody) SetRequestId(v string) *CreateCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCompanyResponseBody) Validate() error {
	return dara.Validate(s)
}
