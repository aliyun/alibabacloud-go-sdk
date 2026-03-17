// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v int32) *UpdateSmartAGAccessPointRequest
	GetAccessPointId() *int32
	SetOwnerAccount(v string) *UpdateSmartAGAccessPointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateSmartAGAccessPointRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateSmartAGAccessPointRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateSmartAGAccessPointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmartAGAccessPointRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *UpdateSmartAGAccessPointRequest
	GetSmartAGId() *string
}

type UpdateSmartAGAccessPointRequest struct {
	// The ID of the access point to which you want to switch.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	AccessPointId *int32  `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// sag-far8v6owtdxlua****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s UpdateSmartAGAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGAccessPointRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGAccessPointRequest) GetAccessPointId() *int32 {
	return s.AccessPointId
}

func (s *UpdateSmartAGAccessPointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateSmartAGAccessPointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmartAGAccessPointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAGAccessPointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmartAGAccessPointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmartAGAccessPointRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UpdateSmartAGAccessPointRequest) SetAccessPointId(v int32) *UpdateSmartAGAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *UpdateSmartAGAccessPointRequest) SetOwnerAccount(v string) *UpdateSmartAGAccessPointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateSmartAGAccessPointRequest) SetOwnerId(v int64) *UpdateSmartAGAccessPointRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmartAGAccessPointRequest) SetRegionId(v string) *UpdateSmartAGAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAGAccessPointRequest) SetResourceOwnerAccount(v string) *UpdateSmartAGAccessPointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmartAGAccessPointRequest) SetResourceOwnerId(v int64) *UpdateSmartAGAccessPointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmartAGAccessPointRequest) SetSmartAGId(v string) *UpdateSmartAGAccessPointRequest {
	s.SmartAGId = &v
	return s
}

func (s *UpdateSmartAGAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
