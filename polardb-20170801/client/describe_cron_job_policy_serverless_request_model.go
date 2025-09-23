// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCronJobPolicyServerlessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeCronJobPolicyServerlessRequest
	GetDBClusterId() *string
	SetJobId(v string) *DescribeCronJobPolicyServerlessRequest
	GetJobId() *string
	SetOwnerAccount(v string) *DescribeCronJobPolicyServerlessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCronJobPolicyServerlessRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCronJobPolicyServerlessRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCronJobPolicyServerlessRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCronJobPolicyServerlessRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCronJobPolicyServerlessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCronJobPolicyServerlessRequest
	GetResourceOwnerId() *int64
}

type DescribeCronJobPolicyServerlessRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 8006e51c-dab3-4602-bc69-4f728002c6ce
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCronJobPolicyServerlessRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCronJobPolicyServerlessRequest) GoString() string {
	return s.String()
}

func (s *DescribeCronJobPolicyServerlessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeCronJobPolicyServerlessRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeCronJobPolicyServerlessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCronJobPolicyServerlessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCronJobPolicyServerlessRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCronJobPolicyServerlessRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCronJobPolicyServerlessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCronJobPolicyServerlessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCronJobPolicyServerlessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCronJobPolicyServerlessRequest) SetDBClusterId(v string) *DescribeCronJobPolicyServerlessRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetJobId(v string) *DescribeCronJobPolicyServerlessRequest {
	s.JobId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetOwnerAccount(v string) *DescribeCronJobPolicyServerlessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetOwnerId(v int64) *DescribeCronJobPolicyServerlessRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetPageNumber(v int32) *DescribeCronJobPolicyServerlessRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetPageSize(v int32) *DescribeCronJobPolicyServerlessRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetRegionId(v string) *DescribeCronJobPolicyServerlessRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetResourceOwnerAccount(v string) *DescribeCronJobPolicyServerlessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) SetResourceOwnerId(v int64) *DescribeCronJobPolicyServerlessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessRequest) Validate() error {
	return dara.Validate(s)
}
