// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserOnlineClientStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeUserOnlineClientStatisticsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUserOnlineClientStatisticsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUserOnlineClientStatisticsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUserOnlineClientStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUserOnlineClientStatisticsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeUserOnlineClientStatisticsRequest
	GetSmartAGId() *string
	SetUserNames(v []*string) *DescribeUserOnlineClientStatisticsRequest
	GetUserNames() []*string
}

type DescribeUserOnlineClientStatisticsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-sfjg*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// doctest
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeUserOnlineClientStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientStatisticsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUserOnlineClientStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserOnlineClientStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserOnlineClientStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUserOnlineClientStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserOnlineClientStatisticsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeUserOnlineClientStatisticsRequest) GetUserNames() []*string {
	return s.UserNames
}

func (s *DescribeUserOnlineClientStatisticsRequest) SetOwnerAccount(v string) *DescribeUserOnlineClientStatisticsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsRequest) SetOwnerId(v int64) *DescribeUserOnlineClientStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsRequest) SetRegionId(v string) *DescribeUserOnlineClientStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeUserOnlineClientStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsRequest) SetResourceOwnerId(v int64) *DescribeUserOnlineClientStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsRequest) SetSmartAGId(v string) *DescribeUserOnlineClientStatisticsRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsRequest) SetUserNames(v []*string) *DescribeUserOnlineClientStatisticsRequest {
	s.UserNames = v
	return s
}

func (s *DescribeUserOnlineClientStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
