// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeGrantRulesToCenRequest
	GetCenId() *string
	SetChildInstanceId(v string) *DescribeGrantRulesToCenRequest
	GetChildInstanceId() *string
	SetChildInstanceOwnerId(v int64) *DescribeGrantRulesToCenRequest
	GetChildInstanceOwnerId() *int64
	SetEnabledIpv6(v bool) *DescribeGrantRulesToCenRequest
	GetEnabledIpv6() *bool
	SetMaxResults(v int64) *DescribeGrantRulesToCenRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeGrantRulesToCenRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeGrantRulesToCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGrantRulesToCenRequest
	GetOwnerId() *int64
	SetProductType(v string) *DescribeGrantRulesToCenRequest
	GetProductType() *string
	SetRegionId(v string) *DescribeGrantRulesToCenRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeGrantRulesToCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGrantRulesToCenRequest
	GetResourceOwnerId() *int64
}

type DescribeGrantRulesToCenRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-nye53d7p3hzyu4****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the network instance to query.
	//
	// example:
	//
	// vpc-rj9gt5nll27onu7****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// 1250123456123456
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// Specifies whether IPv6 is enabled.
	//
	// 1. This parameter applies only when **ProductType*	- is set to **VPC**.
	//
	// 2. Set to **true*	- to return only VPCs with IPv6 enabled, or **false*	- to return only those with IPv6 disabled. If you omit this parameter, the response is not filtered by the IPv6 status.
	//
	// example:
	//
	// true
	EnabledIpv6 *bool `json:"EnabledIpv6,omitempty" xml:"EnabledIpv6,omitempty"`
	// - If you omit this parameter, all entries are returned. The **MaxResults*	- value in the response then indicates the total number of entries.
	//
	// - If you set this parameter, the response is paginated. The **MaxResults*	- value specifies the number of entries to return on each page. Valid values: **1*	- to **100**. The **MaxResults*	- value in the response indicates the number of entries on the current page. Setting **MaxResults*	- to **20*	- is recommended.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results.
	//
	// - Omit this parameter for the first request.
	//
	// - Set this parameter to the **NextToken*	- value from the previous response to retrieve the next page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// - **VPC**: a virtual private cloud (VPC).
	//
	// - **VBR**: a virtual border router (VBR).
	//
	// - **CCN**: a Cloud Connect Network (CCN) instance.
	//
	// - **VPN**: an IPsec-VPN connection.
	//
	// - **ECR**: an Express Connect router.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the region where the network instance is located.
	//
	// You can call the [](t2264556.xdita#)operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeGrantRulesToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeGrantRulesToCenRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeGrantRulesToCenRequest) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *DescribeGrantRulesToCenRequest) GetEnabledIpv6() *bool {
	return s.EnabledIpv6
}

func (s *DescribeGrantRulesToCenRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeGrantRulesToCenRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGrantRulesToCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGrantRulesToCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGrantRulesToCenRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeGrantRulesToCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantRulesToCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGrantRulesToCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGrantRulesToCenRequest) SetCenId(v string) *DescribeGrantRulesToCenRequest {
	s.CenId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetChildInstanceId(v string) *DescribeGrantRulesToCenRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetChildInstanceOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetEnabledIpv6(v bool) *DescribeGrantRulesToCenRequest {
	s.EnabledIpv6 = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetMaxResults(v int64) *DescribeGrantRulesToCenRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetNextToken(v string) *DescribeGrantRulesToCenRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerAccount(v string) *DescribeGrantRulesToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetProductType(v string) *DescribeGrantRulesToCenRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetRegionId(v string) *DescribeGrantRulesToCenRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetResourceOwnerAccount(v string) *DescribeGrantRulesToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetResourceOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) Validate() error {
	return dara.Validate(s)
}
