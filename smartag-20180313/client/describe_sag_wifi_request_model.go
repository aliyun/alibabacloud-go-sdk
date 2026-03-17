// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWifiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagWifiRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagWifiRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagWifiRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagWifiRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagWifiRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagWifiRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagWifiRequest
	GetSmartAGSn() *string
}

type DescribeSagWifiRequest struct {
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
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number (SN) of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSagWifiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWifiRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagWifiRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagWifiRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagWifiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagWifiRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagWifiRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagWifiRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagWifiRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagWifiRequest) SetOwnerAccount(v string) *DescribeSagWifiRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagWifiRequest) SetOwnerId(v int64) *DescribeSagWifiRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagWifiRequest) SetRegionId(v string) *DescribeSagWifiRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagWifiRequest) SetResourceOwnerAccount(v string) *DescribeSagWifiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagWifiRequest) SetResourceOwnerId(v int64) *DescribeSagWifiRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagWifiRequest) SetSmartAGId(v string) *DescribeSagWifiRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagWifiRequest) SetSmartAGSn(v string) *DescribeSagWifiRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagWifiRequest) Validate() error {
	return dara.Validate(s)
}
