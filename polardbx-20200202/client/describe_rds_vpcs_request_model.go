// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRdsVpcsRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeRdsVpcsRequest
	GetZoneId() *string
}

type DescribeRdsVpcsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRdsVpcsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRdsVpcsRequest) SetRegionId(v string) *DescribeRdsVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetZoneId(v string) *DescribeRdsVpcsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) Validate() error {
	return dara.Validate(s)
}
