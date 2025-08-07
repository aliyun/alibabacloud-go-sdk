// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CancelActiveOperationTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelActiveOperationTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelActiveOperationTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CancelActiveOperationTasksRequest
	GetSecurityToken() *string
	SetTaskIds(v string) *CancelActiveOperationTasksRequest
	GetTaskIds() *string
}

type CancelActiveOperationTasksRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >
	//
	// 	- You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the region information about all clusters within a specified account.
	//
	// 	- If you do not specify this parameter, scheduled tasks on your clusters that are deployed in all regions are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The IDs of O\\&M events that are canceled at a time. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2355,2352
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s CancelActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelActiveOperationTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelActiveOperationTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelActiveOperationTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelActiveOperationTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CancelActiveOperationTasksRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *CancelActiveOperationTasksRequest) SetOwnerAccount(v string) *CancelActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetOwnerId(v int64) *CancelActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetRegionId(v string) *CancelActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *CancelActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetResourceOwnerId(v int64) *CancelActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetSecurityToken(v string) *CancelActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetTaskIds(v string) *CancelActiveOperationTasksRequest {
	s.TaskIds = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
