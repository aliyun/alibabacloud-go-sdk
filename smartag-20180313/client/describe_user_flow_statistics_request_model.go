// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserFlowStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeUserFlowStatisticsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUserFlowStatisticsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUserFlowStatisticsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUserFlowStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUserFlowStatisticsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeUserFlowStatisticsRequest
	GetSmartAGId() *string
	SetStatisticsDate(v string) *DescribeUserFlowStatisticsRequest
	GetStatisticsDate() *string
	SetUserNames(v []*string) *DescribeUserFlowStatisticsRequest
	GetUserNames() []*string
}

type DescribeUserFlowStatisticsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the SAG APP instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG APP instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-mfkg*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The month during which the data transfer statistics are collected.
	//
	// If you do not specify a month, the current month is queried.
	//
	// example:
	//
	// 201905
	StatisticsDate *string `json:"StatisticsDate,omitempty" xml:"StatisticsDate,omitempty"`
	// The list of usernames of client accounts. Maximum value of N: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeUserFlowStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserFlowStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowStatisticsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUserFlowStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserFlowStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserFlowStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUserFlowStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserFlowStatisticsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeUserFlowStatisticsRequest) GetStatisticsDate() *string {
	return s.StatisticsDate
}

func (s *DescribeUserFlowStatisticsRequest) GetUserNames() []*string {
	return s.UserNames
}

func (s *DescribeUserFlowStatisticsRequest) SetOwnerAccount(v string) *DescribeUserFlowStatisticsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) SetOwnerId(v int64) *DescribeUserFlowStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) SetRegionId(v string) *DescribeUserFlowStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeUserFlowStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) SetResourceOwnerId(v int64) *DescribeUserFlowStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) SetSmartAGId(v string) *DescribeUserFlowStatisticsRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) SetStatisticsDate(v string) *DescribeUserFlowStatisticsRequest {
	s.StatisticsDate = &v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) SetUserNames(v []*string) *DescribeUserFlowStatisticsRequest {
	s.UserNames = v
	return s
}

func (s *DescribeUserFlowStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
