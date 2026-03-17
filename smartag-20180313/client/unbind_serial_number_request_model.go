// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSerialNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *UnbindSerialNumberRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnbindSerialNumberRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnbindSerialNumberRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnbindSerialNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindSerialNumberRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *UnbindSerialNumberRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *UnbindSerialNumberRequest
	GetSmartAGId() *string
}

type UnbindSerialNumberRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the SAG instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number of the SAG device to be disassociated.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-sh-0-0927-16****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-0phdojgu5tqr1p****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s UnbindSerialNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindSerialNumberRequest) GoString() string {
	return s.String()
}

func (s *UnbindSerialNumberRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnbindSerialNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindSerialNumberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnbindSerialNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindSerialNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindSerialNumberRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UnbindSerialNumberRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UnbindSerialNumberRequest) SetOwnerAccount(v string) *UnbindSerialNumberRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindSerialNumberRequest) SetOwnerId(v int64) *UnbindSerialNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindSerialNumberRequest) SetRegionId(v string) *UnbindSerialNumberRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindSerialNumberRequest) SetResourceOwnerAccount(v string) *UnbindSerialNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindSerialNumberRequest) SetResourceOwnerId(v int64) *UnbindSerialNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindSerialNumberRequest) SetSerialNumber(v string) *UnbindSerialNumberRequest {
	s.SerialNumber = &v
	return s
}

func (s *UnbindSerialNumberRequest) SetSmartAGId(v string) *UnbindSerialNumberRequest {
	s.SmartAGId = &v
	return s
}

func (s *UnbindSerialNumberRequest) Validate() error {
	return dara.Validate(s)
}
