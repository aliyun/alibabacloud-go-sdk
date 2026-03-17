// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRemoteAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagRemoteAccessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagRemoteAccessRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagRemoteAccessRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagRemoteAccessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagRemoteAccessRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *DescribeSagRemoteAccessRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *DescribeSagRemoteAccessRequest
	GetSmartAGId() *string
}

type DescribeSagRemoteAccessRequest struct {
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
	// The serial number of the SAG device.
	//
	// example:
	//
	// sage62x022502****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeSagRemoteAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRemoteAccessRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagRemoteAccessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagRemoteAccessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagRemoteAccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagRemoteAccessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagRemoteAccessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagRemoteAccessRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeSagRemoteAccessRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagRemoteAccessRequest) SetOwnerAccount(v string) *DescribeSagRemoteAccessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagRemoteAccessRequest) SetOwnerId(v int64) *DescribeSagRemoteAccessRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagRemoteAccessRequest) SetRegionId(v string) *DescribeSagRemoteAccessRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagRemoteAccessRequest) SetResourceOwnerAccount(v string) *DescribeSagRemoteAccessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagRemoteAccessRequest) SetResourceOwnerId(v int64) *DescribeSagRemoteAccessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagRemoteAccessRequest) SetSerialNumber(v string) *DescribeSagRemoteAccessRequest {
	s.SerialNumber = &v
	return s
}

func (s *DescribeSagRemoteAccessRequest) SetSmartAGId(v string) *DescribeSagRemoteAccessRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagRemoteAccessRequest) Validate() error {
	return dara.Validate(s)
}
