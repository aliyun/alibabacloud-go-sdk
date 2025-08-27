// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlipayNo(v string) *CostCenterModifyRequest
	GetAlipayNo() *string
	SetDisable(v int64) *CostCenterModifyRequest
	GetDisable() *int64
	SetNumber(v string) *CostCenterModifyRequest
	GetNumber() *string
	SetScope(v int64) *CostCenterModifyRequest
	GetScope() *int64
	SetThirdpartId(v string) *CostCenterModifyRequest
	GetThirdpartId() *string
	SetTitle(v string) *CostCenterModifyRequest
	GetTitle() *string
}

type CostCenterModifyRequest struct {
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

func (s CostCenterModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s CostCenterModifyRequest) GoString() string {
	return s.String()
}

func (s *CostCenterModifyRequest) GetAlipayNo() *string {
	return s.AlipayNo
}

func (s *CostCenterModifyRequest) GetDisable() *int64 {
	return s.Disable
}

func (s *CostCenterModifyRequest) GetNumber() *string {
	return s.Number
}

func (s *CostCenterModifyRequest) GetScope() *int64 {
	return s.Scope
}

func (s *CostCenterModifyRequest) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *CostCenterModifyRequest) GetTitle() *string {
	return s.Title
}

func (s *CostCenterModifyRequest) SetAlipayNo(v string) *CostCenterModifyRequest {
	s.AlipayNo = &v
	return s
}

func (s *CostCenterModifyRequest) SetDisable(v int64) *CostCenterModifyRequest {
	s.Disable = &v
	return s
}

func (s *CostCenterModifyRequest) SetNumber(v string) *CostCenterModifyRequest {
	s.Number = &v
	return s
}

func (s *CostCenterModifyRequest) SetScope(v int64) *CostCenterModifyRequest {
	s.Scope = &v
	return s
}

func (s *CostCenterModifyRequest) SetThirdpartId(v string) *CostCenterModifyRequest {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterModifyRequest) SetTitle(v string) *CostCenterModifyRequest {
	s.Title = &v
	return s
}

func (s *CostCenterModifyRequest) Validate() error {
	return dara.Validate(s)
}
