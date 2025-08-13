// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreatedByEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CheckCreatedByEnabledRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCreatedByEnabledRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckCreatedByEnabledRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckCreatedByEnabledRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *CheckCreatedByEnabledRequest
	GetResourceOwnerId() *string
}

type CheckCreatedByEnabledRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckCreatedByEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCreatedByEnabledRequest) GoString() string {
	return s.String()
}

func (s *CheckCreatedByEnabledRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCreatedByEnabledRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCreatedByEnabledRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCreatedByEnabledRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCreatedByEnabledRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *CheckCreatedByEnabledRequest) SetOwnerAccount(v string) *CheckCreatedByEnabledRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetOwnerId(v int64) *CheckCreatedByEnabledRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetRegionId(v string) *CheckCreatedByEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetResourceOwnerAccount(v string) *CheckCreatedByEnabledRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetResourceOwnerId(v string) *CheckCreatedByEnabledRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) Validate() error {
	return dara.Validate(s)
}
