// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CheckAccountNameRequest
	GetAccountName() *string
	SetDBClusterId(v string) *CheckAccountNameRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CheckAccountNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckAccountNameRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckAccountNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckAccountNameRequest
	GetResourceOwnerId() *int64
}

type CheckAccountNameRequest struct {
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
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

func (s CheckAccountNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountNameRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CheckAccountNameRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckAccountNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckAccountNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckAccountNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckAccountNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckAccountNameRequest) SetAccountName(v string) *CheckAccountNameRequest {
	s.AccountName = &v
	return s
}

func (s *CheckAccountNameRequest) SetDBClusterId(v string) *CheckAccountNameRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckAccountNameRequest) SetOwnerAccount(v string) *CheckAccountNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckAccountNameRequest) SetOwnerId(v int64) *CheckAccountNameRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckAccountNameRequest) SetResourceOwnerAccount(v string) *CheckAccountNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckAccountNameRequest) SetResourceOwnerId(v int64) *CheckAccountNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckAccountNameRequest) Validate() error {
	return dara.Validate(s)
}
