// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAIDBClusterAttributeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAIDBClusterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAIDBClusterAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAIDBClusterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAIDBClusterAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeAIDBClusterAttributeRequest struct {
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

func (s DescribeAIDBClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAIDBClusterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAIDBClusterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAIDBClusterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAIDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeAIDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeRequest) SetOwnerAccount(v string) *DescribeAIDBClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterAttributeRequest) SetOwnerId(v int64) *DescribeAIDBClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAIDBClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeAIDBClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
