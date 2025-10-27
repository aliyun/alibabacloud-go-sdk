// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControllerDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeControllerDetectionRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeControllerDetectionRequest
	GetEndTime() *string
	SetLang(v string) *DescribeControllerDetectionRequest
	GetLang() *string
	SetOwnerAccount(v string) *DescribeControllerDetectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeControllerDetectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeControllerDetectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeControllerDetectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeControllerDetectionRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeControllerDetectionRequest
	GetStartTime() *string
}

type DescribeControllerDetectionRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2024-06-23T02:21Z
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
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2021-05-03T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeControllerDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeControllerDetectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeControllerDetectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeControllerDetectionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeControllerDetectionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeControllerDetectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeControllerDetectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeControllerDetectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeControllerDetectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeControllerDetectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeControllerDetectionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeControllerDetectionRequest) SetDBClusterId(v string) *DescribeControllerDetectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetEndTime(v string) *DescribeControllerDetectionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetLang(v string) *DescribeControllerDetectionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetOwnerAccount(v string) *DescribeControllerDetectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetOwnerId(v int64) *DescribeControllerDetectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetRegionId(v string) *DescribeControllerDetectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetResourceOwnerAccount(v string) *DescribeControllerDetectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetResourceOwnerId(v int64) *DescribeControllerDetectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeControllerDetectionRequest) SetStartTime(v string) *DescribeControllerDetectionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeControllerDetectionRequest) Validate() error {
	return dara.Validate(s)
}
