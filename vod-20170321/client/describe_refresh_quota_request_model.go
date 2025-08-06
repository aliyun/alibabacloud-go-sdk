// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRefreshQuotaRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DescribeRefreshQuotaRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *DescribeRefreshQuotaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *DescribeRefreshQuotaRequest
	GetResourceOwnerId() *string
}

type DescribeRefreshQuotaRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRefreshQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRefreshQuotaRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeRefreshQuotaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRefreshQuotaRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DescribeRefreshQuotaRequest) SetOwnerAccount(v string) *DescribeRefreshQuotaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) SetOwnerId(v string) *DescribeRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) SetResourceOwnerAccount(v string) *DescribeRefreshQuotaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) SetResourceOwnerId(v string) *DescribeRefreshQuotaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) Validate() error {
	return dara.Validate(s)
}
