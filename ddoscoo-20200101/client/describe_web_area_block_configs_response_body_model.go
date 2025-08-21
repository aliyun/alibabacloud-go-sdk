// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAreaBlockConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreaBlockConfigs(v []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) *DescribeWebAreaBlockConfigsResponseBody
	GetAreaBlockConfigs() []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs
	SetRequestId(v string) *DescribeWebAreaBlockConfigsResponseBody
	GetRequestId() *string
}

type DescribeWebAreaBlockConfigsResponseBody struct {
	// An array that consists of the configurations of the Location Blacklist (Domain Names) policy.
	AreaBlockConfigs []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs `json:"AreaBlockConfigs,omitempty" xml:"AreaBlockConfigs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebAreaBlockConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponseBody) GetAreaBlockConfigs() []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs {
	return s.AreaBlockConfigs
}

func (s *DescribeWebAreaBlockConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebAreaBlockConfigsResponseBody) SetAreaBlockConfigs(v []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) *DescribeWebAreaBlockConfigsResponseBody {
	s.AreaBlockConfigs = v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBody) SetRequestId(v string) *DescribeWebAreaBlockConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs struct {
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The configuration of the blocked locations.
	RegionList []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) GetRegionList() []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList {
	return s.RegionList
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) SetDomain(v string) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs {
	s.Domain = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) SetRegionList(v []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs {
	s.RegionList = v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList struct {
	// Indicates whether the location is blocked. Valid values:
	//
	// 	- **0**: yes
	//
	// 	- **1**: no
	//
	// example:
	//
	// 0
	Block *int32 `json:"Block,omitempty" xml:"Block,omitempty"`
	// The name of the location.
	//
	// example:
	//
	// CN-SHANGHAI
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) GetBlock() *int32 {
	return s.Block
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) GetRegion() *string {
	return s.Region
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) SetBlock(v int32) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList {
	s.Block = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) SetRegion(v string) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList {
	s.Region = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) Validate() error {
	return dara.Validate(s)
}
