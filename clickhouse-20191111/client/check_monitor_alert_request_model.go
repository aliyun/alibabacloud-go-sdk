// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMonitorAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckMonitorAlertRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CheckMonitorAlertRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckMonitorAlertRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckMonitorAlertRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckMonitorAlertRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckMonitorAlertRequest
	GetResourceOwnerId() *int64
}

type CheckMonitorAlertRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp13s14l8498l****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckMonitorAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckMonitorAlertRequest) GoString() string {
	return s.String()
}

func (s *CheckMonitorAlertRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckMonitorAlertRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckMonitorAlertRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckMonitorAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckMonitorAlertRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckMonitorAlertRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckMonitorAlertRequest) SetDBClusterId(v string) *CheckMonitorAlertRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetOwnerAccount(v string) *CheckMonitorAlertRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetOwnerId(v int64) *CheckMonitorAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetRegionId(v string) *CheckMonitorAlertRequest {
	s.RegionId = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetResourceOwnerAccount(v string) *CheckMonitorAlertRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetResourceOwnerId(v int64) *CheckMonitorAlertRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckMonitorAlertRequest) Validate() error {
	return dara.Validate(s)
}
