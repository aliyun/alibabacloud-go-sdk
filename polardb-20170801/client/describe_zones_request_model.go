// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeZonesRequest
	GetEngine() *string
	SetExtra(v string) *DescribeZonesRequest
	GetExtra() *string
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeZonesRequest
	GetResourceOwnerId() *int64
}

type DescribeZonesRequest struct {
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// local
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeZonesRequest) GetExtra() *string {
	return s.Extra
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeZonesRequest) SetEngine(v string) *DescribeZonesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeZonesRequest) SetExtra(v string) *DescribeZonesRequest {
	s.Extra = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerId(v int64) *DescribeZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
