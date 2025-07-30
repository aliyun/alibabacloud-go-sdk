// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationObjectModifyStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSynchronizationObjectModifyStatusRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSynchronizationObjectModifyStatusRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSynchronizationObjectModifyStatusRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeSynchronizationObjectModifyStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSynchronizationObjectModifyStatusRequest
	GetResourceGroupId() *string
	SetTaskId(v string) *DescribeSynchronizationObjectModifyStatusRequest
	GetTaskId() *string
}

type DescribeSynchronizationObjectModifyStatusRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data synchronization instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The task ID, which is returned after you call the [ModifySynchronizationObject](https://help.aliyun.com/document_detail/49451.html) operation to modify the objects to be synchronized.
	//
	// This parameter is required.
	//
	// example:
	//
	// k71r16fj13g****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetAccountId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetClientToken(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetOwnerId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetRegionId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetResourceGroupId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetTaskId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) Validate() error {
	return dara.Validate(s)
}
