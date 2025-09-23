// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListPermissionVersionsResponseBody
	GetNextToken() *string
	SetPermissions(v []*ListPermissionVersionsResponseBodyPermissions) *ListPermissionVersionsResponseBody
	GetPermissions() []*ListPermissionVersionsResponseBodyPermissions
	SetRequestId(v string) *ListPermissionVersionsResponseBody
	GetRequestId() *string
}

type ListPermissionVersionsResponseBody struct {
	// The token that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the permission.
	Permissions []*ListPermissionVersionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04677DCA-7C33-464B-8811-1B1DA3C3D197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPermissionVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionVersionsResponseBody) GetPermissions() []*ListPermissionVersionsResponseBodyPermissions {
	return s.Permissions
}

func (s *ListPermissionVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionVersionsResponseBody) SetNextToken(v string) *ListPermissionVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPermissionVersionsResponseBody) SetPermissions(v []*ListPermissionVersionsResponseBodyPermissions) *ListPermissionVersionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionVersionsResponseBody) SetRequestId(v string) *ListPermissionVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPermissionVersionsResponseBodyPermissions struct {
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

func (s ListPermissionVersionsResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionVersionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsResponseBodyPermissions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPermissionVersionsResponseBodyPermissions) GetDefaultPermission() *bool {
	return s.DefaultPermission
}

func (s *ListPermissionVersionsResponseBodyPermissions) GetDefaultVersion() *bool {
	return s.DefaultVersion
}

func (s *ListPermissionVersionsResponseBodyPermissions) GetPermissionName() *string {
	return s.PermissionName
}

func (s *ListPermissionVersionsResponseBodyPermissions) GetPermissionVersion() *string {
	return s.PermissionVersion
}

func (s *ListPermissionVersionsResponseBodyPermissions) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPermissionVersionsResponseBodyPermissions) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetCreateTime(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.CreateTime = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetDefaultPermission(v bool) *ListPermissionVersionsResponseBodyPermissions {
	s.DefaultPermission = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetDefaultVersion(v bool) *ListPermissionVersionsResponseBodyPermissions {
	s.DefaultVersion = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetPermissionName(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.PermissionName = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetPermissionVersion(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.PermissionVersion = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetResourceType(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetUpdateTime(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.UpdateTime = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) Validate() error {
	return dara.Validate(s)
}
