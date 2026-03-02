// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeveloperPbcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListDeveloperPbcsRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *ListDeveloperPbcsRequest
	GetMarketId() *int64
}

type ListDeveloperPbcsRequest struct {
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId  *int64 `json:"marketId,omitempty" xml:"marketId,omitempty"`
}

func (s ListDeveloperPbcsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeveloperPbcsRequest) GoString() string {
	return s.String()
}

func (s *ListDeveloperPbcsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListDeveloperPbcsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListDeveloperPbcsRequest) SetCompanyId(v int64) *ListDeveloperPbcsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListDeveloperPbcsRequest) SetMarketId(v int64) *ListDeveloperPbcsRequest {
	s.MarketId = &v
	return s
}

func (s *ListDeveloperPbcsRequest) Validate() error {
	return dara.Validate(s)
}
