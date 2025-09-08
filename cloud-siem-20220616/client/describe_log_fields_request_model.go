// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogSource(v string) *DescribeLogFieldsRequest
	GetLogSource() *string
	SetLogType(v string) *DescribeLogFieldsRequest
	GetLogType() *string
	SetRegionId(v string) *DescribeLogFieldsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeLogFieldsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeLogFieldsRequest
	GetRoleType() *int32
}

type DescribeLogFieldsRequest struct {
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeLogFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsRequest) GetLogSource() *string {
	return s.LogSource
}

func (s *DescribeLogFieldsRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeLogFieldsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogFieldsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeLogFieldsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeLogFieldsRequest) SetLogSource(v string) *DescribeLogFieldsRequest {
	s.LogSource = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetLogType(v string) *DescribeLogFieldsRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetRegionId(v string) *DescribeLogFieldsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetRoleFor(v int64) *DescribeLogFieldsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetRoleType(v int32) *DescribeLogFieldsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeLogFieldsRequest) Validate() error {
	return dara.Validate(s)
}
