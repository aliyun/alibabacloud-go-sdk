// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachUserENIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DetachUserENIRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DetachUserENIRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachUserENIRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DetachUserENIRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachUserENIRequest
	GetResourceOwnerId() *int64
}

type DetachUserENIRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DetachUserENIRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachUserENIRequest) GoString() string {
	return s.String()
}

func (s *DetachUserENIRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DetachUserENIRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachUserENIRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachUserENIRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachUserENIRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachUserENIRequest) SetDBClusterId(v string) *DetachUserENIRequest {
	s.DBClusterId = &v
	return s
}

func (s *DetachUserENIRequest) SetOwnerAccount(v string) *DetachUserENIRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachUserENIRequest) SetOwnerId(v int64) *DetachUserENIRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachUserENIRequest) SetResourceOwnerAccount(v string) *DetachUserENIRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachUserENIRequest) SetResourceOwnerId(v int64) *DetachUserENIRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachUserENIRequest) Validate() error {
	return dara.Validate(s)
}
