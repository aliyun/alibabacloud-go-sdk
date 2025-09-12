// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceEngineListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2InstanceEngineListRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormV2InstanceEngineListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormV2InstanceEngineListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetLindormV2InstanceEngineListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetLindormV2InstanceEngineListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormV2InstanceEngineListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormV2InstanceEngineListRequest
	GetSecurityToken() *string
}

type GetLindormV2InstanceEngineListRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2InstanceEngineListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceEngineListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceEngineListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormV2InstanceEngineListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormV2InstanceEngineListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormV2InstanceEngineListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormV2InstanceEngineListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormV2InstanceEngineListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormV2InstanceEngineListRequest) SetInstanceId(v string) *GetLindormV2InstanceEngineListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetOwnerAccount(v string) *GetLindormV2InstanceEngineListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetOwnerId(v int64) *GetLindormV2InstanceEngineListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetRegionId(v string) *GetLindormV2InstanceEngineListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetResourceOwnerAccount(v string) *GetLindormV2InstanceEngineListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetResourceOwnerId(v int64) *GetLindormV2InstanceEngineListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetSecurityToken(v string) *GetLindormV2InstanceEngineListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) Validate() error {
	return dara.Validate(s)
}
