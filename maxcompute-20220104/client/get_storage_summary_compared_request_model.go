// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageSummaryComparedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetStorageSummaryComparedRequest
	GetBeginDate() *string
	SetEndDate(v string) *GetStorageSummaryComparedRequest
	GetEndDate() *string
	SetProjects(v []*string) *GetStorageSummaryComparedRequest
	GetProjects() []*string
	SetRegion(v string) *GetStorageSummaryComparedRequest
	GetRegion() *string
	SetTenantId(v string) *GetStorageSummaryComparedRequest
	GetTenantId() *string
}

type GetStorageSummaryComparedRequest struct {
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
	EndDate  *string   `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Projects []*string `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 483212237127906
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetStorageSummaryComparedRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSummaryComparedRequest) GoString() string {
	return s.String()
}

func (s *GetStorageSummaryComparedRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetStorageSummaryComparedRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetStorageSummaryComparedRequest) GetProjects() []*string {
	return s.Projects
}

func (s *GetStorageSummaryComparedRequest) GetRegion() *string {
	return s.Region
}

func (s *GetStorageSummaryComparedRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetStorageSummaryComparedRequest) SetBeginDate(v string) *GetStorageSummaryComparedRequest {
	s.BeginDate = &v
	return s
}

func (s *GetStorageSummaryComparedRequest) SetEndDate(v string) *GetStorageSummaryComparedRequest {
	s.EndDate = &v
	return s
}

func (s *GetStorageSummaryComparedRequest) SetProjects(v []*string) *GetStorageSummaryComparedRequest {
	s.Projects = v
	return s
}

func (s *GetStorageSummaryComparedRequest) SetRegion(v string) *GetStorageSummaryComparedRequest {
	s.Region = &v
	return s
}

func (s *GetStorageSummaryComparedRequest) SetTenantId(v string) *GetStorageSummaryComparedRequest {
	s.TenantId = &v
	return s
}

func (s *GetStorageSummaryComparedRequest) Validate() error {
	return dara.Validate(s)
}
