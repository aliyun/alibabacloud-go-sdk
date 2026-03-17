// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWanListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagWanListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagWanListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagWanListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagWanListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagWanListRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagWanListRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagWanListRequest
	GetSmartAGSn() *string
}

type DescribeSagWanListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
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

func (s DescribeSagWanListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagWanListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagWanListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagWanListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagWanListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagWanListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagWanListRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagWanListRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagWanListRequest) SetOwnerAccount(v string) *DescribeSagWanListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagWanListRequest) SetOwnerId(v int64) *DescribeSagWanListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagWanListRequest) SetRegionId(v string) *DescribeSagWanListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagWanListRequest) SetResourceOwnerAccount(v string) *DescribeSagWanListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagWanListRequest) SetResourceOwnerId(v int64) *DescribeSagWanListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagWanListRequest) SetSmartAGId(v string) *DescribeSagWanListRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagWanListRequest) SetSmartAGSn(v string) *DescribeSagWanListRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagWanListRequest) Validate() error {
	return dara.Validate(s)
}
