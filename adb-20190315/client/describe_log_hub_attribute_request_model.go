// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogHubAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLogHubAttributeRequest
	GetDBClusterId() *string
	SetDeliverName(v string) *DescribeLogHubAttributeRequest
	GetDeliverName() *string
	SetLogStoreName(v string) *DescribeLogHubAttributeRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *DescribeLogHubAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLogHubAttributeRequest
	GetOwnerId() *int64
	SetProjectName(v string) *DescribeLogHubAttributeRequest
	GetProjectName() *string
	SetRegionId(v string) *DescribeLogHubAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLogHubAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLogHubAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeLogHubAttributeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1nz6smy07szs58p
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the log shipping job.
	//
	// This parameter is required.
	//
	// example:
	//
	// sz_sls2adb_kxdpz_af_data
	DeliverName *string `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// game2-sms-log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// dcsz-af-data
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLogHubAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogHubAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLogHubAttributeRequest) GetDeliverName() *string {
	return s.DeliverName
}

func (s *DescribeLogHubAttributeRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeLogHubAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLogHubAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLogHubAttributeRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLogHubAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogHubAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLogHubAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLogHubAttributeRequest) SetDBClusterId(v string) *DescribeLogHubAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetDeliverName(v string) *DescribeLogHubAttributeRequest {
	s.DeliverName = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetLogStoreName(v string) *DescribeLogHubAttributeRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetOwnerAccount(v string) *DescribeLogHubAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetOwnerId(v int64) *DescribeLogHubAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetProjectName(v string) *DescribeLogHubAttributeRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetRegionId(v string) *DescribeLogHubAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLogHubAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetResourceOwnerId(v int64) *DescribeLogHubAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) Validate() error {
	return dara.Validate(s)
}
