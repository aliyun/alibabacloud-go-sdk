// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVpcAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterVpcAttachmentsRequest
	GetCenId() *string
	SetMaxResults(v int32) *ListTransitRouterVpcAttachmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterVpcAttachmentsRequest
	GetNextToken() *string
	SetOrderType(v string) *ListTransitRouterVpcAttachmentsRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *ListTransitRouterVpcAttachmentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterVpcAttachmentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterVpcAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterVpcAttachmentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterVpcAttachmentsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListTransitRouterVpcAttachmentsRequest
	GetStatus() *string
	SetTag(v []*ListTransitRouterVpcAttachmentsRequestTag) *ListTransitRouterVpcAttachmentsRequest
	GetTag() []*ListTransitRouterVpcAttachmentsRequestTag
	SetTransitRouterAttachmentId(v string) *ListTransitRouterVpcAttachmentsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListTransitRouterVpcAttachmentsRequest
	GetTransitRouterId() *string
	SetVpcId(v string) *ListTransitRouterVpcAttachmentsRequest
	GetVpcId() *string
}

type ListTransitRouterVpcAttachmentsRequest struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of entries to return on each page. The default value is **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// - If this is your first query or no more results are available, you do not need to specify this parameter.
	//
	// - For subsequent queries, set this parameter to the `NextToken` value from the previous response.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The entity that pays for the network instance. Valid values:
	//
	// - **PayByCenOwner**: The fees are paid by the account that owns the CEN instance.
	//
	// - **PayByResourceOwner**: The fees are paid by the account that owns the network instance.
	//
	// example:
	//
	// PayByCenOwner
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
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
	// The status of the VPC connection. Valid values:
	//
	// - **Attaching**: The VPC connection is being created.
	//
	// - **Attached**: The VPC connection is created.
	//
	// - **Detaching**: The VPC connection is being deleted.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	//
	// You can specify up to 20 tags.
	Tag []*ListTransitRouterVpcAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC connection.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1kbjcre9vtsebo1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetTag() []*ListTransitRouterVpcAttachmentsRequestTag {
	return s.Tag
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterVpcAttachmentsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetCenId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterVpcAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetNextToken(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetOrderType(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.OrderType = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterVpcAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetRegionId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterVpcAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetStatus(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.Status = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetTag(v []*ListTransitRouterVpcAttachmentsRequestTag) *ListTransitRouterVpcAttachmentsRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetVpcId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.VpcId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) Validate() error {
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

type ListTransitRouterVpcAttachmentsRequestTag struct {
	// The tag key.
	//
	// The tag key must be 1 to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterVpcAttachmentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterVpcAttachmentsRequestTag) SetKey(v string) *ListTransitRouterVpcAttachmentsRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequestTag) SetValue(v string) *ListTransitRouterVpcAttachmentsRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequestTag) Validate() error {
	return dara.Validate(s)
}
