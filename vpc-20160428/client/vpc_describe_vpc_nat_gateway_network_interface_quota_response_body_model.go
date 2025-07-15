// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody
	GetQuota() *int64
	SetRequestId(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody
	GetRequestId() *string
}

type VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody struct {
	// The number of endpoints that can be created.
	//
	// example:
	//
	// 2
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 611CB80C-B6A9-43DB-9E38-0B0AC3D9B58F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) GetQuota() *int64 {
	return s.Quota
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) SetQuota(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody {
	s.Quota = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) SetRequestId(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
