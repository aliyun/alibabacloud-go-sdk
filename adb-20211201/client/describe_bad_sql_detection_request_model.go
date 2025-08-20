// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBadSqlDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeBadSqlDetectionRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeBadSqlDetectionRequest
	GetEndTime() *string
	SetLang(v string) *DescribeBadSqlDetectionRequest
	GetLang() *string
	SetOwnerAccount(v string) *DescribeBadSqlDetectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBadSqlDetectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeBadSqlDetectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBadSqlDetectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBadSqlDetectionRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeBadSqlDetectionRequest
	GetStartTime() *string
}

type DescribeBadSqlDetectionRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2024-11-17T02:16Z
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
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-23T02:18Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBadSqlDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBadSqlDetectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeBadSqlDetectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBadSqlDetectionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBadSqlDetectionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBadSqlDetectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBadSqlDetectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBadSqlDetectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBadSqlDetectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBadSqlDetectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBadSqlDetectionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBadSqlDetectionRequest) SetDBClusterId(v string) *DescribeBadSqlDetectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetEndTime(v string) *DescribeBadSqlDetectionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetLang(v string) *DescribeBadSqlDetectionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetOwnerAccount(v string) *DescribeBadSqlDetectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetOwnerId(v int64) *DescribeBadSqlDetectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetRegionId(v string) *DescribeBadSqlDetectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetResourceOwnerAccount(v string) *DescribeBadSqlDetectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetResourceOwnerId(v int64) *DescribeBadSqlDetectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) SetStartTime(v string) *DescribeBadSqlDetectionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBadSqlDetectionRequest) Validate() error {
	return dara.Validate(s)
}
