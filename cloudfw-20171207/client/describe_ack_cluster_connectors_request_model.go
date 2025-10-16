// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterConnectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeAckClusterConnectorsRequest
	GetClusterId() *string
	SetConnectorName(v string) *DescribeAckClusterConnectorsRequest
	GetConnectorName() *string
	SetLang(v string) *DescribeAckClusterConnectorsRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeAckClusterConnectorsRequest
	GetMemberUid() *string
	SetPageNo(v string) *DescribeAckClusterConnectorsRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeAckClusterConnectorsRequest
	GetPageSize() *string
	SetRegionNo(v string) *DescribeAckClusterConnectorsRequest
	GetRegionNo() *string
	SetVpcId(v string) *DescribeAckClusterConnectorsRequest
	GetVpcId() *string
}

type DescribeAckClusterConnectorsRequest struct {
	// example:
	//
	// 0E0C30C977463****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vpc-uf6b5lyul0x******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAckClusterConnectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAckClusterConnectorsRequest) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *DescribeAckClusterConnectorsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAckClusterConnectorsRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeAckClusterConnectorsRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeAckClusterConnectorsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAckClusterConnectorsRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAckClusterConnectorsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAckClusterConnectorsRequest) SetClusterId(v string) *DescribeAckClusterConnectorsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) SetConnectorName(v string) *DescribeAckClusterConnectorsRequest {
	s.ConnectorName = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) SetLang(v string) *DescribeAckClusterConnectorsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) SetMemberUid(v string) *DescribeAckClusterConnectorsRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) SetPageNo(v string) *DescribeAckClusterConnectorsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) SetPageSize(v string) *DescribeAckClusterConnectorsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) SetRegionNo(v string) *DescribeAckClusterConnectorsRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) SetVpcId(v string) *DescribeAckClusterConnectorsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeAckClusterConnectorsRequest) Validate() error {
	return dara.Validate(s)
}
