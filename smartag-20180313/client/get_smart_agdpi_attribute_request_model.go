// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAGDpiAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetSmartAGDpiAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetSmartAGDpiAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetSmartAGDpiAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetSmartAGDpiAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmartAGDpiAttributeRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *GetSmartAGDpiAttributeRequest
	GetSmartAGId() *string
}

type GetSmartAGDpiAttributeRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-tq3sazs17smldn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s GetSmartAGDpiAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAGDpiAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetSmartAGDpiAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetSmartAGDpiAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmartAGDpiAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSmartAGDpiAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmartAGDpiAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmartAGDpiAttributeRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *GetSmartAGDpiAttributeRequest) SetOwnerAccount(v string) *GetSmartAGDpiAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetSmartAGDpiAttributeRequest) SetOwnerId(v int64) *GetSmartAGDpiAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmartAGDpiAttributeRequest) SetRegionId(v string) *GetSmartAGDpiAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetSmartAGDpiAttributeRequest) SetResourceOwnerAccount(v string) *GetSmartAGDpiAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmartAGDpiAttributeRequest) SetResourceOwnerId(v int64) *GetSmartAGDpiAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmartAGDpiAttributeRequest) SetSmartAGId(v string) *GetSmartAGDpiAttributeRequest {
	s.SmartAGId = &v
	return s
}

func (s *GetSmartAGDpiAttributeRequest) Validate() error {
	return dara.Validate(s)
}
