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
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information of all clusters that are deployed in a specific region, such as the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all regions that are available for your account, such as the region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region in which the TDE key resides.
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
