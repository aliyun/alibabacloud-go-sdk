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
	// The CEN instance ID.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of entries per page for a paged query. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query exists.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router instance.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
	Tag []*ListTransitRouterEcrAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ECR connection ID.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The transit router instance ID.
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

type ListTransitRouterEcrAttachmentsRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value can be an empty string or up to 128 characters in length. It cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
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
