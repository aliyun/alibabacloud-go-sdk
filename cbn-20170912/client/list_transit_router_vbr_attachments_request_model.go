// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVbrAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterVbrAttachmentsRequest
	GetCenId() *string
	SetMaxResults(v int32) *ListTransitRouterVbrAttachmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterVbrAttachmentsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterVbrAttachmentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterVbrAttachmentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterVbrAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterVbrAttachmentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterVbrAttachmentsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*ListTransitRouterVbrAttachmentsRequestTag) *ListTransitRouterVbrAttachmentsRequest
	GetTag() []*ListTransitRouterVbrAttachmentsRequestTag
	SetTransitRouterAttachmentId(v string) *ListTransitRouterVbrAttachmentsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListTransitRouterVbrAttachmentsRequest
	GetTransitRouterId() *string
}

type ListTransitRouterVbrAttachmentsRequest struct {
	// The IDs of the CEN instances.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// dd20****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Enterprise Edition transit router.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*ListTransitRouterVbrAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VBR connection.
	//
	// example:
	//
	// tr-attach-oyf70wfuorwx87****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterVbrAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetTag() []*ListTransitRouterVbrAttachmentsRequestTag {
	return s.Tag
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterVbrAttachmentsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetCenId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterVbrAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetNextToken(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterVbrAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetRegionId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterVbrAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetTag(v []*ListTransitRouterVbrAttachmentsRequestTag) *ListTransitRouterVbrAttachmentsRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterVbrAttachmentsRequestTag struct {
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

func (s ListTransitRouterVbrAttachmentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterVbrAttachmentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterVbrAttachmentsRequestTag) SetKey(v string) *ListTransitRouterVbrAttachmentsRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequestTag) SetValue(v string) *ListTransitRouterVbrAttachmentsRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequestTag) Validate() error {
	return dara.Validate(s)
}
