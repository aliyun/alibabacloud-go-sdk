// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransferHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTransferHistoryRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeTransferHistoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTransferHistoryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTransferHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTransferHistoryRequest
	GetResourceOwnerId() *int64
}

type DescribeTransferHistoryRequest struct {
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

func (s DescribeTransferHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTransferHistoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTransferHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTransferHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTransferHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTransferHistoryRequest) SetDBClusterId(v string) *DescribeTransferHistoryRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetOwnerAccount(v string) *DescribeTransferHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetOwnerId(v int64) *DescribeTransferHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetResourceOwnerAccount(v string) *DescribeTransferHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetResourceOwnerId(v int64) *DescribeTransferHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTransferHistoryRequest) Validate() error {
	return dara.Validate(s)
}
