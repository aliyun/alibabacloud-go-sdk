// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectionCountRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeConnectionCountRecordsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeConnectionCountRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeConnectionCountRecordsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeConnectionCountRecordsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeConnectionCountRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeConnectionCountRecordsRequest
	GetResourceOwnerId() *int64
}

type DescribeConnectionCountRecordsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1jj9xqft1po****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeConnectionCountRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionCountRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeConnectionCountRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeConnectionCountRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeConnectionCountRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeConnectionCountRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeConnectionCountRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeConnectionCountRecordsRequest) SetDBClusterId(v string) *DescribeConnectionCountRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetOwnerAccount(v string) *DescribeConnectionCountRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetOwnerId(v int64) *DescribeConnectionCountRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetRegionId(v string) *DescribeConnectionCountRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetResourceOwnerAccount(v string) *DescribeConnectionCountRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetResourceOwnerId(v int64) *DescribeConnectionCountRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) Validate() error {
	return dara.Validate(s)
}
