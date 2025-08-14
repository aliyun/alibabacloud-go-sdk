// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVpnAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetCenId() *string
	SetMaxResults(v int32) *ListTransitRouterVpnAttachmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterVpnAttachmentsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*ListTransitRouterVpnAttachmentsRequestTag) *ListTransitRouterVpnAttachmentsRequest
	GetTag() []*ListTransitRouterVpnAttachmentsRequestTag
	SetTransitRouterAttachmentId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetTransitRouterId() *string
}

type ListTransitRouterVpnAttachmentsRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-rsgxs8ng2awen2****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. Valid values:
	//
	// 	- If this is your first query and no subsequent queries are to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the tags that are added to the CEN instance.
	//
	// You can query at most 20 tags in each call.
	Tag []*ListTransitRouterVpnAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPN attachment.
	//
	// example:
	//
	// tr-attach-a6p8voaodog5c0****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-p0wm740vjnbaprv0m****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterVpnAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetTag() []*ListTransitRouterVpnAttachmentsRequestTag {
	return s.Tag
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetCenId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterVpnAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetNextToken(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetRegionId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetTag(v []*ListTransitRouterVpnAttachmentsRequestTag) *ListTransitRouterVpnAttachmentsRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterVpnAttachmentsRequestTag struct {
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
	// value_A1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterVpnAttachmentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) SetKey(v string) *ListTransitRouterVpnAttachmentsRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) SetValue(v string) *ListTransitRouterVpnAttachmentsRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) Validate() error {
	return dara.Validate(s)
}
