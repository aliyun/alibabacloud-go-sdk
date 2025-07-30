// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSynchronizationJobAlertRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSynchronizationJobAlertRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSynchronizationJobAlertRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeSynchronizationJobAlertRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSynchronizationJobAlertRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertRequest
	GetSynchronizationJobId() *string
}

type DescribeSynchronizationJobAlertRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data synchronization instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// >  Default value: **Forward**.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The ID of the data synchronization instance. You can call the **DescribeSynchronizationJobs*	- operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtskxz1170c10p****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DescribeSynchronizationJobAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSynchronizationJobAlertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSynchronizationJobAlertRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSynchronizationJobAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSynchronizationJobAlertRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSynchronizationJobAlertRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeSynchronizationJobAlertRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeSynchronizationJobAlertRequest) SetAccountId(v string) *DescribeSynchronizationJobAlertRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetClientToken(v string) *DescribeSynchronizationJobAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetOwnerId(v string) *DescribeSynchronizationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetRegionId(v string) *DescribeSynchronizationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetResourceGroupId(v string) *DescribeSynchronizationJobAlertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) Validate() error {
	return dara.Validate(s)
}
