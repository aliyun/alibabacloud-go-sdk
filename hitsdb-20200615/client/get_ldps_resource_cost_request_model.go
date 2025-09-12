// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsResourceCostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetLdpsResourceCostRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetLdpsResourceCostRequest
	GetInstanceId() *string
	SetJobId(v string) *GetLdpsResourceCostRequest
	GetJobId() *string
	SetOwnerAccount(v string) *GetLdpsResourceCostRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLdpsResourceCostRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetLdpsResourceCostRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetLdpsResourceCostRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLdpsResourceCostRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLdpsResourceCostRequest
	GetSecurityToken() *string
	SetStartTime(v int64) *GetLdpsResourceCostRequest
	GetStartTime() *int64
}

type GetLdpsResourceCostRequest struct {
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetLdpsResourceCostRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsResourceCostRequest) GoString() string {
	return s.String()
}

func (s *GetLdpsResourceCostRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetLdpsResourceCostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLdpsResourceCostRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetLdpsResourceCostRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLdpsResourceCostRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLdpsResourceCostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLdpsResourceCostRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLdpsResourceCostRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLdpsResourceCostRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLdpsResourceCostRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetLdpsResourceCostRequest) SetEndTime(v int64) *GetLdpsResourceCostRequest {
	s.EndTime = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetInstanceId(v string) *GetLdpsResourceCostRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetJobId(v string) *GetLdpsResourceCostRequest {
	s.JobId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetOwnerAccount(v string) *GetLdpsResourceCostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetOwnerId(v int64) *GetLdpsResourceCostRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetRegionId(v string) *GetLdpsResourceCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetResourceOwnerAccount(v string) *GetLdpsResourceCostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetResourceOwnerId(v int64) *GetLdpsResourceCostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetSecurityToken(v string) *GetLdpsResourceCostRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetStartTime(v int64) *GetLdpsResourceCostRequest {
	s.StartTime = &v
	return s
}

func (s *GetLdpsResourceCostRequest) Validate() error {
	return dara.Validate(s)
}
