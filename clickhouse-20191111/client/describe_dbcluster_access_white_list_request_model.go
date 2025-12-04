// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAccessWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterAccessWhiteListRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterAccessWhiteListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAccessWhiteListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterAccessWhiteListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterAccessWhiteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterAccessWhiteListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
