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
	// The CEN instance ID.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The maximum number of entries to return per page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// - If this is the first query or no subsequent query is to be sent, you do not need to specify this parameter.
	//
	// - If a subsequent query is to be sent, set the value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// dd20****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Enterprise Edition transit router instance.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the region ID.
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
	Tag []*ListTransitRouterVbrAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VBR connection.
	//
	// example:
	//
	// tr-attach-oyf70wfuorwx87****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The Enterprise Edition forward routing instance ID.
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

type ListTransitRouterVbrAttachmentsRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value can be an empty string or up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
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
