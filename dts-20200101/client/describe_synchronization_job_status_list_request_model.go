// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobStatusListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSynchronizationJobStatusListRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSynchronizationJobStatusListRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSynchronizationJobStatusListRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeSynchronizationJobStatusListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSynchronizationJobStatusListRequest
	GetResourceGroupId() *string
	SetSynchronizationJobIdListJsonStr(v string) *DescribeSynchronizationJobStatusListRequest
	GetSynchronizationJobIdListJsonStr() *string
}

type DescribeSynchronizationJobStatusListRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data synchronization instances reside. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The IDs of the data synchronization instances. The value is a JSON array. You can call the [DescribeSynchronizationJobs](https://help.aliyun.com/document_detail/49454.html) operation to query the instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["dtsexjk1alb116****","dtskfq1149w254****"]
	SynchronizationJobIdListJsonStr *string `json:"SynchronizationJobIdListJsonStr,omitempty" xml:"SynchronizationJobIdListJsonStr,omitempty"`
}

func (s DescribeSynchronizationJobStatusListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSynchronizationJobStatusListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSynchronizationJobStatusListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSynchronizationJobStatusListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSynchronizationJobStatusListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSynchronizationJobStatusListRequest) GetSynchronizationJobIdListJsonStr() *string {
	return s.SynchronizationJobIdListJsonStr
}

func (s *DescribeSynchronizationJobStatusListRequest) SetAccountId(v string) *DescribeSynchronizationJobStatusListRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetClientToken(v string) *DescribeSynchronizationJobStatusListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetOwnerId(v string) *DescribeSynchronizationJobStatusListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetRegionId(v string) *DescribeSynchronizationJobStatusListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetResourceGroupId(v string) *DescribeSynchronizationJobStatusListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetSynchronizationJobIdListJsonStr(v string) *DescribeSynchronizationJobStatusListRequest {
	s.SynchronizationJobIdListJsonStr = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) Validate() error {
	return dara.Validate(s)
}
