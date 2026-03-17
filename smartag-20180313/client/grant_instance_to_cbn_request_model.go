// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToCbnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnInstanceId(v string) *GrantInstanceToCbnRequest
	GetCcnInstanceId() *string
	SetCenInstanceId(v string) *GrantInstanceToCbnRequest
	GetCenInstanceId() *string
	SetCenUid(v int64) *GrantInstanceToCbnRequest
	GetCenUid() *int64
	SetGrantTrafficService(v bool) *GrantInstanceToCbnRequest
	GetGrantTrafficService() *bool
	SetOwnerAccount(v string) *GrantInstanceToCbnRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantInstanceToCbnRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GrantInstanceToCbnRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GrantInstanceToCbnRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantInstanceToCbnRequest
	GetResourceOwnerId() *int64
}

type GrantInstanceToCbnRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-n2935s1mnwv8i*****
	CcnInstanceId *string `json:"CcnInstanceId,omitempty" xml:"CcnInstanceId,omitempty"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jm*****
	CenInstanceId *string `json:"CenInstanceId,omitempty" xml:"CenInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1250123456123456
	CenUid *int64 `json:"CenUid,omitempty" xml:"CenUid,omitempty"`
	// Specifies whether to grant the CEN instance permissions to manage network traffic from the CCN instance. Valid values:
	//
	// 	- **true**: grants permissions.
	//
	// 	- **false**: does not grant permissions. This is the default value.
	//
	// >  If you set the value to true and the SAG instance connected to the CCN instance has the secure rerouting feature enabled, you cannot revoke the permissions.
	//
	// example:
	//
	// true
	GrantTrafficService *bool   `json:"GrantTrafficService,omitempty" xml:"GrantTrafficService,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CCN instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.htmll) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantInstanceToCbnRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToCbnRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToCbnRequest) GetCcnInstanceId() *string {
	return s.CcnInstanceId
}

func (s *GrantInstanceToCbnRequest) GetCenInstanceId() *string {
	return s.CenInstanceId
}

func (s *GrantInstanceToCbnRequest) GetCenUid() *int64 {
	return s.CenUid
}

func (s *GrantInstanceToCbnRequest) GetGrantTrafficService() *bool {
	return s.GrantTrafficService
}

func (s *GrantInstanceToCbnRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantInstanceToCbnRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantInstanceToCbnRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantInstanceToCbnRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantInstanceToCbnRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantInstanceToCbnRequest) SetCcnInstanceId(v string) *GrantInstanceToCbnRequest {
	s.CcnInstanceId = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetCenInstanceId(v string) *GrantInstanceToCbnRequest {
	s.CenInstanceId = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetCenUid(v int64) *GrantInstanceToCbnRequest {
	s.CenUid = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetGrantTrafficService(v bool) *GrantInstanceToCbnRequest {
	s.GrantTrafficService = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetOwnerAccount(v string) *GrantInstanceToCbnRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetOwnerId(v int64) *GrantInstanceToCbnRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetRegionId(v string) *GrantInstanceToCbnRequest {
	s.RegionId = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetResourceOwnerAccount(v string) *GrantInstanceToCbnRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantInstanceToCbnRequest) SetResourceOwnerId(v int64) *GrantInstanceToCbnRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantInstanceToCbnRequest) Validate() error {
	return dara.Validate(s)
}
