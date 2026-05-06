// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnlineUserCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *DescribeOnlineUserCountRequest
	GetBizType() *int32
	SetOfficeSiteId(v string) *DescribeOnlineUserCountRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeOnlineUserCountRequest
	GetRegionId() *string
	SetSearchRegionId(v string) *DescribeOnlineUserCountRequest
	GetSearchRegionId() *string
}

type DescribeOnlineUserCountRequest struct {
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// cn-shanghai+dir-631324****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
}

func (s DescribeOnlineUserCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineUserCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnlineUserCountRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *DescribeOnlineUserCountRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeOnlineUserCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOnlineUserCountRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeOnlineUserCountRequest) SetBizType(v int32) *DescribeOnlineUserCountRequest {
	s.BizType = &v
	return s
}

func (s *DescribeOnlineUserCountRequest) SetOfficeSiteId(v string) *DescribeOnlineUserCountRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOnlineUserCountRequest) SetRegionId(v string) *DescribeOnlineUserCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOnlineUserCountRequest) SetSearchRegionId(v string) *DescribeOnlineUserCountRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeOnlineUserCountRequest) Validate() error {
	return dara.Validate(s)
}
