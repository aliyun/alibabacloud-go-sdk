// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterMonitorRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterMonitorRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterMonitorRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterMonitorRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterMonitorRequest) SetDBClusterId(v string) *DescribeDBClusterMonitorRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetOwnerAccount(v string) *DescribeDBClusterMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetOwnerId(v int64) *DescribeDBClusterMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetResourceOwnerId(v int64) *DescribeDBClusterMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) Validate() error {
	return dara.Validate(s)
}
