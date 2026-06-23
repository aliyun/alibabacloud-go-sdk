// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RestartInstanceRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *RestartInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartInstanceRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *RestartInstanceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *RestartInstanceRequest
	GetPageSize() *int32
	SetRegionId(v string) *RestartInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RestartInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartInstanceRequest
	GetResourceOwnerId() *int64
	SetRestartTime(v string) *RestartInstanceRequest
	GetRestartTime() *string
}

type RestartInstanceRequest struct {
	// The cluster ID. Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all clusters in the destination region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - 30 (default)
	//
	// - 50
	//
	// - 100
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. Call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduled restart time. The time must be in the yyyy-MM-ddTHH:mmZ format and in UTC.
	//
	// > If this parameter is empty or specifies a time earlier than the current time, the cluster is restarted immediately.
	//
	// example:
	//
	// 2023-03-22T00:00:50Z
	RestartTime *string `json:"RestartTime,omitempty" xml:"RestartTime,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartInstanceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RestartInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RestartInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartInstanceRequest) GetRestartTime() *string {
	return s.RestartTime
}

func (s *RestartInstanceRequest) SetDBClusterId(v string) *RestartInstanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerAccount(v string) *RestartInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerId(v int64) *RestartInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartInstanceRequest) SetPageNumber(v int32) *RestartInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *RestartInstanceRequest) SetPageSize(v int32) *RestartInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *RestartInstanceRequest) SetRegionId(v string) *RestartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerAccount(v string) *RestartInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerId(v int64) *RestartInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartInstanceRequest) SetRestartTime(v string) *RestartInstanceRequest {
	s.RestartTime = &v
	return s
}

func (s *RestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
