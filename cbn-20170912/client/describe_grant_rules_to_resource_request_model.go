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
	// - If you omit this parameter, all entries are returned in a single response. In this case, the **MaxResults*	- field in the response indicates the total number of entries.
	//
	// - If you specify the **MaxResults*	- parameter, the query is paginated. **MaxResults*	- sets the number of entries per page. The value must be an integer from **1*	- to **100**. The **MaxResults*	- value in the response indicates the number of entries on the current page. The recommended value for this parameter is **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. Valid values:
	//
	// - Omit this parameter for the first request.
	//
	// - For subsequent requests, set this parameter to the **NextToken*	- value from the previous response.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// - **VPC**: a Virtual Private Cloud (VPC) instance.
	//
	// - **ExpressConnect**: a Virtual Border Router (VBR) instance.
	//
	// - **VPN**: an IPsec connection.
	//
	// - **ECR**: an ExpressConnect Router (ECR) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID of the network instance.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the network instance.
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
