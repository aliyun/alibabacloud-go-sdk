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
	// ID of cloud computer N. Valid values of N: 1 to 20.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The unsubscription type.
	//
	// Valid values:
	//
	// 	- RemainRefund: refunds the remaining balance and releases resources.
	//
	// 	- RenewRefund: refunds only the renewal fee and adjusts the expiration date accordingly.
	//
	// example:
	//
	// RemainRefund
	RefundType *string `json:"RefundType,omitempty" xml:"RefundType,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
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
