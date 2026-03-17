// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagOnlineClientStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagOnlineClientStatisticsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagOnlineClientStatisticsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagOnlineClientStatisticsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagOnlineClientStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagOnlineClientStatisticsRequest
	GetResourceOwnerId() *int64
	SetSmartAGIds(v []*string) *DescribeSagOnlineClientStatisticsRequest
	GetSmartAGIds() []*string
}

type DescribeSagOnlineClientStatisticsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// sag-va03wf4l4idaj*****
	SmartAGIds []*string `json:"SmartAGIds,omitempty" xml:"SmartAGIds,omitempty" type:"Repeated"`
}

func (s DescribeSagOnlineClientStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagOnlineClientStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagOnlineClientStatisticsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagOnlineClientStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagOnlineClientStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagOnlineClientStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagOnlineClientStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagOnlineClientStatisticsRequest) GetSmartAGIds() []*string {
	return s.SmartAGIds
}

func (s *DescribeSagOnlineClientStatisticsRequest) SetOwnerAccount(v string) *DescribeSagOnlineClientStatisticsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsRequest) SetOwnerId(v int64) *DescribeSagOnlineClientStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsRequest) SetRegionId(v string) *DescribeSagOnlineClientStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeSagOnlineClientStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsRequest) SetResourceOwnerId(v int64) *DescribeSagOnlineClientStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsRequest) SetSmartAGIds(v []*string) *DescribeSagOnlineClientStatisticsRequest {
	s.SmartAGIds = v
	return s
}

func (s *DescribeSagOnlineClientStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
