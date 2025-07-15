// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeTagKeysRequest
	GetKeyword() *string
	SetMaxResult(v int32) *DescribeTagKeysRequest
	GetMaxResult() *int32
	SetNextToken(v string) *DescribeTagKeysRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeTagKeysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTagKeysRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTagKeysRequest
	GetRegionId() *string
	SetResourceId(v []*string) *DescribeTagKeysRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *DescribeTagKeysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTagKeysRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeTagKeysRequest
	GetResourceType() *string
}

type DescribeTagKeysRequest struct {
	// The tag keys.
	//
	// example:
	//
	// keyword
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query or no next query is to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 20 resource IDs.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VSWITCH**: vSwitch
	//
	// 	- **ROUTETABLE**: route table
	//
	// 	- **EIP**: elastic IP address (EIP)
	//
	// 	- **VpnGateway**: VPN gateway
	//
	// 	- **NATGATEWAY**: NAT gateway
	//
	// 	- **COMMONBANDWIDTHPACKAGE**: EIP bandwidth plan
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeTagKeysRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *DescribeTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagKeysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTagKeysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTagKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagKeysRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeTagKeysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTagKeysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagKeysRequest) SetKeyword(v string) *DescribeTagKeysRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeTagKeysRequest) SetMaxResult(v int32) *DescribeTagKeysRequest {
	s.MaxResult = &v
	return s
}

func (s *DescribeTagKeysRequest) SetNextToken(v string) *DescribeTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTagKeysRequest) SetOwnerAccount(v string) *DescribeTagKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagKeysRequest) SetOwnerId(v int64) *DescribeTagKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetRegionId(v string) *DescribeTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceId(v []*string) *DescribeTagKeysRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceOwnerAccount(v string) *DescribeTagKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceOwnerId(v int64) *DescribeTagKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceType(v string) *DescribeTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
