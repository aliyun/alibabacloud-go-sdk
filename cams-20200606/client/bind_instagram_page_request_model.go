// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstagramPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *BindInstagramPageRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *BindInstagramPageRequest
	GetOwnerId() *int64
	SetPageId(v string) *BindInstagramPageRequest
	GetPageId() *string
	SetRegionId(v string) *BindInstagramPageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *BindInstagramPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindInstagramPageRequest
	GetResourceOwnerId() *int64
}

type BindInstagramPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14711522522
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindInstagramPageRequest) String() string {
	return dara.Prettify(s)
}

func (s BindInstagramPageRequest) GoString() string {
	return s.String()
}

func (s *BindInstagramPageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindInstagramPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindInstagramPageRequest) GetPageId() *string {
	return s.PageId
}

func (s *BindInstagramPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindInstagramPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindInstagramPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindInstagramPageRequest) SetInstanceId(v string) *BindInstagramPageRequest {
	s.InstanceId = &v
	return s
}

func (s *BindInstagramPageRequest) SetOwnerId(v int64) *BindInstagramPageRequest {
	s.OwnerId = &v
	return s
}

func (s *BindInstagramPageRequest) SetPageId(v string) *BindInstagramPageRequest {
	s.PageId = &v
	return s
}

func (s *BindInstagramPageRequest) SetRegionId(v string) *BindInstagramPageRequest {
	s.RegionId = &v
	return s
}

func (s *BindInstagramPageRequest) SetResourceOwnerAccount(v string) *BindInstagramPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindInstagramPageRequest) SetResourceOwnerId(v int64) *BindInstagramPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindInstagramPageRequest) Validate() error {
	return dara.Validate(s)
}
