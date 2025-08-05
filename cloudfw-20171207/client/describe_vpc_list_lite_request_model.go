// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcListLiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcListLiteRequest
	GetLang() *string
	SetRegionNo(v string) *DescribeVpcListLiteRequest
	GetRegionNo() *string
	SetSourceIp(v string) *DescribeVpcListLiteRequest
	GetSourceIp() *string
	SetVpcId(v string) *DescribeVpcListLiteRequest
	GetVpcId() *string
	SetVpcName(v string) *DescribeVpcListLiteRequest
	GetVpcName() *string
}

type DescribeVpcListLiteRequest struct {
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
	// The region ID of the VPC.
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 58.34.174.194
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// vpc-shanghai
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcListLiteRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcListLiteRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcListLiteRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcListLiteRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeVpcListLiteRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcListLiteRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcListLiteRequest) SetLang(v string) *DescribeVpcListLiteRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetRegionNo(v string) *DescribeVpcListLiteRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetSourceIp(v string) *DescribeVpcListLiteRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetVpcId(v string) *DescribeVpcListLiteRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetVpcName(v string) *DescribeVpcListLiteRequest {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcListLiteRequest) Validate() error {
	return dara.Validate(s)
}
