// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterMulticastDomainsRequest
	GetCenId() *string
	SetClientToken(v string) *ListTransitRouterMulticastDomainsRequest
	GetClientToken() *string
	SetMaxResults(v int64) *ListTransitRouterMulticastDomainsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTransitRouterMulticastDomainsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterMulticastDomainsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterMulticastDomainsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterMulticastDomainsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterMulticastDomainsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterMulticastDomainsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*ListTransitRouterMulticastDomainsRequestTag) *ListTransitRouterMulticastDomainsRequest
	GetTag() []*ListTransitRouterMulticastDomainsRequestTag
	SetTransitRouterId(v string) *ListTransitRouterMulticastDomainsRequest
	GetTransitRouterId() *string
	SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastDomainsRequest
	GetTransitRouterMulticastDomainId() *string
}

type ListTransitRouterMulticastDomainsRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-a7syd349kne38g****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- If a value is returned for NextToken, specify the value in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the transit router belongs.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	//
	// You can specify up to 20 tags in each call.
	Tag []*ListTransitRouterMulticastDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The transit router ID.
	//
	// example:
	//
	// tr-p0wr9p28r92d598y6****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The multicast domain ID.
	//
	// example:
	//
	// tr-mcast-domain-3r3bvbypxqheej****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
}

func (s ListTransitRouterMulticastDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterMulticastDomainsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTransitRouterMulticastDomainsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastDomainsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastDomainsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterMulticastDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterMulticastDomainsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterMulticastDomainsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterMulticastDomainsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterMulticastDomainsRequest) GetTag() []*ListTransitRouterMulticastDomainsRequestTag {
	return s.Tag
}

func (s *ListTransitRouterMulticastDomainsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterMulticastDomainsRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *ListTransitRouterMulticastDomainsRequest) SetCenId(v string) *ListTransitRouterMulticastDomainsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetClientToken(v string) *ListTransitRouterMulticastDomainsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetMaxResults(v int64) *ListTransitRouterMulticastDomainsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetNextToken(v string) *ListTransitRouterMulticastDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetOwnerAccount(v string) *ListTransitRouterMulticastDomainsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetOwnerId(v int64) *ListTransitRouterMulticastDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetRegionId(v string) *ListTransitRouterMulticastDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterMulticastDomainsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetResourceOwnerId(v int64) *ListTransitRouterMulticastDomainsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetTag(v []*ListTransitRouterMulticastDomainsRequestTag) *ListTransitRouterMulticastDomainsRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetTransitRouterId(v string) *ListTransitRouterMulticastDomainsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastDomainsRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterMulticastDomainsRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// Each key-value pair must be unique. You can specify up to 20 tag values in each call.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterMulticastDomainsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterMulticastDomainsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterMulticastDomainsRequestTag) SetKey(v string) *ListTransitRouterMulticastDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequestTag) SetValue(v string) *ListTransitRouterMulticastDomainsRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsRequestTag) Validate() error {
	return dara.Validate(s)
}
