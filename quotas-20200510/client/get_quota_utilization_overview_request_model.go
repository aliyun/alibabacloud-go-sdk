// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaUtilizationOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaCategory(v string) *GetQuotaUtilizationOverviewRequest
	GetQuotaCategory() *string
}

type GetQuotaUtilizationOverviewRequest struct {
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s GetQuotaUtilizationOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUtilizationOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaUtilizationOverviewRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *GetQuotaUtilizationOverviewRequest) SetQuotaCategory(v string) *GetQuotaUtilizationOverviewRequest {
	s.QuotaCategory = &v
	return s
}

func (s *GetQuotaUtilizationOverviewRequest) Validate() error {
	return dara.Validate(s)
}
