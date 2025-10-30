// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
	SetServiceResourceType(v string) *DescribeZonesRequest
	GetServiceResourceType() *string
}

type DescribeZonesRequest struct {
	// The region ID of the zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetServiceResourceType(v string) *DescribeZonesRequest {
	s.ServiceResourceType = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
