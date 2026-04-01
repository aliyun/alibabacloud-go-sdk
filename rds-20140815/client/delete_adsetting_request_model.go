// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteADSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteADSettingRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeleteADSettingRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DeleteADSettingRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteADSettingRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteADSettingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteADSettingRequest
	GetResourceOwnerId() *int64
}

type DeleteADSettingRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// The ID of the region.
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

func (s DeleteADSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteADSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteADSettingRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteADSettingRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteADSettingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteADSettingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteADSettingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteADSettingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteADSettingRequest) SetClientToken(v string) *DeleteADSettingRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteADSettingRequest) SetDBInstanceId(v string) *DeleteADSettingRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteADSettingRequest) SetOwnerId(v int64) *DeleteADSettingRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteADSettingRequest) SetRegionId(v string) *DeleteADSettingRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteADSettingRequest) SetResourceOwnerAccount(v string) *DeleteADSettingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteADSettingRequest) SetResourceOwnerId(v int64) *DeleteADSettingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteADSettingRequest) Validate() error {
	return dara.Validate(s)
}
