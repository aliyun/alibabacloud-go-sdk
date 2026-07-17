// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeEngineJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeGroup(v string) *ListComputeEngineJobRequest
	GetComputeGroup() *string
	SetEndTime(v int64) *ListComputeEngineJobRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListComputeEngineJobRequest
	GetInstanceId() *string
	SetJobId(v string) *ListComputeEngineJobRequest
	GetJobId() *string
	SetJobName(v string) *ListComputeEngineJobRequest
	GetJobName() *string
	SetOwnerAccount(v string) *ListComputeEngineJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListComputeEngineJobRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListComputeEngineJobRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComputeEngineJobRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListComputeEngineJobRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListComputeEngineJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListComputeEngineJobRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ListComputeEngineJobRequest
	GetSecurityToken() *string
	SetStartTime(v int64) *ListComputeEngineJobRequest
	GetStartTime() *int64
	SetState(v string) *ListComputeEngineJobRequest
	GetState() *string
}

type ListComputeEngineJobRequest struct {
	ComputeGroup *string `json:"ComputeGroup,omitempty" xml:"ComputeGroup,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListComputeEngineJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeEngineJobRequest) GoString() string {
	return s.String()
}

func (s *ListComputeEngineJobRequest) GetComputeGroup() *string {
	return s.ComputeGroup
}

func (s *ListComputeEngineJobRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListComputeEngineJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeEngineJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListComputeEngineJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListComputeEngineJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListComputeEngineJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListComputeEngineJobRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComputeEngineJobRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeEngineJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeEngineJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListComputeEngineJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListComputeEngineJobRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListComputeEngineJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListComputeEngineJobRequest) GetState() *string {
	return s.State
}

func (s *ListComputeEngineJobRequest) SetComputeGroup(v string) *ListComputeEngineJobRequest {
	s.ComputeGroup = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetEndTime(v int64) *ListComputeEngineJobRequest {
	s.EndTime = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetInstanceId(v string) *ListComputeEngineJobRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetJobId(v string) *ListComputeEngineJobRequest {
	s.JobId = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetJobName(v string) *ListComputeEngineJobRequest {
	s.JobName = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetOwnerAccount(v string) *ListComputeEngineJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetOwnerId(v int64) *ListComputeEngineJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetPageNumber(v int32) *ListComputeEngineJobRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetPageSize(v int32) *ListComputeEngineJobRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetRegionId(v string) *ListComputeEngineJobRequest {
	s.RegionId = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetResourceOwnerAccount(v string) *ListComputeEngineJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetResourceOwnerId(v int64) *ListComputeEngineJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetSecurityToken(v string) *ListComputeEngineJobRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetStartTime(v int64) *ListComputeEngineJobRequest {
	s.StartTime = &v
	return s
}

func (s *ListComputeEngineJobRequest) SetState(v string) *ListComputeEngineJobRequest {
	s.State = &v
	return s
}

func (s *ListComputeEngineJobRequest) Validate() error {
	return dara.Validate(s)
}
