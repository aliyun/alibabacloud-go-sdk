// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageSizeSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetStorageSizeSummaryRequest
	GetDate() *string
	SetRegion(v string) *GetStorageSizeSummaryRequest
	GetRegion() *string
	SetTenantId(v string) *GetStorageSizeSummaryRequest
	GetTenantId() *string
}

type GetStorageSizeSummaryRequest struct {
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 483212237127906
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetStorageSizeSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSizeSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetStorageSizeSummaryRequest) GetDate() *string {
	return s.Date
}

func (s *GetStorageSizeSummaryRequest) GetRegion() *string {
	return s.Region
}

func (s *GetStorageSizeSummaryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetStorageSizeSummaryRequest) SetDate(v string) *GetStorageSizeSummaryRequest {
	s.Date = &v
	return s
}

func (s *GetStorageSizeSummaryRequest) SetRegion(v string) *GetStorageSizeSummaryRequest {
	s.Region = &v
	return s
}

func (s *GetStorageSizeSummaryRequest) SetTenantId(v string) *GetStorageSizeSummaryRequest {
	s.TenantId = &v
	return s
}

func (s *GetStorageSizeSummaryRequest) Validate() error {
	return dara.Validate(s)
}
