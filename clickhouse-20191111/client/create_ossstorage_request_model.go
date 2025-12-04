// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOSSStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateOSSStorageRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreateOSSStorageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateOSSStorageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateOSSStorageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateOSSStorageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateOSSStorageRequest
	GetResourceOwnerId() *int64
}

type CreateOSSStorageRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1z3a2hc8dmt****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateOSSStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOSSStorageRequest) GoString() string {
	return s.String()
}

func (s *CreateOSSStorageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateOSSStorageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateOSSStorageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateOSSStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOSSStorageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateOSSStorageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateOSSStorageRequest) SetDBClusterId(v string) *CreateOSSStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateOSSStorageRequest) SetOwnerAccount(v string) *CreateOSSStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateOSSStorageRequest) SetOwnerId(v int64) *CreateOSSStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOSSStorageRequest) SetRegionId(v string) *CreateOSSStorageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOSSStorageRequest) SetResourceOwnerAccount(v string) *CreateOSSStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateOSSStorageRequest) SetResourceOwnerId(v int64) *CreateOSSStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateOSSStorageRequest) Validate() error {
	return dara.Validate(s)
}
