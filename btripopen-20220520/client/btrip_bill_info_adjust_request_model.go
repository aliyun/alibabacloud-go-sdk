// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBtripBillInfoAdjustRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrimaryId(v int64) *BtripBillInfoAdjustRequest
	GetPrimaryId() *int64
	SetThirdPartCostCenterId(v string) *BtripBillInfoAdjustRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartDepartmentId(v string) *BtripBillInfoAdjustRequest
	GetThirdPartDepartmentId() *string
	SetThirdPartInvoiceId(v string) *BtripBillInfoAdjustRequest
	GetThirdPartInvoiceId() *string
	SetThirdPartProjectId(v string) *BtripBillInfoAdjustRequest
	GetThirdPartProjectId() *string
	SetUserId(v string) *BtripBillInfoAdjustRequest
	GetUserId() *string
}

type BtripBillInfoAdjustRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	PrimaryId *int64 `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	// example:
	//
	// GA15131
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// example:
	//
	// GA15131
	ThirdPartDepartmentId *string `json:"third_part_department_id,omitempty" xml:"third_part_department_id,omitempty"`
	// example:
	//
	// GA15131
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// example:
	//
	// GA15131
	ThirdPartProjectId *string `json:"third_part_project_id,omitempty" xml:"third_part_project_id,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s BtripBillInfoAdjustRequest) String() string {
	return dara.Prettify(s)
}

func (s BtripBillInfoAdjustRequest) GoString() string {
	return s.String()
}

func (s *BtripBillInfoAdjustRequest) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *BtripBillInfoAdjustRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *BtripBillInfoAdjustRequest) GetThirdPartDepartmentId() *string {
	return s.ThirdPartDepartmentId
}

func (s *BtripBillInfoAdjustRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *BtripBillInfoAdjustRequest) GetThirdPartProjectId() *string {
	return s.ThirdPartProjectId
}

func (s *BtripBillInfoAdjustRequest) GetUserId() *string {
	return s.UserId
}

func (s *BtripBillInfoAdjustRequest) SetPrimaryId(v int64) *BtripBillInfoAdjustRequest {
	s.PrimaryId = &v
	return s
}

func (s *BtripBillInfoAdjustRequest) SetThirdPartCostCenterId(v string) *BtripBillInfoAdjustRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *BtripBillInfoAdjustRequest) SetThirdPartDepartmentId(v string) *BtripBillInfoAdjustRequest {
	s.ThirdPartDepartmentId = &v
	return s
}

func (s *BtripBillInfoAdjustRequest) SetThirdPartInvoiceId(v string) *BtripBillInfoAdjustRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *BtripBillInfoAdjustRequest) SetThirdPartProjectId(v string) *BtripBillInfoAdjustRequest {
	s.ThirdPartProjectId = &v
	return s
}

func (s *BtripBillInfoAdjustRequest) SetUserId(v string) *BtripBillInfoAdjustRequest {
	s.UserId = &v
	return s
}

func (s *BtripBillInfoAdjustRequest) Validate() error {
	return dara.Validate(s)
}
