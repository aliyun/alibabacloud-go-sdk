// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckKMSAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckKMSAuthorizedRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CheckKMSAuthorizedRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckKMSAuthorizedRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckKMSAuthorizedRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckKMSAuthorizedRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckKMSAuthorizedRequest
	GetResourceOwnerId() *int64
	SetTDERegion(v string) *CheckKMSAuthorizedRequest
	GetTDERegion() *string
}

type CheckKMSAuthorizedRequest struct {
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/2319131.html) operation to query information about all clusters in the destination region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > Call the [DescribeRegions](https://help.aliyun.com/document_detail/2319134.html) operation to query information about the available regions of the destination account, including region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region where the transparent data encryption (TDE) key is located.
	//
	// example:
	//
	// cn-beijing
	TDERegion *string `json:"TDERegion,omitempty" xml:"TDERegion,omitempty"`
}

func (s CheckKMSAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckKMSAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckKMSAuthorizedRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckKMSAuthorizedRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckKMSAuthorizedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckKMSAuthorizedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckKMSAuthorizedRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckKMSAuthorizedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckKMSAuthorizedRequest) GetTDERegion() *string {
	return s.TDERegion
}

func (s *CheckKMSAuthorizedRequest) SetDBClusterId(v string) *CheckKMSAuthorizedRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckKMSAuthorizedRequest) SetOwnerAccount(v string) *CheckKMSAuthorizedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckKMSAuthorizedRequest) SetOwnerId(v int64) *CheckKMSAuthorizedRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckKMSAuthorizedRequest) SetRegionId(v string) *CheckKMSAuthorizedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckKMSAuthorizedRequest) SetResourceOwnerAccount(v string) *CheckKMSAuthorizedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckKMSAuthorizedRequest) SetResourceOwnerId(v int64) *CheckKMSAuthorizedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckKMSAuthorizedRequest) SetTDERegion(v string) *CheckKMSAuthorizedRequest {
	s.TDERegion = &v
	return s
}

func (s *CheckKMSAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
