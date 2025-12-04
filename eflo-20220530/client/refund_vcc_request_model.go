// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundVccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *RefundVccRequest
	GetRegionId() *string
	SetVccId(v string) *RefundVccRequest
	GetVccId() *string
}

type RefundVccRequest struct {
	// Region
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s RefundVccRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundVccRequest) GoString() string {
	return s.String()
}

func (s *RefundVccRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RefundVccRequest) GetVccId() *string {
	return s.VccId
}

func (s *RefundVccRequest) SetRegionId(v string) *RefundVccRequest {
	s.RegionId = &v
	return s
}

func (s *RefundVccRequest) SetVccId(v string) *RefundVccRequest {
	s.VccId = &v
	return s
}

func (s *RefundVccRequest) Validate() error {
	return dara.Validate(s)
}
