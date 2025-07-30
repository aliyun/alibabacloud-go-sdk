// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *ModifyDtsJobPasswordRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ModifyDtsJobPasswordRequest
	GetDtsJobId() *string
	SetEndpoint(v string) *ModifyDtsJobPasswordRequest
	GetEndpoint() *string
	SetPassword(v string) *ModifyDtsJobPasswordRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyDtsJobPasswordRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDtsJobPasswordRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *ModifyDtsJobPasswordRequest
	GetSynchronizationDirection() *string
	SetUserName(v string) *ModifyDtsJobPasswordRequest
	GetUserName() *string
	SetZeroEtlJob(v bool) *ModifyDtsJobPasswordRequest
	GetZeroEtlJob() *bool
}

type ModifyDtsJobPasswordRequest struct {
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// >  You can call the [DescribeMigrationJobs](https://help.aliyun.com/document_detail/208139.html), [DescribeSubscriptionInstances](https://help.aliyun.com/document_detail/49442.html), or [DescribeSynchronizationJobs](https://help.aliyun.com/document_detail/49454.html) operation to query the instance ID
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the DTS task. The DTS task can be a data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The database to which the password belongs. Valid values:
	//
	// 	- **src**: source database.
	//
	// 	- **dest**: destination database.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// src
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The new password.
	//
	// >  This parameter is required and cannot be set to a value that is the same as the current password.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the region where the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// Synchronization direction, with values:
	//
	// - **Forward*	- (default): Forward. - **Reverse**: Reverse.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The account of the source or destination database.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Whether it is a seamless integration (Zero-ETL) task, the value can be:
	//
	// - **false**: No. - **true**: Yes.
	//
	// example:
	//
	// false
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s ModifyDtsJobPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobPasswordRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ModifyDtsJobPasswordRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobPasswordRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModifyDtsJobPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyDtsJobPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobPasswordRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobPasswordRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ModifyDtsJobPasswordRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyDtsJobPasswordRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ModifyDtsJobPasswordRequest) SetDtsInstanceId(v string) *ModifyDtsJobPasswordRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetDtsJobId(v string) *ModifyDtsJobPasswordRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetEndpoint(v string) *ModifyDtsJobPasswordRequest {
	s.Endpoint = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetPassword(v string) *ModifyDtsJobPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetRegionId(v string) *ModifyDtsJobPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetResourceGroupId(v string) *ModifyDtsJobPasswordRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetSynchronizationDirection(v string) *ModifyDtsJobPasswordRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetUserName(v string) *ModifyDtsJobPasswordRequest {
	s.UserName = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetZeroEtlJob(v bool) *ModifyDtsJobPasswordRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) Validate() error {
	return dara.Validate(s)
}
