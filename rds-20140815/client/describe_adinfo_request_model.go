// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeADInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeADInfoRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeADInfoRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeADInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeADInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeADInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeADInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeADInfoRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1k8s41l2o52****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeADInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeADInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeADInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeADInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeADInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeADInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeADInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeADInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeADInfoRequest) SetClientToken(v string) *DescribeADInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeADInfoRequest) SetDBInstanceId(v string) *DescribeADInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeADInfoRequest) SetOwnerId(v int64) *DescribeADInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeADInfoRequest) SetRegionId(v string) *DescribeADInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeADInfoRequest) SetResourceOwnerAccount(v string) *DescribeADInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeADInfoRequest) SetResourceOwnerId(v int64) *DescribeADInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeADInfoRequest) Validate() error {
	return dara.Validate(s)
}
