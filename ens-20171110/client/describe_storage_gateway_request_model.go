// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeStorageGatewayRequest
	GetEnsRegionId() *string
	SetGatewayId(v string) *DescribeStorageGatewayRequest
	GetGatewayId() *string
	SetGatewayType(v string) *DescribeStorageGatewayRequest
	GetGatewayType() *string
	SetPageNumber(v string) *DescribeStorageGatewayRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeStorageGatewayRequest
	GetPageSize() *string
	SetVpcId(v string) *DescribeStorageGatewayRequest
	GetVpcId() *string
}

type DescribeStorageGatewayRequest struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// sgw-****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The type of the gateway. Set this parameter to **1**. **1*	- indicates iSCSI.
	//
	// example:
	//
	// 1
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// n-***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeStorageGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageGatewayRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageGatewayRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeStorageGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DescribeStorageGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *DescribeStorageGatewayRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeStorageGatewayRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeStorageGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeStorageGatewayRequest) SetEnsRegionId(v string) *DescribeStorageGatewayRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeStorageGatewayRequest) SetGatewayId(v string) *DescribeStorageGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeStorageGatewayRequest) SetGatewayType(v string) *DescribeStorageGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *DescribeStorageGatewayRequest) SetPageNumber(v string) *DescribeStorageGatewayRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageGatewayRequest) SetPageSize(v string) *DescribeStorageGatewayRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageGatewayRequest) SetVpcId(v string) *DescribeStorageGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeStorageGatewayRequest) Validate() error {
	return dara.Validate(s)
}
