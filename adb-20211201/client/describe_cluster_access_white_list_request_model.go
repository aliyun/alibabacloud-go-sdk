// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAccessWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeClusterAccessWhiteListRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeClusterAccessWhiteListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterAccessWhiteListRequest
	GetResourceOwnerAccount() *string
}

type DescribeClusterAccessWhiteListRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeClusterAccessWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeClusterAccessWhiteListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterAccessWhiteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterAccessWhiteListRequest) SetDBClusterId(v string) *DescribeClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterAccessWhiteListRequest) SetRegionId(v string) *DescribeClusterAccessWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterAccessWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
