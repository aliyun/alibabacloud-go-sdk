// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachType(v string) *DescribeVpnAttachmentsRequest
	GetAttachType() *string
	SetOwnerAccount(v string) *DescribeVpnAttachmentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnAttachmentsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpnAttachmentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnAttachmentsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnAttachmentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnAttachmentsRequest
	GetResourceOwnerId() *int64
	SetVpnConnectionId(v string) *DescribeVpnAttachmentsRequest
	GetVpnConnectionId() *string
}

type DescribeVpnAttachmentsRequest struct {
	// The type of resource to which the IPsec-VPN connection is attached. Default value: **CEN**.
	//
	// The value can only be **CEN**, which indicates that the system queries IPsec-VPN connections that are associated with transit router instances.
	//
	// example:
	//
	// CEN
	AttachType   *string `json:"AttachType,omitempty" xml:"AttachType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page when using paging. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// > If you do not specify the IPsec-VPN connection ID, the system queries all IPsec-VPN connections that are associated with transit router instances in the current region.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DescribeVpnAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnAttachmentsRequest) GetAttachType() *string {
	return s.AttachType
}

func (s *DescribeVpnAttachmentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnAttachmentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnAttachmentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnAttachmentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnAttachmentsRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnAttachmentsRequest) SetAttachType(v string) *DescribeVpnAttachmentsRequest {
	s.AttachType = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetOwnerAccount(v string) *DescribeVpnAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetOwnerId(v int64) *DescribeVpnAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetPageNumber(v int32) *DescribeVpnAttachmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetPageSize(v int32) *DescribeVpnAttachmentsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetRegionId(v string) *DescribeVpnAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetResourceOwnerAccount(v string) *DescribeVpnAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetResourceOwnerId(v int64) *DescribeVpnAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) SetVpnConnectionId(v string) *DescribeVpnAttachmentsRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}
