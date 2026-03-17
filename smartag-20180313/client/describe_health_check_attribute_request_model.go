// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHcInstanceId(v string) *DescribeHealthCheckAttributeRequest
	GetHcInstanceId() *string
	SetOwnerAccount(v string) *DescribeHealthCheckAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHealthCheckAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHealthCheckAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHealthCheckAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHealthCheckAttributeRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeHealthCheckAttributeRequest
	GetSmartAGId() *string
}

type DescribeHealthCheckAttributeRequest struct {
	// The ID of the health check instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc-1k4ucuq77b56x4****
	HcInstanceId *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeHealthCheckAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckAttributeRequest) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *DescribeHealthCheckAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHealthCheckAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHealthCheckAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHealthCheckAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHealthCheckAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHealthCheckAttributeRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeHealthCheckAttributeRequest) SetHcInstanceId(v string) *DescribeHealthCheckAttributeRequest {
	s.HcInstanceId = &v
	return s
}

func (s *DescribeHealthCheckAttributeRequest) SetOwnerAccount(v string) *DescribeHealthCheckAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHealthCheckAttributeRequest) SetOwnerId(v int64) *DescribeHealthCheckAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHealthCheckAttributeRequest) SetRegionId(v string) *DescribeHealthCheckAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHealthCheckAttributeRequest) SetResourceOwnerAccount(v string) *DescribeHealthCheckAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHealthCheckAttributeRequest) SetResourceOwnerId(v int64) *DescribeHealthCheckAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHealthCheckAttributeRequest) SetSmartAGId(v string) *DescribeHealthCheckAttributeRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeHealthCheckAttributeRequest) Validate() error {
	return dara.Validate(s)
}
