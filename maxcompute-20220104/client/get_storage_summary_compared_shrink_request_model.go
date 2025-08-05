// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageSummaryComparedShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetStorageSummaryComparedShrinkRequest
	GetBeginDate() *string
	SetEndDate(v string) *GetStorageSummaryComparedShrinkRequest
	GetEndDate() *string
	SetProjectsShrink(v string) *GetStorageSummaryComparedShrinkRequest
	GetProjectsShrink() *string
	SetRegion(v string) *GetStorageSummaryComparedShrinkRequest
	GetRegion() *string
	SetTenantId(v string) *GetStorageSummaryComparedShrinkRequest
	GetTenantId() *string
}

type GetStorageSummaryComparedShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20250601
	BeginDate *string `json:"beginDate,omitempty" xml:"beginDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20250604
	EndDate        *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ProjectsShrink *string `json:"projects,omitempty" xml:"projects,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 483212237127906
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetStorageSummaryComparedShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSummaryComparedShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetStorageSummaryComparedShrinkRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetStorageSummaryComparedShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetStorageSummaryComparedShrinkRequest) GetProjectsShrink() *string {
	return s.ProjectsShrink
}

func (s *GetStorageSummaryComparedShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *GetStorageSummaryComparedShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetStorageSummaryComparedShrinkRequest) SetBeginDate(v string) *GetStorageSummaryComparedShrinkRequest {
	s.BeginDate = &v
	return s
}

func (s *GetStorageSummaryComparedShrinkRequest) SetEndDate(v string) *GetStorageSummaryComparedShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetStorageSummaryComparedShrinkRequest) SetProjectsShrink(v string) *GetStorageSummaryComparedShrinkRequest {
	s.ProjectsShrink = &v
	return s
}

func (s *GetStorageSummaryComparedShrinkRequest) SetRegion(v string) *GetStorageSummaryComparedShrinkRequest {
	s.Region = &v
	return s
}

func (s *GetStorageSummaryComparedShrinkRequest) SetTenantId(v string) *GetStorageSummaryComparedShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GetStorageSummaryComparedShrinkRequest) Validate() error {
	return dara.Validate(s)
}
