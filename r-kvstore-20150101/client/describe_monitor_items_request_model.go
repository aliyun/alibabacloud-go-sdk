// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeMonitorItemsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeMonitorItemsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeMonitorItemsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMonitorItemsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeMonitorItemsRequest
	GetSecurityToken() *string
}

type DescribeMonitorItemsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeMonitorItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeMonitorItemsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMonitorItemsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMonitorItemsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMonitorItemsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeMonitorItemsRequest) SetOwnerAccount(v string) *DescribeMonitorItemsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetOwnerId(v int64) *DescribeMonitorItemsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetResourceOwnerAccount(v string) *DescribeMonitorItemsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetResourceOwnerId(v int64) *DescribeMonitorItemsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetSecurityToken(v string) *DescribeMonitorItemsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeMonitorItemsRequest) Validate() error {
	return dara.Validate(s)
}
