// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUserRolesByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *ListWorkspaceUserRolesByUserIdRequest
	GetUserId() *string
}

type ListWorkspaceUserRolesByUserIdRequest struct {
	// The ID of the Quick BI user. This is not your Alibaba Cloud account ID. You can call the [QueryUserInfoByAccount](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryUserInfoByAccount?spm=api-workbench.api_explorer.0.0.672f50daGq9ooV\\&params=%7B%7D\\&tab=DOC\\&sdkStyle=old\\&RegionId=cn-hangzhou) operation to obtain the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// afas-********asfasg
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListWorkspaceUserRolesByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUserRolesByUserIdRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUserRolesByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListWorkspaceUserRolesByUserIdRequest) SetUserId(v string) *ListWorkspaceUserRolesByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
