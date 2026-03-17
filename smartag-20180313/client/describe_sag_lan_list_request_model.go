// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagLanListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagLanListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagLanListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagLanListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagLanListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagLanListRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagLanListRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagLanListRequest
	GetSmartAGSn() *string
}

type DescribeSagLanListRequest struct {
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
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSagLanListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagLanListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagLanListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagLanListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagLanListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagLanListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagLanListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagLanListRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagLanListRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagLanListRequest) SetOwnerAccount(v string) *DescribeSagLanListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagLanListRequest) SetOwnerId(v int64) *DescribeSagLanListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagLanListRequest) SetRegionId(v string) *DescribeSagLanListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagLanListRequest) SetResourceOwnerAccount(v string) *DescribeSagLanListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagLanListRequest) SetResourceOwnerId(v int64) *DescribeSagLanListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagLanListRequest) SetSmartAGId(v string) *DescribeSagLanListRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagLanListRequest) SetSmartAGSn(v string) *DescribeSagLanListRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagLanListRequest) Validate() error {
	return dara.Validate(s)
}
