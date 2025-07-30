// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteSynchronizationJobRequest
	GetAccountId() *string
	SetOwnerId(v string) *DeleteSynchronizationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *DeleteSynchronizationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteSynchronizationJobRequest
	GetResourceGroupId() *string
	SetSynchronizationJobId(v string) *DeleteSynchronizationJobRequest
	GetSynchronizationJobId() *string
}

type DeleteSynchronizationJobRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data synchronization instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组ID。
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the data synchronization instance. You can call the DescribeSynchronizationJobs operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtshn6107ve264****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DeleteSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteSynchronizationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteSynchronizationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSynchronizationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteSynchronizationJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DeleteSynchronizationJobRequest) SetAccountId(v string) *DeleteSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetOwnerId(v string) *DeleteSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetRegionId(v string) *DeleteSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetResourceGroupId(v string) *DeleteSynchronizationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetSynchronizationJobId(v string) *DeleteSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
