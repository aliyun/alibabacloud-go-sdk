// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefundPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *DescribeRefundPriceRequest
	GetDesktopId() []*string
	SetRefundType(v string) *DescribeRefundPriceRequest
	GetRefundType() *string
	SetRegionId(v string) *DescribeRefundPriceRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *DescribeRefundPriceRequest
	GetResellerOwnerUid() *int64
}

type DescribeRefundPriceRequest struct {
	// A list of cloud desktop IDs. You can specify one or more IDs. The number of IDs (N) must be between 1 and 20.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// Refund type.
	//
	// example:
	//
	// RemainRefund
	RefundType *string `json:"RefundType,omitempty" xml:"RefundType,omitempty"`
	// Region ID. Call [DescribeRegions](~~DescribeRegions~~) to get a list of regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s DescribeRefundPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefundPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefundPriceRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *DescribeRefundPriceRequest) GetRefundType() *string {
	return s.RefundType
}

func (s *DescribeRefundPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRefundPriceRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DescribeRefundPriceRequest) SetDesktopId(v []*string) *DescribeRefundPriceRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeRefundPriceRequest) SetRefundType(v string) *DescribeRefundPriceRequest {
	s.RefundType = &v
	return s
}

func (s *DescribeRefundPriceRequest) SetRegionId(v string) *DescribeRefundPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRefundPriceRequest) SetResellerOwnerUid(v int64) *DescribeRefundPriceRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DescribeRefundPriceRequest) Validate() error {
	return dara.Validate(s)
}
