// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeSupabaseInstanceAttributeRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeSupabaseInstanceAttributeRequest
	GetRegionId() *string
}

type DescribeSupabaseInstanceAttributeRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSupabaseInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstanceAttributeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSupabaseInstanceAttributeRequest) SetDBInstanceName(v string) *DescribeSupabaseInstanceAttributeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeRequest) SetRegionId(v string) *DescribeSupabaseInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
