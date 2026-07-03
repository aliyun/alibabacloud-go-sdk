// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeAlertSourceRequest
	GetEndTime() *int64
	SetLevel(v []*string) *DescribeAlertSourceRequest
	GetLevel() []*string
	SetRegionId(v string) *DescribeAlertSourceRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertSourceRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertSourceRequest
	GetRoleType() *int32
	SetStartTime(v int64) *DescribeAlertSourceRequest
	GetStartTime() *int64
}

type DescribeAlertSourceRequest struct {
	// The end of the query time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The threat levels. Valid values:
	//
	// - `serious`: High
	//
	// - `suspicious`: Medium
	//
	// - `remind`: Low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	Level []*string `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
	// The region of the data management center for threat analysis. Select the data management center that corresponds to the region where your assets are located. Valid values:
	//
	// - `cn-hangzhou`: for assets in the Chinese mainland and Hong Kong (China).
	//
	// - `ap-southeast-1`: for assets in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose data you want to view. An administrator uses this parameter to view data from the perspective of a specific member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view. Valid values:
	//
	// - `0`: View data for the current Alibaba Cloud account.
	//
	// - `1`: View data for all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The start of the query time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertSourceRequest) GetLevel() []*string {
	return s.Level
}

func (s *DescribeAlertSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertSourceRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertSourceRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertSourceRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertSourceRequest) SetEndTime(v int64) *DescribeAlertSourceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetLevel(v []*string) *DescribeAlertSourceRequest {
	s.Level = v
	return s
}

func (s *DescribeAlertSourceRequest) SetRegionId(v string) *DescribeAlertSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetRoleFor(v int64) *DescribeAlertSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetRoleType(v int32) *DescribeAlertSourceRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetStartTime(v int64) *DescribeAlertSourceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertSourceRequest) Validate() error {
	return dara.Validate(s)
}
