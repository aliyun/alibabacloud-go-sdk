// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlipayNo(v string) *CostCenterSaveRequest
	GetAlipayNo() *string
	SetDisable(v int64) *CostCenterSaveRequest
	GetDisable() *int64
	SetNumber(v string) *CostCenterSaveRequest
	GetNumber() *string
	SetScope(v int64) *CostCenterSaveRequest
	GetScope() *int64
	SetThirdpartId(v string) *CostCenterSaveRequest
	GetThirdpartId() *string
	SetTitle(v string) *CostCenterSaveRequest
	GetTitle() *string
}

type CostCenterSaveRequest struct {
	// example:
	//
	// a@alipay.com
	AlipayNo *string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	Disable  *int64  `json:"disable,omitempty" xml:"disable,omitempty"`
	// example:
	//
	// 12345
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Scope *int64 `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CostCenterSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s CostCenterSaveRequest) GoString() string {
	return s.String()
}

func (s *CostCenterSaveRequest) GetAlipayNo() *string {
	return s.AlipayNo
}

func (s *CostCenterSaveRequest) GetDisable() *int64 {
	return s.Disable
}

func (s *CostCenterSaveRequest) GetNumber() *string {
	return s.Number
}

func (s *CostCenterSaveRequest) GetScope() *int64 {
	return s.Scope
}

func (s *CostCenterSaveRequest) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *CostCenterSaveRequest) GetTitle() *string {
	return s.Title
}

func (s *CostCenterSaveRequest) SetAlipayNo(v string) *CostCenterSaveRequest {
	s.AlipayNo = &v
	return s
}

func (s *CostCenterSaveRequest) SetDisable(v int64) *CostCenterSaveRequest {
	s.Disable = &v
	return s
}

func (s *CostCenterSaveRequest) SetNumber(v string) *CostCenterSaveRequest {
	s.Number = &v
	return s
}

func (s *CostCenterSaveRequest) SetScope(v int64) *CostCenterSaveRequest {
	s.Scope = &v
	return s
}

func (s *CostCenterSaveRequest) SetThirdpartId(v string) *CostCenterSaveRequest {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterSaveRequest) SetTitle(v string) *CostCenterSaveRequest {
	s.Title = &v
	return s
}

func (s *CostCenterSaveRequest) Validate() error {
	return dara.Validate(s)
}
