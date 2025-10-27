// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutorDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeExecutorDetectionRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeExecutorDetectionRequest
	GetEndTime() *string
	SetLang(v string) *DescribeExecutorDetectionRequest
	GetLang() *string
	SetOwnerAccount(v string) *DescribeExecutorDetectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeExecutorDetectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeExecutorDetectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeExecutorDetectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeExecutorDetectionRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeExecutorDetectionRequest
	GetStartTime() *string
}

type DescribeExecutorDetectionRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// >
	//
	// 	- The end time must be later than the start time.
	//
	// 	- The maximum time range that can be specified is 24 hours.
	//
	// example:
	//
	// 2021-05-27T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language. Valid values:
	//
	// 	- **zh**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2024-10-31T02:06Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeExecutorDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeExecutorDetectionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeExecutorDetectionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExecutorDetectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeExecutorDetectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExecutorDetectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExecutorDetectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeExecutorDetectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeExecutorDetectionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeExecutorDetectionRequest) SetDBClusterId(v string) *DescribeExecutorDetectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetEndTime(v string) *DescribeExecutorDetectionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetLang(v string) *DescribeExecutorDetectionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetOwnerAccount(v string) *DescribeExecutorDetectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetOwnerId(v int64) *DescribeExecutorDetectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetRegionId(v string) *DescribeExecutorDetectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetResourceOwnerAccount(v string) *DescribeExecutorDetectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetResourceOwnerId(v int64) *DescribeExecutorDetectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) SetStartTime(v string) *DescribeExecutorDetectionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeExecutorDetectionRequest) Validate() error {
	return dara.Validate(s)
}
