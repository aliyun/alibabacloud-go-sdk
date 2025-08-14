// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterEcrAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterEcrAttachmentsRequest
	GetCenId() *string
	SetMaxResults(v int32) *ListTransitRouterEcrAttachmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterEcrAttachmentsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterEcrAttachmentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterEcrAttachmentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterEcrAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterEcrAttachmentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterEcrAttachmentsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*ListTransitRouterEcrAttachmentsRequestTag) *ListTransitRouterEcrAttachmentsRequest
	GetTag() []*ListTransitRouterEcrAttachmentsRequestTag
	SetTransitRouterAttachmentId(v string) *ListTransitRouterEcrAttachmentsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListTransitRouterEcrAttachmentsRequest
	GetTransitRouterId() *string
}

type ListTransitRouterEcrAttachmentsRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
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
	// The tags.
	//
	// You can specify at most 20 tags.
	Tag []*ListTransitRouterEcrAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the ECR connection.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-6ehx7q2jze8ch5ji0****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterEcrAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterEcrAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetTag() []*ListTransitRouterEcrAttachmentsRequestTag {
	return s.Tag
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterEcrAttachmentsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetCenId(v string) *ListTransitRouterEcrAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterEcrAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetNextToken(v string) *ListTransitRouterEcrAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterEcrAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterEcrAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetRegionId(v string) *ListTransitRouterEcrAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterEcrAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterEcrAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetTag(v []*ListTransitRouterEcrAttachmentsRequestTag) *ListTransitRouterEcrAttachmentsRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterEcrAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterEcrAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterEcrAttachmentsRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterEcrAttachmentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterEcrAttachmentsRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterEcrAttachmentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterEcrAttachmentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterEcrAttachmentsRequestTag) SetKey(v string) *ListTransitRouterEcrAttachmentsRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequestTag) SetValue(v string) *ListTransitRouterEcrAttachmentsRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsRequestTag) Validate() error {
	return dara.Validate(s)
}
