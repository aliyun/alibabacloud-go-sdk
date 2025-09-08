// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBuyStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeUserBuyStatusRequest
	GetRegionId() *string
	SetSubUserId(v int64) *DescribeUserBuyStatusRequest
	GetSubUserId() *int64
}

type DescribeUserBuyStatusRequest struct {
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 123XXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeUserBuyStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBuyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserBuyStatusRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeUserBuyStatusRequest) SetRegionId(v string) *DescribeUserBuyStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserBuyStatusRequest) SetSubUserId(v int64) *DescribeUserBuyStatusRequest {
	s.SubUserId = &v
	return s
}

func (s *DescribeUserBuyStatusRequest) Validate() error {
	return dara.Validate(s)
}
