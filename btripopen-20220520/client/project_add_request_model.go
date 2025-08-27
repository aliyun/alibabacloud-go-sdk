// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProjectAddRequest
	GetCode() *string
	SetProjectName(v string) *ProjectAddRequest
	GetProjectName() *string
	SetThirdPartCostCenterId(v string) *ProjectAddRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartId(v string) *ProjectAddRequest
	GetThirdPartId() *string
	SetThirdPartInvoiceId(v string) *ProjectAddRequest
	GetThirdPartInvoiceId() *string
}

type ProjectAddRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pro_code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 234
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	// example:
	//
	// 123
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
}

func (s ProjectAddRequest) String() string {
	return dara.Prettify(s)
}

func (s ProjectAddRequest) GoString() string {
	return s.String()
}

func (s *ProjectAddRequest) GetCode() *string {
	return s.Code
}

func (s *ProjectAddRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ProjectAddRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *ProjectAddRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *ProjectAddRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ProjectAddRequest) SetCode(v string) *ProjectAddRequest {
	s.Code = &v
	return s
}

func (s *ProjectAddRequest) SetProjectName(v string) *ProjectAddRequest {
	s.ProjectName = &v
	return s
}

func (s *ProjectAddRequest) SetThirdPartCostCenterId(v string) *ProjectAddRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *ProjectAddRequest) SetThirdPartId(v string) *ProjectAddRequest {
	s.ThirdPartId = &v
	return s
}

func (s *ProjectAddRequest) SetThirdPartInvoiceId(v string) *ProjectAddRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ProjectAddRequest) Validate() error {
	return dara.Validate(s)
}
