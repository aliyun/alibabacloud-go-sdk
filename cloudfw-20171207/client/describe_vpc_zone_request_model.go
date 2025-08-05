// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *DescribeVpcZoneRequest
	GetEnvironment() *string
	SetLang(v string) *DescribeVpcZoneRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcZoneRequest
	GetMemberUid() *string
	SetRegionNo(v string) *DescribeVpcZoneRequest
	GetRegionNo() *string
}

type DescribeVpcZoneRequest struct {
	// The environment. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **TransitRouter**
	//
	// example:
	//
	// VPC
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member in Cloud Firewall.
	//
	// example:
	//
	// 1415189284827022
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeVpcZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcZoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *DescribeVpcZoneRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcZoneRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcZoneRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcZoneRequest) SetEnvironment(v string) *DescribeVpcZoneRequest {
	s.Environment = &v
	return s
}

func (s *DescribeVpcZoneRequest) SetLang(v string) *DescribeVpcZoneRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcZoneRequest) SetMemberUid(v string) *DescribeVpcZoneRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcZoneRequest) SetRegionNo(v string) *DescribeVpcZoneRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcZoneRequest) Validate() error {
	return dara.Validate(s)
}
