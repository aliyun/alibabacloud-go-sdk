// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBackupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterBackupId(v string) *DescribeClusterBackupListRequest
	GetClusterBackupId() *string
	SetEndTime(v string) *DescribeClusterBackupListRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeClusterBackupListRequest
	GetInstanceId() *string
	SetNoShardBackup(v string) *DescribeClusterBackupListRequest
	GetNoShardBackup() *string
	SetOwnerAccount(v string) *DescribeClusterBackupListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeClusterBackupListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeClusterBackupListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterBackupListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeClusterBackupListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterBackupListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterBackupListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeClusterBackupListRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeClusterBackupListRequest
	GetStartTime() *string
}

type DescribeClusterBackupListRequest struct {
	// The backup set ID.
	//
	// example:
	//
	// cb-hyxdof5x9kqbtust
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-13T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-t4nj72oug5r5646qog
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to show backup set information for shards in the instance.
	//
	// 	- **true**: does not show backup set information for shards in the instance.
	//
	// 	- **false*	- (default): shows backup set information for shards in the instance.
	//
	// Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	NoShardBackup *string `json:"NoShardBackup,omitempty" xml:"NoShardBackup,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// 	- 200
	//
	// 	- 300
	//
	// 	- 5
	//
	// 	- 10
	//
	// 	- 15
	//
	// 	- 20
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-03T07:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterBackupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupListRequest) GetClusterBackupId() *string {
	return s.ClusterBackupId
}

func (s *DescribeClusterBackupListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeClusterBackupListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClusterBackupListRequest) GetNoShardBackup() *string {
	return s.NoShardBackup
}

func (s *DescribeClusterBackupListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeClusterBackupListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterBackupListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterBackupListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterBackupListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterBackupListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterBackupListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterBackupListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeClusterBackupListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeClusterBackupListRequest) SetClusterBackupId(v string) *DescribeClusterBackupListRequest {
	s.ClusterBackupId = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetEndTime(v string) *DescribeClusterBackupListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetInstanceId(v string) *DescribeClusterBackupListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetNoShardBackup(v string) *DescribeClusterBackupListRequest {
	s.NoShardBackup = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetOwnerAccount(v string) *DescribeClusterBackupListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetOwnerId(v int64) *DescribeClusterBackupListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetPageNumber(v int32) *DescribeClusterBackupListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetPageSize(v int32) *DescribeClusterBackupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetRegionId(v string) *DescribeClusterBackupListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetResourceOwnerAccount(v string) *DescribeClusterBackupListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetResourceOwnerId(v int64) *DescribeClusterBackupListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetSecurityToken(v string) *DescribeClusterBackupListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeClusterBackupListRequest) SetStartTime(v string) *DescribeClusterBackupListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterBackupListRequest) Validate() error {
	return dara.Validate(s)
}
