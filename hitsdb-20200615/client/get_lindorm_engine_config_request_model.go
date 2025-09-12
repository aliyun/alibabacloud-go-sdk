// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngineType(v string) *GetLindormEngineConfigRequest
	GetEngineType() *string
	SetInstanceId(v string) *GetLindormEngineConfigRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormEngineConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormEngineConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetLindormEngineConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetLindormEngineConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormEngineConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormEngineConfigRequest
	GetSecurityToken() *string
}

type GetLindormEngineConfigRequest struct {
	// This parameter is required.
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *GetLindormEngineConfigRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *GetLindormEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormEngineConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormEngineConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormEngineConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormEngineConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormEngineConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormEngineConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormEngineConfigRequest) SetEngineType(v string) *GetLindormEngineConfigRequest {
	s.EngineType = &v
	return s
}

func (s *GetLindormEngineConfigRequest) SetInstanceId(v string) *GetLindormEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormEngineConfigRequest) SetOwnerAccount(v string) *GetLindormEngineConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormEngineConfigRequest) SetOwnerId(v int64) *GetLindormEngineConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormEngineConfigRequest) SetRegionId(v string) *GetLindormEngineConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormEngineConfigRequest) SetResourceOwnerAccount(v string) *GetLindormEngineConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormEngineConfigRequest) SetResourceOwnerId(v int64) *GetLindormEngineConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormEngineConfigRequest) SetSecurityToken(v string) *GetLindormEngineConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
