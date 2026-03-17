// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSagInstanceToCcnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnInstanceId(v string) *GrantSagInstanceToCcnRequest
	GetCcnInstanceId() *string
	SetCcnUid(v int64) *GrantSagInstanceToCcnRequest
	GetCcnUid() *int64
	SetGrantTrafficService(v bool) *GrantSagInstanceToCcnRequest
	GetGrantTrafficService() *bool
	SetOwnerAccount(v string) *GrantSagInstanceToCcnRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantSagInstanceToCcnRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GrantSagInstanceToCcnRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GrantSagInstanceToCcnRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantSagInstanceToCcnRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *GrantSagInstanceToCcnRequest
	GetSmartAGId() *string
}

type GrantSagInstanceToCcnRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-n2935s1mnwv8i*****
	CcnInstanceId *string `json:"CcnInstanceId,omitempty" xml:"CcnInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the CCN instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	CcnUid *int64 `json:"CcnUid,omitempty" xml:"CcnUid,omitempty"`
	// Specifies whether to grant the CCN instance permissions to manage network traffic from the SAG instance.
	//
	// After the CCN instance is granted the permissions, the CCN instance can redirect the network traffic sent from the SAG instance to the Internet to Cloud Security Access Service (CSAS) for security audit.
	//
	// 	- **true**: grants permissions.
	//
	// 	- **false**: does not grant permissions.
	//
	// >  If you set the value to true and the SAG instance connected to the CCN instance has the secure rerouting feature enabled, you cannot revoke the permissions.
	//
	// example:
	//
	// true
	GrantTrafficService *bool   `json:"GrantTrafficService,omitempty" xml:"GrantTrafficService,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
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
	// sag-tzirqx07bvcngm****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s GrantSagInstanceToCcnRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantSagInstanceToCcnRequest) GoString() string {
	return s.String()
}

func (s *GrantSagInstanceToCcnRequest) GetCcnInstanceId() *string {
	return s.CcnInstanceId
}

func (s *GrantSagInstanceToCcnRequest) GetCcnUid() *int64 {
	return s.CcnUid
}

func (s *GrantSagInstanceToCcnRequest) GetGrantTrafficService() *bool {
	return s.GrantTrafficService
}

func (s *GrantSagInstanceToCcnRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantSagInstanceToCcnRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantSagInstanceToCcnRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantSagInstanceToCcnRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantSagInstanceToCcnRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantSagInstanceToCcnRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *GrantSagInstanceToCcnRequest) SetCcnInstanceId(v string) *GrantSagInstanceToCcnRequest {
	s.CcnInstanceId = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetCcnUid(v int64) *GrantSagInstanceToCcnRequest {
	s.CcnUid = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetGrantTrafficService(v bool) *GrantSagInstanceToCcnRequest {
	s.GrantTrafficService = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetOwnerAccount(v string) *GrantSagInstanceToCcnRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetOwnerId(v int64) *GrantSagInstanceToCcnRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetRegionId(v string) *GrantSagInstanceToCcnRequest {
	s.RegionId = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetResourceOwnerAccount(v string) *GrantSagInstanceToCcnRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetResourceOwnerId(v int64) *GrantSagInstanceToCcnRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) SetSmartAGId(v string) *GrantSagInstanceToCcnRequest {
	s.SmartAGId = &v
	return s
}

func (s *GrantSagInstanceToCcnRequest) Validate() error {
	return dara.Validate(s)
}
