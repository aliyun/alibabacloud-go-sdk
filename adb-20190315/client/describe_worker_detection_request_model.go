// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkerDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeWorkerDetectionRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeWorkerDetectionRequest
	GetEndTime() *string
	SetLang(v string) *DescribeWorkerDetectionRequest
	GetLang() *string
	SetOwnerAccount(v string) *DescribeWorkerDetectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeWorkerDetectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeWorkerDetectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeWorkerDetectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeWorkerDetectionRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeWorkerDetectionRequest
	GetStartTime() *string
}

type DescribeWorkerDetectionRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-xxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2024-04-25T02:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
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
	// >  You can query data only within the last 15 days.
	//
	// example:
	//
	// 2011-06-01T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeWorkerDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeWorkerDetectionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeWorkerDetectionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWorkerDetectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeWorkerDetectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeWorkerDetectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWorkerDetectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeWorkerDetectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeWorkerDetectionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeWorkerDetectionRequest) SetDBClusterId(v string) *DescribeWorkerDetectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetEndTime(v string) *DescribeWorkerDetectionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetLang(v string) *DescribeWorkerDetectionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetOwnerAccount(v string) *DescribeWorkerDetectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetOwnerId(v int64) *DescribeWorkerDetectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetRegionId(v string) *DescribeWorkerDetectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetResourceOwnerAccount(v string) *DescribeWorkerDetectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetResourceOwnerId(v int64) *DescribeWorkerDetectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) SetStartTime(v string) *DescribeWorkerDetectionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeWorkerDetectionRequest) Validate() error {
	return dara.Validate(s)
}
