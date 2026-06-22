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
	// The ACK cluster ID.
	//
	// example:
	//
	// 0E0C30C977463****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the ACK cluster connector.
	//
	// example:
	//
	// test
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// The language in which the unhealthy reason of the ACK cluster connector health status is displayed.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Alibaba Cloud UID of the account to which the ACK cluster resource belongs.
	//
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceRegionList](~~DescribeAccessInstanceRegionList~~): Queries the list of synchronization node regions.
	//
	// > For more information about the regions supported by ACK cluster connectors in Cloud Firewall, see [ACK cluster synchronization nodes](https://help.aliyun.com/document_detail/2865120.html).
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The instance ID of the VPC-connected instance to which the ACK cluster belongs.
	//
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
