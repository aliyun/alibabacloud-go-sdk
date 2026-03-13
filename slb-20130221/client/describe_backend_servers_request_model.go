// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeBackendServersRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeBackendServersRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackendServersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeBackendServersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackendServersRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DescribeBackendServersRequest
	GetTags() *string
	SetAccessKeyId(v string) *DescribeBackendServersRequest
	GetAccessKeyId() *string
}

type DescribeBackendServersRequest struct {
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// This parameter is required.
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s DescribeBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackendServersRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackendServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackendServersRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeBackendServersRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeBackendServersRequest) SetListenerPort(v int32) *DescribeBackendServersRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeBackendServersRequest) SetLoadBalancerId(v string) *DescribeBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeBackendServersRequest) SetOwnerAccount(v string) *DescribeBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackendServersRequest) SetOwnerId(v int64) *DescribeBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackendServersRequest) SetRegionId(v string) *DescribeBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackendServersRequest) SetResourceOwnerAccount(v string) *DescribeBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackendServersRequest) SetResourceOwnerId(v int64) *DescribeBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackendServersRequest) SetTags(v string) *DescribeBackendServersRequest {
	s.Tags = &v
	return s
}

func (s *DescribeBackendServersRequest) SetAccessKeyId(v string) *DescribeBackendServersRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
