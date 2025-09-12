// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceEngineListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormInstanceEngineListRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormInstanceEngineListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormInstanceEngineListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetLindormInstanceEngineListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetLindormInstanceEngineListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormInstanceEngineListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormInstanceEngineListRequest
	GetSecurityToken() *string
}

type GetLindormInstanceEngineListRequest struct {
	// Instance ID, which can be obtained by calling the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormInstanceEngineListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceEngineListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormInstanceEngineListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormInstanceEngineListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormInstanceEngineListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormInstanceEngineListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormInstanceEngineListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormInstanceEngineListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormInstanceEngineListRequest) SetInstanceId(v string) *GetLindormInstanceEngineListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetOwnerAccount(v string) *GetLindormInstanceEngineListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetOwnerId(v int64) *GetLindormInstanceEngineListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetRegionId(v string) *GetLindormInstanceEngineListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceEngineListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetResourceOwnerId(v int64) *GetLindormInstanceEngineListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetSecurityToken(v string) *GetLindormInstanceEngineListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) Validate() error {
	return dara.Validate(s)
}
