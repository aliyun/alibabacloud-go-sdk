// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbnormalPatternDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAbnormalPatternDetectionRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeAbnormalPatternDetectionRequest
	GetEndTime() *string
	SetLang(v string) *DescribeAbnormalPatternDetectionRequest
	GetLang() *string
	SetOwnerAccount(v string) *DescribeAbnormalPatternDetectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAbnormalPatternDetectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAbnormalPatternDetectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAbnormalPatternDetectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAbnormalPatternDetectionRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeAbnormalPatternDetectionRequest
	GetStartTime() *string
}

type DescribeAbnormalPatternDetectionRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2024-11-17T02:16Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language. Valid values:
	//
	// 	- **zh*	- (default): simplified Chinese.
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
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2021-09-30T00:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAbnormalPatternDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalPatternDetectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalPatternDetectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAbnormalPatternDetectionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAbnormalPatternDetectionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAbnormalPatternDetectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAbnormalPatternDetectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAbnormalPatternDetectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAbnormalPatternDetectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAbnormalPatternDetectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAbnormalPatternDetectionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAbnormalPatternDetectionRequest) SetDBClusterId(v string) *DescribeAbnormalPatternDetectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetEndTime(v string) *DescribeAbnormalPatternDetectionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetLang(v string) *DescribeAbnormalPatternDetectionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetOwnerAccount(v string) *DescribeAbnormalPatternDetectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetOwnerId(v int64) *DescribeAbnormalPatternDetectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetRegionId(v string) *DescribeAbnormalPatternDetectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetResourceOwnerAccount(v string) *DescribeAbnormalPatternDetectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetResourceOwnerId(v int64) *DescribeAbnormalPatternDetectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) SetStartTime(v string) *DescribeAbnormalPatternDetectionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionRequest) Validate() error {
	return dara.Validate(s)
}
