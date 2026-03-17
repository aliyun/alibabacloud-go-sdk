// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagHaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagHaRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagHaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagHaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagHaRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagHaRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagHaRequest
	GetSmartAGSn() *string
}

type DescribeSagHaRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Smart Access Gateway (SAG) instance is deployed.
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
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSagHaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagHaRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagHaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagHaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagHaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagHaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagHaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagHaRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagHaRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagHaRequest) SetOwnerAccount(v string) *DescribeSagHaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagHaRequest) SetOwnerId(v int64) *DescribeSagHaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagHaRequest) SetRegionId(v string) *DescribeSagHaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagHaRequest) SetResourceOwnerAccount(v string) *DescribeSagHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagHaRequest) SetResourceOwnerId(v int64) *DescribeSagHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagHaRequest) SetSmartAGId(v string) *DescribeSagHaRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagHaRequest) SetSmartAGSn(v string) *DescribeSagHaRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagHaRequest) Validate() error {
	return dara.Validate(s)
}
