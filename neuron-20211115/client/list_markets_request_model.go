// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListMarketsRequest
	GetCompanyId() *int64
}

type ListMarketsRequest struct {
	// example:
	//
	// 41
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
}

func (s ListMarketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMarketsRequest) GoString() string {
	return s.String()
}

func (s *ListMarketsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListMarketsRequest) SetCompanyId(v int64) *ListMarketsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListMarketsRequest) Validate() error {
	return dara.Validate(s)
}
