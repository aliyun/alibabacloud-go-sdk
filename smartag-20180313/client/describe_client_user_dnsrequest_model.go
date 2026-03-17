// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientUserDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeClientUserDNSRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeClientUserDNSRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClientUserDNSRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClientUserDNSRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClientUserDNSRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeClientUserDNSRequest
	GetSmartAGId() *string
}

type DescribeClientUserDNSRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-9uyg53s6juhpxv****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeClientUserDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientUserDNSRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientUserDNSRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeClientUserDNSRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClientUserDNSRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClientUserDNSRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClientUserDNSRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClientUserDNSRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeClientUserDNSRequest) SetOwnerAccount(v string) *DescribeClientUserDNSRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClientUserDNSRequest) SetOwnerId(v int64) *DescribeClientUserDNSRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClientUserDNSRequest) SetRegionId(v string) *DescribeClientUserDNSRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClientUserDNSRequest) SetResourceOwnerAccount(v string) *DescribeClientUserDNSRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClientUserDNSRequest) SetResourceOwnerId(v int64) *DescribeClientUserDNSRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClientUserDNSRequest) SetSmartAGId(v string) *DescribeClientUserDNSRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeClientUserDNSRequest) Validate() error {
	return dara.Validate(s)
}
