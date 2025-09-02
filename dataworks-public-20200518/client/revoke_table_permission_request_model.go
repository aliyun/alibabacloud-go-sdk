// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTablePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v string) *RevokeTablePermissionRequest
	GetActions() *string
	SetMaxComputeProjectName(v string) *RevokeTablePermissionRequest
	GetMaxComputeProjectName() *string
	SetRevokeUserId(v string) *RevokeTablePermissionRequest
	GetRevokeUserId() *string
	SetRevokeUserName(v string) *RevokeTablePermissionRequest
	GetRevokeUserName() *string
	SetTableName(v string) *RevokeTablePermissionRequest
	GetTableName() *string
	SetWorkspaceId(v int64) *RevokeTablePermissionRequest
	GetWorkspaceId() *int64
}

type RevokeTablePermissionRequest struct {
	// The permissions that you want to revoke. Separate multiple permissions with commas (,). You can revoke only the SELECT, DESCRIBE, and DOWNLOAD permissions on MaxCompute tables.
	//
	// This parameter is required.
	//
	// example:
	//
	// Select,Describe
	Actions *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The name of the MaxCompute project to which the table belongs. You can log on to the DataWorks console and go to the SettingCenter page to obtain the name of the MaxCompute project that you associate with the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// aMaxcomputeProjectName
	MaxComputeProjectName *string `json:"MaxComputeProjectName,omitempty" xml:"MaxComputeProjectName,omitempty"`
	// The ID of the Alibaba Cloud account from which you want to revoke permissions. You can log on to the DataWorks console and go to the Security Settings page to obtain the ID. You must specify either this parameter or the RevokeUserName parameter. If you specify both this parameter and the RevokeUserName parameter and the parameter values are different, the value of this parameter prevails.
	//
	// example:
	//
	// 267842600408993176
	RevokeUserId *string `json:"RevokeUserId,omitempty" xml:"RevokeUserId,omitempty"`
	// The Alibaba Cloud account from which you want to revoke permissions. Specify this parameter in the format that is the same as the format of the account used to access the MaxCompute project.
	//
	// 	- If the account is an Alibaba Cloud account, the value is in the ALIYUN$+Account name format.
	//
	// 	- If the account is a RAM user, the value is in the RAM$+Account name format.
	//
	// You must specify either this parameter or the RevokeUserId parameter. If you specify both this parameter and the RevokeUserId parameter and the parameter values are different, the value of the RevokeUserId parameter prevails.
	//
	// example:
	//
	// RAM$dataworks_3h1_1:stsramuser
	RevokeUserName *string `json:"RevokeUserName,omitempty" xml:"RevokeUserName,omitempty"`
	// The name of the MaxCompute table. You can call the [SearchMetaTables](https://help.aliyun.com/document_detail/173919.html) operation to query the name of the MaxCompute table.
	//
	// This parameter is required.
	//
	// example:
	//
	// aTableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the DataWorks workspace with which the MaxCompute project is associated. You can log on to the DataWorks console and go to the Workspace page to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RevokeTablePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeTablePermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeTablePermissionRequest) GetActions() *string {
	return s.Actions
}

func (s *RevokeTablePermissionRequest) GetMaxComputeProjectName() *string {
	return s.MaxComputeProjectName
}

func (s *RevokeTablePermissionRequest) GetRevokeUserId() *string {
	return s.RevokeUserId
}

func (s *RevokeTablePermissionRequest) GetRevokeUserName() *string {
	return s.RevokeUserName
}

func (s *RevokeTablePermissionRequest) GetTableName() *string {
	return s.TableName
}

func (s *RevokeTablePermissionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *RevokeTablePermissionRequest) SetActions(v string) *RevokeTablePermissionRequest {
	s.Actions = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetMaxComputeProjectName(v string) *RevokeTablePermissionRequest {
	s.MaxComputeProjectName = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetRevokeUserId(v string) *RevokeTablePermissionRequest {
	s.RevokeUserId = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetRevokeUserName(v string) *RevokeTablePermissionRequest {
	s.RevokeUserName = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetTableName(v string) *RevokeTablePermissionRequest {
	s.TableName = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetWorkspaceId(v int64) *RevokeTablePermissionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RevokeTablePermissionRequest) Validate() error {
	return dara.Validate(s)
}
