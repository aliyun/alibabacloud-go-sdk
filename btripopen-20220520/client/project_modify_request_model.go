// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProjectModifyRequest
	GetCode() *string
	SetHasManager(v bool) *ProjectModifyRequest
	GetHasManager() *bool
	SetProjectName(v string) *ProjectModifyRequest
	GetProjectName() *string
	SetThirdPartCostCenterId(v string) *ProjectModifyRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartId(v string) *ProjectModifyRequest
	GetThirdPartId() *string
	SetThirdPartInvoiceId(v string) *ProjectModifyRequest
	GetThirdPartInvoiceId() *string
}

type ProjectModifyRequest struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	HasManager *bool   `json:"has_manager,omitempty" xml:"has_manager,omitempty"`
	// This parameter is required.
	ProjectName           *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// This parameter is required.
	ThirdPartId        *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
}

func (s ProjectModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s ProjectModifyRequest) GoString() string {
	return s.String()
}

func (s *ProjectModifyRequest) GetCode() *string {
	return s.Code
}

func (s *ProjectModifyRequest) GetHasManager() *bool {
	return s.HasManager
}

func (s *ProjectModifyRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ProjectModifyRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *ProjectModifyRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *ProjectModifyRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ProjectModifyRequest) SetCode(v string) *ProjectModifyRequest {
	s.Code = &v
	return s
}

func (s *ProjectModifyRequest) SetHasManager(v bool) *ProjectModifyRequest {
	s.HasManager = &v
	return s
}

func (s *ProjectModifyRequest) SetProjectName(v string) *ProjectModifyRequest {
	s.ProjectName = &v
	return s
}

func (s *ProjectModifyRequest) SetThirdPartCostCenterId(v string) *ProjectModifyRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *ProjectModifyRequest) SetThirdPartId(v string) *ProjectModifyRequest {
	s.ThirdPartId = &v
	return s
}

func (s *ProjectModifyRequest) SetThirdPartInvoiceId(v string) *ProjectModifyRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ProjectModifyRequest) Validate() error {
	return dara.Validate(s)
}
