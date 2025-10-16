// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeAckClustersRequest
	GetClusterId() *string
	SetClusterName(v string) *DescribeAckClustersRequest
	GetClusterName() *string
	SetClusterSpec(v string) *DescribeAckClustersRequest
	GetClusterSpec() *string
	SetConnectorStatus(v string) *DescribeAckClustersRequest
	GetConnectorStatus() *string
	SetMemberUid(v string) *DescribeAckClustersRequest
	GetMemberUid() *string
	SetPageNo(v string) *DescribeAckClustersRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeAckClustersRequest
	GetPageSize() *string
	SetRegionNo(v string) *DescribeAckClustersRequest
	GetRegionNo() *string
}

type DescribeAckClustersRequest struct {
	// example:
	//
	// cb0f5640b1b2d404cad6ba21509d7847b
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// ack-cluster-name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// example:
	//
	// ready
	ConnectorStatus *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeAckClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeAckClustersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAckClustersRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAckClustersRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeAckClustersRequest) GetConnectorStatus() *string {
	return s.ConnectorStatus
}

func (s *DescribeAckClustersRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeAckClustersRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeAckClustersRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAckClustersRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAckClustersRequest) SetClusterId(v string) *DescribeAckClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeAckClustersRequest) SetClusterName(v string) *DescribeAckClustersRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeAckClustersRequest) SetClusterSpec(v string) *DescribeAckClustersRequest {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeAckClustersRequest) SetConnectorStatus(v string) *DescribeAckClustersRequest {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeAckClustersRequest) SetMemberUid(v string) *DescribeAckClustersRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAckClustersRequest) SetPageNo(v string) *DescribeAckClustersRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAckClustersRequest) SetPageSize(v string) *DescribeAckClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAckClustersRequest) SetRegionNo(v string) *DescribeAckClustersRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAckClustersRequest) Validate() error {
	return dara.Validate(s)
}
