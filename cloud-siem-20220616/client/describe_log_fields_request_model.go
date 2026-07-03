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
	// The log source for the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The log source for the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The region where the threat analysis Management Hub is located. Select the region of the Management Hub based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose view the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
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
