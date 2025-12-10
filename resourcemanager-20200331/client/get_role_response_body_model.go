// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRoleResponseBody
	GetRequestId() *string
	SetRole(v *GetRoleResponseBodyRole) *GetRoleResponseBody
	GetRole() *GetRoleResponseBodyRole
}

type GetRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the role.
	Role *GetRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s GetRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoleResponseBody) GetRole() *GetRoleResponseBodyRole {
	return s.Role
}

func (s *GetRoleResponseBody) SetRequestId(v string) *GetRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoleResponseBody) SetRole(v *GetRoleResponseBodyRole) *GetRoleResponseBody {
	s.Role = v
	return s
}

func (s *GetRoleResponseBody) Validate() error {
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The document of the policy in which the identity that can assume the role is specified.
	//
	// example:
	//
	// { \\"Statement\\": [ { \\"Action\\": \\"sts:AssumeRole\\", \\"Effect\\": \\"Allow\\", \\"Principal\\": { \\"RAM\\": \\"acs:ram::12345678901234****:root\\" } } ], \\"Version\\": \\"1\\" }
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the role was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the role is a service-linked role.
	//
	// example:
	//
	// true
	IsServiceLinkedRole *bool `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	// The information of the most recent deletion task.
	LatestDeletionTask *GetRoleResponseBodyRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" type:"Struct"`
	// The maximum session duration of the role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the role.
	//
	// example:
	//
	// 90123456789****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the role after authorization.
	//
	// example:
	//
	// ECSAdmin@role.123456.onaliyunservice.com
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	// The time when the role was updated.
	//
	// example:
	//
	// 2016-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetRoleResponseBodyRole) String() string {
	return dara.Prettify(s)
}

func (s GetRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBodyRole) GetArn() *string {
	return s.Arn
}

func (s *GetRoleResponseBodyRole) GetAssumeRolePolicyDocument() *string {
	return s.AssumeRolePolicyDocument
}

func (s *GetRoleResponseBodyRole) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetRoleResponseBodyRole) GetDescription() *string {
	return s.Description
}

func (s *GetRoleResponseBodyRole) GetIsServiceLinkedRole() *bool {
	return s.IsServiceLinkedRole
}

func (s *GetRoleResponseBodyRole) GetLatestDeletionTask() *GetRoleResponseBodyRoleLatestDeletionTask {
	return s.LatestDeletionTask
}

func (s *GetRoleResponseBodyRole) GetMaxSessionDuration() *int64 {
	return s.MaxSessionDuration
}

func (s *GetRoleResponseBodyRole) GetRoleId() *string {
	return s.RoleId
}

func (s *GetRoleResponseBodyRole) GetRoleName() *string {
	return s.RoleName
}

func (s *GetRoleResponseBodyRole) GetRolePrincipalName() *string {
	return s.RolePrincipalName
}

func (s *GetRoleResponseBodyRole) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetRoleResponseBodyRole) SetArn(v string) *GetRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *GetRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetCreateDate(v string) *GetRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetDescription(v string) *GetRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetIsServiceLinkedRole(v bool) *GetRoleResponseBodyRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetLatestDeletionTask(v *GetRoleResponseBodyRoleLatestDeletionTask) *GetRoleResponseBodyRole {
	s.LatestDeletionTask = v
	return s
}

func (s *GetRoleResponseBodyRole) SetMaxSessionDuration(v int64) *GetRoleResponseBodyRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetRoleId(v string) *GetRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetRoleName(v string) *GetRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetRolePrincipalName(v string) *GetRoleResponseBodyRole {
	s.RolePrincipalName = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetUpdateDate(v string) *GetRoleResponseBodyRole {
	s.UpdateDate = &v
	return s
}

func (s *GetRoleResponseBodyRole) Validate() error {
	if s.LatestDeletionTask != nil {
		if err := s.LatestDeletionTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRoleResponseBodyRoleLatestDeletionTask struct {
	// The time when the deletion task was created.
	//
	// example:
	//
	// 2018-10-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the deletion task.
	//
	// example:
	//
	// ECSAdmin/cc61514b-26eb-4453-ab53-b142eb702a3d
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
}

func (s GetRoleResponseBodyRoleLatestDeletionTask) String() string {
	return dara.Prettify(s)
}

func (s GetRoleResponseBodyRoleLatestDeletionTask) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBodyRoleLatestDeletionTask) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetRoleResponseBodyRoleLatestDeletionTask) GetDeletionTaskId() *string {
	return s.DeletionTaskId
}

func (s *GetRoleResponseBodyRoleLatestDeletionTask) SetCreateDate(v string) *GetRoleResponseBodyRoleLatestDeletionTask {
	s.CreateDate = &v
	return s
}

func (s *GetRoleResponseBodyRoleLatestDeletionTask) SetDeletionTaskId(v string) *GetRoleResponseBodyRoleLatestDeletionTask {
	s.DeletionTaskId = &v
	return s
}

func (s *GetRoleResponseBodyRoleLatestDeletionTask) Validate() error {
	return dara.Validate(s)
}
