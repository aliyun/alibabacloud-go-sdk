// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCronJobPolicyServerlessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CancelCronJobPolicyServerlessRequest
	GetDBClusterId() *string
	SetJobId(v string) *CancelCronJobPolicyServerlessRequest
	GetJobId() *string
	SetOwnerAccount(v string) *CancelCronJobPolicyServerlessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelCronJobPolicyServerlessRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelCronJobPolicyServerlessRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelCronJobPolicyServerlessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelCronJobPolicyServerlessRequest
	GetResourceOwnerId() *int64
}

type CancelCronJobPolicyServerlessRequest struct {
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// b3e7b3d3-027d-4fcc-9f92-5c5f2363e141
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelCronJobPolicyServerlessRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCronJobPolicyServerlessRequest) GoString() string {
	return s.String()
}

func (s *CancelCronJobPolicyServerlessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelCronJobPolicyServerlessRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelCronJobPolicyServerlessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelCronJobPolicyServerlessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelCronJobPolicyServerlessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCronJobPolicyServerlessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelCronJobPolicyServerlessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelCronJobPolicyServerlessRequest) SetDBClusterId(v string) *CancelCronJobPolicyServerlessRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelCronJobPolicyServerlessRequest) SetJobId(v string) *CancelCronJobPolicyServerlessRequest {
	s.JobId = &v
	return s
}

func (s *CancelCronJobPolicyServerlessRequest) SetOwnerAccount(v string) *CancelCronJobPolicyServerlessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelCronJobPolicyServerlessRequest) SetOwnerId(v int64) *CancelCronJobPolicyServerlessRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelCronJobPolicyServerlessRequest) SetRegionId(v string) *CancelCronJobPolicyServerlessRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCronJobPolicyServerlessRequest) SetResourceOwnerAccount(v string) *CancelCronJobPolicyServerlessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelCronJobPolicyServerlessRequest) SetResourceOwnerId(v int64) *CancelCronJobPolicyServerlessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelCronJobPolicyServerlessRequest) Validate() error {
	return dara.Validate(s)
}
