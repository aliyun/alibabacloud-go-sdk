// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagPortListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagPortListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagPortListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagPortListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagPortListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagPortListRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagPortListRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagPortListRequest
	GetSmartAGSn() *string
}

type DescribeSagPortListRequest struct {
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
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSagPortListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagPortListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagPortListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagPortListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagPortListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagPortListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagPortListRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagPortListRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagPortListRequest) SetOwnerAccount(v string) *DescribeSagPortListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagPortListRequest) SetOwnerId(v int64) *DescribeSagPortListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagPortListRequest) SetRegionId(v string) *DescribeSagPortListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagPortListRequest) SetResourceOwnerAccount(v string) *DescribeSagPortListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagPortListRequest) SetResourceOwnerId(v int64) *DescribeSagPortListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagPortListRequest) SetSmartAGId(v string) *DescribeSagPortListRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagPortListRequest) SetSmartAGSn(v string) *DescribeSagPortListRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagPortListRequest) Validate() error {
	return dara.Validate(s)
}
