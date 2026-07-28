// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysForExpressConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeTagKeysForExpressConnectRequest
	GetKeyword() *string
	SetMaxResult(v int32) *DescribeTagKeysForExpressConnectRequest
	GetMaxResult() *int32
	SetNextToken(v string) *DescribeTagKeysForExpressConnectRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeTagKeysForExpressConnectRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTagKeysForExpressConnectRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTagKeysForExpressConnectRequest
	GetRegionId() *string
	SetResourceId(v []*string) *DescribeTagKeysForExpressConnectRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *DescribeTagKeysForExpressConnectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTagKeysForExpressConnectRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeTagKeysForExpressConnectRequest
	GetResourceType() *string
}

type DescribeTagKeysForExpressConnectRequest struct {
	// The keyword of the tag name. Fuzzy match is supported. You can enter a keyword to query all tags that contain the keyword.
	//
	// example:
	//
	// keyword
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries per page for a paged query. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If this is the first query or there is no next query, leave this parameter empty.
	//
	// - If there is a next query, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID. You can specify up to 20 resource IDs.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - **PHYSICALCONNECTION**: Express Connect circuit instance.
	//
	// - **VIRTUALBORDERROUTER**: Virtual Border Router.
	//
	// - **ROUTERINTERFACE**: VBR uplink.
	//
	// - **TRAFFICQOS**: QoS policy.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagKeysForExpressConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysForExpressConnectRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysForExpressConnectRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeTagKeysForExpressConnectRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *DescribeTagKeysForExpressConnectRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagKeysForExpressConnectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTagKeysForExpressConnectRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTagKeysForExpressConnectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagKeysForExpressConnectRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeTagKeysForExpressConnectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTagKeysForExpressConnectRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTagKeysForExpressConnectRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagKeysForExpressConnectRequest) SetKeyword(v string) *DescribeTagKeysForExpressConnectRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetMaxResult(v int32) *DescribeTagKeysForExpressConnectRequest {
	s.MaxResult = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetNextToken(v string) *DescribeTagKeysForExpressConnectRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetOwnerAccount(v string) *DescribeTagKeysForExpressConnectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetOwnerId(v int64) *DescribeTagKeysForExpressConnectRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetRegionId(v string) *DescribeTagKeysForExpressConnectRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetResourceId(v []*string) *DescribeTagKeysForExpressConnectRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetResourceOwnerAccount(v string) *DescribeTagKeysForExpressConnectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetResourceOwnerId(v int64) *DescribeTagKeysForExpressConnectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) SetResourceType(v string) *DescribeTagKeysForExpressConnectRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectRequest) Validate() error {
	return dara.Validate(s)
}
