// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEngineJobLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetComputeEngineJobLogRequest
	GetInstanceId() *string
	SetJobId(v string) *GetComputeEngineJobLogRequest
	GetJobId() *string
	SetOwnerAccount(v string) *GetComputeEngineJobLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetComputeEngineJobLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *GetComputeEngineJobLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetComputeEngineJobLogRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetComputeEngineJobLogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetComputeEngineJobLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetComputeEngineJobLogRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetComputeEngineJobLogRequest
	GetSecurityToken() *string
}

type GetComputeEngineJobLogRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetComputeEngineJobLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEngineJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetComputeEngineJobLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetComputeEngineJobLogRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetComputeEngineJobLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetComputeEngineJobLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetComputeEngineJobLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetComputeEngineJobLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetComputeEngineJobLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeEngineJobLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetComputeEngineJobLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetComputeEngineJobLogRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetComputeEngineJobLogRequest) SetInstanceId(v string) *GetComputeEngineJobLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetJobId(v string) *GetComputeEngineJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetOwnerAccount(v string) *GetComputeEngineJobLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetOwnerId(v int64) *GetComputeEngineJobLogRequest {
	s.OwnerId = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetPageNumber(v int32) *GetComputeEngineJobLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetPageSize(v int32) *GetComputeEngineJobLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetRegionId(v string) *GetComputeEngineJobLogRequest {
	s.RegionId = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetResourceOwnerAccount(v string) *GetComputeEngineJobLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetResourceOwnerId(v int64) *GetComputeEngineJobLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) SetSecurityToken(v string) *GetComputeEngineJobLogRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetComputeEngineJobLogRequest) Validate() error {
	return dara.Validate(s)
}
