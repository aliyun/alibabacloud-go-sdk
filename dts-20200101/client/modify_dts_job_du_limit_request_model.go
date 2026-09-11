// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobDuLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *ModifyDtsJobDuLimitRequest
	GetDtsJobId() *string
	SetDuLimit(v int64) *ModifyDtsJobDuLimitRequest
	GetDuLimit() *int64
	SetOwnerId(v string) *ModifyDtsJobDuLimitRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyDtsJobDuLimitRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDtsJobDuLimitRequest
	GetResourceGroupId() *string
}

type ModifyDtsJobDuLimitRequest struct {
	// The ID of the DTS migration, synchronization, or change tracking task.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsxxx
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The DU upper limit of the task.
	//
	// > The minimum value is **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DuLimit *int64  `json:"DuLimit,omitempty" xml:"DuLimit,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyDtsJobDuLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobDuLimitRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobDuLimitRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobDuLimitRequest) GetDuLimit() *int64 {
	return s.DuLimit
}

func (s *ModifyDtsJobDuLimitRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyDtsJobDuLimitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobDuLimitRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobDuLimitRequest) SetDtsJobId(v string) *ModifyDtsJobDuLimitRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobDuLimitRequest) SetDuLimit(v int64) *ModifyDtsJobDuLimitRequest {
	s.DuLimit = &v
	return s
}

func (s *ModifyDtsJobDuLimitRequest) SetOwnerId(v string) *ModifyDtsJobDuLimitRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDtsJobDuLimitRequest) SetRegionId(v string) *ModifyDtsJobDuLimitRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobDuLimitRequest) SetResourceGroupId(v string) *ModifyDtsJobDuLimitRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobDuLimitRequest) Validate() error {
	return dara.Validate(s)
}
