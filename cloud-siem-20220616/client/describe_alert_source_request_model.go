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
	// The end time of the query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The threat level. The value is a JSON array. Valid values:
	//
	// - serious: High.
	//
	// - suspicious: Medium.
	//
	// - remind: Low.
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	Level []*string `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
	// The region where the data management center of Threat Analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The start time of the query. Unit: milliseconds.
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
