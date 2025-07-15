// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourceStockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcpSpecId(v string) *CheckResourceStockRequest
	GetAcpSpecId() *string
	SetAmount(v int32) *CheckResourceStockRequest
	GetAmount() *int32
	SetBizRegionId(v string) *CheckResourceStockRequest
	GetBizRegionId() *string
	SetGpuAcceleration(v bool) *CheckResourceStockRequest
	GetGpuAcceleration() *bool
	SetZoneId(v string) *CheckResourceStockRequest
	GetZoneId() *string
}

type CheckResourceStockRequest struct {
	// Specification ID.
	//
	// example:
	//
	// acp.basic.small
	AcpSpecId *string `json:"AcpSpecId,omitempty" xml:"AcpSpecId,omitempty"`
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId     *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	GpuAcceleration *bool   `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The availability zone of the resource.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckResourceStockRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceStockRequest) GoString() string {
	return s.String()
}

func (s *CheckResourceStockRequest) GetAcpSpecId() *string {
	return s.AcpSpecId
}

func (s *CheckResourceStockRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CheckResourceStockRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CheckResourceStockRequest) GetGpuAcceleration() *bool {
	return s.GpuAcceleration
}

func (s *CheckResourceStockRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CheckResourceStockRequest) SetAcpSpecId(v string) *CheckResourceStockRequest {
	s.AcpSpecId = &v
	return s
}

func (s *CheckResourceStockRequest) SetAmount(v int32) *CheckResourceStockRequest {
	s.Amount = &v
	return s
}

func (s *CheckResourceStockRequest) SetBizRegionId(v string) *CheckResourceStockRequest {
	s.BizRegionId = &v
	return s
}

func (s *CheckResourceStockRequest) SetGpuAcceleration(v bool) *CheckResourceStockRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CheckResourceStockRequest) SetZoneId(v string) *CheckResourceStockRequest {
	s.ZoneId = &v
	return s
}

func (s *CheckResourceStockRequest) Validate() error {
	return dara.Validate(s)
}
