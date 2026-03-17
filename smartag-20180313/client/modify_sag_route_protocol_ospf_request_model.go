// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRouteProtocolOspfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v int32) *ModifySagRouteProtocolOspfRequest
	GetAreaId() *int32
	SetAuthenticationType(v string) *ModifySagRouteProtocolOspfRequest
	GetAuthenticationType() *string
	SetDeadTime(v int32) *ModifySagRouteProtocolOspfRequest
	GetDeadTime() *int32
	SetHelloTime(v int32) *ModifySagRouteProtocolOspfRequest
	GetHelloTime() *int32
	SetMd5Key(v string) *ModifySagRouteProtocolOspfRequest
	GetMd5Key() *string
	SetMd5KeyId(v int32) *ModifySagRouteProtocolOspfRequest
	GetMd5KeyId() *int32
	SetOwnerAccount(v string) *ModifySagRouteProtocolOspfRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagRouteProtocolOspfRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifySagRouteProtocolOspfRequest
	GetPassword() *string
	SetRegionId(v string) *ModifySagRouteProtocolOspfRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagRouteProtocolOspfRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagRouteProtocolOspfRequest
	GetResourceOwnerId() *int64
	SetRouterId(v string) *ModifySagRouteProtocolOspfRequest
	GetRouterId() *string
	SetSmartAGId(v string) *ModifySagRouteProtocolOspfRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagRouteProtocolOspfRequest
	GetSmartAGSn() *string
}

type ModifySagRouteProtocolOspfRequest struct {
	// The ID of the OSPF area.
	//
	// Valid values: **1 to 2147483647**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	AreaId *int32 `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The authentication method. Valid values:
	//
	// 	- **NONE**: does not enable authentication.
	//
	// 	- **CLEARTEXT**: uses plaintext authentication. You must enter the plaintext password.
	//
	// 	- **MD5**: uses MD5 authentication. You must enter the MD5 key ID and the MD5 key.
	//
	// This parameter is required.
	//
	// example:
	//
	// NONE
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The timeout period of OSPF. Unit: seconds.
	//
	// Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	DeadTime *int32 `json:"DeadTime,omitempty" xml:"DeadTime,omitempty"`
	// The time interval at which Hello packets are sent. Unit: seconds.
	//
	// Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	HelloTime *int32 `json:"HelloTime,omitempty" xml:"HelloTime,omitempty"`
	// The MD5 key value.
	//
	// Valid values: **1 to 47**.
	//
	// >  This parameter is required only if **AuthenticationType*	- is set to **MD5**.
	//
	// example:
	//
	// 5
	Md5Key *string `json:"Md5Key,omitempty" xml:"Md5Key,omitempty"`
	// The ID of the MD5 key.
	//
	// Valid values: **1 to 2147483647**.
	//
	// >  This parameter is required only if **AuthenticationType*	- is set to **MD5**.
	//
	// example:
	//
	// 7
	Md5KeyId     *int32  `json:"Md5KeyId,omitempty" xml:"Md5KeyId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The plaintext password.
	//
	// The password must be 1 to 8 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// >  This parameter is required only if **AuthenticationType*	- is set to **CLEARTEXT**.
	//
	// example:
	//
	// 1212****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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
	// The ID of the router that has OSPF enabled.
	//
	// Set the value to an IPv4 address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30***
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s ModifySagRouteProtocolOspfRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRouteProtocolOspfRequest) GoString() string {
	return s.String()
}

func (s *ModifySagRouteProtocolOspfRequest) GetAreaId() *int32 {
	return s.AreaId
}

func (s *ModifySagRouteProtocolOspfRequest) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *ModifySagRouteProtocolOspfRequest) GetDeadTime() *int32 {
	return s.DeadTime
}

func (s *ModifySagRouteProtocolOspfRequest) GetHelloTime() *int32 {
	return s.HelloTime
}

func (s *ModifySagRouteProtocolOspfRequest) GetMd5Key() *string {
	return s.Md5Key
}

func (s *ModifySagRouteProtocolOspfRequest) GetMd5KeyId() *int32 {
	return s.Md5KeyId
}

func (s *ModifySagRouteProtocolOspfRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagRouteProtocolOspfRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagRouteProtocolOspfRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifySagRouteProtocolOspfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagRouteProtocolOspfRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagRouteProtocolOspfRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagRouteProtocolOspfRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *ModifySagRouteProtocolOspfRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagRouteProtocolOspfRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagRouteProtocolOspfRequest) SetAreaId(v int32) *ModifySagRouteProtocolOspfRequest {
	s.AreaId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetAuthenticationType(v string) *ModifySagRouteProtocolOspfRequest {
	s.AuthenticationType = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetDeadTime(v int32) *ModifySagRouteProtocolOspfRequest {
	s.DeadTime = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetHelloTime(v int32) *ModifySagRouteProtocolOspfRequest {
	s.HelloTime = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetMd5Key(v string) *ModifySagRouteProtocolOspfRequest {
	s.Md5Key = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetMd5KeyId(v int32) *ModifySagRouteProtocolOspfRequest {
	s.Md5KeyId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetOwnerAccount(v string) *ModifySagRouteProtocolOspfRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetOwnerId(v int64) *ModifySagRouteProtocolOspfRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetPassword(v string) *ModifySagRouteProtocolOspfRequest {
	s.Password = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetRegionId(v string) *ModifySagRouteProtocolOspfRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetResourceOwnerAccount(v string) *ModifySagRouteProtocolOspfRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetResourceOwnerId(v int64) *ModifySagRouteProtocolOspfRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetRouterId(v string) *ModifySagRouteProtocolOspfRequest {
	s.RouterId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetSmartAGId(v string) *ModifySagRouteProtocolOspfRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) SetSmartAGSn(v string) *ModifySagRouteProtocolOspfRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagRouteProtocolOspfRequest) Validate() error {
	return dara.Validate(s)
}
