// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesForExpressConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagResourcesForExpressConnectRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagResourcesForExpressConnectRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTagResourcesForExpressConnectRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTagResourcesForExpressConnectRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTagResourcesForExpressConnectRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListTagResourcesForExpressConnectRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *ListTagResourcesForExpressConnectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTagResourcesForExpressConnectRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ListTagResourcesForExpressConnectRequest
	GetResourceType() *string
	SetTag(v []*ListTagResourcesForExpressConnectRequestTag) *ListTagResourcesForExpressConnectRequest
	GetTag() []*ListTagResourcesForExpressConnectRequestTag
}

type ListTagResourcesForExpressConnectRequest struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the resource resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **PHYSICALCONNECTION**: Express Connect circuit.
	//
	// 	- **VIRTUALBORDERROUTER**: virtual border router (VBR).
	//
	// 	- **ROUTERINTERFACE**: router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*ListTagResourcesForExpressConnectRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesForExpressConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForExpressConnectRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForExpressConnectRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagResourcesForExpressConnectRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesForExpressConnectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTagResourcesForExpressConnectRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagResourcesForExpressConnectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesForExpressConnectRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListTagResourcesForExpressConnectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagResourcesForExpressConnectRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTagResourcesForExpressConnectRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesForExpressConnectRequest) GetTag() []*ListTagResourcesForExpressConnectRequestTag {
	return s.Tag
}

func (s *ListTagResourcesForExpressConnectRequest) SetMaxResults(v int32) *ListTagResourcesForExpressConnectRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetNextToken(v string) *ListTagResourcesForExpressConnectRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetOwnerAccount(v string) *ListTagResourcesForExpressConnectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetOwnerId(v int64) *ListTagResourcesForExpressConnectRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetRegionId(v string) *ListTagResourcesForExpressConnectRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetResourceId(v []*string) *ListTagResourcesForExpressConnectRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetResourceOwnerAccount(v string) *ListTagResourcesForExpressConnectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetResourceOwnerId(v int64) *ListTagResourcesForExpressConnectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetResourceType(v string) *ListTagResourcesForExpressConnectRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) SetTag(v []*ListTagResourcesForExpressConnectRequestTag) *ListTagResourcesForExpressConnectRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesForExpressConnectRequest) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesForExpressConnectRequestTag struct {
	// The key of the tag to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource. You can specify up to 20 tag values The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesForExpressConnectRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForExpressConnectRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForExpressConnectRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesForExpressConnectRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesForExpressConnectRequestTag) SetKey(v string) *ListTagResourcesForExpressConnectRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequestTag) SetValue(v string) *ListTagResourcesForExpressConnectRequestTag {
	s.Value = &v
	return s
}

func (s *ListTagResourcesForExpressConnectRequestTag) Validate() error {
	return dara.Validate(s)
}
