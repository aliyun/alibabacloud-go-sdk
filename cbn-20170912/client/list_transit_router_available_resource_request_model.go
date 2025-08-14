// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ListTransitRouterAvailableResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterAvailableResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterAvailableResourceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterAvailableResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterAvailableResourceRequest
	GetResourceOwnerId() *int64
	SetSupportMulticast(v bool) *ListTransitRouterAvailableResourceRequest
	GetSupportMulticast() *bool
}

type ListTransitRouterAvailableResourceRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Enterprise Edition transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to query only the zones in which the multicast feature is supported. Valid values:
	//
	// 	- **true**
	//
	//     After you call **ListTransitRouterAvailableResource**, if no zone is returned, it indicates that the Enterprise Edition transit router does not support the multicast feature.
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	SupportMulticast *bool `json:"SupportMulticast,omitempty" xml:"SupportMulticast,omitempty"`
}

func (s ListTransitRouterAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterAvailableResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterAvailableResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterAvailableResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterAvailableResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterAvailableResourceRequest) GetSupportMulticast() *bool {
	return s.SupportMulticast
}

func (s *ListTransitRouterAvailableResourceRequest) SetOwnerAccount(v string) *ListTransitRouterAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetOwnerId(v int64) *ListTransitRouterAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetRegionId(v string) *ListTransitRouterAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetResourceOwnerAccount(v string) *ListTransitRouterAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetResourceOwnerId(v int64) *ListTransitRouterAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetSupportMulticast(v bool) *ListTransitRouterAvailableResourceRequest {
	s.SupportMulticast = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
