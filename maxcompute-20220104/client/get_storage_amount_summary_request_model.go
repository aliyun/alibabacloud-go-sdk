// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageAmountSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetStorageAmountSummaryRequest
	GetDate() *string
	SetRegion(v string) *GetStorageAmountSummaryRequest
	GetRegion() *string
	SetTenantId(v string) *GetStorageAmountSummaryRequest
	GetTenantId() *string
}

type GetStorageAmountSummaryRequest struct {
	// The date for statistics collection. The value is at the day level and must be in the YYYYMMdd format.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The tenant ID. You can view the tenant ID by logging on to the MaxCompute console and choosing **Tenant Management*	- > **Tenant Properties*	- in the left-side navigation pane.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetStorageAmountSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAmountSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetStorageAmountSummaryRequest) GetDate() *string {
	return s.Date
}

func (s *GetStorageAmountSummaryRequest) GetRegion() *string {
	return s.Region
}

func (s *GetStorageAmountSummaryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetStorageAmountSummaryRequest) SetDate(v string) *GetStorageAmountSummaryRequest {
	s.Date = &v
	return s
}

func (s *GetStorageAmountSummaryRequest) SetRegion(v string) *GetStorageAmountSummaryRequest {
	s.Region = &v
	return s
}

func (s *GetStorageAmountSummaryRequest) SetTenantId(v string) *GetStorageAmountSummaryRequest {
	s.TenantId = &v
	return s
}

func (s *GetStorageAmountSummaryRequest) Validate() error {
	return dara.Validate(s)
}
