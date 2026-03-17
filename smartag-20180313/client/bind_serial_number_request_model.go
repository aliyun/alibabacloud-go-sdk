// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSerialNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *BindSerialNumberRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BindSerialNumberRequest
	GetOwnerId() *int64
	SetRegionId(v string) *BindSerialNumberRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *BindSerialNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindSerialNumberRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *BindSerialNumberRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *BindSerialNumberRequest
	GetSmartAGId() *string
}

type BindSerialNumberRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sage62x021922****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SAG instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-r79m060r6oy55******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s BindSerialNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s BindSerialNumberRequest) GoString() string {
	return s.String()
}

func (s *BindSerialNumberRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BindSerialNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindSerialNumberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindSerialNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindSerialNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindSerialNumberRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *BindSerialNumberRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *BindSerialNumberRequest) SetOwnerAccount(v string) *BindSerialNumberRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindSerialNumberRequest) SetOwnerId(v int64) *BindSerialNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *BindSerialNumberRequest) SetRegionId(v string) *BindSerialNumberRequest {
	s.RegionId = &v
	return s
}

func (s *BindSerialNumberRequest) SetResourceOwnerAccount(v string) *BindSerialNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindSerialNumberRequest) SetResourceOwnerId(v int64) *BindSerialNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindSerialNumberRequest) SetSerialNumber(v string) *BindSerialNumberRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindSerialNumberRequest) SetSmartAGId(v string) *BindSerialNumberRequest {
	s.SmartAGId = &v
	return s
}

func (s *BindSerialNumberRequest) Validate() error {
	return dara.Validate(s)
}
