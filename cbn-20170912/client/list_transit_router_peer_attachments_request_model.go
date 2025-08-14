// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterPeerAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterPeerAttachmentsRequest
	GetCenId() *string
	SetMaxResults(v int32) *ListTransitRouterPeerAttachmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterPeerAttachmentsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterPeerAttachmentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterPeerAttachmentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterPeerAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterPeerAttachmentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterPeerAttachmentsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*ListTransitRouterPeerAttachmentsRequestTag) *ListTransitRouterPeerAttachmentsRequest
	GetTag() []*ListTransitRouterPeerAttachmentsRequestTag
	SetTransitRouterAttachmentId(v string) *ListTransitRouterPeerAttachmentsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListTransitRouterPeerAttachmentsRequest
	GetTransitRouterId() *string
}

type ListTransitRouterPeerAttachmentsRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of entries returned per page. Default value: **20**. Maximum value: **100**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// 1
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Enterprise Edition transit router is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the tag.
	//
	// You can specify at most 20 tags in each call.
	Tag []*ListTransitRouterPeerAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the inter-region connection.
	//
	// example:
	//
	// tr-attach-5u4qbayfv2io5v****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterPeerAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetTag() []*ListTransitRouterPeerAttachmentsRequestTag {
	return s.Tag
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterPeerAttachmentsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetCenId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterPeerAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetNextToken(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterPeerAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetRegionId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterPeerAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetTag(v []*ListTransitRouterPeerAttachmentsRequestTag) *ListTransitRouterPeerAttachmentsRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterPeerAttachmentsRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterPeerAttachmentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterPeerAttachmentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterPeerAttachmentsRequestTag) SetKey(v string) *ListTransitRouterPeerAttachmentsRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequestTag) SetValue(v string) *ListTransitRouterPeerAttachmentsRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequestTag) Validate() error {
	return dara.Validate(s)
}
