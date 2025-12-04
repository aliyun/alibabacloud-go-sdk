// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorDataReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMonitorDataReportRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreateMonitorDataReportRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateMonitorDataReportRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateMonitorDataReportRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateMonitorDataReportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMonitorDataReportRequest
	GetResourceOwnerId() *int64
}

type CreateMonitorDataReportRequest struct {
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

func (s CreateMonitorDataReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorDataReportRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorDataReportRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMonitorDataReportRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateMonitorDataReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMonitorDataReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorDataReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMonitorDataReportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMonitorDataReportRequest) SetDBClusterId(v string) *CreateMonitorDataReportRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetOwnerAccount(v string) *CreateMonitorDataReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetOwnerId(v int64) *CreateMonitorDataReportRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetRegionId(v string) *CreateMonitorDataReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetResourceOwnerAccount(v string) *CreateMonitorDataReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetResourceOwnerId(v int64) *CreateMonitorDataReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMonitorDataReportRequest) Validate() error {
	return dara.Validate(s)
}
