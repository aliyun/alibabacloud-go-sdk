// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachUserENIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AttachUserENIRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *AttachUserENIRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachUserENIRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AttachUserENIRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachUserENIRequest
	GetResourceOwnerId() *int64
}

type AttachUserENIRequest struct {
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

func (s AttachUserENIRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachUserENIRequest) GoString() string {
	return s.String()
}

func (s *AttachUserENIRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AttachUserENIRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachUserENIRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachUserENIRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachUserENIRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachUserENIRequest) SetDBClusterId(v string) *AttachUserENIRequest {
	s.DBClusterId = &v
	return s
}

func (s *AttachUserENIRequest) SetOwnerAccount(v string) *AttachUserENIRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachUserENIRequest) SetOwnerId(v int64) *AttachUserENIRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachUserENIRequest) SetResourceOwnerAccount(v string) *AttachUserENIRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachUserENIRequest) SetResourceOwnerId(v int64) *AttachUserENIRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachUserENIRequest) Validate() error {
	return dara.Validate(s)
}
