// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagManagementPortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagManagementPortRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagManagementPortRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagManagementPortRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagManagementPortRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagManagementPortRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagManagementPortRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagManagementPortRequest
	GetSmartAGSn() *string
}

type DescribeSagManagementPortRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-jbauqzw5ildnud****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sage62x021922****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSagManagementPortRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagManagementPortRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagManagementPortRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagManagementPortRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagManagementPortRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagManagementPortRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagManagementPortRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagManagementPortRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagManagementPortRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagManagementPortRequest) SetOwnerAccount(v string) *DescribeSagManagementPortRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagManagementPortRequest) SetOwnerId(v int64) *DescribeSagManagementPortRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagManagementPortRequest) SetRegionId(v string) *DescribeSagManagementPortRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagManagementPortRequest) SetResourceOwnerAccount(v string) *DescribeSagManagementPortRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagManagementPortRequest) SetResourceOwnerId(v int64) *DescribeSagManagementPortRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagManagementPortRequest) SetSmartAGId(v string) *DescribeSagManagementPortRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagManagementPortRequest) SetSmartAGSn(v string) *DescribeSagManagementPortRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagManagementPortRequest) Validate() error {
	return dara.Validate(s)
}
