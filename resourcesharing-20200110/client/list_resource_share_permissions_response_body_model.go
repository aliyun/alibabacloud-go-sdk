// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceSharePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListResourceSharePermissionsResponseBody
	GetNextToken() *string
	SetPermissions(v []*ListResourceSharePermissionsResponseBodyPermissions) *ListResourceSharePermissionsResponseBody
	GetPermissions() []*ListResourceSharePermissionsResponseBodyPermissions
	SetRequestId(v string) *ListResourceSharePermissionsResponseBody
	GetRequestId() *string
}

type ListResourceSharePermissionsResponseBody struct {
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the permissions.
	Permissions []*ListResourceSharePermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2F23CFB6-A721-4E90-AC1E-0E30FA8B45DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceSharePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceSharePermissionsResponseBody) GetPermissions() []*ListResourceSharePermissionsResponseBodyPermissions {
	return s.Permissions
}

func (s *ListResourceSharePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceSharePermissionsResponseBody) SetNextToken(v string) *ListResourceSharePermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBody) SetPermissions(v []*ListResourceSharePermissionsResponseBodyPermissions) *ListResourceSharePermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListResourceSharePermissionsResponseBody) SetRequestId(v string) *ListResourceSharePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceSharePermissionsResponseBodyPermissions struct {
	// The creation time.
	//
	// example:
	//
	// 2020-12-07T07:39:01.818Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the permission is the default permission. Valid values:
	//
	// 	- false: The permission is not the default permission.
	//
	// 	- true: The permission is the default permission.
	//
	// example:
	//
	// true
	DefaultPermission *bool `json:"DefaultPermission,omitempty" xml:"DefaultPermission,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// 	- false: The version is not the default version.
	//
	// 	- true: The version is the default version.
	//
	// example:
	//
	// true
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// AliyunRSDefaultPermissionVSwitch
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	//
	// example:
	//
	// v1
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-12-07T07:39:01.818Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceSharePermissionsResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharePermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) GetDefaultPermission() *bool {
	return s.DefaultPermission
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) GetDefaultVersion() *bool {
	return s.DefaultVersion
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) GetPermissionName() *string {
	return s.PermissionName
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) GetPermissionVersion() *string {
	return s.PermissionVersion
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetCreateTime(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.CreateTime = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetDefaultPermission(v bool) *ListResourceSharePermissionsResponseBodyPermissions {
	s.DefaultPermission = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetDefaultVersion(v bool) *ListResourceSharePermissionsResponseBodyPermissions {
	s.DefaultVersion = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetPermissionName(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.PermissionName = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetPermissionVersion(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.PermissionVersion = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetResourceType(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetUpdateTime(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) Validate() error {
	return dara.Validate(s)
}
