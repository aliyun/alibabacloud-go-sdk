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
	// The instance ID of the data migration, synchronization, or subscribe instance.
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration, synchronization, or change tracking task.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The database to which the account belongs. Valid values:
	//
	// - **src**: the source database.
	//
	// - **dest**: the destination database.
	//
	// > This parameter is required.
	//
	// example:
	//
	// src
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The new password for the database account.
	//
	// > This parameter is required and must be different from the current password.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - Default value: **Forward**.
	//
	// - This parameter is required only when the synchronization topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The username of the database account to modify.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Specifies whether the node is a seamless integration (Zero-ETL) node. Valid values:
	//
	// - **true**
	//
	// - **false**.
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
