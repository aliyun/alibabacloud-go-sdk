// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodesParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBNodesParametersRequest
	GetDBClusterId() *string
	SetDBNodeIds(v string) *DescribeDBNodesParametersRequest
	GetDBNodeIds() *string
	SetOwnerAccount(v string) *DescribeDBNodesParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBNodesParametersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBNodesParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBNodesParametersRequest
	GetResourceOwnerId() *int64
}

type DescribeDBNodesParametersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node ID. You can specify multiple node IDs. Separate multiple node IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// pi-****************,pi-****************
	DBNodeIds            *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBNodesParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodesParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBNodesParametersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBNodesParametersRequest) GetDBNodeIds() *string {
	return s.DBNodeIds
}

func (s *DescribeDBNodesParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBNodesParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBNodesParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBNodesParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBNodesParametersRequest) SetDBClusterId(v string) *DescribeDBNodesParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBNodesParametersRequest) SetDBNodeIds(v string) *DescribeDBNodesParametersRequest {
	s.DBNodeIds = &v
	return s
}

func (s *DescribeDBNodesParametersRequest) SetOwnerAccount(v string) *DescribeDBNodesParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBNodesParametersRequest) SetOwnerId(v int64) *DescribeDBNodesParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBNodesParametersRequest) SetResourceOwnerAccount(v string) *DescribeDBNodesParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBNodesParametersRequest) SetResourceOwnerId(v int64) *DescribeDBNodesParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBNodesParametersRequest) Validate() error {
	return dara.Validate(s)
}
