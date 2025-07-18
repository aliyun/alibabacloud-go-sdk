// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeAvailableResourcesRequest
	GetChargeType() *string
	SetRegion(v string) *DescribeAvailableResourcesRequest
	GetRegion() *string
	SetZoneId(v string) *DescribeAvailableResourcesRequest
	GetZoneId() *string
}

type DescribeAvailableResourcesRequest struct {
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeAvailableResourcesRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeAvailableResourcesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourcesRequest) SetChargeType(v string) *DescribeAvailableResourcesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetRegion(v string) *DescribeAvailableResourcesRequest {
	s.Region = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetZoneId(v string) *DescribeAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) Validate() error {
	return dara.Validate(s)
}
