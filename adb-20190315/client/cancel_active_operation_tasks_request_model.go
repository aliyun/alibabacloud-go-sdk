// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *CancelActiveOperationTasksRequest
	GetIds() *string
	SetOwnerAccount(v string) *CancelActiveOperationTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelActiveOperationTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelActiveOperationTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CancelActiveOperationTasksRequest
	GetSecurityToken() *string
}

type CancelActiveOperationTasksRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// am-xxxxxxxx
	Ids          *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CancelActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksRequest) GetIds() *string {
	return s.Ids
}

func (s *CancelActiveOperationTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelActiveOperationTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelActiveOperationTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelActiveOperationTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelActiveOperationTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CancelActiveOperationTasksRequest) SetIds(v string) *CancelActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetOwnerAccount(v string) *CancelActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetOwnerId(v int64) *CancelActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetRegionId(v string) *CancelActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *CancelActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetResourceOwnerId(v int64) *CancelActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetSecurityToken(v string) *CancelActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
