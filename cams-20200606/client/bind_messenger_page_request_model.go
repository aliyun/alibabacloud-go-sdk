// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindMessengerPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *BindMessengerPageRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *BindMessengerPageRequest
	GetOwnerId() *int64
	SetPageId(v string) *BindMessengerPageRequest
	GetPageId() *string
	SetRegionId(v string) *BindMessengerPageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *BindMessengerPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindMessengerPageRequest
	GetResourceOwnerId() *int64
}

type BindMessengerPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 181916005005216
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindMessengerPageRequest) String() string {
	return dara.Prettify(s)
}

func (s BindMessengerPageRequest) GoString() string {
	return s.String()
}

func (s *BindMessengerPageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindMessengerPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindMessengerPageRequest) GetPageId() *string {
	return s.PageId
}

func (s *BindMessengerPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindMessengerPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindMessengerPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindMessengerPageRequest) SetInstanceId(v string) *BindMessengerPageRequest {
	s.InstanceId = &v
	return s
}

func (s *BindMessengerPageRequest) SetOwnerId(v int64) *BindMessengerPageRequest {
	s.OwnerId = &v
	return s
}

func (s *BindMessengerPageRequest) SetPageId(v string) *BindMessengerPageRequest {
	s.PageId = &v
	return s
}

func (s *BindMessengerPageRequest) SetRegionId(v string) *BindMessengerPageRequest {
	s.RegionId = &v
	return s
}

func (s *BindMessengerPageRequest) SetResourceOwnerAccount(v string) *BindMessengerPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindMessengerPageRequest) SetResourceOwnerId(v int64) *BindMessengerPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindMessengerPageRequest) Validate() error {
	return dara.Validate(s)
}
