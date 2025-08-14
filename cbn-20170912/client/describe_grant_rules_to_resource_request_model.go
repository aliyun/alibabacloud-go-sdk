// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeGrantRulesToResourceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeGrantRulesToResourceRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeGrantRulesToResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGrantRulesToResourceRequest
	GetOwnerId() *int64
	SetProductType(v string) *DescribeGrantRulesToResourceRequest
	GetProductType() *string
	SetRegionId(v string) *DescribeGrantRulesToResourceRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeGrantRulesToResourceRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeGrantRulesToResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGrantRulesToResourceRequest
	GetResourceOwnerId() *int64
}

type DescribeGrantRulesToResourceRequest struct {
	// 	- If you do not specify a value for **MaxResults**, entries are returned in one response. After you send the request, the value of **MaxResults*	- includes all entries.
	//
	// 	- If you specify a value for **MaxResults**, entries are returned in batches. The value of **MaxResults*	- indicates the total number of entries returned per batch. Valid values: **1*	- to **100**. After you send the request, the value of **MaxResults*	- indicates the number of entries returned in the current response. We recommend that you set **MaxResults*	- to **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of the **NextToken*	- parameter.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of network instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **ExpressConnect**: virtual border router (VBR)
	//
	// 	- **VPN**: IPsec-VPN connection
	//
	// 	- **ECR**: Express Connect Router (ECR)
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID of the network instance.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The network instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-p0wfut1iqauelpdpb****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeGrantRulesToResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToResourceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeGrantRulesToResourceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGrantRulesToResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGrantRulesToResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGrantRulesToResourceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeGrantRulesToResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantRulesToResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeGrantRulesToResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGrantRulesToResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGrantRulesToResourceRequest) SetMaxResults(v int32) *DescribeGrantRulesToResourceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetNextToken(v string) *DescribeGrantRulesToResourceRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetOwnerAccount(v string) *DescribeGrantRulesToResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetOwnerId(v int64) *DescribeGrantRulesToResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetProductType(v string) *DescribeGrantRulesToResourceRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetRegionId(v string) *DescribeGrantRulesToResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetResourceId(v string) *DescribeGrantRulesToResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetResourceOwnerAccount(v string) *DescribeGrantRulesToResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) SetResourceOwnerId(v int64) *DescribeGrantRulesToResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToResourceRequest) Validate() error {
	return dara.Validate(s)
}
