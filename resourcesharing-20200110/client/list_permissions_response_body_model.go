// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListPermissionsResponseBody
	GetNextToken() *string
	SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody
	GetPermissions() []*ListPermissionsResponseBodyPermissions
	SetRequestId(v string) *ListPermissionsResponseBody
	GetRequestId() *string
}

type ListPermissionsResponseBody struct {
	// The token that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the permission.
	Permissions []*ListPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04677DCA-7C33-464B-8811-1B1DA3C3D197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionsResponseBody) GetPermissions() []*ListPermissionsResponseBodyPermissions {
	return s.Permissions
}

func (s *ListPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionsResponseBody) SetNextToken(v string) *ListPermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionsResponseBody) SetRequestId(v string) *ListPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionsResponseBody) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPermissionsResponseBodyPermissions struct {
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

func (s ListPermissionsResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPermissionsResponseBodyPermissions) GetDefaultPermission() *bool {
	return s.DefaultPermission
}

func (s *ListPermissionsResponseBodyPermissions) GetDefaultVersion() *bool {
	return s.DefaultVersion
}

func (s *ListPermissionsResponseBodyPermissions) GetPermissionName() *string {
	return s.PermissionName
}

func (s *ListPermissionsResponseBodyPermissions) GetPermissionVersion() *string {
	return s.PermissionVersion
}

func (s *ListPermissionsResponseBodyPermissions) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPermissionsResponseBodyPermissions) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPermissionsResponseBodyPermissions) SetCreateTime(v string) *ListPermissionsResponseBodyPermissions {
	s.CreateTime = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetDefaultPermission(v bool) *ListPermissionsResponseBodyPermissions {
	s.DefaultPermission = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetDefaultVersion(v bool) *ListPermissionsResponseBodyPermissions {
	s.DefaultVersion = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionName(v string) *ListPermissionsResponseBodyPermissions {
	s.PermissionName = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionVersion(v string) *ListPermissionsResponseBodyPermissions {
	s.PermissionVersion = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetResourceType(v string) *ListPermissionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetUpdateTime(v string) *ListPermissionsResponseBodyPermissions {
	s.UpdateTime = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) Validate() error {
	return dara.Validate(s)
}
