// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterializedViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyMaterializedViewRequest
	GetDBClusterId() *string
	SetDbName(v string) *ModifyMaterializedViewRequest
	GetDbName() *string
	SetEnableDelayAlert(v bool) *ModifyMaterializedViewRequest
	GetEnableDelayAlert() *bool
	SetEnableFailureAlert(v bool) *ModifyMaterializedViewRequest
	GetEnableFailureAlert() *bool
	SetGroupName(v string) *ModifyMaterializedViewRequest
	GetGroupName() *string
	SetLatencyTolerance(v int32) *ModifyMaterializedViewRequest
	GetLatencyTolerance() *int32
	SetOwnerAccount(v string) *ModifyMaterializedViewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyMaterializedViewRequest
	GetOwnerId() *int64
	SetQueryWrite(v bool) *ModifyMaterializedViewRequest
	GetQueryWrite() *bool
	SetRegionId(v string) *ModifyMaterializedViewRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyMaterializedViewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyMaterializedViewRequest
	GetResourceOwnerId() *int64
	SetViewName(v string) *ModifyMaterializedViewRequest
	GetViewName() *string
}

type ModifyMaterializedViewRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database where the materialized view resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// myDB
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Enable the refresh delay alert. Valid values:
	//
	// 	- true: Enables alert.
	//
	// 	- false: Disables alert.
	//
	// example:
	//
	// false
	EnableDelayAlert *bool `json:"EnableDelayAlert,omitempty" xml:"EnableDelayAlert,omitempty"`
	// Specifies whether to send alerts when the refresh task fails. Valid values:
	//
	// 	- true: Send alerts.
	//
	// 	- false: Alerts disabled.
	//
	// example:
	//
	// false
	EnableFailureAlert *bool `json:"EnableFailureAlert,omitempty" xml:"EnableFailureAlert,omitempty"`
	// The name of the resource group to which the materialized view is bound.
	//
	// example:
	//
	// res_1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Refresh delay tolerance (in minutes).
	//
	// example:
	//
	// 2
	LatencyTolerance *int32  `json:"LatencyTolerance,omitempty" xml:"LatencyTolerance,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable query rewrite. Valid values:
	//
	// 	- true: Enables query rewrite.
	//
	// 	- false: Disables query rewrite.
	//
	// example:
	//
	// true
	QueryWrite *bool `json:"QueryWrite,omitempty" xml:"QueryWrite,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the materialized view.
	//
	// This parameter is required.
	//
	// example:
	//
	// mv_1
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s ModifyMaterializedViewRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterializedViewRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaterializedViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyMaterializedViewRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyMaterializedViewRequest) GetEnableDelayAlert() *bool {
	return s.EnableDelayAlert
}

func (s *ModifyMaterializedViewRequest) GetEnableFailureAlert() *bool {
	return s.EnableFailureAlert
}

func (s *ModifyMaterializedViewRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyMaterializedViewRequest) GetLatencyTolerance() *int32 {
	return s.LatencyTolerance
}

func (s *ModifyMaterializedViewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyMaterializedViewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyMaterializedViewRequest) GetQueryWrite() *bool {
	return s.QueryWrite
}

func (s *ModifyMaterializedViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMaterializedViewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyMaterializedViewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyMaterializedViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *ModifyMaterializedViewRequest) SetDBClusterId(v string) *ModifyMaterializedViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetDbName(v string) *ModifyMaterializedViewRequest {
	s.DbName = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetEnableDelayAlert(v bool) *ModifyMaterializedViewRequest {
	s.EnableDelayAlert = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetEnableFailureAlert(v bool) *ModifyMaterializedViewRequest {
	s.EnableFailureAlert = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetGroupName(v string) *ModifyMaterializedViewRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetLatencyTolerance(v int32) *ModifyMaterializedViewRequest {
	s.LatencyTolerance = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetOwnerAccount(v string) *ModifyMaterializedViewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetOwnerId(v int64) *ModifyMaterializedViewRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetQueryWrite(v bool) *ModifyMaterializedViewRequest {
	s.QueryWrite = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetRegionId(v string) *ModifyMaterializedViewRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetResourceOwnerAccount(v string) *ModifyMaterializedViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetResourceOwnerId(v int64) *ModifyMaterializedViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMaterializedViewRequest) SetViewName(v string) *ModifyMaterializedViewRequest {
	s.ViewName = &v
	return s
}

func (s *ModifyMaterializedViewRequest) Validate() error {
	return dara.Validate(s)
}
